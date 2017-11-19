// +build CargoEntities

package Server

import (
	"encoding/json"
	"log"
	"strings"
	"unsafe"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/Utility"
)

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_EntityEntityPrototype() {

	var entityEntityProto EntityPrototype
	entityEntityProto.TypeName = "CargoEntities.Entity"
	entityEntityProto.IsAbstract = true
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Error")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.LogEntry")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Log")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Project")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.TextMessage")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Computer")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.File")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Notification")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Account")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.User")
	entityEntityProto.SubstitutionGroup = append(entityEntityProto.SubstitutionGroup, "CargoEntities.Group")
	entityEntityProto.Ids = append(entityEntityProto.Ids, "UUID")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "UUID")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 0)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")
	entityEntityProto.Indexs = append(entityEntityProto.Indexs, "ParentUuid")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "ParentUuid")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 1)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")
	entityEntityProto.Fields = append(entityEntityProto.Fields, "ParentLnk")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.string")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 2)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	entityEntityProto.Ids = append(entityEntityProto.Ids, "M_id")
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 3)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, true)
	entityEntityProto.Fields = append(entityEntityProto.Fields, "M_id")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "xs.ID")
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "")

	/** associations of Entity **/
	entityEntityProto.FieldsOrder = append(entityEntityProto.FieldsOrder, 4)
	entityEntityProto.FieldsVisibility = append(entityEntityProto.FieldsVisibility, false)
	entityEntityProto.Fields = append(entityEntityProto.Fields, "M_entitiesPtr")
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "undefined")
	entityEntityProto.FieldsDefaultValue = append(entityEntityProto.FieldsDefaultValue, "undefined")
	entityEntityProto.FieldsType = append(entityEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&entityEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			Parameter
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ParameterEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Parameter
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesParameterEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ParameterEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesParameterExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Parameter).TYPENAME = "CargoEntities.Parameter"
		object.(*CargoEntities.Parameter).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Parameter", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Parameter).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Parameter).UUID
			}
			return val.(*CargoEntities_ParameterEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_ParameterEntity)
	if object == nil {
		entity.object = new(CargoEntities.Parameter)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Parameter)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Parameter"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ParameterEntity) GetTypeName() string {
	return "CargoEntities.Parameter"
}
func (this *CargoEntities_ParameterEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_ParameterEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_ParameterEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_ParameterEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_ParameterEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_ParameterEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_ParameterEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ParameterEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_ParameterEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_ParameterEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_ParameterEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ParameterEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_ParameterEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_ParameterEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ParameterEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ParameterEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ParameterEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ParameterEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ParameterEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ParameterEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ParameterEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_ParameterEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ParameterEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ParameterEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ParameterEntityPrototype() {

	var parameterEntityProto EntityPrototype
	parameterEntityProto.TypeName = "CargoEntities.Parameter"
	parameterEntityProto.Ids = append(parameterEntityProto.Ids, "UUID")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "UUID")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 0)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.Indexs = append(parameterEntityProto.Indexs, "ParentUuid")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "ParentUuid")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 1)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "ParentLnk")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 2)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")

	/** members of Parameter **/
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 3)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_name")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 4)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_type")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.string")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "")
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 5)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, true)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_isArray")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "xs.boolean")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "false")

	/** associations of Parameter **/
	parameterEntityProto.FieldsOrder = append(parameterEntityProto.FieldsOrder, 6)
	parameterEntityProto.FieldsVisibility = append(parameterEntityProto.FieldsVisibility, false)
	parameterEntityProto.Fields = append(parameterEntityProto.Fields, "M_parametersPtr")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "undefined")
	parameterEntityProto.FieldsDefaultValue = append(parameterEntityProto.FieldsDefaultValue, "undefined")
	parameterEntityProto.FieldsType = append(parameterEntityProto.FieldsType, "CargoEntities.Parameter:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&parameterEntityProto)

}

/** Create **/
func (this *CargoEntities_ParameterEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Parameter **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_isArray")

	/** associations of Parameter **/
	query.Fields = append(query.Fields, "M_parametersPtr")

	var ParameterInfo []interface{}

	ParameterInfo = append(ParameterInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ParameterInfo = append(ParameterInfo, this.GetParentPtr().GetUuid())
		ParameterInfo = append(ParameterInfo, this.GetParentLnk())
	} else {
		ParameterInfo = append(ParameterInfo, "")
		ParameterInfo = append(ParameterInfo, "")
	}

	/** members of Parameter **/
	ParameterInfo = append(ParameterInfo, this.object.M_name)
	ParameterInfo = append(ParameterInfo, this.object.M_type)
	ParameterInfo = append(ParameterInfo, this.object.M_isArray)

	/** associations of Parameter **/

	/** Save parameters type Parameter **/
	/** attribute Parameter has no method GetId, must be an error here...*/
	ParameterInfo = append(ParameterInfo, "")
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ParameterInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ParameterInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ParameterEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ParameterEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Parameter **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_isArray")

	/** associations of Parameter **/
	query.Fields = append(query.Fields, "M_parametersPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Parameter...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Parameter)
		this.object.TYPENAME = "CargoEntities.Parameter"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Parameter **/

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** type **/
		if results[0][4] != nil {
			this.object.M_type = results[0][4].(string)
		}

		/** isArray **/
		if results[0][5] != nil {
			this.object.M_isArray = results[0][5].(bool)
		}

		/** associations of Parameter **/

		/** parametersPtr **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Parameter"
				id_ := refTypeName + "$$" + id
				this.object.M_parametersPtr = id
				GetServer().GetEntityManager().appendReference("parametersPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesParameterEntityFromObject(object *CargoEntities.Parameter) *CargoEntities_ParameterEntity {
	return this.NewCargoEntitiesParameterEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ParameterEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesParameterExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Parameter"
	query.Indexs = append(query.Indexs, "M_name="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ParameterEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_ParameterEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Action
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ActionEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Action
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesActionEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ActionEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesActionExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Action).TYPENAME = "CargoEntities.Action"
		object.(*CargoEntities.Action).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Action", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Action).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Action).UUID
			}
			return val.(*CargoEntities_ActionEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_ActionEntity)
	if object == nil {
		entity.object = new(CargoEntities.Action)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Action)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Action"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ActionEntity) GetTypeName() string {
	return "CargoEntities.Action"
}
func (this *CargoEntities_ActionEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_ActionEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_ActionEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_ActionEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_ActionEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_ActionEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_ActionEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ActionEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_ActionEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_ActionEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_ActionEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ActionEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_ActionEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_ActionEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ActionEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ActionEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ActionEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ActionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ActionEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ActionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ActionEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_ActionEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ActionEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Action"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ActionEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ActionEntityPrototype() {

	var actionEntityProto EntityPrototype
	actionEntityProto.TypeName = "CargoEntities.Action"
	actionEntityProto.Ids = append(actionEntityProto.Ids, "UUID")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "UUID")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 0)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.Indexs = append(actionEntityProto.Indexs, "ParentUuid")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "ParentUuid")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 1)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.Fields = append(actionEntityProto.Fields, "ParentLnk")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 2)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")

	/** members of Action **/
	actionEntityProto.Ids = append(actionEntityProto.Ids, "M_name")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 3)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_name")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.ID")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 4)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_doc")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "xs.string")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 5)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_parameters")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "[]")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]CargoEntities.Parameter")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 6)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_results")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "[]")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "[]CargoEntities.Parameter")
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 7)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, true)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_accessType")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "1")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "enum:AccessType_Hidden:AccessType_Public:AccessType_Restricted")

	/** associations of Action **/
	actionEntityProto.FieldsOrder = append(actionEntityProto.FieldsOrder, 8)
	actionEntityProto.FieldsVisibility = append(actionEntityProto.FieldsVisibility, false)
	actionEntityProto.Fields = append(actionEntityProto.Fields, "M_entitiesPtr")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "undefined")
	actionEntityProto.FieldsDefaultValue = append(actionEntityProto.FieldsDefaultValue, "undefined")
	actionEntityProto.FieldsType = append(actionEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&actionEntityProto)

}

/** Create **/
func (this *CargoEntities_ActionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Action"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Action **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_doc")
	query.Fields = append(query.Fields, "M_parameters")
	query.Fields = append(query.Fields, "M_results")
	query.Fields = append(query.Fields, "M_accessType")

	/** associations of Action **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var ActionInfo []interface{}

	ActionInfo = append(ActionInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ActionInfo = append(ActionInfo, this.GetParentPtr().GetUuid())
		ActionInfo = append(ActionInfo, this.GetParentLnk())
	} else {
		ActionInfo = append(ActionInfo, "")
		ActionInfo = append(ActionInfo, "")
	}

	/** members of Action **/
	ActionInfo = append(ActionInfo, this.object.M_name)
	ActionInfo = append(ActionInfo, this.object.M_doc)

	/** Save parameters type Parameter **/
	parametersIds := make([]string, 0)
	lazy_parameters := this.lazyMap["M_parameters"] != nil && len(this.object.M_parameters) == 0
	if !lazy_parameters {
		for i := 0; i < len(this.object.M_parameters); i++ {
			parametersEntity := GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), this.object.M_parameters[i].UUID, this.object.M_parameters[i])
			parametersIds = append(parametersIds, parametersEntity.GetUuid())
			parametersEntity.AppendReferenced("parameters", this)
			this.AppendChild("parameters", parametersEntity)
			if parametersEntity.NeedSave() {
				parametersEntity.SaveEntity()
			}
		}
	} else {
		parametersIds = this.lazyMap["M_parameters"].([]string)
	}
	parametersStr, _ := json.Marshal(parametersIds)
	ActionInfo = append(ActionInfo, string(parametersStr))

	/** Save results type Parameter **/
	resultsIds := make([]string, 0)
	lazy_results := this.lazyMap["M_results"] != nil && len(this.object.M_results) == 0
	if !lazy_results {
		for i := 0; i < len(this.object.M_results); i++ {
			resultsEntity := GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), this.object.M_results[i].UUID, this.object.M_results[i])
			resultsIds = append(resultsIds, resultsEntity.GetUuid())
			resultsEntity.AppendReferenced("results", this)
			this.AppendChild("results", resultsEntity)
			if resultsEntity.NeedSave() {
				resultsEntity.SaveEntity()
			}
		}
	} else {
		resultsIds = this.lazyMap["M_results"].([]string)
	}
	resultsStr, _ := json.Marshal(resultsIds)
	ActionInfo = append(ActionInfo, string(resultsStr))

	/** Save accessType type AccessType **/
	if this.object.M_accessType == CargoEntities.AccessType_Hidden {
		ActionInfo = append(ActionInfo, 0)
	} else if this.object.M_accessType == CargoEntities.AccessType_Public {
		ActionInfo = append(ActionInfo, 1)
	} else if this.object.M_accessType == CargoEntities.AccessType_Restricted {
		ActionInfo = append(ActionInfo, 2)
	} else {
		ActionInfo = append(ActionInfo, 0)
	}

	/** associations of Action **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		ActionInfo = append(ActionInfo, this.object.M_entitiesPtr)
	} else {
		ActionInfo = append(ActionInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ActionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ActionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ActionEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ActionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Action"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Action **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_doc")
	query.Fields = append(query.Fields, "M_parameters")
	query.Fields = append(query.Fields, "M_results")
	query.Fields = append(query.Fields, "M_accessType")

	/** associations of Action **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Action...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Action)
		this.object.TYPENAME = "CargoEntities.Action"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Action **/

		/** name **/
		if results[0][3] != nil {
			this.object.M_name = results[0][3].(string)
		}

		/** doc **/
		if results[0][4] != nil {
			this.object.M_doc = results[0][4].(string)
		}

		/** parameters **/
		if results[0][5] != nil {
			uuidsStr := results[0][5].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var parametersEntity *CargoEntities_ParameterEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							parametersEntity = instance.(*CargoEntities_ParameterEntity)
						} else {
							parametersEntity = GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), uuids[i], nil)
							parametersEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(parametersEntity)
						}
						parametersEntity.AppendReferenced("parameters", this)
						this.AppendChild("parameters", parametersEntity)
					}
				} else {
					this.lazyMap["M_parameters"] = uuids
				}
			}
		}

		/** results **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var resultsEntity *CargoEntities_ParameterEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							resultsEntity = instance.(*CargoEntities_ParameterEntity)
						} else {
							resultsEntity = GetServer().GetEntityManager().NewCargoEntitiesParameterEntity(this.GetUuid(), uuids[i], nil)
							resultsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(resultsEntity)
						}
						resultsEntity.AppendReferenced("results", this)
						this.AppendChild("results", resultsEntity)
					}
				} else {
					this.lazyMap["M_results"] = uuids
				}
			}
		}

		/** accessType **/
		if results[0][7] != nil {
			enumIndex := results[0][7].(int)
			if enumIndex == 0 {
				this.object.M_accessType = CargoEntities.AccessType_Hidden
			} else if enumIndex == 1 {
				this.object.M_accessType = CargoEntities.AccessType_Public
			} else if enumIndex == 2 {
				this.object.M_accessType = CargoEntities.AccessType_Restricted
			}
		}

		/** associations of Action **/

		/** entitiesPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesActionEntityFromObject(object *CargoEntities.Action) *CargoEntities_ActionEntity {
	return this.NewCargoEntitiesActionEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ActionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesActionExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Action"
	query.Indexs = append(query.Indexs, "M_name="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ActionEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_ActionEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Error
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ErrorEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Error
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesErrorEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ErrorEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesErrorExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Error).TYPENAME = "CargoEntities.Error"
		object.(*CargoEntities.Error).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Error", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Error).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Error).UUID
			}
			return val.(*CargoEntities_ErrorEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_ErrorEntity)
	if object == nil {
		entity.object = new(CargoEntities.Error)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Error)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Error"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ErrorEntity) GetTypeName() string {
	return "CargoEntities.Error"
}
func (this *CargoEntities_ErrorEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_ErrorEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_ErrorEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_ErrorEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_ErrorEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_ErrorEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_ErrorEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ErrorEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_ErrorEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_ErrorEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_ErrorEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ErrorEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_ErrorEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_ErrorEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ErrorEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ErrorEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ErrorEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ErrorEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ErrorEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ErrorEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ErrorEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_ErrorEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ErrorEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Error"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ErrorEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ErrorEntityPrototype() {

	var errorEntityProto EntityPrototype
	errorEntityProto.TypeName = "CargoEntities.Error"
	errorEntityProto.SuperTypeNames = append(errorEntityProto.SuperTypeNames, "CargoEntities.Entity")
	errorEntityProto.SuperTypeNames = append(errorEntityProto.SuperTypeNames, "CargoEntities.Message")
	errorEntityProto.Ids = append(errorEntityProto.Ids, "UUID")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "UUID")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 0)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")
	errorEntityProto.Indexs = append(errorEntityProto.Indexs, "ParentUuid")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "ParentUuid")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 1)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")
	errorEntityProto.Fields = append(errorEntityProto.Fields, "ParentLnk")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 2)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	errorEntityProto.Ids = append(errorEntityProto.Ids, "M_id")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 3)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_id")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.ID")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 4)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_body")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")

	/** members of Error **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 5)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_errorPath")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.string")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 6)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_code")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "xs.int")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "0")
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 7)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, true)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_accountRef")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "CargoEntities.Account:Ref")

	/** associations of Error **/
	errorEntityProto.FieldsOrder = append(errorEntityProto.FieldsOrder, 8)
	errorEntityProto.FieldsVisibility = append(errorEntityProto.FieldsVisibility, false)
	errorEntityProto.Fields = append(errorEntityProto.Fields, "M_entitiesPtr")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsDefaultValue = append(errorEntityProto.FieldsDefaultValue, "undefined")
	errorEntityProto.FieldsType = append(errorEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&errorEntityProto)

}

