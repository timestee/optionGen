package optiongen

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"myitcv.io/gogenerate"
)

type FieldType int

const (
	FieldTypeFunc FieldType = iota
	FieldTypeVar
)

type fileOptionGen struct {
	FilePath   string
	FileName   string
	PkgName    string
	ImportPath []string

	ClassList         map[string]bool
	ClassOptionFields map[string][]optionField
}

type optionField struct {
	FieldType FieldType
	Name      string
	Type      string
	Body      string
}

type templateData struct {
	ClassOptionInfo     map[string][]optionInfo
	ClassOptionTypeName map[string]string
}

type optionInfo struct {
	FieldType      FieldType
	Name           string
	OptionFuncName string
	GenOptionFunc  bool
	Type           template.HTML
	Body           template.HTML
}

func (g fileOptionGen) gen(optionWithStructName bool) {
	needGen := false
	for _, need := range g.ClassList {
		needGen = needGen || need
	}
	if !needGen {
		return
	}

	buf := BufWrite{
		buf: bytes.NewBuffer(nil),
	}

	buf.wln(fmt.Sprintf("// Code generated by %s. DO NOT EDIT.", OptionGen))
	buf.wln(fmt.Sprintf("// %s: %s", OptionGen, "github.com/timestee/optionGen"))
	buf.wln()
	buf.wf("package %v\n", g.PkgName)

	for _, importPath := range g.ImportPath {
		buf.wf("import %v\n", importPath)
	}

	tmp := templateData{
		ClassOptionInfo:     make(map[string][]optionInfo),
		ClassOptionTypeName: make(map[string]string),
	}
	for className, exist := range g.ClassList {
		if exist {
			for _, val := range g.ClassOptionFields[className] {
				name := strings.Trim(val.Name, "\"")
				funcName := "With"
				if optionWithStructName {
					funcName = funcName + strings.Title(className)
				}
				if strings.HasSuffix(funcName, "Options") {
					funcName = funcName[:len(funcName)-1]
				}
				if strings.HasSuffix(funcName, "Opts") {
					funcName = funcName[:len(funcName)-1]
				}

				funcName += strings.Title(name)
				if strings.HasPrefix(val.Type, "(") && strings.HasSuffix(val.Type, ")") {
					val.Type = val.Type[1 : len(val.Type)-1]
				}
				tmp.ClassOptionInfo[className] = append(tmp.ClassOptionInfo[className], optionInfo{
					FieldType:      val.FieldType,
					Name:           name,
					GenOptionFunc:  !strings.HasSuffix(name, "_") && !strings.HasSuffix(name, "Inner"),
					OptionFuncName: funcName,
					Type:           template.HTML(val.Type),
					Body:           template.HTML(val.Body),
				})
			}
			optionTypeName := className + "Option"
			if strings.HasSuffix(className, "Options") {
				optionTypeName = className[:len(className)-1]
			}
			if strings.HasSuffix(className, "Opts") {
				optionTypeName = className[:len(className)-1]
			}
			tmp.ClassOptionTypeName[className] = optionTypeName
		}
	}

	t := template.Must(template.New("tmp").Parse(templateTextWithPreviousSupport))

	err := t.Execute(buf.buf, tmp)
	if err != nil {
		log.Fatalf("cannot execute template: %v", err)
	}

	if strings.HasPrefix(g.FileName, "gen_") {
		g.FileName = strings.TrimLeft(g.FileName, "gen_")
	}

	genName := gogenerate.NameFile(g.FileName, OptionGen)
	source, err := goimportsBuf(buf.buf)
	if err != nil {
		log.Fatalln("goimports:", err.Error())
	}

	if err := ioutil.WriteFile(genName, source.Bytes(), 0644); err != nil {
		log.Fatalf("could not write %v: %v", genName, err)
	}
	if Verbose {
		log.Println(fmt.Sprintf("%s/%s", g.PkgName, genName))
	}
}

func goimportsBuf(buf *bytes.Buffer) (*bytes.Buffer, error) {
	out := bytes.NewBuffer(nil)
	cmd := exec.Command("goimports")
	cmd.Stdin = buf
	cmd.Stdout = out

	err := cmd.Run()

	return out, err
}

