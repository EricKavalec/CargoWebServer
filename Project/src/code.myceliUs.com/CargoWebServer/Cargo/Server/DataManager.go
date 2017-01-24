package Server

import (
	"errors"
	"log"
	"os"
	"sync"

	"code.myceliUs.com/CargoWebServer/Cargo/Config/CargoConfig"
	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
)

////////////////////////////////////////////////////////////////////////////////
//              			Store id's
////////////////////////////////////////////////////////////////////////////////
const (
	// Use to store computer, group and user info

	// The bpmn informations...
	BPMN20DB = "BPMN20"
	BPMNDIDB = "BPMN20"
	DCDB     = "BPMN20"
	DIDB     = "BPMN20"

	// The runtime database...
	BPMS_RuntimeDB = "BPMS_Runtime"

	// The persistence db
	CargoEntitiesDB = "CargoEntities"

	// The configuration db
	CargoConfigDB = "CargoConfig"
)

/**
 * Data manager can be use to retreive data inside a data store, like Sql server
 * from file like xml file or any other source...
 */
type DataManager struct {
	/** This contain connection to know dataStore **/
	m_dataStores map[string]DataStore
	/**
	 * Use to protected the entitiesMap access...
	 */
	sync.RWMutex
}

var dataManager *DataManager

func (this *Server) GetDataManager() *DataManager {
	if dataManager == nil {
		dataManager = newDataManager()
	}
	return dataManager
}

/**
 * This is the accessing function to dataStore...
 */
func newDataManager() *DataManager {

	// Register dynamic type here...
	dataManager := new(DataManager)
	dataManager.m_dataStores = make(map[string]DataStore)

	/** Now I will initialyse data store one by one... **/
	defaultStoreConfigurations := GetServer().GetConfigurationManager().GetDefaultDataStoreConfigurations()

	for i := 0; i < len(defaultStoreConfigurations); i++ {
		store, err := NewDataStore(defaultStoreConfigurations[i])
		if err != nil {
			log.Fatal(err)
		}
		dataManager.m_dataStores[store.GetId()] = store
	}

	/** Return the data manager pointer... **/
	return dataManager
}

////////////////////////////////////////////////////////////////////////////////
//                   		DataManager function
////////////////////////////////////////////////////////////////////////////////
func (this *DataManager) Initialize() {
	// Here I will get the datastore configuration...
	storeConfigurations := GetServer().GetConfigurationManager().GetDataStoreConfigurations()

	for i := 0; i < len(storeConfigurations); i++ {
		store, err := NewDataStore(storeConfigurations[i])
		if err != nil {
			log.Fatal(err)
		}

		this.m_dataStores[store.GetId()] = store
	}
}

func (this *DataManager) Start() {
	log.Println("--> Start DataManager")
}

func (this *DataManager) Stop() {
	log.Println("--> Stop DataManager")
	this.close()
}

////////////////////////////////////////////////////////////////////////////////
// private function
////////////////////////////////////////////////////////////////////////////////

/**
 * Access a store with here given name...
 */
func (this *DataManager) getDataStore(name string) DataStore {
	this.Lock()
	defer this.Unlock()

	store := this.m_dataStores[name]
	if store == nil {
		log.Println("Store with name ", name, " is not configure!!!")
	}
	return store
}

/**
 * Remove a dataStore from the map
 */
func (this *DataManager) removeDataStore(name string) {
	this.Lock()
	defer this.Unlock()

	// Close the connection.
	this.m_dataStores[name].Close()

	// Delete the reference from the database.
	delete(this.m_dataStores, name)
}

/**
 * Execute a query that read information from the store and
 * return the result and an array of interface...
 */
func (this *DataManager) readData(storeName string, query string, fieldsType []interface{}, params []interface{}) ([][]interface{}, error) {

	store := this.getDataStore(storeName)
	if store == nil {
		return nil, errors.New("The datastore '" + storeName + "' does not exist.")
	}
	data, err := store.Read(query, fieldsType, params)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
	}
	return data, err
}