/** Create **/
func (this *CargoEntities_ErrorEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Error"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Error **/
	query.Fields = append(query.Fields, "M_errorPath")
	query.Fields = append(query.Fields, "M_code")
	query.Fields = append(query.Fields, "M_accountRef")

	/** associations of Error **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var ErrorInfo []interface{}

	ErrorInfo = append(ErrorInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ErrorInfo = append(ErrorInfo, this.GetParentPtr().GetUuid())
		ErrorInfo = append(ErrorInfo, this.GetParentLnk())
	} else {
		ErrorInfo = append(ErrorInfo, "")
		ErrorInfo = append(ErrorInfo, "")
	}

	/** members of Entity **/
	ErrorInfo = append(ErrorInfo, this.object.M_id)

	/** members of Message **/
	ErrorInfo = append(ErrorInfo, this.object.M_body)

	/** members of Error **/
	ErrorInfo = append(ErrorInfo, this.object.M_errorPath)
	ErrorInfo = append(ErrorInfo, this.object.M_code)

	/** Save accountRef type Account **/
	if len(this.object.M_accountRef) > 0 {
		ErrorInfo = append(ErrorInfo, this.object.M_accountRef)
	} else {
		ErrorInfo = append(ErrorInfo, "")
	}

	/** associations of Error **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		ErrorInfo = append(ErrorInfo, this.object.M_entitiesPtr)
	} else {
		ErrorInfo = append(ErrorInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ErrorInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ErrorInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ErrorEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ErrorEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Error"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Error **/
	query.Fields = append(query.Fields, "M_errorPath")
	query.Fields = append(query.Fields, "M_code")
	query.Fields = append(query.Fields, "M_accountRef")

	/** associations of Error **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Error...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Error)
		this.object.TYPENAME = "CargoEntities.Error"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Message **/

		/** body **/
		if results[0][4] != nil {
			this.object.M_body = results[0][4].(string)
		}

		/** members of Error **/

		/** errorPath **/
		if results[0][5] != nil {
			this.object.M_errorPath = results[0][5].(string)
		}

		/** code **/
		if results[0][6] != nil {
			this.object.M_code = results[0][6].(int)
		}

		/** accountRef **/
		if results[0][7] != nil {
			id := results[0][7].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_accountRef = id
				GetServer().GetEntityManager().appendReference("accountRef", this.object.UUID, id_)
			}
		}

		/** associations of Error **/

		/** entitiesPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesErrorEntityFromObject(object *CargoEntities.Error) *CargoEntities_ErrorEntity {
	return this.NewCargoEntitiesErrorEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ErrorEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesErrorExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Error"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ErrorEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_ErrorEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			LogEntry
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_LogEntryEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.LogEntry
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesLogEntryEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_LogEntryEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesLogEntryExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.LogEntry).TYPENAME = "CargoEntities.LogEntry"
		object.(*CargoEntities.LogEntry).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.LogEntry", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.LogEntry).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.LogEntry).UUID
			}
			return val.(*CargoEntities_LogEntryEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_LogEntryEntity)
	if object == nil {
		entity.object = new(CargoEntities.LogEntry)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.LogEntry)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.LogEntry"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_LogEntryEntity) GetTypeName() string {
	return "CargoEntities.LogEntry"
}
func (this *CargoEntities_LogEntryEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_LogEntryEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_LogEntryEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_LogEntryEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_LogEntryEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_LogEntryEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_LogEntryEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_LogEntryEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_LogEntryEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_LogEntryEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_LogEntryEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_LogEntryEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_LogEntryEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_LogEntryEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_LogEntryEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_LogEntryEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_LogEntryEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_LogEntryEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_LogEntryEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_LogEntryEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_LogEntryEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_LogEntryEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_LogEntryEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_LogEntryEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_LogEntryEntityPrototype() {

	var logEntryEntityProto EntityPrototype
	logEntryEntityProto.TypeName = "CargoEntities.LogEntry"
	logEntryEntityProto.SuperTypeNames = append(logEntryEntityProto.SuperTypeNames, "CargoEntities.Entity")
	logEntryEntityProto.Ids = append(logEntryEntityProto.Ids, "UUID")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "UUID")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 0)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")
	logEntryEntityProto.Indexs = append(logEntryEntityProto.Indexs, "ParentUuid")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "ParentUuid")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 1)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "ParentLnk")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.string")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 2)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	logEntryEntityProto.Ids = append(logEntryEntityProto.Ids, "M_id")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 3)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_id")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.ID")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "")

	/** members of LogEntry **/
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 4)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_creationTime")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "xs.date")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "new Date()")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 5)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, true)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_entityRef")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Entity:Ref")

	/** associations of LogEntry **/
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 6)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_loggerPtr")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Log:Ref")
	logEntryEntityProto.FieldsOrder = append(logEntryEntityProto.FieldsOrder, 7)
	logEntryEntityProto.FieldsVisibility = append(logEntryEntityProto.FieldsVisibility, false)
	logEntryEntityProto.Fields = append(logEntryEntityProto.Fields, "M_entitiesPtr")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsDefaultValue = append(logEntryEntityProto.FieldsDefaultValue, "undefined")
	logEntryEntityProto.FieldsType = append(logEntryEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&logEntryEntityProto)

}

/** Create **/
func (this *CargoEntities_LogEntryEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of LogEntry **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_entityRef")

	/** associations of LogEntry **/
	query.Fields = append(query.Fields, "M_loggerPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var LogEntryInfo []interface{}

	LogEntryInfo = append(LogEntryInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		LogEntryInfo = append(LogEntryInfo, this.GetParentPtr().GetUuid())
		LogEntryInfo = append(LogEntryInfo, this.GetParentLnk())
	} else {
		LogEntryInfo = append(LogEntryInfo, "")
		LogEntryInfo = append(LogEntryInfo, "")
	}

	/** members of Entity **/
	LogEntryInfo = append(LogEntryInfo, this.object.M_id)

	/** members of LogEntry **/
	LogEntryInfo = append(LogEntryInfo, this.object.M_creationTime)

	/** Save entityRef type Entity **/
	if len(this.object.M_entityRef) > 0 {
		LogEntryInfo = append(LogEntryInfo, this.object.M_entityRef)
	} else {
		LogEntryInfo = append(LogEntryInfo, "")
	}

	/** associations of LogEntry **/

	/** Save logger type Log **/
	if len(this.object.M_loggerPtr) > 0 {
		LogEntryInfo = append(LogEntryInfo, this.object.M_loggerPtr)
	} else {
		LogEntryInfo = append(LogEntryInfo, "")
	}

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		LogEntryInfo = append(LogEntryInfo, this.object.M_entitiesPtr)
	} else {
		LogEntryInfo = append(LogEntryInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), LogEntryInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), LogEntryInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_LogEntryEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_LogEntryEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of LogEntry **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_entityRef")

	/** associations of LogEntry **/
	query.Fields = append(query.Fields, "M_loggerPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of LogEntry...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.LogEntry)
		this.object.TYPENAME = "CargoEntities.LogEntry"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of LogEntry **/

		/** creationTime **/
		if results[0][4] != nil {
			this.object.M_creationTime = results[0][4].(int64)
		}

		/** entityRef **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entity"
				id_ := refTypeName + "$$" + id
				this.object.M_entityRef = id
				GetServer().GetEntityManager().appendReference("entityRef", this.object.UUID, id_)
			}
		}

		/** associations of LogEntry **/

		/** loggerPtr **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Log"
				id_ := refTypeName + "$$" + id
				this.object.M_loggerPtr = id
				GetServer().GetEntityManager().appendReference("loggerPtr", this.object.UUID, id_)
			}
		}

		/** entitiesPtr **/
		if results[0][7] != nil {
			id := results[0][7].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesLogEntryEntityFromObject(object *CargoEntities.LogEntry) *CargoEntities_LogEntryEntity {
	return this.NewCargoEntitiesLogEntryEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_LogEntryEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesLogEntryExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.LogEntry"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_LogEntryEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_LogEntryEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Log
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_LogEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Log
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesLogEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_LogEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesLogExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Log).TYPENAME = "CargoEntities.Log"
		object.(*CargoEntities.Log).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Log", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Log).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Log).UUID
			}
			return val.(*CargoEntities_LogEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_LogEntity)
	if object == nil {
		entity.object = new(CargoEntities.Log)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Log)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Log"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_LogEntity) GetTypeName() string {
	return "CargoEntities.Log"
}
func (this *CargoEntities_LogEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_LogEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_LogEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_LogEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_LogEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_LogEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_LogEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_LogEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_LogEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_LogEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_LogEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_LogEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_LogEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_LogEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_LogEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_LogEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_LogEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_LogEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_LogEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_LogEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_LogEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_LogEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_LogEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Log"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_LogEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_LogEntityPrototype() {

	var logEntityProto EntityPrototype
	logEntityProto.TypeName = "CargoEntities.Log"
	logEntityProto.SuperTypeNames = append(logEntityProto.SuperTypeNames, "CargoEntities.Entity")
	logEntityProto.Ids = append(logEntityProto.Ids, "UUID")
	logEntityProto.Fields = append(logEntityProto.Fields, "UUID")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 0)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")
	logEntityProto.Indexs = append(logEntityProto.Indexs, "ParentUuid")
	logEntityProto.Fields = append(logEntityProto.Fields, "ParentUuid")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 1)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")
	logEntityProto.Fields = append(logEntityProto.Fields, "ParentLnk")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.string")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 2)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	logEntityProto.Ids = append(logEntityProto.Ids, "M_id")
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 3)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, true)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_id")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "xs.ID")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "")

	/** members of Log **/
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 4)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, true)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_entries")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "[]")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "[]CargoEntities.LogEntry")

	/** associations of Log **/
	logEntityProto.FieldsOrder = append(logEntityProto.FieldsOrder, 5)
	logEntityProto.FieldsVisibility = append(logEntityProto.FieldsVisibility, false)
	logEntityProto.Fields = append(logEntityProto.Fields, "M_entitiesPtr")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "undefined")
	logEntityProto.FieldsDefaultValue = append(logEntityProto.FieldsDefaultValue, "undefined")
	logEntityProto.FieldsType = append(logEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&logEntityProto)

}

/** Create **/
func (this *CargoEntities_LogEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Log"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Log **/
	query.Fields = append(query.Fields, "M_entries")

	/** associations of Log **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var LogInfo []interface{}

	LogInfo = append(LogInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		LogInfo = append(LogInfo, this.GetParentPtr().GetUuid())
		LogInfo = append(LogInfo, this.GetParentLnk())
	} else {
		LogInfo = append(LogInfo, "")
		LogInfo = append(LogInfo, "")
	}

	/** members of Entity **/
	LogInfo = append(LogInfo, this.object.M_id)

	/** members of Log **/

	/** Save entries type LogEntry **/
	entriesIds := make([]string, 0)
	lazy_entries := this.lazyMap["M_entries"] != nil && len(this.object.M_entries) == 0
	if !lazy_entries {
		for i := 0; i < len(this.object.M_entries); i++ {
			entriesEntity := GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), this.object.M_entries[i].UUID, this.object.M_entries[i])
			entriesIds = append(entriesIds, entriesEntity.GetUuid())
			entriesEntity.AppendReferenced("entries", this)
			this.AppendChild("entries", entriesEntity)
			if entriesEntity.NeedSave() {
				entriesEntity.SaveEntity()
			}
		}
	} else {
		entriesIds = this.lazyMap["M_entries"].([]string)
	}
	entriesStr, _ := json.Marshal(entriesIds)
	LogInfo = append(LogInfo, string(entriesStr))

	/** associations of Log **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		LogInfo = append(LogInfo, this.object.M_entitiesPtr)
	} else {
		LogInfo = append(LogInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), LogInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), LogInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_LogEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_LogEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Log"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Log **/
	query.Fields = append(query.Fields, "M_entries")

	/** associations of Log **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Log...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Log)
		this.object.TYPENAME = "CargoEntities.Log"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Log **/

		/** entries **/
		if results[0][4] != nil {
			uuidsStr := results[0][4].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var entriesEntity *CargoEntities_LogEntryEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							entriesEntity = instance.(*CargoEntities_LogEntryEntity)
						} else {
							entriesEntity = GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), uuids[i], nil)
							entriesEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(entriesEntity)
						}
						entriesEntity.AppendReferenced("entries", this)
						this.AppendChild("entries", entriesEntity)
					}
				} else {
					this.lazyMap["M_entries"] = uuids
				}
			}
		}

		/** associations of Log **/

		/** entitiesPtr **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesLogEntityFromObject(object *CargoEntities.Log) *CargoEntities_LogEntity {
	return this.NewCargoEntitiesLogEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_LogEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesLogExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Log"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_LogEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_LogEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Project
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ProjectEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Project
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesProjectEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ProjectEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesProjectExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Project).TYPENAME = "CargoEntities.Project"
		object.(*CargoEntities.Project).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Project", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Project).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Project).UUID
			}
			return val.(*CargoEntities_ProjectEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_ProjectEntity)
	if object == nil {
		entity.object = new(CargoEntities.Project)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Project)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Project"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ProjectEntity) GetTypeName() string {
	return "CargoEntities.Project"
}
func (this *CargoEntities_ProjectEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_ProjectEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_ProjectEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_ProjectEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_ProjectEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_ProjectEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_ProjectEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ProjectEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_ProjectEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_ProjectEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_ProjectEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ProjectEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_ProjectEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_ProjectEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ProjectEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ProjectEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ProjectEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ProjectEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ProjectEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ProjectEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ProjectEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_ProjectEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ProjectEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Project"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ProjectEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ProjectEntityPrototype() {

	var projectEntityProto EntityPrototype
	projectEntityProto.TypeName = "CargoEntities.Project"
	projectEntityProto.SuperTypeNames = append(projectEntityProto.SuperTypeNames, "CargoEntities.Entity")
	projectEntityProto.Ids = append(projectEntityProto.Ids, "UUID")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "UUID")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 0)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")
	projectEntityProto.Indexs = append(projectEntityProto.Indexs, "ParentUuid")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "ParentUuid")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 1)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")
	projectEntityProto.Fields = append(projectEntityProto.Fields, "ParentLnk")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 2)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	projectEntityProto.Ids = append(projectEntityProto.Ids, "M_id")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 3)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_id")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.ID")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")

	/** members of Project **/
	projectEntityProto.Indexs = append(projectEntityProto.Indexs, "M_name")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 4)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_name")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "xs.string")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "")
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 5)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, true)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_filesRef")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "undefined")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "[]")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "[]CargoEntities.File:Ref")

	/** associations of Project **/
	projectEntityProto.FieldsOrder = append(projectEntityProto.FieldsOrder, 6)
	projectEntityProto.FieldsVisibility = append(projectEntityProto.FieldsVisibility, false)
	projectEntityProto.Fields = append(projectEntityProto.Fields, "M_entitiesPtr")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "undefined")
	projectEntityProto.FieldsDefaultValue = append(projectEntityProto.FieldsDefaultValue, "undefined")
	projectEntityProto.FieldsType = append(projectEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&projectEntityProto)

}

