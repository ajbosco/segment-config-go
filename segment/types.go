package segment

import (
	"fmt"
	"time"
)

// Workspace defines the struct for the workspace object
type Workspace struct {
	Name        string    `json:"name,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	ID          string    `json:"id,omitempty"`
	CreateTime  time.Time `json:"create_time,omitempty"`
}

// Sources defines the struct for the sources object
type Sources struct {
	Sources []Source `json:"sources,omitempty"`
}

// Source defines the struct for the source object
type Source struct {
	Name          string        `json:"name,omitempty"`
	CatalogName   string        `json:"catalog_name,omitempty"`
	Parent        string        `json:"parent,omitempty"`
	WriteKeys     []string      `json:"write_keys,omitempty"`
	LibraryConfig LibraryConfig `json:"library_config,omitempty"`
	CreateTime    time.Time     `json:"create_time,omitempty"`
}

// CommonEventSettings provides accepted values for CommonTrackEventOnViolations, CommonIdentifyEventOnViolations and CommonGroupEventOnViolations
type CommonEventSettings string

const (
	Allow CommonEventSettings = "ALLOW"

	// Only for use with CommonTrackEventOnViolations
	OmitProps CommonEventSettings = "OMIT_PROPERTIES"

	// Only for use with CommonIdentifyEventOnViolations and CommonIdentifyEventOnViolations
	OmitTraits CommonEventSettings = "OMIT_TRAITS"

	Block CommonEventSettings = "BLOCK"
)

type SourceConfig struct {
	Name                                string              `json:"name,omitempty"`
	Parent                              string              `json:"parent,omitempty"`
	AllowUnplannedTrackEvents           bool                `json:"allow_unplanned_track_events,omitempty"`
	AllowUnplannedIdentifyTraits        bool                `json:"allow_unplanned_identify_traits,omitempty"`
	AllowUnplannedGroupTraits           bool                `json:"allow_unplanned_group_traits,omitempty"`
	AllowTrackEventOnViolations         bool                `json:"allow_track_event_on_violations,omitempty"`
	AllowIdentifyTraitsOnViolations     bool                `json:"allow_identify_traits_on_violations,omitempty"`
	AllowGroupTraitsOnViolations        bool                `json:"allow_group_traits_on_violations,omitempty"`
	AllowUnplannedTrackEventsProperties bool                `json:"allow_unplanned_track_event_properties,omitempty"`
	AllowTrackPropertiesOnViolations    bool                `json:"allow_track_properties_on_violations,omitempty"`
	ForwardingBlockedEventsTo           string              `json:"forwarding_blocked_events_to,omitempty"`
	ForwardingViolationsTo              string              `json:"forwarding_violations_to,omitempty"`
	CommonTrackEventOnViolations        CommonEventSettings `json:"common_track_event_on_violations,omitempty"`
	CommonIdentifyEventOnViolations     CommonEventSettings `json:"common_identify_event_on_violations,omitempty"`
	CommonGroupEventOnViolations        CommonEventSettings `json:"common_group_event_on_violations,omitempty"`
}

type sourceConfigUpdateRequest struct {
	Config     SourceConfig `json:"schema_config,omitempty"`
	UpdateMask UpdateMask   `json:"update_mask,omitempty"`
}

// LibraryConfig contains information about a source's library
type LibraryConfig struct {
	MetricsEnabled       bool   `json:"metrics_enabled,omitempty"`
	RetryQueue           bool   `json:"retry_queue,omitempty"`
	CrossDomainIDEnabled bool   `json:"cross_domain_id_enabled,omitempty"`
	APIHost              string `json:"api_host,omitempty"`
}

// Destinations defines the struct for the destination object
type Destinations struct {
	Destinations []Destination `json:"destinations,omitempty"`
}

// Destination defines the struct for the destination object
type Destination struct {
	Name           string              `json:"name,omitempty"`
	Parent         string              `json:"parent,omitempty"`
	DisplayName    string              `json:"display_name,omitempty"`
	Enabled        bool                `json:"enabled,omitempty"`
	ConnectionMode string              `json:"connection_mode,omitempty"`
	Configs        []DestinationConfig `json:"config,omitempty"`
	CreateTime     time.Time           `json:"create_time,omitempty"`
	UpdateTime     time.Time           `json:"update_time,omitempty"`
}

// DestinationConfig contains information about how a Destination is configured
type DestinationConfig struct {
	Name        string      `json:"name,omitempty"`
	DisplayName string      `json:"display_name,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	Type        string      `json:"type,omitempty"`
}

