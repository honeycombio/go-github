package github

import (
	"context"
	"net/http"
	"testing"
)

func TestRepositoriesService_PassFailDeploymentGate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/runs/1/deployment_protection_rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
	})

	ctx := context.Background()
	_, err := client.Repositories.PassFailDeploymentGate(ctx, "o", "r", 1, &PassFailDeploymentGateRequest{
		State:           String("approved"),
		Comment:         String("comment"),
		EnvironmentName: String("prod"),
	})
	if err != nil {
		t.Errorf("Repositories.DeleteDeploymentBranchPolicy returned error: %v", err)
	}

	const methodName = "PassFailDeploymentGate"
	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		return client.Repositories.PassFailDeploymentGate(ctx, "o", "r", 1, &PassFailDeploymentGateRequest{
			State:           String("approved"),
			Comment:         String("comment"),
			EnvironmentName: String("prod"),
		})
	})
}
