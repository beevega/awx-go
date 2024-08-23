package awx

import (
	"context"
	"errors"
	"fmt"
)

// JobTemplateService implements awx job template apis.
type JobTemplateService struct {
	Requester *Requester
}

// ListJobTemplates shows a list of job templates.
func (jt *JobTemplateService) ListJobTemplates(ctx context.Context, params map[string]string) (*ListJobTemplates, error) {
	result := ListJobTemplates{}
	endpoint := "/api/v2/job_templates/"

	_, err := jt.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Launch lauchs a job with the job template
//
//	monitor
//	timeout TIMEOUT
//	wait
//	extra_vars JSON/YAML
//	inventory ID
//	scm_branch TEXT
//	limit TEXT
//	job_tags TEXT
//	skip_tags TEXT
//	job_type {run,check}
//	verbosity {0,1,2,3,4,5}
//	diff_mode BOOLEAN
//	credentials [ID, ID, ...]
//	credential_passwords JSON/YAML
//
// You can limit the collection with the following patterns:
//
// 1 Host = Just the hostname
// Multiple hosts = host1,host2,host3
// Group = Groupname
// Multiple Groups = Group1:Group2:Group3
// Exclude Group in Group = Group1:!Group4
func (jt *JobTemplateService) Launch(ctx context.Context, id int, data map[string]interface{}) (*JobLaunch, error) {
	result := JobLaunch{}
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d/launch/", id)

	_, err := jt.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	// in case invalid job id return
	if result.Job == 0 {
		return nil, errors.New("invalid job id 0")
	}

	return &result, nil
}

// CreateJobTemplate creates a job template
//
//	name TEXT *REQUIRED
//	description TEXT
//	job_type {run,check} *REQUIRED
//	inventory ID *REQUIRED
//	project ID *REQUIRED
//	playbook TEXT
//	scm_branch TEXT
//	forks INTEGER
//	limit TEXT
//	verbosity {0,1,2,3,4,5}
//	extra_vars JSON/YAML
//	job_tags TEXT
//	force_handlers BOOLEAN
//	skip_tags TEXT
//	start_at_task TEXT
//	timeout INTEGER
//	use_fact_cache BOOLEAN
//	host_config_key TEXT
//	ask_scm_branch_on_launch BOOLEAN
//	ask_diff_mode_on_launch BOOLEAN
//	ask_variables_on_launch BOOLEAN
//	ask_limit_on_launch BOOLEAN
//	ask_tags_on_launch BOOLEAN
//	ask_skip_tags_on_launch BOOLEAN
//	ask_job_type_on_launch BOOLEAN
//	ask_verbosity_on_launch BOOLEAN
//	ask_inventory_on_launch BOOLEAN
//	ask_credential_on_launch BOOLEAN
//	survey_enabled BOOLEAN
//	become_enabled BOOLEAN
//	diff_mode BOOLEAN
//	allow_simultaneous BOOLEAN
//	custom_virtualenv TEXT
//	job_slice_count INTEGER
//	webhook_service {,github,gitlab}
//	webhook_credential ID
func (jt *JobTemplateService) CreateJobTemplate(ctx context.Context, data map[string]interface{}) (*JobTemplate, error) {
	result := JobTemplate{}
	endpoint := "/api/v2/job_templates/"

	validate, status := ValidateParams(data, []string{"name", "job_type", "inventory", "project"})
	if !status {
		return nil, fmt.Errorf("mandatory input arguments are absent: %s", validate)
	}

	_, err := jt.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateJobTemplate updates a job template
//
//	name TEXT
//	description TEXT
//	job_type {run,check}
//	inventory ID
//	project ID
//	playbook TEXT
//	scm_branch TEXT
//	forks INTEGER
//	limit TEXT
//	verbosity {0,1,2,3,4,5}
//	extra_vars JSON/YAML
//	job_tags TEXT
//	force_handlers BOOLEAN
//	skip_tags TEXT
//	start_at_task TEXT
//	timeout INTEGER
//	use_fact_cache BOOLEAN
//	host_config_key TEXT
//	ask_scm_branch_on_launch BOOLEAN
//	ask_diff_mode_on_launch BOOLEAN
//	ask_variables_on_launch BOOLEAN
//	ask_limit_on_launch BOOLEAN
//	ask_tags_on_launch BOOLEAN
//	ask_skip_tags_on_launch BOOLEAN
//	ask_job_type_on_launch BOOLEAN
//	ask_verbosity_on_launch BOOLEAN
//	ask_inventory_on_launch BOOLEAN
//	ask_credential_on_launch BOOLEAN
//	survey_enabled BOOLEAN
//	become_enabled BOOLEAN
//	diff_mode BOOLEAN
//	allow_simultaneous BOOLEAN
//	custom_virtualenv TEXT
//	job_slice_count INTEGER
//	webhook_service {github,gitlab}
//	webhook_credential ID
func (jt *JobTemplateService) UpdateJobTemplate(ctx context.Context, id int, data map[string]interface{}) (*JobTemplate, error) {
	result := JobTemplate{}
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d", id)

	_, err := jt.Requester.Patch(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteJobTemplate deletes a job template
func (jt *JobTemplateService) DeleteJobTemplate(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d", id)

	_, err := jt.Requester.Delete(ctx, endpoint)
	if err != nil {
		return err
	}

	return nil
}