// UpdateMask contains information for updating Destinations and Sources
type UpdateMask struct {
	Paths []string `json:"paths,omitempty"`
}

type sourceCreateRequest struct {
	Source Source `json:"source,omitempty"`
}

type destinationCreateRequest struct {
	Destination Destination `json:"destination,omitempty"`
}

type destinationUpdateRequest struct {
	Destination Destination `json:"destination,omitempty"`
	UpdateMask  UpdateMask  `json:"update_mask,omitempty"`
}

// TrackingPlans is a list of tracking plans
type TrackingPlans struct {
	TrackingPlans []TrackingPlan `json:"tracking_plans,omitempty"`
}

// TrackingPlan contains information about a tracking plan
type TrackingPlan struct {
	Name        string    `json:"name,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Rules       RuleSet   `json:"rules,omitempty"`
	CreateTime  time.Time `json:"create_time,omitempty"`
	UpdateTime  time.Time `json:"update_time,omitempty"`
}

// RuleSet contains a set of different rules about the tracking plan
type RuleSet struct {
	Global   Rules   `json:"global,omitempty"`
	Events   []Event `json:"events,omitempty"`
	Identify Rules   `json:"identify,omitempty"`
	Group    Rules   `json:"group,omitempty"`
}

// Rules contains information about a specific type of rules of the tracking plan
type Rules struct {
	Schema     string         `json:"$schema,omitempty"`
	Type       string         `json:"type,omitempty"`
	Properties RuleProperties `json:"properties,omitempty"`
	Required   []string       `json:"required,omitempty"`
}

// RuleProperties contains the different properties of a specific type of rules
type RuleProperties struct {
	Context    Properties `json:"context,omitempty"`
	Properties Properties `json:"properties,omitempty"`
	Traits     Properties `json:"traits,omitempty"`
}

// Properties contains information about a specific type of rule properties
type Properties struct {
	Properties map[string]Property `json:"properties,omitempty"`
	Required   []string            `json:"required,omitempty"`
	Type       string              `json:"type,omitempty"`
}

// Property contains information of a single property
type Property struct {
	Description string              `json:"description,omitempty"`
	Type        interface{}         `json:"type,omitempty"`
	Pattern     *string             `json:"pattern,omitempty"`
	Format      *string             `json:"format,omitempty"`
	Items       *Property           `json:"items,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
	Required    []string            `json:"required,omitempty"`
}

// Event contains information about a single event of the tracking plan
type Event struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Rules       Rules  `json:"rules,omitempty"`
	Version     *int   `json:"version,omitempty"`
}

type trackingPlanCreateRequest struct {
	TrackingPlan TrackingPlan `json:"tracking_plan,omitempty"`
}

type trackingPlanUpdateRequest struct {
	UpdateMask   UpdateMask   `json:"update_mask,omitempty"`
	TrackingPlan TrackingPlan `json:"tracking_plan,omitempty"`
}

type TrackingPlanSourceConnections struct {
	Connections []TrackingPlanSourceConnection `json:"connections,omitempty"`
}

// TrackingPlanSourceConnection represents the link between a tracking plan and a source
type TrackingPlanSourceConnection struct {
	// Source is the full name of the source (including the workspace path)
	Source         string `json:"source_name,omitempty"`
	TrackingPlanId string `json:"tracking_plan_id,omitempty"`
}

type trackingPlanSourceConnectionCreateRequest struct {
	Name string `json:"source_name"`
}

type SegmentApiError struct {
	Message string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (err *SegmentApiError) Error() string {
	return fmt.Sprintf("Segment Error %d: %s", err.Code, err.Message)
}
