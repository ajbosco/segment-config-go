package segment

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspaces_GetWorkspace(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"name": "workspaces/myworkspace",
			"display_name": "My Workspace",
			"id": "jwt9cirmwq"
		  }`)
	})

	actual, err := client.GetWorkspace()
	assert.NoError(t, err)

	expected := Workspace{
		Name:        "workspaces/myworkspace",
		DisplayName: "My Workspace",
		ID:          "jwt9cirmwq"}
	assert.Equal(t, expected, actual)
}
