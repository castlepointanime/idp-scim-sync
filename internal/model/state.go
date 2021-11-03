package model

import (
	"encoding/json"

	"github.com/slashdevops/idp-scim-sync/internal/hash"
)

const (
	// StateSchemaVersion is the current schema version for the state file.
	StateSchemaVersion = "1.0.0"
)

// StateResources is a list of resources in the state, groups, users and groups and their users.
type StateResources struct {
	Groups        GroupsResult        `json:"groups"`
	Users         UsersResult         `json:"users"`
	GroupsMembers GroupsMembersResult `json:"groupsMembers"`
}

// State is the state of the system.
type State struct {
	SchemaVersion string         `json:"schemaVersion"`
	CodeVersion   string         `json:"codeVersion"`
	LastSync      string         `json:"lastSync"`
	HashCode      string         `json:"hashCode"`
	Resources     StateResources `json:"resources"`
}

// MarshalJSON marshals the state to JSON.
func (s *State) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(*s, "", "  ")
}

// SetHashCode is a helper function to avoid errors when calculating hash code.
// this method discards fields that are not used in the hash calculation.
// only fields comming from the Identity Provider are used.
func (s *State) SetHashCode() {
	s.HashCode = hash.Get(State{
		Resources: s.Resources,
	})
}
