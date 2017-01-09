/*
 * The text contain in the page...
 */

var languageInfo = {
	"en": {
		"Username": "Username",
		"Password": "Password",
		"RememberMe": "Remember me",
		"Login": "Login"
	},
	"fr": {
		"Username": "Nom d'usagé",
		"Password": "Mot de passe",
		"RememberMe": "Rester connecté",
		"Login": "Se connecter"
	}
}

// Set the text...
server.languageManager.appendLanguageInfo(languageInfo)

/*
 * The login page.
 */

var LoginPage = function (loginCallback) {

	this.id = "loginPage"
	this.loginCallback = loginCallback

	// Here I will create the content...
	this.loginDiv = new Element(null, {
		"tag": "div",
		"class": "login_page",
		"style": "width: 100%; height: 100%; padding-top: 0px; overflow:hidden;",
		"id":"loginPage",
		"draggable": "false"
	})

	this.loginDiv.appendElement({ "tag": "div", "class": "login-form" }).down()

		// The username input...
		.appendElement({ "tag": "div", "class": "input-container" }).down()
		.appendElement({ "tag": "input", "id": "username_input", "class": "fancy username", "name": "username", "required": "required", "type": "text" })
		.appendElement({ "tag": "label", "id": "user_placeholder", "class": "placeholder" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-user" })
		.appendElement({ "tag": "span", "style": "margin-left:3px; display:inline;", "id": "username_label" })
		.appendElement({ "tag": "span", "style": "margin-left:3px; display:inline;", "id": "username_error" }).up()
		.appendElement({ "tag": "div", "class": "line" }).up()

		// The password
		.appendElement({ "tag": "div", "class": "input-container" }).down()
		.appendElement({ "tag": "input", "class": "fancy password", "id": "password_input", "name": "password", "required": "required", "type": "password" })
		.appendElement({ "tag": "label", "id": "password_placeholder", "class": "placeholder" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-lock" })
		.appendElement({ "tag": "span", "style": "margin-left:3px;", "id": "password_label" })
		.appendElement({ "tag": "span", "style": "margin-left:3px;", "id": "password_error" }).up()
		.appendElement({ "tag": "div", "class": "line" }).up()

		// The remeber me button...
		.appendElement({ "tag": "p" }).down()
		.appendElement({ "tag": "label", "for": "remember_me" }).down()
		.appendElement({ "tag": "input", "name": "remember_me", "type": "checkbox", "id": "remember_me", "value": "forever" })
		.appendElement({ "tag": "span", "style": "margin-left:3px;", "id": "remember_me_label" }).up().up()

		// The button
		.appendElement({ "tag": "p" }).down()
		.appendElement({ "tag": "button", "id": "login_button", "type": "reset", "class": "button blue button-login" }).down()

	// Get the reference...
	this.loginButton = this.loginDiv.getChildById("login_button")
	this.rememeberMe = this.loginDiv.getChildById("remember_me")
	this.usernameInput = this.loginDiv.getChildById("username_input")
	this.passwordInput = this.loginDiv.getChildById("password_input")

	// Set the text...
	server.languageManager.setElementText(this.loginDiv.getChildById("username_label"), "Username")
	server.languageManager.setElementText(this.loginDiv.getChildById("password_label"), "Password")
	server.languageManager.setElementText(this.loginDiv.getChildById("remember_me_label"), "RememberMe")
	server.languageManager.setElementText(this.loginButton, "Login")

	// The error display element...
	this.usernameError = this.loginDiv.getChildById("username_error")
	this.passwordError = this.loginDiv.getChildById("password_error")

	// The login success event...
	server.sessionManager.attach(this, LoginEvent, function (evt, loginPanel) {
		// I will reinit the panel here...
		/*if (evt.dataMap["entity"].TYPENAME == entityPanel.typeName) {
			if (evt.dataMap["entity"] && entityPanel.entity != null) {
				if (entityPanel.entity.UUID == evt.dataMap["entity"].UUID) {
					entityPanel.init(entityPanel.proto)
					entityPanel.setEntity(server.entityManager.entities[evt.dataMap["entity"].UUID])
				}
			}
		}*/

	})

	// The error handling...
	server.errorManager.attach(this.usernameError.element, "ACCOUNT_DOESNT_EXIST_ERROR", function (usernameError, usernameInput, passwordInput, rememeberMe) {
		return function (err) {
			server.languageManager.setElementText(usernameError.element, "ACCOUNT_DOESNT_EXIST_ERROR")
			// Display error message for 1 second and reset the inputs...
			setTimeout(function (usernameError, passwordInput, usernameInput, rememeberMe) {
				return function () {
					usernameError.element.textContent = ""
					usernameInput.element.value = ""
					passwordInput.element.value = ""
					passwordInput.element.blur()
					usernameInput.element.focus()
					rememeberMe.element.checked = false
					localStorage.removeItem('_remember_me_');
				}
			} (usernameError, passwordInput, usernameInput, rememeberMe), 1000)

		}
	} (this.usernameError, this.usernameInput, this.passwordInput, this.rememeberMe))

	server.errorManager.attach(this.passwordError.element, "PASSWORD_MISMATCH_ERROR", function (passwordError, passwordInput, rememeberMe) {
		return function (err) {
			server.languageManager.setElementText(passwordError.element, "PASSWORD_MISMATCH_ERROR")
			setTimeout(function (passwordError, passwordInput, usernameInput) {
				return function () {
					passwordError.element.textContent = ""
					passwordInput.element.value = ""
					passwordInput.element.focus()
					rememeberMe.element.checked = false
					localStorage.removeItem('_remember_me_');
				}
			} (passwordError, passwordInput, rememeberMe), 1000)
		}
	} (this.passwordError, this.passwordInput, this.rememeberMe))

	// The event handling...
	this.loginButton.element.onclick = function (usernameInput, passwordInput, loginCallback, loginPage) {
		return function () {
			/* get the value **/
			userName = usernameInput.element.value
			password = passwordInput.element.value

			/* Here I will call the login... **/
			if (password.length > 0 && userName.length > 0) {
				server.sessionManager.login(userName, password, function (result, caller) {
					caller.loginCallback(result)
					var rememberMe = localStorage.getItem('_remember_me_')
					if (rememberMe != undefined) {
						// In that case I will put the session info in the 
						var loginInfo = { "userName": caller.userName, "password": caller.password }
						// local storage...
						localStorage.setItem('_remember_me_', JSON.stringify(loginInfo))
					}
				},
					function (error, caller) {
						/* Nothing todo here */
					}, { "loginCallback": loginCallback, "loginPage": loginPage, "userName": userName, "password": password })
			} else {
				if (userName.length == 0) {
					passwordInput.element.value = ""
					usernameInput.element.focus()
				} else {
					passwordInput.element.focus()
				}
			}
		}
	} (this.usernameInput, this.passwordInput, this.loginCallback, this)

	// Now the key down event...
	this.usernameInput.element.onkeydown = function (usernameInput, passwordInput) {
		return function (e) {
			var code = (e.keyCode ? e.keyCode : e.which);
			if (code == 13 && usernameInput.element.value.length > 0) { //Enter keycode
				passwordInput.element.value = ""
				passwordInput.element.focus()
			}
		}
	} (this.usernameInput, this.passwordInput)

	// Same as login button...
	this.passwordInput.element.onkeydown = function (passwordInput, loginButton) {
		return function (e) {
			var code = (e.keyCode ? e.keyCode : e.which);
			if (code == 13 && passwordInput.element.value.length > 0) { //Enter keycode
				loginButton.element.click()
			}
		}
	} (this.passwordInput, this.loginButton)

	// Now the remember me...
	this.rememeberMe.element.onclick = function (rememeberMe) {
		return function () {
			var remember = rememeberMe.element.checked
			if (remember) {
				// Here i set rememberme on the local store...
				localStorage.setItem('_remember_me_', "");
			} else {
				// Remove the key...
				localStorage.removeItem('_remember_me_');
			}
		}
	} (this.rememeberMe)

	if (localStorage.getItem("_remember_me_") != undefined) {
		// Here I will retreive the session info create an event and send it to the server.
		var loginInfo = JSON.parse(localStorage.getItem("_remember_me_"))

		server.sessionManager.login(loginInfo.userName, loginInfo.password,
			// Success Callback
			function (result, caller) {
				caller.loginCallback(result)
			},
			// Error Callback
			function (err, caller) {

			}
			, { "loginCallback": this.loginCallback, "userName": loginInfo.userName, "password": loginInfo.password })

	}

	return this
}