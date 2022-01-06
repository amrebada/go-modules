package core

func ErrorsToJSON(err []error) []string {
	var errors []string
	for _, e := range err {
		errors = append(errors, e.Error())
	}
	return errors
}
