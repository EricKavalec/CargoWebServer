/*
* (C) Copyright 2016 Mycelius SA (http://mycelius.com/).
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

/**
* @fileOverview Events functionnalities.
* @author Dave Courtois, Eric Kavalec
* @version 1.0
*/

// Default channels and events
// TODO revoir cette section
AccountEvent = "AccountEvent"
AccountRegisterSuccessEvent = 0
AccountConfirmationSucessEvent = 1
SessionEvent = "SessionEvent"
LoginEvent = 4
LogoutEvent = 5
StateChangeEvent = 6
BpmnEvent = "BpmnEvent"
NewProcessInstanceEvent = 7
UpdateProcessInstanceEvent = 8
NewDefinitionsEvent = 9
DeleteDefinitionsEvent = 10
UpdateDefinitionsEvent = 11
EntityEvent = "EntityEvent"
NewEntityEvent = 12
UpdateEntityEvent = 13
DeleteEntityEvent = 14
OpenEntityEvent = 15
CloseEntityEvent = 16
FileEvent = "FileEvent"
NewFileEvent = 17
DeleteFileEvent = 18
UpdateFileEvent = 19
TableEvent = "TableEvent"
DeleteRowEvent = 22
NewRowEvent = 23
UpdateRowEvent = 24
SecurityEvent = "SecurityEvent"
NewRoleEvent = 25
DeleteRoleEvent = 26
UpdateRoleEvent = 27
ProjectEvent = "ProjectEvent"
EmailEvent = "EmailEvent"

/**
* EventHub contructor
* @constructor
* @param {string} channelId The id of the channel of events to manage
* @returns {EventHub}
* @stability 2
* @public true
*/
var EventHub = function (channelId) {

    this.id = randomUUID();

    this.channelId = channelId;

    this.observers = {}

    return this
}

/**
 * Register the hub as a listener to a given channel.
 */
EventHub.prototype.registerListener = function () {
    // Append to the event handler
    server.eventHandler.addEventListener(this,
        // callback
        function (listener) {
            console.log("Listener" + listener.id + "was registered to the channel ", listener.channelId)
        }
    )
}

/**
* Attach observer to a specific event.
* @param obeserver The observer to attach.
* @param eventId The event id.
* @param {function} updateFct The function to execute when the event is received.
* @stability 1
* @public true
*/
EventHub.prototype.attach = function (observer, eventId, updateFct) {
    observer.observable = this

    if (observer.id == undefined) {
        // observer needs a UUID
        observer.id = randomUUID()
    }

    if (this.observers[eventId] == undefined) {
        this.observers[eventId] = []
    }

    var observerExistsForEventId = false
    for (var i = 0; i < this.observers[eventId].length; i++) {
        if (this.observers[eventId][i].id == observer.id) {
            // only on obeserver with the same id are allowed.
            observerExistsForEventId = true
        }
    }

    if (!observerExistsForEventId) {
        this.observers[eventId].push(observer)
    }

    if (observer.updateFunctions == undefined){
        observer.updateFunctions = {}
    }

    observer.updateFunctions[this.id + "_" + eventId] = updateFct
}

/**
* Detach observer from event.
* @param obeserver The to detach
* @param eventId The event id
* @stability 1
* @public true
*/
EventHub.prototype.detach = function (observer, eventId) {
    if (observer.observable != null) {
        observer.observable = null
    }
    if (observer.updateFunctions != undefined) {
        if (observer.updateFunctions[this.id + "_" + eventId] != null) {
            delete observer.updateFunctions[this.id + "_" + eventId]
            if (Object.keys(observer.updateFunctions).length == 0) {
                this.observers[eventId].pop(observer)
            }
        }
    }
}

/**
* When an event is received, the observer callback function is called.
* @param evt The event to dispatch.
* @stability 1
* @public false
*/
EventHub.prototype.onEvent = function (evt) {
    console.log("Event received: ", evt)
    var observers = this.observers[evt.code]
    if (observers != undefined) {
        for (var i = 0; i < observers.length; i++) {
            if (observers[i].updateFunctions != undefined) {
                if (observers[i].updateFunctions[this.id + "_" + evt.code] != null) {
                    observers[i].updateFunctions[this.id + "_" + evt.code](evt, observers[i])
                } else {
                    if (Object.keys(observers[i].updateFunctions).length == 0) {
                        this.observers[eventId].pop(observers[i])
                    }
                }
            }
        }
    }
}

/**
* EventChannel constructor
* @constructor
* Each event type has its own channel.
* @param id The channel id.
* @returns {EventChannel}
* @stability 1
* @public unknown
*/
var EventChannel = function (id) {
    // The event id
    this.id = id

    this.listeners = {}

    return this
}

/**
* @param evt
* @stability 1
* @public true
*/
EventChannel.prototype.broadcastEvent = function (evt) {
    for (var l in this.listeners) {
        var listener = this.listeners[l]
        listener.onEvent(evt)
    }
}