/** Create **/
func (this *CargoEntities_ProjectEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Project"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Project **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_filesRef")

	/** associations of Project **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var ProjectInfo []interface{}

	ProjectInfo = append(ProjectInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ProjectInfo = append(ProjectInfo, this.GetParentPtr().GetUuid())
		ProjectInfo = append(ProjectInfo, this.GetParentLnk())
	} else {
		ProjectInfo = append(ProjectInfo, "")
		ProjectInfo = append(ProjectInfo, "")
	}

	/** members of Entity **/
	ProjectInfo = append(ProjectInfo, this.object.M_id)

	/** members of Project **/
	ProjectInfo = append(ProjectInfo, this.object.M_name)

	/** Save filesRef type File **/
	filesRefStr, _ := json.Marshal(this.object.M_filesRef)
	ProjectInfo = append(ProjectInfo, string(filesRefStr))

	/** associations of Project **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		ProjectInfo = append(ProjectInfo, this.object.M_entitiesPtr)
	} else {
		ProjectInfo = append(ProjectInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ProjectInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ProjectInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ProjectEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ProjectEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Project"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Project **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_filesRef")

	/** associations of Project **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Project...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Project)
		this.object.TYPENAME = "CargoEntities.Project"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Project **/

		/** name **/
		if results[0][4] != nil {
			this.object.M_name = results[0][4].(string)
		}

		/** filesRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.File"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_filesRef = append(this.object.M_filesRef, ids[i])
					GetServer().GetEntityManager().appendReference("filesRef", this.object.UUID, id_)
				}
			}
		}

		/** associations of Project **/

		/** entitiesPtr **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesProjectEntityFromObject(object *CargoEntities.Project) *CargoEntities_ProjectEntity {
	return this.NewCargoEntitiesProjectEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ProjectEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesProjectExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Project"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ProjectEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_ProjectEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_MessageEntityPrototype() {

	var messageEntityProto EntityPrototype
	messageEntityProto.TypeName = "CargoEntities.Message"
	messageEntityProto.IsAbstract = true
	messageEntityProto.SuperTypeNames = append(messageEntityProto.SuperTypeNames, "CargoEntities.Entity")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.Error")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.Notification")
	messageEntityProto.SubstitutionGroup = append(messageEntityProto.SubstitutionGroup, "CargoEntities.TextMessage")
	messageEntityProto.Ids = append(messageEntityProto.Ids, "UUID")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "UUID")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 0)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")
	messageEntityProto.Indexs = append(messageEntityProto.Indexs, "ParentUuid")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "ParentUuid")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 1)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")
	messageEntityProto.Fields = append(messageEntityProto.Fields, "ParentLnk")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 2)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	messageEntityProto.Ids = append(messageEntityProto.Ids, "M_id")
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 3)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, true)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_id")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.ID")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 4)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, true)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_body")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "xs.string")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "")

	/** associations of Message **/
	messageEntityProto.FieldsOrder = append(messageEntityProto.FieldsOrder, 5)
	messageEntityProto.FieldsVisibility = append(messageEntityProto.FieldsVisibility, false)
	messageEntityProto.Fields = append(messageEntityProto.Fields, "M_entitiesPtr")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "undefined")
	messageEntityProto.FieldsDefaultValue = append(messageEntityProto.FieldsDefaultValue, "undefined")
	messageEntityProto.FieldsType = append(messageEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&messageEntityProto)

}

////////////////////////////////////////////////////////////////////////////////
//              			Notification
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_NotificationEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Notification
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesNotificationEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_NotificationEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesNotificationExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Notification).TYPENAME = "CargoEntities.Notification"
		object.(*CargoEntities.Notification).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Notification", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Notification).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Notification).UUID
			}
			return val.(*CargoEntities_NotificationEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_NotificationEntity)
	if object == nil {
		entity.object = new(CargoEntities.Notification)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Notification)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Notification"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_NotificationEntity) GetTypeName() string {
	return "CargoEntities.Notification"
}
func (this *CargoEntities_NotificationEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_NotificationEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_NotificationEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_NotificationEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_NotificationEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_NotificationEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_NotificationEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_NotificationEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_NotificationEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_NotificationEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_NotificationEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_NotificationEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_NotificationEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_NotificationEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_NotificationEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_NotificationEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_NotificationEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_NotificationEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_NotificationEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_NotificationEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_NotificationEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_NotificationEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_NotificationEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_NotificationEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_NotificationEntityPrototype() {

	var notificationEntityProto EntityPrototype
	notificationEntityProto.TypeName = "CargoEntities.Notification"
	notificationEntityProto.SuperTypeNames = append(notificationEntityProto.SuperTypeNames, "CargoEntities.Entity")
	notificationEntityProto.SuperTypeNames = append(notificationEntityProto.SuperTypeNames, "CargoEntities.Message")
	notificationEntityProto.Ids = append(notificationEntityProto.Ids, "UUID")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "UUID")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 0)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")
	notificationEntityProto.Indexs = append(notificationEntityProto.Indexs, "ParentUuid")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "ParentUuid")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 1)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "ParentLnk")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 2)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	notificationEntityProto.Ids = append(notificationEntityProto.Ids, "M_id")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 3)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_id")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.ID")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 4)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_body")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")

	/** members of Notification **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 5)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_fromRef")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Account:Ref")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 6)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_toRef")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Account:Ref")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 7)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_type")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.string")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "")
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 8)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, true)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_code")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "xs.int")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "0")

	/** associations of Notification **/
	notificationEntityProto.FieldsOrder = append(notificationEntityProto.FieldsOrder, 9)
	notificationEntityProto.FieldsVisibility = append(notificationEntityProto.FieldsVisibility, false)
	notificationEntityProto.Fields = append(notificationEntityProto.Fields, "M_entitiesPtr")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsDefaultValue = append(notificationEntityProto.FieldsDefaultValue, "undefined")
	notificationEntityProto.FieldsType = append(notificationEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&notificationEntityProto)

}

/** Create **/
func (this *CargoEntities_NotificationEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Notification **/
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_code")

	/** associations of Notification **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var NotificationInfo []interface{}

	NotificationInfo = append(NotificationInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		NotificationInfo = append(NotificationInfo, this.GetParentPtr().GetUuid())
		NotificationInfo = append(NotificationInfo, this.GetParentLnk())
	} else {
		NotificationInfo = append(NotificationInfo, "")
		NotificationInfo = append(NotificationInfo, "")
	}

	/** members of Entity **/
	NotificationInfo = append(NotificationInfo, this.object.M_id)

	/** members of Message **/
	NotificationInfo = append(NotificationInfo, this.object.M_body)

	/** members of Notification **/

	/** Save fromRef type Account **/
	if len(this.object.M_fromRef) > 0 {
		NotificationInfo = append(NotificationInfo, this.object.M_fromRef)
	} else {
		NotificationInfo = append(NotificationInfo, "")
	}

	/** Save toRef type Account **/
	if len(this.object.M_toRef) > 0 {
		NotificationInfo = append(NotificationInfo, this.object.M_toRef)
	} else {
		NotificationInfo = append(NotificationInfo, "")
	}
	NotificationInfo = append(NotificationInfo, this.object.M_type)
	NotificationInfo = append(NotificationInfo, this.object.M_code)

	/** associations of Notification **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		NotificationInfo = append(NotificationInfo, this.object.M_entitiesPtr)
	} else {
		NotificationInfo = append(NotificationInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), NotificationInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), NotificationInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_NotificationEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_NotificationEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of Notification **/
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_type")
	query.Fields = append(query.Fields, "M_code")

	/** associations of Notification **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Notification...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Notification)
		this.object.TYPENAME = "CargoEntities.Notification"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Message **/

		/** body **/
		if results[0][4] != nil {
			this.object.M_body = results[0][4].(string)
		}

		/** members of Notification **/

		/** fromRef **/
		if results[0][5] != nil {
			id := results[0][5].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_fromRef = id
				GetServer().GetEntityManager().appendReference("fromRef", this.object.UUID, id_)
			}
		}

		/** toRef **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_toRef = id
				GetServer().GetEntityManager().appendReference("toRef", this.object.UUID, id_)
			}
		}

		/** type **/
		if results[0][7] != nil {
			this.object.M_type = results[0][7].(string)
		}

		/** code **/
		if results[0][8] != nil {
			this.object.M_code = results[0][8].(int)
		}

		/** associations of Notification **/

		/** entitiesPtr **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesNotificationEntityFromObject(object *CargoEntities.Notification) *CargoEntities_NotificationEntity {
	return this.NewCargoEntitiesNotificationEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_NotificationEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesNotificationExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Notification"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_NotificationEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_NotificationEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			TextMessage
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_TextMessageEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.TextMessage
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesTextMessageEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_TextMessageEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesTextMessageExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.TextMessage).TYPENAME = "CargoEntities.TextMessage"
		object.(*CargoEntities.TextMessage).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.TextMessage", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.TextMessage).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.TextMessage).UUID
			}
			return val.(*CargoEntities_TextMessageEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_TextMessageEntity)
	if object == nil {
		entity.object = new(CargoEntities.TextMessage)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.TextMessage)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.TextMessage"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_TextMessageEntity) GetTypeName() string {
	return "CargoEntities.TextMessage"
}
func (this *CargoEntities_TextMessageEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_TextMessageEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_TextMessageEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_TextMessageEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_TextMessageEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_TextMessageEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_TextMessageEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_TextMessageEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_TextMessageEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_TextMessageEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_TextMessageEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_TextMessageEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_TextMessageEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_TextMessageEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_TextMessageEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_TextMessageEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_TextMessageEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_TextMessageEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_TextMessageEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_TextMessageEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_TextMessageEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_TextMessageEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_TextMessageEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_TextMessageEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_TextMessageEntityPrototype() {

	var textMessageEntityProto EntityPrototype
	textMessageEntityProto.TypeName = "CargoEntities.TextMessage"
	textMessageEntityProto.SuperTypeNames = append(textMessageEntityProto.SuperTypeNames, "CargoEntities.Entity")
	textMessageEntityProto.SuperTypeNames = append(textMessageEntityProto.SuperTypeNames, "CargoEntities.Message")
	textMessageEntityProto.Ids = append(textMessageEntityProto.Ids, "UUID")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "UUID")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 0)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")
	textMessageEntityProto.Indexs = append(textMessageEntityProto.Indexs, "ParentUuid")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "ParentUuid")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 1)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "ParentLnk")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 2)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	textMessageEntityProto.Ids = append(textMessageEntityProto.Ids, "M_id")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 3)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_id")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.ID")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** members of Message **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 4)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_body")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** members of TextMessage **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 5)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_creationTime")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.date")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "new Date()")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 6)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_fromRef")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Account:Ref")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 7)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_toRef")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Account:Ref")
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 8)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, true)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_title")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "xs.string")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "")

	/** associations of TextMessage **/
	textMessageEntityProto.FieldsOrder = append(textMessageEntityProto.FieldsOrder, 9)
	textMessageEntityProto.FieldsVisibility = append(textMessageEntityProto.FieldsVisibility, false)
	textMessageEntityProto.Fields = append(textMessageEntityProto.Fields, "M_entitiesPtr")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsDefaultValue = append(textMessageEntityProto.FieldsDefaultValue, "undefined")
	textMessageEntityProto.FieldsType = append(textMessageEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&textMessageEntityProto)

}

/** Create **/
func (this *CargoEntities_TextMessageEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of TextMessage **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_title")

	/** associations of TextMessage **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var TextMessageInfo []interface{}

	TextMessageInfo = append(TextMessageInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		TextMessageInfo = append(TextMessageInfo, this.GetParentPtr().GetUuid())
		TextMessageInfo = append(TextMessageInfo, this.GetParentLnk())
	} else {
		TextMessageInfo = append(TextMessageInfo, "")
		TextMessageInfo = append(TextMessageInfo, "")
	}

	/** members of Entity **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_id)

	/** members of Message **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_body)

	/** members of TextMessage **/
	TextMessageInfo = append(TextMessageInfo, this.object.M_creationTime)

	/** Save fromRef type Account **/
	if len(this.object.M_fromRef) > 0 {
		TextMessageInfo = append(TextMessageInfo, this.object.M_fromRef)
	} else {
		TextMessageInfo = append(TextMessageInfo, "")
	}

	/** Save toRef type Account **/
	if len(this.object.M_toRef) > 0 {
		TextMessageInfo = append(TextMessageInfo, this.object.M_toRef)
	} else {
		TextMessageInfo = append(TextMessageInfo, "")
	}
	TextMessageInfo = append(TextMessageInfo, this.object.M_title)

	/** associations of TextMessage **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		TextMessageInfo = append(TextMessageInfo, this.object.M_entitiesPtr)
	} else {
		TextMessageInfo = append(TextMessageInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), TextMessageInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), TextMessageInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_TextMessageEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_TextMessageEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Message **/
	query.Fields = append(query.Fields, "M_body")

	/** members of TextMessage **/
	query.Fields = append(query.Fields, "M_creationTime")
	query.Fields = append(query.Fields, "M_fromRef")
	query.Fields = append(query.Fields, "M_toRef")
	query.Fields = append(query.Fields, "M_title")

	/** associations of TextMessage **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of TextMessage...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.TextMessage)
		this.object.TYPENAME = "CargoEntities.TextMessage"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Message **/

		/** body **/
		if results[0][4] != nil {
			this.object.M_body = results[0][4].(string)
		}

		/** members of TextMessage **/

		/** creationTime **/
		if results[0][5] != nil {
			this.object.M_creationTime = results[0][5].(int64)
		}

		/** fromRef **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_fromRef = id
				GetServer().GetEntityManager().appendReference("fromRef", this.object.UUID, id_)
			}
		}

		/** toRef **/
		if results[0][7] != nil {
			id := results[0][7].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_toRef = id
				GetServer().GetEntityManager().appendReference("toRef", this.object.UUID, id_)
			}
		}

		/** title **/
		if results[0][8] != nil {
			this.object.M_title = results[0][8].(string)
		}

		/** associations of TextMessage **/

		/** entitiesPtr **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesTextMessageEntityFromObject(object *CargoEntities.TextMessage) *CargoEntities_TextMessageEntity {
	return this.NewCargoEntitiesTextMessageEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_TextMessageEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesTextMessageExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.TextMessage"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_TextMessageEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_TextMessageEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Session
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_SessionEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Session
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesSessionEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_SessionEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesSessionExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Session).TYPENAME = "CargoEntities.Session"
		object.(*CargoEntities.Session).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Session", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Session).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Session).UUID
			}
			return val.(*CargoEntities_SessionEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_SessionEntity)
	if object == nil {
		entity.object = new(CargoEntities.Session)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Session)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Session"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_SessionEntity) GetTypeName() string {
	return "CargoEntities.Session"
}
func (this *CargoEntities_SessionEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_SessionEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_SessionEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_SessionEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_SessionEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_SessionEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_SessionEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_SessionEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_SessionEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_SessionEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_SessionEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_SessionEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_SessionEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_SessionEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_SessionEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_SessionEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_SessionEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_SessionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_SessionEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_SessionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_SessionEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_SessionEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_SessionEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Session"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_SessionEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_SessionEntityPrototype() {

	var sessionEntityProto EntityPrototype
	sessionEntityProto.TypeName = "CargoEntities.Session"
	sessionEntityProto.Ids = append(sessionEntityProto.Ids, "UUID")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "UUID")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 0)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")
	sessionEntityProto.Indexs = append(sessionEntityProto.Indexs, "ParentUuid")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "ParentUuid")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 1)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "ParentLnk")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.string")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 2)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")

	/** members of Session **/
	sessionEntityProto.Ids = append(sessionEntityProto.Ids, "M_id")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 3)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_id")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.ID")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 4)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_startTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.date")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "new Date()")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 5)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_endTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.date")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "new Date()")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 6)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_statusTime")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "xs.date")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "new Date()")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 7)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_sessionState")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "1")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "enum:SessionState_Online:SessionState_Away:SessionState_Offline")
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 8)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, true)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_computerRef")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "CargoEntities.Computer:Ref")

	/** associations of Session **/
	sessionEntityProto.FieldsOrder = append(sessionEntityProto.FieldsOrder, 9)
	sessionEntityProto.FieldsVisibility = append(sessionEntityProto.FieldsVisibility, false)
	sessionEntityProto.Fields = append(sessionEntityProto.Fields, "M_accountPtr")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsDefaultValue = append(sessionEntityProto.FieldsDefaultValue, "undefined")
	sessionEntityProto.FieldsType = append(sessionEntityProto.FieldsType, "CargoEntities.Account:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&sessionEntityProto)

}

