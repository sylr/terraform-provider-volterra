//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package namespace

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GetActiveServicePoliciesRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetActiveServicePoliciesRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetActiveServicePoliciesRequest) DeepCopy() *GetActiveServicePoliciesRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetActiveServicePoliciesRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetActiveServicePoliciesRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetActiveServicePoliciesRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetActiveServicePoliciesRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetActiveServicePoliciesRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetActiveServicePoliciesRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetActiveServicePoliciesRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetActiveServicePoliciesRequest got type %s", t)
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
var DefaultGetActiveServicePoliciesRequestValidator = func() *ValidateGetActiveServicePoliciesRequest {
	v := &ValidateGetActiveServicePoliciesRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetActiveServicePoliciesRequestValidator() db.Validator {
	return DefaultGetActiveServicePoliciesRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetActiveServicePoliciesResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetActiveServicePoliciesResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetActiveServicePoliciesResponse) DeepCopy() *GetActiveServicePoliciesResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetActiveServicePoliciesResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetActiveServicePoliciesResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetActiveServicePoliciesResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetActiveServicePoliciesResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetActiveServicePoliciesResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetServicePoliciesDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *GetActiveServicePoliciesResponse) GetServicePoliciesDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, vref := range m.GetServicePolicies() {
		if vref == nil {
			return nil, fmt.Errorf("GetActiveServicePoliciesResponse.service_policies[%d] has a nil value", i)
		}
		vdRef := db.NewDirectRefForView(vref)
		vdRef.SetKind("service_policy.Object")
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "service_policy.Object",
			RefdTenant: vref.Tenant,
			RefdNS:     vref.Namespace,
			RefdName:   vref.Name,
			DRField:    "service_policies",
			Ref:        vdRef,
		})
	}

	return drInfos, nil
}

// GetServicePoliciesDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetActiveServicePoliciesResponse) GetServicePoliciesDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "service_policy.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: service_policy")
	}
	for i, vref := range m.GetServicePolicies() {
		if vref == nil {
			return nil, fmt.Errorf("GetActiveServicePoliciesResponse.service_policies[%d] has a nil value", i)
		}
		ref := &ves_io_schema.ObjectRefType{
			Kind:      "service_policy.Object",
			Tenant:    vref.Tenant,
			Namespace: vref.Namespace,
			Name:      vref.Name,
		}
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGetActiveServicePoliciesResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetActiveServicePoliciesResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetActiveServicePoliciesResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetActiveServicePoliciesResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["service_policies"]; exists {

		vOpts := append(opts, db.WithValidateField("service_policies"))
		for idx, item := range m.GetServicePolicies() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetActiveServicePoliciesResponseValidator = func() *ValidateGetActiveServicePoliciesResponse {
	v := &ValidateGetActiveServicePoliciesResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["service_policies"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	return v
}()

func GetActiveServicePoliciesResponseValidator() db.Validator {
	return DefaultGetActiveServicePoliciesResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetFastACLsForInternetVIPsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetFastACLsForInternetVIPsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetFastACLsForInternetVIPsRequest) DeepCopy() *GetFastACLsForInternetVIPsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetFastACLsForInternetVIPsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetFastACLsForInternetVIPsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetFastACLsForInternetVIPsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetFastACLsForInternetVIPsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetFastACLsForInternetVIPsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetFastACLsForInternetVIPsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetFastACLsForInternetVIPsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetFastACLsForInternetVIPsRequest got type %s", t)
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
var DefaultGetFastACLsForInternetVIPsRequestValidator = func() *ValidateGetFastACLsForInternetVIPsRequest {
	v := &ValidateGetFastACLsForInternetVIPsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetFastACLsForInternetVIPsRequestValidator() db.Validator {
	return DefaultGetFastACLsForInternetVIPsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetFastACLsForInternetVIPsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetFastACLsForInternetVIPsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetFastACLsForInternetVIPsResponse) DeepCopy() *GetFastACLsForInternetVIPsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetFastACLsForInternetVIPsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetFastACLsForInternetVIPsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetFastACLsForInternetVIPsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetFastACLsForInternetVIPsResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetFastACLsForInternetVIPsResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetFastAclsDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *GetFastACLsForInternetVIPsResponse) GetFastAclsDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, vref := range m.GetFastAcls() {
		if vref == nil {
			return nil, fmt.Errorf("GetFastACLsForInternetVIPsResponse.fast_acls[%d] has a nil value", i)
		}
		vdRef := db.NewDirectRefForView(vref)
		vdRef.SetKind("fast_acl.Object")
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "fast_acl.Object",
			RefdTenant: vref.Tenant,
			RefdNS:     vref.Namespace,
			RefdName:   vref.Name,
			DRField:    "fast_acls",
			Ref:        vdRef,
		})
	}

	return drInfos, nil
}

// GetFastAclsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetFastACLsForInternetVIPsResponse) GetFastAclsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "fast_acl.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: fast_acl")
	}
	for i, vref := range m.GetFastAcls() {
		if vref == nil {
			return nil, fmt.Errorf("GetFastACLsForInternetVIPsResponse.fast_acls[%d] has a nil value", i)
		}
		ref := &ves_io_schema.ObjectRefType{
			Kind:      "fast_acl.Object",
			Tenant:    vref.Tenant,
			Namespace: vref.Namespace,
			Name:      vref.Name,
		}
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGetFastACLsForInternetVIPsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetFastACLsForInternetVIPsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetFastACLsForInternetVIPsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetFastACLsForInternetVIPsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["fast_acls"]; exists {

		vOpts := append(opts, db.WithValidateField("fast_acls"))
		for idx, item := range m.GetFastAcls() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetFastACLsForInternetVIPsResponseValidator = func() *ValidateGetFastACLsForInternetVIPsResponse {
	v := &ValidateGetFastACLsForInternetVIPsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["fast_acls"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	return v
}()

func GetFastACLsForInternetVIPsResponseValidator() db.Validator {
	return DefaultGetFastACLsForInternetVIPsResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SetActiveServicePoliciesRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetActiveServicePoliciesRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetActiveServicePoliciesRequest) DeepCopy() *SetActiveServicePoliciesRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetActiveServicePoliciesRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetActiveServicePoliciesRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetActiveServicePoliciesRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetActiveServicePoliciesRequestValidator().Validate(ctx, m, opts...)
}

func (m *SetActiveServicePoliciesRequest) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetServicePoliciesDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *SetActiveServicePoliciesRequest) GetServicePoliciesDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, vref := range m.GetServicePolicies() {
		if vref == nil {
			return nil, fmt.Errorf("SetActiveServicePoliciesRequest.service_policies[%d] has a nil value", i)
		}
		vdRef := db.NewDirectRefForView(vref)
		vdRef.SetKind("service_policy.Object")
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "service_policy.Object",
			RefdTenant: vref.Tenant,
			RefdNS:     vref.Namespace,
			RefdName:   vref.Name,
			DRField:    "service_policies",
			Ref:        vdRef,
		})
	}

	return drInfos, nil
}

// GetServicePoliciesDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *SetActiveServicePoliciesRequest) GetServicePoliciesDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "service_policy.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: service_policy")
	}
	for i, vref := range m.GetServicePolicies() {
		if vref == nil {
			return nil, fmt.Errorf("SetActiveServicePoliciesRequest.service_policies[%d] has a nil value", i)
		}
		ref := &ves_io_schema.ObjectRefType{
			Kind:      "service_policy.Object",
			Tenant:    vref.Tenant,
			Namespace: vref.Namespace,
			Name:      vref.Name,
		}
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateSetActiveServicePoliciesRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetActiveServicePoliciesRequest) ServicePoliciesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_views.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ves_io_schema_views.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for service_policies")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_views.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_views.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated service_policies")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items service_policies")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateSetActiveServicePoliciesRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetActiveServicePoliciesRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetActiveServicePoliciesRequest got type %s", t)
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

	if fv, exists := v.FldValidators["service_policies"]; exists {
		vOpts := append(opts, db.WithValidateField("service_policies"))
		if err := fv(ctx, m.GetServicePolicies(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetActiveServicePoliciesRequestValidator = func() *ValidateSetActiveServicePoliciesRequest {
	v := &ValidateSetActiveServicePoliciesRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhServicePolicies := v.ServicePoliciesValidationRuleHandler
	rulesServicePolicies := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "32",
	}
	vFn, err = vrhServicePolicies(rulesServicePolicies)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SetActiveServicePoliciesRequest.service_policies: %s", err)
		panic(errMsg)
	}
	v.FldValidators["service_policies"] = vFn

	return v
}()

func SetActiveServicePoliciesRequestValidator() db.Validator {
	return DefaultSetActiveServicePoliciesRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetActiveServicePoliciesResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetActiveServicePoliciesResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetActiveServicePoliciesResponse) DeepCopy() *SetActiveServicePoliciesResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetActiveServicePoliciesResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetActiveServicePoliciesResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetActiveServicePoliciesResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetActiveServicePoliciesResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetActiveServicePoliciesResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetActiveServicePoliciesResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetActiveServicePoliciesResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetActiveServicePoliciesResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetActiveServicePoliciesResponseValidator = func() *ValidateSetActiveServicePoliciesResponse {
	v := &ValidateSetActiveServicePoliciesResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetActiveServicePoliciesResponseValidator() db.Validator {
	return DefaultSetActiveServicePoliciesResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SetFastACLsForInternetVIPsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetFastACLsForInternetVIPsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetFastACLsForInternetVIPsRequest) DeepCopy() *SetFastACLsForInternetVIPsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetFastACLsForInternetVIPsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetFastACLsForInternetVIPsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetFastACLsForInternetVIPsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetFastACLsForInternetVIPsRequestValidator().Validate(ctx, m, opts...)
}

func (m *SetFastACLsForInternetVIPsRequest) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetFastAclsDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *SetFastACLsForInternetVIPsRequest) GetFastAclsDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, vref := range m.GetFastAcls() {
		if vref == nil {
			return nil, fmt.Errorf("SetFastACLsForInternetVIPsRequest.fast_acls[%d] has a nil value", i)
		}
		vdRef := db.NewDirectRefForView(vref)
		vdRef.SetKind("fast_acl.Object")
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "fast_acl.Object",
			RefdTenant: vref.Tenant,
			RefdNS:     vref.Namespace,
			RefdName:   vref.Name,
			DRField:    "fast_acls",
			Ref:        vdRef,
		})
	}

	return drInfos, nil
}

// GetFastAclsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *SetFastACLsForInternetVIPsRequest) GetFastAclsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "fast_acl.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: fast_acl")
	}
	for i, vref := range m.GetFastAcls() {
		if vref == nil {
			return nil, fmt.Errorf("SetFastACLsForInternetVIPsRequest.fast_acls[%d] has a nil value", i)
		}
		ref := &ves_io_schema.ObjectRefType{
			Kind:      "fast_acl.Object",
			Tenant:    vref.Tenant,
			Namespace: vref.Namespace,
			Name:      vref.Name,
		}
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateSetFastACLsForInternetVIPsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetFastACLsForInternetVIPsRequest) FastAclsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_views.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ves_io_schema_views.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for fast_acls")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_views.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_views.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated fast_acls")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items fast_acls")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateSetFastACLsForInternetVIPsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetFastACLsForInternetVIPsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetFastACLsForInternetVIPsRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["fast_acls"]; exists {
		vOpts := append(opts, db.WithValidateField("fast_acls"))
		if err := fv(ctx, m.GetFastAcls(), vOpts...); err != nil {
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
var DefaultSetFastACLsForInternetVIPsRequestValidator = func() *ValidateSetFastACLsForInternetVIPsRequest {
	v := &ValidateSetFastACLsForInternetVIPsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhFastAcls := v.FastAclsValidationRuleHandler
	rulesFastAcls := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "128",
	}
	vFn, err = vrhFastAcls(rulesFastAcls)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SetFastACLsForInternetVIPsRequest.fast_acls: %s", err)
		panic(errMsg)
	}
	v.FldValidators["fast_acls"] = vFn

	return v
}()

func SetFastACLsForInternetVIPsRequestValidator() db.Validator {
	return DefaultSetFastACLsForInternetVIPsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetFastACLsForInternetVIPsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetFastACLsForInternetVIPsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetFastACLsForInternetVIPsResponse) DeepCopy() *SetFastACLsForInternetVIPsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetFastACLsForInternetVIPsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetFastACLsForInternetVIPsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetFastACLsForInternetVIPsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetFastACLsForInternetVIPsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetFastACLsForInternetVIPsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetFastACLsForInternetVIPsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetFastACLsForInternetVIPsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetFastACLsForInternetVIPsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetFastACLsForInternetVIPsResponseValidator = func() *ValidateSetFastACLsForInternetVIPsResponse {
	v := &ValidateSetFastACLsForInternetVIPsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetFastACLsForInternetVIPsResponseValidator() db.Validator {
	return DefaultSetFastACLsForInternetVIPsResponseValidator
}