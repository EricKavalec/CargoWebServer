package Server

import (
	b64 "encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"sync"

	"code.myceliUs.com/CargoWebServer/Cargo/Utility"
)

/**
 * The message processor processes the incomming message received.
 */
type MessageProcessor struct {
	// Run until abortedByEnvironment is false
	abortedByEnvironment chan bool

	// The received message.
	m_incomingChannel chan *message

	// The message to be sent..
	m_outgoingChannel chan *message

	// The pending message.
	m_pendingRequestChannel chan *message

	// This map will contain chunk of message larger
	// than the allow transfert size.
	m_pendingMsgChunk map[string][][]byte

	// The pending request will keep the request in memory
	// before a response will be receive.
	m_pendingRequest map[string]*message

	// The message(s) that waiting for response.
	m_pendingMsg map[string][]*message

	/**
	 * Use to protected the entitiesMap access.
	 */
	sync.RWMutex
}

/**
 * Create a new message processor.
 */
func newMessageProcessor() *MessageProcessor {
	p := new(MessageProcessor)

	// Incoming message channel
	p.m_incomingChannel = make(chan *message)

	// Outgoing message channel
	p.m_outgoingChannel = make(chan *message)

	// Pending message.
	p.m_pendingRequest = make(map[string]*message)
	p.m_pendingRequestChannel = make(chan *message)

	p.m_pendingMsg = make(map[string][]*message)
	p.m_pendingMsgChunk = make(map[string][][]byte)

	// Channel to stop the message proecessing.
	p.abortedByEnvironment = make(chan bool)
	return p
}

func (this *MessageProcessor) run() {
	for {
		select {
		case m := <-this.m_incomingChannel:
			// Process the incomming message.
			go this.processIncomming(m)

		case m := <-this.m_outgoingChannel:
			// Process the outgoing message.
			go this.processOutgoing(m)

		case m := <-this.m_pendingRequestChannel:
			// Process the pending message.
			go this.appendPendingRequest(m)

		case done := <-this.abortedByEnvironment:
			if done {
				return
			}
		}
	}
}

/**
 * That function determine the max message size
 */
func getMaxMessageSize() int {
	return 17740
}

//////////////////////////////////////////////////////////////////////////////////
// Synchronize access function.
//////////////////////////////////////////////////////////////////////////////////

///////////////////////////// Request functions //////////////////////////////////
func (this *MessageProcessor) setPendingRequest(m *message) {
	this.Lock()
	defer this.Unlock()

	//log.Println("append pending request with id ", m.GetId())
	this.m_pendingRequest[*m.msg.Rqst.Id] = m

}

/**
 * Always use that function to process a request, don't send the request
 * directly with the connection.
 */
func (this *MessageProcessor) appendPendingRequest(m *message) {
	this.Lock()
	defer this.Unlock()

	var isOpen = false

	for i := 0; i < len(m.to); i++ {
		if m.to[i].IsOpen() {
			isOpen = true
		}
	}

	// I will keep the reference to the request to be able
	// to made the action later.
	if isOpen {
		//log.Println("append pending request with id ", m.GetId())
		this.m_pendingRequest[*m.msg.Rqst.Id] = m
		this.m_outgoingChannel <- m
	}
}

/**
 * Get the pending request with is id.
 */
func (this *MessageProcessor) getPendingRequestById(id string) *message {
	this.Lock()
	defer this.Unlock()
	return this.m_pendingRequest[id]
}

/**
 * The map of pending request. Each action is reponsible
 * to remove it pending request from that map.
 */
func (this *MessageProcessor) isPendingRequestExist(m *message) bool {

	if this.getPendingRequestById(m.GetId()) != nil {
		return true
	}
	return false
}

/**
 * When the request is process the message can be remove from the pending request.
 */
func (this *MessageProcessor) removePendingRequest(m *message) {
	this.Lock()
	defer this.Unlock()
	delete(this.m_pendingRequest, m.GetId())
}

////////////////////////////////  Response function //////////////////////////////
/**
 * Always use that function to process a response, don't send the resonponse
 * directly with the connection.
 */
func (this *MessageProcessor) appendResponse(m *message) {
	this.Lock()
	defer this.Unlock()

	var isOpen = false

	for i := 0; i < len(m.to); i++ {
		if m.to[i].IsOpen() {
			isOpen = true
		}
	}

	// I will keep the reference to the request to be able
	// to made the action later.
	if isOpen {
		this.m_outgoingChannel <- m
	}

}

