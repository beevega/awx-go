package awx

import (
	"context"
	"fmt"
)

// HostService implements awx Hosts apis.
type HostService struct {
	Requester *Requester
}

// ListHosts shows list of awx Hosts.
func (h *HostService) ListHosts(ctx context.Context, params map[string]string) (*ListHosts, error) {
	result := ListHosts{}
	endpoint := "/api/v2/hosts/"

	_, err := h.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateHost creates an awx Host.
//
//	name TEXT
//	description TEXT
//	inventory ID
//	enabled BOOLEAN
//	instance_id TEXT
//	variables JSON/YAML
func (h *HostService) CreateHost(ctx context.Context, data map[string]interface{}) (*Host, error) {
	result := Host{}
	endpoint := "/api/v2/hosts/"

	validate, status := ValidateParams(data, []string{"name", "inventory"})
	if !status {
		return nil, fmt.Errorf("mandatory input arguments are absent: %s", validate)
	}

	// Add check if Host exists and return proper error
	_, err := h.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateHost update an awx Host
func (h *HostService) UpdateHost(ctx context.Context, id int, data map[string]interface{}) (*Host, error) {
	result := Host{}
	endpoint := fmt.Sprintf("/api/v2/hosts/%d", id)

	_, err := h.Requester.Patch(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AssociateGroup update an awx Host
func (h *HostService) AssociateGroup(ctx context.Context, id int, data map[string]interface{}) (*Host, error) {
	result := Host{}
	endpoint := fmt.Sprintf("/api/v2/hosts/%d/groups/", id)
	data["associate"] = true

	validate, status := ValidateParams(data, []string{"id"})
	if !status {
		return nil, fmt.Errorf("mandatory input arguments are absent: %s", validate)
	}

	_, err := h.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DisAssociateGroup update an awx Host
func (h *HostService) DisAssociateGroup(ctx context.Context, id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := Host{}
	endpoint := fmt.Sprintf("/api/v2/hosts/%d/groups/", id)
	data["disassociate"] = true

	validate, status := ValidateParams(data, []string{"id"})
	if !status {
		return nil, fmt.Errorf("mandatory input arguments are absent: %s", validate)
	}

	_, err := h.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteHost delete an awx Host.
func (h *HostService) DeleteHost(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("/api/v2/hosts/%d", id)

	_, err := h.Requester.Delete(ctx, endpoint)
	if err != nil {
		return err
	}

	return nil
}
