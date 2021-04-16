package segment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ListSources returns all sources for a workspace
func (c *Client) ListSources() (Sources, error) {
	var s Sources
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s", WorkspacesEndpoint, c.workspace, SourceEndpoint),
		nil)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return s, errors.Wrap(err, "failed to unmarshal sources response")
	}

	return s, nil
}

// GetSource returns information about a source
func (c *Client) GetSource(srcName string) (Source, error) {
	var s Source
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s",
			WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName),
		nil)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return s, errors.Wrap(err, "failed to unmarshal source response")
	}

	return s, nil
}

// CreateSource creates a new source
func (c *Client) CreateSource(srcName string, catName string) (Source, error) {
	var s Source
	srcFullName := fmt.Sprintf("%s/%s/%s/%s",
		WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName)
	src := Source{
		Name:        srcFullName,
		CatalogName: catName,
	}
	req := sourceCreateRequest{src}
	data, err := c.doRequest(http.MethodPost,
		fmt.Sprintf("%s/%s/%s",
			WorkspacesEndpoint, c.workspace, SourceEndpoint),
		req)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return s, errors.Wrap(err, "failed to unmarshal source response")
	}

	return s, nil
}

// DeleteSource deletes a source from the workspace
func (c *Client) DeleteSource(srcName string) error {
	_, err := c.doRequest(http.MethodDelete,
		fmt.Sprintf("%s/%s/%s/%s",
			WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName),
		nil)
	if err != nil {
		return err
	}

	return nil
}

// GetSourceConfig retrieves the schema config of a given source
// API Doc: https://reference.segmentapis.com/#c74efb9b-b09e-4072-8da1-ba6ca60e6a78
func (c *Client) GetSourceConfig(srcName string) (SourceConfig, error) {
	var result SourceConfig

	response, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s/schema-config", WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName), nil)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// UpdateSourceConfig updates the schema config of a given source
// API Doc: https://reference.segmentapis.com/#af54244f-4ec7-4e78-96e9-8966dd18e56f
func (c *Client) UpdateSourceConfig(srcName string, config SourceConfig) (SourceConfig, error) {
	var result SourceConfig

	req := sourceConfigUpdateRequest{
		Config: config,
		UpdateMask: UpdateMask{Paths: []string{
			"schema_config.allow_unplanned_track_events",
			"schema_config.allow_unplanned_identify_traits",
			"schema_config.allow_unplanned_group_traits",
			"schema_config.forwarding_blocked_events_to",
			"schema_config.allow_unplanned_track_event_properties",
			"schema_config.allow_track_event_on_violations",
			"schema_config.allow_identify_traits_on_violations",
			"schema_config.allow_group_traits_on_violations",
			"schema_config.forwarding_violations_to",
			"schema_config.common_track_event_on_violations",
			"schema_config.common_identify_event_on_violations",
			"schema_config.common_group_event_on_violations",
		}},
	}

	response, err := c.doRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/%s/schema-config", WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName), req)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
