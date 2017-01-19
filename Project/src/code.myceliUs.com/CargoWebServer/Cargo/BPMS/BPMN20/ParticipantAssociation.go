package BPMN20

import(
"encoding/xml"
)

type ParticipantAssociation struct{

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

	/** members of ParticipantAssociation **/
	m_innerParticipantRef *Participant
	/** If the ref is a string and not an object **/
	M_innerParticipantRef string
	m_outerParticipantRef *Participant
	/** If the ref is a string and not an object **/
	M_outerParticipantRef string


	/** Associations **/
	m_callConversationPtr *CallConversation
	/** If the ref is a string and not an object **/
	M_callConversationPtr string
	m_collaborationPtr Collaboration
	/** If the ref is a string and not an object **/
	M_collaborationPtr string
	m_callChoreographyActivityPtr *CallChoreography
	/** If the ref is a string and not an object **/
	M_callChoreographyActivityPtr string
	m_lanePtr []*Lane
	/** If the ref is a string and not an object **/
	M_lanePtr []string
	m_outgoingPtr []*Association
	/** If the ref is a string and not an object **/
	M_outgoingPtr []string
	m_incomingPtr []*Association
	/** If the ref is a string and not an object **/
	M_incomingPtr []string
}

/** Xml parser for ParticipantAssociation **/
type XsdParticipantAssociation struct {
	XMLName xml.Name	`xml:"participantAssociation"`
	/** BaseElement **/
	M_documentation	[]*XsdDocumentation	`xml:"documentation,omitempty"`
	M_extensionElements	*XsdExtensionElements	`xml:"extensionElements,omitempty"`
	M_id	string	`xml:"id,attr"`
//	M_other	string	`xml:",innerxml"`


	M_innerParticipantRef	string	`xml:"innerParticipantRef"`
	M_outerParticipantRef	string	`xml:"outerParticipantRef"`

}
/** UUID **/
func (this *ParticipantAssociation) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *ParticipantAssociation) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *ParticipantAssociation) SetId(ref interface{}){
	this.NeedSave = true
	this.M_id = ref.(string)
}

/** Remove reference Id **/

/** Other **/
func (this *ParticipantAssociation) GetOther() interface{}{
	return this.M_other
}

/** Init reference Other **/
func (this *ParticipantAssociation) SetOther(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_other = ref.(string)
	}else{
		this.m_other = ref.(interface{})
	}
}

/** Remove reference Other **/

/** ExtensionElements **/
func (this *ParticipantAssociation) GetExtensionElements() *ExtensionElements{
	return this.M_extensionElements
}

/** Init reference ExtensionElements **/
func (this *ParticipantAssociation) SetExtensionElements(ref interface{}){
	this.NeedSave = true
	this.M_extensionElements = ref.(*ExtensionElements)
}

/** Remove reference ExtensionElements **/

/** ExtensionDefinitions **/
func (this *ParticipantAssociation) GetExtensionDefinitions() []*ExtensionDefinition{
	return this.M_extensionDefinitions
}

