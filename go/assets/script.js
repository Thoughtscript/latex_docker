const XHR = function (url, data=undefined, ismultipart=false) {
    return new Promise(function (resolve, reject) {
        let xhr = new XMLHttpRequest()

        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 201 || xhr.status === 200) return resolve(xhr.response)
                else return reject("Request failed - please try again!")
            }
        }

        xhr.open('POST', url, true)
        if (!ismultipart) xhr.setRequestHeader('Content-Type','application/json')
        xhr.withCredentials = false

        if (data) xhr.send(data)
        else xhr.send()
    })
}

const get_el = (id, func) => {
    const EL = document.getElementById(id)
    if (EL) EL.onclick = func
}

window.onload = function () {

    /*
     * Button handler for "Compile and parse raw LaTeX file in Docker
     * Container into viewable PDF" event
     */

    const R = function (e) {
            e.preventDefault()

            XHR("https://localhost/api/pdf/make").then(
                function (success) {
                const r = document.getElementById("result")
                const n = document.createTextNode(success)
                console.log(success)
                r.appendChild(n)
                const br = document.createElement("br")
                r.appendChild(br)

                // Force rerender
                document.getElementById("pdfviewer").src=""
                document.getElementById("pdfviewer").src="paper.pdf"
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
    
    get_el("recompile", R)

    /*
     * Button handler for "Rendering LaTeX in-browser"
     */

   const RR = function(e) {
        e.preventDefault()

        var els = document.getElementById("rawlatexwrapper").childNodes
        els[els.length - 1].remove()

        var text = document.getElementById("rawlatex").value
        var generator = new latexjs.HtmlGenerator({ hyphenate: false })
        generator = latexjs.parse(text, { generator: generator })
        document.getElementById("rawlatexwrapper").appendChild(generator.domFragment())
    }

    get_el("render", RR)

    /*
     * Button handler for "Upload pdf file" event
     */

    const U = function(e) {
        e.preventDefault()

        /*
          https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest_API/Using_FormData_Objects
          Although the above says not to use: "multipart/form-data"

          It's required by GoLang r.ParseMultipartForm
          However, explicitly setting the "Content-Type" to "multipart/form-data"

          Will result in multipart boundary issues - the solution: 
          don't set Content-Type when using FormData
          (the browser will automatically add the correct header w/ boundary)
        */
        var data = new FormData(document.getElementById("fileform"))

         XHR("https://localhost:443/api/pdf/upload", data, true).then(
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

    get_el("uploadbtn", U)

    /*
     * Button handler for "Save raw LaTeX to Docker file" event
     */

    const S = function(e) {
        e.preventDefault()

        var text = document.getElementById("rawlatex").value

        XHR("https://localhost/api/latex/save", text).then(
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

    get_el("save", S)
}