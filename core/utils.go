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

func ExtractPathParameters(path string) []string {
	pathParts := strings.Split(path, "/")
	parameters := []string{}
	for _, pathPart := range pathParts {
		pathPart = strings.ReplaceAll(pathPart, "/", "")
		if strings.HasPrefix(pathPart, ":") {
			parameters = append(parameters, pathPart[1:])
		}
	}

	return parameters
}

func GenerateOperationId(path string, controllerName string, method HttpMethods, version string) string {
	pathParts := strings.Split(path, "/")
	pathString := ""
	for _, pathPart := range pathParts {
		pathString += Capitalize(strings.ReplaceAll(pathPart, ":", ""))
	}

	return fmt.Sprintf("%s%s%s%s", Capitalize(version), Capitalize(controllerName), Capitalize(string(method)), pathString)
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

func ExtractSchemaName(typeName string) string {
	parts := strings.Split(typeName, ".")
	schemaName := ""
	for _, part := range parts {
		schemaName += Capitalize(part)
	}
	return schemaName
}
