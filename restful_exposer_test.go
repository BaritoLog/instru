package instru

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestRestfulServer_Start(t *testing.T) {
	expectedBody := []byte(`{"metrics":{"test01":{"_counter":{"total":1,"events":{"success":1}}}}}`)

	instr := NewInstrumentation()
	instr.Count("test01").Event("success")

	exposer := NewRestfulExposer(":65500")

	go exposer.Expose(instr)
	defer exposer.Stop()

	resp, err := http.Get("http://localhost:65500")

	FatalIfError(t, err)
	FatalIf(t, resp.StatusCode != 200, "wrong ")

	body, _ := ioutil.ReadAll(resp.Body)
	FatalIf(t, bytes.Compare(body, expectedBody) != 0, "got wrong body")
}
