package BPMN20

import(
"encoding/xml"
)

type Resource struct{

	/** The entity UUID **/
	UUID string
	/** The entity TypeName **/
	TYPENAME string
	/** If the entity value has change... **/
	NeedSave bool

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of BaseElement **/
	M_id string
	m_other interface{}
	/** If the ref is a string and not an object **/
	M_other string
	M_extensionElements *ExtensionElements
	M_extensionDefinitions []*ExtensionDefinition
	M_extensionValues []*ExtensionAttributeValue
	M_documentation []*Documentation

	/** members of RootElement **/
	/** No members **/

	/** members of Resource **/
	M_name string
	M_resourceParameter []*ResourceParameter


	/** Associations **/
	m_activityResourcePtr []ResourceRole
	/** If the ref is a string and not an object **/
	M_activityResourcePtr []string
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
	m_definitionsPtr *Definitions
	/** If the ref is a string and not an object **/
	M_definitionsPtr string
}

/** Xml parser for Resource **/
type XsdResource struct {
	XMLName xml.Name	`xml:"resource"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	/** RootElement **/


	M_resourceParameter	[]*XsdResourceParameter	`xml:"resourceParameter,omitempty"`
	M_name	string	`xml:"name,attr"`

}
/** UUID **/
func (this *Resource) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *Resource) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *Resource) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *Resource) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *Resource) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *Resource) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *Resource) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *Resource) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *Resource) SetExtensionDefinitions(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionDefinitionss []*ExtensionDefinition
	for i:=0; i<len(this.M_extensionDefinitions); i++ {
		if this.M_extensionDefinitions[i].GetName() != ref.(*ExtensionDefinition).GetName() {
			extensionDefinitionss = append(extensionDefinitionss, this.M_extensionDefinitions[i])
		} else {
			isExist = true
			extensionDefinitionss = append(extensionDefinitionss, ref.(*ExtensionDefinition))
		}
	}
	if !isExist {
		extensionDefinitionss = append(extensionDefinitionss, ref.(*ExtensionDefinition))
	}
	this.M_extensionDefinitions = extensionDefinitionss
}

/** Remove reference ExtensionDefinitions **/

/** ExtensionValues **/
func (this *Resource) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *Resource) SetExtensionValues(ref interface{}){
	this.NeedSave = true
	isExist := false
	var extensionValuess []*ExtensionAttributeValue
	for i:=0; i<len(this.M_extensionValues); i++ {
		if this.M_extensionValues[i].GetUUID() != ref.(*ExtensionAttributeValue).GetUUID() {
			extensionValuess = append(extensionValuess, this.M_extensionValues[i])
		} else {
			isExist = true
			extensionValuess = append(extensionValuess, ref.(*ExtensionAttributeValue))
		}
	}
	if !isExist {
		extensionValuess = append(extensionValuess, ref.(*ExtensionAttributeValue))
	}
	this.M_extensionValues = extensionValuess
}

/** Remove reference ExtensionValues **/

/** Documentation **/
func (this *Resource) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *Resource) SetDocumentation(ref interface{}){
	this.NeedSave = true
	isExist := false
	var documentations []*Documentation
	for i:=0; i<len(this.M_documentation); i++ {
		if this.M_documentation[i].GetUUID() != ref.(BaseElement).GetUUID() {
			documentations = append(documentations, this.M_documentation[i])
		} else {
			isExist = true
			documentations = append(documentations, ref.(*Documentation))
		}
	}
	if !isExist {
		documentations = append(documentations, ref.(*Documentation))
	}
	this.M_documentation = documentations
}

/** Remove reference Documentation **/
func (this *Resource) RemoveDocumentation(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	documentation_ := make([]*Documentation, 0)
	for i := 0; i < len(this.M_documentation); i++ {
		if toDelete.GetUUID() != this.M_documentation[i].GetUUID() {
			documentation_ = append(documentation_, this.M_documentation[i])
		}
	}
	this.M_documentation = documentation_
}

/** Name **/
func (this *Resource) GetName() string{
	return this.M_name
}

/** Init reference Name **/
func (this *Resource) SetName(ref interface{}){
	this.NeedSave = true
	this.M_name = ref.(string)
}

/** Remove reference Name **/

/** ResourceParameter **/
func (this *Resource) GetResourceParameter() []*ResourceParameter{
	return this.M_resourceParameter
}

/** Init reference ResourceParameter **/
func (this *Resource) SetResourceParameter(ref interface{}){
	this.NeedSave = true
	isExist := false
	var resourceParameters []*ResourceParameter
	for i:=0; i<len(this.M_resourceParameter); i++ {
		if this.M_resourceParameter[i].GetUUID() != ref.(BaseElement).GetUUID() {
			resourceParameters = append(resourceParameters, this.M_resourceParameter[i])
		} else {
			isExist = true
			resourceParameters = append(resourceParameters, ref.(*ResourceParameter))
		}
	}
	if !isExist {
		resourceParameters = append(resourceParameters, ref.(*ResourceParameter))
	}
	this.M_resourceParameter = resourceParameters
}

/** Remove reference ResourceParameter **/
func (this *Resource) RemoveResourceParameter(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	resourceParameter_ := make([]*ResourceParameter, 0)
	for i := 0; i < len(this.M_resourceParameter); i++ {
		if toDelete.GetUUID() != this.M_resourceParameter[i].GetUUID() {
			resourceParameter_ = append(resourceParameter_, this.M_resourceParameter[i])
		}
	}
	this.M_resourceParameter = resourceParameter_
}

/** ActivityResource **/
func (this *Resource) GetActivityResourcePtr() []ResourceRole{
	return this.m_activityResourcePtr
}

/** Init reference ActivityResource **/
func (this *Resource) SetActivityResourcePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_activityResourcePtr); i++ {
			if this.M_activityResourcePtr[i] == refStr {
				return
			}
		}
		this.M_activityResourcePtr = append(this.M_activityResourcePtr, ref.(string))
	}else{
		this.RemoveActivityResourcePtr(ref)
		this.m_activityResourcePtr = append(this.m_activityResourcePtr, ref.(ResourceRole))
		this.M_activityResourcePtr = append(this.M_activityResourcePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference ActivityResource **/
func (this *Resource) RemoveActivityResourcePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	activityResourcePtr_ := make([]ResourceRole, 0)
	activityResourcePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_activityResourcePtr); i++ {
		if toDelete.GetUUID() != this.m_activityResourcePtr[i].(BaseElement).GetUUID() {
			activityResourcePtr_ = append(activityResourcePtr_, this.m_activityResourcePtr[i])
			activityResourcePtrUuid = append(activityResourcePtrUuid, this.M_activityResourcePtr[i])
		}
	}
	this.m_activityResourcePtr = activityResourcePtr_
	this.M_activityResourcePtr = activityResourcePtrUuid
}

/** Lane **/
func (this *Resource) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *Resource) SetLanePtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_lanePtr); i++ {
			if this.M_lanePtr[i] == refStr {
				return
			}
		}
		this.M_lanePtr = append(this.M_lanePtr, ref.(string))
	}else{
		this.RemoveLanePtr(ref)
		this.m_lanePtr = append(this.m_lanePtr, ref.(*Lane))
		this.M_lanePtr = append(this.M_lanePtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Lane **/
func (this *Resource) RemoveLanePtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	lanePtr_ := make([]*Lane, 0)
	lanePtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_lanePtr); i++ {
		if toDelete.GetUUID() != this.m_lanePtr[i].GetUUID() {
			lanePtr_ = append(lanePtr_, this.m_lanePtr[i])
			lanePtrUuid = append(lanePtrUuid, this.M_lanePtr[i])
		}
	}
	this.m_lanePtr = lanePtr_
	this.M_lanePtr = lanePtrUuid
}

/** Outgoing **/
func (this *Resource) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *Resource) SetOutgoingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_outgoingPtr); i++ {
			if this.M_outgoingPtr[i] == refStr {
				return
			}
		}
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(string))
	}else{
		this.RemoveOutgoingPtr(ref)
		this.m_outgoingPtr = append(this.m_outgoingPtr, ref.(*Association))
		this.M_outgoingPtr = append(this.M_outgoingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Outgoing **/
func (this *Resource) RemoveOutgoingPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	outgoingPtr_ := make([]*Association, 0)
	outgoingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_outgoingPtr); i++ {
		if toDelete.GetUUID() != this.m_outgoingPtr[i].GetUUID() {
			outgoingPtr_ = append(outgoingPtr_, this.m_outgoingPtr[i])
			outgoingPtrUuid = append(outgoingPtrUuid, this.M_outgoingPtr[i])
		}
	}
	this.m_outgoingPtr = outgoingPtr_
	this.M_outgoingPtr = outgoingPtrUuid
}

/** Incoming **/
func (this *Resource) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *Resource) SetIncomingPtr(ref interface{}){
	this.NeedSave = true
	if refStr, ok := ref.(string); ok {
		for i:=0; i < len(this.M_incomingPtr); i++ {
			if this.M_incomingPtr[i] == refStr {
				return
			}
		}
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(string))
	}else{
		this.RemoveIncomingPtr(ref)
		this.m_incomingPtr = append(this.m_incomingPtr, ref.(*Association))
		this.M_incomingPtr = append(this.M_incomingPtr, ref.(BaseElement).GetUUID())
	}
}

/** Remove reference Incoming **/
func (this *Resource) RemoveIncomingPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	incomingPtr_ := make([]*Association, 0)
	incomingPtrUuid := make([]string, 0)
	for i := 0; i < len(this.m_incomingPtr); i++ {
		if toDelete.GetUUID() != this.m_incomingPtr[i].GetUUID() {
			incomingPtr_ = append(incomingPtr_, this.m_incomingPtr[i])
			incomingPtrUuid = append(incomingPtrUuid, this.M_incomingPtr[i])
		}
	}
	this.m_incomingPtr = incomingPtr_
	this.M_incomingPtr = incomingPtrUuid
}

/** Definitions **/
func (this *Resource) GetDefinitionsPtr() *Definitions{
	return this.m_definitionsPtr
}

/** Init reference Definitions **/
func (this *Resource) SetDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_definitionsPtr = ref.(string)
	}else{
		this.m_definitionsPtr = ref.(*Definitions)
		this.M_definitionsPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Definitions **/
func (this *Resource) RemoveDefinitionsPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_definitionsPtr.GetUUID() {
		this.m_definitionsPtr = nil
		this.M_definitionsPtr = ""
	}
}