////////////////////////////// Pending message ///////////////////////////////////
func (this *MessageProcessor) createPendingMessages(m *message) {
	this.Lock()

	//Get the max size.
	maxSize := getMaxMessageSize()

	// So here I will chunk the file into smaler section and
	// send the messages to the client.
	count := len(m.GetBytes()) / maxSize

	// Round up here.
	if len(m.GetBytes())%maxSize > 0 {
		count++
	}

	id := m.GetId()

	// Create the message array
	this.m_pendingMsg[id] = make([]*message, count)
	for i := 0; i < count; i++ {
		// So here I w ill create the slice.
		var bytesSlice []byte
		var startIndex = i * maxSize

		if startIndex+maxSize < len(m.GetBytes()) {
			bytesSlice = m.GetBytes()[startIndex : startIndex+maxSize]
		} else {
			bytesSlice = m.GetBytes()[startIndex:]
		}

		transferMsg := new(message)
		transferMsg.to = m.to
		transferMsg.msg = new(Message)
		transferMsg.msg.Id = &id
		index_ := int32(i)
		total := int32(count)
		transferMsg.msg.Index = &index_
		transferMsg.msg.Total = &total

		transferMsg.msg.Type = new(Message_MessageType)
		*transferMsg.msg.Type = Message_TRANSFER

		transferMsg.msg.Data = make([]byte, len(bytesSlice))
		copy(transferMsg.msg.Data, bytesSlice)

		this.m_pendingMsg[id][i] = transferMsg
	}
	// So here I will start the pending message processing.
	this.Unlock()
	this.processPendingMessage(id)
}

func (this *MessageProcessor) isPending(id string) bool {
	this.Lock()
	defer this.Unlock()
	if _, ok := this.m_pendingMsg[id]; ok {
		return true
	}
	return false
}

/**
 * Get the pending request with is id.
 */
func (this *MessageProcessor) getPendingMessagesById(id string) []*message {
	this.Lock()
	defer this.Unlock()
	return this.m_pendingMsg[id]
}

/**
 * Remove the first messages from the pending list and retrun it
 */
func (this *MessageProcessor) popPendingMessages(id string) *message {
	this.Lock()
	defer this.Unlock()

	messages := this.m_pendingMsg[id]
	if len(messages) > 0 {
		msg := messages[0]
		if *msg.msg.Type == Message_TRANSFER {
			// I will now remove the first item from the array.
			this.m_pendingMsg[id] = make([]*message, 0)
			for i := 1; i < len(messages); i++ {
				this.m_pendingMsg[id] = append(this.m_pendingMsg[id], messages[i])
			}

			return msg
		}
	} else {
		// Remove the pendeing message from the list.
		delete(this.m_pendingMsg, id)
	}
	return nil
}

////////////////////////////// chunk message /////////////////////////////////

/**
 * Create an empty array of bytes.
 */
func (this *MessageProcessor) createChunkMessages(messageId string, container [][]byte) {
	this.Lock()
	defer this.Unlock()

	this.m_pendingMsgChunk[messageId] = container
}

/**
 * Get the chunked message.
 */
func (this *MessageProcessor) getChunkMessagesById(id string) [][]byte {
	this.Lock()
	defer this.Unlock()

	return this.m_pendingMsgChunk[id]
}

/**
 * Determine if the chunck message exist or not.
 */
func (this *MessageProcessor) isChunkMessagesExist(m *message) bool {
	if this.getChunkMessagesById(m.GetId()) != nil {
		return true
	}
	return false
}

func (this *MessageProcessor) removeChunkMessages(messageId string) {
	this.Lock()
	defer this.Unlock()
	delete(this.m_pendingMsgChunk, messageId)
}

//////////////////////////////////////////////////////////////////////////////////
// Processing functions.
//////////////////////////////////////////////////////////////////////////////////

/**
 * Process Is use to execute the action associated whit the request.
 */
