// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import "errors"

const (
	Read Permissions = 1 << iota
	// TODO: Handle Allocate permission
	Allocate
	Write

	None Permissions = 0
	All              = Read | Allocate | Write
)

var errUnexpectedLength = errors.New("permissions should be exactly one byte")

// StateKey holds the name of the key and its permission (Read/Allocate/Write). By default,
// initialization of Keys with duplicate key will not work. And to prevent duplicate
// insertions from overriding the original permissions, use the Add function below.
type Keys map[string]Permissions

// All acceptable permission options
type Permissions byte

// Transactions are expected to use this to prevent
// overriding of key permissions
func (k Keys) Add(name string, permission Permissions) {
	// Transaction's permissions are the union of all
	// state keys from both Actions and Auth
	k[name] |= permission
}

// Has returns true if [p] has all the permissions that are contained in require
func (p Permissions) Has(require Permissions) bool {
	return require&^p == 0
}

// Populates bitset using the values from Permissions struct
// We don't need to worry about adding invalid permission bits
// since our struct is defined with all valid ones
func (p Permissions) ToBytes() []byte {
	return []byte{byte(p)}
}

// Populates Permissions struct with the permissions bytes
// Any permission bits that were not defined will be dropped
func FromBytes(permissionBytes []byte) (Permissions, error) {
	if len(permissionBytes) != 1 {
		return None, errUnexpectedLength
	}
	return Permissions(permissionBytes[0]) & All, nil
}