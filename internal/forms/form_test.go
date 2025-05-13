package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("Form failed validation when it was valid")
	}

	form.Errors.Add("name", "This field cannot be blank")
	isValid = form.Valid()
	if isValid {
		t.Error("Form passed validation when it was invalid")
	}
}

func TestForm_Required(t *testing.T) {

	r := httptest.NewRequest("POST", "/", nil)
	form := New(r.PostForm)

	form.Required("name", "number")

	if form.Valid() {
		t.Error("Form shows required when fields missing")
	}

	postedData := url.Values{}
	postedData.Set("name", "John")
	postedData.Set("number", "42")

	r, _ = http.NewRequest("POST", "/", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("name", "number")
	if !form.Valid() {
		t.Error("Form does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {

	r, _ := http.NewRequest("POST", "/", nil)
	form := New(r.PostForm)

	has := form.Has("age")
	if has {
		t.Error("Form reports as containing a field it does not")
	}

	postedData := url.Values{}
	postedData.Add("name", "John")
	form = New(postedData)
	has = form.Has("name")
	if !has {
		t.Error("Form reports as not containing a field it actually contains")
	}
}

func TestForm_MinLength(t *testing.T) {

	postedData := url.Values{}
	postedData.Set("name", "John")
	postedData.Set("number", "42")

	form := New(postedData)
	form.MinLength("name", 5)
	if form.Valid() {
		t.Error("Field length reporting as too short when length is valid")
	}

	if form.Errors.Get("name") == "" {
		t.Error("Field should contain min length error message")
	}

	form.MinLength("number", 2)
	if !form.Valid() && !form.MinLength("number", 2) {
		t.Error("Form does not have required fields when it does")
	}

	if form.Errors.Get("number") != "" {
		t.Error("Field should not contain min length error message")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("email_address")

	if form.Valid() {
		t.Error("Form does not have required fields when it does")
	}

	postedData = url.Values{}
	postedData.Add("email", "mark@mark.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Form reported an invalid email when it was valid")
	}

	postedData = url.Values{}
	postedData.Add("email", "invalidemailaddress.com")

	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("Form reported email as valid when it was invalid")
	}

	//ther gmail trick
	postedData = url.Values{}
	postedData.Add("email", "fatherandson+cat@gmail.com")

	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Form reported email as valid when it was invalid")
	}

}
