Create templates for different HTML sections

## Nesting templates

Define a new template: 
``` go
{{define "TemplateName"}}
insert content here
{{end}}
```

Reference the template: 
``` go
{{template "TemplateName"}}
```

You reference the template by the template-name. 

> The naming of the file is only relevant in case you manually parse the template with the `template.ParseFiles` method