func (this *MessageProcessor) processIncomming(m *message) {
	msg := m.msg

	if *msg.Type == Message_REQUEST {

		// I will create the new action
		a := newAction(msg.GetRqst().GetMethod(), m)

		// append the message to the pending request.
		this.setPendingRequest(m)

		// Now the parameters.
		for _, param := range msg.GetRqst().GetParams() {
			if param.GetType() == Data_DOUBLE {
				val, err := strconv.ParseFloat(string(param.GetDataBytes()), 64)
				if err != nil {
					panic(err)
				}
				a.Params = append(a.Params, val)
			} else if param.GetType() == Data_INTEGER {

				val, err := strconv.ParseInt(string(param.GetDataBytes()), 10, 64)
				if err != nil {
					panic(err)
				}
				a.Params = append(a.Params, val)
			} else if param.GetType() == Data_BOOLEAN {

				val, err := strconv.ParseBool(string(param.GetDataBytes()))
				if err != nil {
					panic(err)
				}
				a.Params = append(a.Params, val)
			} else if param.GetType() == Data_STRING {
				val := string(param.GetDataBytes())
				a.Params = append(a.Params, val)

			} else if param.GetType() == Data_BYTES {
				a.Params = append(a.Params, param.GetDataBytes())

			} else if param.GetType() == Data_JSON_STR {
				val := string(param.GetDataBytes())
				val_, err := b64.StdEncoding.DecodeString(val)
				if err == nil {
					val = string(val_)
				}

				// Only registered type will be process sucessfully here.
				// how the server will be able to know what to do otherwise.
				if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {
					// It contain an array of values to be init
					var values interface{}
					if param.GetTypeName() == "[]string" {
						values = make([]string, 0)
					} else {
						values = make([]interface{}, 0)
					}

					err = json.Unmarshal([]byte(val), &values)
					if err == nil {
						p, err := Utility.InitializeStructures(values.([]interface{}), param.GetTypeName())
						if err == nil {
							a.Params = append(a.Params, p.Interface())
						} else {
							//log.Println("Error:", err)
							// Here I will try to create a the array of object.
							if err.Error() == "NotDynamicObject" {
								p, err := Utility.InitializeArray(values.([]interface{}), param.GetTypeName())
								if err == nil {
									a.Params = append(a.Params, p.Interface())
								}
							}
						}
					}

				} else {
					// It contain an object.
					var valMap map[string]interface{}
					err = json.Unmarshal([]byte(val), &valMap)
					if err == nil {
						p, err := Utility.InitializeStructure(valMap)
						if err != nil {
							log.Println("Error:", err)
							a.Params = append(a.Params, valMap)
						} else {
							a.Params = append(a.Params, p.Interface())
						}
					} else {
						// I will set a nil value to the parameter in that case.
						a.Params = append(a.Params, nil)
					}
				}
			}
		}

		go a.execute()

		// Execute in it own goroutine.
		this.removePendingRequest(m)

	} else if *msg.Type == Message_RESPONSE {
		// If the response is in the pending message I will process the next message.
		if this.isPending(msg.Rsp.GetId()) {
			//do something here
			this.processPendingMessage(msg.Rsp.GetId())
		} else {
			// Here I received a response from the client so I will process it.
			rqst := this.getPendingRequestById(msg.Rsp.GetId())

			//log.Println("The response for message ", rqst.GetId(), " was succefully received!")
			this.removePendingRequest(rqst)

			// Here I will execute the successCallback if some is define.
			if rqst.successCallback != nil {
				// Call the successCallback.
				rqst.successCallback(m)
			}
		}

	} else if *msg.Type == Message_ERROR {
		err := msg.GetErr()
		log.Printf("Error ", err.GetMessage())
		// An error was encounter by the client.
	} else if *msg.Type == Message_EVENT {

		// When the client throw an event this is the place where
		// I handle it.
		evt := msg.GetEvt()
		// I will process the event.
		GetServer().GetEventManager().BroadcastEvent(evt)

	} else if *msg.Type == Message_TRANSFER {

		total := int(msg.GetTotal())
		messageId := msg.GetId()
		index := int(msg.GetIndex())
		chunk := this.getChunkMessagesById(messageId)
		if chunk != nil {
			// So here it's not the first message receive for the file.
			chunk[index] = msg.GetData()
			if index == total-1 {
				// In that case it's the last message.
				data := make([]byte, 0) // create the buffer that will contain the data.
				for i := 0; i < total; i++ {
					// Here I will try to decode the string if it'encoded.
					val_, err := b64.StdEncoding.DecodeString(string(chunk[i]))
					if err == nil {
						chunk[i] = val_
					}
					data = append(data, chunk[i]...)
				}

				// Release the memory for that message.
				this.removeChunkMessages(messageId)

				// and process the action...
				originMsg, err := NewMessageFromData(data, m.from)
				if err == nil {
					this.m_incomingChannel <- originMsg
				} else {
					log.Println("Error: ", err)
				}
			}

		} else {
			// The chunk is not there so I will insert it.
			// -- firt i will create a new array whit the necessary space.
			container := make([][]byte, total, total)
			container[0] = msg.GetData()
			this.createChunkMessages(messageId, container)

		}

		// Here I will send back an empty response to tell the other end
		// that the message is process in order to continue the transfer.
		to := make([]connection, 1)
		to[0] = m.from
		result := make([]*MessageData, 0)
		resultMsg, _ := NewResponseMessage(messageId, result, to)
		this.m_outgoingChannel <- resultMsg
	}
}

/**
 * Process outgoing message to be sent to the client.
 */
func (this *MessageProcessor) processOutgoing(m *message) {

	// Get the max message size.
	maxSize := getMaxMessageSize()

	// Here I will send the message to the client.
	if *m.msg.Type == Message_REQUEST || *m.msg.Type == Message_RESPONSE {

		// Request or Response.
		if len(m.GetBytes()) < maxSize {
			for i := 0; i < len(m.to); i++ {
				m.to[i].Send(m.GetBytes())
			}
		} else {
			// so here I will split the message in multiple part
			// and send it.
			this.createPendingMessages(m)
		}
	} else if *m.msg.Type == Message_EVENT {
		// Event
		log.Println("----------------------------------> ln 541 message processor want to send event message!")
	} else if *m.msg.Type == Message_ERROR {
		// Error
		log.Println("----------------------------------> ln 544 message processor want to send error message!")
	} else if *m.msg.Type == Message_TRANSFER {
		// Transfer
		for i := 0; i < len(m.to); i++ {
			if m.to[i].IsOpen() {
				m.to[i].Send(m.GetBytes())
			}
		}
	}
}

/**
 * Process pending message one by one.
 */
func (this *MessageProcessor) processPendingMessage(id string) {

	msg := this.popPendingMessages(id)
	if msg != nil {
		this.m_outgoingChannel <- msg
	}
}