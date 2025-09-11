package awx

import (
	"context"
	"fmt"
)

// Enum of job statuses.
const (
	JobStatusNew        = "new"
	JobStatusPending    = "pending"
	JobStatusWaiting    = "waiting"
	JobStatusRunning    = "running"
	JobStatusSuccessful = "successful"
	JobStatusFailed     = "failed"
	JobStatusError      = "error"
	JobStatusCanceled   = "canceled"
)

// JobService implements awx job apis.
type JobService struct {
	Requester *Requester
}

// GetJob shows the details of a job.
func (j *JobService) GetJob(ctx context.Context, id int, params map[string]string) (*Job, error) {
	result := Job{}
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/", id)

	_, err := j.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CancelJob cancels a job.
func (j *JobService) CancelJob(ctx context.Context, id int, data map[string]interface{}) (*CancelJobResponse, error) {
	result := CancelJobResponse{}
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/cancel/", id)

	_, err := j.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RelaunchJob relaunch a job.
func (j *JobService) RelaunchJob(ctx context.Context, id int, data map[string]interface{}) (*JobLaunch, error) {
	result := JobLaunch{}
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/relaunch/", id)

	_, err := j.Requester.Post(ctx, endpoint, data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetHostSummaries get a job hosts summaries.
func (j *JobService) GetHostSummaries(ctx context.Context, id int, params map[string]string) (*HostSummaries, error) {
	result := HostSummaries{}
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/job_host_summaries/", id)

	_, err := j.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetJobEvents get a list of job events.
func (j *JobService) GetJobEvents(ctx context.Context, id int, params map[string]string) (*JobEvents, error) {
	result := JobEvents{}
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/job_events/", id)

	_, err := j.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetJobStdout get a stdout result of finished job.
func (j *JobService) GetJobStdout(ctx context.Context, id int, params map[string]string) (*JobStdout, error) {
	result := JobStdout{}
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/stdout/", id)

	// Hard set format=json
	params["format"] = "json"

	_, err := j.Requester.Get(ctx, endpoint, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
