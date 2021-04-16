package segment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// ListTrackingPlans lists all the tracking plans in the workspace
func (c *Client) ListTrackingPlans() (TrackingPlans, error) {
	var tps TrackingPlans
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
		nil)
	if err != nil {
		return tps, err
	}
	err = json.Unmarshal(data, &tps)
	if err != nil {
		return tps, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tps, nil
}

// GetTrackingPlan gets a specific tracking plan from segment
func (c *Client) GetTrackingPlan(trackingPlanID string) (TrackingPlan, error) {
	var tp TrackingPlan
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, trackingPlanID),
		nil)
	if err != nil {
		return tp, err
	}

	err = json.Unmarshal(data, &tp)
	if err != nil {
		return tp, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tp, nil
}

// CreateTrackingPlan creates tracking plan
func (c *Client) CreateTrackingPlan(data TrackingPlan) (TrackingPlan, error) {
	var tp TrackingPlan
	tpCreateReq := trackingPlanCreateRequest{
		TrackingPlan: data,
	}
	responseBody, err := c.doRequest(http.MethodPost,
		fmt.Sprintf("%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
		tpCreateReq)

	if err != nil {
		return tp, err
	}
	err = json.Unmarshal(responseBody, &tp)
	if err != nil {
		return tp, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tp, nil
}

// UpdateTrackingPlan updates a tracking plan
func (c *Client) UpdateTrackingPlan(trackingPlanID string, data TrackingPlan) (TrackingPlan, error) {
	var tp TrackingPlan

	um := UpdateMask{
		Paths: []string{"tracking_plan.display_name", "tracking_plan.rules"},
	}
	tpUpdateReq := trackingPlanUpdateRequest{
		UpdateMask:   um,
		TrackingPlan: data,
	}
	responseBody, err := c.doRequest(http.MethodPut,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, trackingPlanID),
		tpUpdateReq)

	if err != nil {
		return tp, err
	}
	err = json.Unmarshal(responseBody, &tp)
	if err != nil {
		return tp, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tp, nil
}

// DeleteTrackingPlan Deletes a tracking plan
func (c *Client) DeleteTrackingPlan(trackingPlanID string) error {

	_, err := c.doRequest(http.MethodDelete,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, trackingPlanID),
		nil)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateTrackingPlanSourceConnection(planId string, sourceName string) error {
	data, err := c.doRequest(http.MethodPost,
		fmt.Sprintf("%s/%s/%s/%s/source-connections",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, planId),
		trackingPlanSourceConnectionCreateRequest{Name: fmt.Sprintf("workspaces/%s/sources/%s", c.workspace, sourceName)})
	if err != nil {
		return err
	}

	var result TrackingPlanSourceConnection
	if err := json.Unmarshal(data, &result); err != nil {
		return errors.Errorf("Unexpected response body: %s", string(data))
	}

	return nil
}

func (c *Client) ListTrackingPlanSources(planId string) ([]TrackingPlanSourceConnection, error) {
	var connections TrackingPlanSourceConnections
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s/source-connections",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, planId),
		nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &connections); err != nil {
		return connections.Connections, err
	}

	return connections.Connections, nil
}

func (c *Client) DeleteTrackingPlanSourceConnection(planId string, sourceName string) error {
	data, err := c.doRequest(http.MethodDelete,
		fmt.Sprintf("%s/%s/%s/%s/source-connections/%s",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, planId, sourceName),
		nil)
	if err != nil {
		return err
	}

	if strings.TrimSpace(string(data)) != "{}" {
		return errors.Errorf("Unexpected response body: %s", string(data))
	}

	return nil
}

// TODO: Batch create tracking plan source connections
