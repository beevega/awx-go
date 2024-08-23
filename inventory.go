package awx

import (
	"context"
	"fmt"
)

// InventoriesService implements awx inventories apis.
type InventoriesService struct {
	Requester *Requester
}

// ListInventories shows list of awx inventories.
func (i *InventoriesService) ListInventories(ctx context.Context, params map[string]string) (*ListInventories, error) {
	result := ListInventories{}
	endpoint := "/api/v2/inventories/"

	_, err := i.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateInventory creates an awx inventory.
func (i *InventoriesService) CreateInventory(ctx context.Context, data map[string]interface{}) (*Inventory, error) {
	result := Inventory{}
	endpoint := "/api/v2/inventories/"

	validate, status := ValidateParams(data, []string{"name", "organization"})
	if !status {
		return nil, fmt.Errorf("mandatory input arguments are absent: %s", validate)
	}

	// Add check if inventory exists and return proper error
	_, err := i.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateInventory update an awx inventory
func (i *InventoriesService) UpdateInventory(ctx context.Context, id int, data map[string]interface{}) (*Inventory, error) {
	result := Inventory{}
	endpoint := fmt.Sprintf("/api/v2/inventories/%d", id)

	_, err := i.Requester.Patch(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetInventory retrives the inventory information from its ID or Name
func (i *InventoriesService) GetInventory(ctx context.Context, id int) (*Inventory, error) {
	result := Inventory{}
	endpoint := fmt.Sprintf("/api/v2/inventories/%d", id)

	_, err := i.Requester.Get(ctx, endpoint, &result, map[string]string{})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteInventory delete an inventory from AWX
func (i *InventoriesService) DeleteInventory(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("/api/v2/inventories/%d", id)

	_, err := i.Requester.Delete(ctx, endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (i *InventoriesService) SyncInventorySourcesByInventoryID(ctx context.Context, id int) ([]*InventoryUpdate, error) {
	result := make([]*InventoryUpdate, 0)
	endpoint := fmt.Sprintf("/api/v2/inventories/%d/update_inventory_sources/", id)

	_, err := i.Requester.Post(ctx, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
