package core

import (
	"fmt"
	"strings"
)

func ErrorsToJSON(err []error) []string {
	var errors []string
	for _, e := range err {
		errors = append(errors, e.Error())
	}
	return errors
}

func GenerateOperationId(path string, controllerName string, method string, version string) string {
	pathParts := strings.Split(path, "/")
	pathString := ""
	for _, pathPart := range pathParts {
		pathString += Capitalize(pathPart)
	}

	return fmt.Sprintf("%s%s%s%s", Capitalize(version), Capitalize(controllerName), Capitalize(method), pathString)
}

func GenerateFullPath(path string, version string, controllerPath string) string {
	return fmt.Sprintf("%s%s%s", AppendSlashLeft(version), AppendSlashLeft(controllerPath), AppendSlashLeft(path))
}

func AppendSlashLeft(s string) string {
	if s == "" || s == "/" {
		return ""
	}
	if strings.HasPrefix(s, "/") {
		return s
	}
	return "/" + s
}

func Capitalize(s string) string {
	if len(s) <= 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(string(s[1:]))
}
