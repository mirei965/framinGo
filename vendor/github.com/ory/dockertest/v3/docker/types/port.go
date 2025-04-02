// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Port An open port on a container
// swagger:model Port
type Port struct {

	// IP
	IP string `json:"IP,omitempty"`

	// Port on the container
	// Required: true
	PrivatePort uint16 `json:"PrivatePort"`

	// Port exposed on the host
	PublicPort uint16 `json:"PublicPort,omitempty"`

	// type
	// Required: true
	Type string `json:"Type"`
}
