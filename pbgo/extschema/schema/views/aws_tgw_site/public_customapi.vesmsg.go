//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package aws_tgw_site

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

func (m *SetVPCIpPrefixesRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPCIpPrefixesRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPCIpPrefixesRequest) DeepCopy() *SetVPCIpPrefixesRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPCIpPrefixesRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPCIpPrefixesRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPCIpPrefixesRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPCIpPrefixesRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPCIpPrefixesRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPCIpPrefixesRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPCIpPrefixesRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPCIpPrefixesRequest got type %s", t)
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

	if fv, exists := v.FldValidators["vpc_ip_prefixes"]; exists {

		vOpts := append(opts, db.WithValidateField("vpc_ip_prefixes"))
		for key, value := range m.GetVpcIpPrefixes() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPCIpPrefixesRequestValidator = func() *ValidateSetVPCIpPrefixesRequest {
	v := &ValidateSetVPCIpPrefixesRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["vpc_ip_prefixes"] = VPCIpPrefixesTypeValidator().Validate

	return v
}()

func SetVPCIpPrefixesRequestValidator() db.Validator {
	return DefaultSetVPCIpPrefixesRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPCIpPrefixesResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPCIpPrefixesResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPCIpPrefixesResponse) DeepCopy() *SetVPCIpPrefixesResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPCIpPrefixesResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPCIpPrefixesResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPCIpPrefixesResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPCIpPrefixesResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPCIpPrefixesResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPCIpPrefixesResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPCIpPrefixesResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPCIpPrefixesResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPCIpPrefixesResponseValidator = func() *ValidateSetVPCIpPrefixesResponse {
	v := &ValidateSetVPCIpPrefixesResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetVPCIpPrefixesResponseValidator() db.Validator {
	return DefaultSetVPCIpPrefixesResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPNTunnelsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPNTunnelsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPNTunnelsRequest) DeepCopy() *SetVPNTunnelsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPNTunnelsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPNTunnelsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPNTunnelsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPNTunnelsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPNTunnelsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPNTunnelsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPNTunnelsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPNTunnelsRequest got type %s", t)
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

	if fv, exists := v.FldValidators["tunnels"]; exists {

		vOpts := append(opts, db.WithValidateField("tunnels"))
		for idx, item := range m.GetTunnels() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPNTunnelsRequestValidator = func() *ValidateSetVPNTunnelsRequest {
	v := &ValidateSetVPNTunnelsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tunnels"] = AWSVPNTunnelConfigTypeValidator().Validate

	return v
}()

func SetVPNTunnelsRequestValidator() db.Validator {
	return DefaultSetVPNTunnelsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPNTunnelsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPNTunnelsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPNTunnelsResponse) DeepCopy() *SetVPNTunnelsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPNTunnelsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPNTunnelsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPNTunnelsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPNTunnelsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPNTunnelsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPNTunnelsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPNTunnelsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPNTunnelsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPNTunnelsResponseValidator = func() *ValidateSetVPNTunnelsResponse {
	v := &ValidateSetVPNTunnelsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetVPNTunnelsResponseValidator() db.Validator {
	return DefaultSetVPNTunnelsResponseValidator
}
