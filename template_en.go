package autodoc

import "text/template"

var apiTemplateEN = template.Must(template.New("api").Parse(`### {{ .Name }}

{{ .Desc }}

#### Request URL:
- {{ .URI }}

#### Request Method:
- {{ .Method }}

#### Request Parameters:

{{ .RequestModel }}

#### Request Example:

{{ .RequestExample }}

#### Return Parameters:

{{ .ResponseModel }}

#### Response Example:

{{ .ResponseExample }}

------
`))

var headerEN = []string{"Field Name", "Type", "Remark"}

var typeEN = "Type Enum"
