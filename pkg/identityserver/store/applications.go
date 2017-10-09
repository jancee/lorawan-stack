// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package store

import (
	"github.com/TheThingsNetwork/ttn/pkg/identityserver/store/sql/factory"
	"github.com/TheThingsNetwork/ttn/pkg/identityserver/types"
	"github.com/TheThingsNetwork/ttn/pkg/ttnpb"
)

// ApplicationStore is a store that holds Applications.
type ApplicationStore interface {
	// Create creates a new application.
	Create(app types.Application) error

	// GetByID finds the application by ID and retrieves it.
	GetByID(appID string) (types.Application, error)

	// ListByUser returns the applications to which an user is a collaborator.
	ListByUser(userID string) ([]types.Application, error)

	// Update updates the application.
	Update(app types.Application) error

	// Archive sets the ArchivedAt field of the application to the current timestamp.
	Archive(appID string) error

	// AddAPIKey adds a new application API key to a given application.
	AddAPIKey(appID string, key ttnpb.APIKey) error

	// RemoveAPIKey removes an application API key from a given application.
	RemoveAPIKey(appID string, keyName string) error

	// SetCollaborator inserts or updates a collaborator within an application.
	// If the list of rights is empty the collaborator will be unset.
	SetCollaborator(appID string, collaborator ttnpb.Collaborator) error

	// ListCollaborators retrieves all the collaborators from an application.
	ListCollaborators(appID string) ([]ttnpb.Collaborator, error)

	// ListUserRights returns the rights a given User has for an Application.
	ListUserRights(appID string, userID string) ([]ttnpb.Right, error)

	// LoadAttributes loads extra attributes into the Application.
	LoadAttributes(app types.Application) error

	// WriteAttributes writes the extra attributes on the Application if it is an
	// Attributer to the store.
	WriteAttributes(app types.Application, result types.Application) error

	// SetFactory allows to replace the default ttnpb.Application factory.
	SetFactory(factory factory.ApplicationFactory)
}
