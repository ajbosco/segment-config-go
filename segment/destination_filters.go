package segment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ListDestinations returns all destinations for a source
func (c *Client) ListDestinationFilters(srcName string, destinationName string) ([]DestinationFilter, error) {
	var d DestinationFilters
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
