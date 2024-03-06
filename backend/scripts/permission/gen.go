package main

import (
	"fmt"
	"github.com/serenize/snaker"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	SOURCE_FILE        = "../../../proto/permission/permissions.yaml"
	TEMPLATE_FILE      = "../../../proto/permission/permissions-%s.tmpl"
	TEMPLATE_FILE_NAME = "permissions-%s.tmpl"

	TARGET_GO_FILE = "../../permissions/permissions.go"
	TARGET_TS_FILE = "../../../frontend/src/permissions/index.ts"
)

func main() {
	cmdPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	var allPermissions []string
	for resource, permissions := range readFile(path.Join(cmdPath, SOURCE_FILE)) {
		for _, permission := range permissions {
			allPermissions = append(allPermissions, strings.Join([]string{UppercaseFirst(resource), UppercaseFirst(permission)}, ""))
		}
	}
	genGoFile(allPermissions, cmdPath)
	genTsFile(allPermissions, cmdPath)
}

func genGoFile(permissions []string, cmdPath string) {
	tmpl, err := template.New(fmt.Sprintf(TEMPLATE_FILE_NAME, "go")).Funcs(template.FuncMap{
		"getPermissionStr":      getPermissionStr,
		"getPermissionEnumName": getPermissionEnumName,
	}).ParseFiles(path.Join(cmdPath, fmt.Sprintf(TEMPLATE_FILE, "go")))
	if err != nil {
		panic(err)
	}
	os.Remove(path.Join(cmdPath, TARGET_GO_FILE))
	file, err := os.Create(path.Join(cmdPath, TARGET_GO_FILE))
	if err != nil {
		panic(err)
	}
	defer func() {
		file.Close()
		os.Chmod(path.Join(cmdPath, TARGET_GO_FILE), 0444)
	}()
	err = tmpl.Execute(file, map[string][]string{
		"permissions": permissions,
	})
	if err != nil {
		panic(err)
	}
}

func genTsFile(permissions []string, cmdPath string) {
	tmpl, err := template.New(fmt.Sprintf(TEMPLATE_FILE_NAME, "ts")).Funcs(template.FuncMap{
		"getPermissionStr":      getPermissionStr,
		"getPermissionEnumName": getPermissionEnumName,
	}).ParseFiles(path.Join(cmdPath, fmt.Sprintf(TEMPLATE_FILE, "ts")))
	if err != nil {
		panic(err)
	}
	os.Remove(path.Join(cmdPath, TARGET_TS_FILE))
	file, err := os.Create(path.Join(cmdPath, TARGET_TS_FILE))
	if err != nil {
		panic(err)
	}
	defer func() {
		file.Close()
		os.Chmod(path.Join(cmdPath, TARGET_TS_FILE), 0444)
	}()
	err = tmpl.Execute(file, map[string][]string{
		"permissions": permissions,
	})
	if err != nil {
		panic(err)
	}
}

func readFile(p string) map[string][]string {
	data, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}
	result := make(map[string][]string)
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func UppercaseFirst(word string) string {
	length := len(word)
	if length == 0 {
		return ""
	}
	remaining := word[1:]
	first := strings.ToUpper(string(word[0]))
	return strings.Join([]string{first, remaining}, "")
}

func getPermissionStr(name string) string {
	return strings.ReplaceAll(snaker.CamelToSnake(name), "_", "-")
}

func getPermissionEnumName(name string) string {
	return strings.ToUpper(snaker.CamelToSnake(name))
}