/**
* Singleton used to receive and send events. 
* Alternative to using event listeners and channels.
* @constructor
* @returns {EventHandler}
* @stability 1
* @public true
*/
var EventHandler = function () {
    /**
     * @property channels The channel 
     */
     this.channels = {}

     return this
 }

/**
* Append a new event manager.
* @param listener The listener to append.
* @param callback The function to call when an event happen.
* @stability 1
* @public true
*/
EventHandler.prototype.addEventListener = function (listener, callback) {
    /* Add it to the local event listener **/
    if (this.channels[listener.channelId] == undefined) {
        this.channels[listener.channelId] = new EventChannel(listener.channelId)
    }
    // append the listener
    this.channels[listener.channelId].listeners[listener.id] = listener

    /* Append to the remote event listener **/
    // Create a request
    var p1 = new RpcData({ "name": "name", "type": 2, "dataBytes": utf8_to_b64(listener.channelId) });

    var params = new Array();
    params[0] = p1;
    // Register this listener to the server.
    var rqst = new Request(randomUUID(), server.conn, "RegisterListener", params,
        // Progress callback
        function () { },
        // Success callback
        function (result, listener) {
            // calling success call back function
            callback(listener);
        },
        // Error callback.
        function () {

        }, listener);
    rqst.send();
}

/**
* Remove the listener and close the channel if is empty
* @param listener The listener to remove.
* @param callback The function to call when an event happen.
*/
EventHandler.prototype.removeEventManager = function (listener, callback) {
    /* Delete the local listener **/
    if (this.channels[listener.channelId] != undefined) {
        if (this.channels[listener.channelId].listeners[listener.id] != undefined) {
            delete this.channels[listener.channelId].listeners[listener.id]
        }
        if (Object.keys(this.channels[listener.channelId]).length == 0) {
            delete this.channels[listener.channelId]
        }
    }

    /* Delete the remote listener **/
    // Create a request
    var p1 = new RpcData({ "name": "name", "type": 2, "dataBytes": utf8_to_b64(listener.channelId) });

    var params = new Array();
    params[0] = p1;

    // Register this listener to the server.
    var rqst = new Request(randomUUID(), server.conn, "UnregisterListener", params, function () {
        // I will call the success call back function
        callback();
    });
    rqst.send();
}

/*
* Server side script
*/
function AppendEventFilter(filter, channelId) {
    server.GetEventManager().AppendEventFilter(filter, channelId, messageId, sessionId)
}

/**
* Append a new filter to a listener
* @param {string} filter The filter to append
* @param {string} channelId The id of the channel. 
* @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
* @param {function} errorCallback In case of error.
* @param {object} caller A place to store object from the request context and get it back from the response context.
*/
EventHandler.prototype.appendEventFilter = function (filter, channelId, successCallback, errorCallback, caller) {
    var params = []
    params.push(createRpcData(filter, "STRING", "filter"))
    params.push(createRpcData(channelId, "STRING", "channelId"))

    // Call it on the server.
    server.executeJsFunction(
        AppendEventFilter.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            // display the message in the console.
            console.log(errMsg)
            // call the immediate error callback.
            caller.errorCallback(errMsg, caller.caller)
            // dispatch the message.
            server.errorManager.onError(errMsg)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
        )
}

/**
* Broadcast an event localy over a given channel
* var evt = {"code":OpenEntityEvent, "channelId":FileEvent, "dataMap":{"fileInfo": file}}
* server.eventHandler.broadcastLocalEvent(evt)
* @param evt The event to broadcast locally
* @stability 1
* @public true
*/
EventHandler.prototype.broadcastLocalEvent = function (evt) {
    var channel = this.channels[evt.name]
    if (channel != undefined) {

        channel.broadcastEvent(evt)
    }
}

/*
* Server side script
*/
function BroadcastNetworkEvent(evtNumber, eventName, eventDatas) {
    // Call the method.
    server.GetEventManager().BroadcastEventData(evtNumber, eventName, eventDatas, messageId, sessionId)
}

/**
* Broadcast event over the network.
* @param {int} evtNumber The event number.
* @param {string} channelId The event type.
* @param {MessageData} eventDatas An array of Message Data structures.
* Here is an example To send a file open event over the network.
* var entityInfo = {"TYPENAME":"Server.MessageData", "Name":"entityInfo", "Value":file.stringify()}
* server.eventHandler.broadcastNetworkEvent(OpenEntityEvent, EntityEvent, [entityInfo], function(){}, function(){}, undefined) 
*/
EventHandler.prototype.broadcastNetworkEvent = function (evtNumber, evtName, eventDatas, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(createRpcData(evtNumber, "INTEGER", "evtNumber"))
    params.push(createRpcData(channelId, "STRING", "channelId"))
    params.push(createRpcData(eventDatas, "JSON_STR", "eventDatas"))

    // Call it on the server.
    server.executeJsFunction(
        BroadcastNetworkEvent.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
        )
}
