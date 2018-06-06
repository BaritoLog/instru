package instru

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestRestfulServer_Start(t *testing.T) {
	expectedBody := []byte(`{"metrics":{"test01":{"_counter":{"Events":{"success":1},"Total":1}}}}`)

	instr := &instrumentation{}
	err := json.Unmarshal(expectedBody, instr)
	FatalIfError(t, err)

	exposer := NewRestfulExposer(":65500")

	go exposer.Expose(instr)
	defer exposer.Stop()

	resp, err := http.Get("http://localhost:65500")

	FatalIfError(t, err)
	FatalIf(t, resp.StatusCode != 200, "wrong ")

	body, _ := ioutil.ReadAll(resp.Body)
	FatalIf(t, bytes.Compare(body, expectedBody) != 0, "got wrong body")
}
