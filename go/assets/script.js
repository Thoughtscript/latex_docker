const XHR = function (url) {
    return new Promise(function (resolve, reject) {
        let xhr = new XMLHttpRequest()

        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 201) return resolve(xhr.response)
                else return reject("Request failed - please try again!")
            }
        }

        xhr.open('POST', url, true)
        xhr.setRequestHeader('Content-Type', 'application/json')
        xhr.withCredentials = false

        xhr.send()
    })
}


window.onload = function () {

    document.getElementById("recompile").onclick = function (e) {
        e.preventDefault()

        XHR("https://localhost/api/pdf/make").then(
            function (success) {
                const r = document.getElementById("result")
                const n = document.createTextNode(success)
                console.log(success)
                r.appendChild(n)
                const br = document.createElement("br")
                r.appendChild(br)
            },
            function (failure) {
                const r = document.getElementById("result")
                const n = document.createTextNode(failure)
                console.log(failure)
                r.appendChild(n)
                const br = document.createElement("br")
                r.appendChild(br)
            }
        )
    }
}