/**
 * Execute a query that create a new data. The data contains the new
 * value to insert in the DB.
 */
func (this *DataManager) createData(storeName string, query string, d []interface{}) (lastId interface{}, err error) {
	store := this.getDataStore(storeName)
	if store == nil {
		return nil, errors.New("Data store '" + storeName + " does not exist.")
	}
	lastId, err = store.Create(query, d)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
	}
	return
}

func (this *DataManager) deleteData(storeName string, query string, params []interface{}) (err error) {
	store := this.getDataStore(storeName)
	if store == nil {
		return errors.New("Data store " + storeName + " does not exist.")
	}

	err = store.Delete(query, params)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
	}
	return
}

func (this *DataManager) updateData(storeName string, query string, fields []interface{}, params []interface{}) (err error) {
	store := this.getDataStore(storeName)
	if store == nil {
		return errors.New("Data store " + storeName + " does not exist.")
	}
	err = store.Update(query, fields, params)
	if err != nil {
		err = errors.New("Query '" + query + "' failed with error '" + err.Error() + "'.")
	}
	return
}

func (this *DataManager) createDataStore(storeId string, storeType CargoConfig.DataStoreType, storeVendor CargoConfig.DataStoreVendor) (DataStore, *CargoEntities.Error) {

	if !Utility.IsValidVariableName(storeId) {
		cargoError := NewError(Utility.FileLine(), INVALID_VARIABLE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' is not valid."))
		return nil, cargoError
	}

	if this.getDataStore(storeId) != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ALREADY_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' already exists."))
		return nil, cargoError
	}

	// Create the new store here.
	var storeConfig CargoConfig.DataStoreConfiguration
	storeConfig.M_id = storeId
	storeConfig.M_dataStoreVendor = storeVendor
	storeConfig.M_dataStoreType = storeType
	storeConfig.NeedSave = true

	// Create the store here.
	store, err := NewDataStore(storeConfig)
	if err == nil {
		// Append the new dataStore configuration.
		GetServer().GetConfigurationManager().m_activeConfigurations.SetDataStoreConfigs(&storeConfig)
		// Save it.
		GetServer().GetConfigurationManager().m_configurationEntity.SaveEntity()
		this.Lock()
		this.m_dataStores[storeId] = store
		this.Unlock()
	} else {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to create dataStore with id '"+storeId+"' and with error '"+err.Error()+"'."))
		return nil, cargoError
	}

	return store, nil
}

func (this *DataManager) deleteDataStore(storeId string) *CargoEntities.Error {

	if !Utility.IsValidVariableName(storeId) {
		cargoError := NewError(Utility.FileLine(), INVALID_VARIABLE_NAME_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' is not valid."))
		return cargoError
	}

	if this.getDataStore(storeId) == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_DOESNT_EXIST_ERROR, SERVER_ERROR_CODE, errors.New("The storeId '"+storeId+"' doesn't exist."))
		log.Println("------> Store with id", storeId, "dosen't exist!")
		return cargoError
	}

	// Delete the dataStore configuration
	dataStoreConfigurationUuid := CargoConfigDataStoreConfigurationExists(storeId)
	dataStoreConfigurationEntity, errObj := GetServer().GetEntityManager().getEntityByUuid(dataStoreConfigurationUuid)

	// In case of the configuration is not already deleted...
	if errObj == nil {
		dataStoreConfigurationEntity.DeleteEntity()
	}

	// Remove the storeObject from the storeMap
	this.removeDataStore(storeId)

	// Delete the directory
	filePath := GetServer().GetConfigurationManager().GetDataPath() + "/" + storeId
	err := os.RemoveAll(filePath)

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Failed to delete directory '"+filePath+"' with error '"+err.Error()+"'."))
		log.Println("------> Fail to remove ", storeId, err)
		return cargoError
	}

	return nil

}

func (this *DataManager) close() {
	this.Lock()
	defer this.Unlock()

	// Close the data manager.
	for _, v := range this.m_dataStores {
		v.Close()
	}

}

