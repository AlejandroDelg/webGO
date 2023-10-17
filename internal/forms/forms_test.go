package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("first_name", "Alex")

	r.Form = postedData
	form := New(r.Form)

	err := form.Has("first_name")
	if err == false {
		t.Error("field not found")
	}
}

func TestForm_NotHas(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	postedData.Add("first_name", "Alex")

	r.Form = postedData
	form := New(r.Form)

	err := form.Has("last_name")
	if err != false {
		t.Error("field found")
	}
}
func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatisthis", nil)
	postedData := url.Values{}
	postedData.Add("first_name", "Alex")
	postedData.Add("email", "aaa@a.com")

	r.PostForm = postedData

	form := New(r.PostForm)
	if !(form.IsEmail("aaa@a.com")) {
		t.Error("form doest have email")
	}

	r = httptest.NewRequest("POST", "/whatisthis", nil)
	postedData = url.Values{}
	postedData.Add("email", "aaa")
	r.PostForm = postedData
	form = New(r.PostForm)
	if form.IsEmail("aaa") {
		t.Error("incorrect email")
	}
}
func TestForm_MinLength(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatisthis", nil)
	postedData := url.Values{}
	postedData.Add("first_name", "Alex")

	r.Form = postedData
	form := New(r.Form)

	err := form.MinLength("first_name", 3)

	if err == false {
		t.Error("error in min length")
	}
}
