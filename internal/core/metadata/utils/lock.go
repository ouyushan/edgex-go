//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import "sync"

type CapacityCheckLock struct {
	mutex sync.RWMutex
}

func NewCapacityCheckLock() *CapacityCheckLock {
	return &CapacityCheckLock{}
}

func (c *CapacityCheckLock) Lock() {
	c.mutex.Lock()
}
func (c *CapacityCheckLock) Unlock() {
	c.mutex.Unlock()
}

// ProfileAssignmentLock guards the assignment of a device profile against its deletion: a device
// being written is not in the database yet, so the deletion check cannot see it referencing the profile.
type ProfileAssignmentLock struct {
	mutex sync.RWMutex
}

func NewProfileAssignmentLock() *ProfileAssignmentLock {
	return &ProfileAssignmentLock{}
}

// RLock is held by the device writes assigning a profile, which may run concurrently with each other.
func (d *ProfileAssignmentLock) RLock() {
	d.mutex.RLock()
}
func (d *ProfileAssignmentLock) RUnlock() {
	d.mutex.RUnlock()
}

// Lock is held by the device profile deletion, which waits for the in-flight device writes.
func (d *ProfileAssignmentLock) Lock() {
	d.mutex.Lock()
}
func (d *ProfileAssignmentLock) Unlock() {
	d.mutex.Unlock()
}
