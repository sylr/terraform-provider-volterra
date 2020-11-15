//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CustomCreateRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomCreateRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomCreateRequest) DeepCopy() *CustomCreateRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomCreateRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomCreateRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomCreateRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomCreateRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCustomCreateRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomCreateRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomCreateRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomCreateRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_groups"]; exists {

		vOpts := append(opts, db.WithValidateField("api_groups"))
		for idx, item := range m.GetApiGroups() {
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

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
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
var DefaultCustomCreateRequestValidator = func() *ValidateCustomCreateRequest {
	v := &ValidateCustomCreateRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectCreateMetaTypeValidator().Validate

	v.FldValidators["spec"] = CreateSpecTypeValidator().Validate

	return v
}()

func CustomCreateRequestValidator() db.Validator {
	return DefaultCustomCreateRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CustomGetRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomGetRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomGetRequest) DeepCopy() *CustomGetRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomGetRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomGetRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomGetRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomGetRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCustomGetRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomGetRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomGetRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomGetRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCustomGetRequestValidator = func() *ValidateCustomGetRequest {
	v := &ValidateCustomGetRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CustomGetRequestValidator() db.Validator {
	return DefaultCustomGetRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CustomGetResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomGetResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomGetResponse) DeepCopy() *CustomGetResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomGetResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomGetResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomGetResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomGetResponseValidator().Validate(ctx, m, opts...)
}

func (m *CustomGetResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo

	return drInfos, nil
}

type ValidateCustomGetResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomGetResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomGetResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomGetResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_groups"]; exists {

		vOpts := append(opts, db.WithValidateField("api_groups"))
		for idx, item := range m.GetApiGroups() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["object"]; exists {

		vOpts := append(opts, db.WithValidateField("object"))
		if err := fv(ctx, m.GetObject(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCustomGetResponseValidator = func() *ValidateCustomGetResponse {
	v := &ValidateCustomGetResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["object"] = ObjectValidator().Validate

	return v
}()

func CustomGetResponseValidator() db.Validator {
	return DefaultCustomGetResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *CustomListRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomListRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomListRequest) DeepCopy() *CustomListRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomListRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomListRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomListRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomListRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCustomListRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomListRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomListRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomListRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCustomListRequestValidator = func() *ValidateCustomListRequest {
	v := &ValidateCustomListRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CustomListRequestValidator() db.Validator {
	return DefaultCustomListRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CustomListResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomListResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomListResponse) DeepCopy() *CustomListResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomListResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomListResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomListResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomListResponseValidator().Validate(ctx, m, opts...)
}

func (m *CustomListResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetItemsDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *CustomListResponse) GetItemsDRefInfo() ([]db.DRefInfo, error) {
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

type ValidateCustomListResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomListResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomListResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomListResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCustomListResponseValidator = func() *ValidateCustomListResponse {
	v := &ValidateCustomListResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["items"] = RoleValidator().Validate

	return v
}()

func CustomListResponseValidator() db.Validator {
	return DefaultCustomListResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *CustomReplaceRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomReplaceRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomReplaceRequest) DeepCopy() *CustomReplaceRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomReplaceRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomReplaceRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomReplaceRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomReplaceRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCustomReplaceRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomReplaceRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomReplaceRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomReplaceRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_groups"]; exists {

		vOpts := append(opts, db.WithValidateField("api_groups"))
		for idx, item := range m.GetApiGroups() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
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

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCustomReplaceRequestValidator = func() *ValidateCustomReplaceRequest {
	v := &ValidateCustomReplaceRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["spec"] = ReplaceSpecTypeValidator().Validate

	return v
}()

func CustomReplaceRequestValidator() db.Validator {
	return DefaultCustomReplaceRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *Role) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Role) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Role) DeepCopy() *Role {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Role{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Role) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Role) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RoleValidator().Validate(ctx, m, opts...)
}

func (m *Role) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetGetSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *Role) GetGetSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.GetSpec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.GetSpec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "get_spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateRole struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRole) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Role)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Role got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_groups"]; exists {

		vOpts := append(opts, db.WithValidateField("api_groups"))
		for idx, item := range m.GetApiGroups() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["disabled"]; exists {

		vOpts := append(opts, db.WithValidateField("disabled"))
		if err := fv(ctx, m.GetDisabled(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["get_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("get_spec"))
		if err := fv(ctx, m.GetGetSpec(), vOpts...); err != nil {
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

	if fv, exists := v.FldValidators["tenant"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant"))
		if err := fv(ctx, m.GetTenant(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["uid"]; exists {

		vOpts := append(opts, db.WithValidateField("uid"))
		if err := fv(ctx, m.GetUid(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRoleValidator = func() *ValidateRole {
	v := &ValidateRole{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["get_spec"] = GetSpecTypeValidator().Validate

	return v
}()

func RoleValidator() db.Validator {
	return DefaultRoleValidator
}