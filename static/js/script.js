var myCodeMirror = CodeMirror(document.getElementById("editor"), {
    value: "# Title\n\nSome test2"
});


myCodeMirror.on("change", function(cm, change) {
    console.log("Change");
});

$.get( "/kbapi/testfile", function( data ) {
    myCodeMirror.setValue(data);
});

