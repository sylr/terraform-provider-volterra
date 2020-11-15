//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package site

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

// Create CustomStateAPI GRPC Client satisfying server.CustomClient
type CustomStateAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomStateAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomStateAPIGrpcClient) doRPCSetState(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetStateReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.SetStateReq", yamlReq)
	}
	rsp, err := c.grpcClient.SetState(ctx, req, opts...)
	return rsp, err
}

func (c *CustomStateAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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
	return rsp, nil
}

func NewCustomStateAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomStateAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomStateAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["SetState"] = ccl.doRPCSetState

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomStateAPI REST Client satisfying server.CustomClient
type CustomStateAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomStateAPIRestClient) doRPCSetState(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetStateReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.SetStateReq: %s", yamlReq, err)
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
		q.Add("state", fmt.Sprintf("%v", req.State))

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
	pbRsp := &SetStateResp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.site.SetStateResp", body)
	}
	return pbRsp, nil
}

func (c *CustomStateAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomStateAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomStateAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["SetState"] = ccl.doRPCSetState

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomStateAPIInprocClient

// INPROC Client (satisfying CustomStateAPIClient interface)
type CustomStateAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomStateAPIInprocClient) SetState(ctx context.Context, in *SetStateReq, opts ...grpc.CallOption) (*SetStateResp, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.site.CustomStateAPI")
	cah, ok := ah.(CustomStateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomStateAPISrv", ah)
	}

	var (
		rsp *SetStateResp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.site.SetStateReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomStateAPI.SetState' operation on 'site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	rsp, err = cah.SetState(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.site.SetStateResp", rsp)...)

	return rsp, nil
}

func NewCustomStateAPIInprocClient(svc svcfw.Service) CustomStateAPIClient {
	return &CustomStateAPIInprocClient{svc: svc}
}

// RegisterGwCustomStateAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomStateAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomStateAPIHandlerClient(ctx, mux, NewCustomStateAPIInprocClient(s))
}

var CustomStateAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Site status",
        "description": "Site objects can be in various states and this state defined how (and if) site is functional.\nObject transitions are limited by state machine so only some transitions can be triggered by \nuser and transition is always depending on previous state, e.g. site in UGPRADING state can't\nmoved to different state by user and it's necessary to wait for the system to change state",
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
        "/public/namespaces/{namespace}/site/{name}/state": {
            "post": {
                "summary": "Set site state",
                "description": "Request changing site state but this request goes through validation as some\ntrainsitions are not allowed.\nIt can be used to decomission site by sending state DECOMISSIONING. Example of \nforbidden state is PROVISIONING and UPGRADING.",
                "operationId": "ves.io.schema.site.CustomStateAPI.SetState",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/siteSetStateResp"
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
                            "$ref": "#/definitions/siteSetStateReq"
                        }
                    }
                ],
                "tags": [
                    "CustomStateAPI"
                ],
                "x-ves-proto-rpc": "ves.io.schema.site.CustomStateAPI.SetState"
            },
            "x-displayname": "Site Status",
            "x-ves-proto-service": "ves.io.schema.site.CustomStateAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "siteSetStateReq": {
            "type": "object",
            "description": "Set status of the site",
            "title": "Set state request",
            "x-displayname": "Set Status Request",
            "x-ves-proto-message": "ves.io.schema.site.SetStateReq",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Site name\n\nExample: - \"ce398\"-\nRequired: YES",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "ce398",
                    "x-ves-required": "true"
                },
                "namespace": {
                    "type": "string",
                    "description": " Site namespace\n\nExample: - \"system\"-\nRequired: YES",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "system",
                    "x-ves-required": "true"
                },
                "state": {
                    "description": " Desired (target) state for site (3 = STANDBY)\n\nExample: - 3 -\nRequired: YES",
                    "title": "State",
                    "$ref": "#/definitions/siteSiteState",
                    "x-displayname": "State",
                    "x-ves-required": "true"
                }
            }
        },
        "siteSetStateResp": {
            "type": "object",
            "description": "Response for set state request, empty because the only resturned information\nis currently error message",
            "title": "Set state responde",
            "x-displayname": "Set Status Response",
            "x-ves-proto-message": "ves.io.schema.site.SetStateResp"
        },
        "siteSiteState": {
            "type": "string",
            "description": "State of Site defines in which operational state site itself is.\n\nSite is online and operational.\nSite is in provisioning state. For instance during site deployment or switching to different connected Regional Edge.\nSite is in process of upgrade. It transition to ONLINE or FAILED state.\nSite is in Standby before goes to ONLINE. This is mainly for Regional Edge sites to do their verification before they go to ONLINE state.\nSite is in failed state. It failed during provisioning or upgrade phase. Site Status Objects contain more details.\nReregistration was requested\nReregistration is in progress and maurice is waiting for nodes\nSite deletion is in progress\nSite is waiting for registration",
            "title": "SiteState",
            "enum": [
                "ONLINE",
                "PROVISIONING",
                "UPGRADING",
                "STANDBY",
                "FAILED",
                "REREGISTRATION",
                "WAITINGNODES",
                "DECOMMISSIONING",
                "WAITING_FOR_REGISTRATION"
            ],
            "default": "ONLINE",
            "x-displayname": "Site State",
            "x-ves-proto-enum": "ves.io.schema.site.SiteState"
        }
    },
    "x-displayname": "Site",
    "x-ves-proto-file": "ves.io/schema/site/public_customapi.proto"
}`
