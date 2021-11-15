package main

import "html/template"

var PageLayout_Template = template.Must(template.New("PageLayout").Parse(`
<html>
<head>
    <style>
        input,
        label,
        button, 
		a {
            font-size: xx-large;
        }
        .center {
            margin: auto;
            width: 50%;
            padding: 10px;
        }
    </style>
</head>
<body>
<script>
  async function fetchAndCopy(urlEscapedName) {
	let response = await fetch("/fetch?name="+urlEscapedName);
	let data = await response.text();
    copyToClipboard(data);
  }
  // Copies a string to the clipboard. Must be called from within an
  // event handler such as click. May return false if it failed, but
  // this is not always possible. Browser support for Chrome 43+,
  // Firefox 42+, Safari 10+, Edge and Internet Explorer 10+.
  // Internet Explorer: The clipboard feature may be disabled by
  // an administrator. By default a prompt is shown the first
  // time the clipboard is used (per session).
  function copyToClipboard(text) {
	if (window.clipboardData && window.clipboardData.setData) {
		// Internet Explorer-specific code path to prevent textarea being shown while dialog is visible.
		return window.clipboardData.setData("Text", text);
	} else {
		if (document.queryCommandSupported && document.queryCommandSupported("copy")) {
		  var textarea = document.createElement("textarea");
		  textarea.textContent = text;
		  textarea.style.position = "fixed";  // Prevent scrolling to bottom of page in Microsoft Edge.
		  document.body.appendChild(textarea);
		  textarea.select();
		  try {
			  return document.execCommand("copy");  // Security exception may be thrown by some browsers.
		  }
		  catch (ex) {
			  console.warn("Copy to clipboard failed.", ex);
			  return false;
		  }
		  finally {
			  document.body.removeChild(textarea);
		  }
	  	} else {
			// fallback to dialog method
			window.prompt("Copy to clipboard: Ctrl+C, Enter", text);
	  	}
	}
  }  
</script>

<script>
function lookup(what) {
	document.location = "/?q="+what
}
</script>

<div class="center">
	<input id="q" type="search">
	<button
		id="search"
		onclick="javascript:lookup(document.getElementById('q').value);">Search</button>
	{{.Render "TableOfKeys"}}	
</div>

<script>
document.getElementById("q")
    .addEventListener("keyup", function(event) {
    event.preventDefault();
    if (event.keyCode === 13) {
        document.getElementById("search").click();
    } else {
		console.log(event);
	}
});
</script>

</body>
</html>
`))
