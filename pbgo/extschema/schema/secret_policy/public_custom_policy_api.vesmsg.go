//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package secret_policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *RecoverRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RecoverRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RecoverRequest) DeepCopy() *RecoverRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RecoverRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RecoverRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RecoverRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RecoverRequestValidator().Validate(ctx, m, opts...)
}

type ValidateRecoverRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRecoverRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateRecoverRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateRecoverRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RecoverRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RecoverRequest got type %s", t)
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
var DefaultRecoverRequestValidator = func() *ValidateRecoverRequest {
	v := &ValidateRecoverRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RecoverRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RecoverRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	return v
}()

func RecoverRequestValidator() db.Validator {
	return DefaultRecoverRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *RecoverResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RecoverResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RecoverResponse) DeepCopy() *RecoverResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RecoverResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RecoverResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RecoverResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RecoverResponseValidator().Validate(ctx, m, opts...)
}

type ValidateRecoverResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRecoverResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RecoverResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RecoverResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRecoverResponseValidator = func() *ValidateRecoverResponse {
	v := &ValidateRecoverResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RecoverResponseValidator() db.Validator {
	return DefaultRecoverResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SoftDeleteRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SoftDeleteRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SoftDeleteRequest) DeepCopy() *SoftDeleteRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SoftDeleteRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SoftDeleteRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SoftDeleteRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SoftDeleteRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSoftDeleteRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSoftDeleteRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateSoftDeleteRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateSoftDeleteRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SoftDeleteRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SoftDeleteRequest got type %s", t)
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
var DefaultSoftDeleteRequestValidator = func() *ValidateSoftDeleteRequest {
	v := &ValidateSoftDeleteRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SoftDeleteRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SoftDeleteRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	return v
}()

func SoftDeleteRequestValidator() db.Validator {
	return DefaultSoftDeleteRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SoftDeleteResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SoftDeleteResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SoftDeleteResponse) DeepCopy() *SoftDeleteResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SoftDeleteResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SoftDeleteResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SoftDeleteResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SoftDeleteResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSoftDeleteResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSoftDeleteResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SoftDeleteResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SoftDeleteResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSoftDeleteResponseValidator = func() *ValidateSoftDeleteResponse {
	v := &ValidateSoftDeleteResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SoftDeleteResponseValidator() db.Validator {
	return DefaultSoftDeleteResponseValidator
}
