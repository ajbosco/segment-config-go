# segment-config-go

[![Github Actions](https://github.com/uswitch/segment-config-go/workflows/build/badge.svg?branch=master&event=push)](https://github.com/uswitch/segment-config-go/actions?workflow=build)
[![Go Report Card](https://goreportcard.com/badge/github.com/uswitch/segment-config-go?style=flat-square)](https://goreportcard.com/report/github.com/uswitch/segment-config-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/uswitch/segment-config-go/segment)

segment-config-go is a Go client library for accessing the [Segment Config](https://segment.com/docs/config-api/) API.

This library allows you to do the following programmatically:

* List all your Segment sources and destinations
* Create [sources](https://segment.com/docs/sources/)
* Create or modify [destinations](https://segment.com/docs/destinations/)
* Enable and disable destinations
* Create, list or modify [tracking plans](https://segment.com/docs/protocols/tracking-plan/create/)

## Authentication

segment-config-go requires a Segment Personal Access Token for authentication. You can generate one with the appropriate access by following the steps in the Segment [documentation](https://segment.com/docs/config-api/authentication/)

## Usage

```go
import "github.com/uswitch/segment-config-go/segment"
```

Construct a new Segment client with your access token and Segment workspace. For example:

```go
accessToken :=  os.Getenv("ACCESS_TOKEN")
segmentWorkspace :=  os.Getenv("SEGMENT_WORKSPACE")

client := segment.NewClient(accessToken, segmentWorkspace)
```

Now you can interact with the API to do things like list all [sources](https://segment.com/docs/sources/) in your workspace:

```go
sources, err := c.ListSources()
```

List [destinations](https://segment.com/docs/destinations/) for a given source:

```go
destinations, err := c.ListDestinations("your-source")
```

Create a new [source](https://segment.com/docs/sources/):

```go
source, err := c.CreateSource("your-source", "catalog/sources/javascript")
```

Create a new [destination](https://segment.com/docs/destinations/):

```go
source, err := c.CreateDestination("your-source", "google-analytics", "cloud", false, nil)
```

Create a new [tracking plan](https://segment.com/docs/protocols/tracking-plan/create/):

```go
tp := TrackingPlan{
    DisplayName: "Your Tracking Plan",
    Rules: RuleSet{
        Global: Rules{
            Schema: "http://json-schema.org/draft-04/schema#",
            Type:   "object",
            Properties: RuleProperties{
                Context: Properties{
                    Type: "object",
                    Properties: map[string]Property{},
                },
                Properties: Properties{},
                Traits:     Properties{},
            },
        },
        Events: []Event{
            {
                Name:        "Test Event",
                Description: "A simple test event",
                Rules: Rules{
                    Schema: "http://json-schema.org/draft-07/schema#",
                    Type:   "object",
                    Properties: RuleProperties{
                        Traits: Properties{},
                        Properties: Properties{
                            Required: []string{"user_id"},
                            Type:     "object",
                            Properties: map[string]Property{
                                "user_id": {
                                    Description: "unique id of the user",
                                    Type:        []string{"string"},
                                },
                            },
                        },
                    },
                },
            },
        },
    },
}
trackingPlan, err := c.CreateTrackingPlan(tp)
```

Get an existing [tracking plan](https://segment.com/docs/protocols/tracking-plan/create/):

```go
trackingPlan, err := c.GetTrackingPlan("rs_123abc")
```

List all [tracking plans](https://segment.com/docs/protocols/tracking-plan/create/):

```go
trackingPlans, err := c.ListTrackingPlans()
```

Update an existing [tracking plan](https://segment.com/docs/protocols/tracking-plan/create/):
```go
tp := TrackingPlan{
    DisplayName: "Your Tracking Plan",
    Rules: RuleSet{
        Global: Rules{
            Schema: "http://json-schema.org/draft-04/schema#",
            Type:   "object",
            Properties: RuleProperties{
                Context: Properties{
                    Type: "object",
                    Properties: map[string]Property{},
                },
                Properties: Properties{},
                Traits:     Properties{},
            },
        },
        Events: []Event{
            {
                Name:        "Test Event",
                Description: "A simple test event",
                Rules: Rules{
                    Schema: "http://json-schema.org/draft-07/schema#",
                    Type:   "object",
                    Properties: RuleProperties{
                        Traits: Properties{},
                        Properties: Properties{
                            Required: []string{"user_id"},
                            Type:     "object",
                            Properties: map[string]Property{
                                "user_id": {
                                    Description: "unique id of the user",
                                    Type:        []string{"string"},
                                },
                            },
                        },
                    },
                },
            },
        },
    },
}
trackingPlan, err := c.UpdateTrackingPlan("rs_123abc", tp)
```

Delete an existing [tracking plan](https://segment.com/docs/protocols/tracking-plan/create/):

```go
err := client.DeleteTrackingPlan("rs_123abc")
```
