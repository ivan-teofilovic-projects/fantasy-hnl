const registerButton = document.getElementById("registerButton")
const loginButton = document.getElementById("loginButton")

function loginData() {
    var data = {}
    const inputValues = document.getElementsByClassName("loginForm")
    for (var i = 0; i < inputValues.length; i++) {
        data[inputValues[i].name] = inputValues[i].value
    }
    for (var i = 0; i < inputValues.length; i++) {
        inputValues[i].value = ""
    }
    console.log
    postData('/signin', data)
}

async function postData(url = '', data = {}) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data) 
    })
    return response
}

function sendData() {
    var data = {}
    const inputValues = document.getElementsByClassName("inputForm")
    for (var i = 0; i < inputValues.length; i++) {
        data[inputValues[i].name] = inputValues[i].value
    }
    for (var i = 0; i < inputValues.length; i++) {
        inputValues[i].value = ""
    }

    postData('/register', data)
}

async function verifyUser() {
    const response = await fetch('/verification.html')
    return response
}

registerButton.addEventListener('click', function(){
    sendData()
    verifyUser()
})

loginButton.addEventListener('click', function(){
    loginData()
})