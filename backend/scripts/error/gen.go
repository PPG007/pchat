package main

import (
	"bytes"
	"fmt"
	"github.com/serenize/snaker"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	SOURCE_FILE = "../../../proto/error/errors.yaml"

	TARGET_GO_FILE = "../../errors/codes.go"
	TARGET_TS_FILE = "../../../frontend/src/errors/index.ts"
)

type ErrorDefinition struct {
	Prefix int      `yaml:"prefix"`
	Codes  []string `yaml:"codes"`
}

type ErrorGroup struct {
	Errors []Error
}

type Error struct {
	Code int
	Name string
}

func main() {
	cmdPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	var groups []ErrorGroup
	for group, definition := range readFile(path.Join(cmdPath, SOURCE_FILE)) {
		var errors []Error
		for i, code := range definition.Codes {
			errors = append(errors, Error{
				Code: cast.ToInt(fmt.Sprintf("%d%03d", definition.Prefix, i)),
				Name: strings.ToUpper(snaker.CamelToSnake(fmt.Sprintf("err_%s_%s", group, code))),
			})
		}
		groups = append(groups, ErrorGroup{
			errors,
		})
	}
	genGoFile(groups, path.Join(cmdPath, TARGET_GO_FILE))
	genTsFile(groups, path.Join(cmdPath, TARGET_TS_FILE))
}

func readFile(p string) map[string]ErrorDefinition {
	data, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}
	result := map[string]ErrorDefinition{}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	return result
}

const TEMPLATE_GO = `// Generated file, DO NOT EDIT!
package errors

{{range .Groups}}
const(
	{{- range .Errors}}
	{{.Name}} = {{.Code}}
	{{- end}}
)
{{end}}
`

func genGoFile(groups []ErrorGroup, fileName string) {
	eg := struct {
		Groups []ErrorGroup
	}{groups}
	os.Remove(fileName)
	t := template.Must(template.New("go").Parse(TEMPLATE_GO))
	buffer := bytes.NewBuffer(nil)
	err := t.Execute(buffer, eg)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	os.Chmod(fileName, 0444)
}

const TEMPLATE_TS = `// Generated file, DO NOT EDIT!
export enum ErrorCode {
  {{range .Groups}}
  {{- range .Errors}}
  {{.Name}} = {{.Code}},
  {{- end}}
  {{end}}
}

export const fromNumber = (code: number) => {
  if (code in ErrorCode) {
    return code as ErrorCode
  }
  return ErrorCode.ERR_COMMON_UNKNOWN
}
`

func genTsFile(groups []ErrorGroup, fileName string) {
	eg := struct {
		Groups []ErrorGroup
	}{groups}
	os.Remove(fileName)
	t := template.Must(template.New("ts").Parse(TEMPLATE_TS))
	buffer := bytes.NewBuffer(nil)
	err := t.Execute(buffer, eg)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	os.Chmod(fileName, 0444)
}
