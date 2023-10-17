package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
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
func (f *Form) Has(field string) bool {
	if f.Get(field) != "" {
		return true
	}
	f.Errors.Add(field, "this field cannot be blank")
	return false

}

func (f *Form) Required(fields ...string) {
	for _, t := range fields {

		value := f.Get(t)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(t, "this field cannot be blank")
		}
	}
}

func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	println("x es igual a : ", x)
	if len(x) < length {
		f.Errors.Add(x, fmt.Sprintf("this field must be at least %d long", length))
		return false
	}
	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(field) {
		f.Errors.Add(field, "is not and email")
		return false
	}
	return true

}
