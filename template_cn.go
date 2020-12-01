package autodoc

import "text/template"

var apiTemplateCN = template.Must(template.New("api").Parse(`### {{ .Name }}

{{ .Desc }}

#### 请求URL:
- {{ .URI }}

#### 请求方式:
- {{ .Method }}

#### 请求参数:

{{ .RequestModel }}

#### 请求示例:

{{ .RequestExample }}

#### 返回参数:

{{ .ResponseModel }}

#### 返回示例:

{{ .ResponseExample }}

------
`))

var headerCN = []string{"字段名", "类型", "备注"}

var typeCN = "枚举类型"
