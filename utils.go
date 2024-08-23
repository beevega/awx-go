package awx

import (
	"context"
	"fmt"
	"time"
)

// WaitForJobFinish ожидает что у задания будет один из статусов, указывающих на завершение задания.
// Перечень статусов:
// successful
// failed
// error
// canceled
func WaitForSuccessJobFinish(c *Client, id int, secs int) error {
	return waitFor(secs, func() (bool, error) {

		job, err := c.JobService.GetJob(context.Background(), id, map[string]string{})
		if err != nil {
			return false, err
		}

		switch job.Status {
		case JobStatusSuccessful:
			return true, nil
		case JobStatusFailed:
			return true, fmt.Errorf("task finished with bad status: %s", JobStatusFailed)
		case JobStatusError:
			return true, fmt.Errorf("task finished with bad status: %s", JobStatusError)
		case JobStatusCanceled:
			return true, fmt.Errorf("task finished with bad status: %s", JobStatusCanceled)
		}

		return false, nil
	})
}

// WaitFor polls a predicate function, once per second, up to a timeout limit.
// This is useful to wait for a resource to transition to a certain state.
// To handle situations when the predicate might hang indefinitely, the
// predicate will be prematurely cancelled after the timeout.
// Resource packages will wrap this in a more convenient function that's
// specific to a certain resource, but it can also be useful on its own.
func waitFor(timeout int, predicate func() (bool, error)) error {
	type WaitForResult struct {
		Success bool
		Error   error
	}

	start := time.Now().Unix()

	for {
		// If a timeout is set, and that's been exceeded, shut it down.
		if timeout >= 0 && time.Now().Unix()-start >= int64(timeout) {
			return fmt.Errorf("a timeout occurred")
		}

		time.Sleep(1 * time.Second)

		var result WaitForResult
		ch := make(chan bool, 1)
		go func() {
			defer close(ch)
			satisfied, err := predicate()
			result.Success = satisfied
			result.Error = err
		}()

		select {
		case <-ch:
			if result.Error != nil {
				return result.Error
			}
			if result.Success {
				return nil
			}
		// If the predicate has not finished by the timeout, cancel it.
		case <-time.After(time.Duration(timeout) * time.Second):
			return fmt.Errorf("a timeout occurred")
		}
	}
}
