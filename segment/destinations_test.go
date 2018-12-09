package segment

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDestinations_ListDestinations(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"destinations": [
				{
					"name": "workspaces/myworkspace/sources/js/destinations/google-analytics",
					"parent": "workspaces/myworkspace/sources/js",
					"display_name": "Google Analytics",
					"enabled": true,
					"connection_mode": "CLOUD",
					"config": [
						{
							"name": "workspaces/myworkspace/sources/js/destinations/google-analytics/config/domain",
							"display_name": "Cookie Domain Name",
							"value": "",
							"type": "string"
						}
					]
				}
			],
			"next_page_token": ""
		}`)
	})

	actual, err := client.ListDestinations(testSource)
	assert.NoError(t, err)

	expected := Destinations{
		Destinations: []Destination{
			{
				Name:           "workspaces/myworkspace/sources/js/destinations/google-analytics",
				Parent:         "workspaces/myworkspace/sources/js",
				DisplayName:    "Google Analytics",
				Enabled:        true,
				ConnectionMode: "CLOUD",
				Configs: []DestinationConfig{
					{
						Name:        "workspaces/myworkspace/sources/js/destinations/google-analytics/config/domain",
						DisplayName: "Cookie Domain Name",
						Value:       "",
						Type:        "string"},
				}}}}
	assert.Equal(t, expected, actual)
}
func TestDestinations_GetDestination(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	testDest := "test-dest"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint, testDest)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"name": "workspaces/myworkspace/sources/js/destinations/google-analytics",
			"parent": "workspaces/myworkspace/sources/js",
			"display_name": "Google Analytics",
			"enabled": true,
			"connection_mode": "CLOUD",
			"config": [
				{
					"name": "workspaces/myworkspace/sources/js/destinations/google-analytics/config/domain",
					"display_name": "Cookie Domain Name",
					"value": "",
					"type": "string"
				}
			]
		}`)
	})

	actual, err := client.GetDestination(testSource, testDest)
	assert.NoError(t, err)

	expected := Destination{
		Name:           "workspaces/myworkspace/sources/js/destinations/google-analytics",
		Parent:         "workspaces/myworkspace/sources/js",
		DisplayName:    "Google Analytics",
		Enabled:        true,
		ConnectionMode: "CLOUD",
		Configs: []DestinationConfig{
			{
				Name:        "workspaces/myworkspace/sources/js/destinations/google-analytics/config/domain",
				DisplayName: "Cookie Domain Name",
				Value:       "",
				Type:        "string"},
		}}
	assert.Equal(t, expected, actual)
}

func TestDestinations_CreateDestination(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	testDest := Destination{
		Name:           "workspaces/myworkspace/sources/js/destinations/google-analytics",
		Parent:         "workspaces/myworkspace/sources/js",
		DisplayName:    "Google Analytics",
		Enabled:        true,
		ConnectionMode: "CLOUD"}

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"name": "workspaces/myworkspace/sources/js/destinations/google-analytics",
			"parent": "workspaces/myworkspace/sources/js",
			"display_name": "Google Analytics",
			"enabled": true,
			"connection_mode": "CLOUD"
		}`)
	})

	expected := Destination{
		Name:           "workspaces/myworkspace/sources/js/destinations/google-analytics",
		Parent:         "workspaces/myworkspace/sources/js",
		DisplayName:    "Google Analytics",
		Enabled:        true,
		ConnectionMode: "CLOUD"}

	actual, err := client.CreateDestination(testSource, testDest)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestDestinations_DeleteDestination(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	testDest := "test-dest"

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint, testDest)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
	})

	err := client.DeleteDestinaton(testSource, testDest)
	assert.NoError(t, err)
}

func TestDestinations_UpdateDestination(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	testDest := "test-dest"
	testDestUpdate := Destination{
		Name:           "workspaces/myworkspace/sources/js/destinations/google-analytics",
		Parent:         "workspaces/myworkspace/sources/js",
		DisplayName:    "Google Analytics",
		Enabled:        true,
		ConnectionMode: "CLOUD"}
	testUpdateMask := UpdateMask{
		Paths: []string{"destination.enabled"},
	}

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint, testDest)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"name": "workspaces/myworkspace/sources/js/destinations/google-analytics",
			"parent": "workspaces/myworkspace/sources/js",
			"display_name": "Google Analytics",
			"enabled": false,
			"connection_mode": "CLOUD"
		}`)
	})

	expected := Destination{
		Name:           "workspaces/myworkspace/sources/js/destinations/google-analytics",
		Parent:         "workspaces/myworkspace/sources/js",
		DisplayName:    "Google Analytics",
		Enabled:        false,
		ConnectionMode: "CLOUD"}

	actual, err := client.UpdateDestination(testSource, testDest, testDestUpdate, testUpdateMask)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}
