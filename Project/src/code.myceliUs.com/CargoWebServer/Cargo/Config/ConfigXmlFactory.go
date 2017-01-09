package Config
import (
	"golang.org/x/net/html/charset"
	"encoding/xml"
	"strings"
	"os"
	"log"
	"path/filepath"
	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	"code.myceliUs.com/CargoWebServer/Cargo/Config/CargoConfig"
)
type ConfigXmlFactory struct {
	m_references map[string] interface{}
	m_object map[string]map[string][]string
}

/** Initialization function from xml file **/
func (this *ConfigXmlFactory)InitXml(inputPath string, object *CargoConfig.Configurations) error{
	this.m_references = make(map[string]interface{})
	this.m_object = make(map[string]map[string][]string)
	xmlFilePath, err := filepath.Abs(inputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	reader, err := os.Open(xmlFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var xmlElement *CargoConfig.XsdConfigurations
	xmlElement = new(CargoConfig.XsdConfigurations)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(xmlElement); err != nil {
		return err
	}
	this.InitConfigurations(xmlElement, object)
	for ref0, refMap := range this.m_object {
		refOwner := this.m_references[ref0]
		if refOwner != nil {
			for ref1, _ := range refMap {
				refs := refMap[ref1]
				for i:=0; i<len(refs); i++{
					ref:= this.m_references[refs[i]]
					if  ref != nil {
						params := make([]interface {},0)
						params = append(params,ref)
						methodName := "Set" + strings.ToUpper(ref1[0:1]) + ref1[1:]
						Utility.CallMethod(refOwner, methodName, params )
					}else{
						params := make([]interface {},0)
						params = append(params,refs[i])
						methodName := "Set" + strings.ToUpper(ref1[0:1]) + ref1[1:]
						Utility.CallMethod(refOwner, methodName, params)
					}
				}
			}
		}
	}
	return nil
}

/** Serialization to xml file **/
func (this *ConfigXmlFactory)SerializeXml(outputPath string, toSerialize *CargoConfig.Configurations) error{
	xmlFilePath, err := filepath.Abs(outputPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fo, err := os.Create(xmlFilePath)
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	var xmlElement *CargoConfig.XsdConfigurations
	xmlElement = new(CargoConfig.XsdConfigurations)

	this.SerialyzeConfigurations(xmlElement, toSerialize)
	output, err := xml.MarshalIndent(xmlElement, "  ", "    ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fileContent := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n")
	fileContent = append(fileContent, output...)
	_, err = fo.Write(fileContent)
	return nil
}

/** inititialisation of ServerConfiguration **/
func (this *ConfigXmlFactory) InitServerConfiguration(xmlElement *CargoConfig.XsdServerConfiguration,object *CargoConfig.ServerConfiguration){
	log.Println("Initialize ServerConfiguration")

	/** ServerConfiguration **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	object.M_ipv4= xmlElement.M_ipv4

	/** Configuration **/
	object.M_hostName= xmlElement.M_hostName

	/** Configuration **/
	object.M_port= xmlElement.M_port

	/** Configuration **/
	object.M_applicationsPath= xmlElement.M_applicationsPath

	/** Configuration **/
	object.M_dataPath= xmlElement.M_dataPath

	/** Configuration **/
	object.M_scriptsPath= xmlElement.M_scriptsPath

	/** Configuration **/
	object.M_definitionsPath= xmlElement.M_definitionsPath

	/** Configuration **/
	object.M_schemasPath= xmlElement.M_schemasPath

	/** Configuration **/
	object.M_tmpPath= xmlElement.M_tmpPath

	/** Configuration **/
	object.M_binPath= xmlElement.M_binPath
}

/** inititialisation of ApplicationConfiguration **/
func (this *ConfigXmlFactory) InitApplicationConfiguration(xmlElement *CargoConfig.XsdApplicationConfiguration,object *CargoConfig.ApplicationConfiguration){
	log.Println("Initialize ApplicationConfiguration")

	/** ApplicationConfiguration **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	object.M_indexPage= xmlElement.M_indexPage
}

/** inititialisation of SmtpConfiguration **/
func (this *ConfigXmlFactory) InitSmtpConfiguration(xmlElement *CargoConfig.XsdSmtpConfiguration,object *CargoConfig.SmtpConfiguration){
	log.Println("Initialize SmtpConfiguration")

	/** SmtpConfiguration **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	object.M_hostName= xmlElement.M_hostName

	/** Configuration **/
	object.M_ipv4= xmlElement.M_ipv4

	/** Configuration **/
	object.M_port= xmlElement.M_port

	/** Configuration **/
	object.M_user= xmlElement.M_user

	/** Configuration **/
	object.M_pwd= xmlElement.M_pwd
}

/** inititialisation of LdapConfiguration **/
func (this *ConfigXmlFactory) InitLdapConfiguration(xmlElement *CargoConfig.XsdLdapConfiguration,object *CargoConfig.LdapConfiguration){
	log.Println("Initialize LdapConfiguration")

	/** LdapConfiguration **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	object.M_hostName= xmlElement.M_hostName

	/** Configuration **/
	object.M_ipv4= xmlElement.M_ipv4

	/** Configuration **/
	object.M_port= xmlElement.M_port

	/** Configuration **/
	object.M_user= xmlElement.M_user

	/** Configuration **/
	object.M_pwd= xmlElement.M_pwd

	/** Configuration **/
	object.M_domain= xmlElement.M_domain

	/** Configuration **/
	object.M_searchBase= xmlElement.M_searchBase
}

/** inititialisation of DataStoreConfiguration **/
func (this *ConfigXmlFactory) InitDataStoreConfiguration(xmlElement *CargoConfig.XsdDataStoreConfiguration,object *CargoConfig.DataStoreConfiguration){
	log.Println("Initialize DataStoreConfiguration")

	/** DataStoreConfiguration **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	object.M_hostName= xmlElement.M_hostName

	/** Configuration **/
	object.M_ipv4= xmlElement.M_ipv4

	/** Configuration **/
	object.M_port= xmlElement.M_port

	/** Configuration **/
	object.M_user= xmlElement.M_user

	/** Configuration **/
	object.M_pwd= xmlElement.M_pwd

	/** DataStoreType **/
	if xmlElement.M_dataStoreType=="##SQL_STORE"{
		object.M_dataStoreType=CargoConfig.DataStoreType_SQL_STORE
	} else if xmlElement.M_dataStoreType=="##KEY_VALUE_STORE"{
		object.M_dataStoreType=CargoConfig.DataStoreType_KEY_VALUE_STORE
	}

	/** DataStoreVendor **/
	if xmlElement.M_dataStoreVendor=="##MYCELIUS"{
		object.M_dataStoreVendor=CargoConfig.DataStoreVendor_MYCELIUS
	} else if xmlElement.M_dataStoreVendor=="##MYSQL"{
		object.M_dataStoreVendor=CargoConfig.DataStoreVendor_MYSQL
	} else if xmlElement.M_dataStoreVendor=="##MSSQL"{
		object.M_dataStoreVendor=CargoConfig.DataStoreVendor_MSSQL
	}
}

/** inititialisation of Configurations **/
func (this *ConfigXmlFactory) InitConfigurations(xmlElement *CargoConfig.XsdConfigurations,object *CargoConfig.Configurations){
	log.Println("Initialize Configurations")

	/** Init serverConfiguration **/
	if object.M_serverConfig== nil{
		object.M_serverConfig= new(CargoConfig.ServerConfiguration)
	}
	this.InitServerConfiguration(&xmlElement.M_serverConfig,object.M_serverConfig)

		/** association initialisation **/

	/** Init applicationConfiguration **/
	object.M_applicationConfigs= make([]*CargoConfig.ApplicationConfiguration,0)
	for i:=0;i<len(xmlElement.M_applicationConfigs); i++{
		val:=new(CargoConfig.ApplicationConfiguration)
		this.InitApplicationConfiguration(xmlElement.M_applicationConfigs[i],val)
		object.M_applicationConfigs= append(object.M_applicationConfigs, val)

		/** association initialisation **/
	}

	/** Init smtpConfiguration **/
	object.M_smtpConfigs= make([]*CargoConfig.SmtpConfiguration,0)
	for i:=0;i<len(xmlElement.M_smtpConfigs); i++{
		val:=new(CargoConfig.SmtpConfiguration)
		this.InitSmtpConfiguration(xmlElement.M_smtpConfigs[i],val)
		object.M_smtpConfigs= append(object.M_smtpConfigs, val)

		/** association initialisation **/
	}

	/** Init ldapConfiguration **/
	object.M_ldapConfigs= make([]*CargoConfig.LdapConfiguration,0)
	for i:=0;i<len(xmlElement.M_ldapConfigs); i++{
		val:=new(CargoConfig.LdapConfiguration)
		this.InitLdapConfiguration(xmlElement.M_ldapConfigs[i],val)
		object.M_ldapConfigs= append(object.M_ldapConfigs, val)

		/** association initialisation **/
	}

	/** Init dataStoreConfiguration **/
	object.M_dataStoreConfigs= make([]*CargoConfig.DataStoreConfiguration,0)
	for i:=0;i<len(xmlElement.M_dataStoreConfigs); i++{
		val:=new(CargoConfig.DataStoreConfiguration)
		this.InitDataStoreConfiguration(xmlElement.M_dataStoreConfigs[i],val)
		object.M_dataStoreConfigs= append(object.M_dataStoreConfigs, val)

		/** association initialisation **/
	}

	/** Configurations **/
	object.M_id= xmlElement.M_id

	/** Configurations **/
	object.M_name= xmlElement.M_name

	/** Configurations **/
	object.M_version= xmlElement.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of LdapConfiguration **/
func (this *ConfigXmlFactory) SerialyzeLdapConfiguration(xmlElement *CargoConfig.XsdLdapConfiguration,object *CargoConfig.LdapConfiguration){
	if xmlElement == nil{
		return
	}

	/** LdapConfiguration **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	xmlElement.M_hostName= object.M_hostName

	/** Configuration **/
	xmlElement.M_ipv4= object.M_ipv4

	/** Configuration **/
	xmlElement.M_port= object.M_port

	/** Configuration **/
	xmlElement.M_user= object.M_user

	/** Configuration **/
	xmlElement.M_pwd= object.M_pwd

	/** Configuration **/
	xmlElement.M_domain= object.M_domain

	/** Configuration **/
	xmlElement.M_searchBase= object.M_searchBase
}

/** serialysation of DataStoreConfiguration **/
func (this *ConfigXmlFactory) SerialyzeDataStoreConfiguration(xmlElement *CargoConfig.XsdDataStoreConfiguration,object *CargoConfig.DataStoreConfiguration){
	if xmlElement == nil{
		return
	}

	/** DataStoreConfiguration **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	xmlElement.M_hostName= object.M_hostName

	/** Configuration **/
	xmlElement.M_ipv4= object.M_ipv4

	/** Configuration **/
	xmlElement.M_port= object.M_port

	/** Configuration **/
	xmlElement.M_user= object.M_user

	/** Configuration **/
	xmlElement.M_pwd= object.M_pwd

	/** DataStoreType **/
	if object.M_dataStoreType==CargoConfig.DataStoreType_SQL_STORE{
		xmlElement.M_dataStoreType="##SQL_STORE"
	} else if object.M_dataStoreType==CargoConfig.DataStoreType_KEY_VALUE_STORE{
		xmlElement.M_dataStoreType="##KEY_VALUE_STORE"
	}

	/** DataStoreVendor **/
	if object.M_dataStoreVendor==CargoConfig.DataStoreVendor_MYCELIUS{
		xmlElement.M_dataStoreVendor="##MYCELIUS"
	} else if object.M_dataStoreVendor==CargoConfig.DataStoreVendor_MYSQL{
		xmlElement.M_dataStoreVendor="##MYSQL"
	} else if object.M_dataStoreVendor==CargoConfig.DataStoreVendor_MSSQL{
		xmlElement.M_dataStoreVendor="##MSSQL"
	}
}

/** serialysation of Configurations **/
func (this *ConfigXmlFactory) SerialyzeConfigurations(xmlElement *CargoConfig.XsdConfigurations,object *CargoConfig.Configurations){
	if xmlElement == nil{
		return
	}

	/** Serialyze ServerConfiguration **/

	/** Now I will save the value of serverConfig **/
	if object.M_serverConfig!=nil{
		this.SerialyzeServerConfiguration(&xmlElement.M_serverConfig,object.M_serverConfig)
	}

	/** Serialyze ApplicationConfiguration **/
	if len(object.M_applicationConfigs) > 0 {
		xmlElement.M_applicationConfigs= make([]*CargoConfig.XsdApplicationConfiguration,0)
	}

	/** Now I will save the value of applicationConfigs **/
	for i:=0; i<len(object.M_applicationConfigs);i++{
		xmlElement.M_applicationConfigs=append(xmlElement.M_applicationConfigs,new(CargoConfig.XsdApplicationConfiguration))
		this.SerialyzeApplicationConfiguration(xmlElement.M_applicationConfigs[i],object.M_applicationConfigs[i])
	}

	/** Serialyze SmtpConfiguration **/
	if len(object.M_smtpConfigs) > 0 {
		xmlElement.M_smtpConfigs= make([]*CargoConfig.XsdSmtpConfiguration,0)
	}

	/** Now I will save the value of smtpConfigs **/
	for i:=0; i<len(object.M_smtpConfigs);i++{
		xmlElement.M_smtpConfigs=append(xmlElement.M_smtpConfigs,new(CargoConfig.XsdSmtpConfiguration))
		this.SerialyzeSmtpConfiguration(xmlElement.M_smtpConfigs[i],object.M_smtpConfigs[i])
	}

	/** Serialyze LdapConfiguration **/
	if len(object.M_ldapConfigs) > 0 {
		xmlElement.M_ldapConfigs= make([]*CargoConfig.XsdLdapConfiguration,0)
	}

	/** Now I will save the value of ldapConfigs **/
	for i:=0; i<len(object.M_ldapConfigs);i++{
		xmlElement.M_ldapConfigs=append(xmlElement.M_ldapConfigs,new(CargoConfig.XsdLdapConfiguration))
		this.SerialyzeLdapConfiguration(xmlElement.M_ldapConfigs[i],object.M_ldapConfigs[i])
	}

	/** Serialyze DataStoreConfiguration **/
	if len(object.M_dataStoreConfigs) > 0 {
		xmlElement.M_dataStoreConfigs= make([]*CargoConfig.XsdDataStoreConfiguration,0)
	}

	/** Now I will save the value of dataStoreConfigs **/
	for i:=0; i<len(object.M_dataStoreConfigs);i++{
		xmlElement.M_dataStoreConfigs=append(xmlElement.M_dataStoreConfigs,new(CargoConfig.XsdDataStoreConfiguration))
		this.SerialyzeDataStoreConfiguration(xmlElement.M_dataStoreConfigs[i],object.M_dataStoreConfigs[i])
	}

	/** Configurations **/
	xmlElement.M_id= object.M_id

	/** Configurations **/
	xmlElement.M_name= object.M_name

	/** Configurations **/
	xmlElement.M_version= object.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of ServerConfiguration **/
func (this *ConfigXmlFactory) SerialyzeServerConfiguration(xmlElement *CargoConfig.XsdServerConfiguration,object *CargoConfig.ServerConfiguration){
	if xmlElement == nil{
		return
	}

	/** ServerConfiguration **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	xmlElement.M_ipv4= object.M_ipv4

	/** Configuration **/
	xmlElement.M_hostName= object.M_hostName

	/** Configuration **/
	xmlElement.M_port= object.M_port

	/** Configuration **/
	xmlElement.M_applicationsPath= object.M_applicationsPath

	/** Configuration **/
	xmlElement.M_dataPath= object.M_dataPath

	/** Configuration **/
	xmlElement.M_scriptsPath= object.M_scriptsPath

	/** Configuration **/
	xmlElement.M_definitionsPath= object.M_definitionsPath

	/** Configuration **/
	xmlElement.M_schemasPath= object.M_schemasPath

	/** Configuration **/
	xmlElement.M_tmpPath= object.M_tmpPath

	/** Configuration **/
	xmlElement.M_binPath= object.M_binPath
}

/** serialysation of ApplicationConfiguration **/
func (this *ConfigXmlFactory) SerialyzeApplicationConfiguration(xmlElement *CargoConfig.XsdApplicationConfiguration,object *CargoConfig.ApplicationConfiguration){
	if xmlElement == nil{
		return
	}

	/** ApplicationConfiguration **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	xmlElement.M_indexPage= object.M_indexPage
}

/** serialysation of SmtpConfiguration **/
func (this *ConfigXmlFactory) SerialyzeSmtpConfiguration(xmlElement *CargoConfig.XsdSmtpConfiguration,object *CargoConfig.SmtpConfiguration){
	if xmlElement == nil{
		return
	}

	/** SmtpConfiguration **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Configuration **/
	xmlElement.M_hostName= object.M_hostName

	/** Configuration **/
	xmlElement.M_ipv4= object.M_ipv4

	/** Configuration **/
	xmlElement.M_port= object.M_port

	/** Configuration **/
	xmlElement.M_user= object.M_user

	/** Configuration **/
	xmlElement.M_pwd= object.M_pwd
}