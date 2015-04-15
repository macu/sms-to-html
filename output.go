package main

import (
	"html/template"
	"io"
	"log"
)

var outTemplate = template.Must(template.New("output").Parse(`
<!doctype html>
<html lang="en-CA">

<head>
	<meta charset="utf-8">
	<title>Messages history</title>
</head>

<body>

{{range .SMSList}}
	<article class="sms">
		<span>
			SMS
		{{if .IsIncoming}}
			received on {{.ReadableDate}} from {{.ContactName}}
		{{else}}
			sent on {{.ReadableDate}} to {{.ContactName}}
		{{end}}
		</span>
		<p>{{.Body}}</p>
	</article>
{{end}}

{{range .MMSList}}
	<article class="mms">
	<span>
		MMS
	{{if .IsIncoming}}
		received on {{.ReadableDate}} from {{.ContactName}}
	{{else}}
		sent on {{.ReadableDate}} to {{.ContactName}}
	{{end}}
	</span>
	</article>
{{end}}

</body>

</html>
`))

func writeMessages(messages *Messages, output io.Writer) {
	err := outTemplate.Execute(output, messages)
	if err != nil {
		log.Fatal(err)
	}
}
