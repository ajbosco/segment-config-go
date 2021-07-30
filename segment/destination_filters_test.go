package segment

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleFilter1JSON = `
{
	"name": "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKl",
	"if": "type != 'track'",
	"actions": [
	  {
		"type": "drop_event"
	  }
	],
	"title": "Only allow track events",
	"description": "We don't need identify and page calls",
	"enabled": true
}
`

var sampleFilter1 = DestinationFilter{
	Name:        "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKl",
	Conditions:  "type != 'track'",
	Actions:     []DestinationFilterAction{NewDropEventAction()},
	Title:       "Only allow track events",
	Description: "We don't need identify and page calls",
	IsEnabled:   true,
}

const sampleFilter2JSON = `
{
  "name": "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
  "if": "!(event = \"Session Started\" or event = \"Order Completed\")",
  "actions": [
	{
        "type": "blacklist_fields",
		"fields": {
			"context": { "fields": ["foo"] },
			"properties": { "fields": ["bar"] },
			"traits": { "fields": ["baz"] }
		}
    }
  ],
  "title": "Only allow Session Started and Order Completed events",
  "description": "Those events are used for this and that",
  "enabled": false
}
`

var sampleFilter2 = DestinationFilter{
	Name:       "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
	Conditions: "!(event = \"Session Started\" or event = \"Order Completed\")",
	Actions: []DestinationFilterAction{NewBlockListEventAction(EventDescription{
		Context:    EventFieldsSelection{Fields: []string{"foo"}},
		Properties: EventFieldsSelection{Fields: []string{"bar"}},
		Traits:     EventFieldsSelection{Fields: []string{"baz"}},
	})},
	Title:       "Only allow Session Started and Order Completed events",
	Description: "Those events are used for this and that",
	IsEnabled:   false,
}

const sampleFilter3JSON = `
{
  "name": "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
  "if": "!(event = \"Session Started\" or event = \"Order Completed\")",
  "actions": [
	{
        "type": "whitelist_fields",
		"fields": {
			"context": { "fields": ["baz"] },
			"properties": { "fields": ["bar"] },
			"traits": { "fields": ["foo"] }
		}
    }
  ],
  "title": "Only allow Session Started and Order Completed events",
  "description": "Those events are used for this and that",
  "enabled": false
}
`

var sampleFilter3 = DestinationFilter{
	Name:       "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
	Conditions: "!(event = \"Session Started\" or event = \"Order Completed\")",
	Actions: []DestinationFilterAction{NewAllowListEventAction(EventDescription{
		Context:    EventFieldsSelection{Fields: []string{"baz"}},
		Properties: EventFieldsSelection{Fields: []string{"bar"}},
		Traits:     EventFieldsSelection{Fields: []string{"foo"}},
	})},
	Title:       "Only allow Session Started and Order Completed events",
	Description: "Those events are used for this and that",
	IsEnabled:   false,
}

const sampleFilter4JSON = `
{
  "name": "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
  "if": "!(event = \"Session Started\" or event = \"Order Completed\")",
  "actions": [
	{
        "type": "sample_event",
		"percent": 0.6,
		"path": "userId"
    }
  ],
  "title": "Only allow Session Started and Order Completed events",
  "description": "Those events are used for this and that",
  "enabled": true
}
`

var sampleFilter4 = DestinationFilter{
	Name:        "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKz",
	Conditions:  "!(event = \"Session Started\" or event = \"Order Completed\")",
	Actions:     []DestinationFilterAction{NewSamplingEventAction(0.6, "userId")},
	Title:       "Only allow Session Started and Order Completed events",
	Description: "Those events are used for this and that",
	IsEnabled:   true,
}

const sampleFilter5JSON = `
{
	"name": "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKl",
	"if": "type != 'track'",
	"actions": [
		{
			"type": "drop_event"
		},
		{
			"type": "sample_event",
			"percent": 0.6,
			"path": "userId"
		}
	],
	"title": "Only allow track events",
	"description": "We don't need identify and page calls",
	"enabled": true
}
`

var sampleFilter5 = DestinationFilter{
	Name:       "workspaces/test-workspace/sources/test-source/destinations/test-dest/filters/df_1JjE7f3gz8OXU6lVbUGuDhHcyKl",
	Conditions: "type != 'track'",
	Actions: []DestinationFilterAction{
		NewDropEventAction(),
		NewSamplingEventAction(0.6, "userId"),
	},
	Title:       "Only allow track events",
	Description: "We don't need identify and page calls",
	IsEnabled:   true,
}

