// +build Config

package Config

import(
	"encoding/xml"
)

type ServerConfiguration struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** The parent uuid if there is some. **/
	ParentUuid string
	/** The relation name with the parent. **/
	ParentLnk string
	/** If the entity value has change... **/
	NeedSave bool
	/** Get entity by uuid function **/
	getEntityByUuid func(string)(interface{}, error)

	/** members of Configuration **/
	M_id string

	/** members of ServerConfiguration **/
	M_hostName string
	M_ipv4 string
	M_serverPort int
	M_serviceContainerPort int
	M_applicationsPath string
	M_dataPath string
	M_scriptsPath string
	M_definitionsPath string
	M_schemasPath string
	M_tmpPath string
	M_binPath string


	/** Associations **/
	m_parentPtr *Configurations
	/** If the ref is a string and not an object **/
	M_parentPtr string
}

/** Xml parser for ServerConfiguration **/
type XsdServerConfiguration struct {
	XMLName xml.Name	`xml:"serverConfiguration"`
	/** Configuration **/
	M_id	string	`xml:"id,attr"`


	M_ipv4	string	`xml:"ipv4,attr"`
	M_hostName	string	`xml:"hostName,attr"`
	M_serverPort	int	`xml:"serverPort,attr"`
	M_serviceContainerPort	int	`xml:"serviceContainerPort,attr"`
	M_applicationsPath	string	`xml:"applicationsPath,attr"`
	M_dataPath	string	`xml:"dataPath,attr"`
	M_scriptsPath	string	`xml:"scriptsPath,attr"`
	M_definitionsPath	string	`xml:"definitionsPath,attr"`
	M_schemasPath	string	`xml:"schemasPath,attr"`
	M_tmpPath	string	`xml:"tmpPath,attr"`
	M_binPath	string	`xml:"binPath,attr"`

}
/***************** Entity **************************/

/** UUID **/
func (this *ServerConfiguration) GetUuid() string{
	return this.UUID
}
func (this *ServerConfiguration) SetUuid(uuid string){
	this.UUID = uuid
}

/** Return the array of entity id's without it uuid **/
func (this *ServerConfiguration) Ids() []interface{} {
	ids := make([]interface{}, 0)
	ids = append(ids, this.M_id)
	return ids
}

/** The type name **/
func (this *ServerConfiguration) GetTypeName() string{
	this.TYPENAME = "Config.ServerConfiguration"
	return this.TYPENAME
}

/** Return the entity parent UUID **/
func (this *ServerConfiguration) GetParentUuid() string{
	return this.ParentUuid
}

/** Set it parent UUID **/
func (this *ServerConfiguration) SetParentUuid(parentUuid string){
	this.ParentUuid = parentUuid
}

/** Return it relation with it parent, only one parent is possible by entity. **/
func (this *ServerConfiguration) GetParentLnk() string{
	return this.ParentLnk
}
func (this *ServerConfiguration) SetParentLnk(parentLnk string){
	this.ParentLnk = parentLnk
}

/** Evaluate if an entity needs to be saved. **/
func (this *ServerConfiguration) IsNeedSave() bool{
	return this.NeedSave
}
func (this *ServerConfiguration) ResetNeedSave(){
	this.NeedSave=false
}

/** Give access to entity manager GetEntityByUuid function from Entities package. **/
func (this *ServerConfiguration) SetEntityGetter(fct func(uuid string)(interface{}, error)){
	this.getEntityByUuid = fct
}

/** Id **/
func (this *ServerConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ServerConfiguration) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Id **/

