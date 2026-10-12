// In the Temporal Go SDK, a Workflow Definition is an exported function with two additional requirements:
// it must accept workflow.Context as the first input parameter, and it must return error. 

// Your Workflow function can optionally return another value, which you'll use to return the result of the Workflow Execution. 
// https://docs.temporal.io/develop/go/workflows/basics?_gl=1*1s6qaay*_gcl_aw*R0NMLjE3OTA3NDAwMjguQ2p3S0NBand3LTNWQmhBY0Vpd0F3VVVJdTNPdUlxSEtIYUtoMHpMZ1BYOTloM1hrd2x5NVlKYWhCajhmMTJvQmpqQUkxb0FETm9HZ0FCb0NJcUVRQXZEX0J3RQ..*_gcl_au*MTQ1MDU4Njg4OC4xNzkwNjQ2NzI2Li0uLS4xNzkwNjQ2NzI1LjEyMjE4ODAyNjIuMTc5MDc0MDAyOC4xNzkwNzQwMDI3*_ga*MjQ1NzI0MTcwLjE3OTA2NDY3MTE.*_ga_R90Q9SJD3D*czE3OTA3NDI1ODgkbzQkZzEkdDE3OTA3NDI2MDQkajQ0JGwwJGgw#develop-workflows

// Temporal Workflows must be deterministic so that Temporal can replay your Workflow in the event of a crash. 
// That's why you call Activities from your Workflow code. 
// Activities don't have the same determinism constraints that Workflows have.
package iplocate

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// GetAddressFromIP is the Temporal Workflow that retrieves the IP address and location info.
func GetAddressFromIP(ctx workflow.Context, name string) (string, error) {
	// Define the activity options, including the retry policy
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second, //amount of time that must elapse before the first retry occurs
			MaximumInterval:    time.Minute, //maximum interval between retries
			BackoffCoefficient: 2,           //how much the retry interval increases
			// MaximumAttempts: 5, // Uncomment this if you want to limit attempts
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var ipActivities *IPActivities

	var ip string
	err := workflow.ExecuteActivity(ctx, ipActivities.GetIP).Get(ctx, &ip)
	if err != nil {
		return "", fmt.Errorf("Failed to get IP: %s", err)
	}

	var location string
	err = workflow.ExecuteActivity(ctx, ipActivities.GetLocationInfo, ip).Get(ctx, &location)
	if err != nil {
		return "", fmt.Errorf("Failed to get location: %s", err)
	}
	return fmt.Sprintf("Hello, %s. Your IP is %s and your location is %s", name, ip, location), nil
}