var allFilters = []struct {
	filter     DestinationFilter
	filterJSON string
}{
	{filter: sampleFilter1, filterJSON: sampleFilter1JSON},
	{filter: sampleFilter2, filterJSON: sampleFilter2JSON},
	{filter: sampleFilter3, filterJSON: sampleFilter3JSON},
	{filter: sampleFilter4, filterJSON: sampleFilter4JSON},
	{filter: sampleFilter5, filterJSON: sampleFilter5JSON},
}

func jsonFilters() []string {
	filters := []string{}
	for _, f := range allFilters {
		filters = append(filters, f.filterJSON)
	}
	return filters
}

func inRequest(filterJSON string) string {
	return fmt.Sprintf(`
		{
			"filter": %s,
			"update_mask": {
				"paths": ["if", "actions", "title", "description", "enabled"]
			}
		}
	`, filterJSON)
}

func TestDestinationFilters_ListFilters(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	testDest := "test-dest"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s/%s",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint, testDest, DestinationFiltersEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"filters": [%s]
		}`, strings.Join(jsonFilters(), ","))
	})

	actual, err := client.ListDestinationFilters(testSource, testDest)
	assert.NoError(t, err)

	expected := []DestinationFilter{}
	for _, f := range allFilters {
		expected = append(expected, f.filter)
	}
	assert.Equal(t, expected, actual)
}

func withValidRequest(t *testing.T, method string, body string, handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if body != "" {
			raw, err := ioutil.ReadAll(r.Body)
			assert.NoError(t, err)
			defer r.Body.Close()
			assert.JSONEq(t, body, string(raw))
		}
		assert.Equal(t, method, r.Method)

		handler(w, r)
	}
}

func TestDestinationFilters_CreateFilter(t *testing.T) {
	testSource := "test-source"
	testDest := "test-dest"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s/%s",
		apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource, DestinationEndpoint, testDest, DestinationFiltersEndpoint)

	for _, testCase := range allFilters {
		setup()
		defer teardown()

		mux.HandleFunc(endpoint, withValidRequest(t, "POST", inRequest(testCase.filterJSON), func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testCase.filterJSON)
		}))

		t.Run(fmt.Sprintf("CreateFilter for %s", testCase.filterJSON), func(t *testing.T) {
			expected := testCase.filter
			actual, err := client.CreateDestinationFilter(testSource, testDest, testCase.filter)
			assert.NoError(t, err)

			assert.EqualValues(t, expected, *actual)
		})
	}
}

func TestDestinationFilters_UpdateFilter(t *testing.T) {
	testSource := "test-source"
	testDest := "test-dest"

	for _, testCase := range allFilters {
		setup()
		defer teardown()
		endpoint := fmt.Sprintf("/%s/%s", apiVersion, testCase.filter.Name)

		mux.HandleFunc(endpoint, withValidRequest(t, "PATCH", inRequest(testCase.filterJSON), func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testCase.filterJSON)
		}))

		t.Run(fmt.Sprintf("UpdateFilter for %s", testCase.filterJSON), func(t *testing.T) {
			expected := testCase.filter
			actual, err := client.UpdateDestinationFilter(testSource, testDest, testCase.filter)
			assert.NoError(t, err)

			assert.EqualValues(t, expected, *actual)
		})
	}
}

func TestDestinationFilters_GetFilter(t *testing.T) {
	testSource := "test-source"
	testDest := "test-dest"

	for _, testCase := range allFilters {
		setup()
		defer teardown()
		endpoint := fmt.Sprintf("/%s/%s", apiVersion, testCase.filter.Name)

		mux.HandleFunc(endpoint, withValidRequest(t, "GET", "", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testCase.filterJSON)
		}))

		t.Run(fmt.Sprintf("GetFilter for %s", testCase.filterJSON), func(t *testing.T) {
			expected := testCase.filter
			fmt.Println(testCase.filter.Name)
			fmt.Println(pathToId(testCase.filter.Name))
			actual, err := client.GetDestinationFilter(testSource, testDest, pathToId(testCase.filter.Name))
			assert.NoError(t, err)
			assert.EqualValues(t, expected, *actual)
		})
	}
}

func TestDestinationFilters_DeleteFilter(t *testing.T) {
	testSource := "test-source"
	testDest := "test-dest"

	for _, testCase := range allFilters {
		setup()
		defer teardown()
		endpoint := fmt.Sprintf("/%s/%s", apiVersion, testCase.filter.Name)

		mux.HandleFunc(endpoint, withValidRequest(t, "DELETE", "", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{}`)
		}))

		t.Run(fmt.Sprintf("DeleteFilter for %s", testCase.filterJSON), func(t *testing.T) {
			fmt.Println(testCase.filter.Name)
			fmt.Println(pathToId(testCase.filter.Name))
			err := client.DeleteDestinationFilter(testSource, testDest, pathToId(testCase.filter.Name))
			assert.NoError(t, err)
		})
	}
}

func pathToId(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}

	return path
}
