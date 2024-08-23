package awx

import (
	"net/http"
)

type Client struct {
	JobTemplateService   *JobTemplateService
	InventoriesService   *InventoriesService
	HostService          *HostService
	JobService           *JobService
	OrganizationsService *OrganizationsService
}

func NewClient(baseURL string, username string, password string) (*Client, error) {

	tokenAuth := BasicAuth{
		Username: username,
		Password: password,
	}

	requester := Requester{
		Base:   baseURL,
		Auth:   &tokenAuth,
		Client: http.DefaultClient,
	}

	client := Client{
		JobTemplateService: &JobTemplateService{
			Requester: &requester,
		},
		InventoriesService: &InventoriesService{
			Requester: &requester,
		},
		HostService: &HostService{
			Requester: &requester,
		},
		JobService: &JobService{
			Requester: &requester,
		},
		OrganizationsService: &OrganizationsService{
			Requester: &requester,
		},
	}

	return &client, nil
}

func NewClientWithToken(baseURL string, token string) (*Client, error) {

	tokenAuth := TokenAuth{
		Token: token,
	}

	requester := Requester{
		Base:   baseURL,
		Auth:   &tokenAuth,
		Client: http.DefaultClient,
	}

	client := Client{
		JobTemplateService: &JobTemplateService{
			Requester: &requester,
		},
		InventoriesService: &InventoriesService{
			Requester: &requester,
		},
		HostService: &HostService{
			Requester: &requester,
		},
		JobService: &JobService{
			Requester: &requester,
		},
		OrganizationsService: &OrganizationsService{
			Requester: &requester,
		},
	}

	return &client, nil
}
