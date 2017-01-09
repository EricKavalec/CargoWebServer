package Persistence
import (
	"golang.org/x/net/html/charset"
	"encoding/xml"
	"strings"
	"os"
	"log"
	"path/filepath"
	"code.google.com/p/go-uuid/uuid"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
)
type PersistenceXmlFactory struct {
	m_references map[string] interface{}
	m_object map[string]map[string][]string
}

/** Initialization function from xml file **/
func (this *PersistenceXmlFactory)InitXml(inputPath string, object *CargoEntities.Entities) error{
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
	var xmlElement *CargoEntities.XsdEntities
	xmlElement = new(CargoEntities.XsdEntities)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(xmlElement); err != nil {
		return err
	}
	this.InitEntities(xmlElement, object)
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
func (this *PersistenceXmlFactory)SerializeXml(outputPath string, toSerialize *CargoEntities.Entities) error{
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
	var xmlElement *CargoEntities.XsdEntities
	xmlElement = new(CargoEntities.XsdEntities)

	this.SerialyzeEntities(xmlElement, toSerialize)
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

/** inititialisation of Entities **/
func (this *PersistenceXmlFactory) InitEntities(xmlElement *CargoEntities.XsdEntities,object *CargoEntities.Entities){
	log.Println("Initialize Entities")

	/** Init toRef **/
	object.M_entities= make([]CargoEntities.Entity,0)
	for i:=0;i<len(xmlElement.M_entities_0); i++{
		val:=new(CargoEntities.Account)
		this.InitAccount(xmlElement.M_entities_0[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init filesRef **/
	for i:=0;i<len(xmlElement.M_entities_1); i++{
		val:=new(CargoEntities.File)
		this.InitFile(xmlElement.M_entities_1[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init project **/
	for i:=0;i<len(xmlElement.M_entities_2); i++{
		val:=new(CargoEntities.Project)
		this.InitProject(xmlElement.M_entities_2[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init computerRef **/
	for i:=0;i<len(xmlElement.M_entities_3); i++{
		val:=new(CargoEntities.Computer)
		this.InitComputer(xmlElement.M_entities_3[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init membersRef **/
	for i:=0;i<len(xmlElement.M_entities_4); i++{
		val:=new(CargoEntities.User)
		this.InitUser(xmlElement.M_entities_4[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init memberOfRef **/
	for i:=0;i<len(xmlElement.M_entities_5); i++{
		val:=new(CargoEntities.Group)
		this.InitGroup(xmlElement.M_entities_5[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init logEntry **/
	for i:=0;i<len(xmlElement.M_entities_6); i++{
		val:=new(CargoEntities.LogEntry)
		this.InitLogEntry(xmlElement.M_entities_6[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init log **/
	for i:=0;i<len(xmlElement.M_entities_7); i++{
		val:=new(CargoEntities.Log)
		this.InitLog(xmlElement.M_entities_7[i],val)
		object.M_entities= append(object.M_entities, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init rolesRef **/
	object.M_roles= make([]*CargoEntities.Role,0)
	for i:=0;i<len(xmlElement.M_roles); i++{
		val:=new(CargoEntities.Role)
		this.InitRole(xmlElement.M_roles[i],val)
		object.M_roles= append(object.M_roles, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Init restriction **/
	object.M_restrictions= make([]*CargoEntities.Restriction,0)
	for i:=0;i<len(xmlElement.M_restrictions); i++{
		val:=new(CargoEntities.Restriction)
		this.InitRestriction(xmlElement.M_restrictions[i],val)
		object.M_restrictions= append(object.M_restrictions, val)

		/** association initialisation **/
		val.SetEntitiesPtr(object)
	}

	/** Entities **/
	object.M_id= xmlElement.M_id

	/** Entities **/
	object.M_name= xmlElement.M_name

	/** Entities **/
	object.M_version= xmlElement.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Account **/
func (this *PersistenceXmlFactory) InitAccount(xmlElement *CargoEntities.XsdAccount,object *CargoEntities.Account){
	log.Println("Initialize Account")

	/** Account **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref userRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	if xmlElement.M_userRef !=nil {
		if _, ok:= this.m_object[object.M_id]["userRef"]; !ok {
			this.m_object[object.M_id]["userRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["userRef"] = append(this.m_object[object.M_id]["userRef"], *xmlElement.M_userRef)
	}


	/** Init ref rolesRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_rolesRef); i++ {
		if _, ok:= this.m_object[object.M_id]["rolesRef"]; !ok {
			this.m_object[object.M_id]["rolesRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["rolesRef"] = append(this.m_object[object.M_id]["rolesRef"], xmlElement.M_rolesRef[i])
	}


	/** Init session **/
	object.M_sessions= make([]*CargoEntities.Session,0)
	for i:=0;i<len(xmlElement.M_sessions); i++{
		val:=new(CargoEntities.Session)
		this.InitSession(xmlElement.M_sessions[i],val)
		object.M_sessions= append(object.M_sessions, val)

		/** association initialisation **/
		val.SetAccountPtr(object)
	}

	/** Entity **/
	object.M_name= xmlElement.M_name

	/** Entity **/
	object.M_password= xmlElement.M_password

	/** Entity **/
	object.M_email= xmlElement.M_email
}

/** inititialisation of Restriction **/
func (this *PersistenceXmlFactory) InitRestriction(xmlElement *CargoEntities.XsdRestriction,object *CargoEntities.Restriction){
	log.Println("Initialize Restriction")
	this.m_references["Restriction"] = object
}

/** inititialisation of Group **/
func (this *PersistenceXmlFactory) InitGroup(xmlElement *CargoEntities.XsdGroup,object *CargoEntities.Group){
	log.Println("Initialize Group")

	/** Group **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref membersRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_membersRef); i++ {
		if _, ok:= this.m_object[object.M_id]["membersRef"]; !ok {
			this.m_object[object.M_id]["membersRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["membersRef"] = append(this.m_object[object.M_id]["membersRef"], xmlElement.M_membersRef[i])
	}


	/** Entity **/
	object.M_name= xmlElement.M_name
}

/** inititialisation of Role **/
func (this *PersistenceXmlFactory) InitRole(xmlElement *CargoEntities.XsdRole,object *CargoEntities.Role){
	log.Println("Initialize Role")

	/** Role **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Computer **/
func (this *PersistenceXmlFactory) InitComputer(xmlElement *CargoEntities.XsdComputer,object *CargoEntities.Computer){
	log.Println("Initialize Computer")

	/** Computer **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** OsType **/
	if xmlElement.M_osType=="##Unknown"{
		object.M_osType=CargoEntities.OsType_Unknown
	} else if xmlElement.M_osType=="##Linux"{
		object.M_osType=CargoEntities.OsType_Linux
	} else if xmlElement.M_osType=="##Windows7"{
		object.M_osType=CargoEntities.OsType_Windows7
	} else if xmlElement.M_osType=="##Windows8"{
		object.M_osType=CargoEntities.OsType_Windows8
	} else if xmlElement.M_osType=="##Windows10"{
		object.M_osType=CargoEntities.OsType_Windows10
	} else if xmlElement.M_osType=="##OSX"{
		object.M_osType=CargoEntities.OsType_OSX
	} else if xmlElement.M_osType=="##IOS"{
		object.M_osType=CargoEntities.OsType_IOS
	} else{
		object.M_osType=CargoEntities.OsType_Unknown
 	}

	/** PlatformType **/
	if xmlElement.M_platformType=="##Unknown"{
		object.M_platformType=CargoEntities.PlatformType_Unknown
	} else if xmlElement.M_platformType=="##Tablet"{
		object.M_platformType=CargoEntities.PlatformType_Tablet
	} else if xmlElement.M_platformType=="##Phone"{
		object.M_platformType=CargoEntities.PlatformType_Phone
	} else if xmlElement.M_platformType=="##Desktop"{
		object.M_platformType=CargoEntities.PlatformType_Desktop
	} else if xmlElement.M_platformType=="##Laptop"{
		object.M_platformType=CargoEntities.PlatformType_Laptop
	} else{
		object.M_platformType=CargoEntities.PlatformType_Unknown
 	}

	/** Entity **/
	object.M_name= xmlElement.M_name

	/** Entity **/
	object.M_ipv4= xmlElement.M_ipv4
}

/** inititialisation of File **/
func (this *PersistenceXmlFactory) InitFile(xmlElement *CargoEntities.XsdFile,object *CargoEntities.File){
	log.Println("Initialize File")

	/** File **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init filesRef **/
	object.M_files= make([]*CargoEntities.File,0)
	for i:=0;i<len(xmlElement.M_files); i++{
		val:=new(CargoEntities.File)
		this.InitFile(xmlElement.M_files[i],val)
		object.M_files= append(object.M_files, val)

		/** association initialisation **/
	}

	/** Entity **/
	object.M_name= xmlElement.M_name

	/** Entity **/
	object.M_path= xmlElement.M_path

	/** Entity **/
	object.M_size= xmlElement.M_size

	/** Entity **/
	object.M_modeTime= xmlElement.M_modeTime

	/** Entity **/
	object.M_isDir= xmlElement.M_isDir

	/** Entity **/
	object.M_checksum= xmlElement.M_checksum

	/** Entity **/
	object.M_data= xmlElement.M_data

	/** Entity **/
	object.M_thumbnail= xmlElement.M_thumbnail

	/** Entity **/
	object.M_mime= xmlElement.M_mime

	/** FileType **/
	if xmlElement.M_fileType=="##dbFile"{
		object.M_fileType=CargoEntities.FileType_DbFile
	} else if xmlElement.M_fileType=="##diskFile"{
		object.M_fileType=CargoEntities.FileType_DiskFile
	} else{
		object.M_fileType=CargoEntities.FileType_DbFile
 	}
}

/** inititialisation of Log **/
func (this *PersistenceXmlFactory) InitLog(xmlElement *CargoEntities.XsdLog,object *CargoEntities.Log){
	log.Println("Initialize Log")

	/** Init logEntry **/
	object.M_entries= make([]*CargoEntities.LogEntry,0)
	for i:=0;i<len(xmlElement.M_entries); i++{
		val:=new(CargoEntities.LogEntry)
		this.InitLogEntry(xmlElement.M_entries[i],val)
		object.M_entries= append(object.M_entries, val)

		/** association initialisation **/
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Session **/
func (this *PersistenceXmlFactory) InitSession(xmlElement *CargoEntities.XsdSession,object *CargoEntities.Session){
	log.Println("Initialize Session")

	/** Init ref computerRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	if xmlElement.M_computerRef !=nil {
		if _, ok:= this.m_object[object.M_id]["computerRef"]; !ok {
			this.m_object[object.M_id]["computerRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["computerRef"] = append(this.m_object[object.M_id]["computerRef"], *xmlElement.M_computerRef)
	}


	/** Session **/
	object.M_id= xmlElement.M_id

	/** SessionState **/
	if xmlElement.M_sessionState=="##Online"{
		object.M_sessionState=CargoEntities.SessionState_Online
	} else if xmlElement.M_sessionState=="##Away"{
		object.M_sessionState=CargoEntities.SessionState_Away
	} else if xmlElement.M_sessionState=="##Offline"{
		object.M_sessionState=CargoEntities.SessionState_Offline
	} else{
		object.M_sessionState=CargoEntities.SessionState_Online
 	}

	/** Session **/
	object.M_startTime= xmlElement.M_startTime

	/** Session **/
	object.M_endTime= xmlElement.M_endTime
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** inititialisation of Project **/
func (this *PersistenceXmlFactory) InitProject(xmlElement *CargoEntities.XsdProject,object *CargoEntities.Project){
	log.Println("Initialize Project")

	/** Project **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref filesRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_filesRef); i++ {
		if _, ok:= this.m_object[object.M_id]["filesRef"]; !ok {
			this.m_object[object.M_id]["filesRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["filesRef"] = append(this.m_object[object.M_id]["filesRef"], xmlElement.M_filesRef[i])
	}


	/** Entity **/
	object.M_name= xmlElement.M_name
}

/** inititialisation of User **/
func (this *PersistenceXmlFactory) InitUser(xmlElement *CargoEntities.XsdUser,object *CargoEntities.User){
	log.Println("Initialize User")

	/** User **/
	object.M_id= xmlElement.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Init ref memberOfRef **/
	if len(object.M_id) == 0 {
		object.M_id=uuid.NewRandom().String()
		this.m_references[object.M_id] = object
	}
	if _, ok:= this.m_object[object.M_id]; !ok {
		this.m_object[object.M_id]=make(map[string][]string)
	}
	for i:=0; i < len(xmlElement.M_memberOfRef); i++ {
		if _, ok:= this.m_object[object.M_id]["memberOfRef"]; !ok {
			this.m_object[object.M_id]["memberOfRef"]=make([]string,0)
		}
		this.m_object[object.M_id]["memberOfRef"] = append(this.m_object[object.M_id]["memberOfRef"], xmlElement.M_memberOfRef[i])
	}


	/** Entity **/
	object.M_firstName= xmlElement.M_firstName

	/** Entity **/
	object.M_lastName= xmlElement.M_lastName

	/** Entity **/
	object.M_middle= xmlElement.M_middle

	/** Entity **/
	object.M_email= xmlElement.M_email

	/** Entity **/
	object.M_phone= xmlElement.M_phone
}

/** serialysation of Session **/
func (this *PersistenceXmlFactory) SerialyzeSession(xmlElement *CargoEntities.XsdSession,object *CargoEntities.Session){
	if xmlElement == nil{
		return
	}

	/** Serialyze ref computerRef **/
	xmlElement.M_computerRef=&object.M_computerRef


	/** Session **/
	xmlElement.M_id= object.M_id

	/** SessionState **/
	if object.M_sessionState==CargoEntities.SessionState_Online{
		xmlElement.M_sessionState="##Online"
	} else if object.M_sessionState==CargoEntities.SessionState_Away{
		xmlElement.M_sessionState="##Away"
	} else if object.M_sessionState==CargoEntities.SessionState_Offline{
		xmlElement.M_sessionState="##Offline"
	} else{
		xmlElement.M_sessionState="##Online"
 	}

	/** Session **/
	xmlElement.M_startTime= object.M_startTime

	/** Session **/
	xmlElement.M_endTime= object.M_endTime
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Role **/
func (this *PersistenceXmlFactory) SerialyzeRole(xmlElement *CargoEntities.XsdRole,object *CargoEntities.Role){
	if xmlElement == nil{
		return
	}

	/** Role **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Restriction **/
func (this *PersistenceXmlFactory) SerialyzeRestriction(xmlElement *CargoEntities.XsdRestriction,object *CargoEntities.Restriction){
	if xmlElement == nil{
		return
	}
	this.m_references["Restriction"] = object
}

/** serialysation of User **/
func (this *PersistenceXmlFactory) SerialyzeUser(xmlElement *CargoEntities.XsdUser,object *CargoEntities.User){
	if xmlElement == nil{
		return
	}

	/** User **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref memberOfRef **/
		xmlElement.M_memberOfRef= object.M_memberOfRef


	/** Entity **/
	xmlElement.M_firstName= object.M_firstName

	/** Entity **/
	xmlElement.M_lastName= object.M_lastName

	/** Entity **/
	xmlElement.M_middle= object.M_middle

	/** Entity **/
	xmlElement.M_email= object.M_email

	/** Entity **/
	xmlElement.M_phone= object.M_phone
}

/** serialysation of Log **/
func (this *PersistenceXmlFactory) SerialyzeLog(xmlElement *CargoEntities.XsdLog,object *CargoEntities.Log){
	if xmlElement == nil{
		return
	}

	/** Serialyze LogEntry **/
	if len(object.M_entries) > 0 {
		xmlElement.M_entries= make([]*CargoEntities.XsdLogEntry,0)
	}

	/** Now I will save the value of entries **/
	for i:=0; i<len(object.M_entries);i++{
		xmlElement.M_entries=append(xmlElement.M_entries,new(CargoEntities.XsdLogEntry))
		this.SerialyzeLogEntry(xmlElement.M_entries[i],object.M_entries[i])
	}
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of Account **/
func (this *PersistenceXmlFactory) SerialyzeAccount(xmlElement *CargoEntities.XsdAccount,object *CargoEntities.Account){
	if xmlElement == nil{
		return
	}

	/** Account **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref userRef **/
	xmlElement.M_userRef=&object.M_userRef


	/** Serialyze ref rolesRef **/
		xmlElement.M_rolesRef= object.M_rolesRef


	/** Serialyze Session **/
	if len(object.M_sessions) > 0 {
		xmlElement.M_sessions= make([]*CargoEntities.XsdSession,0)
	}

	/** Now I will save the value of sessions **/
	for i:=0; i<len(object.M_sessions);i++{
		xmlElement.M_sessions=append(xmlElement.M_sessions,new(CargoEntities.XsdSession))
		this.SerialyzeSession(xmlElement.M_sessions[i],object.M_sessions[i])
	}

	/** Serialyze Message **/

	/** Now I will save the value of messages **/
	for i:=0; i<len(object.M_messages);i++{
		switch v:= object.M_messages[i].(type){
		}
	}

	/** Entity **/
	xmlElement.M_name= object.M_name

	/** Entity **/
	xmlElement.M_password= object.M_password

	/** Entity **/
	xmlElement.M_email= object.M_email
}

/** serialysation of Project **/
func (this *PersistenceXmlFactory) SerialyzeProject(xmlElement *CargoEntities.XsdProject,object *CargoEntities.Project){
	if xmlElement == nil{
		return
	}

	/** Project **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref filesRef **/
		xmlElement.M_filesRef= object.M_filesRef


	/** Entity **/
	xmlElement.M_name= object.M_name
}

/** serialysation of Computer **/
func (this *PersistenceXmlFactory) SerialyzeComputer(xmlElement *CargoEntities.XsdComputer,object *CargoEntities.Computer){
	if xmlElement == nil{
		return
	}

	/** Computer **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** OsType **/
	if object.M_osType==CargoEntities.OsType_Unknown{
		xmlElement.M_osType="##Unknown"
	} else if object.M_osType==CargoEntities.OsType_Linux{
		xmlElement.M_osType="##Linux"
	} else if object.M_osType==CargoEntities.OsType_Windows7{
		xmlElement.M_osType="##Windows7"
	} else if object.M_osType==CargoEntities.OsType_Windows8{
		xmlElement.M_osType="##Windows8"
	} else if object.M_osType==CargoEntities.OsType_Windows10{
		xmlElement.M_osType="##Windows10"
	} else if object.M_osType==CargoEntities.OsType_OSX{
		xmlElement.M_osType="##OSX"
	} else if object.M_osType==CargoEntities.OsType_IOS{
		xmlElement.M_osType="##IOS"
	} else{
		xmlElement.M_osType="##Unknown"
 	}

	/** PlatformType **/
	if object.M_platformType==CargoEntities.PlatformType_Unknown{
		xmlElement.M_platformType="##Unknown"
	} else if object.M_platformType==CargoEntities.PlatformType_Tablet{
		xmlElement.M_platformType="##Tablet"
	} else if object.M_platformType==CargoEntities.PlatformType_Phone{
		xmlElement.M_platformType="##Phone"
	} else if object.M_platformType==CargoEntities.PlatformType_Desktop{
		xmlElement.M_platformType="##Desktop"
	} else if object.M_platformType==CargoEntities.PlatformType_Laptop{
		xmlElement.M_platformType="##Laptop"
	} else{
		xmlElement.M_platformType="##Unknown"
 	}

	/** Entity **/
	xmlElement.M_name= object.M_name

	/** Entity **/
	xmlElement.M_ipv4= object.M_ipv4
}

/** serialysation of Entities **/
func (this *PersistenceXmlFactory) SerialyzeEntities(xmlElement *CargoEntities.XsdEntities,object *CargoEntities.Entities){
	if xmlElement == nil{
		return
	}

	/** Serialyze Entity **/
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_0= make([]*CargoEntities.XsdAccount,0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_1= make([]*CargoEntities.XsdFile,0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_2= make([]*CargoEntities.XsdProject,0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_3= make([]*CargoEntities.XsdComputer,0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_4= make([]*CargoEntities.XsdUser,0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_5= make([]*CargoEntities.XsdGroup,0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_6= make([]*CargoEntities.XsdLogEntry,0)
	}
	if len(object.M_entities) > 0 {
		xmlElement.M_entities_7= make([]*CargoEntities.XsdLog,0)
	}

	/** Now I will save the value of entities **/
	for i:=0; i<len(object.M_entities);i++{
		switch v:= object.M_entities[i].(type){
			case *CargoEntities.Account:
				xmlElement.M_entities_0=append(xmlElement.M_entities_0,new(CargoEntities.XsdAccount))
				this.SerialyzeAccount(xmlElement.M_entities_0[len(xmlElement.M_entities_0)-1],v)
				log.Println("Serialyze Entities:entities:Account")
			case *CargoEntities.File:
				xmlElement.M_entities_1=append(xmlElement.M_entities_1,new(CargoEntities.XsdFile))
				this.SerialyzeFile(xmlElement.M_entities_1[len(xmlElement.M_entities_1)-1],v)
				log.Println("Serialyze Entities:entities:File")
			case *CargoEntities.Project:
				xmlElement.M_entities_2=append(xmlElement.M_entities_2,new(CargoEntities.XsdProject))
				this.SerialyzeProject(xmlElement.M_entities_2[len(xmlElement.M_entities_2)-1],v)
				log.Println("Serialyze Entities:entities:Project")
			case *CargoEntities.Computer:
				xmlElement.M_entities_3=append(xmlElement.M_entities_3,new(CargoEntities.XsdComputer))
				this.SerialyzeComputer(xmlElement.M_entities_3[len(xmlElement.M_entities_3)-1],v)
				log.Println("Serialyze Entities:entities:Computer")
			case *CargoEntities.User:
				xmlElement.M_entities_4=append(xmlElement.M_entities_4,new(CargoEntities.XsdUser))
				this.SerialyzeUser(xmlElement.M_entities_4[len(xmlElement.M_entities_4)-1],v)
				log.Println("Serialyze Entities:entities:User")
			case *CargoEntities.Group:
				xmlElement.M_entities_5=append(xmlElement.M_entities_5,new(CargoEntities.XsdGroup))
				this.SerialyzeGroup(xmlElement.M_entities_5[len(xmlElement.M_entities_5)-1],v)
				log.Println("Serialyze Entities:entities:Group")
			case *CargoEntities.LogEntry:
				xmlElement.M_entities_6=append(xmlElement.M_entities_6,new(CargoEntities.XsdLogEntry))
				this.SerialyzeLogEntry(xmlElement.M_entities_6[len(xmlElement.M_entities_6)-1],v)
				log.Println("Serialyze Entities:entities:LogEntry")
			case *CargoEntities.Log:
				xmlElement.M_entities_7=append(xmlElement.M_entities_7,new(CargoEntities.XsdLog))
				this.SerialyzeLog(xmlElement.M_entities_7[len(xmlElement.M_entities_7)-1],v)
				log.Println("Serialyze Entities:entities:Log")
		}
	}

	/** Serialyze Role **/
	if len(object.M_roles) > 0 {
		xmlElement.M_roles= make([]*CargoEntities.XsdRole,0)
	}

	/** Now I will save the value of roles **/
	for i:=0; i<len(object.M_roles);i++{
		xmlElement.M_roles=append(xmlElement.M_roles,new(CargoEntities.XsdRole))
		this.SerialyzeRole(xmlElement.M_roles[i],object.M_roles[i])
	}

	/** Serialyze Restriction **/
	if len(object.M_restrictions) > 0 {
		xmlElement.M_restrictions= make([]*CargoEntities.XsdRestriction,0)
	}

	/** Now I will save the value of restrictions **/
	for i:=0; i<len(object.M_restrictions);i++{
		xmlElement.M_restrictions=append(xmlElement.M_restrictions,new(CargoEntities.XsdRestriction))
		this.SerialyzeRestriction(xmlElement.M_restrictions[i],object.M_restrictions[i])
	}

	/** Entities **/
	xmlElement.M_id= object.M_id

	/** Entities **/
	xmlElement.M_name= object.M_name

	/** Entities **/
	xmlElement.M_version= object.M_version
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}
}

/** serialysation of File **/
func (this *PersistenceXmlFactory) SerialyzeFile(xmlElement *CargoEntities.XsdFile,object *CargoEntities.File){
	if xmlElement == nil{
		return
	}

	/** File **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze File **/
	if len(object.M_files) > 0 {
		xmlElement.M_files= make([]*CargoEntities.XsdFile,0)
	}

	/** Now I will save the value of files **/
	for i:=0; i<len(object.M_files);i++{
		xmlElement.M_files=append(xmlElement.M_files,new(CargoEntities.XsdFile))
		this.SerialyzeFile(xmlElement.M_files[i],object.M_files[i])
	}

	/** Entity **/
	xmlElement.M_name= object.M_name

	/** Entity **/
	xmlElement.M_path= object.M_path

	/** Entity **/
	xmlElement.M_size= object.M_size

	/** Entity **/
	xmlElement.M_modeTime= object.M_modeTime

	/** Entity **/
	xmlElement.M_isDir= object.M_isDir

	/** Entity **/
	xmlElement.M_checksum= object.M_checksum

	/** Entity **/
	xmlElement.M_data= object.M_data

	/** Entity **/
	xmlElement.M_thumbnail= object.M_thumbnail

	/** Entity **/
	xmlElement.M_mime= object.M_mime

	/** FileType **/
	if object.M_fileType==CargoEntities.FileType_DbFile{
		xmlElement.M_fileType="##dbFile"
	} else if object.M_fileType==CargoEntities.FileType_DiskFile{
		xmlElement.M_fileType="##diskFile"
	} else{
		xmlElement.M_fileType="##dbFile"
 	}
}

/** serialysation of Group **/
func (this *PersistenceXmlFactory) SerialyzeGroup(xmlElement *CargoEntities.XsdGroup,object *CargoEntities.Group){
	if xmlElement == nil{
		return
	}

	/** Group **/
	xmlElement.M_id= object.M_id
	if len(object.M_id) > 0 {
		this.m_references[object.M_id] = object
	}

	/** Serialyze ref membersRef **/
		xmlElement.M_membersRef= object.M_membersRef


	/** Entity **/
	xmlElement.M_name= object.M_name
}