package segment

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testTrackingPlanID            = "rs_123abc"
	testTrackingPlanErrorResponse = `{"name": 123}`
	testTrackingPlanResponse      = `{
		"name": "workspaces/test/tracking-plans/rs_123abc",
		"display_name": "Test Tracking Plan",
		"rules": {
		  "global": {
			"$schema": "http://json-schema.org/draft-04/schema#",
			"type": "object",
			"properties": {
			  "context": {
				"type": "object",
				"properties": {
				  "context_prop_1": {
					"type": [
					  "object"
					]
				  }
				},
				"required": [
				  "context_prop_1"
				]
			  },
			  "properties": {},
			  "traits": {}
			}
		  },
		  "events": [
			{
			  "name": "Test Event Clicked",
			  "description": "A simple test event",
			  "rules": {
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": {
				  "traits": {},
				  "properties": {
					"required": [
					  "user_id",
					  "email"
					],
					"type": "object",
					"properties": {
					  "user_id": {
						"description": "unique id of the user",
						"type": [
						  "string"
						]
					  },
					  "email": {
						"description": "user email",
						"type": [
						  "string"
						]
					  },
					  "test_prop": {
						"description": "test prop",
						"type": [
						  "integer"
						]
					  }
					},
					"context": {}
				  }
				},
				"version": 1
			  }
			}
		  ]
		}
	  }`
)

func TestTrackingPlans_ListTrackingPlans(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"tracking_plans": [
				{
			  		"name": "workspaces/test/tracking-plans/rs_123abc",
					"display_name": "Test Tracking Plan 1",
					"rules": {
						"events": []
					}
			 	},
				{
					"name": "workspaces/test/tracking-plans/rs_456def",
					"display_name": "Test Tracking Plan 2",
					"rules": {
					 	"events": []
					}
				}
			]
		}`)
	})

	actual, err := client.ListTrackingPlans()
	assert.NoError(t, err)

	expected := TrackingPlans{
		TrackingPlans: []TrackingPlan{
			{
				Name:        "workspaces/test/tracking-plans/rs_123abc",
				DisplayName: "Test Tracking Plan 1",
				Rules: RuleSet{
					Events: []Event{},
				},
			},
			{
				Name:        "workspaces/test/tracking-plans/rs_456def",
				DisplayName: "Test Tracking Plan 2",
				Rules: RuleSet{
					Events: []Event{},
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestTrackingPlans_GetTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testTrackingPlanResponse)
	})

	actual, err := client.GetTrackingPlan(testTrackingPlanID)
	assert.NoError(t, err)

	version := 1
	expected := TrackingPlan{
		Name:        "workspaces/test/tracking-plans/rs_123abc",
		DisplayName: "Test Tracking Plan",
		Rules: RuleSet{
			Global: Rules{
				Schema: "http://json-schema.org/draft-04/schema#",
				Type:   "object",
				Properties: RuleProperties{
					Context: Properties{
						Type: "object",
						Properties: map[string]Property{
							"context_prop_1": {
								Type: []string{"object"},
							},
						},
						Required: []string{"context_prop_1"},
					},
					Properties: Properties{},
					Traits:     Properties{},
				},
			},
			Events: []Event{
				{
					Name:        "Test Event Clicked",
					Description: "A simple test event",
					Rules: Rules{
						Schema: "http://json-schema.org/draft-07/schema#",
						Type:   "object",
						Properties: RuleProperties{
							Traits: Properties{},
							Properties: Properties{
								Required: []string{"user_id", "email"},
								Type:     "object",
								Properties: map[string]Property{
									"user_id": {
										Description: "unique id of the user",
										Type:        []string{"string"},
									},
									"email": {
										Description: "user email",
										Type:        []string{"string"},
									},
									"test_prop": {
										Description: "test prop",
										Type:        []string{"integer"},
									},
								},
							},
							Context: Properties{},
						},
						Version: &version,
					},
				},
			},
		},
	}
	assert.Equal(t, expected, actual)
}

func TestTrackingPlans_CreateTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testTrackingPlanResponse)
	})

	version := 1
	expected := TrackingPlan{
		DisplayName: "Test Tracking Plan",
		Rules: RuleSet{
			Global: Rules{
				Schema: "http://json-schema.org/draft-04/schema#",
				Type:   "object",
				Properties: RuleProperties{
					Context: Properties{
						Type: "object",
						Properties: map[string]Property{
							"context_prop_1": {
								Type: []string{"object"},
							},
						},
						Required: []string{"context_prop_1"},
					},
					Properties: Properties{},
					Traits:     Properties{},
				},
			},
			Events: []Event{
				{
					Name:        "Test Event Clicked",
					Description: "A simple test event",
					Rules: Rules{
						Schema: "http://json-schema.org/draft-07/schema#",
						Type:   "object",
						Properties: RuleProperties{
							Traits: Properties{},
							Properties: Properties{
								Required: []string{"user_id", "email"},
								Type:     "object",
								Properties: map[string]Property{
									"user_id": {
										Description: "unique id of the user",
										Type:        []string{"string"},
									},
									"email": {
										Description: "user email",
										Type:        []string{"string"},
									},
									"test_prop": {
										Description: "test prop",
										Type:        []string{"integer"},
									},
								},
							},
						},
						Version: &version,
					},
				},
			},
		},
	}

	actual, err := client.CreateTrackingPlan(expected)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual.Name)
	assert.Equal(t, expected.DisplayName, actual.DisplayName)
	assert.Equal(t, expected.Rules, actual.Rules)
}

func TestTrackingPlans_UpdateTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testTrackingPlanResponse)
	})

	version := 1
	expected := TrackingPlan{
		DisplayName: "Test Tracking Plan",
		Rules: RuleSet{
			Global: Rules{
				Schema: "http://json-schema.org/draft-04/schema#",
				Type:   "object",
				Properties: RuleProperties{
					Context: Properties{
						Type: "object",
						Properties: map[string]Property{
							"context_prop_1": {
								Type: []string{"object"},
							},
						},
						Required: []string{"context_prop_1"},
					},
					Properties: Properties{},
					Traits:     Properties{},
				},
			},
			Events: []Event{
				{
					Name:        "Test Event Clicked",
					Description: "A simple test event",
					Rules: Rules{
						Schema: "http://json-schema.org/draft-07/schema#",
						Type:   "object",
						Properties: RuleProperties{
							Traits: Properties{},
							Properties: Properties{
								Required: []string{"user_id", "email"},
								Type:     "object",
								Properties: map[string]Property{
									"user_id": {
										Description: "unique id of the user",
										Type:        []string{"string"},
									},
									"email": {
										Description: "user email",
										Type:        []string{"string"},
									},
									"test_prop": {
										Description: "test prop",
										Type:        []string{"integer"},
									},
								},
							},
						},
						Version: &version,
					},
				},
			},
		},
	}

	actual, err := client.UpdateTrackingPlan(testTrackingPlanID, expected)
	assert.NoError(t, err)
	assert.Equal(t, "workspaces/test/tracking-plans/rs_123abc", actual.Name)
	assert.Equal(t, expected.DisplayName, actual.DisplayName)
	assert.Equal(t, expected.Rules, actual.Rules)
}

func TestTrackingPlans_DeleteTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{}")
	})

	err := client.DeleteTrackingPlan(testTrackingPlanID)
	assert.NoError(t, err)
}