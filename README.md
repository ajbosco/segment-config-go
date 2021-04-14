# segment-config-go

[![Github Actions](https://github.com/ajbosco/segment-config-go/workflows/build/badge.svg?branch=master&event=push)](https://github.com/ajbosco/segment-config-go/actions?workflow=build)
[![Go Report Card](https://goreportcard.com/badge/github.com/ajbosco/segment-config-go?style=flat-square)](https://goreportcard.com/report/github.com/ajbosco/segment-config-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/ajbosco/segment-config-go/segment)

segment-config-go is a Go client library for accessing the [Segment Config](https://segment.com/docs/config-api/) API.

This library allows you to do the following programmatically:

* List all your Segment sources and destinations
* Create [sources](https://segment.com/docs/sources/) 
* Create or modify [destinations](https://segment.com/docs/destinations/)
* Enable and disable destinations

## Authentication

segment-config-go requires a Segment Personal Access Token for authentication. You can generate one with the appropriate access by following the steps in the Segment [documentation](https://segment.com/docs/config-api/authentication/)

## Usage

```go
import "github.com/ajbosco/segment-config-go/segment"
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

