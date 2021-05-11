//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package user

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

// Create CustomPrivateAPI GRPC Client satisfying server.CustomClient
type CustomPrivateAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomPrivateAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomPrivateAPIGrpcClient) doRPCUpdateLastLogin(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &LastLoginUpdateRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.user.LastLoginUpdateRequest", yamlReq)
	}
	rsp, err := c.grpcClient.UpdateLastLogin(ctx, req, opts...)
	return rsp, err
}

func (c *CustomPrivateAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomPrivateAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomPrivateAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomPrivateAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["UpdateLastLogin"] = ccl.doRPCUpdateLastLogin

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomPrivateAPI REST Client satisfying server.CustomClient
type CustomPrivateAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomPrivateAPIRestClient) doRPCUpdateLastLogin(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &LastLoginUpdateRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.user.LastLoginUpdateRequest: %s", yamlReq, err)
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
		q.Add("last_login_timestamp", fmt.Sprintf("%v", req.LastLoginTimestamp))
		q.Add("tenant", fmt.Sprintf("%v", req.Tenant))
		q.Add("user", fmt.Sprintf("%v", req.User))

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
	pbRsp := &LastLoginUpdateResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.user.LastLoginUpdateResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomPrivateAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomPrivateAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomPrivateAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["UpdateLastLogin"] = ccl.doRPCUpdateLastLogin

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomPrivateAPIInprocClient

// INPROC Client (satisfying CustomPrivateAPIClient interface)
type CustomPrivateAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomPrivateAPIInprocClient) UpdateLastLogin(ctx context.Context, in *LastLoginUpdateRequest, opts ...grpc.CallOption) (*LastLoginUpdateResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.user.CustomPrivateAPI")
	cah, ok := ah.(CustomPrivateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomPrivateAPISrv", ah)
	}

	var (
		rsp *LastLoginUpdateResponse
		err error
	)

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.UpdateLastLogin(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}

func NewCustomPrivateAPIInprocClient(svc svcfw.Service) CustomPrivateAPIClient {
	return &CustomPrivateAPIInprocClient{svc: svc}
}

// RegisterGwCustomPrivateAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomPrivateAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomPrivateAPIHandlerClient(ctx, mux, NewCustomPrivateAPIInprocClient(s))
}

var CustomPrivateAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "User Custom Private API",
        "description": "Custom private APIs for user management",
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
        "/private/custom/namespaces/system/users/update_last_login": {
            "post": {
                "summary": "Update Last Login",
                "description": "API to update last login timestamp of user",
                "operationId": "ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/userLastLoginUpdateResponse"
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
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userLastLoginUpdateRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-user-CustomPrivateAPI-UpdateLastLogin"
                },
                "x-ves-proto-rpc": "ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin"
            },
            "x-displayname": "Custom Private API",
            "x-ves-proto-service": "ves.io.schema.user.CustomPrivateAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        }
    },
    "definitions": {
        "userLastLoginUpdateRequest": {
            "type": "object",
            "description": "Request to update user login timestamp.",
            "title": "LastLoginUpdateRequest",
            "x-displayname": "Last Login Update Request",
            "x-ves-proto-message": "ves.io.schema.user.LastLoginUpdateRequest",
            "properties": {
                "last_login_timestamp": {
                    "type": "string",
                    "description": " Last successful login timestamp of the user .",
                    "title": "last_login_timestamp",
                    "format": "date-time",
                    "x-displayname": "Last Login Timestamp"
                },
                "tenant": {
                    "type": "string",
                    "description": " Tenant ID of the tenant user belongs to.\n\nExample: - \"company1-as432s\"-",
                    "title": "tenant",
                    "x-displayname": "Tenant",
                    "x-ves-example": "company1-as432s"
                },
                "user": {
                    "type": "string",
                    "description": " User ID of the user. typically email id\n\nExample: - \"user@company1.com\"-",
                    "title": "user",
                    "x-displayname": "User",
                    "x-ves-example": "user@company1.com"
                }
            }
        },
        "userLastLoginUpdateResponse": {
            "type": "object",
            "title": "LastLoginUpdateResponse",
            "x-displayname": "Last Login Update Response",
            "x-ves-proto-message": "ves.io.schema.user.LastLoginUpdateResponse"
        }
    },
    "x-displayname": "User custom private API",
    "x-ves-proto-file": "ves.io/schema/user/private_customapi.proto"
}`