package models

import (
	"github.com/porter-dev/porter/internal/models/integrations"
	"gorm.io/gorm"
)

// Registry is an integration that can connect to a Docker image registry via
// a specific auth mechanism
type Registry struct {
	gorm.Model

	// Name of the registry
	Name string `json:"name"`

	// URL of the registry
	URL string `json:"url"`

	// The project that this integration belongs to
	ProjectID uint `json:"project_id"`

	// The infra id, if registry was provisioned with Porter
	InfraID uint `json:"infra_id"`

	// ------------------------------------------------------------------
	// All fields below this line are encrypted before storage
	// ------------------------------------------------------------------

	GCPIntegrationID uint
	AWSIntegrationID uint
	DOIntegrationID  uint

	// A token cache that can be used by an auth mechanism (integration), if desired
	TokenCache integrations.RegTokenCache
}

// RegistryExternal is an external Registry to be shared over REST
type RegistryExternal struct {
	ID uint `json:"id"`

	// The project that this integration belongs to
	ProjectID uint `json:"project_id"`

	// Name of the registry
	Name string `json:"name"`

	// URL of the registry
	URL string `json:"url"`

	// The integration service for this registry
	Service integrations.IntegrationService `json:"service"`

	// The infra id, if registry was provisioned with Porter
	InfraID uint `json:"infra_id"`
}

// Externalize generates an external Registry to be shared over REST
func (r *Registry) Externalize() *RegistryExternal {
	var serv integrations.IntegrationService

	if r.AWSIntegrationID != 0 {
		serv = integrations.ECR
	} else if r.GCPIntegrationID != 0 {
		serv = integrations.GCR
	} else if r.DOIntegrationID != 0 {
		serv = integrations.DOCR
	}

	return &RegistryExternal{
		ID:        r.ID,
		ProjectID: r.ProjectID,
		Name:      r.Name,
		URL:       r.URL,
		Service:   serv,
		InfraID:   r.InfraID,
	}
}
