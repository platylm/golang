package greeting_test

import (
	. "greeting"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Hello_Get_Func(t *testing.T) {
	expected := `{"txt":"hello world"}`
	r, _ := http.NewRequest(http.MethodGet, "/hello?name=world", nil)
	w := httptest.NewRecorder()

	Hello(w, r)

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func Test_Hello_Post_Func(t *testing.T) {
	expected := `{"txt":"hello world"}`
	msg := `{"name":"world"}`
	r, _ := http.NewRequest(http.MethodPost, "/hello-post", strings.NewReader(msg))
	w := httptest.NewRecorder()

	HelloPost(w, r)

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func Test_Greeting_Message(t *testing.T) {
	expected := "hello world"

	actual := GreetingMessage("world")

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
