// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package provisioner

import (
	"context"

	"github.com/daytonaio/daytona/pkg/containerregistry"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/provider/manager"
	"github.com/daytonaio/daytona/pkg/target"
	"github.com/daytonaio/daytona/pkg/workspace"
)

type IProvisioner interface {
	CreateTarget(target *target.Target, targetConfig *provider.TargetConfig) error
	StartTarget(target *target.Target, targetConfig *provider.TargetConfig) error
	StopTarget(target *target.Target, targetConfig *provider.TargetConfig) error
	GetTargetInfo(ctx context.Context, target *target.Target, targetConfig *provider.TargetConfig) (*target.TargetInfo, error)
	DestroyTarget(target *target.Target, targetConfig *provider.TargetConfig) error

	CreateWorkspace(workspace *workspace.Workspace, targetConfig *provider.TargetConfig, cr *containerregistry.ContainerRegistry, gc *gitprovider.GitProviderConfig) error
	DestroyWorkspace(workspace *workspace.Workspace, targetConfig *provider.TargetConfig) error
	StartWorkspace(workspace *workspace.Workspace, targetConfig *provider.TargetConfig) error
	StopWorkspace(workspace *workspace.Workspace, targetConfig *provider.TargetConfig) error
	GetWorkspaceInfo(ctx context.Context, workspace *workspace.Workspace, targetConfig *provider.TargetConfig) (*workspace.WorkspaceInfo, error)
}

type ProvisionerConfig struct {
	ProviderManager manager.IProviderManager
}

func NewProvisioner(config ProvisionerConfig) IProvisioner {
	return &Provisioner{
		providerManager: config.ProviderManager,
	}
}

type Provisioner struct {
	providerManager manager.IProviderManager
}
