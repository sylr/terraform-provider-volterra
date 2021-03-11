//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package aws_tgw_site

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// Create CustomAPI GRPC Client satisfying server.CustomClient
type CustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomAPIGrpcClient) doRPCSetTGWInfo(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetTGWInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_tgw_site.SetTGWInfoRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SetTGWInfo(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) doRPCSetVPCIpPrefixes(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetVPCIpPrefixesRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SetVPCIpPrefixes(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) doRPCSetVPNTunnels(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetVPNTunnelsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_tgw_site.SetVPNTunnelsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SetVPNTunnels(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}
	if cco.YAMLReq == "" {
		return nil, fmt.Errorf("Error, empty request body")
	}
	ctx = client.AddHdrsToCtx(cco.Headers, ctx)

	rsp, err := rpcFn(ctx, cco.YAMLReq, cco.GrpcCallOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using GRPC")
	}
	if cco.OutCallResponse != nil {
		cco.OutCallResponse.ProtoMsg = rsp
	}
	return rsp, nil
}

func NewCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["SetTGWInfo"] = ccl.doRPCSetTGWInfo

	rpcFns["SetVPCIpPrefixes"] = ccl.doRPCSetVPCIpPrefixes

	rpcFns["SetVPNTunnels"] = ccl.doRPCSetVPNTunnels

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPI REST Client satisfying server.CustomClient
type CustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomAPIRestClient) doRPCSetTGWInfo(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetTGWInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_tgw_site.SetTGWInfoRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		newReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP POST request for custom API")
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("tgw_info", fmt.Sprintf("%v", req.TgwInfo))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &SetTGWInfoResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.views.aws_tgw_site.SetTGWInfoResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) doRPCSetVPCIpPrefixes(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetVPCIpPrefixesRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		newReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP POST request for custom API")
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("vpc_ip_prefixes", fmt.Sprintf("%v", req.VpcIpPrefixes))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &SetVPCIpPrefixesResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) doRPCSetVPNTunnels(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetVPNTunnelsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_tgw_site.SetVPNTunnelsRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		newReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP POST request for custom API")
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("tunnels", fmt.Sprintf("%v", req.Tunnels))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &SetVPNTunnelsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.views.aws_tgw_site.SetVPNTunnelsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}

	rsp, err := rpcFn(ctx, cco)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using Rest")
	}
	return rsp, nil
}

func NewCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["SetTGWInfo"] = ccl.doRPCSetTGWInfo

	rpcFns["SetVPCIpPrefixes"] = ccl.doRPCSetVPCIpPrefixes

	rpcFns["SetVPNTunnels"] = ccl.doRPCSetVPNTunnels

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type CustomAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomAPIInprocClient) SetTGWInfo(ctx context.Context, in *SetTGWInfoRequest, opts ...grpc.CallOption) (*SetTGWInfoResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.views.aws_tgw_site.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPISrv", ah)
	}

	var (
		rsp *SetTGWInfoResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.views.aws_tgw_site.SetTGWInfoRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SetTGWInfo' operation on 'aws_tgw_site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.views.aws_tgw_site.CustomAPI.SetTGWInfo"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SetTGWInfo(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.views.aws_tgw_site.SetTGWInfoResponse", rsp)...)

	return rsp, nil
}
func (c *CustomAPIInprocClient) SetVPCIpPrefixes(ctx context.Context, in *SetVPCIpPrefixesRequest, opts ...grpc.CallOption) (*SetVPCIpPrefixesResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.views.aws_tgw_site.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPISrv", ah)
	}

	var (
		rsp *SetVPCIpPrefixesResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SetVPCIpPrefixes' operation on 'aws_tgw_site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPCIpPrefixes"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SetVPCIpPrefixes(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesResponse", rsp)...)

	return rsp, nil
}
func (c *CustomAPIInprocClient) SetVPNTunnels(ctx context.Context, in *SetVPNTunnelsRequest, opts ...grpc.CallOption) (*SetVPNTunnelsResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.views.aws_tgw_site.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPISrv", ah)
	}

	var (
		rsp *SetVPNTunnelsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.views.aws_tgw_site.SetVPNTunnelsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SetVPNTunnels' operation on 'aws_tgw_site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPNTunnels"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SetVPNTunnels(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.views.aws_tgw_site.SetVPNTunnelsResponse", rsp)...)

	return rsp, nil
}

