// +build Config

package Config

import(
	"encoding/xml"
	"code.myceliUs.com/Utility"
)

type OAuth2Refresh struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)
	/** Use to put the entity in the cache **/
	setEntity func(interface{})
	/** Generate the entity uuid **/
	generateUuid func(interface{}) string

	/** members of OAuth2Refresh **/
	M_id string
	M_access string


	/** Associations **/
	M_parentPtr string
}

/** Xml parser for OAuth2Refresh **/
type XsdOAuth2Refresh struct {
	XMLName xml.Name	`xml:"oauth2Refresh"`
	M_id	string	`xml:"id,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *OAuth2Refresh) GetUuid() string{
	if len(this.UUID) == 0 {
		this.SetUuid(this.generateUuid(this))
	}
	return this.UUID
}
func (this *OAuth2Refresh) SetUuid(uuid string){
	this.UUID = uuid
}

func (this *OAuth2Refresh) SetFieldValue(field string, value interface{}) error{
	return Utility.SetProperty(this, field, value)
}

func (this *OAuth2Refresh) GetFieldValue(field string) interface{}{
	return Utility.GetProperty(this, field)
}

/** Return the array of entity id's without it uuid **/
func (this *OAuth2Refresh) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *OAuth2Refresh) GetTypeName() string{
	this.TYPENAME = "Config.OAuth2Refresh"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *OAuth2Refresh) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *OAuth2Refresh) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Refresh) GetParentLnk() string{
	return this.ParentLnk
}
func (this *OAuth2Refresh) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

func (this *OAuth2Refresh) GetParent() interface{}{
	parent, err := this.getEntityByUuid(this.ParentUuid)
	if err != nil {
		return nil
	}
	return parent
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *OAuth2Refresh) GetChilds() []interface{}{
	var childs []interface{}
	return childs
}
/** Return the list of all childs uuid **/
func (this *OAuth2Refresh) GetChildsUuid() []string{
	var childs []string
	return childs
}
/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *OAuth2Refresh) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}
/** Use it the set the entity on the cache. **/
func (this *OAuth2Refresh) SetEntitySetter(fct func(entity interface{})){
	this.setEntity = fct
}
/** Set the uuid generator function **/
func (this *OAuth2Refresh) SetUuidGenerator(fct func(entity interface{}) string){
	this.generateUuid = fct
}

func (this *OAuth2Refresh) GetId()string{
	return this.M_id
}

func (this *OAuth2Refresh) SetId(val string){
	this.M_id= val
}




func (this *OAuth2Refresh) GetAccess()*OAuth2Access{
	entity, err := this.getEntityByUuid(this.M_access)
	if err == nil {
		return entity.(*OAuth2Access)
	}
	return nil
}

func (this *OAuth2Refresh) SetAccess(val *OAuth2Access){
	this.M_access= val.GetUuid()
	this.setEntity(this)
}


func (this *OAuth2Refresh) ResetAccess(){
	this.M_access= ""
}


func (this *OAuth2Refresh) GetParentPtr()*OAuth2Configuration{
	entity, err := this.getEntityByUuid(this.M_parentPtr)
	if err == nil {
		return entity.(*OAuth2Configuration)
	}
	return nil
}

func (this *OAuth2Refresh) SetParentPtr(val *OAuth2Configuration){
	this.M_parentPtr= val.GetUuid()
	this.setEntity(this)
}


func (this *OAuth2Refresh) ResetParentPtr(){
	this.M_parentPtr= ""
}