/** Create **/
func (this *CargoEntities_SessionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Session"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Session **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_startTime")
	query.Fields = append(query.Fields, "M_endTime")
	query.Fields = append(query.Fields, "M_statusTime")
	query.Fields = append(query.Fields, "M_sessionState")
	query.Fields = append(query.Fields, "M_computerRef")

	/** associations of Session **/
	query.Fields = append(query.Fields, "M_accountPtr")

	var SessionInfo []interface{}

	SessionInfo = append(SessionInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		SessionInfo = append(SessionInfo, this.GetParentPtr().GetUuid())
		SessionInfo = append(SessionInfo, this.GetParentLnk())
	} else {
		SessionInfo = append(SessionInfo, "")
		SessionInfo = append(SessionInfo, "")
	}

	/** members of Session **/
	SessionInfo = append(SessionInfo, this.object.M_id)
	SessionInfo = append(SessionInfo, this.object.M_startTime)
	SessionInfo = append(SessionInfo, this.object.M_endTime)
	SessionInfo = append(SessionInfo, this.object.M_statusTime)

	/** Save sessionState type SessionState **/
	if this.object.M_sessionState == CargoEntities.SessionState_Online {
		SessionInfo = append(SessionInfo, 0)
	} else if this.object.M_sessionState == CargoEntities.SessionState_Away {
		SessionInfo = append(SessionInfo, 1)
	} else if this.object.M_sessionState == CargoEntities.SessionState_Offline {
		SessionInfo = append(SessionInfo, 2)
	} else {
		SessionInfo = append(SessionInfo, 0)
	}

	/** Save computerRef type Computer **/
	if len(this.object.M_computerRef) > 0 {
		SessionInfo = append(SessionInfo, this.object.M_computerRef)
	} else {
		SessionInfo = append(SessionInfo, "")
	}

	/** associations of Session **/

	/** Save account type Account **/
	if len(this.object.M_accountPtr) > 0 {
		SessionInfo = append(SessionInfo, this.object.M_accountPtr)
	} else {
		SessionInfo = append(SessionInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), SessionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), SessionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_SessionEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_SessionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Session"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Session **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_startTime")
	query.Fields = append(query.Fields, "M_endTime")
	query.Fields = append(query.Fields, "M_statusTime")
	query.Fields = append(query.Fields, "M_sessionState")
	query.Fields = append(query.Fields, "M_computerRef")

	/** associations of Session **/
	query.Fields = append(query.Fields, "M_accountPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Session...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Session)
		this.object.TYPENAME = "CargoEntities.Session"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Session **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** startTime **/
		if results[0][4] != nil {
			this.object.M_startTime = results[0][4].(int64)
		}

		/** endTime **/
		if results[0][5] != nil {
			this.object.M_endTime = results[0][5].(int64)
		}

		/** statusTime **/
		if results[0][6] != nil {
			this.object.M_statusTime = results[0][6].(int64)
		}

		/** sessionState **/
		if results[0][7] != nil {
			enumIndex := results[0][7].(int)
			if enumIndex == 0 {
				this.object.M_sessionState = CargoEntities.SessionState_Online
			} else if enumIndex == 1 {
				this.object.M_sessionState = CargoEntities.SessionState_Away
			} else if enumIndex == 2 {
				this.object.M_sessionState = CargoEntities.SessionState_Offline
			}
		}

		/** computerRef **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Computer"
				id_ := refTypeName + "$$" + id
				this.object.M_computerRef = id
				GetServer().GetEntityManager().appendReference("computerRef", this.object.UUID, id_)
			}
		}

		/** associations of Session **/

		/** accountPtr **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Account"
				id_ := refTypeName + "$$" + id
				this.object.M_accountPtr = id
				GetServer().GetEntityManager().appendReference("accountPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesSessionEntityFromObject(object *CargoEntities.Session) *CargoEntities_SessionEntity {
	return this.NewCargoEntitiesSessionEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_SessionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesSessionExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Session"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_SessionEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_SessionEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Role
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_RoleEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Role
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesRoleEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_RoleEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesRoleExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Role).TYPENAME = "CargoEntities.Role"
		object.(*CargoEntities.Role).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Role", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Role).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Role).UUID
			}
			return val.(*CargoEntities_RoleEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_RoleEntity)
	if object == nil {
		entity.object = new(CargoEntities.Role)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Role)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Role"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_RoleEntity) GetTypeName() string {
	return "CargoEntities.Role"
}
func (this *CargoEntities_RoleEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_RoleEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_RoleEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_RoleEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_RoleEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_RoleEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_RoleEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_RoleEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_RoleEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_RoleEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_RoleEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_RoleEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_RoleEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_RoleEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_RoleEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_RoleEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_RoleEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_RoleEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_RoleEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_RoleEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_RoleEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_RoleEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_RoleEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Role"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_RoleEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_RoleEntityPrototype() {

	var roleEntityProto EntityPrototype
	roleEntityProto.TypeName = "CargoEntities.Role"
	roleEntityProto.Ids = append(roleEntityProto.Ids, "UUID")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "UUID")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 0)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")
	roleEntityProto.Indexs = append(roleEntityProto.Indexs, "ParentUuid")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "ParentUuid")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 1)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")
	roleEntityProto.Fields = append(roleEntityProto.Fields, "ParentLnk")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.string")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 2)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")

	/** members of Role **/
	roleEntityProto.Ids = append(roleEntityProto.Ids, "M_id")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 3)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_id")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "xs.ID")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 4)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_accounts")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "[]")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]CargoEntities.Account:Ref")
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 5)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, true)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_actions")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "[]")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "[]CargoEntities.Action:Ref")

	/** associations of Role **/
	roleEntityProto.FieldsOrder = append(roleEntityProto.FieldsOrder, 6)
	roleEntityProto.FieldsVisibility = append(roleEntityProto.FieldsVisibility, false)
	roleEntityProto.Fields = append(roleEntityProto.Fields, "M_entitiesPtr")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsDefaultValue = append(roleEntityProto.FieldsDefaultValue, "undefined")
	roleEntityProto.FieldsType = append(roleEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&roleEntityProto)

}

/** Create **/
func (this *CargoEntities_RoleEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Role"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Role **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_accounts")
	query.Fields = append(query.Fields, "M_actions")

	/** associations of Role **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var RoleInfo []interface{}

	RoleInfo = append(RoleInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		RoleInfo = append(RoleInfo, this.GetParentPtr().GetUuid())
		RoleInfo = append(RoleInfo, this.GetParentLnk())
	} else {
		RoleInfo = append(RoleInfo, "")
		RoleInfo = append(RoleInfo, "")
	}

	/** members of Role **/
	RoleInfo = append(RoleInfo, this.object.M_id)

	/** Save accounts type Account **/
	accountsStr, _ := json.Marshal(this.object.M_accounts)
	RoleInfo = append(RoleInfo, string(accountsStr))

	/** Save actions type Action **/
	actionsStr, _ := json.Marshal(this.object.M_actions)
	RoleInfo = append(RoleInfo, string(actionsStr))

	/** associations of Role **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		RoleInfo = append(RoleInfo, this.object.M_entitiesPtr)
	} else {
		RoleInfo = append(RoleInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), RoleInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), RoleInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_RoleEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_RoleEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Role"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Role **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_accounts")
	query.Fields = append(query.Fields, "M_actions")

	/** associations of Role **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Role...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Role)
		this.object.TYPENAME = "CargoEntities.Role"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Role **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** accounts **/
		if results[0][4] != nil {
			idsStr := results[0][4].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Account"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_accounts = append(this.object.M_accounts, ids[i])
					GetServer().GetEntityManager().appendReference("accounts", this.object.UUID, id_)
				}
			}
		}

		/** actions **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Action"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_actions = append(this.object.M_actions, ids[i])
					GetServer().GetEntityManager().appendReference("actions", this.object.UUID, id_)
				}
			}
		}

		/** associations of Role **/

		/** entitiesPtr **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesRoleEntityFromObject(object *CargoEntities.Role) *CargoEntities_RoleEntity {
	return this.NewCargoEntitiesRoleEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_RoleEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesRoleExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Role"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_RoleEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_RoleEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Account
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_AccountEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Account
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesAccountEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_AccountEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesAccountExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Account).TYPENAME = "CargoEntities.Account"
		object.(*CargoEntities.Account).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Account", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Account).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Account).UUID
			}
			return val.(*CargoEntities_AccountEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_AccountEntity)
	if object == nil {
		entity.object = new(CargoEntities.Account)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Account)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Account"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_AccountEntity) GetTypeName() string {
	return "CargoEntities.Account"
}
func (this *CargoEntities_AccountEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_AccountEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_AccountEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_AccountEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_AccountEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_AccountEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_AccountEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_AccountEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_AccountEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_AccountEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_AccountEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_AccountEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_AccountEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_AccountEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_AccountEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_AccountEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_AccountEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_AccountEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_AccountEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_AccountEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_AccountEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_AccountEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_AccountEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Account"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_AccountEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_AccountEntityPrototype() {

	var accountEntityProto EntityPrototype
	accountEntityProto.TypeName = "CargoEntities.Account"
	accountEntityProto.SuperTypeNames = append(accountEntityProto.SuperTypeNames, "CargoEntities.Entity")
	accountEntityProto.Ids = append(accountEntityProto.Ids, "UUID")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "UUID")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 0)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.Indexs = append(accountEntityProto.Indexs, "ParentUuid")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "ParentUuid")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 1)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.Fields = append(accountEntityProto.Fields, "ParentLnk")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 2)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	accountEntityProto.Ids = append(accountEntityProto.Ids, "M_id")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 3)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_id")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.ID")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")

	/** members of Account **/
	accountEntityProto.Indexs = append(accountEntityProto.Indexs, "M_name")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 4)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_name")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 5)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_password")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 6)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_email")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "xs.string")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 7)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_sessions")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Session")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 8)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_messages")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Message")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 9)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_userRef")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "CargoEntities.User:Ref")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 10)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_rolesRef")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Role:Ref")
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 11)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, true)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_permissionsRef")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "[]")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "[]CargoEntities.Permission:Ref")

	/** associations of Account **/
	accountEntityProto.FieldsOrder = append(accountEntityProto.FieldsOrder, 12)
	accountEntityProto.FieldsVisibility = append(accountEntityProto.FieldsVisibility, false)
	accountEntityProto.Fields = append(accountEntityProto.Fields, "M_entitiesPtr")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsDefaultValue = append(accountEntityProto.FieldsDefaultValue, "undefined")
	accountEntityProto.FieldsType = append(accountEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&accountEntityProto)

}

