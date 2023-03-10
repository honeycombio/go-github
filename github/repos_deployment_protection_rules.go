// Copyright 2023 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
)

// PassFailDeploymentGateRequest represents a deployment branch policy request.
type PassFailDeploymentGateRequest struct {
	State           *string `json:"state,omitempty"`
	Comment         *string `json:"comment,omitempty"`
	EnvironmentName *string `json:"environment_name,omitempty"`
}

// PassFailDeploymentGate is a POST request to determine if a deployment gate should pass or fail
func (s *RepositoriesService) PassFailDeploymentGate(ctx context.Context, callbackURL string, request *PassFailDeploymentGateRequest) (*Response, error) {
	req, err := s.client.NewRequest("POST", callbackURL, request)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
