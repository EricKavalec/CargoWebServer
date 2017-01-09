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
 * @fileOverview Session managemement functionality.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The session manager regroups functionalities relatead to session such as
 * login, logout and users sessions informations.
 * @constructor
 * @extends EventManager
 */
var SessionManager = function (id) {

    // Keeps a list of user's' sessions localy.
    this.sessions = {}

    /**
     * @property {string} activeSessionAccountId The current user acount id.
     */
    this.activeSessionAccountId = null

    if (server == undefined) {
        return
    }

    if (id == undefined) {
        id = randomUUID()
    }

    EventManager.call(this, id, SessionEvent)

    return this
}

SessionManager.prototype = new EventManager(null);
SessionManager.prototype.constructor = SessionManager;

/*
 * Dispatch event.
 */
SessionManager.prototype.onEvent = function (evt) {
    EventManager.prototype.onEvent.call(this, evt)
}

SessionManager.prototype.RegisterListener = function () {
    // Append to the event handler.
    server.eventHandler.AddEventManager(this,
        // callback
        function () {
            console.log("Session manager is registered")
        }
    )
}

/*
 * Sever side code.
 */
function Login(name, password) {
    var newSession = null
    newSession = server.GetSessionManager().Login(name, password, messageId, sessionId)
    return newSession
}

/**
 * Authenticate the user's account name and password on the server.
 * @see LoginPage
 * @param {string} name The account name
 * @param {string} password The account password
 * @param {function} successCallback The function to execute in case of authentification success
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SessionManager.prototype.login = function (name, password, successCallback, errorCallback, caller) {
    var params = []
    params.push(new RpcData({ "name": "name", "type": 2, "dataBytes": utf8_to_b64(name) }))
    params.push(new RpcData({ "name": "password", "type": 2, "dataBytes": utf8_to_b64(password) }))

    // Call it on the server.
    server.executeJsFunction(
        Login.toString(), // The function to execute remotely on server
        params, // The parameters to pass to this function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            if (result[0] != null) {
                // I will use the entity manager to initialyse the session entity.
                server.entityManager.getEntityByUuid(result[0].UUID,
                    function (session, caller) {
                        server.sessionManager.activeSessionAccountId = session.M_accountPtr
                        session["set_M_accountPtr_" + session.M_accountPtr + "_ref"](function (caller, session) {
                            return function (accountPtr) {
                                // call the callback after the session is intialysed.
                                if (accountPtr["set_M_userRef_" + accountPtr.M_userRef + "_ref"] != undefined) {
                                    accountPtr["set_M_userRef_" + accountPtr.M_userRef + "_ref"](function (session, caller) {
                                        return function () {
                                            caller.successCallback(session, caller.caller)
                                        }
                                    } (session, caller))
                                } else {
                                    caller.successCallback(session, caller.caller)
                                }
                            }
                        } (caller, session))

                    }, function (errMsg, caller) {

                    }, caller
                )
            }
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

/*
 * Sever side code.
 */
function GetActiveSessions() {
    var sessions = null
    sessions = server.GetSessionManager().GetActiveSessions()
    return sessions
}

/**
 * Retreive information about all actives users sessions on the server.
 * @param {function} successCallback When the result is send back by the server, the parameter named result contain all sessions infromations.
 * @param {function} errorCallback The function to execute in case of error.
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SessionManager.prototype.getActiveSessions = function (successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []

    // Call it on the server.
    server.executeJsFunction(
        GetActiveSessions.toString(), // The function to execute remotely on server
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

/*
 * Sever side code.
 */
function GetActiveSessionByAccountId(accountId) {
    var sessions = null
    sessions = server.GetSessionManager().GetActiveSessionByAccountId(accountId)
    return sessions
}

/**
 * Get the list of all active session on the server for a given account name
 * @param {string} accountId The account name.
 * @param {function} successCallback The function to execute in case of success.
 * @param {function} errorCallback The function to execute in case of error. 
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SessionManager.prototype.getActiveSessionByAccountId = function (accountId, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(new RpcData({ "name": "accountId", "type": 2, "dataBytes": utf8_to_b64(accountId) }))

    // Call it on the server.
    server.executeJsFunction(
        GetActiveSessionByAccountId.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            caller.successCallback(result, caller.caller)
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

/*
 * Sever side code.
 */
function UpdateSessionState(state) {
    var err = server.GetSessionManager().UpdateSessionState(state, messageId, sessionId)
    return err
}

/**
 * Change the state of a given session. 
 * @param {int} state 1: Online, 2:Away, other: Offline.
 * @param {function} successCallback The function to execute in case of success.
 * @param {function} errorCallback The function to execute in case of error. 
 * @param {object} caller A place to store object from the request context and get it back from the response context.
 */
SessionManager.prototype.updateSessionState = function (state, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(new RpcData({ "name": "state", "type": 1, "dataBytes": utf8_to_b64(state) }))

    // Call it on the server.
    server.executeJsFunction(
        UpdateSessionState.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            if (caller != null) {
                caller.successCallback(result, caller.caller)
            } else {
                caller.successCallback(result, null)
            }
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

/*
 * Sever side code.
 */
function Logout(toCloseId) {
    var err = server.GetSessionManager().Logout(toCloseId, messageId, sessionId)
    return err
}

/**
 * Close a user session from the server. A logout event is throw to inform other participant that the session
 * is closed.
 * @param {function} successCallback The function to execute in case of success.
 * @param {function} errorCallback The function to execute in case of error. 
 * @param {object} caller A place to store object from the request context and get it back from the response context. 
 */
SessionManager.prototype.logout = function (sessionId, successCallback, errorCallback, caller) {
    // server is the client side singleton.
    var params = []
    params.push(new RpcData({ "name": "name", "type": 2, "dataBytes": utf8_to_b64(sessionId) }))
    // Call it on the server.
    server.executeJsFunction(
        Logout.toString(), // The function to execute remotely on server
        params, // The parameters to pass to that function
        function (index, total, caller) { // The progress callback
            // Nothing special to do here.
        },
        function (result, caller) {
            console.log(result[0])
            caller.successCallback(result, caller.caller)
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