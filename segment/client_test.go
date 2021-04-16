package segment

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const (
	testToken     = "test-token"
	testWorkspace = "test-workspace"
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(testToken, testWorkspace)
	client.baseURL = server.URL
}

func teardown() {
	server.Close()
}

func testClientDefaultBaseURL(t *testing.T, c *Client) {
	if c.baseURL == "" || c.baseURL != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.baseURL, defaultBaseURL)
	}
}

func Test_NewClient(t *testing.T) {
	c := NewClient(testToken, testWorkspace)
	testClientDefaultBaseURL(t, c)
}

func Test_doRequest(t *testing.T) {
	setup()
	defer teardown()

	testData := `{"testing":"things"}`

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, testData)
	})

	actual, err := client.doRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	expected := []byte(testData)
	assert.Equal(t, expected, actual)
}

func Test_doRequest_httpError_notFound(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "Not Found", 404)
	})

	_, err := client.doRequest(http.MethodGet, "/", nil)
	assert.Error(t, err)
}

func Test_doRequest_httpError_badRequest(t *testing.T) {
	setup()
	defer teardown()

	expected := SegmentApiError{Code: 5, Message: "foo"}
	errorJson := `{ "error": "foo", "code": 5 }`

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, errorJson, 400)
	})

	_, err := client.doRequest(http.MethodGet, "/", nil)
	assert.EqualError(t, err, expected.Error())
}

func Test_doRequest_httpError_badRequestUnstructured(t *testing.T) {
	setup()
	defer teardown()

	expected := "the request is invalid"

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad request", 400)
	})

	_, err := client.doRequest(http.MethodGet, "/", nil)
	assert.EqualError(t, err, expected)
}
