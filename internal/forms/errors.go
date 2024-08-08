package forms

type errors map[string][]string

// Add Adds an error message for a given from field
func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message for the given field
func (e errors) Get(field string) string {
	errorString := e[field]
	if len(errorString) == 0 {
		return ""
	}
	return errorString[0]
}
