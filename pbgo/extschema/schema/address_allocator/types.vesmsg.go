//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package address_allocator

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

func (m *AllocationScheme) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AllocationScheme) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AllocationScheme) DeepCopy() *AllocationScheme {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AllocationScheme{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AllocationScheme) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AllocationScheme) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AllocationSchemeValidator().Validate(ctx, m, opts...)
}

type ValidateAllocationScheme struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAllocationScheme) AllocationUnitValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for allocation_unit")
	}

	return validatorFn, nil
}

func (v *ValidateAllocationScheme) LocalInterfaceAddressOffsetValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for local_interface_address_offset")
	}

	return validatorFn, nil
}

func (v *ValidateAllocationScheme) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AllocationScheme)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AllocationScheme got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["allocation_unit"]; exists {

		vOpts := append(opts, db.WithValidateField("allocation_unit"))
		if err := fv(ctx, m.GetAllocationUnit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["local_interface_address_offset"]; exists {

		vOpts := append(opts, db.WithValidateField("local_interface_address_offset"))
		if err := fv(ctx, m.GetLocalInterfaceAddressOffset(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["local_interface_address_type"]; exists {

		vOpts := append(opts, db.WithValidateField("local_interface_address_type"))
		if err := fv(ctx, m.GetLocalInterfaceAddressType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAllocationSchemeValidator = func() *ValidateAllocationScheme {
	v := &ValidateAllocationScheme{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAllocationUnit := v.AllocationUnitValidationRuleHandler
	rulesAllocationUnit := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gte":       "0",
		"ves.io.schema.rules.uint32.lte":       "32",
	}
	vFn, err = vrhAllocationUnit(rulesAllocationUnit)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AllocationScheme.allocation_unit: %s", err)
		panic(errMsg)
	}
	v.FldValidators["allocation_unit"] = vFn

	vrhLocalInterfaceAddressOffset := v.LocalInterfaceAddressOffsetValidationRuleHandler
	rulesLocalInterfaceAddressOffset := map[string]string{
		"ves.io.schema.rules.uint32.gte": "0",
		"ves.io.schema.rules.uint32.lte": "32",
	}
	vFn, err = vrhLocalInterfaceAddressOffset(rulesLocalInterfaceAddressOffset)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AllocationScheme.local_interface_address_offset: %s", err)
		panic(errMsg)
	}
	v.FldValidators["local_interface_address_offset"] = vFn

	return v
}()

func AllocationSchemeValidator() db.Validator {
	return DefaultAllocationSchemeValidator
}

// augmented methods on protoc/std generated struct

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) ModeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(AllocatorMode)
		return int32(i)
	}
	// AllocatorMode_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, AllocatorMode_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mode")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) AddressAllocationSchemeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for address_allocation_scheme")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := AllocationSchemeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) AddressPoolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for address_pool")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for address_pool")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated address_pool")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items address_pool")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address_allocation_scheme"]; exists {

		vOpts := append(opts, db.WithValidateField("address_allocation_scheme"))
		if err := fv(ctx, m.GetAddressAllocationScheme(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["address_pool"]; exists {
		vOpts := append(opts, db.WithValidateField("address_pool"))
		if err := fv(ctx, m.GetAddressPool(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMode := v.ModeValidationRuleHandler
	rulesMode := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMode(rulesMode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.mode: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mode"] = vFn

	vrhAddressAllocationScheme := v.AddressAllocationSchemeValidationRuleHandler
	rulesAddressAllocationScheme := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAddressAllocationScheme(rulesAddressAllocationScheme)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.address_allocation_scheme: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address_allocation_scheme"] = vFn

	vrhAddressPool := v.AddressPoolValidationRuleHandler
	rulesAddressPool := map[string]string{
		"ves.io.schema.rules.message.required":                  "true",
		"ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
		"ves.io.schema.rules.repeated.max_items":                "32",
		"ves.io.schema.rules.repeated.min_items":                "1",
	}
	vFn, err = vrhAddressPool(rulesAddressPool)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.address_pool: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address_pool"] = vFn

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) ModeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(AllocatorMode)
		return int32(i)
	}
	// AllocatorMode_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, AllocatorMode_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mode")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) AddressAllocationSchemeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for address_allocation_scheme")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := AllocationSchemeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) AddressPoolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for address_pool")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for address_pool")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated address_pool")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items address_pool")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address_allocation_scheme"]; exists {

		vOpts := append(opts, db.WithValidateField("address_allocation_scheme"))
		if err := fv(ctx, m.GetAddressAllocationScheme(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["address_pool"]; exists {
		vOpts := append(opts, db.WithValidateField("address_pool"))
		if err := fv(ctx, m.GetAddressPool(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMode := v.ModeValidationRuleHandler
	rulesMode := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMode(rulesMode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.mode: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mode"] = vFn

	vrhAddressAllocationScheme := v.AddressAllocationSchemeValidationRuleHandler
	rulesAddressAllocationScheme := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAddressAllocationScheme(rulesAddressAllocationScheme)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.address_allocation_scheme: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address_allocation_scheme"] = vFn

	vrhAddressPool := v.AddressPoolValidationRuleHandler
	rulesAddressPool := map[string]string{
		"ves.io.schema.rules.message.required":                  "true",
		"ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
		"ves.io.schema.rules.repeated.max_items":                "32",
		"ves.io.schema.rules.repeated.min_items":                "1",
	}
	vFn, err = vrhAddressPool(rulesAddressPool)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.address_pool: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address_pool"] = vFn

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) ModeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(AllocatorMode)
		return int32(i)
	}
	// AllocatorMode_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, AllocatorMode_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mode")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) AddressAllocationSchemeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for address_allocation_scheme")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := AllocationSchemeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) AddressPoolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for address_pool")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for address_pool")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated address_pool")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items address_pool")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address_allocation_scheme"]; exists {

		vOpts := append(opts, db.WithValidateField("address_allocation_scheme"))
		if err := fv(ctx, m.GetAddressAllocationScheme(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["address_pool"]; exists {
		vOpts := append(opts, db.WithValidateField("address_pool"))
		if err := fv(ctx, m.GetAddressPool(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["allocation_map"]; exists {

		vOpts := append(opts, db.WithValidateField("allocation_map"))
		if err := fv(ctx, m.GetAllocationMap(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMode := v.ModeValidationRuleHandler
	rulesMode := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMode(rulesMode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.mode: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mode"] = vFn

	vrhAddressAllocationScheme := v.AddressAllocationSchemeValidationRuleHandler
	rulesAddressAllocationScheme := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAddressAllocationScheme(rulesAddressAllocationScheme)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.address_allocation_scheme: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address_allocation_scheme"] = vFn

	vrhAddressPool := v.AddressPoolValidationRuleHandler
	rulesAddressPool := map[string]string{
		"ves.io.schema.rules.message.required":                  "true",
		"ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
		"ves.io.schema.rules.repeated.max_items":                "32",
		"ves.io.schema.rules.repeated.min_items":                "1",
	}
	vFn, err = vrhAddressPool(rulesAddressPool)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.address_pool: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address_pool"] = vFn

	v.FldValidators["allocation_map"] = NodePrefixMapTypeValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *NodePrefixMapType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NodePrefixMapType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NodePrefixMapType) DeepCopy() *NodePrefixMapType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NodePrefixMapType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NodePrefixMapType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NodePrefixMapType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NodePrefixMapTypeValidator().Validate(ctx, m, opts...)
}

type ValidateNodePrefixMapType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNodePrefixMapType) EndpointsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemKeyRules := db.GetMapStringKeyRules(rules)
	itemKeyFn, err := db.NewStringValidationRuleHandler(itemKeyRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item key ValidationRuleHandler for endpoints")
	}
	itemsValidatorFn := func(ctx context.Context, kv map[string]*NodePrefixType, opts ...db.ValidateOpt) error {
		for key, value := range kv {
			if err := itemKeyFn(ctx, key, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element with key %v", key))
			}
			if err := NodePrefixTypeValidator().Validate(ctx, value, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("value for element with key %v", key))
			}
		}
		return nil
	}
	mapValFn, err := db.NewMapValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Map ValidationRuleHandler for endpoints")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.(map[string]*NodePrefixType)
		if !ok {
			return fmt.Errorf("Map validation expected map[ string ]*NodePrefixType, got %T", val)
		}
		if err := mapValFn(ctx, len(elems), opts...); err != nil {
			return errors.Wrap(err, "map endpoints")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items endpoints")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateNodePrefixMapType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NodePrefixMapType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NodePrefixMapType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["endpoints"]; exists {
		vOpts := append(opts, db.WithValidateField("endpoints"))
		if err := fv(ctx, m.GetEndpoints(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultNodePrefixMapTypeValidator = func() *ValidateNodePrefixMapType {
	v := &ValidateNodePrefixMapType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEndpoints := v.EndpointsValidationRuleHandler
	rulesEndpoints := map[string]string{
		"ves.io.schema.rules.map.keys.string.max_len": "256",
		"ves.io.schema.rules.map.keys.string.min_len": "1",
		"ves.io.schema.rules.map.max_pairs":           "128",
	}
	vFn, err = vrhEndpoints(rulesEndpoints)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for NodePrefixMapType.endpoints: %s", err)
		panic(errMsg)
	}
	v.FldValidators["endpoints"] = vFn

	return v
}()

func NodePrefixMapTypeValidator() db.Validator {
	return DefaultNodePrefixMapTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *NodePrefixType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NodePrefixType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NodePrefixType) DeepCopy() *NodePrefixType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NodePrefixType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NodePrefixType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NodePrefixType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NodePrefixTypeValidator().Validate(ctx, m, opts...)
}

type ValidateNodePrefixType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNodePrefixType) PrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for prefix")
	}

	return validatorFn, nil
}

func (v *ValidateNodePrefixType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NodePrefixType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NodePrefixType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["prefix"]; exists {

		vOpts := append(opts, db.WithValidateField("prefix"))
		if err := fv(ctx, m.GetPrefix(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultNodePrefixTypeValidator = func() *ValidateNodePrefixType {
	v := &ValidateNodePrefixType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPrefix := v.PrefixValidationRuleHandler
	rulesPrefix := map[string]string{
		"ves.io.schema.rules.string.ipv4_prefix": "true",
	}
	vFn, err = vrhPrefix(rulesPrefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for NodePrefixType.prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["prefix"] = vFn

	return v
}()

func NodePrefixTypeValidator() db.Validator {
	return DefaultNodePrefixTypeValidator
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.AddressAllocationScheme = f.GetAddressAllocationScheme()
	m.AddressPool = f.GetAddressPool()
	m.Mode = f.GetMode()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.AddressAllocationScheme = m1.AddressAllocationScheme
	f.AddressPool = m1.AddressPool
	f.Mode = m1.Mode
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.AddressAllocationScheme = f.GetAddressAllocationScheme()
	m.AddressPool = f.GetAddressPool()
	m.Mode = f.GetMode()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.AddressAllocationScheme = m1.AddressAllocationScheme
	f.AddressPool = m1.AddressPool
	f.Mode = m1.Mode
}