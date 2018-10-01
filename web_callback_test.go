package instru

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestWebCallback(t *testing.T) {
	expectedBody := []byte(`{"metrics":{"test01":{"_counter":{"total":1,"events":{"success":1}}}}}`)
	instr := NewInstrumentation()
	instr.Count("test01").Event("success")

	var gotBody []byte
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotBody, _ = ioutil.ReadAll(r.Body)
	}))
	defer ts.Close()

	callback := NewWebCallback(ts.URL)
	err := callback.OnCallback(instr)

	FatalIfError(t, err)
	FatalIf(t, bytes.Compare(gotBody, expectedBody) != 0, "got wrong body")
}

func TestWebCallback_ClientError(t *testing.T) {
	callback := NewWebCallback("wrong-url")
	err := callback.OnCallback(NewInstrumentation())

	FatalIfWrongError(t, err, `Post wrong-url: unsupported protocol scheme ""`)
}

func TestWebCallback_StatusCodeNotOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	callback := NewWebCallback(ts.URL)
	err := callback.OnCallback(NewInstrumentation())
	FatalIfWrongError(t, err, "Got status code 500")
}
