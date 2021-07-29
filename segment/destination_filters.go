package segment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var updateMask = newUpdateMask("if", "actions", "title", "description", "enabled")

// ListDestinations returns all destinations for a source
func (c *Client) ListDestinationFilters(srcName string, destinationName string) ([]DestinationFilter, error) {
	var d destinationFiltersListResponse
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s",
			WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName, DestinationEndpoint, destinationName, DestinationFiltersEndpoint),
		nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal destinations response")
	}

	return d.Filters, nil
}

func (c *Client) CreateDestinationFilter(srcName string, destinationName string, filter DestinationFilter) (*DestinationFilter, error) {
	data, err := c.doRequest(http.MethodPost,
		fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s",
			WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName, DestinationEndpoint, destinationName, DestinationFiltersEndpoint),
		destinationFilterCRURequest{Filter: filter, UpdateMask: updateMask})
	if err != nil {
		return nil, err
	}

	var result DestinationFilter
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal filter")
	}

	return &result, nil
}

func (c *Client) UpdateDestinationFilter(srcName string, destinationName string, filter DestinationFilter) (*DestinationFilter, error) {
	data, err := c.doRequest(http.MethodPatch, filter.Name, destinationFilterCRURequest{Filter: filter, UpdateMask: updateMask})
	if err != nil {
		return nil, err
	}

	var result DestinationFilter
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal filter")
	}

	return &result, nil
}

func (c *Client) GetDestinationFilter(srcName string, destinationName string, filterId string) (*DestinationFilter, error) {
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s",
			WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName, DestinationEndpoint, destinationName, DestinationFiltersEndpoint, filterId),
		nil)
	if err != nil {
		return nil, err
	}

	var d destinationFilterCRURequest
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal filter")
	}

	return &d.Filter, nil
}

func (c *Client) DeleteDestinationFilter(srcName string, destinationName string, filterId string) error {
	_, err := c.doRequest(http.MethodDelete,
		fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s",
			WorkspacesEndpoint, c.workspace, SourceEndpoint, srcName, DestinationEndpoint, destinationName, DestinationFiltersEndpoint, filterId),
		nil)

	return err
}