/** Init reference ExtensionDefinitions **/
func (this *ParticipantAssociation) SetExtensionDefinitions(ref interface{}){
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
func (this *ParticipantAssociation) GetExtensionValues() []*ExtensionAttributeValue{
	return this.M_extensionValues
}

/** Init reference ExtensionValues **/
func (this *ParticipantAssociation) SetExtensionValues(ref interface{}){
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
func (this *ParticipantAssociation) GetDocumentation() []*Documentation{
	return this.M_documentation
}

/** Init reference Documentation **/
func (this *ParticipantAssociation) SetDocumentation(ref interface{}){
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
func (this *ParticipantAssociation) RemoveDocumentation(ref interface{}){
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

/** InnerParticipantRef **/
func (this *ParticipantAssociation) GetInnerParticipantRef() *Participant{
	return this.m_innerParticipantRef
}

/** Init reference InnerParticipantRef **/
func (this *ParticipantAssociation) SetInnerParticipantRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_innerParticipantRef = ref.(string)
	}else{
		this.m_innerParticipantRef = ref.(*Participant)
		this.M_innerParticipantRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference InnerParticipantRef **/
func (this *ParticipantAssociation) RemoveInnerParticipantRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_innerParticipantRef.GetUUID() {
		this.m_innerParticipantRef = nil
		this.M_innerParticipantRef = ""
	}
}

/** OuterParticipantRef **/
func (this *ParticipantAssociation) GetOuterParticipantRef() *Participant{
	return this.m_outerParticipantRef
}

/** Init reference OuterParticipantRef **/
func (this *ParticipantAssociation) SetOuterParticipantRef(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_outerParticipantRef = ref.(string)
	}else{
		this.m_outerParticipantRef = ref.(*Participant)
		this.M_outerParticipantRef = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference OuterParticipantRef **/
func (this *ParticipantAssociation) RemoveOuterParticipantRef(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_outerParticipantRef.GetUUID() {
		this.m_outerParticipantRef = nil
		this.M_outerParticipantRef = ""
	}
}

/** CallConversation **/
func (this *ParticipantAssociation) GetCallConversationPtr() *CallConversation{
	return this.m_callConversationPtr
}

/** Init reference CallConversation **/
func (this *ParticipantAssociation) SetCallConversationPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_callConversationPtr = ref.(string)
	}else{
		this.m_callConversationPtr = ref.(*CallConversation)
		this.M_callConversationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CallConversation **/
func (this *ParticipantAssociation) RemoveCallConversationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_callConversationPtr.GetUUID() {
		this.m_callConversationPtr = nil
		this.M_callConversationPtr = ""
	}
}

/** Collaboration **/
func (this *ParticipantAssociation) GetCollaborationPtr() Collaboration{
	return this.m_collaborationPtr
}

/** Init reference Collaboration **/
func (this *ParticipantAssociation) SetCollaborationPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_collaborationPtr = ref.(string)
	}else{
		this.m_collaborationPtr = ref.(Collaboration)
		this.M_collaborationPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference Collaboration **/
func (this *ParticipantAssociation) RemoveCollaborationPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_collaborationPtr.(BaseElement).GetUUID() {
		this.m_collaborationPtr = nil
		this.M_collaborationPtr = ""
	}
}

/** CallChoreographyActivity **/
func (this *ParticipantAssociation) GetCallChoreographyActivityPtr() *CallChoreography{
	return this.m_callChoreographyActivityPtr
}

/** Init reference CallChoreographyActivity **/
func (this *ParticipantAssociation) SetCallChoreographyActivityPtr(ref interface{}){
	this.NeedSave = true
	if _, ok := ref.(string); ok {
		this.M_callChoreographyActivityPtr = ref.(string)
	}else{
		this.m_callChoreographyActivityPtr = ref.(*CallChoreography)
		this.M_callChoreographyActivityPtr = ref.(BaseElement).GetUUID()
	}
}

/** Remove reference CallChoreographyActivity **/
func (this *ParticipantAssociation) RemoveCallChoreographyActivityPtr(ref interface{}){
	this.NeedSave = true
	toDelete := ref.(BaseElement)
	if toDelete.GetUUID() == this.m_callChoreographyActivityPtr.GetUUID() {
		this.m_callChoreographyActivityPtr = nil
		this.M_callChoreographyActivityPtr = ""
	}
}

/** Lane **/
func (this *ParticipantAssociation) GetLanePtr() []*Lane{
	return this.m_lanePtr
}

/** Init reference Lane **/
func (this *ParticipantAssociation) SetLanePtr(ref interface{}){
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
func (this *ParticipantAssociation) RemoveLanePtr(ref interface{}){
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
func (this *ParticipantAssociation) GetOutgoingPtr() []*Association{
	return this.m_outgoingPtr
}

/** Init reference Outgoing **/
func (this *ParticipantAssociation) SetOutgoingPtr(ref interface{}){
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
func (this *ParticipantAssociation) RemoveOutgoingPtr(ref interface{}){
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
func (this *ParticipantAssociation) GetIncomingPtr() []*Association{
	return this.m_incomingPtr
}

/** Init reference Incoming **/
func (this *ParticipantAssociation) SetIncomingPtr(ref interface{}){
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
func (this *ParticipantAssociation) RemoveIncomingPtr(ref interface{}){
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
