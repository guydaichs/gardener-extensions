// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/internal/imagevector"
	"github.com/gardener/gardener-extensions/pkg/terraformer"

	"github.com/gardener/gardener/pkg/logger"
	"k8s.io/client-go/rest"
)

const (
	// TerraformVarServiceAccount is the name of the terraform service account environment variable.
	TerraformVarServiceAccount = "TF_VAR_SERVICEACCOUNT"
)

// TerraformerVariablesEnvironmentFromServiceAccount computes the Terraformer variables environment from the
// given ServiceAccount.
func TerraformerVariablesEnvironmentFromServiceAccount(account *ServiceAccount) (map[string]string, error) {
	var buf bytes.Buffer
	if err := json.Compact(&buf, account.Raw); err != nil {
		return nil, err
	}

	return map[string]string{
		TerraformVarServiceAccount: buf.String(),
	}, nil
}

// NewTerraformer initializes a new Terraformer that has the ServiceAccount credentials.
func NewTerraformer(
	restConfig *rest.Config,
	serviceAccount *ServiceAccount,
	purpose,
	namespace,
	name string,
) (terraformer.Terraformer, error) {
	tf, err := terraformer.NewForConfig(logger.NewLogger("info"), restConfig, purpose, namespace, name, imagevector.TerraformerImage())
	if err != nil {
		return nil, err
	}

	variables, err := TerraformerVariablesEnvironmentFromServiceAccount(serviceAccount)
	if err != nil {
		return nil, err
	}

	return tf.
		SetVariablesEnvironment(variables).
		SetJobBackoffLimit(0).
		SetActiveDeadlineSeconds(630).
		SetDeadlineCleaning(5 * time.Minute).
		SetDeadlinePod(15 * time.Minute).
		SetDeadlineJob(15 * time.Minute), nil
}