/** Create **/
func (this *CargoEntities_AccountEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Account"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Account **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_password")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_sessions")
	query.Fields = append(query.Fields, "M_messages")
	query.Fields = append(query.Fields, "M_userRef")
	query.Fields = append(query.Fields, "M_rolesRef")
	query.Fields = append(query.Fields, "M_permissionsRef")

	/** associations of Account **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var AccountInfo []interface{}

	AccountInfo = append(AccountInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		AccountInfo = append(AccountInfo, this.GetParentPtr().GetUuid())
		AccountInfo = append(AccountInfo, this.GetParentLnk())
	} else {
		AccountInfo = append(AccountInfo, "")
		AccountInfo = append(AccountInfo, "")
	}

	/** members of Entity **/
	AccountInfo = append(AccountInfo, this.object.M_id)

	/** members of Account **/
	AccountInfo = append(AccountInfo, this.object.M_name)
	AccountInfo = append(AccountInfo, this.object.M_password)
	AccountInfo = append(AccountInfo, this.object.M_email)

	/** Save sessions type Session **/
	sessionsIds := make([]string, 0)
	lazy_sessions := this.lazyMap["M_sessions"] != nil && len(this.object.M_sessions) == 0
	if !lazy_sessions {
		for i := 0; i < len(this.object.M_sessions); i++ {
			sessionsEntity := GetServer().GetEntityManager().NewCargoEntitiesSessionEntity(this.GetUuid(), this.object.M_sessions[i].UUID, this.object.M_sessions[i])
			sessionsIds = append(sessionsIds, sessionsEntity.GetUuid())
			sessionsEntity.AppendReferenced("sessions", this)
			this.AppendChild("sessions", sessionsEntity)
			if sessionsEntity.NeedSave() {
				sessionsEntity.SaveEntity()
			}
		}
	} else {
		sessionsIds = this.lazyMap["M_sessions"].([]string)
	}
	sessionsStr, _ := json.Marshal(sessionsIds)
	AccountInfo = append(AccountInfo, string(sessionsStr))

	/** Save messages type Message **/
	messagesIds := make([]string, 0)
	lazy_messages := this.lazyMap["M_messages"] != nil && len(this.object.M_messages) == 0
	if !lazy_messages {
		for i := 0; i < len(this.object.M_messages); i++ {
			switch v := this.object.M_messages[i].(type) {
			case *CargoEntities.Error:
				messagesEntity := GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), v.UUID, v)
				messagesIds = append(messagesIds, messagesEntity.GetUuid())
				messagesEntity.AppendReferenced("messages", this)
				this.AppendChild("messages", messagesEntity)
				if messagesEntity.NeedSave() {
					messagesEntity.SaveEntity()
				}
			case *CargoEntities.Notification:
				messagesEntity := GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), v.UUID, v)
				messagesIds = append(messagesIds, messagesEntity.GetUuid())
				messagesEntity.AppendReferenced("messages", this)
				this.AppendChild("messages", messagesEntity)
				if messagesEntity.NeedSave() {
					messagesEntity.SaveEntity()
				}
			case *CargoEntities.TextMessage:
				messagesEntity := GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), v.UUID, v)
				messagesIds = append(messagesIds, messagesEntity.GetUuid())
				messagesEntity.AppendReferenced("messages", this)
				this.AppendChild("messages", messagesEntity)
				if messagesEntity.NeedSave() {
					messagesEntity.SaveEntity()
				}
			}
		}
	} else {
		messagesIds = this.lazyMap["M_messages"].([]string)
	}
	messagesStr, _ := json.Marshal(messagesIds)
	AccountInfo = append(AccountInfo, string(messagesStr))

	/** Save userRef type User **/
	if len(this.object.M_userRef) > 0 {
		AccountInfo = append(AccountInfo, this.object.M_userRef)
	} else {
		AccountInfo = append(AccountInfo, "")
	}

	/** Save rolesRef type Role **/
	rolesRefStr, _ := json.Marshal(this.object.M_rolesRef)
	AccountInfo = append(AccountInfo, string(rolesRefStr))

	/** Save permissionsRef type Permission **/
	permissionsRefStr, _ := json.Marshal(this.object.M_permissionsRef)
	AccountInfo = append(AccountInfo, string(permissionsRefStr))

	/** associations of Account **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		AccountInfo = append(AccountInfo, this.object.M_entitiesPtr)
	} else {
		AccountInfo = append(AccountInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), AccountInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), AccountInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_AccountEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_AccountEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Account"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Account **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_password")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_sessions")
	query.Fields = append(query.Fields, "M_messages")
	query.Fields = append(query.Fields, "M_userRef")
	query.Fields = append(query.Fields, "M_rolesRef")
	query.Fields = append(query.Fields, "M_permissionsRef")

	/** associations of Account **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Account...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Account)
		this.object.TYPENAME = "CargoEntities.Account"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Account **/

		/** name **/
		if results[0][4] != nil {
			this.object.M_name = results[0][4].(string)
		}

		/** password **/
		if results[0][5] != nil {
			this.object.M_password = results[0][5].(string)
		}

		/** email **/
		if results[0][6] != nil {
			this.object.M_email = results[0][6].(string)
		}

		/** sessions **/
		if results[0][7] != nil {
			uuidsStr := results[0][7].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var sessionsEntity *CargoEntities_SessionEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							sessionsEntity = instance.(*CargoEntities_SessionEntity)
						} else {
							sessionsEntity = GetServer().GetEntityManager().NewCargoEntitiesSessionEntity(this.GetUuid(), uuids[i], nil)
							sessionsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(sessionsEntity)
						}
						sessionsEntity.AppendReferenced("sessions", this)
						this.AppendChild("sessions", sessionsEntity)
					}
				} else {
					this.lazyMap["M_sessions"] = uuids
				}
			}
		}

		/** messages **/
		if results[0][8] != nil {
			uuidsStr := results[0][8].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					typeName := uuids[i][0:strings.Index(uuids[i], "%")]
					if err != nil {
						log.Println("type ", typeName, " not found!")
						return err
					}
					if typeName == "CargoEntities.Error" {
						if len(uuids[i]) > 0 {
							var messagesEntity *CargoEntities_ErrorEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								messagesEntity = instance.(*CargoEntities_ErrorEntity)
							} else {
								messagesEntity = GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), uuids[i], nil)
								messagesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(messagesEntity)
							}
							messagesEntity.AppendReferenced("messages", this)
							this.AppendChild("messages", messagesEntity)
						}
					} else if typeName == "CargoEntities.Notification" {
						if len(uuids[i]) > 0 {
							var messagesEntity *CargoEntities_NotificationEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								messagesEntity = instance.(*CargoEntities_NotificationEntity)
							} else {
								messagesEntity = GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), uuids[i], nil)
								messagesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(messagesEntity)
							}
							messagesEntity.AppendReferenced("messages", this)
							this.AppendChild("messages", messagesEntity)
						}
					} else if typeName == "CargoEntities.TextMessage" {
						if len(uuids[i]) > 0 {
							var messagesEntity *CargoEntities_TextMessageEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								messagesEntity = instance.(*CargoEntities_TextMessageEntity)
							} else {
								messagesEntity = GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), uuids[i], nil)
								messagesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(messagesEntity)
							}
							messagesEntity.AppendReferenced("messages", this)
							this.AppendChild("messages", messagesEntity)
						}
					}
				} else {
					this.lazyMap["M_messages"] = uuids
				}
			}
		}

		/** userRef **/
		if results[0][9] != nil {
			id := results[0][9].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.User"
				id_ := refTypeName + "$$" + id
				this.object.M_userRef = id
				GetServer().GetEntityManager().appendReference("userRef", this.object.UUID, id_)
			}
		}

		/** rolesRef **/
		if results[0][10] != nil {
			idsStr := results[0][10].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Role"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_rolesRef = append(this.object.M_rolesRef, ids[i])
					GetServer().GetEntityManager().appendReference("rolesRef", this.object.UUID, id_)
				}
			}
		}

		/** permissionsRef **/
		if results[0][11] != nil {
			idsStr := results[0][11].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Permission"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_permissionsRef = append(this.object.M_permissionsRef, ids[i])
					GetServer().GetEntityManager().appendReference("permissionsRef", this.object.UUID, id_)
				}
			}
		}

		/** associations of Account **/

		/** entitiesPtr **/
		if results[0][12] != nil {
			id := results[0][12].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesAccountEntityFromObject(object *CargoEntities.Account) *CargoEntities_AccountEntity {
	return this.NewCargoEntitiesAccountEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_AccountEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesAccountExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Account"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		var query EntityQuery
		query.TypeName = "CargoEntities.Account"
		query.Indexs = append(query.Indexs, "M_name="+val)
		query.Fields = append(query.Fields, "UUID")
		var fieldsType []interface{} // not use...
		var params []interface{}
		queryStr, _ := json.Marshal(query)
		results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
		if err != nil || len(results) == 0 {
			return ""
		}
		return results[0][0].(string)
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_AccountEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_AccountEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Computer
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_ComputerEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Computer
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesComputerEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_ComputerEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesComputerExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Computer).TYPENAME = "CargoEntities.Computer"
		object.(*CargoEntities.Computer).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Computer", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Computer).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Computer).UUID
			}
			return val.(*CargoEntities_ComputerEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_ComputerEntity)
	if object == nil {
		entity.object = new(CargoEntities.Computer)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Computer)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Computer"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_ComputerEntity) GetTypeName() string {
	return "CargoEntities.Computer"
}
func (this *CargoEntities_ComputerEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_ComputerEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_ComputerEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_ComputerEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_ComputerEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_ComputerEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_ComputerEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_ComputerEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_ComputerEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_ComputerEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_ComputerEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_ComputerEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_ComputerEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_ComputerEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_ComputerEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_ComputerEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_ComputerEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_ComputerEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_ComputerEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_ComputerEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_ComputerEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_ComputerEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_ComputerEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_ComputerEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_ComputerEntityPrototype() {

	var computerEntityProto EntityPrototype
	computerEntityProto.TypeName = "CargoEntities.Computer"
	computerEntityProto.SuperTypeNames = append(computerEntityProto.SuperTypeNames, "CargoEntities.Entity")
	computerEntityProto.Ids = append(computerEntityProto.Ids, "UUID")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "UUID")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 0)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.Indexs = append(computerEntityProto.Indexs, "ParentUuid")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "ParentUuid")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 1)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.Fields = append(computerEntityProto.Fields, "ParentLnk")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 2)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	computerEntityProto.Ids = append(computerEntityProto.Ids, "M_id")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 3)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_id")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.ID")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")

	/** members of Computer **/
	computerEntityProto.Indexs = append(computerEntityProto.Indexs, "M_name")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 4)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_name")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 5)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_ipv4")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "xs.string")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 6)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_osType")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "1")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "enum:OsType_Unknown:OsType_Linux:OsType_Windows7:OsType_Windows8:OsType_Windows10:OsType_OSX:OsType_IOS")
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 7)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, true)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_platformType")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "1")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "enum:PlatformType_Unknown:PlatformType_Tablet:PlatformType_Phone:PlatformType_Desktop:PlatformType_Laptop")

	/** associations of Computer **/
	computerEntityProto.FieldsOrder = append(computerEntityProto.FieldsOrder, 8)
	computerEntityProto.FieldsVisibility = append(computerEntityProto.FieldsVisibility, false)
	computerEntityProto.Fields = append(computerEntityProto.Fields, "M_entitiesPtr")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "undefined")
	computerEntityProto.FieldsDefaultValue = append(computerEntityProto.FieldsDefaultValue, "undefined")
	computerEntityProto.FieldsType = append(computerEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&computerEntityProto)

}

/** Create **/
func (this *CargoEntities_ComputerEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Computer **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_osType")
	query.Fields = append(query.Fields, "M_platformType")

	/** associations of Computer **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var ComputerInfo []interface{}

	ComputerInfo = append(ComputerInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		ComputerInfo = append(ComputerInfo, this.GetParentPtr().GetUuid())
		ComputerInfo = append(ComputerInfo, this.GetParentLnk())
	} else {
		ComputerInfo = append(ComputerInfo, "")
		ComputerInfo = append(ComputerInfo, "")
	}

	/** members of Entity **/
	ComputerInfo = append(ComputerInfo, this.object.M_id)

	/** members of Computer **/
	ComputerInfo = append(ComputerInfo, this.object.M_name)
	ComputerInfo = append(ComputerInfo, this.object.M_ipv4)

	/** Save osType type OsType **/
	if this.object.M_osType == CargoEntities.OsType_Unknown {
		ComputerInfo = append(ComputerInfo, 0)
	} else if this.object.M_osType == CargoEntities.OsType_Linux {
		ComputerInfo = append(ComputerInfo, 1)
	} else if this.object.M_osType == CargoEntities.OsType_Windows7 {
		ComputerInfo = append(ComputerInfo, 2)
	} else if this.object.M_osType == CargoEntities.OsType_Windows8 {
		ComputerInfo = append(ComputerInfo, 3)
	} else if this.object.M_osType == CargoEntities.OsType_Windows10 {
		ComputerInfo = append(ComputerInfo, 4)
	} else if this.object.M_osType == CargoEntities.OsType_OSX {
		ComputerInfo = append(ComputerInfo, 5)
	} else if this.object.M_osType == CargoEntities.OsType_IOS {
		ComputerInfo = append(ComputerInfo, 6)
	} else {
		ComputerInfo = append(ComputerInfo, 0)
	}

	/** Save platformType type PlatformType **/
	if this.object.M_platformType == CargoEntities.PlatformType_Unknown {
		ComputerInfo = append(ComputerInfo, 0)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Tablet {
		ComputerInfo = append(ComputerInfo, 1)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Phone {
		ComputerInfo = append(ComputerInfo, 2)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Desktop {
		ComputerInfo = append(ComputerInfo, 3)
	} else if this.object.M_platformType == CargoEntities.PlatformType_Laptop {
		ComputerInfo = append(ComputerInfo, 4)
	} else {
		ComputerInfo = append(ComputerInfo, 0)
	}

	/** associations of Computer **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		ComputerInfo = append(ComputerInfo, this.object.M_entitiesPtr)
	} else {
		ComputerInfo = append(ComputerInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), ComputerInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), ComputerInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_ComputerEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_ComputerEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Computer **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_ipv4")
	query.Fields = append(query.Fields, "M_osType")
	query.Fields = append(query.Fields, "M_platformType")

	/** associations of Computer **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Computer...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Computer)
		this.object.TYPENAME = "CargoEntities.Computer"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Computer **/

		/** name **/
		if results[0][4] != nil {
			this.object.M_name = results[0][4].(string)
		}

		/** ipv4 **/
		if results[0][5] != nil {
			this.object.M_ipv4 = results[0][5].(string)
		}

		/** osType **/
		if results[0][6] != nil {
			enumIndex := results[0][6].(int)
			if enumIndex == 0 {
				this.object.M_osType = CargoEntities.OsType_Unknown
			} else if enumIndex == 1 {
				this.object.M_osType = CargoEntities.OsType_Linux
			} else if enumIndex == 2 {
				this.object.M_osType = CargoEntities.OsType_Windows7
			} else if enumIndex == 3 {
				this.object.M_osType = CargoEntities.OsType_Windows8
			} else if enumIndex == 4 {
				this.object.M_osType = CargoEntities.OsType_Windows10
			} else if enumIndex == 5 {
				this.object.M_osType = CargoEntities.OsType_OSX
			} else if enumIndex == 6 {
				this.object.M_osType = CargoEntities.OsType_IOS
			}
		}

		/** platformType **/
		if results[0][7] != nil {
			enumIndex := results[0][7].(int)
			if enumIndex == 0 {
				this.object.M_platformType = CargoEntities.PlatformType_Unknown
			} else if enumIndex == 1 {
				this.object.M_platformType = CargoEntities.PlatformType_Tablet
			} else if enumIndex == 2 {
				this.object.M_platformType = CargoEntities.PlatformType_Phone
			} else if enumIndex == 3 {
				this.object.M_platformType = CargoEntities.PlatformType_Desktop
			} else if enumIndex == 4 {
				this.object.M_platformType = CargoEntities.PlatformType_Laptop
			}
		}

		/** associations of Computer **/

		/** entitiesPtr **/
		if results[0][8] != nil {
			id := results[0][8].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesComputerEntityFromObject(object *CargoEntities.Computer) *CargoEntities_ComputerEntity {
	return this.NewCargoEntitiesComputerEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_ComputerEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesComputerExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Computer"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_ComputerEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_ComputerEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Permission
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_PermissionEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Permission
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesPermissionEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_PermissionEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesPermissionExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Permission).TYPENAME = "CargoEntities.Permission"
		object.(*CargoEntities.Permission).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Permission", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Permission).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Permission).UUID
			}
			return val.(*CargoEntities_PermissionEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_PermissionEntity)
	if object == nil {
		entity.object = new(CargoEntities.Permission)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Permission)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Permission"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_PermissionEntity) GetTypeName() string {
	return "CargoEntities.Permission"
}
func (this *CargoEntities_PermissionEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_PermissionEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_PermissionEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_PermissionEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_PermissionEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_PermissionEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_PermissionEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_PermissionEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_PermissionEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_PermissionEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_PermissionEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_PermissionEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_PermissionEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_PermissionEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_PermissionEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_PermissionEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_PermissionEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_PermissionEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_PermissionEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_PermissionEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_PermissionEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_PermissionEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_PermissionEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_PermissionEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_PermissionEntityPrototype() {

	var permissionEntityProto EntityPrototype
	permissionEntityProto.TypeName = "CargoEntities.Permission"
	permissionEntityProto.Ids = append(permissionEntityProto.Ids, "UUID")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "UUID")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 0)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")
	permissionEntityProto.Indexs = append(permissionEntityProto.Indexs, "ParentUuid")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "ParentUuid")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 1)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "ParentLnk")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.string")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 2)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")

	/** members of Permission **/
	permissionEntityProto.Ids = append(permissionEntityProto.Ids, "M_id")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 3)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_id")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.ID")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 4)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_types")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "xs.int")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "0")
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 5)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, true)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_accountsRef")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "undefined")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "[]")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "[]CargoEntities.Account:Ref")

	/** associations of Permission **/
	permissionEntityProto.FieldsOrder = append(permissionEntityProto.FieldsOrder, 6)
	permissionEntityProto.FieldsVisibility = append(permissionEntityProto.FieldsVisibility, false)
	permissionEntityProto.Fields = append(permissionEntityProto.Fields, "M_entitiesPtr")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "undefined")
	permissionEntityProto.FieldsDefaultValue = append(permissionEntityProto.FieldsDefaultValue, "undefined")
	permissionEntityProto.FieldsType = append(permissionEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&permissionEntityProto)

}

