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
 * @fileOverview The definiton of Server class.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The sever class is a representation of a distant server into a local object.
 * It regroup other manager object to give access to different services like file management, data manager etc.
 * @param {string} hostName The domain name of the server.
 * @param {string} ipv4 The ip v4 adress of the server.
 * @param {int} port The port number of the sever.
 * @constructor
 */
var Server = function (hostName, ipv4, port) {

    /**
     * @property {string} hostName The server side domain name.
     * @example www.cargoWebserver.com
     * @example localhost
     */
    this.hostName = hostName

    /**
     * @property {string} ipv4 the ip adress of the server.
     * @example 127.0.0.1
     */
    this.ipv4 = ipv4

    /**
     * @property {int} port the port number.
     * @example 8080
     */
    this.port = port

    /**
     * @property {string} sessionId Identify the client with the server.
     */
    this.sessionId = null

    /**
     * @property {object} conn The web socket connection between the client and the server. 
     */
    this.conn = null

    /**
     * @property {LanguageManager} languageManager Manage the language in GUI. 
     */
    this.languageManager = new LanguageManager()
    this.languageManager.appendLanguageInfo(languageInfo)

    /**
     * @property {EventHandler} eventHandler Manage the network event. 
     */
    this.eventHandler = new EventHandler()

    /**
     * @property {ErrorManager} errorManager Manage errors 
     */
    this.errorManager = null

    /**
     * @property {AccountManager} accountManager Accounts related functions. 
     */
    this.accountManager = null

    /**
     * @property {SessionManager} sessionManager Users sessions related functions. 
     */
    this.sessionManager = null

    /**
     * @property {fileManager} fileManager File manupulation related functions
     */
    this.fileManager = null

    /**
     * @property {entityManager} entityManager Object persistence functionalities
     */
    this.entityManager = null

    /**
     * @property {dataManager} dataManager Database access functionalities
     */
    this.dataManager = null

    /**
     * @property {emailManager} emailManager Email functionalities
     */
    this.emailManager = null

    /**
     * @property {securityManager} securityManager Security functionalities
     */
    this.securityManager = null


    this.projectManager = null
    this.workflowManager = null

    return this;
}

/**
 * That function is use to execute javascript function on the server side and get the result back
 * if there any.
 * @param {string} sessionId  The active session id.
 * @param {string} functionSrc  The source code of the function itself.
 * @param {array} functionParams  The list of function parameters.
 * @param {callback} progressCallback The function called when information is transfer from the server to client.
 * @param {callback} successCallback The success function to execute if the function success
 * @param {callback} errorCallback Call in case of error.
 * @param {object} caller This is the caller object of the function, it can be use full if other action are to be call on it.
 */
Server.prototype.executeJsFunction = function (functionSrc, functionParams, progressCallback, successCallback, errorCallback, caller) {
    var params = []
    params.push(new RpcData({ "name": "functionSrc", "type": 2, "dataBytes": utf8_to_b64(functionSrc) }))

    for (var i = 0; i < functionParams.length; i++) {
        params.push(functionParams[i])
    }

    var rqst = new Request(randomUUID(), server.conn, "ExecuteJsFunction", params,
        // Progress call back
        function (index, total, caller) {
            progressCallback(index, total, caller)
        },
        // success call back
        function (id, results, caller) {
            // Format the results.
            var result = results["result"]
            // Call the successCallback whit the information.
            successCallback(result, caller)
            return true;
        },
        // Error call back
        function (errMsg, caller) {
            errorCallback(errMsg, caller)
        },
        caller
    );
    rqst.send()
}

/*
 * Handle the messages received.
 */
Server.prototype.handleMessage = function (conn, data) {

    // I will decode the message.
    var msg = RpcMessage.decode(data);

    // Here the sever ask for something.
    if (msg.type === 0) {
        // Request
        //console.log("request receive.");
        // I will create the Rpc request object.
        var request = new Request(msg.rqst.id, conn, msg.rqst.method, msg.rqst.params, null, null, null);
        request.process()
    }
    // Here I receive an answer from the server.
    else if (msg.type === 1) {
        // Response
        //console.log("Response received.");
        var response = new Response(msg.rsp.id, conn, msg.rsp.results, null, null, null);
        if (pendingMessage[response.id] != undefined) {
            processPendingMessage(response.id)
        } else {
            var rqst = pendingRequest[response.id]
            if (rqst != undefined) {
                delete pendingRequest[rqst.id];
                response.execute(rqst);
            }
        }

    }
    // I receive an error from the server.
    else if (msg.type === 2) {
        // error
        //console.log("Error received.");
        // Create the error message.
        var err = new ErrorMsg(msg.err.id, conn, msg.err.code, msg.err.message, msg.err.data, null, null, null)
        err.catch()

    }
    // I receive event from the server.
    else if (msg.type === 3) {
        // event
        //console.log("Event received.");
        var evt = new EventMsg(msg.evt.id, conn, msg.evt.name, msg.evt.code, msg.evt.evtData, null, null, null)
        server.eventHandler.BroadcastEvent(evt)
    }
    // This is a transfer message use whit chunk message.
    else if (msg.type === 4) {

        // Here I need to save the message into the pending message map.
        if (pendingMessage[msg.id] == undefined) {
            // I will create the array first.
            pendingMessage[msg.id] = new Array()
        }

        // Now I will save the message.
        pendingMessage[msg.id][msg.index] = msg

        if (pendingRequest[msg.id].progressCallback != null) {
            pendingRequest[msg.id].progressCallback(msg.index, msg.total, pendingRequest[msg.id].caller)
        }

        // I will now create the answer for that transfer message and send it back
        // to the server.
        var results = new Array()
        var resp = new Response(msg.id, conn, results, null, null, null)
        resp.send()

        // If the message is fully transfer.
        if (msg.index == msg.total - 1) {
            // Here is the case of fragmented response. so I need to reconstruct the response from
            // the data contain in the assembled message.
            var data = []
            for (var i = 0; i < pendingMessage[msg.id].length; i++) {
                var begin = pendingMessage[msg.id][i].data.offset;
                var end = pendingMessage[msg.id][i].data.limit;
                data[i] = pendingMessage[msg.id][i].data.view.subarray(begin, end)
            }

            // Now I will delete the message.
            delete pendingMessage[msg.id]


            var fileReader = new FileReader();
            fileReader.onload = function () {
                if (this.result != null) {
                    var arrayBuffer = this.result;
                    server.handleMessage(server.conn, arrayBuffer);
                } else {
                    console.log("File data cannot be read!!!")
                }
            };

            fileReader.readAsArrayBuffer(new Blob(data));
        }
    }
};

