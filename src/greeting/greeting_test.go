package greeting_test

import (
	. "greeting"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Hello_Func(t *testing.T) {
	expected := `{"txt::"hello world"}`
	r, _ := http.NewRequest(http.MethodGet, "/hello?name=world", nil)
	w := httptest.NewRecorder()

	Hello(w, r)

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	actual := string(body)

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