/** Create **/
func (this *CargoEntities_PermissionEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Permission **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_types")
	query.Fields = append(query.Fields, "M_accountsRef")

	/** associations of Permission **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var PermissionInfo []interface{}

	PermissionInfo = append(PermissionInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		PermissionInfo = append(PermissionInfo, this.GetParentPtr().GetUuid())
		PermissionInfo = append(PermissionInfo, this.GetParentLnk())
	} else {
		PermissionInfo = append(PermissionInfo, "")
		PermissionInfo = append(PermissionInfo, "")
	}

	/** members of Permission **/
	PermissionInfo = append(PermissionInfo, this.object.M_id)
	PermissionInfo = append(PermissionInfo, this.object.M_types)

	/** Save accountsRef type Account **/
	accountsRefStr, _ := json.Marshal(this.object.M_accountsRef)
	PermissionInfo = append(PermissionInfo, string(accountsRefStr))

	/** associations of Permission **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		PermissionInfo = append(PermissionInfo, this.object.M_entitiesPtr)
	} else {
		PermissionInfo = append(PermissionInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), PermissionInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), PermissionInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_PermissionEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_PermissionEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Permission **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_types")
	query.Fields = append(query.Fields, "M_accountsRef")

	/** associations of Permission **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Permission...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Permission)
		this.object.TYPENAME = "CargoEntities.Permission"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Permission **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** types **/
		if results[0][4] != nil {
			this.object.M_types = results[0][4].(int)
		}

		/** accountsRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Account"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_accountsRef = append(this.object.M_accountsRef, ids[i])
					GetServer().GetEntityManager().appendReference("accountsRef", this.object.UUID, id_)
				}
			}
		}

		/** associations of Permission **/

		/** entitiesPtr **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesPermissionEntityFromObject(object *CargoEntities.Permission) *CargoEntities_PermissionEntity {
	return this.NewCargoEntitiesPermissionEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_PermissionEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesPermissionExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Permission"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_PermissionEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_PermissionEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			File
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_FileEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.File
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesFileEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_FileEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesFileExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.File).TYPENAME = "CargoEntities.File"
		object.(*CargoEntities.File).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.File", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.File).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.File).UUID
			}
			return val.(*CargoEntities_FileEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_FileEntity)
	if object == nil {
		entity.object = new(CargoEntities.File)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.File)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.File"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_FileEntity) GetTypeName() string {
	return "CargoEntities.File"
}
func (this *CargoEntities_FileEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_FileEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_FileEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_FileEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_FileEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_FileEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_FileEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_FileEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_FileEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_FileEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_FileEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_FileEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_FileEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_FileEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_FileEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_FileEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_FileEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_FileEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_FileEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_FileEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_FileEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_FileEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_FileEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.File"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_FileEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_FileEntityPrototype() {

	var fileEntityProto EntityPrototype
	fileEntityProto.TypeName = "CargoEntities.File"
	fileEntityProto.SuperTypeNames = append(fileEntityProto.SuperTypeNames, "CargoEntities.Entity")
	fileEntityProto.Ids = append(fileEntityProto.Ids, "UUID")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "UUID")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 0)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.Indexs = append(fileEntityProto.Indexs, "ParentUuid")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "ParentUuid")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 1)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.Fields = append(fileEntityProto.Fields, "ParentLnk")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 2)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	fileEntityProto.Ids = append(fileEntityProto.Ids, "M_id")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 3)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_id")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.ID")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")

	/** members of File **/
	fileEntityProto.Indexs = append(fileEntityProto.Indexs, "M_name")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 4)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_name")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 5)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_path")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 6)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_size")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.int")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "0")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 7)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_modeTime")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.date")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "new Date()")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 8)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_isDir")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.boolean")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "false")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 9)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_checksum")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 10)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_data")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 11)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_thumbnail")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 12)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_mime")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "xs.string")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 13)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_files")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "[]")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "[]CargoEntities.File")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 14)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, true)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_fileType")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "1")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "enum:FileType_DbFile:FileType_DiskFile")

	/** associations of File **/
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 15)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_parentDirPtr")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "CargoEntities.File:Ref")
	fileEntityProto.FieldsOrder = append(fileEntityProto.FieldsOrder, 16)
	fileEntityProto.FieldsVisibility = append(fileEntityProto.FieldsVisibility, false)
	fileEntityProto.Fields = append(fileEntityProto.Fields, "M_entitiesPtr")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsDefaultValue = append(fileEntityProto.FieldsDefaultValue, "undefined")
	fileEntityProto.FieldsType = append(fileEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&fileEntityProto)

}

/** Create **/
func (this *CargoEntities_FileEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.File"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of File **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_path")
	query.Fields = append(query.Fields, "M_size")
	query.Fields = append(query.Fields, "M_modeTime")
	query.Fields = append(query.Fields, "M_isDir")
	query.Fields = append(query.Fields, "M_checksum")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_thumbnail")
	query.Fields = append(query.Fields, "M_mime")
	query.Fields = append(query.Fields, "M_files")
	query.Fields = append(query.Fields, "M_fileType")

	/** associations of File **/
	query.Fields = append(query.Fields, "M_parentDirPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var FileInfo []interface{}

	FileInfo = append(FileInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		FileInfo = append(FileInfo, this.GetParentPtr().GetUuid())
		FileInfo = append(FileInfo, this.GetParentLnk())
	} else {
		FileInfo = append(FileInfo, "")
		FileInfo = append(FileInfo, "")
	}

	/** members of Entity **/
	FileInfo = append(FileInfo, this.object.M_id)

	/** members of File **/
	FileInfo = append(FileInfo, this.object.M_name)
	FileInfo = append(FileInfo, this.object.M_path)
	FileInfo = append(FileInfo, this.object.M_size)
	FileInfo = append(FileInfo, this.object.M_modeTime)
	FileInfo = append(FileInfo, this.object.M_isDir)
	FileInfo = append(FileInfo, this.object.M_checksum)
	FileInfo = append(FileInfo, this.object.M_data)
	FileInfo = append(FileInfo, this.object.M_thumbnail)
	FileInfo = append(FileInfo, this.object.M_mime)

	/** Save files type File **/
	filesIds := make([]string, 0)
	lazy_files := this.lazyMap["M_files"] != nil && len(this.object.M_files) == 0
	if !lazy_files {
		for i := 0; i < len(this.object.M_files); i++ {
			filesEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), this.object.M_files[i].UUID, this.object.M_files[i])
			filesIds = append(filesIds, filesEntity.GetUuid())
			filesEntity.AppendReferenced("files", this)
			this.AppendChild("files", filesEntity)
			if filesEntity.NeedSave() {
				filesEntity.SaveEntity()
			}
		}
	} else {
		filesIds = this.lazyMap["M_files"].([]string)
	}
	filesStr, _ := json.Marshal(filesIds)
	FileInfo = append(FileInfo, string(filesStr))

	/** Save fileType type FileType **/
	if this.object.M_fileType == CargoEntities.FileType_DbFile {
		FileInfo = append(FileInfo, 0)
	} else if this.object.M_fileType == CargoEntities.FileType_DiskFile {
		FileInfo = append(FileInfo, 1)
	} else {
		FileInfo = append(FileInfo, 0)
	}

	/** associations of File **/

	/** Save parentDir type File **/
	if len(this.object.M_parentDirPtr) > 0 {
		FileInfo = append(FileInfo, this.object.M_parentDirPtr)
	} else {
		FileInfo = append(FileInfo, "")
	}

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		FileInfo = append(FileInfo, this.object.M_entitiesPtr)
	} else {
		FileInfo = append(FileInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), FileInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), FileInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_FileEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_FileEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.File"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of File **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_path")
	query.Fields = append(query.Fields, "M_size")
	query.Fields = append(query.Fields, "M_modeTime")
	query.Fields = append(query.Fields, "M_isDir")
	query.Fields = append(query.Fields, "M_checksum")
	query.Fields = append(query.Fields, "M_data")
	query.Fields = append(query.Fields, "M_thumbnail")
	query.Fields = append(query.Fields, "M_mime")
	query.Fields = append(query.Fields, "M_files")
	query.Fields = append(query.Fields, "M_fileType")

	/** associations of File **/
	query.Fields = append(query.Fields, "M_parentDirPtr")
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of File...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.File)
		this.object.TYPENAME = "CargoEntities.File"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of File **/

		/** name **/
		if results[0][4] != nil {
			this.object.M_name = results[0][4].(string)
		}

		/** path **/
		if results[0][5] != nil {
			this.object.M_path = results[0][5].(string)
		}

		/** size **/
		if results[0][6] != nil {
			this.object.M_size = results[0][6].(int)
		}

		/** modeTime **/
		if results[0][7] != nil {
			this.object.M_modeTime = results[0][7].(int64)
		}

		/** isDir **/
		if results[0][8] != nil {
			this.object.M_isDir = results[0][8].(bool)
		}

		/** checksum **/
		if results[0][9] != nil {
			this.object.M_checksum = results[0][9].(string)
		}

		/** data **/
		if results[0][10] != nil {
			this.object.M_data = results[0][10].(string)
		}

		/** thumbnail **/
		if results[0][11] != nil {
			this.object.M_thumbnail = results[0][11].(string)
		}

		/** mime **/
		if results[0][12] != nil {
			this.object.M_mime = results[0][12].(string)
		}

		/** files **/
		if results[0][13] != nil {
			uuidsStr := results[0][13].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var filesEntity *CargoEntities_FileEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							filesEntity = instance.(*CargoEntities_FileEntity)
						} else {
							filesEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), uuids[i], nil)
							filesEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(filesEntity)
						}
						filesEntity.AppendReferenced("files", this)
						this.AppendChild("files", filesEntity)
					}
				} else {
					this.lazyMap["M_files"] = uuids
				}
			}
		}

		/** fileType **/
		if results[0][14] != nil {
			enumIndex := results[0][14].(int)
			if enumIndex == 0 {
				this.object.M_fileType = CargoEntities.FileType_DbFile
			} else if enumIndex == 1 {
				this.object.M_fileType = CargoEntities.FileType_DiskFile
			}
		}

		/** associations of File **/

		/** parentDirPtr **/
		if results[0][15] != nil {
			id := results[0][15].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.File"
				id_ := refTypeName + "$$" + id
				this.object.M_parentDirPtr = id
				GetServer().GetEntityManager().appendReference("parentDirPtr", this.object.UUID, id_)
			}
		}

		/** entitiesPtr **/
		if results[0][16] != nil {
			id := results[0][16].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesFileEntityFromObject(object *CargoEntities.File) *CargoEntities_FileEntity {
	return this.NewCargoEntitiesFileEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_FileEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesFileExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.File"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_FileEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_FileEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			User
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_UserEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.User
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesUserEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_UserEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesUserExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.User).TYPENAME = "CargoEntities.User"
		object.(*CargoEntities.User).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.User", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.User).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.User).UUID
			}
			return val.(*CargoEntities_UserEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_UserEntity)
	if object == nil {
		entity.object = new(CargoEntities.User)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.User)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.User"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_UserEntity) GetTypeName() string {
	return "CargoEntities.User"
}
func (this *CargoEntities_UserEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_UserEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_UserEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_UserEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_UserEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_UserEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_UserEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_UserEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_UserEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_UserEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_UserEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_UserEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_UserEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_UserEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_UserEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_UserEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_UserEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_UserEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_UserEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_UserEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_UserEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_UserEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_UserEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.User"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_UserEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_UserEntityPrototype() {

	var userEntityProto EntityPrototype
	userEntityProto.TypeName = "CargoEntities.User"
	userEntityProto.SuperTypeNames = append(userEntityProto.SuperTypeNames, "CargoEntities.Entity")
	userEntityProto.Ids = append(userEntityProto.Ids, "UUID")
	userEntityProto.Fields = append(userEntityProto.Fields, "UUID")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 0)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.Indexs = append(userEntityProto.Indexs, "ParentUuid")
	userEntityProto.Fields = append(userEntityProto.Fields, "ParentUuid")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 1)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.Fields = append(userEntityProto.Fields, "ParentLnk")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 2)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	userEntityProto.Ids = append(userEntityProto.Ids, "M_id")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 3)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_id")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.ID")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")

	/** members of User **/
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 4)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_firstName")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 5)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_lastName")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 6)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_middle")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 7)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_phone")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 8)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_email")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "xs.string")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 9)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_memberOfRef")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "[]")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]CargoEntities.Group:Ref")
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 10)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, true)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_accounts")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "[]")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "[]CargoEntities.Account:Ref")

	/** associations of User **/
	userEntityProto.FieldsOrder = append(userEntityProto.FieldsOrder, 11)
	userEntityProto.FieldsVisibility = append(userEntityProto.FieldsVisibility, false)
	userEntityProto.Fields = append(userEntityProto.Fields, "M_entitiesPtr")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsDefaultValue = append(userEntityProto.FieldsDefaultValue, "undefined")
	userEntityProto.FieldsType = append(userEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&userEntityProto)

}

/** Create **/
func (this *CargoEntities_UserEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.User"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of User **/
	query.Fields = append(query.Fields, "M_firstName")
	query.Fields = append(query.Fields, "M_lastName")
	query.Fields = append(query.Fields, "M_middle")
	query.Fields = append(query.Fields, "M_phone")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_memberOfRef")
	query.Fields = append(query.Fields, "M_accounts")

	/** associations of User **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var UserInfo []interface{}

	UserInfo = append(UserInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		UserInfo = append(UserInfo, this.GetParentPtr().GetUuid())
		UserInfo = append(UserInfo, this.GetParentLnk())
	} else {
		UserInfo = append(UserInfo, "")
		UserInfo = append(UserInfo, "")
	}

	/** members of Entity **/
	UserInfo = append(UserInfo, this.object.M_id)

	/** members of User **/
	UserInfo = append(UserInfo, this.object.M_firstName)
	UserInfo = append(UserInfo, this.object.M_lastName)
	UserInfo = append(UserInfo, this.object.M_middle)
	UserInfo = append(UserInfo, this.object.M_phone)
	UserInfo = append(UserInfo, this.object.M_email)

	/** Save memberOfRef type Group **/
	memberOfRefStr, _ := json.Marshal(this.object.M_memberOfRef)
	UserInfo = append(UserInfo, string(memberOfRefStr))

	/** Save accounts type Account **/
	accountsStr, _ := json.Marshal(this.object.M_accounts)
	UserInfo = append(UserInfo, string(accountsStr))

	/** associations of User **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		UserInfo = append(UserInfo, this.object.M_entitiesPtr)
	} else {
		UserInfo = append(UserInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), UserInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), UserInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_UserEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_UserEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.User"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of User **/
	query.Fields = append(query.Fields, "M_firstName")
	query.Fields = append(query.Fields, "M_lastName")
	query.Fields = append(query.Fields, "M_middle")
	query.Fields = append(query.Fields, "M_phone")
	query.Fields = append(query.Fields, "M_email")
	query.Fields = append(query.Fields, "M_memberOfRef")
	query.Fields = append(query.Fields, "M_accounts")

	/** associations of User **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of User...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.User)
		this.object.TYPENAME = "CargoEntities.User"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of User **/

		/** firstName **/
		if results[0][4] != nil {
			this.object.M_firstName = results[0][4].(string)
		}

		/** lastName **/
		if results[0][5] != nil {
			this.object.M_lastName = results[0][5].(string)
		}

		/** middle **/
		if results[0][6] != nil {
			this.object.M_middle = results[0][6].(string)
		}

		/** phone **/
		if results[0][7] != nil {
			this.object.M_phone = results[0][7].(string)
		}

		/** email **/
		if results[0][8] != nil {
			this.object.M_email = results[0][8].(string)
		}

		/** memberOfRef **/
		if results[0][9] != nil {
			idsStr := results[0][9].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Group"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_memberOfRef = append(this.object.M_memberOfRef, ids[i])
					GetServer().GetEntityManager().appendReference("memberOfRef", this.object.UUID, id_)
				}
			}
		}

		/** accounts **/
		if results[0][10] != nil {
			idsStr := results[0][10].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.Account"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_accounts = append(this.object.M_accounts, ids[i])
					GetServer().GetEntityManager().appendReference("accounts", this.object.UUID, id_)
				}
			}
		}

		/** associations of User **/

		/** entitiesPtr **/
		if results[0][11] != nil {
			id := results[0][11].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesUserEntityFromObject(object *CargoEntities.User) *CargoEntities_UserEntity {
	return this.NewCargoEntitiesUserEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_UserEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesUserExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.User"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_UserEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_UserEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Group
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_GroupEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Group
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesGroupEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_GroupEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesGroupExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Group).TYPENAME = "CargoEntities.Group"
		object.(*CargoEntities.Group).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Group", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Group).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Group).UUID
			}
			return val.(*CargoEntities_GroupEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_GroupEntity)
	if object == nil {
		entity.object = new(CargoEntities.Group)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Group)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Group"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_GroupEntity) GetTypeName() string {
	return "CargoEntities.Group"
}
func (this *CargoEntities_GroupEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_GroupEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_GroupEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_GroupEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_GroupEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_GroupEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_GroupEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_GroupEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_GroupEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_GroupEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_GroupEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_GroupEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_GroupEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_GroupEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_GroupEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_GroupEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_GroupEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_GroupEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_GroupEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_GroupEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_GroupEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_GroupEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_GroupEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Group"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_GroupEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_GroupEntityPrototype() {

	var groupEntityProto EntityPrototype
	groupEntityProto.TypeName = "CargoEntities.Group"
	groupEntityProto.SuperTypeNames = append(groupEntityProto.SuperTypeNames, "CargoEntities.Entity")
	groupEntityProto.Ids = append(groupEntityProto.Ids, "UUID")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "UUID")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 0)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")
	groupEntityProto.Indexs = append(groupEntityProto.Indexs, "ParentUuid")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "ParentUuid")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 1)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")
	groupEntityProto.Fields = append(groupEntityProto.Fields, "ParentLnk")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 2)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")

	/** members of Entity **/
	groupEntityProto.Ids = append(groupEntityProto.Ids, "M_id")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 3)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_id")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.ID")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")

	/** members of Group **/
	groupEntityProto.Indexs = append(groupEntityProto.Indexs, "M_name")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 4)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_name")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "xs.string")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "")
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 5)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, true)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_membersRef")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "undefined")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "[]")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "[]CargoEntities.User:Ref")

	/** associations of Group **/
	groupEntityProto.FieldsOrder = append(groupEntityProto.FieldsOrder, 6)
	groupEntityProto.FieldsVisibility = append(groupEntityProto.FieldsVisibility, false)
	groupEntityProto.Fields = append(groupEntityProto.Fields, "M_entitiesPtr")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "undefined")
	groupEntityProto.FieldsDefaultValue = append(groupEntityProto.FieldsDefaultValue, "undefined")
	groupEntityProto.FieldsType = append(groupEntityProto.FieldsType, "CargoEntities.Entities:Ref")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&groupEntityProto)

}

