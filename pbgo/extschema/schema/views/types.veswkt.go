//
// Copyright (c) 2019 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package views

import (
	"gopkg.volterra.us/stdlib/db"
)

// DeepCopy2 allows ObjectRefType to satisfy ViewDirectRef
func (m *ObjectRefType) DeepCopy2() db.ViewDirectRef {
	return m.DeepCopy()
}
