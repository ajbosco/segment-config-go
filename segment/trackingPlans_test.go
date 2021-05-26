package segment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
					"type": "object"
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
			  "version": 1,
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
						"type": "string"
					  },
					  "created": {
						  "description": "a datetime string",
						  "type": "string",
						  "format": "date-time"
					  },
					  "email": {
						"description": "user email",
						"pattern": "@",
						"type": "string"
					  },
					  "test_prop": {
						"description": "test prop",
						"type": "integer"
					  },
					  "enum_prop": {
						"type": "string",
						"enum": ["foo", "bar", null]
					  },
					  "array_prop": {
						"type": "array",
						"description": "array prop",
						"items": {
						  "type": "string",
						  "description": ""
						}
					  },
					  "prop_obj": {
						"description": "object prop",
						"type": "object",
						"properties": {
						  "prop_str": {
							"type": ["string", "null"],
							"description": ""
						  },
						  "prop_obj_nested": {
							"type": "object",
							"description": "nested object prop",
							"properties": {
							  "prop_bool": {
								"type": "boolean",
								"description": ""
							  }
							},
							"required": []
						  }
						},
						"required": [
						  "prop_str",
						  "prop_obj_nested"
						]
					  }
					},
					"context": {}
				  }
				}
			  }
			}
		  ]
		}
	  }`
	testSourcesResponse = `{
		"connections": [
			{
				"source_name": "workspaces/myworkspace/sources/test_source1",
				"tracking_plan_id": "rs_123abc"
			},
			{
				"source_name": "workspaces/myworkspace/sources/test_source2",
				"tracking_plan_id": "rs_123abc"
			}
		]
	}`
)

func newString(value string) *string {
	return &value
}

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
	atsign := "@"
	datetimeFormat := "date-time"
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
								Type: "object",
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
					Version:     &version,
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
										Type:        "string",
									},
									"created": {
										Description: "a datetime string",
										Format:      &datetimeFormat,
										Type:        "string",
									},
									"email": {
										Description: "user email",
										Pattern:     &atsign,
										Type:        "string",
									},
									"test_prop": {
										Description: "test prop",
										Type:        "integer",
									},
									"enum_prop": {
										Type: "string",
										Enum: []*string{newString("foo"), newString("bar"), nil},
									},
									"array_prop": {
										Description: "array prop",
										Type:        "array",
										Items: &Property{
											Description: "",
											Type:        "string",
										},
									},
									"prop_obj": {
										Description: "object prop",
										Required:    []string{"prop_str", "prop_obj_nested"},
										Type:        "object",
										Properties: map[string]Property{
											"prop_str": {
												Type:        []interface{}{"string", "null"},
												Description: "",
											},
											"prop_obj_nested": {
												Type:        "object",
												Description: "nested object prop",
												Properties: map[string]Property{
													"prop_bool": {
														Type:        "boolean",
														Description: "",
													},
												},
												Required: []string{},
											},
										},
									},
								},
							},
							Context: Properties{},
						},
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
	atsign := "@"
	datetimeFormat := "date-time"
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
								Type: "object",
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
					Version:     &version,
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
										Type:        "string",
									},
									"created": {
										Description: "a datetime string",
										Format:      &datetimeFormat,
										Type:        "string",
									},
									"email": {
										Description: "user email",
										Pattern:     &atsign,
										Type:        "string",
									},
									"test_prop": {
										Description: "test prop",
										Type:        "integer",
									},
									"enum_prop": {
										Type: "string",
										Enum: []*string{newString("foo"), newString("bar"), nil},
									},
									"array_prop": {
										Description: "array prop",
										Type:        "array",
										Items: &Property{
											Description: "",
											Type:        "string",
										},
									},
									"prop_obj": {
										Description: "object prop",
										Required:    []string{"prop_str", "prop_obj_nested"},
										Type:        "object",
										Properties: map[string]Property{
											"prop_str": {
												Type:        []interface{}{"string", "null"},
												Description: "",
											},
											"prop_obj_nested": {
												Type:        "object",
												Description: "nested object prop",
												Properties: map[string]Property{
													"prop_bool": {
														Type:        "boolean",
														Description: "",
													},
												},
												Required: []string{},
											},
										},
									},
								},
							},
						},
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
	atsign := "@"
	datetimeFormat := "date-time"
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
								Type: "object",
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
					Version:     &version,
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
										Type:        "string",
									},
									"created": {
										Description: "a datetime string",
										Format:      &datetimeFormat,
										Type:        "string",
									},
									"email": {
										Description: "user email",
										Pattern:     &atsign,
										Type:        "string",
									},
									"test_prop": {
										Description: "test prop",
										Type:        "integer",
									},
									"enum_prop": {
										Type: "string",
										Enum: []*string{newString("foo"), newString("bar"), nil},
									},
									"array_prop": {
										Description: "array prop",
										Type:        "array",
										Items: &Property{
											Description: "",
											Type:        "string",
										},
									},
									"prop_obj": {
										Description: "object prop",
										Required:    []string{"prop_str", "prop_obj_nested"},
										Type:        "object",
										Properties: map[string]Property{
											"prop_str": {
												Type:        []interface{}{"string", "null"},
												Description: "",
											},
											"prop_obj_nested": {
												Type:        "object",
												Description: "nested object prop",
												Properties: map[string]Property{
													"prop_bool": {
														Type:        "boolean",
														Description: "",
													},
												},
												Required: []string{},
											},
										},
									},
								},
							},
						},
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
func TestTrackingPlans_ListTrackingPlansSourceConnections(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/source-connections", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint, testTrackingPlanID)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprint(w, testSourcesResponse)
		}
	})

	actual, err := client.ListTrackingPlanSources(testTrackingPlanID)
	assert.NoError(t, err)

	expected := []TrackingPlanSourceConnection{
		{Source: "workspaces/myworkspace/sources/test_source1", TrackingPlanId: testTrackingPlanID},
		{Source: "workspaces/myworkspace/sources/test_source2", TrackingPlanId: testTrackingPlanID},
	}

	assert.Equal(t, expected, actual)
}

func TestTrackingPlans_CreateTrackingPlansSourceConnection(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/source-connections", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint, testTrackingPlanID)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		var j bytes.Buffer
		b, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.NoError(t, json.Compact(&j, b))
		if r.Method == http.MethodPost && j.String() == fmt.Sprintf(`{"source_name":"workspaces/%s/sources/test_source1"}`, testWorkspace) {
			fmt.Fprintf(w, `{
				"source_name": "test_source1",
				"tracking_plan_id": "%s"
			}`, testTrackingPlanID)
		}
	})

	err := client.CreateTrackingPlanSourceConnection(testTrackingPlanID, "test_source1")
	assert.NoError(t, err)
}

func TestTrackingPlans_DeleteTrackingPlansSourceConnection(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/source-connections/%s", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint, testTrackingPlanID, "test_source1")

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			fmt.Fprint(w, "{}")
		}
	})

	err := client.DeleteTrackingPlanSourceConnection(testTrackingPlanID, "test_source1")
	assert.NoError(t, err)
}
