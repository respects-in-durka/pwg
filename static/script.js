const editor = ace.edit("editor");
ace.config.set('basePath', 'path');
editor.session.setMode("ace/mode/golang");

function run() {
    fetch("/eval", {
        method: "POST",
        body: editor.getValue()
    }).then(data => data.text()).then(text => {
        let element = document.createElement("div")
        element.setAttribute('id', 'response')
        document.getElementsByClassName("response")[0].appendChild(element)
        if (text.startsWith("<html>")) {
            let start = text.indexOf("<pre");
            let end = text.lastIndexOf("</pre>") + 6;
            document.getElementById("response").innerHTML = text.substring(start, end)
        } else {
            document.getElementById("response").innerText = text;
        }
    })
}