/**
 * Get the current session id.
 */
Server.prototype.setSessionId = function () {
    var params = new Array();
    // Register this listener to the server.
    var rqst = new Request(randomUUID(), server.conn, "GetSessionId", params,
        // Progress callback
        function () { },
        // Success callback
        function (id, results, server) {
            // Keep the session id...
            server.sessionId = results.result

            // Each application must contain a main.
            if (main != null) {
                // Go to the main entry point
                // Append the listener for the entity.
                // The session listener.
                //server.entityManager.getEntityPrototypes("BPMN20", function(result, caller){
                //server.entityManager.getEntityPrototypes("BPMS_Runtime", function(result, caller){
                server.entityManager.getEntityPrototypes("CargoConfig", function (result, caller) {
                    server.entityManager.getEntityPrototypes("CargoEntities", function (result, caller) {
                        main()
                    }, function () {/* Error callback */ }, null)
                }, function () {/* Error callback */ }, null)
                // }, function(){/* Error callback */}, null)
                // }, function(){/* Error callback */}, null)

            } else {
                // I will show the project manager page.
                // TODOO error 404
                //alert("Error 404, Cargo handler.js line 121")
            }
        },
        // Error callback...
        function () {

        }, this);
    rqst.send();
}


/*
 * Sever side code.
 */
function SetRootPath(path) {
    // Call set root path on the server...
    server.SetRootPath(path)
}

/**
 * Set the application root path.
 * @param {string} path The application root path on the server.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
Server.prototype.setRootPath = function (rootPath, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(new RpcData({ "name": "rootPath", "type": 2, "dataBytes": utf8_to_b64(rootPath) }))

    // Call it on the server.
    server.executeJsFunction(
        SetRootPath.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

function Connect(host, port) {
    server.Connect(host, port)
}

/**
 * Create a connection to another server from the current server.
 * @param {string} host The host name of the server to connect with.
 * @param {int} port The port number of the distant host.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
Server.prototype.connect = function (host, port, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(new RpcData({ "name": "host", "type": 2, "dataBytes": utf8_to_b64(host) }))
    params.push(new RpcData({ "name": "port", "type": 1, "dataBytes": utf8_to_b64(port) }))

    // Call it on the server.
    server.executeJsFunction(
        Connect.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}

function Disconnect(host, port) {
    server.Disconnect(host, port)
}

/**
 * Close the connection to another server from the current server.
 * @param {string} host The host name of the server to connect with.
 * @param {int} port The port number of the distant host.
 * @param {function} successCallback The function is call in case of success and the result parameter contain objects we looking for.
 * @param {function} errorCallback In case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
Server.prototype.disconnect = function (host, port, successCallback, errorCallback, caller) {

    // server is the client side singleton.
    var params = []
    params.push(new RpcData({ "name": "host", "type": 2, "dataBytes": utf8_to_b64(host) }))
    params.push(new RpcData({ "name": "port", "type": 1, "dataBytes": utf8_to_b64(port) }))

    // Call it on the server.
    server.executeJsFunction(
        Disconnect.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}


function Stop() {
    server.Stop()
}

/**
 * Close the server.
 */
Server.prototype.stop = function (successCallback, errorCallback, caller) {
    // server is the client side singleton...
    var params = []
    // Call it on the server.
    server.executeJsFunction(
        Stop.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            //console.log(result)
            caller.successCallback(result[0], caller.caller)
        },
        function (errMsg, caller) {
            console.log(errMsg)
            server.errorManager.onError(errMsg)
            caller.errorCallback(errMsg, caller.caller)
        }, // Error callback
        { "caller": caller, "successCallback": successCallback, "errorCallback": errorCallback } // The caller
    )
}