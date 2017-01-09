package CargoConfig

import(
"encoding/xml"
)

type ServerConfiguration struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Configuration **/
	M_id string

	/** members of ServerConfiguration **/
	M_hostName string
	M_ipv4 string
	M_port int
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
	M_port	int	`xml:"port,attr"`
	M_applicationsPath	string	`xml:"applicationsPath,attr"`
	M_dataPath	string	`xml:"dataPath,attr"`
	M_scriptsPath	string	`xml:"scriptsPath,attr"`
	M_definitionsPath	string	`xml:"definitionsPath,attr"`
	M_schemasPath	string	`xml:"schemasPath,attr"`
	M_tmpPath	string	`xml:"tmpPath,attr"`
	M_binPath	string	`xml:"binPath,attr"`

}
/** UUID **/
func (this *ServerConfiguration) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ServerConfiguration) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ServerConfiguration) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** HostName **/
func (this *ServerConfiguration) GetHostName() string{
	return this.M_hostName
}

/** Init reference HostName **/
func (this *ServerConfiguration) SetHostName(ref interface{}){
	this.NeedSave = true
	this.M_hostName = ref.(string)
}

/** Remove reference HostName **/

/** Ipv4 **/
func (this *ServerConfiguration) GetIpv4() string{
	return this.M_ipv4
}

/** Init reference Ipv4 **/
func (this *ServerConfiguration) SetIpv4(ref interface{}){
	this.NeedSave = true
	this.M_ipv4 = ref.(string)
}

/** Remove reference Ipv4 **/

/** Port **/
func (this *ServerConfiguration) GetPort() int{
	return this.M_port
}

/** Init reference Port **/
func (this *ServerConfiguration) SetPort(ref interface{}){
	this.NeedSave = true
	this.M_port = ref.(int)
}

/** Remove reference Port **/

/** ApplicationsPath **/
func (this *ServerConfiguration) GetApplicationsPath() string{
	return this.M_applicationsPath
}

/** Init reference ApplicationsPath **/
func (this *ServerConfiguration) SetApplicationsPath(ref interface{}){
	this.NeedSave = true
	this.M_applicationsPath = ref.(string)
}

/** Remove reference ApplicationsPath **/

/** DataPath **/
func (this *ServerConfiguration) GetDataPath() string{
	return this.M_dataPath
}

/** Init reference DataPath **/
func (this *ServerConfiguration) SetDataPath(ref interface{}){
	this.NeedSave = true
	this.M_dataPath = ref.(string)
}

/** Remove reference DataPath **/

/** ScriptsPath **/
func (this *ServerConfiguration) GetScriptsPath() string{
	return this.M_scriptsPath
}

/** Init reference ScriptsPath **/
func (this *ServerConfiguration) SetScriptsPath(ref interface{}){
	this.NeedSave = true
	this.M_scriptsPath = ref.(string)
}

/** Remove reference ScriptsPath **/

/** DefinitionsPath **/
func (this *ServerConfiguration) GetDefinitionsPath() string{
	return this.M_definitionsPath
}

/** Init reference DefinitionsPath **/
func (this *ServerConfiguration) SetDefinitionsPath(ref interface{}){
	this.NeedSave = true
	this.M_definitionsPath = ref.(string)
}

/** Remove reference DefinitionsPath **/

/** SchemasPath **/
func (this *ServerConfiguration) GetSchemasPath() string{
	return this.M_schemasPath
}

/** Init reference SchemasPath **/
func (this *ServerConfiguration) SetSchemasPath(ref interface{}){
	this.NeedSave = true
	this.M_schemasPath = ref.(string)
}

/** Remove reference SchemasPath **/

/** TmpPath **/
func (this *ServerConfiguration) GetTmpPath() string{
	return this.M_tmpPath
}

/** Init reference TmpPath **/
func (this *ServerConfiguration) SetTmpPath(ref interface{}){
	this.NeedSave = true
	this.M_tmpPath = ref.(string)
}

/** Remove reference TmpPath **/

/** BinPath **/
func (this *ServerConfiguration) GetBinPath() string{
	return this.M_binPath
}

/** Init reference BinPath **/
func (this *ServerConfiguration) SetBinPath(ref interface{}){
	this.NeedSave = true
	this.M_binPath = ref.(string)
}

/** Remove reference BinPath **/

/** Parent **/
func (this *ServerConfiguration) GetParentPtr() *Configurations{
	return this.m_parentPtr
}

/** Init reference Parent **/
func (this *ServerConfiguration) SetParentPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_parentPtr = ref.(string)
	}else{
		this.m_parentPtr = ref.(*Configurations)
		this.M_parentPtr = ref.(*Configurations).GetUUID()
	}
}

/** Remove reference Parent **/
func (this *ServerConfiguration) RemoveParentPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(*Configurations)
	if toDelete.GetUUID() == this.m_parentPtr.GetUUID() {
		this.m_parentPtr = nil
		this.M_parentPtr = ""
	}
}