/** Create **/
func (this *CargoEntities_GroupEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Group"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Group **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_membersRef")

	/** associations of Group **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	var GroupInfo []interface{}

	GroupInfo = append(GroupInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		GroupInfo = append(GroupInfo, this.GetParentPtr().GetUuid())
		GroupInfo = append(GroupInfo, this.GetParentLnk())
	} else {
		GroupInfo = append(GroupInfo, "")
		GroupInfo = append(GroupInfo, "")
	}

	/** members of Entity **/
	GroupInfo = append(GroupInfo, this.object.M_id)

	/** members of Group **/
	GroupInfo = append(GroupInfo, this.object.M_name)

	/** Save membersRef type User **/
	membersRefStr, _ := json.Marshal(this.object.M_membersRef)
	GroupInfo = append(GroupInfo, string(membersRefStr))

	/** associations of Group **/

	/** Save entities type Entities **/
	if len(this.object.M_entitiesPtr) > 0 {
		GroupInfo = append(GroupInfo, this.object.M_entitiesPtr)
	} else {
		GroupInfo = append(GroupInfo, "")
	}
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), GroupInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), GroupInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_GroupEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_GroupEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Group"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entity **/
	query.Fields = append(query.Fields, "M_id")

	/** members of Group **/
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_membersRef")

	/** associations of Group **/
	query.Fields = append(query.Fields, "M_entitiesPtr")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Group...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Group)
		this.object.TYPENAME = "CargoEntities.Group"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entity **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** members of Group **/

		/** name **/
		if results[0][4] != nil {
			this.object.M_name = results[0][4].(string)
		}

		/** membersRef **/
		if results[0][5] != nil {
			idsStr := results[0][5].(string)
			ids := make([]string, 0)
			err := json.Unmarshal([]byte(idsStr), &ids)
			if err != nil {
				return err
			}
			for i := 0; i < len(ids); i++ {
				if len(ids[i]) > 0 {
					refTypeName := "CargoEntities.User"
					id_ := refTypeName + "$$" + ids[i]
					this.object.M_membersRef = append(this.object.M_membersRef, ids[i])
					GetServer().GetEntityManager().appendReference("membersRef", this.object.UUID, id_)
				}
			}
		}

		/** associations of Group **/

		/** entitiesPtr **/
		if results[0][6] != nil {
			id := results[0][6].(string)
			if len(id) > 0 {
				refTypeName := "CargoEntities.Entities"
				id_ := refTypeName + "$$" + id
				this.object.M_entitiesPtr = id
				GetServer().GetEntityManager().appendReference("entitiesPtr", this.object.UUID, id_)
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesGroupEntityFromObject(object *CargoEntities.Group) *CargoEntities_GroupEntity {
	return this.NewCargoEntitiesGroupEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_GroupEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesGroupExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Group"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_GroupEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_GroupEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

////////////////////////////////////////////////////////////////////////////////
//              			Entities
////////////////////////////////////////////////////////////////////////////////
/** local type **/
type CargoEntities_EntitiesEntity struct {
	/** not the object id, except for the definition **/
	childsUuid     []string
	referencesUuid []string
	lazyMap        map[string]interface{}
	lazy           bool
	referenced     []EntityRef
	object         *CargoEntities.Entities
}

/** Constructor function **/
func (this *EntityManager) NewCargoEntitiesEntitiesEntity(parentUuid string, objectId string, object interface{}) *CargoEntities_EntitiesEntity {
	var uuidStr string
	if len(objectId) > 0 {
		if Utility.IsValidEntityReferenceName(objectId) {
			uuidStr = objectId
		} else {
			uuidStr = CargoEntitiesEntitiesExists(objectId)
		}
	}
	if object != nil {
		object.(*CargoEntities.Entities).TYPENAME = "CargoEntities.Entities"
		object.(*CargoEntities.Entities).ParentUuid = parentUuid
	}
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype("CargoEntities.Entities", "CargoEntities")
	if len(uuidStr) > 0 {
		if object != nil {
			object.(*CargoEntities.Entities).UUID = uuidStr
		}
		if val, ok := this.contain(uuidStr); ok {
			if object != nil {
				this.setObjectValues(val, object)

				uuidStr = object.(*CargoEntities.Entities).UUID
			}
			return val.(*CargoEntities_EntitiesEntity)
		}
	} else {
		if len(prototype.Ids) == 1 {
			// Here there is a new entity...
			uuidStr = prototype.TypeName + "%" + Utility.RandomUUID()
		} else {
			var keyInfo string
			if len(parentUuid) > 0 {
				keyInfo += parentUuid + ":"
			}
			keyInfo += prototype.TypeName + ":"
			for i := 1; i < len(prototype.Ids); i++ {
				var getter = "Get" + strings.ToUpper(prototype.Ids[i][2:3]) + prototype.Ids[i][3:]
				params := make([]interface{}, 0)
				value, _ := Utility.CallMethod(object, getter, params)
				keyInfo += Utility.ToString(value)
				// Append underscore for readability in case of problem...
				if i < len(prototype.Ids)-1 {
					keyInfo += "_"
				}
			}

			// The uuid is in that case a MD5 value.
			uuidStr = prototype.TypeName + "%" + Utility.GenerateUUID(keyInfo)
		}
	}
	entity := new(CargoEntities_EntitiesEntity)
	if object == nil {
		entity.object = new(CargoEntities.Entities)
		entity.SetNeedSave(true)
	} else {
		entity.object = object.(*CargoEntities.Entities)
		entity.SetNeedSave(true)
	}
	entity.lazyMap = make(map[string]interface{})
	entity.object.TYPENAME = "CargoEntities.Entities"

	entity.object.UUID = uuidStr
	entity.object.ParentUuid = parentUuid
	entity.SetInit(false)
	this.insert(entity)
	return entity
}

/** Entity functions **/
func (this *CargoEntities_EntitiesEntity) GetTypeName() string {
	return "CargoEntities.Entities"
}
func (this *CargoEntities_EntitiesEntity) GetUuid() string {
	return this.object.UUID
}
func (this *CargoEntities_EntitiesEntity) GetParentUuid() string {
	return this.object.ParentUuid
}
func (this *CargoEntities_EntitiesEntity) GetParentPtr() Entity {
	parentPtr, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetParentUuid(), true)
	return parentPtr
}

func (this *CargoEntities_EntitiesEntity) SetParentLnk(lnk string) {
	this.object.ParentLnk = lnk
}

func (this *CargoEntities_EntitiesEntity) GetParentLnk() string {
	return this.object.ParentLnk
}
func (this *CargoEntities_EntitiesEntity) AppendReferenced(name string, owner Entity) {
	if owner.GetUuid() == this.GetUuid() {
		return
	}
	var ref EntityRef
	ref.Name = name
	ref.OwnerUuid = owner.GetUuid()
	for i := 0; i < len(this.referenced); i++ {
		if this.referenced[i].Name == ref.Name && this.referenced[i].OwnerUuid == ref.OwnerUuid {
			return
		}
	}
	this.referenced = append(this.referenced, ref)
}

func (this *CargoEntities_EntitiesEntity) GetReferenced() []EntityRef {
	return this.referenced
}

func (this *CargoEntities_EntitiesEntity) GetSize() uint {
	return uint(unsafe.Sizeof(*this.object))
}

func (this *CargoEntities_EntitiesEntity) RemoveReferenced(name string, owner Entity) {
	var referenced []EntityRef
	referenced = make([]EntityRef, 0)
	for i := 0; i < len(this.referenced); i++ {
		ref := this.referenced[i]
		if !(ref.Name == name && ref.OwnerUuid == owner.GetUuid()) {
			referenced = append(referenced, ref)
		}
	}
	// Set the reference.
	this.referenced = referenced
}

func (this *CargoEntities_EntitiesEntity) RemoveReference(name string, reference Entity) {
	refsUuid := make([]string, 0)
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid != reference.GetUuid() {
			refsUuid = append(refsUuid, reference.GetUuid())
		}
	}
	// Set the new array...
	this.SetReferencesUuid(refsUuid)
	var removeMethode = "Remove" + strings.ToUpper(name[2:3]) + name[3:]
	params := make([]interface{}, 1)
	params[0] = reference.GetObject()
	Utility.CallMethod(this.GetObject(), removeMethode, params)
}

func (this *CargoEntities_EntitiesEntity) GetChildsUuid() []string {
	return this.childsUuid
}

func (this *CargoEntities_EntitiesEntity) SetChildsUuid(childsUuid []string) {
	this.childsUuid = childsUuid
}

/**
 * Remove a child uuid form the list of child in an entity.
 */
func (this *CargoEntities_EntitiesEntity) RemoveChild(name string, uuid string) {
	childsUuid := make([]string, 0)
	params := make([]interface{}, 1)
	for i := 0; i < len(this.GetChildsUuid()); i++ {
		if this.GetChildsUuid()[i] != uuid {
			childsUuid = append(childsUuid, this.GetChildsUuid()[i])
		} else {
			entity, _ := GetServer().GetEntityManager().getEntityByUuid(this.GetChildsUuid()[i], false)
			params[0] = entity.GetObject()
		}
	}

	this.childsUuid = childsUuid
	var removeMethode = "Remove" + strings.ToUpper(name[0:1]) + name[1:]
	if params[0] != nil {
		Utility.CallMethod(this.GetObject(), removeMethode, params)
	}
}

func (this *CargoEntities_EntitiesEntity) GetReferencesUuid() []string {
	return this.referencesUuid
}

func (this *CargoEntities_EntitiesEntity) SetReferencesUuid(refsUuid []string) {
	this.referencesUuid = refsUuid
}

func (this *CargoEntities_EntitiesEntity) GetObject() interface{} {
	return this.object
}

func (this *CargoEntities_EntitiesEntity) NeedSave() bool {
	return this.object.NeedSave
}

func (this *CargoEntities_EntitiesEntity) SetNeedSave(needSave bool) {
	this.object.NeedSave = needSave
}

func (this *CargoEntities_EntitiesEntity) IsInit() bool {
	return this.object.IsInit
}

func (this *CargoEntities_EntitiesEntity) SetInit(isInit bool) {
	this.object.IsInit = isInit
}

func (this *CargoEntities_EntitiesEntity) IsLazy() bool {
	return this.lazy
}

func (this *CargoEntities_EntitiesEntity) GetChecksum() string {
	mapValues, _ := Utility.ToMap(this.object)
	return Utility.GetChecksum(mapValues)
}

func (this *CargoEntities_EntitiesEntity) Exist() bool {
	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"
	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return false
	}
	return len(results[0][0].(string)) > 0

}

/**
* Return the entity prototype.
 */
func (this *CargoEntities_EntitiesEntity) GetPrototype() *EntityPrototype {
	typeName := this.GetTypeName()
	prototype, _ := GetServer().GetEntityManager().getEntityPrototype(typeName, typeName[0:strings.Index(typeName, ".")])
	return prototype
}

/** Entity Prototype creation **/
func (this *EntityManager) create_CargoEntities_EntitiesEntityPrototype() {

	var entitiesEntityProto EntityPrototype
	entitiesEntityProto.TypeName = "CargoEntities.Entities"
	entitiesEntityProto.Ids = append(entitiesEntityProto.Ids, "UUID")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "UUID")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 0)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.Indexs = append(entitiesEntityProto.Indexs, "ParentUuid")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "ParentUuid")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 1)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "ParentLnk")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 2)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, false)
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")

	/** members of Entities **/
	entitiesEntityProto.Ids = append(entitiesEntityProto.Ids, "M_id")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 3)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_id")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.ID")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 4)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_name")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 5)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_version")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "xs.string")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 6)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_entities")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Entity")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 7)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_roles")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Role")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 8)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_permissions")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Permission")
	entitiesEntityProto.FieldsOrder = append(entitiesEntityProto.FieldsOrder, 9)
	entitiesEntityProto.FieldsVisibility = append(entitiesEntityProto.FieldsVisibility, true)
	entitiesEntityProto.Fields = append(entitiesEntityProto.Fields, "M_actions")
	entitiesEntityProto.FieldsDefaultValue = append(entitiesEntityProto.FieldsDefaultValue, "[]")
	entitiesEntityProto.FieldsType = append(entitiesEntityProto.FieldsType, "[]CargoEntities.Action")

	store := GetServer().GetDataManager().getDataStore(CargoEntitiesDB).(*KeyValueDataStore)
	store.SetEntityPrototype(&entitiesEntityProto)

}