/** HostName **/
func (this *ServerConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *ServerConfiguration) SetHostName(ref interface{}){
	if this.M_hostName != ref.(string) {
		this.M_hostName = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *ServerConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *ServerConfiguration) SetIpv4(ref interface{}){
	if this.M_ipv4 != ref.(string) {
		this.M_ipv4 = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference Ipv4 **/

/** ServerPort **/
func (this *ServerConfiguration) GetServerPort() int{
	return this.M_serverPort
}

/** Init reference ServerPort **/
func (this *ServerConfiguration) SetServerPort(ref interface{}){
	if this.M_serverPort != ref.(int) {
		this.M_serverPort = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference ServerPort **/

/** ServiceContainerPort **/
func (this *ServerConfiguration) GetServiceContainerPort() int{
	return this.M_serviceContainerPort
}

/** Init reference ServiceContainerPort **/
func (this *ServerConfiguration) SetServiceContainerPort(ref interface{}){
	if this.M_serviceContainerPort != ref.(int) {
		this.M_serviceContainerPort = ref.(int)
		this.NeedSave = true
	}
}

/** Remove reference ServiceContainerPort **/

/** ApplicationsPath **/
func (this *ServerConfiguration) GetApplicationsPath() string{
	return this.M_applicationsPath
}

/** Init reference ApplicationsPath **/
func (this *ServerConfiguration) SetApplicationsPath(ref interface{}){
	if this.M_applicationsPath != ref.(string) {
		this.M_applicationsPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference ApplicationsPath **/

/** DataPath **/
func (this *ServerConfiguration) GetDataPath() string{
	return this.M_dataPath
}

/** Init reference DataPath **/
func (this *ServerConfiguration) SetDataPath(ref interface{}){
	if this.M_dataPath != ref.(string) {
		this.M_dataPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference DataPath **/

/** ScriptsPath **/
func (this *ServerConfiguration) GetScriptsPath() string{
	return this.M_scriptsPath
}

/** Init reference ScriptsPath **/
func (this *ServerConfiguration) SetScriptsPath(ref interface{}){
	if this.M_scriptsPath != ref.(string) {
		this.M_scriptsPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference ScriptsPath **/

/** DefinitionsPath **/
func (this *ServerConfiguration) GetDefinitionsPath() string{
	return this.M_definitionsPath
}

/** Init reference DefinitionsPath **/
func (this *ServerConfiguration) SetDefinitionsPath(ref interface{}){
	if this.M_definitionsPath != ref.(string) {
		this.M_definitionsPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference DefinitionsPath **/

/** SchemasPath **/
func (this *ServerConfiguration) GetSchemasPath() string{
	return this.M_schemasPath
}

/** Init reference SchemasPath **/
func (this *ServerConfiguration) SetSchemasPath(ref interface{}){
	if this.M_schemasPath != ref.(string) {
		this.M_schemasPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference SchemasPath **/

/** TmpPath **/
func (this *ServerConfiguration) GetTmpPath() string{
	return this.M_tmpPath
}

/** Init reference TmpPath **/
func (this *ServerConfiguration) SetTmpPath(ref interface{}){
	if this.M_tmpPath != ref.(string) {
		this.M_tmpPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference TmpPath **/

/** BinPath **/
func (this *ServerConfiguration) GetBinPath() string{
	return this.M_binPath
}

/** Init reference BinPath **/
func (this *ServerConfiguration) SetBinPath(ref interface{}){
	if this.M_binPath != ref.(string) {
		this.M_binPath = ref.(string)
		this.NeedSave = true
	}
}

/** Remove reference BinPath **/

/** Parent **/
func (this *ServerConfiguration) GetParentPtr() *Configurations{
	if this.m_parentPtr == nil {
		entity, err := this.getEntityByUuid(this.M_parentPtr)
		if err == nil {
			this.m_parentPtr = entity.(*Configurations)
		}
	}
	return this.m_parentPtr
}
func (this *ServerConfiguration) GetParentPtrStr() string{
	return this.M_parentPtr
}

/** Init reference Parent **/
func (this *ServerConfiguration) SetParentPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_parentPtr != ref.(string) {
			this.M_parentPtr = ref.(string)
			this.NeedSave = true
		}
	}else{
		if this.M_parentPtr != ref.(*Configurations).GetUuid() {
			this.M_parentPtr = ref.(*Configurations).GetUuid()
			this.NeedSave = true
		}
		this.m_parentPtr = ref.(*Configurations)
	}
}

/** Remove reference Parent **/
func (this *ServerConfiguration) RemoveParentPtr(ref interface{}){
	toDelete := ref.(*Configurations)
	if this.m_parentPtr!= nil {
		if toDelete.GetUuid() == this.m_parentPtr.GetUuid() {
			this.m_parentPtr = nil
			this.M_parentPtr = ""
			this.NeedSave = true
		}
	}
}
