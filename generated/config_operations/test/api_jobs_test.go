/*
Config Operations Testing JobsAPIService
*/
package config_operations

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_config_operations_JobsAPIService_List tests listing jobs
// This is a read-only operation that retrieves the current list of configuration jobs
func Test_config_operations_JobsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigOperationsTestClient(t)

	fmt.Printf("Listing configuration jobs\n")

	// Make the list request to the API
	req := client.JobsAPI.ListJobs(context.Background())
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the request was successful
	require.NoError(t, err, "Failed to list jobs")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")

	// Assert response properties
	require.NotNil(t, res, "Response should not be nil")

	// Get the data using the GetData method
	jobs := res.GetData()
	fmt.Printf("Successfully retrieved %d jobs (limit: %d, offset: %d, total: %d)\n",
		len(jobs), res.Limit, res.Offset, res.Total)

	// If there are jobs, verify the structure
	if len(jobs) > 0 {
		firstJob := jobs[0]
		assert.NotEmpty(t, firstJob.Id, "Job should have an ID")
		assert.NotEmpty(t, firstJob.JobType, "Job should have a type")
		assert.NotEmpty(t, firstJob.StatusStr, "Job should have a status string")

		fmt.Printf("Sample job - ID: %s, Type: %s, Status: %s\n",
			firstJob.Id, firstJob.JobType, firstJob.StatusStr)
	} else {
		fmt.Printf("No jobs found in the system\n")
	}

	t.Logf("Successfully listed jobs")
}

// Test_config_operations_JobsAPIService_GetByID tests retrieving a specific job by ID
// This test first lists jobs to find a valid ID, then retrieves that specific job
func Test_config_operations_JobsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigOperationsTestClient(t)

	// First, list jobs to get a valid job ID
	listReq := client.JobsAPI.ListJobs(context.Background())
	listRes, _, listErr := listReq.Execute()
	if listErr != nil {
		handleAPIError(listErr)
	}
	require.NoError(t, listErr, "Failed to list jobs")

	// Get the jobs data
	jobs := listRes.GetData()

	// Skip test if no jobs exist
	if len(jobs) == 0 {
		t.Skip("No jobs available to test GetByID - skipping test")
		return
	}

	// Get the first job's ID (string type)
	jobIDStr := jobs[0].Id
	fmt.Printf("Retrieving job with ID: %s\n", jobIDStr)

	// Retrieve the specific job by ID
	req := client.JobsAPI.GetJobsByID(context.Background(), jobIDStr)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the request was successful
	require.NoError(t, err, "Failed to get job by ID")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")

	// Assert response properties
	require.NotNil(t, res, "Response should not be nil")

	// Get the data - API returns array with jobs
	retrievedJobs := res.GetData()
	require.Greater(t, len(retrievedJobs), 0, "Should have at least one job in response")

	// Verify we got the job we requested
	foundJob := retrievedJobs[0]
	assert.Equal(t, jobIDStr, foundJob.Id, "Retrieved job ID should match requested ID")

	fmt.Printf("Successfully retrieved job - ID: %s, Type: %s, Status: %s\n",
		foundJob.Id, foundJob.JobType, foundJob.StatusStr)

	t.Logf("Successfully retrieved job by ID: %s", jobIDStr)
}
