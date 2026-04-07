// Package datamammoth implements a Packer builder plugin that creates server
// images on the DataMammoth platform.
//
// The builder provisions a temporary server, runs provisioners (shell, Ansible,
// etc.), creates a snapshot, and tears down the server.
package datamammoth

import (
	"context"
	"fmt"
)

// Config holds the builder configuration parsed from HCL/JSON.
type Config struct {
	// APIKey is the DataMammoth API key (dm_live_... or dm_test_...).
	APIKey string `mapstructure:"api_key" required:"true"`

	// BaseURL overrides the API base URL (for self-hosted or staging).
	BaseURL string `mapstructure:"base_url"`

	// ProductID is the hosting product to use for the temporary server.
	ProductID string `mapstructure:"product_id" required:"true"`

	// Region is the datacenter region (e.g., "EU", "US-central").
	Region string `mapstructure:"region" required:"true"`

	// ImageID is the base OS image to start from.
	ImageID string `mapstructure:"image_id" required:"true"`

	// SnapshotName is the name for the resulting snapshot.
	SnapshotName string `mapstructure:"snapshot_name"`

	// SSHUsername for connecting to the temporary server.
	SSHUsername string `mapstructure:"ssh_username"`
}

// Builder implements the packer.Builder interface.
type Builder struct {
	config Config
}

// Prepare validates the configuration.
func (b *Builder) Prepare(raws ...interface{}) ([]string, []string, error) {
	// In production, this would use mapstructure to decode raws into b.config
	// and validate required fields.
	return nil, nil, nil
}

// Run executes the build:
//  1. Create a temporary server via POST /api/v2/servers
//  2. Wait for provisioning (poll task)
//  3. Connect via SSH and yield to Packer provisioners
//  4. Create snapshot via POST /api/v2/servers/{id}/snapshots
//  5. Terminate the temporary server
//  6. Return the snapshot ID as an artifact
func (b *Builder) Run(ctx context.Context, ui interface{}, hook interface{}) (interface{}, error) {
	return nil, fmt.Errorf("packer-plugin-datamammoth: Run() not yet implemented")
}

// PluginSet returns the Packer plugin set with the datamammoth builder registered.
// In production, this would return a *plugin.Set. For now it returns a placeholder.
func PluginSet() *PluginSetPlaceholder {
	return &PluginSetPlaceholder{}
}

// PluginSetPlaceholder is a stand-in until the real packer-plugin-sdk is wired up.
type PluginSetPlaceholder struct{}
