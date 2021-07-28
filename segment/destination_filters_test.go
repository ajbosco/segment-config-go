package segment

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleFilter1JSON = `
{
	"name": "workspaces/myworkspace/sources/test-source/destinations/google-analytics/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKl",
	"if": "type != 'track'",
	"actions": [
	  {
		"type": "drop_event",
		"fields": {}
	  }
	],
	"title": "Only allow track events",
	"description": "We don't need identify and page calls",
	"enabled": true
}
`

const sampleFilter2JSON = `
{
  "name": "workspaces/myworkspace/sources/test-source/destinations/google-analytics/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
  "if": "!(event = \"Session Started\" or event = \"Order Completed\")",
  "actions": [
	{
	  "type": "drop_event",
	  "fields": {}
	}
  ],
  "title": "Only allow Session Started and Order Completed events",
  "description": "",
  "enabled": false
}
`

func TestDestinationFilters_ListFilters(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	testDest := "test-dest"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s/%s",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint, testDest, DestinationFiltersEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"filters": [%s, %s]
		}`, sampleFilter1JSON, sampleFilter2JSON)
	})

	actual, err := client.ListDestinationFilters(testSource, testDest)
	assert.NoError(t, err)

	expected := []DestinationFilter{
		{
			Name:        "workspaces/myworkspace/sources/test-source/destinations/google-analytics/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKl",
			Conditions:  "type != 'track'",
			Actions:     []DestinationFilterAction{NewDropEventAction()},
			Title:       "Only allow track events",
			Description: "We don't need identify and page calls",
			IsEnabled:   true,
		},
		{
			Name:        "workspaces/myworkspace/sources/test-source/destinations/google-analytics/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
			Conditions:  "!(event = \"Session Started\" or event = \"Order Completed\")",
			Actions:     []DestinationFilterAction{NewDropEventAction()},
			Title:       "Only allow Session Started and Order Completed events",
			Description: "",
			IsEnabled:   false,
		},
	}
	assert.Equal(t, expected, actual)
}
