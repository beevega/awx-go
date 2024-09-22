package awx

import (
	"context"
	"fmt"
)

// HostService implements awx Hosts apis.
type GroupService struct {
	Requester *Requester
}

// ListGroups shows list of awx Groups.
//
// Each group data structure includes the following fields:
//
//	id: Database ID for this group. (integer)
//	type: Data type for this group. (choice)
//	url: URL for this group. (string)
//	related: Data structure with URLs of related resources. (object)
//	summary_fields: Data structure with name/description for related resources.
//		The output for some objects may be limited for performance reasons. (object)
//	created: Timestamp when this group was created. (datetime)
//	modified: Timestamp when this group was last modified. (datetime)
//	name: Name of this group. (string)
//	description: Optional description of this group. (string)
//	inventory: (id)
//	variables: Group variables in JSON or YAML format. (json)
func (g *GroupService) ListGroups(ctx context.Context, params map[string]string) (*ListGroups, error) {
	result := ListGroups{}
	endpoint := "/api/v2/groups/"

	_, err := g.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	// if err := CheckResponse(resp); err != nil {
	// 	return nil, result, err
	// }

	return &result, nil
}

// ListGroupsByInventoryId shows list of groups that created in specify inventory.
func (g *GroupService) ListGroupsByInventoryId(ctx context.Context, inventoryId int) (*ListGroups, error) {
	result := ListGroups{}
	endpoint := fmt.Sprintf("/api/v2/inventories/%d/groups/", inventoryId)

	_, err := g.Requester.Get(ctx, endpoint, &result, nil)
	if err != nil {
		return nil, err
	}

	// if err := CheckResponse(resp); err != nil {
	// 	return nil, result, err
	// }

	return &result, nil
}

// CreateGroup creates an awx Group.
func (g *GroupService) CreateGroup(ctx context.Context, data map[string]interface{}) (*Group, error) {
	result := Group{}
	endpoint := "/api/v2/groups/"

	validate, status := ValidateParams(data, []string{"name", "inventory"})
	if !status {
		return nil, fmt.Errorf("mandatory input arguments are absent: %s", validate)
	}

	_, err := g.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	// if err := CheckResponse(resp); err != nil {
	// 	return nil, err
	// }

	return &result, nil
}

// UpdateGroup update an awx group.
func (g *GroupService) UpdateGroup(ctx context.Context, id int, data map[string]interface{}) (*Group, error) {
	result := Group{}
	endpoint := fmt.Sprintf("/api/v2/groups/%d", id)

	_, err := g.Requester.Patch(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	// if err := CheckResponse(resp); err != nil {
	// 	return nil, err
	// }

	return &result, nil
}

// DeleteGroup delete an awx Group.
func (g *GroupService) DeleteGroup(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("/api/v2/groups/%d", id)

	_, err := g.Requester.Delete(ctx, endpoint)
	if err != nil {
		return err
	}

	// if err := CheckResponse(resp); err != nil {
	// 	return nil, err
	// }

	return nil
}

func (g *GroupService) AddHostToGroup(ctx context.Context, id int, inventoryId int, name string) error {
	endpoint := fmt.Sprintf("/api/v2/groups/%d/hosts/", id)

	payload := map[string]interface{}{
		"inventory": inventoryId,
		"name":      name,
	}

	_, err := g.Requester.Post(ctx, endpoint, payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) AddChildrenToGroup(ctx context.Context, id int, childGroupId int) (*Group, error) {
	result := Group{}
	endpoint := fmt.Sprintf("/api/v2/groups/%d/children/", id)

	payload := map[string]int{
		"id": childGroupId,
	}

	_, err := g.Requester.Post(ctx, endpoint, payload, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