func NewCustomAPIInprocClient(svc svcfw.Service) CustomAPIClient {
	return &CustomAPIInprocClient{svc: svc}
}

// RegisterGwCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomAPIHandlerClient(ctx, mux, NewCustomAPIInprocClient(s))
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "AWS TGW site",
        "description": "AWS TGW site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in AWS VPC.\nIt can be used to either automatically create or Manually assisted site creation in AWS TGW.\n\nView will create following child objects.\n\n* Site",
        "version": "version not set"
    },
    "schemes": [
        "http",
        "https"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "tags": null,
    "paths": {
        "/public/namespaces/{namespace}/aws_tgw_site/{name}/set_tgw_info": {
            "post": {
                "summary": "Configure TGW Information",
                "description": "Configure TGW Information like tgw-id, volterra site's vpc-id",
                "operationId": "ves.io.schema.views.aws_tgw_site.CustomAPI.SetTGWInfo",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/aws_tgw_siteSetTGWInfoResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "name",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aws_tgw_siteSetTGWInfoRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-aws_tgw_site-CustomAPI-SetTGWInfo"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.aws_tgw_site.CustomAPI.SetTGWInfo"
            },
            "x-displayname": "Custom API for AWS TGW site",
            "x-ves-proto-service": "ves.io.schema.views.aws_tgw_site.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/aws_tgw_site/{name}/set_vpc_ip_prefixes": {
            "post": {
                "summary": "Configure VPC IP prefixes",
                "description": "Configure VPC IP prefix set",
                "operationId": "ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPCIpPrefixes",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/aws_tgw_siteSetVPCIpPrefixesResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "name",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aws_tgw_siteSetVPCIpPrefixesRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-aws_tgw_site-CustomAPI-SetVPCIpPrefixes"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPCIpPrefixes"
            },
            "x-displayname": "Custom API for AWS TGW site",
            "x-ves-proto-service": "ves.io.schema.views.aws_tgw_site.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/aws_tgw_site/{name}/set_vpn_tunnels": {
            "post": {
                "summary": "Configure VPN tunnels",
                "description": "Configure VPC IP prefix set",
                "operationId": "ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPNTunnels",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/aws_tgw_siteSetVPNTunnelsResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "name",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aws_tgw_siteSetVPNTunnelsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-aws_tgw_site-CustomAPI-SetVPNTunnels"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPNTunnels"
            },
            "x-displayname": "Custom API for AWS TGW site",
            "x-ves-proto-service": "ves.io.schema.views.aws_tgw_site.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "aws_tgw_siteAWSTGWInfoConfigType": {
            "type": "object",
            "description": "AWS tgw information like tgw-id and site's vpc-id",
            "title": "AWS TGW Information Config",
            "x-displayname": "AWS TGW Information Config",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.AWSTGWInfoConfigType",
            "properties": {
                "tgw_id": {
                    "type": "string",
                    "description": " TGW ID populated by AWS\n\nExample: - \"tgw-12345678\"-\nRequired: YES",
                    "title": "TGW ID",
                    "x-displayname": "TGW ID",
                    "x-ves-example": "tgw-12345678",
                    "x-ves-required": "true"
                },
                "vpc_id": {
                    "type": "string",
                    "description": " VPC ID where the volterra site exists\n\nExample: - \"vpc-12345678\"-\nRequired: YES",
                    "title": "VPC ID",
                    "x-displayname": "VPC ID",
                    "x-ves-example": "vpc-12345678",
                    "x-ves-required": "true"
                }
            }
        },
        "aws_tgw_siteAWSVPNTunnelConfigType": {
            "type": "object",
            "description": "Remote IP for VPN tunnels of a node",
            "title": "AWS VPN Tunnel Config",
            "x-displayname": "AWS VPN Tunnel Config",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.AWSVPNTunnelConfigType",
            "properties": {
                "node_id": {
                    "type": "string",
                    "description": " Volterra Node ID for which this tunnel is configured\n\nExample: - \"ves-node-id-xxxxxx\"-\nRequired: YES",
                    "title": "Volterra Node ID",
                    "x-displayname": "Volterra Node ID",
                    "x-ves-example": "ves-node-id-xxxxxx",
                    "x-ves-required": "true"
                },
                "node_name": {
                    "type": "string",
                    "description": " Name of the node for which this tunnel is configured\n\nExample: - \"master-0\"-\nRequired: YES",
                    "title": "Name of the Node",
                    "x-displayname": "Name of the Node",
                    "x-ves-example": "master-0",
                    "x-ves-required": "true"
                },
                "tunnel_remote_ip": {
                    "type": "array",
                    "description": " Remote IP(s) for up to two tunnels\n\nExample: - \"3.4.5.6\"-\nRequired: YES",
                    "title": "Remote IP(s)",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Remote IP(s)",
                    "x-ves-example": "3.4.5.6",
                    "x-ves-required": "true"
                }
            }
        },
        "aws_tgw_siteSetTGWInfoRequest": {
            "type": "object",
            "description": "Request to configure TGW Information",
            "title": "Request to configure TGW Information",
            "x-displayname": "Request to configure TGW Information",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.SetTGWInfoRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of the object to be configured\n\nExample: - \"aws-tgw-site-1\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "aws-tgw-site-1"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "tgw_info": {
                    "description": " AWS TGW Info Config",
                    "title": "AWS TGW Info Config",
                    "$ref": "#/definitions/aws_tgw_siteAWSTGWInfoConfigType",
                    "x-displayname": "AWS TGW Info Config"
                }
            }
        },
        "aws_tgw_siteSetTGWInfoResponse": {
            "type": "object",
            "description": "Response to configure TGW info",
            "title": "Response to configure TGW info",
            "x-displayname": "Response to configure TGW info",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.SetTGWInfoResponse"
        },
        "aws_tgw_siteSetVPCIpPrefixesRequest": {
            "type": "object",
            "description": "Request to configure VPC IP prefix set",
            "title": "Request to configure VPC IP prefix set",
            "x-displayname": "Request to configure VPC IP prefix set",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of the object to be configured\n\nExample: - \"aws-tgw-site-1\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "aws-tgw-site-1"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "vpc_ip_prefixes": {
                    "type": "object",
                    "description": " IP prefixes of subnets in the VPC.",
                    "title": "VPC IP Prefixes",
                    "x-displayname": "VPC IP Prefixes"
                }
            }
        },
        "aws_tgw_siteSetVPCIpPrefixesResponse": {
            "type": "object",
            "description": "Response to configure VPC IP prefix set",
            "title": "Response to configure VPC IP prefix set",
            "x-displayname": "Response to configure VPC IP prefix set",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesResponse"
        },
        "aws_tgw_siteSetVPNTunnelsRequest": {
            "type": "object",
            "description": "Request to configure VPN Tunnels",
            "title": "Request to configure VPN Tunnels",
            "x-displayname": "Request to configure VPN Tunnels",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.SetVPNTunnelsRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of the object to be configured\n\nExample: - \"aws-tgw-site-1\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "aws-tgw-site-1"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "tunnels": {
                    "type": "array",
                    "description": " AWS VPN Tunner Config",
                    "title": "AWS VPN tunnel config",
                    "items": {
                        "$ref": "#/definitions/aws_tgw_siteAWSVPNTunnelConfigType"
                    },
                    "x-displayname": "AWS VPN Tunnel Config"
                }
            }
        },
        "aws_tgw_siteSetVPNTunnelsResponse": {
            "type": "object",
            "description": "Response to configure VPN Tunnels",
            "title": "Response to configure VPN Tunnels",
            "x-displayname": "Response to configure VPN Tunnels",
            "x-ves-proto-message": "ves.io.schema.views.aws_tgw_site.SetVPNTunnelsResponse"
        },
        "aws_tgw_siteVPCIpPrefixesType": {
            "type": "object",
            "description": "x-displayName: \"VPC IP prefixes\"\nVPC IP prefixes",
            "title": "VPC IP prefixes",
            "properties": {
                "prefixes": {
                    "type": "array",
                    "description": "x-displayName: \"Prefixes\"\nx-required\nx-example: \"['10.2.1.0/24', '192.168.8.0/29', '10.7.64.160/27']\"\nAn unordered list of IP prefixes.",
                    "title": "IP prefixes",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    },
    "x-displayname": "Configure AWS TGW Site",
    "x-ves-proto-file": "ves.io/schema/views/aws_tgw_site/public_customapi.proto"
}`
