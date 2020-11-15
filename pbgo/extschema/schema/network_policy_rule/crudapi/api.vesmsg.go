//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package crudapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_network_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy_rule"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *ObjectCreateReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectCreateReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectCreateReq) DeepCopy() *ObjectCreateReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectCreateReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectCreateReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectCreateReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectCreateReqValidator().Validate(ctx, m, opts...)
}

func (m *ObjectCreateReq) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetSystemMetadataDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *ObjectCreateReq) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.Spec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.Spec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

// GetDRefInfo for the field's type
func (m *ObjectCreateReq) GetSystemMetadataDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.SystemMetadata == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.SystemMetadata.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "system_metadata." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateObjectCreateReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectCreateReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectCreateReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectCreateReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectCreateReqValidator = func() *ValidateObjectCreateReq {
	v := &ValidateObjectCreateReq{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["system_metadata"] = ves_io_schema.SystemObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = ves_io_schema_network_policy_rule.SpecTypeValidator().Validate

	return v
}()

func ObjectCreateReqValidator() db.Validator {
	return DefaultObjectCreateReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectCreateRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectCreateRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectCreateRsp) DeepCopy() *ObjectCreateRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectCreateRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectCreateRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectCreateRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectCreateRspValidator().Validate(ctx, m, opts...)
}

func (m *ObjectCreateRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetSystemMetadataDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *ObjectCreateRsp) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.Spec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.Spec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