/** Create **/
func (this *CargoEntities_EntitiesEntity) SaveEntity() {
	if this.object.NeedSave == false {
		return
	}

	if this.lazy == true {
		this.InitEntity(this.GetUuid(), false)
	}

	this.SetNeedSave(false)
	this.SetInit(true)
	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entities **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_entities")
	query.Fields = append(query.Fields, "M_roles")
	query.Fields = append(query.Fields, "M_permissions")
	query.Fields = append(query.Fields, "M_actions")

	var EntitiesInfo []interface{}

	EntitiesInfo = append(EntitiesInfo, this.GetUuid())
	if this.GetParentPtr() != nil {
		EntitiesInfo = append(EntitiesInfo, this.GetParentPtr().GetUuid())
		EntitiesInfo = append(EntitiesInfo, this.GetParentLnk())
	} else {
		EntitiesInfo = append(EntitiesInfo, "")
		EntitiesInfo = append(EntitiesInfo, "")
	}

	/** members of Entities **/
	EntitiesInfo = append(EntitiesInfo, this.object.M_id)
	EntitiesInfo = append(EntitiesInfo, this.object.M_name)
	EntitiesInfo = append(EntitiesInfo, this.object.M_version)

	/** Save entities type Entity **/
	entitiesIds := make([]string, 0)
	lazy_entities := this.lazyMap["M_entities"] != nil && len(this.object.M_entities) == 0
	if !lazy_entities {
		for i := 0; i < len(this.object.M_entities); i++ {
			switch v := this.object.M_entities[i].(type) {
			case *CargoEntities.Notification:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.Account:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesAccountEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.User:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesUserEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.Group:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesGroupEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.Error:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.TextMessage:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.LogEntry:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.Log:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesLogEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.Project:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesProjectEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.Computer:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesComputerEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			case *CargoEntities.File:
				entitiesEntity := GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), v.UUID, v)
				entitiesIds = append(entitiesIds, entitiesEntity.GetUuid())
				entitiesEntity.AppendReferenced("entities", this)
				this.AppendChild("entities", entitiesEntity)
				if entitiesEntity.NeedSave() {
					entitiesEntity.SaveEntity()
				}
			}
		}
	} else {
		entitiesIds = this.lazyMap["M_entities"].([]string)
	}
	entitiesStr, _ := json.Marshal(entitiesIds)
	EntitiesInfo = append(EntitiesInfo, string(entitiesStr))

	/** Save roles type Role **/
	rolesIds := make([]string, 0)
	lazy_roles := this.lazyMap["M_roles"] != nil && len(this.object.M_roles) == 0
	if !lazy_roles {
		for i := 0; i < len(this.object.M_roles); i++ {
			rolesEntity := GetServer().GetEntityManager().NewCargoEntitiesRoleEntity(this.GetUuid(), this.object.M_roles[i].UUID, this.object.M_roles[i])
			rolesIds = append(rolesIds, rolesEntity.GetUuid())
			rolesEntity.AppendReferenced("roles", this)
			this.AppendChild("roles", rolesEntity)
			if rolesEntity.NeedSave() {
				rolesEntity.SaveEntity()
			}
		}
	} else {
		rolesIds = this.lazyMap["M_roles"].([]string)
	}
	rolesStr, _ := json.Marshal(rolesIds)
	EntitiesInfo = append(EntitiesInfo, string(rolesStr))

	/** Save permissions type Permission **/
	permissionsIds := make([]string, 0)
	lazy_permissions := this.lazyMap["M_permissions"] != nil && len(this.object.M_permissions) == 0
	if !lazy_permissions {
		for i := 0; i < len(this.object.M_permissions); i++ {
			permissionsEntity := GetServer().GetEntityManager().NewCargoEntitiesPermissionEntity(this.GetUuid(), this.object.M_permissions[i].UUID, this.object.M_permissions[i])
			permissionsIds = append(permissionsIds, permissionsEntity.GetUuid())
			permissionsEntity.AppendReferenced("permissions", this)
			this.AppendChild("permissions", permissionsEntity)
			if permissionsEntity.NeedSave() {
				permissionsEntity.SaveEntity()
			}
		}
	} else {
		permissionsIds = this.lazyMap["M_permissions"].([]string)
	}
	permissionsStr, _ := json.Marshal(permissionsIds)
	EntitiesInfo = append(EntitiesInfo, string(permissionsStr))

	/** Save actions type Action **/
	actionsIds := make([]string, 0)
	lazy_actions := this.lazyMap["M_actions"] != nil && len(this.object.M_actions) == 0
	if !lazy_actions {
		for i := 0; i < len(this.object.M_actions); i++ {
			actionsEntity := GetServer().GetEntityManager().NewCargoEntitiesActionEntity(this.GetUuid(), this.object.M_actions[i].UUID, this.object.M_actions[i])
			actionsIds = append(actionsIds, actionsEntity.GetUuid())
			actionsEntity.AppendReferenced("actions", this)
			this.AppendChild("actions", actionsEntity)
			if actionsEntity.NeedSave() {
				actionsEntity.SaveEntity()
			}
		}
	} else {
		actionsIds = this.lazyMap["M_actions"].([]string)
	}
	actionsStr, _ := json.Marshal(actionsIds)
	EntitiesInfo = append(EntitiesInfo, string(actionsStr))
	eventData := make([]*MessageData, 1)
	msgData := new(MessageData)
	msgData.Name = "entity"
	msgData.Value = this.GetObject()
	eventData[0] = msgData
	var err error
	var evt *Event
	if this.Exist() == true {
		evt, _ = NewEvent(UpdateEntityEvent, EntityEvent, eventData)
		var params []interface{}
		query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())
		queryStr, _ := json.Marshal(query)
		err = GetServer().GetDataManager().updateData(CargoEntitiesDB, string(queryStr), EntitiesInfo, params)
	} else {
		evt, _ = NewEvent(NewEntityEvent, EntityEvent, eventData)
		queryStr, _ := json.Marshal(query)
		_, err = GetServer().GetDataManager().createData(CargoEntitiesDB, string(queryStr), EntitiesInfo)
	}
	if err == nil {
		GetServer().GetEntityManager().insert(this)
		GetServer().GetEntityManager().setReferences(this)
		GetServer().GetEventManager().BroadcastEvent(evt)
	}
}

/** Read **/
func (this *CargoEntities_EntitiesEntity) InitEntity(id string, lazy bool) error {
	if this.object.IsInit == true {
		entity, err := GetServer().GetEntityManager().getEntityByUuid(id, lazy)
		if err == nil {
			// Return the already initialyse entity.
			this = entity.(*CargoEntities_EntitiesEntity)
			return nil
		}
		// I must reinit the entity if the entity manager dosent have it.
		this.object.IsInit = false
	}
	this.lazy = lazy

	// Set the reference on the map
	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"

	query.Fields = append(query.Fields, "UUID")
	query.Fields = append(query.Fields, "ParentUuid")
	query.Fields = append(query.Fields, "ParentLnk")

	/** members of Entities **/
	query.Fields = append(query.Fields, "M_id")
	query.Fields = append(query.Fields, "M_name")
	query.Fields = append(query.Fields, "M_version")
	query.Fields = append(query.Fields, "M_entities")
	query.Fields = append(query.Fields, "M_roles")
	query.Fields = append(query.Fields, "M_permissions")
	query.Fields = append(query.Fields, "M_actions")

	query.Indexs = append(query.Indexs, "UUID="+this.GetUuid())

	var fieldsType []interface{} // not use...
	var params []interface{}
	var results [][]interface{}
	var err error
	queryStr, _ := json.Marshal(query)

	results, err = GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil {
		return err
	}
	// Initialisation of information of Entities...
	if len(results) > 0 {

		/** initialyzation of the entity object **/
		this.object = new(CargoEntities.Entities)
		this.object.TYPENAME = "CargoEntities.Entities"

		this.object.UUID = results[0][0].(string)
		this.object.ParentUuid = results[0][1].(string)
		this.object.ParentLnk = results[0][2].(string)

		/** members of Entities **/

		/** id **/
		if results[0][3] != nil {
			this.object.M_id = results[0][3].(string)
		}

		/** name **/
		if results[0][4] != nil {
			this.object.M_name = results[0][4].(string)
		}

		/** version **/
		if results[0][5] != nil {
			this.object.M_version = results[0][5].(string)
		}

		/** entities **/
		if results[0][6] != nil {
			uuidsStr := results[0][6].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					typeName := uuids[i][0:strings.Index(uuids[i], "%")]
					if err != nil {
						log.Println("type ", typeName, " not found!")
						return err
					}
					if typeName == "CargoEntities.Computer" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_ComputerEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_ComputerEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesComputerEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.File" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_FileEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_FileEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesFileEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.Error" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_ErrorEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_ErrorEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesErrorEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.LogEntry" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_LogEntryEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_LogEntryEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesLogEntryEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.Log" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_LogEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_LogEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesLogEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.Project" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_ProjectEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_ProjectEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesProjectEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.TextMessage" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_TextMessageEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_TextMessageEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesTextMessageEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.Notification" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_NotificationEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_NotificationEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesNotificationEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.Account" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_AccountEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_AccountEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesAccountEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.User" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_UserEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_UserEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesUserEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					} else if typeName == "CargoEntities.Group" {
						if len(uuids[i]) > 0 {
							var entitiesEntity *CargoEntities_GroupEntity
							if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
								entitiesEntity = instance.(*CargoEntities_GroupEntity)
							} else {
								entitiesEntity = GetServer().GetEntityManager().NewCargoEntitiesGroupEntity(this.GetUuid(), uuids[i], nil)
								entitiesEntity.InitEntity(uuids[i], lazy)
								GetServer().GetEntityManager().insert(entitiesEntity)
							}
							entitiesEntity.AppendReferenced("entities", this)
							this.AppendChild("entities", entitiesEntity)
						}
					}
				} else {
					this.lazyMap["M_entities"] = uuids
				}
			}
		}

		/** roles **/
		if results[0][7] != nil {
			uuidsStr := results[0][7].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var rolesEntity *CargoEntities_RoleEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							rolesEntity = instance.(*CargoEntities_RoleEntity)
						} else {
							rolesEntity = GetServer().GetEntityManager().NewCargoEntitiesRoleEntity(this.GetUuid(), uuids[i], nil)
							rolesEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(rolesEntity)
						}
						rolesEntity.AppendReferenced("roles", this)
						this.AppendChild("roles", rolesEntity)
					}
				} else {
					this.lazyMap["M_roles"] = uuids
				}
			}
		}

		/** permissions **/
		if results[0][8] != nil {
			uuidsStr := results[0][8].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var permissionsEntity *CargoEntities_PermissionEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							permissionsEntity = instance.(*CargoEntities_PermissionEntity)
						} else {
							permissionsEntity = GetServer().GetEntityManager().NewCargoEntitiesPermissionEntity(this.GetUuid(), uuids[i], nil)
							permissionsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(permissionsEntity)
						}
						permissionsEntity.AppendReferenced("permissions", this)
						this.AppendChild("permissions", permissionsEntity)
					}
				} else {
					this.lazyMap["M_permissions"] = uuids
				}
			}
		}

		/** actions **/
		if results[0][9] != nil {
			uuidsStr := results[0][9].(string)
			uuids := make([]string, 0)
			err := json.Unmarshal([]byte(uuidsStr), &uuids)
			if err != nil {
				return err
			}
			for i := 0; i < len(uuids); i++ {
				if !lazy {
					if len(uuids[i]) > 0 {
						var actionsEntity *CargoEntities_ActionEntity
						if instance, ok := GetServer().GetEntityManager().contain(uuids[i]); ok {
							actionsEntity = instance.(*CargoEntities_ActionEntity)
						} else {
							actionsEntity = GetServer().GetEntityManager().NewCargoEntitiesActionEntity(this.GetUuid(), uuids[i], nil)
							actionsEntity.InitEntity(uuids[i], lazy)
							GetServer().GetEntityManager().insert(actionsEntity)
						}
						actionsEntity.AppendReferenced("actions", this)
						this.AppendChild("actions", actionsEntity)
					}
				} else {
					this.lazyMap["M_actions"] = uuids
				}
			}
		}
	}

	// set need save to false.
	this.SetNeedSave(false)
	// set init done.
	this.SetInit(true)
	// Init the references...
	GetServer().GetEntityManager().InitEntity(this, lazy)
	return nil
}

/** instantiate a new entity from an existing object. **/
func (this *EntityManager) NewCargoEntitiesEntitiesEntityFromObject(object *CargoEntities.Entities) *CargoEntities_EntitiesEntity {
	return this.NewCargoEntitiesEntitiesEntity("", object.UUID, object)
}

/** Delete **/
func (this *CargoEntities_EntitiesEntity) DeleteEntity() {
	GetServer().GetEntityManager().deleteEntity(this)
}

/** Exists **/
func CargoEntitiesEntitiesExists(val string) string {
	var query EntityQuery
	query.TypeName = "CargoEntities.Entities"
	query.Indexs = append(query.Indexs, "M_id="+val)
	query.Fields = append(query.Fields, "UUID")
	var fieldsType []interface{} // not use...
	var params []interface{}
	queryStr, _ := json.Marshal(query)
	results, err := GetServer().GetDataManager().readData(CargoEntitiesDB, string(queryStr), fieldsType, params)
	if err != nil || len(results) == 0 {
		return ""
	}
	return results[0][0].(string)
}

/** Append child entity into parent entity. **/
func (this *CargoEntities_EntitiesEntity) AppendChild(attributeName string, child Entity) error {

	// Append child if is not there...
	if !Utility.Contains(this.childsUuid, child.GetUuid()) {
		this.childsUuid = append(this.childsUuid, child.GetUuid())
	}
	// Set this as parent in the child
	child.SetParentLnk("M_" + attributeName)

	params := make([]interface{}, 1)
	params[0] = child.GetObject()
	attributeName = strings.Replace(attributeName, "M_", "", -1)
	methodName := "Set" + strings.ToUpper(attributeName[0:1]) + attributeName[1:]
	_, invalidMethod := Utility.CallMethod(this.object, methodName, params)
	if invalidMethod != nil {
		return invalidMethod.(error)
	}
	return nil
}

/** Append reference entity into parent entity. **/
func (this *CargoEntities_EntitiesEntity) AppendReference(reference Entity) {

	// Here i will append the reference uuid
	index := -1
	for i := 0; i < len(this.referencesUuid); i++ {
		refUuid := this.referencesUuid[i]
		if refUuid == reference.GetUuid() {
			index = i
			break
		}
	}
	if index == -1 {
		this.referencesUuid = append(this.referencesUuid, reference.GetUuid())
	}
}

/** Register the entity to the dynamic typing system. **/
func (this *EntityManager) registerCargoEntitiesObjects() {
	Utility.RegisterType((*CargoEntities.Parameter)(nil))
	Utility.RegisterType((*CargoEntities.Action)(nil))
	Utility.RegisterType((*CargoEntities.Error)(nil))
	Utility.RegisterType((*CargoEntities.LogEntry)(nil))
	Utility.RegisterType((*CargoEntities.Log)(nil))
	Utility.RegisterType((*CargoEntities.Project)(nil))
	Utility.RegisterType((*CargoEntities.Notification)(nil))
	Utility.RegisterType((*CargoEntities.TextMessage)(nil))
	Utility.RegisterType((*CargoEntities.Session)(nil))
	Utility.RegisterType((*CargoEntities.Role)(nil))
	Utility.RegisterType((*CargoEntities.Account)(nil))
	Utility.RegisterType((*CargoEntities.Computer)(nil))
	Utility.RegisterType((*CargoEntities.Permission)(nil))
	Utility.RegisterType((*CargoEntities.File)(nil))
	Utility.RegisterType((*CargoEntities.User)(nil))
	Utility.RegisterType((*CargoEntities.Group)(nil))
	Utility.RegisterType((*CargoEntities.Entities)(nil))
}

/** Create entity prototypes contain in a package **/
func (this *EntityManager) createCargoEntitiesPrototypes() {
	this.create_CargoEntities_EntityEntityPrototype()
	this.create_CargoEntities_ParameterEntityPrototype()
	this.create_CargoEntities_ActionEntityPrototype()
	this.create_CargoEntities_ErrorEntityPrototype()
	this.create_CargoEntities_LogEntryEntityPrototype()
	this.create_CargoEntities_LogEntityPrototype()
	this.create_CargoEntities_ProjectEntityPrototype()
	this.create_CargoEntities_MessageEntityPrototype()
	this.create_CargoEntities_NotificationEntityPrototype()
	this.create_CargoEntities_TextMessageEntityPrototype()
	this.create_CargoEntities_SessionEntityPrototype()
	this.create_CargoEntities_RoleEntityPrototype()
	this.create_CargoEntities_AccountEntityPrototype()
	this.create_CargoEntities_ComputerEntityPrototype()
	this.create_CargoEntities_PermissionEntityPrototype()
	this.create_CargoEntities_FileEntityPrototype()
	this.create_CargoEntities_UserEntityPrototype()
	this.create_CargoEntities_GroupEntityPrototype()
	this.create_CargoEntities_EntitiesEntityPrototype()
}
