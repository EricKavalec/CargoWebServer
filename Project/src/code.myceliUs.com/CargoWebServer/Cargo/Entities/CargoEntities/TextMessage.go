// +build CargoEntities

package CargoEntities

import(
	"encoding/xml"
)

type TextMessage struct{

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

	/** If the entity is fully initialyse **/
	IsInit   bool

	/** members of Entity **/
	M_id string

	/** members of Message **/
	M_body string

	/** members of TextMessage **/
	M_creationTime int64
	m_fromRef *Account
	/** If the ref is a string and not an object **/
	M_fromRef string
	m_toRef *Account
	/** If the ref is a string and not an object **/
	M_toRef string
	M_title string


	/** Associations **/
	m_entitiesPtr *Entities
	/** If the ref is a string and not an object **/
	M_entitiesPtr string
}

/** Xml parser for TextMessage **/
type XsdTextMessage struct {
	XMLName xml.Name	`xml:"textMessage"`
	/** Entity **/
	M_id	string	`xml:"id,attr"`


	/** Message **/
	M_body	string	`xml:"body,attr"`


	M_fromRef	*string	`xml:"fromRef"`
	M_toRef	*string	`xml:"toRef"`
	M_title	string	`xml:"title,attr"`
	M_creationTime	int64	`xml:"creationTime,attr"`

}
/** UUID **/
func (this *TextMessage) GetUUID() string{
	return this.UUID
}

/** Id **/
func (this *TextMessage) GetId() string{
	return this.M_id
}

/** Init reference Id **/
func (this *TextMessage) SetId(ref interface{}){
	if this.M_id != ref.(string) {
		this.M_id = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Id **/

/** Body **/
func (this *TextMessage) GetBody() string{
	return this.M_body
}

/** Init reference Body **/
func (this *TextMessage) SetBody(ref interface{}){
	if this.M_body != ref.(string) {
		this.M_body = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Body **/

/** CreationTime **/
func (this *TextMessage) GetCreationTime() int64{
	return this.M_creationTime
}

/** Init reference CreationTime **/
func (this *TextMessage) SetCreationTime(ref interface{}){
	if this.M_creationTime != ref.(int64) {
		this.M_creationTime = ref.(int64)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference CreationTime **/

/** FromRef **/
func (this *TextMessage) GetFromRef() *Account{
	return this.m_fromRef
}

/** Init reference FromRef **/
func (this *TextMessage) SetFromRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_fromRef != ref.(string) {
			this.M_fromRef = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_fromRef != ref.(Entity).GetUUID() {
			this.M_fromRef = ref.(Entity).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_fromRef = ref.(*Account)
	}
}

/** Remove reference FromRef **/
func (this *TextMessage) RemoveFromRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_fromRef!= nil {
		if toDelete.GetUUID() == this.m_fromRef.GetUUID() {
			this.m_fromRef = nil
			this.M_fromRef = ""
			this.NeedSave = true
		}
	}
}

/** ToRef **/
func (this *TextMessage) GetToRef() *Account{
	return this.m_toRef
}

/** Init reference ToRef **/
func (this *TextMessage) SetToRef(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_toRef != ref.(string) {
			this.M_toRef = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_toRef != ref.(Entity).GetUUID() {
			this.M_toRef = ref.(Entity).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_toRef = ref.(*Account)
	}
}

/** Remove reference ToRef **/
func (this *TextMessage) RemoveToRef(ref interface{}){
	toDelete := ref.(Entity)
	if this.m_toRef!= nil {
		if toDelete.GetUUID() == this.m_toRef.GetUUID() {
			this.m_toRef = nil
			this.M_toRef = ""
			this.NeedSave = true
		}
	}
}

/** Title **/
func (this *TextMessage) GetTitle() string{
	return this.M_title
}

/** Init reference Title **/
func (this *TextMessage) SetTitle(ref interface{}){
	if this.M_title != ref.(string) {
		this.M_title = ref.(string)
		if this.IsInit == true {			this.NeedSave = true
		}
	}
}

/** Remove reference Title **/

/** Entities **/
func (this *TextMessage) GetEntitiesPtr() *Entities{
	return this.m_entitiesPtr
}

/** Init reference Entities **/
func (this *TextMessage) SetEntitiesPtr(ref interface{}){
	if _, ok := ref.(string); ok {
		if this.M_entitiesPtr != ref.(string) {
			this.M_entitiesPtr = ref.(string)
			if this.IsInit == true {				this.NeedSave = true
			}
		}
	}else{
		if this.M_entitiesPtr != ref.(*Entities).GetUUID() {
			this.M_entitiesPtr = ref.(*Entities).GetUUID()
			if this.IsInit == true {				this.NeedSave = true
			}
		}
		this.m_entitiesPtr = ref.(*Entities)
	}
}

/** Remove reference Entities **/
func (this *TextMessage) RemoveEntitiesPtr(ref interface{}){
	toDelete := ref.(*Entities)
	if this.m_entitiesPtr!= nil {
		if toDelete.GetUUID() == this.m_entitiesPtr.GetUUID() {
			this.m_entitiesPtr = nil
			this.M_entitiesPtr = ""
			this.NeedSave = true
		}
	}
}