////////////////////////////////////////////////////////////////////////////////
// public function
////////////////////////////////////////////////////////////////////////////////
/**
 * Execute a query that read information from the store and
 * return the result and an array of interface...
 */
func (this *DataManager) Ping(storeName string, messageId string, sessionId string) {
	store := this.getDataStore(storeName)
	if store == nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("The datastore '"+storeName+"' does not exist."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
	err := store.Ping()

	if err != nil {
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, errors.New("Fail to ping the data store "+err.Error()+"'."))
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
		return
	}
}

/**
 * Execute a query that read information from the store and
 * return the result and an array of interface...
 */
func (this *DataManager) Read(storeName string, query string, fieldsType []interface{}, params []interface{}, messageId string, sessionId string) [][]interface{} {
	data, err := this.readData(storeName, query, fieldsType, params)
	if err != nil {
		// Create the error message
		cargoError := NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err)
		GetServer().reportErrorMessage(messageId, sessionId, cargoError)
	}
	return data
}

/**
 * Execute a query that create a new data. The data contain de the new
 * value to insert in the DB.
 */
func (this *DataManager) Create(storeName string, query string, d []interface{}, messageId string, sessionId string) interface{} {
	lastId, err := this.createData(storeName, query, d)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err))
		return -1
	}
	return lastId
}

/**
 * Update the data.
 */
func (this *DataManager) Update(storeName string, query string, fields []interface{}, params []interface{}, messageId string, sessionId string) {
	err := this.updateData(storeName, query, fields, params)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err))
	}
}

/**
 * Remove the data.
 */
func (this *DataManager) Delete(storeName string, query string, params []interface{}, messageId string, sessionId string) {
	err := this.deleteData(storeName, query, params)
	if err != nil {
		GetServer().reportErrorMessage(messageId, sessionId, NewError(Utility.FileLine(), DATASTORE_ERROR, SERVER_ERROR_CODE, err))
	}
}

/**
 * Create a new data store.
 */
func (this *DataManager) CreateDataStore(storeId string, storeType int64, storeVendor int64, messageId string, sessionId string) {

	_, errObj := this.createDataStore(storeId, CargoConfig.DataStoreType(storeType), CargoConfig.DataStoreVendor(storeVendor))
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

/**
 * Delete a new data store.
 */
func (this *DataManager) DeleteDataStore(storeId string, messageId string, sessionId string) {

	errObj := this.deleteDataStore(storeId)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
	}
}

////////////////////////////////////////////////////////////////////////////////
//                              DataStore
////////////////////////////////////////////////////////////////////////////////

/**
 * This is the factory function that create the correct store depending
 * of he's information.
 */
func NewDataStore(info CargoConfig.DataStoreConfiguration) (DataStore, error) {
	var err error
	if info.M_dataStoreType == CargoConfig.DataStoreType_SQL_STORE {
		dataStore, err := NewSqlDataStore(info)
		return dataStore, err
	} else if info.M_dataStoreType == CargoConfig.DataStoreType_KEY_VALUE_STORE {
		dataStore, err := NewKeyValueDataStore(info)
		return dataStore, err
	}

	return nil, err
}

/**
 * DataStore is use to store data and do CRUD operation on it...
 */
type DataStore interface {
	/**
	 * Connection related stuff
	 */
	GetId() string

	/**
	 * Test if there's a connection with the server...
	 */
	Ping() error

	/** Crud interface **/
	Create(query string, data []interface{}) (lastId interface{}, err error)

	/**
	 * Param are filter to discard some element...
	 */
	Read(query string, fieldsType []interface{}, params []interface{}) ([][]interface{}, error)

	/**
	 * Update
	 */
	Update(query string, fields []interface{}, params []interface{}) error

	/**
	 * Delete values that match given parameter...
	 */
	Delete(query string, params []interface{}) error

	/**
	 * Close the data store, remove all connections or lnk to the data store.
	 */
	Close() error

	/**
	 * Open the data store connection.
	 */
	Connect() error
}
