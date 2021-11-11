package main

import "html/template"

var PageLayout_Template = template.Must(template.New("PageLayout").Parse(`
<html>
<body>
{{.Render "TableOfKeys"}}
</body>
</html>
`))
