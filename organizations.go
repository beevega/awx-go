package awx

import "context"

// OrganizationsService implements awx groups apis.
type OrganizationsService struct {
	Requester *Requester
}

// Create shows list of awx groups.
func (i *OrganizationsService) List(ctx context.Context, params map[string]string) (*ListOrganizations, error) {
	result := ListOrganizations{}
	endpoint := "/api/v2/organizations/"

	_, err := i.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Create creates an awx group.
// func (i *OrganizationsService) Create(ctx context.Context, data map[string]interface{}) (*Inventory, error) {
// 	result := Organization{}
// 	endpoint := "/api/v2/organizations/"

// 	validate, status := ValidateParams(data, []string{"name"})
// 	if !status {
// 		return nil, fmt.Errorf("mandatory input arguments are absent: %s", validate)
// 	}

// 	// Add check if inventory exists and return proper error
// 	_, err := i.Requester.Post(ctx, endpoint, data, &result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &result, nil
// }
