package segment

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSources_ListSources(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"sources": [
			  {
				"name": "workspaces/myworkspace/sources/ios",
				"parent": "workspaces/myworkspace",
				"catalog_name": "catalog/sources/ios",
				"write_keys": [
				  "F4xIkfwZiNBeYI0A9y1Cikhgi9dy7gX7"
				],
				"library_config": {
				  "metrics_enabled": false,
				  "retry_queue": false,
				  "cross_domain_id_enabled": false,
				  "api_host": ""
				}
			  },
			  {
				"name": "workspaces/myworkspace/sources/js",
				"parent": "workspaces/myworkspace",
				"catalog_name": "catalog/sources/javascript",
				"write_keys": [
				  "fbSIEioXNHc61VQgnmAHsHoDISjD6an7"
				],
				"library_config": {
				  "metrics_enabled": false,
				  "retry_queue": false,
				  "cross_domain_id_enabled": false,
				  "api_host": ""
				}
			  }
			],
			"next_page_token": ""
		  }`)
	})

	actual, err := client.ListSources()
	assert.NoError(t, err)

	expected := Sources{Sources: []Source{
		{
			Name:        "workspaces/myworkspace/sources/ios",
			CatalogName: "catalog/sources/ios",
			Parent:      "workspaces/myworkspace",
			WriteKeys:   []string{"F4xIkfwZiNBeYI0A9y1Cikhgi9dy7gX7"},
			LibraryConfig: LibraryConfig{
				MetricsEnabled:       false,
				RetryQueue:           false,
				CrossDomainIDEnabled: false,
				APIHost:              ""}},
		{
			Name:        "workspaces/myworkspace/sources/js",
			CatalogName: "catalog/sources/javascript",
			Parent:      "workspaces/myworkspace",
			WriteKeys:   []string{"fbSIEioXNHc61VQgnmAHsHoDISjD6an7"},
			LibraryConfig: LibraryConfig{
				MetricsEnabled:       false,
				RetryQueue:           false,
				CrossDomainIDEnabled: false,
				APIHost:              ""}}}}
	assert.Equal(t, expected, actual)
}

func TestSources_GetSource(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"name": "workspaces/myworkspace/sources/js",
			"parent": "workspaces/myworkspace",
			"catalog_name": "catalog/sources/ios",
			"write_keys": [
			  "F4xIkfwZiNBeYI0A9y1Cikhgi9dy7gX7"
			],
			"library_config": {
			  "metrics_enabled": false,
			  "retry_queue": false,
			  "cross_domain_id_enabled": false,
			  "api_host": ""
			}
		  }`)
	})

	actual, err := client.GetSource(testSource)
	assert.NoError(t, err)

	expected := Source{
		Name:        "workspaces/myworkspace/sources/js",
		CatalogName: "catalog/sources/ios",
		Parent:      "workspaces/myworkspace",
		WriteKeys:   []string{"F4xIkfwZiNBeYI0A9y1Cikhgi9dy7gX7"},
		LibraryConfig: LibraryConfig{
			MetricsEnabled:       false,
			RetryQueue:           false,
			CrossDomainIDEnabled: false,
			APIHost:              ""}}
	assert.Equal(t, expected, actual)
}

func TestSources_CreateSource(t *testing.T) {
	setup()
	defer teardown()

	testSrcName := "your-source"
	testCatName := "catalog/sources/javascript"

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"name": "workspaces/myworkspace/sources/js",
			"parent": "workspaces/myworkspace",
			"catalog_name": "catalog/sources/ios",
			"write_keys": [
			  "F4xIkfwZiNBeYI0A9y1Cikhgi9dy7gX7"
			],
			"library_config": {
			  "metrics_enabled": false,
			  "retry_queue": false,
			  "cross_domain_id_enabled": false,
			  "api_host": ""
			}
		  }`)
	})

	expected := Source{
		Name:        "workspaces/myworkspace/sources/js",
		CatalogName: "catalog/sources/ios",
		Parent:      "workspaces/myworkspace",
		WriteKeys:   []string{"F4xIkfwZiNBeYI0A9y1Cikhgi9dy7gX7"},
		LibraryConfig: LibraryConfig{
			MetricsEnabled:       false,
			RetryQueue:           false,
			CrossDomainIDEnabled: false,
			APIHost:              ""}}

	actual, err := client.CreateSource(testSrcName, testCatName)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSources_DeleteSource(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
	})

	err := client.DeleteSource(testSource)
	assert.NoError(t, err)
}