const templateTextWithPreviousSupport = `
{{- range $className, $optionList := .ClassOptionInfo }}

type {{ $className }} struct {
	{{- range $index, $option := $optionList }}
		{{ $option.Name }} {{ $option.Type }}
	{{- end }}
}

func (cc *{{ $className }}) SetOption(opt {{index $.ClassOptionTypeName $className}}) {
	_ = opt(cc)
}

func (cc *{{ $className }}) GetSetOption(opt {{index $.ClassOptionTypeName $className}}) {{index $.ClassOptionTypeName $className}} {
	return opt(cc)
}

type {{index $.ClassOptionTypeName $className}} func(cc *{{$className}}) {{index $.ClassOptionTypeName $className}}
{{ range $index, $option := $optionList }}

{{- if eq $option.GenOptionFunc true }}
	func {{$option.OptionFuncName}}(v {{$option.Type}}) {{index $.ClassOptionTypeName $className}}   { return func(cc *{{$className}}) {{index $.ClassOptionTypeName $className}} {
		previous := cc.{{$option.Name}}
		cc.{{$option.Name}} = v
		return {{$option.OptionFuncName}}(previous)
} }
{{- end }}

{{- end }}

func New{{$className}}(opts ... {{index $.ClassOptionTypeName $className}}) *{{ $className }} {
	cc := newDefault{{ $className }}()
	for _, opt := range opts  {
		_ = opt(cc)
	}
	if watchDog{{$className}} != nil {
		watchDog{{$className}}(cc)
	}
	return cc
}

func Install{{$className}}WatchDog(dog func(cc *{{$className}})) {
	watchDog{{$className}} = dog
}

var watchDog{{$className}} func(cc *{{$className}})

var default{{index $.ClassOptionTypeName $className}}s = [...]{{index $.ClassOptionTypeName $className}} {
{{- range $index, $option := $optionList }}
	{{- if eq $option.GenOptionFunc true }}
		{{- if eq $option.FieldType 0 }}
			{{$option.OptionFuncName}}({{ $option.Type }} {{ $option.Body}}),
		{{- else }}
			{{$option.OptionFuncName}}({{ $option.Body}}),
		{{- end }}
	{{- end }}
{{- end }}
	}

func newDefault{{ $className }} () *{{ $className }} {
	cc := &{{ $className }}{
{{- range $index, $option := $optionList }}
	{{- if eq $option.GenOptionFunc false }}
		{{- if eq $option.FieldType 0 }}
			{{$option.Name}}: {{ $option.Type }} {{ $option.Body}},
		{{- else }}
			{{$option.Name}}: {{ $option.Body}},
		{{- end }}
	{{- end }}
{{- end }}
	}

	for _, opt := range default{{index $.ClassOptionTypeName $className}}s  {
		_ = opt(cc)
	}

	return cc
}

{{ end }}
`

const templateText = `
{{- range $className, $optionList := .ClassOptionInfo }}
type {{ $className }} struct {
	{{- range $index, $option := $optionList }}
		{{ $option.Name }} {{ $option.Type }}
	{{- end }}
}

func (cc *{{ $className }}) ApplyOption(opt {{index $.ClassOptionTypeName $className}}) {
	opt(cc)
}

type {{index $.ClassOptionTypeName $className}} func(cc *{{$className}})
{{ range $index, $option := $optionList }}

{{- if eq $option.GenOptionFunc true }}
	func {{$option.OptionFuncName}}(v {{$option.Type}}) {{index $.ClassOptionTypeName $className}}   { return func(cc *{{$className}}) {cc.{{$option.Name}} = v } }
{{- end }}

{{- end }}

func New{{$className}}(opts ... {{index $.ClassOptionTypeName $className}}) *{{ $className }} {
	cc := newDefault{{ $className }}()
	for _, opt := range opts  {
		_ = opt(cc)
	}
	if watchDog{{$className}} != nil {
		watchDog{{$className}}(cc)
	}
	return cc
}

func Install{{$className}}WatchDog(dog {{index $.ClassOptionTypeName $className}}) {
	watchDog{{$className}} = dog
}

var watchDog{{$className}} {{index $.ClassOptionTypeName $className}}

var default{{index $.ClassOptionTypeName $className}}s = [...]{{index $.ClassOptionTypeName $className}} {
{{- range $index, $option := $optionList }}
	{{- if eq $option.GenOptionFunc true }}
		{{- if eq $option.FieldType 0 }}
			{{$option.OptionFuncName}}({{ $option.Type }} {{ $option.Body}}),
		{{- else }}
			{{$option.OptionFuncName}}({{ $option.Body}}),
		{{- end }}
	{{- end }}
{{- end }}
	}

func newDefault{{ $className }} () *{{ $className }} {
	cc := &{{ $className }}{
{{- range $index, $option := $optionList }}
	{{- if eq $option.GenOptionFunc false }}
		{{- if eq $option.FieldType 0 }}
			{{$option.Name}}: {{ $option.Type }} {{ $option.Body}},
		{{- else }}
			{{$option.Name}}: {{ $option.Body}},
		{{- end }}
	{{- end }}
{{- end }}
	}

	for _, opt := range default{{index $.ClassOptionTypeName $className}}s  {
		_ = opt(cc)
	}

	return cc
}

{{ end }}
`
