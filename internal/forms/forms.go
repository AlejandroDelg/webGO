package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, request *http.Request) bool {
	if request.Form.Get(field) != "" {
		return true
	}
	f.Errors.Add(field, "this field cannot be blank")
	return false

}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
