//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_virtual_host "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
)

// resourceVolterraVirtualHost is implementation of Volterra's VirtualHost resources
func resourceVolterraVirtualHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraVirtualHostCreate,
		Read:   resourceVolterraVirtualHostRead,
		Update: resourceVolterraVirtualHostUpdate,
		Delete: resourceVolterraVirtualHostDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"add_location": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"advertise_policies": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"buffer_policy": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"max_request_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"max_request_time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"captcha_challenge": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cookie_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"enable_captcha_challenge": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"js_challenge": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cookie_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"enable_js_challenge": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"js_script_delay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"no_challenge": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"compression_params": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"content_length": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"content_type": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"disable_on_etag_header": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"remove_accept_encoding_header": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"cors_policy": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allow_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"allow_headers": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"allow_methods": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"allow_origin": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"allow_origin_regex": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"disabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"expose_headers": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"max_age": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"maximum_age": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"custom_errors": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"disable_default_error_pages": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_dns_resolve": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"domains": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"dynamic_reverse_proxy": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"resolution_network": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"resolution_network_type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"resolve_endpoint_dynamically": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"max_request_header_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"proxy": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"rate_limiter": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"rate_limiter_allowed_prefixes": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"request_headers_to_add": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"append": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"request_headers_to_remove": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"response_headers_to_add": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"append": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"response_headers_to_remove": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"retry_policy": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"back_off": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"base_interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"num_retries": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"per_try_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"retriable_status_codes": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						"retry_on": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"routes": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"temporary_user_blocking": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"tls_parameters": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"common_params": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cipher_suites": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"maximum_protocol_version": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"minimum_protocol_version": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"tls_certificates": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate_url": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"private_key": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"secret_encoding_type": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"blindfold_secret_info": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"decryption_provider": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"location": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"store_provider": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"clear_secret_info": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"provider": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"url": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vault_secret_info": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"key": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"location": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"provider": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"secret_encoding": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"version": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},

															"wingman_secret_info": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"name": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},

									"trusted_ca_url": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"validation_params": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"skip_hostname_verification": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"trusted_ca_url": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"use_volterra_trusted_ca_url": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"verify_subject_alt_names": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},

						"require_client_certificate": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"user_identification": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"waf_type": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"waf": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"waf": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"tenant": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"waf_rules": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"waf_rules": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"tenant": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraVirtualHostCreate creates VirtualHost resource
func resourceVolterraVirtualHostCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_virtual_host.CreateSpecType{}
	createReq := &ves_io_schema_virtual_host.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("add_location"); ok && !isIntfNil(v) {

		createSpec.AddLocation =
			v.(bool)
	}

	if v, ok := d.GetOk("advertise_policies"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		advertisePoliciesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.AdvertisePolicies = advertisePoliciesInt
		for i, ps := range sl {

			apMapToStrVal := ps.(map[string]interface{})
			advertisePoliciesInt[i] = &ves_io_schema.ObjectRefType{}

			advertisePoliciesInt[i].Kind = "advertise_policy"

			if v, ok := apMapToStrVal["name"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Name = v.(string)
			}

			if v, ok := apMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Namespace = v.(string)
			}

			if v, ok := apMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Tenant = v.(string)
			}

			if v, ok := apMapToStrVal["uid"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("buffer_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		bufferPolicy := &ves_io_schema.BufferConfigType{}
		createSpec.BufferPolicy = bufferPolicy
		for _, set := range sl {

			bufferPolicyMapStrToI := set.(map[string]interface{})

			if w, ok := bufferPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
				bufferPolicy.Disabled = w.(bool)
			}

			if w, ok := bufferPolicyMapStrToI["max_request_bytes"]; ok && !isIntfNil(w) {
				bufferPolicy.MaxRequestBytes = w.(uint32)
			}

			if w, ok := bufferPolicyMapStrToI["max_request_time"]; ok && !isIntfNil(w) {
				bufferPolicy.MaxRequestTime = w.(uint32)
			}

		}

	}

	challengeTypeTypeFound := false

	if v, ok := d.GetOk("captcha_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_virtual_host.CreateSpecType_CaptchaChallenge{}
		challengeTypeInt.CaptchaChallenge = &ves_io_schema_virtual_host.CaptchaChallengeType{}
		createSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_captcha_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.EnableCaptchaChallenge = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("js_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_virtual_host.CreateSpecType_JsChallenge{}
		challengeTypeInt.JsChallenge = &ves_io_schema_virtual_host.JavascriptChallengeType{}
		createSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_js_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.EnableJsChallenge = v.(bool)
			}

			if v, ok := cs["js_script_delay"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.JsScriptDelay = uint32(v.(int))
			}

		}

	}

	if v, ok := d.GetOk("no_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true

		if v.(bool) {
			challengeTypeInt := &ves_io_schema_virtual_host.CreateSpecType_NoChallenge{}
			challengeTypeInt.NoChallenge = &ves_io_schema.Empty{}
			createSpec.ChallengeType = challengeTypeInt
		}

	}

	if v, ok := d.GetOk("compression_params"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		compressionParams := &ves_io_schema_virtual_host.CompressionType{}
		createSpec.CompressionParams = compressionParams
		for _, set := range sl {

			compressionParamsMapStrToI := set.(map[string]interface{})

			if w, ok := compressionParamsMapStrToI["content_length"]; ok && !isIntfNil(w) {
				compressionParams.ContentLength = w.(uint32)
			}

			if w, ok := compressionParamsMapStrToI["content_type"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				compressionParams.ContentType = ls
			}

			if w, ok := compressionParamsMapStrToI["disable_on_etag_header"]; ok && !isIntfNil(w) {
				compressionParams.DisableOnEtagHeader = w.(bool)
			}

			if w, ok := compressionParamsMapStrToI["remove_accept_encoding_header"]; ok && !isIntfNil(w) {
				compressionParams.RemoveAcceptEncodingHeader = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("cors_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		corsPolicy := &ves_io_schema.CorsPolicy{}
		createSpec.CorsPolicy = corsPolicy
		for _, set := range sl {

			corsPolicyMapStrToI := set.(map[string]interface{})

			if w, ok := corsPolicyMapStrToI["allow_credentials"]; ok && !isIntfNil(w) {
				corsPolicy.AllowCredentials = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["allow_headers"]; ok && !isIntfNil(w) {
				corsPolicy.AllowHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_methods"]; ok && !isIntfNil(w) {
				corsPolicy.AllowMethods = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_origin"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOrigin = ls
			}

			if w, ok := corsPolicyMapStrToI["allow_origin_regex"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOriginRegex = ls
			}

			if w, ok := corsPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
				corsPolicy.Disabled = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["expose_headers"]; ok && !isIntfNil(w) {
				corsPolicy.ExposeHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["max_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaxAge = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["maximum_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaximumAge = w.(int32)
			}

		}

	}

	if v, ok := d.GetOk("disable_default_error_pages"); ok && !isIntfNil(v) {

		createSpec.DisableDefaultErrorPages =
			v.(bool)
	}

	if v, ok := d.GetOk("disable_dns_resolve"); ok && !isIntfNil(v) {

		createSpec.DisableDnsResolve =
			v.(bool)
	}

	if v, ok := d.GetOk("domains"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		createSpec.Domains = ls

	}

	if v, ok := d.GetOk("dynamic_reverse_proxy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		dynamicReverseProxy := &ves_io_schema_virtual_host.DynamicReverseProxyType{}
		createSpec.DynamicReverseProxy = dynamicReverseProxy
		for _, set := range sl {

			dynamicReverseProxyMapStrToI := set.(map[string]interface{})

			if v, ok := dynamicReverseProxyMapStrToI["resolution_network"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				resolutionNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				dynamicReverseProxy.ResolutionNetwork = resolutionNetworkInt
				for i, ps := range sl {

					rnMapToStrVal := ps.(map[string]interface{})
					resolutionNetworkInt[i] = &ves_io_schema.ObjectRefType{}

					resolutionNetworkInt[i].Kind = "virtual_network"

					if v, ok := rnMapToStrVal["name"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Name = v.(string)
					}

					if v, ok := rnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Namespace = v.(string)
					}

					if v, ok := rnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Tenant = v.(string)
					}

					if v, ok := rnMapToStrVal["uid"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Uid = v.(string)
					}

				}

			}

			if v, ok := dynamicReverseProxyMapStrToI["resolution_network_type"]; ok && !isIntfNil(v) {

				dynamicReverseProxy.ResolutionNetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

			}

			if w, ok := dynamicReverseProxyMapStrToI["resolve_endpoint_dynamically"]; ok && !isIntfNil(w) {
				dynamicReverseProxy.ResolveEndpointDynamically = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("idle_timeout"); ok && !isIntfNil(v) {

		createSpec.IdleTimeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("max_request_header_size"); ok && !isIntfNil(v) {

		createSpec.MaxRequestHeaderSize =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("proxy"); ok && !isIntfNil(v) {

		createSpec.Proxy = ves_io_schema_virtual_host.ProxyType(ves_io_schema_virtual_host.ProxyType_value[v.(string)])

	}

	if v, ok := d.GetOk("rate_limiter"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rateLimiterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.RateLimiter = rateLimiterInt
		for i, ps := range sl {

			rlMapToStrVal := ps.(map[string]interface{})
			rateLimiterInt[i] = &ves_io_schema.ObjectRefType{}

			rateLimiterInt[i].Kind = "rate_limiter"

			if v, ok := rlMapToStrVal["name"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Name = v.(string)
			}

			if v, ok := rlMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Namespace = v.(string)
			}

			if v, ok := rlMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Tenant = v.(string)
			}

			if v, ok := rlMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("rate_limiter_allowed_prefixes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rateLimiterAllowedPrefixesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.RateLimiterAllowedPrefixes = rateLimiterAllowedPrefixesInt
		for i, ps := range sl {

			rlapMapToStrVal := ps.(map[string]interface{})
			rateLimiterAllowedPrefixesInt[i] = &ves_io_schema.ObjectRefType{}

			rateLimiterAllowedPrefixesInt[i].Kind = "ip_prefix_set"

			if v, ok := rlapMapToStrVal["name"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Name = v.(string)
			}

			if v, ok := rlapMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Namespace = v.(string)
			}

			if v, ok := rlapMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Tenant = v.(string)
			}

			if v, ok := rlapMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("request_headers_to_add"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		requestHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
		createSpec.RequestHeadersToAdd = requestHeadersToAdd
		for i, set := range sl {
			requestHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

			requestHeadersToAddMapStrToI := set.(map[string]interface{})

			if w, ok := requestHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
				requestHeadersToAdd[i].Append = w.(bool)
			}

			if w, ok := requestHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
				requestHeadersToAdd[i].Name = w.(string)
			}

			if w, ok := requestHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
				requestHeadersToAdd[i].Value = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("request_headers_to_remove"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		createSpec.RequestHeadersToRemove = ls

	}

	if v, ok := d.GetOk("response_headers_to_add"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		responseHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
		createSpec.ResponseHeadersToAdd = responseHeadersToAdd
		for i, set := range sl {
			responseHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

			responseHeadersToAddMapStrToI := set.(map[string]interface{})

			if w, ok := responseHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
				responseHeadersToAdd[i].Append = w.(bool)
			}

			if w, ok := responseHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
				responseHeadersToAdd[i].Name = w.(string)
			}

			if w, ok := responseHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
				responseHeadersToAdd[i].Value = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("response_headers_to_remove"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		createSpec.ResponseHeadersToRemove = ls

	}

	if v, ok := d.GetOk("retry_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		retryPolicy := &ves_io_schema.RetryPolicyType{}
		createSpec.RetryPolicy = retryPolicy
		for _, set := range sl {

			retryPolicyMapStrToI := set.(map[string]interface{})

			if v, ok := retryPolicyMapStrToI["back_off"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				backOff := &ves_io_schema.RetryBackOff{}
				retryPolicy.BackOff = backOff
				for _, set := range sl {

					backOffMapStrToI := set.(map[string]interface{})

					if w, ok := backOffMapStrToI["base_interval"]; ok && !isIntfNil(w) {
						backOff.BaseInterval = w.(uint32)
					}

					if w, ok := backOffMapStrToI["max_interval"]; ok && !isIntfNil(w) {
						backOff.MaxInterval = w.(uint32)
					}

				}

			}

			if w, ok := retryPolicyMapStrToI["num_retries"]; ok && !isIntfNil(w) {
				retryPolicy.NumRetries = w.(uint32)
			}

			if w, ok := retryPolicyMapStrToI["per_try_timeout"]; ok && !isIntfNil(w) {
				retryPolicy.PerTryTimeout = w.(uint32)
			}

			if w, ok := retryPolicyMapStrToI["retriable_status_codes"]; ok && !isIntfNil(w) {
				ls := make([]uint32, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				retryPolicy.RetriableStatusCodes = ls
			}

			if w, ok := retryPolicyMapStrToI["retry_on"]; ok && !isIntfNil(w) {
				retryPolicy.RetryOn = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Routes = routesInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			routesInt[i] = &ves_io_schema.ObjectRefType{}

			routesInt[i].Kind = "route"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				routesInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				routesInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				routesInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				routesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("temporary_user_blocking"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		temporaryUserBlocking := &ves_io_schema_virtual_host.TemporaryUserBlockingType{}
		createSpec.TemporaryUserBlocking = temporaryUserBlocking
		for _, set := range sl {

			temporaryUserBlockingMapStrToI := set.(map[string]interface{})

			if w, ok := temporaryUserBlockingMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				temporaryUserBlocking.CustomPage = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("tls_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsParameters := &ves_io_schema.DownstreamTlsParamsType{}
		createSpec.TlsParameters = tlsParameters
		for _, set := range sl {

			tlsParametersMapStrToI := set.(map[string]interface{})

			if v, ok := tlsParametersMapStrToI["common_params"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				commonParams := &ves_io_schema.TlsParamsType{}
				tlsParameters.CommonParams = commonParams
				for _, set := range sl {

					commonParamsMapStrToI := set.(map[string]interface{})

					if w, ok := commonParamsMapStrToI["cipher_suites"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						commonParams.CipherSuites = ls
					}

					if v, ok := commonParamsMapStrToI["maximum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MaximumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["minimum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MinimumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						commonParams.TlsCertificates = tlsCertificates
						for i, set := range sl {
							tlsCertificates[i] = &ves_io_schema.TlsCertificateType{}

							tlsCertificatesMapStrToI := set.(map[string]interface{})

							if w, ok := tlsCertificatesMapStrToI["certificate_url"]; ok && !isIntfNil(w) {
								tlsCertificates[i].CertificateUrl = w.(string)
							}

							if w, ok := tlsCertificatesMapStrToI["description"]; ok && !isIntfNil(w) {
								tlsCertificates[i].Description = w.(string)
							}

							if v, ok := tlsCertificatesMapStrToI["private_key"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								privateKey := &ves_io_schema.SecretType{}
								tlsCertificates[i].PrivateKey = privateKey
								for _, set := range sl {

									privateKeyMapStrToI := set.(map[string]interface{})

									if v, ok := privateKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										privateKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := privateKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["decryption_provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.BlindfoldSecretInfo.DecryptionProvider = v.(string)
											}

											if v, ok := cs["location"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.BlindfoldSecretInfo.Location = v.(string)
											}

											if v, ok := cs["store_provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.BlindfoldSecretInfo.StoreProvider = v.(string)
											}

										}

									}

									if v, ok := privateKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
											}

											if v, ok := cs["url"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
											}

										}

									}

									if v, ok := privateKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["key"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Key = v.(string)
											}

											if v, ok := cs["location"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Location = v.(string)
											}

											if v, ok := cs["provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Provider = v.(string)
											}

											if v, ok := cs["secret_encoding"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.SecretEncoding = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											if v, ok := cs["version"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Version = uint32(v.(int))
											}

										}

									}

									if v, ok := privateKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["name"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)
											}

										}

									}

								}

							}

						}

					}

					if w, ok := commonParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
						commonParams.TrustedCaUrl = w.(string)
					}

					if v, ok := commonParamsMapStrToI["validation_params"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						validationParams := &ves_io_schema.TlsValidationParamsType{}
						commonParams.ValidationParams = validationParams
						for _, set := range sl {

							validationParamsMapStrToI := set.(map[string]interface{})

							if w, ok := validationParamsMapStrToI["skip_hostname_verification"]; ok && !isIntfNil(w) {
								validationParams.SkipHostnameVerification = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.TrustedCaUrl = w.(string)
							}

							if w, ok := validationParamsMapStrToI["use_volterra_trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.UseVolterraTrustedCaUrl = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["verify_subject_alt_names"]; ok && !isIntfNil(w) {
								ls := make([]string, len(w.([]interface{})))
								for i, v := range w.([]interface{}) {
									ls[i] = v.(string)
								}
								validationParams.VerifySubjectAltNames = ls
							}

						}

					}

				}

			}

			if w, ok := tlsParametersMapStrToI["require_client_certificate"]; ok && !isIntfNil(w) {
				tlsParameters.RequireClientCertificate = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("user_identification"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		userIdentificationInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.UserIdentification = userIdentificationInt
		for i, ps := range sl {

			uiMapToStrVal := ps.(map[string]interface{})
			userIdentificationInt[i] = &ves_io_schema.ObjectRefType{}

			userIdentificationInt[i].Kind = "user_identification"

			if v, ok := uiMapToStrVal["name"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Name = v.(string)
			}

			if v, ok := uiMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Namespace = v.(string)
			}

			if v, ok := uiMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Tenant = v.(string)
			}

			if v, ok := uiMapToStrVal["uid"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("waf_type"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		wafType := &ves_io_schema.WafType{}
		createSpec.WafType = wafType
		for _, set := range sl {

			wafTypeMapStrToI := set.(map[string]interface{})

			refTypeTypeFound := false

			if v, ok := wafTypeMapStrToI["waf"]; ok && !isIntfNil(v) && !refTypeTypeFound {

				refTypeTypeFound = true
				refTypeInt := &ves_io_schema.WafType_Waf{}
				refTypeInt.Waf = &ves_io_schema.WafRefType{}
				wafType.RefType = refTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["waf"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						wafInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						refTypeInt.Waf.Waf = wafInt
						for i, ps := range sl {

							wMapToStrVal := ps.(map[string]interface{})
							wafInt[i] = &ves_io_schema.ObjectRefType{}

							wafInt[i].Kind = "waf"

							if v, ok := wMapToStrVal["name"]; ok && !isIntfNil(v) {
								wafInt[i].Name = v.(string)
							}

							if v, ok := wMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								wafInt[i].Namespace = v.(string)
							}

							if v, ok := wMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								wafInt[i].Tenant = v.(string)
							}

							if v, ok := wMapToStrVal["uid"]; ok && !isIntfNil(v) {
								wafInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := wafTypeMapStrToI["waf_rules"]; ok && !isIntfNil(v) && !refTypeTypeFound {

				refTypeTypeFound = true
				refTypeInt := &ves_io_schema.WafType_WafRules{}
				refTypeInt.WafRules = &ves_io_schema.WafRulesRefType{}
				wafType.RefType = refTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["waf_rules"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						wafRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						refTypeInt.WafRules.WafRules = wafRulesInt
						for i, ps := range sl {

							wrMapToStrVal := ps.(map[string]interface{})
							wafRulesInt[i] = &ves_io_schema.ObjectRefType{}

							wafRulesInt[i].Kind = "waf_rules"

							if v, ok := wrMapToStrVal["name"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Name = v.(string)
							}

							if v, ok := wrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Namespace = v.(string)
							}

							if v, ok := wrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Tenant = v.(string)
							}

							if v, ok := wrMapToStrVal["uid"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra VirtualHost object with struct: %+v", createReq)

	createVirtualHostResp, err := client.CreateObject(context.Background(), ves_io_schema_virtual_host.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating VirtualHost: %s", err)
	}
	d.SetId(createVirtualHostResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraVirtualHostRead(d, meta)
}

func resourceVolterraVirtualHostRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_virtual_host.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualHost %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualHost %q: %s", d.Id(), err)
	}
	return setVirtualHostFields(client, d, resp)
}

func setVirtualHostFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraVirtualHostUpdate updates VirtualHost resource
func resourceVolterraVirtualHostUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_virtual_host.ReplaceSpecType{}
	updateReq := &ves_io_schema_virtual_host.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("add_location"); ok && !isIntfNil(v) {

		updateSpec.AddLocation =
			v.(bool)
	}

	if v, ok := d.GetOk("advertise_policies"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		advertisePoliciesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.AdvertisePolicies = advertisePoliciesInt
		for i, ps := range sl {

			apMapToStrVal := ps.(map[string]interface{})
			advertisePoliciesInt[i] = &ves_io_schema.ObjectRefType{}

			advertisePoliciesInt[i].Kind = "advertise_policy"

			if v, ok := apMapToStrVal["name"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Name = v.(string)
			}

			if v, ok := apMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Namespace = v.(string)
			}

			if v, ok := apMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Tenant = v.(string)
			}

			if v, ok := apMapToStrVal["uid"]; ok && !isIntfNil(v) {
				advertisePoliciesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("buffer_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		bufferPolicy := &ves_io_schema.BufferConfigType{}
		updateSpec.BufferPolicy = bufferPolicy
		for _, set := range sl {

			bufferPolicyMapStrToI := set.(map[string]interface{})

			if w, ok := bufferPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
				bufferPolicy.Disabled = w.(bool)
			}

			if w, ok := bufferPolicyMapStrToI["max_request_bytes"]; ok && !isIntfNil(w) {
				bufferPolicy.MaxRequestBytes = w.(uint32)
			}

			if w, ok := bufferPolicyMapStrToI["max_request_time"]; ok && !isIntfNil(w) {
				bufferPolicy.MaxRequestTime = w.(uint32)
			}

		}

	}

	challengeTypeTypeFound := false

	if v, ok := d.GetOk("captcha_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_virtual_host.ReplaceSpecType_CaptchaChallenge{}
		challengeTypeInt.CaptchaChallenge = &ves_io_schema_virtual_host.CaptchaChallengeType{}
		updateSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_captcha_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.EnableCaptchaChallenge = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("js_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_virtual_host.ReplaceSpecType_JsChallenge{}
		challengeTypeInt.JsChallenge = &ves_io_schema_virtual_host.JavascriptChallengeType{}
		updateSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_js_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.EnableJsChallenge = v.(bool)
			}

			if v, ok := cs["js_script_delay"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.JsScriptDelay = uint32(v.(int))
			}

		}

	}

	if v, ok := d.GetOk("no_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true

		if v.(bool) {
			challengeTypeInt := &ves_io_schema_virtual_host.ReplaceSpecType_NoChallenge{}
			challengeTypeInt.NoChallenge = &ves_io_schema.Empty{}
			updateSpec.ChallengeType = challengeTypeInt
		}

	}

	if v, ok := d.GetOk("compression_params"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		compressionParams := &ves_io_schema_virtual_host.CompressionType{}
		updateSpec.CompressionParams = compressionParams
		for _, set := range sl {

			compressionParamsMapStrToI := set.(map[string]interface{})

			if w, ok := compressionParamsMapStrToI["content_length"]; ok && !isIntfNil(w) {
				compressionParams.ContentLength = w.(uint32)
			}

			if w, ok := compressionParamsMapStrToI["content_type"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				compressionParams.ContentType = ls
			}

			if w, ok := compressionParamsMapStrToI["disable_on_etag_header"]; ok && !isIntfNil(w) {
				compressionParams.DisableOnEtagHeader = w.(bool)
			}

			if w, ok := compressionParamsMapStrToI["remove_accept_encoding_header"]; ok && !isIntfNil(w) {
				compressionParams.RemoveAcceptEncodingHeader = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("cors_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		corsPolicy := &ves_io_schema.CorsPolicy{}
		updateSpec.CorsPolicy = corsPolicy
		for _, set := range sl {

			corsPolicyMapStrToI := set.(map[string]interface{})

			if w, ok := corsPolicyMapStrToI["allow_credentials"]; ok && !isIntfNil(w) {
				corsPolicy.AllowCredentials = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["allow_headers"]; ok && !isIntfNil(w) {
				corsPolicy.AllowHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_methods"]; ok && !isIntfNil(w) {
				corsPolicy.AllowMethods = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_origin"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOrigin = ls
			}

			if w, ok := corsPolicyMapStrToI["allow_origin_regex"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOriginRegex = ls
			}

			if w, ok := corsPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
				corsPolicy.Disabled = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["expose_headers"]; ok && !isIntfNil(w) {
				corsPolicy.ExposeHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["max_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaxAge = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["maximum_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaximumAge = w.(int32)
			}

		}

	}

	if v, ok := d.GetOk("disable_default_error_pages"); ok && !isIntfNil(v) {

		updateSpec.DisableDefaultErrorPages =
			v.(bool)
	}

	if v, ok := d.GetOk("disable_dns_resolve"); ok && !isIntfNil(v) {

		updateSpec.DisableDnsResolve =
			v.(bool)
	}

	if v, ok := d.GetOk("domains"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		updateSpec.Domains = ls

	}

	if v, ok := d.GetOk("dynamic_reverse_proxy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		dynamicReverseProxy := &ves_io_schema_virtual_host.DynamicReverseProxyType{}
		updateSpec.DynamicReverseProxy = dynamicReverseProxy
		for _, set := range sl {

			dynamicReverseProxyMapStrToI := set.(map[string]interface{})

			if v, ok := dynamicReverseProxyMapStrToI["resolution_network"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				resolutionNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				dynamicReverseProxy.ResolutionNetwork = resolutionNetworkInt
				for i, ps := range sl {

					rnMapToStrVal := ps.(map[string]interface{})
					resolutionNetworkInt[i] = &ves_io_schema.ObjectRefType{}

					resolutionNetworkInt[i].Kind = "virtual_network"

					if v, ok := rnMapToStrVal["name"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Name = v.(string)
					}

					if v, ok := rnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Namespace = v.(string)
					}

					if v, ok := rnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Tenant = v.(string)
					}

					if v, ok := rnMapToStrVal["uid"]; ok && !isIntfNil(v) {
						resolutionNetworkInt[i].Uid = v.(string)
					}

				}

			}

			if v, ok := dynamicReverseProxyMapStrToI["resolution_network_type"]; ok && !isIntfNil(v) {

				dynamicReverseProxy.ResolutionNetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

			}

			if w, ok := dynamicReverseProxyMapStrToI["resolve_endpoint_dynamically"]; ok && !isIntfNil(w) {
				dynamicReverseProxy.ResolveEndpointDynamically = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("idle_timeout"); ok && !isIntfNil(v) {

		updateSpec.IdleTimeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("max_request_header_size"); ok && !isIntfNil(v) {

		updateSpec.MaxRequestHeaderSize =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("proxy"); ok && !isIntfNil(v) {

		updateSpec.Proxy = ves_io_schema_virtual_host.ProxyType(ves_io_schema_virtual_host.ProxyType_value[v.(string)])

	}

	if v, ok := d.GetOk("rate_limiter"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rateLimiterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.RateLimiter = rateLimiterInt
		for i, ps := range sl {

			rlMapToStrVal := ps.(map[string]interface{})
			rateLimiterInt[i] = &ves_io_schema.ObjectRefType{}

			rateLimiterInt[i].Kind = "rate_limiter"

			if v, ok := rlMapToStrVal["name"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Name = v.(string)
			}

			if v, ok := rlMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Namespace = v.(string)
			}

			if v, ok := rlMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Tenant = v.(string)
			}

			if v, ok := rlMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("rate_limiter_allowed_prefixes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rateLimiterAllowedPrefixesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.RateLimiterAllowedPrefixes = rateLimiterAllowedPrefixesInt
		for i, ps := range sl {

			rlapMapToStrVal := ps.(map[string]interface{})
			rateLimiterAllowedPrefixesInt[i] = &ves_io_schema.ObjectRefType{}

			rateLimiterAllowedPrefixesInt[i].Kind = "ip_prefix_set"

			if v, ok := rlapMapToStrVal["name"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Name = v.(string)
			}

			if v, ok := rlapMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Namespace = v.(string)
			}

			if v, ok := rlapMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Tenant = v.(string)
			}

			if v, ok := rlapMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rateLimiterAllowedPrefixesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("request_headers_to_add"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		requestHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
		updateSpec.RequestHeadersToAdd = requestHeadersToAdd
		for i, set := range sl {
			requestHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

			requestHeadersToAddMapStrToI := set.(map[string]interface{})

			if w, ok := requestHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
				requestHeadersToAdd[i].Append = w.(bool)
			}

			if w, ok := requestHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
				requestHeadersToAdd[i].Name = w.(string)
			}

			if w, ok := requestHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
				requestHeadersToAdd[i].Value = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("request_headers_to_remove"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		updateSpec.RequestHeadersToRemove = ls

	}

	if v, ok := d.GetOk("response_headers_to_add"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		responseHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
		updateSpec.ResponseHeadersToAdd = responseHeadersToAdd
		for i, set := range sl {
			responseHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

			responseHeadersToAddMapStrToI := set.(map[string]interface{})

			if w, ok := responseHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
				responseHeadersToAdd[i].Append = w.(bool)
			}

			if w, ok := responseHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
				responseHeadersToAdd[i].Name = w.(string)
			}

			if w, ok := responseHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
				responseHeadersToAdd[i].Value = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("response_headers_to_remove"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		updateSpec.ResponseHeadersToRemove = ls

	}

	if v, ok := d.GetOk("retry_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		retryPolicy := &ves_io_schema.RetryPolicyType{}
		updateSpec.RetryPolicy = retryPolicy
		for _, set := range sl {

			retryPolicyMapStrToI := set.(map[string]interface{})

			if v, ok := retryPolicyMapStrToI["back_off"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				backOff := &ves_io_schema.RetryBackOff{}
				retryPolicy.BackOff = backOff
				for _, set := range sl {

					backOffMapStrToI := set.(map[string]interface{})

					if w, ok := backOffMapStrToI["base_interval"]; ok && !isIntfNil(w) {
						backOff.BaseInterval = w.(uint32)
					}

					if w, ok := backOffMapStrToI["max_interval"]; ok && !isIntfNil(w) {
						backOff.MaxInterval = w.(uint32)
					}

				}

			}

			if w, ok := retryPolicyMapStrToI["num_retries"]; ok && !isIntfNil(w) {
				retryPolicy.NumRetries = w.(uint32)
			}

			if w, ok := retryPolicyMapStrToI["per_try_timeout"]; ok && !isIntfNil(w) {
				retryPolicy.PerTryTimeout = w.(uint32)
			}

			if w, ok := retryPolicyMapStrToI["retriable_status_codes"]; ok && !isIntfNil(w) {
				ls := make([]uint32, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				retryPolicy.RetriableStatusCodes = ls
			}

			if w, ok := retryPolicyMapStrToI["retry_on"]; ok && !isIntfNil(w) {
				retryPolicy.RetryOn = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Routes = routesInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			routesInt[i] = &ves_io_schema.ObjectRefType{}

			routesInt[i].Kind = "route"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				routesInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				routesInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				routesInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				routesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("temporary_user_blocking"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		temporaryUserBlocking := &ves_io_schema_virtual_host.TemporaryUserBlockingType{}
		updateSpec.TemporaryUserBlocking = temporaryUserBlocking
		for _, set := range sl {

			temporaryUserBlockingMapStrToI := set.(map[string]interface{})

			if w, ok := temporaryUserBlockingMapStrToI["custom_page"]; ok && !isIntfNil(w) {
				temporaryUserBlocking.CustomPage = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("tls_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsParameters := &ves_io_schema.DownstreamTlsParamsType{}
		updateSpec.TlsParameters = tlsParameters
		for _, set := range sl {

			tlsParametersMapStrToI := set.(map[string]interface{})

			if v, ok := tlsParametersMapStrToI["common_params"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				commonParams := &ves_io_schema.TlsParamsType{}
				tlsParameters.CommonParams = commonParams
				for _, set := range sl {

					commonParamsMapStrToI := set.(map[string]interface{})

					if w, ok := commonParamsMapStrToI["cipher_suites"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						commonParams.CipherSuites = ls
					}

					if v, ok := commonParamsMapStrToI["maximum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MaximumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["minimum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MinimumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						commonParams.TlsCertificates = tlsCertificates
						for i, set := range sl {
							tlsCertificates[i] = &ves_io_schema.TlsCertificateType{}

							tlsCertificatesMapStrToI := set.(map[string]interface{})

							if w, ok := tlsCertificatesMapStrToI["certificate_url"]; ok && !isIntfNil(w) {
								tlsCertificates[i].CertificateUrl = w.(string)
							}

							if w, ok := tlsCertificatesMapStrToI["description"]; ok && !isIntfNil(w) {
								tlsCertificates[i].Description = w.(string)
							}

							if v, ok := tlsCertificatesMapStrToI["private_key"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								privateKey := &ves_io_schema.SecretType{}
								tlsCertificates[i].PrivateKey = privateKey
								for _, set := range sl {

									privateKeyMapStrToI := set.(map[string]interface{})

									if v, ok := privateKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										privateKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := privateKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["decryption_provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.BlindfoldSecretInfo.DecryptionProvider = v.(string)
											}

											if v, ok := cs["location"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.BlindfoldSecretInfo.Location = v.(string)
											}

											if v, ok := cs["store_provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.BlindfoldSecretInfo.StoreProvider = v.(string)
											}

										}

									}

									if v, ok := privateKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
											}

											if v, ok := cs["url"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
											}

										}

									}

									if v, ok := privateKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["key"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Key = v.(string)
											}

											if v, ok := cs["location"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Location = v.(string)
											}

											if v, ok := cs["provider"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Provider = v.(string)
											}

											if v, ok := cs["secret_encoding"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.SecretEncoding = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											if v, ok := cs["version"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.VaultSecretInfo.Version = uint32(v.(int))
											}

										}

									}

									if v, ok := privateKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["name"]; ok && !isIntfNil(v) {

												secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)
											}

										}

									}

								}

							}

						}

					}

					if w, ok := commonParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
						commonParams.TrustedCaUrl = w.(string)
					}

					if v, ok := commonParamsMapStrToI["validation_params"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						validationParams := &ves_io_schema.TlsValidationParamsType{}
						commonParams.ValidationParams = validationParams
						for _, set := range sl {

							validationParamsMapStrToI := set.(map[string]interface{})

							if w, ok := validationParamsMapStrToI["skip_hostname_verification"]; ok && !isIntfNil(w) {
								validationParams.SkipHostnameVerification = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.TrustedCaUrl = w.(string)
							}

							if w, ok := validationParamsMapStrToI["use_volterra_trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.UseVolterraTrustedCaUrl = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["verify_subject_alt_names"]; ok && !isIntfNil(w) {
								ls := make([]string, len(w.([]interface{})))
								for i, v := range w.([]interface{}) {
									ls[i] = v.(string)
								}
								validationParams.VerifySubjectAltNames = ls
							}

						}

					}

				}

			}

			if w, ok := tlsParametersMapStrToI["require_client_certificate"]; ok && !isIntfNil(w) {
				tlsParameters.RequireClientCertificate = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("user_identification"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		userIdentificationInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.UserIdentification = userIdentificationInt
		for i, ps := range sl {

			uiMapToStrVal := ps.(map[string]interface{})
			userIdentificationInt[i] = &ves_io_schema.ObjectRefType{}

			userIdentificationInt[i].Kind = "user_identification"

			if v, ok := uiMapToStrVal["name"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Name = v.(string)
			}

			if v, ok := uiMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Namespace = v.(string)
			}

			if v, ok := uiMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Tenant = v.(string)
			}

			if v, ok := uiMapToStrVal["uid"]; ok && !isIntfNil(v) {
				userIdentificationInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("waf_type"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		wafType := &ves_io_schema.WafType{}
		updateSpec.WafType = wafType
		for _, set := range sl {

			wafTypeMapStrToI := set.(map[string]interface{})

			refTypeTypeFound := false

			if v, ok := wafTypeMapStrToI["waf"]; ok && !isIntfNil(v) && !refTypeTypeFound {

				refTypeTypeFound = true
				refTypeInt := &ves_io_schema.WafType_Waf{}
				refTypeInt.Waf = &ves_io_schema.WafRefType{}
				wafType.RefType = refTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["waf"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						wafInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						refTypeInt.Waf.Waf = wafInt
						for i, ps := range sl {

							wMapToStrVal := ps.(map[string]interface{})
							wafInt[i] = &ves_io_schema.ObjectRefType{}

							wafInt[i].Kind = "waf"

							if v, ok := wMapToStrVal["name"]; ok && !isIntfNil(v) {
								wafInt[i].Name = v.(string)
							}

							if v, ok := wMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								wafInt[i].Namespace = v.(string)
							}

							if v, ok := wMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								wafInt[i].Tenant = v.(string)
							}

							if v, ok := wMapToStrVal["uid"]; ok && !isIntfNil(v) {
								wafInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := wafTypeMapStrToI["waf_rules"]; ok && !isIntfNil(v) && !refTypeTypeFound {

				refTypeTypeFound = true
				refTypeInt := &ves_io_schema.WafType_WafRules{}
				refTypeInt.WafRules = &ves_io_schema.WafRulesRefType{}
				wafType.RefType = refTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["waf_rules"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						wafRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						refTypeInt.WafRules.WafRules = wafRulesInt
						for i, ps := range sl {

							wrMapToStrVal := ps.(map[string]interface{})
							wafRulesInt[i] = &ves_io_schema.ObjectRefType{}

							wafRulesInt[i].Kind = "waf_rules"

							if v, ok := wrMapToStrVal["name"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Name = v.(string)
							}

							if v, ok := wrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Namespace = v.(string)
							}

							if v, ok := wrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Tenant = v.(string)
							}

							if v, ok := wrMapToStrVal["uid"]; ok && !isIntfNil(v) {
								wafRulesInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra VirtualHost obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_virtual_host.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating VirtualHost: %s", err)
	}

	return resourceVolterraVirtualHostRead(d, meta)
}

func resourceVolterraVirtualHostDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_virtual_host.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualHost %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualHost before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra VirtualHost obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_virtual_host.ObjectType, namespace, name)
}