// GetDRefInfo for the field's type
func (m *ObjectCreateRsp) GetSystemMetadataDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.SystemMetadata == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.SystemMetadata.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "system_metadata." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateObjectCreateRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectCreateRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectCreateRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectCreateRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["err"]; exists {

		vOpts := append(opts, db.WithValidateField("err"))
		if err := fv(ctx, m.GetErr(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_uid"]; exists {

		vOpts := append(opts, db.WithValidateField("object_uid"))
		if err := fv(ctx, m.GetObjectUid(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectCreateRspValidator = func() *ValidateObjectCreateRsp {
	v := &ValidateObjectCreateRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["system_metadata"] = ves_io_schema.SystemObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = ves_io_schema_network_policy_rule.SpecTypeValidator().Validate

	return v
}()

func ObjectCreateRspValidator() db.Validator {
	return DefaultObjectCreateRspValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectDeleteReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectDeleteReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectDeleteReq) DeepCopy() *ObjectDeleteReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectDeleteReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectDeleteReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectDeleteReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectDeleteReqValidator().Validate(ctx, m, opts...)
}

type ValidateObjectDeleteReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectDeleteReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectDeleteReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectDeleteReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["object_uid"]; exists {

		vOpts := append(opts, db.WithValidateField("object_uid"))
		if err := fv(ctx, m.GetObjectUid(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectDeleteReqValidator = func() *ValidateObjectDeleteReq {
	v := &ValidateObjectDeleteReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ObjectDeleteReqValidator() db.Validator {
	return DefaultObjectDeleteReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectDeleteRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectDeleteRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectDeleteRsp) DeepCopy() *ObjectDeleteRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectDeleteRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectDeleteRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectDeleteRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectDeleteRspValidator().Validate(ctx, m, opts...)
}

type ValidateObjectDeleteRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectDeleteRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectDeleteRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectDeleteRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["err"]; exists {

		vOpts := append(opts, db.WithValidateField("err"))
		if err := fv(ctx, m.GetErr(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectDeleteRspValidator = func() *ValidateObjectDeleteRsp {
	v := &ValidateObjectDeleteRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ObjectDeleteRspValidator() db.Validator {
	return DefaultObjectDeleteRspValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectGetReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectGetReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectGetReq) DeepCopy() *ObjectGetReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectGetReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectGetReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectGetReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectGetReqValidator().Validate(ctx, m, opts...)
}

type ValidateObjectGetReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectGetReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectGetReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectGetReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["all_backrefs"]; exists {

		vOpts := append(opts, db.WithValidateField("all_backrefs"))
		if err := fv(ctx, m.GetAllBackrefs(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["backref_types"]; exists {

		vOpts := append(opts, db.WithValidateField("backref_types"))
		for idx, item := range m.GetBackrefTypes() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["include_referred_id"]; exists {

		vOpts := append(opts, db.WithValidateField("include_referred_id"))
		if err := fv(ctx, m.GetIncludeReferredId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_uid"]; exists {

		vOpts := append(opts, db.WithValidateField("object_uid"))
		if err := fv(ctx, m.GetObjectUid(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectGetReqValidator = func() *ValidateObjectGetReq {
	v := &ValidateObjectGetReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ObjectGetReqValidator() db.Validator {
	return DefaultObjectGetReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectGetRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectGetRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectGetRsp) DeepCopy() *ObjectGetRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectGetRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectGetRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectGetRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectGetRspValidator().Validate(ctx, m, opts...)
}

func (m *ObjectGetRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetSystemMetadataDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *ObjectGetRsp) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.Spec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.Spec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

// GetDRefInfo for the field's type
func (m *ObjectGetRsp) GetSystemMetadataDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.SystemMetadata == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.SystemMetadata.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "system_metadata." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateObjectGetRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectGetRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectGetRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectGetRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ent_backrefs"]; exists {

		vOpts := append(opts, db.WithValidateField("ent_backrefs"))
		for idx, item := range m.GetEntBackrefs() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["err"]; exists {

		vOpts := append(opts, db.WithValidateField("err"))
		if err := fv(ctx, m.GetErr(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["resource_version"]; exists {

		vOpts := append(opts, db.WithValidateField("resource_version"))
		if err := fv(ctx, m.GetResourceVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectGetRspValidator = func() *ValidateObjectGetRsp {
	v := &ValidateObjectGetRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["system_metadata"] = ves_io_schema.SystemObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = ves_io_schema_network_policy_rule.SpecTypeValidator().Validate

	v.FldValidators["status"] = ves_io_schema_network_policy_rule.StatusObjectValidator().Validate

	return v
}()

func ObjectGetRspValidator() db.Validator {
	return DefaultObjectGetRspValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectListReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectListReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectListReq) DeepCopy() *ObjectListReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectListReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectListReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectListReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectListReqValidator().Validate(ctx, m, opts...)
}

type ValidateObjectListReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectListReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectListReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectListReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["include_referred_id"]; exists {

		vOpts := append(opts, db.WithValidateField("include_referred_id"))
		if err := fv(ctx, m.GetIncludeReferredId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["label_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("label_filter"))
		if err := fv(ctx, m.GetLabelFilter(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace_filter"))
		for idx, item := range m.GetNamespaceFilter() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["report_fields"]; exists {

		vOpts := append(opts, db.WithValidateField("report_fields"))
		for idx, item := range m.GetReportFields() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["resource_version"]; exists {

		vOpts := append(opts, db.WithValidateField("resource_version"))
		if err := fv(ctx, m.GetResourceVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant_filter"))
		for idx, item := range m.GetTenantFilter() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectListReqValidator = func() *ValidateObjectListReq {
	v := &ValidateObjectListReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ObjectListReqValidator() db.Validator {
	return DefaultObjectListReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectListRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectListRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectListRsp) DeepCopy() *ObjectListRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectListRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectListRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectListRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectListRspValidator().Validate(ctx, m, opts...)
}

func (m *ObjectListRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetItemsDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *ObjectListRsp) GetItemsDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.Items == nil {
		return []db.DRefInfo{}, nil
	}

	for idx, e := range m.Items {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, err
		}
		for _, dri := range driSet {
			dri.DRField = fmt.Sprintf("items[%v].%s", idx, dri.DRField)
			drInfos = append(drInfos, dri)
		}
	}

	return drInfos, err
}

type ValidateObjectListRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectListRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectListRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectListRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["err"]; exists {

		vOpts := append(opts, db.WithValidateField("err"))
		if err := fv(ctx, m.GetErr(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["uids"]; exists {

		vOpts := append(opts, db.WithValidateField("uids"))
		for idx, item := range m.GetUids() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectListRspValidator = func() *ValidateObjectListRsp {
	v := &ValidateObjectListRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["items"] = ObjectListRspItemValidator().Validate

	return v
}()

func ObjectListRspValidator() db.Validator {
	return DefaultObjectListRspValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectListRspItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectListRspItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectListRspItem) DeepCopy() *ObjectListRspItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectListRspItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectListRspItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectListRspItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectListRspItemValidator().Validate(ctx, m, opts...)
}

func (m *ObjectListRspItem) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetSystemMetadataDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *ObjectListRspItem) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.Spec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.Spec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

// GetDRefInfo for the field's type
func (m *ObjectListRspItem) GetSystemMetadataDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.SystemMetadata == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.SystemMetadata.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "system_metadata." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateObjectListRspItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectListRspItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectListRspItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectListRspItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["labels"]; exists {

		vOpts := append(opts, db.WithValidateField("labels"))
		for key, value := range m.GetLabels() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_uid"]; exists {

		vOpts := append(opts, db.WithValidateField("object_uid"))
		if err := fv(ctx, m.GetObjectUid(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant"))
		if err := fv(ctx, m.GetTenant(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectListRspItemValidator = func() *ValidateObjectListRspItem {
	v := &ValidateObjectListRspItem{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["system_metadata"] = ves_io_schema.SystemObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = ves_io_schema_network_policy_rule.SpecTypeValidator().Validate

	v.FldValidators["status"] = ves_io_schema_network_policy_rule.StatusObjectValidator().Validate

	return v
}()

func ObjectListRspItemValidator() db.Validator {
	return DefaultObjectListRspItemValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectReplaceReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectReplaceReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectReplaceReq) DeepCopy() *ObjectReplaceReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectReplaceReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectReplaceReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectReplaceReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectReplaceReqValidator().Validate(ctx, m, opts...)
}

func (m *ObjectReplaceReq) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *ObjectReplaceReq) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.Spec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.Spec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateObjectReplaceReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectReplaceReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectReplaceReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectReplaceReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_uid"]; exists {

		vOpts := append(opts, db.WithValidateField("object_uid"))
		if err := fv(ctx, m.GetObjectUid(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["resource_version"]; exists {

		vOpts := append(opts, db.WithValidateField("resource_version"))
		if err := fv(ctx, m.GetResourceVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectReplaceReqValidator = func() *ValidateObjectReplaceReq {
	v := &ValidateObjectReplaceReq{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = ves_io_schema_network_policy_rule.SpecTypeValidator().Validate

	return v
}()

func ObjectReplaceReqValidator() db.Validator {
	return DefaultObjectReplaceReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectReplaceRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectReplaceRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectReplaceRsp) DeepCopy() *ObjectReplaceRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectReplaceRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectReplaceRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectReplaceRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectReplaceRspValidator().Validate(ctx, m, opts...)
}

func (m *ObjectReplaceRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetSystemMetadataDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *ObjectReplaceRsp) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.Spec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.Spec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

// GetDRefInfo for the field's type
func (m *ObjectReplaceRsp) GetSystemMetadataDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.SystemMetadata == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.SystemMetadata.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "system_metadata." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateObjectReplaceRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectReplaceRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectReplaceRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectReplaceRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["err"]; exists {

		vOpts := append(opts, db.WithValidateField("err"))
		if err := fv(ctx, m.GetErr(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectReplaceRspValidator = func() *ValidateObjectReplaceRsp {
	v := &ValidateObjectReplaceRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["system_metadata"] = ves_io_schema.SystemObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = ves_io_schema_network_policy_rule.SpecTypeValidator().Validate

	return v
}()

func ObjectReplaceRspValidator() db.Validator {
	return DefaultObjectReplaceRspValidator
}

func (m *ObjectCreateReq) FromObject(e db.Entry) {
	f := e.DeepCopy().(*ves_io_schema_network_policy_rule.DBObject)
	_ = f
	m.Metadata = f.GetMetadata()
	m.Spec = f.GetSpec()
	m.SystemMetadata = f.GetSystemMetadata()
}

func (m *ObjectCreateReq) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*ves_io_schema_network_policy_rule.DBObject)
	_ = f
	f.Metadata = m1.Metadata
	f.Spec = m1.Spec
	f.SystemMetadata = m1.SystemMetadata
}

func (m *ObjectCreateRsp) FromObject(e db.Entry) {
	f := e.DeepCopy().(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	m.Metadata = f.GetMetadata()

	m.Spec = f.GetSpec()
	m.SystemMetadata = f.GetSystemMetadata()
}

func (m *ObjectCreateRsp) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	f.Metadata = m1.Metadata

	f.Spec = m1.Spec
	f.SystemMetadata = m1.SystemMetadata
}

func (m *ObjectGetRsp) FromObject(e db.Entry) {
	f := e.DeepCopy().(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	m.Metadata = f.GetMetadata()

	m.Spec = f.GetSpec()

	m.SystemMetadata = f.GetSystemMetadata()
}

func (m *ObjectGetRsp) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	f.Metadata = m1.Metadata

	f.Spec = m1.Spec

	f.SystemMetadata = m1.SystemMetadata
}

func (m *ObjectListRspItem) FromObject(e db.Entry) {
	f := e.DeepCopy().(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	m.Metadata = f.GetMetadata()

	m.Spec = f.GetSpec()

	m.SystemMetadata = f.GetSystemMetadata()

}

func (m *ObjectListRspItem) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	f.Metadata = m1.Metadata

	f.Spec = m1.Spec

	f.SystemMetadata = m1.SystemMetadata

}

func (m *ObjectReplaceReq) FromObject(e db.Entry) {
	f := e.DeepCopy().(*ves_io_schema_network_policy_rule.DBObject)
	_ = f
	m.Metadata = f.GetMetadata()

	m.Spec = f.GetSpec()
}

func (m *ObjectReplaceReq) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*ves_io_schema_network_policy_rule.DBObject)
	_ = f
	f.Metadata = m1.Metadata

	f.Spec = m1.Spec
}

func (m *ObjectReplaceRsp) FromObject(e db.Entry) {
	f := e.DeepCopy().(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	m.Metadata = f.GetMetadata()
	m.Spec = f.GetSpec()
	m.SystemMetadata = f.GetSystemMetadata()
}

func (m *ObjectReplaceRsp) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*ves_io_schema_network_policy_rule.DBObject)
	_ = f

	f.Metadata = m1.Metadata
	f.Spec = m1.Spec
	f.SystemMetadata = m1.SystemMetadata
}
