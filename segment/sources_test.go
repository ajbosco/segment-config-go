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

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, _ *http.Request) {
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

func TestSources_UpdateSourceConfig(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/schema-config", apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, `{
			"name": "workspaces/%s/sources/%s/schema-config",
			"parent": "sources/js",
			"allow_unplanned_track_events": false,
			"allow_unplanned_identify_traits": true,
			"allow_unplanned_group_traits": false,
			"forwarding_blocked_events_to": "forwarding_blocked_events_to_source_slug",
			"allow_unplanned_track_event_properties": true,
			"allow_track_event_on_violations": false,
			"allow_identify_traits_on_violations": true,
			"allow_group_traits_on_violations": false,
			"forwarding_violations_to": "forwarding_violations_to_source_slug",
			"allow_track_properties_on_violations": false,
			"common_track_event_on_violations": "ALLOW",
			"common_identify_event_on_violations": "ALLOW",
			"common_group_event_on_violations": "ALLOW"
		}`, testWorkspace, testSource)
	})

	expected := SourceConfig{
		Name:                                "workspaces/" + testWorkspace + "/sources/" + testSource + "/schema-config",
		Parent:                              "sources/js",
		AllowUnplannedTrackEvents:           false,
		AllowUnplannedIdentifyTraits:        true,
		AllowUnplannedGroupTraits:           false,
		ForwardingBlockedEventsTo:           "forwarding_blocked_events_to_source_slug",
		AllowUnplannedTrackEventsProperties: true,
		AllowTrackEventOnViolations:         false,
		AllowIdentifyTraitsOnViolations:     true,
		AllowGroupTraitsOnViolations:        false,
		ForwardingViolationsTo:              "forwarding_violations_to_source_slug",
		AllowTrackPropertiesOnViolations:    false,
		CommonTrackEventOnViolations:        Allow,
		CommonIdentifyEventOnViolations:     Allow,
		CommonGroupEventOnViolations:        Allow,
	}

	returned, err := client.UpdateSourceConfig(testSource, expected)
	assert.NoError(t, err)
	assert.Equal(t, expected, returned)
}

func TestSources_GetSourceConfig(t *testing.T) {
	setup()
	defer teardown()

	testSource := "test-source"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/schema-config", apiVersion, WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSource)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, `{
			"name": "workspaces/%s/sources/%s/schema-config",
			"parent": "sources/js",
			"allow_unplanned_track_events": true,
			"allow_unplanned_identify_traits": true,
			"allow_unplanned_group_traits": true,
			"forwarding_blocked_events_to": "forwarding_blocked_events_to_source_slug",
			"allow_unplanned_track_event_properties": true,
			"allow_track_event_on_violations": true,
			"allow_identify_traits_on_violations": true,
			"allow_group_traits_on_violations": true,
			"forwarding_violations_to": "forwarding_violations_to_source_slug",
			"allow_track_properties_on_violations": true,
			"common_track_event_on_violations": "ALLOW",
			"common_identify_event_on_violations": "OMIT_TRAITS",
			"common_group_event_on_violations": "BLOCK"
		}`, testWorkspace, testSource)
	})

	expected := SourceConfig{
		Name:                                "workspaces/" + testWorkspace + "/sources/" + testSource + "/schema-config",
		Parent:                              "sources/js",
		AllowUnplannedTrackEvents:           true,
		AllowUnplannedIdentifyTraits:        true,
		AllowUnplannedGroupTraits:           true,
		ForwardingBlockedEventsTo:           "forwarding_blocked_events_to_source_slug",
		AllowUnplannedTrackEventsProperties: true,
		AllowTrackEventOnViolations:         true,
		AllowIdentifyTraitsOnViolations:     true,
		AllowGroupTraitsOnViolations:        true,
		ForwardingViolationsTo:              "forwarding_violations_to_source_slug",
		AllowTrackPropertiesOnViolations:    true,
		CommonTrackEventOnViolations:        Allow,
		CommonIdentifyEventOnViolations:     OmitTraits,
		CommonGroupEventOnViolations:        Block,
	}

	returned, err := client.GetSourceConfig(testSource)
	assert.NoError(t, err)
	assert.Equal(t, expected, returned)
}
