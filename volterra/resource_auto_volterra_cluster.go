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
	ves_io_schema_cluster "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cluster"
)

// resourceVolterraCluster is implementation of Volterra's Cluster resources
func resourceVolterraCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraClusterCreate,
		Read:   resourceVolterraClusterRead,
		Update: resourceVolterraClusterUpdate,
		Delete: resourceVolterraClusterDelete,

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

			"circuit_breaker": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"max_requests": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"pending_requests": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"priority": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"retries": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"connection_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"default_subset": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"endpoint_selection": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"endpoint_subsets": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"keys": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"endpoints": {

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

			"fallback_policy": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"health_checks": {

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

			"http2_options": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"http_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"loadbalancer_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"outlier_detection": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"base_ejection_time": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"consecutive_5xx": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"consecutive_gateway_failure": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"interval": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"max_ejection_percent": {
							Type:     schema.TypeInt,
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

						"sni": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraClusterCreate creates Cluster resource
func resourceVolterraClusterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_cluster.CreateSpecType{}
	createReq := &ves_io_schema_cluster.CreateRequest{
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

	if v, ok := d.GetOk("circuit_breaker"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		circuitBreaker := &ves_io_schema_cluster.CircuitBreaker{}
		createSpec.CircuitBreaker = circuitBreaker
		for _, set := range sl {

			circuitBreakerMapStrToI := set.(map[string]interface{})

			if w, ok := circuitBreakerMapStrToI["connection_limit"]; ok && !isIntfNil(w) {
				circuitBreaker.ConnectionLimit = w.(uint32)
			}

			if w, ok := circuitBreakerMapStrToI["max_requests"]; ok && !isIntfNil(w) {
				circuitBreaker.MaxRequests = w.(uint32)
			}

			if w, ok := circuitBreakerMapStrToI["pending_requests"]; ok && !isIntfNil(w) {
				circuitBreaker.PendingRequests = w.(uint32)
			}

			if v, ok := circuitBreakerMapStrToI["priority"]; ok && !isIntfNil(v) {

				circuitBreaker.Priority = ves_io_schema.RoutingPriority(ves_io_schema.RoutingPriority_value[v.(string)])

			}

			if w, ok := circuitBreakerMapStrToI["retries"]; ok && !isIntfNil(w) {
				circuitBreaker.Retries = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("connection_timeout"); ok && !isIntfNil(v) {

		createSpec.ConnectionTimeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("default_subset"); ok && !isIntfNil(v) {

		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
		}
		createSpec.DefaultSubset = ms
	}

	if v, ok := d.GetOk("endpoint_selection"); ok && !isIntfNil(v) {

		createSpec.EndpointSelection = ves_io_schema_cluster.EndpointSelectionPolicy(ves_io_schema_cluster.EndpointSelectionPolicy_value[v.(string)])

	}

	if v, ok := d.GetOk("endpoint_subsets"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		endpointSubsets := make([]*ves_io_schema_cluster.EndpointSubsetSelectorType, len(sl))
		createSpec.EndpointSubsets = endpointSubsets
		for i, set := range sl {
			endpointSubsets[i] = &ves_io_schema_cluster.EndpointSubsetSelectorType{}

			endpointSubsetsMapStrToI := set.(map[string]interface{})

			if w, ok := endpointSubsetsMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				endpointSubsets[i].Keys = ls
			}

		}

	}

	if v, ok := d.GetOk("endpoints"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		endpointsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Endpoints = endpointsInt
		for i, ps := range sl {

			eMapToStrVal := ps.(map[string]interface{})
			endpointsInt[i] = &ves_io_schema.ObjectRefType{}

			endpointsInt[i].Kind = "endpoint"

			if v, ok := eMapToStrVal["name"]; ok && !isIntfNil(v) {
				endpointsInt[i].Name = v.(string)
			}

			if v, ok := eMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				endpointsInt[i].Namespace = v.(string)
			}

			if v, ok := eMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				endpointsInt[i].Tenant = v.(string)
			}

			if v, ok := eMapToStrVal["uid"]; ok && !isIntfNil(v) {
				endpointsInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("fallback_policy"); ok && !isIntfNil(v) {

		createSpec.FallbackPolicy = ves_io_schema_cluster.SubsetFallbackPolicy(ves_io_schema_cluster.SubsetFallbackPolicy_value[v.(string)])

	}

	if v, ok := d.GetOk("health_checks"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		healthChecksInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.HealthChecks = healthChecksInt
		for i, ps := range sl {

			hcMapToStrVal := ps.(map[string]interface{})
			healthChecksInt[i] = &ves_io_schema.ObjectRefType{}

			healthChecksInt[i].Kind = "healthcheck"

			if v, ok := hcMapToStrVal["name"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Name = v.(string)
			}

			if v, ok := hcMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Namespace = v.(string)
			}

			if v, ok := hcMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Tenant = v.(string)
			}

			if v, ok := hcMapToStrVal["uid"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("http2_options"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		http2Options := &ves_io_schema_cluster.Http2ProtocolOptions{}
		createSpec.Http2Options = http2Options
		for _, set := range sl {

			http2OptionsMapStrToI := set.(map[string]interface{})

			if w, ok := http2OptionsMapStrToI["enabled"]; ok && !isIntfNil(w) {
				http2Options.Enabled = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("http_idle_timeout"); ok && !isIntfNil(v) {

		createSpec.HttpIdleTimeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("loadbalancer_algorithm"); ok && !isIntfNil(v) {

		createSpec.LoadbalancerAlgorithm = ves_io_schema_cluster.LoadbalancerAlgorithm(ves_io_schema_cluster.LoadbalancerAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("outlier_detection"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		outlierDetection := &ves_io_schema_cluster.OutlierDetectionType{}
		createSpec.OutlierDetection = outlierDetection
		for _, set := range sl {

			outlierDetectionMapStrToI := set.(map[string]interface{})

			if w, ok := outlierDetectionMapStrToI["base_ejection_time"]; ok && !isIntfNil(w) {
				outlierDetection.BaseEjectionTime = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["consecutive_5xx"]; ok && !isIntfNil(w) {
				outlierDetection.Consecutive_5Xx = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["consecutive_gateway_failure"]; ok && !isIntfNil(w) {
				outlierDetection.ConsecutiveGatewayFailure = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["interval"]; ok && !isIntfNil(w) {
				outlierDetection.Interval = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["max_ejection_percent"]; ok && !isIntfNil(w) {
				outlierDetection.MaxEjectionPercent = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("tls_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsParameters := &ves_io_schema.UpstreamTlsParamsType{}
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

			if w, ok := tlsParametersMapStrToI["sni"]; ok && !isIntfNil(w) {
				tlsParameters.Sni = w.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra Cluster object with struct: %+v", createReq)

	createClusterResp, err := client.CreateObject(context.Background(), ves_io_schema_cluster.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Cluster: %s", err)
	}
	d.SetId(createClusterResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraClusterRead(d, meta)
}

func resourceVolterraClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_cluster.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Cluster %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Cluster %q: %s", d.Id(), err)
	}
	return setClusterFields(client, d, resp)
}

func setClusterFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraClusterUpdate updates Cluster resource
func resourceVolterraClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_cluster.ReplaceSpecType{}
	updateReq := &ves_io_schema_cluster.ReplaceRequest{
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

	if v, ok := d.GetOk("circuit_breaker"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		circuitBreaker := &ves_io_schema_cluster.CircuitBreaker{}
		updateSpec.CircuitBreaker = circuitBreaker
		for _, set := range sl {

			circuitBreakerMapStrToI := set.(map[string]interface{})

			if w, ok := circuitBreakerMapStrToI["connection_limit"]; ok && !isIntfNil(w) {
				circuitBreaker.ConnectionLimit = w.(uint32)
			}

			if w, ok := circuitBreakerMapStrToI["max_requests"]; ok && !isIntfNil(w) {
				circuitBreaker.MaxRequests = w.(uint32)
			}

			if w, ok := circuitBreakerMapStrToI["pending_requests"]; ok && !isIntfNil(w) {
				circuitBreaker.PendingRequests = w.(uint32)
			}

			if v, ok := circuitBreakerMapStrToI["priority"]; ok && !isIntfNil(v) {

				circuitBreaker.Priority = ves_io_schema.RoutingPriority(ves_io_schema.RoutingPriority_value[v.(string)])

			}

			if w, ok := circuitBreakerMapStrToI["retries"]; ok && !isIntfNil(w) {
				circuitBreaker.Retries = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("connection_timeout"); ok && !isIntfNil(v) {

		updateSpec.ConnectionTimeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("default_subset"); ok && !isIntfNil(v) {

		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
		}
		updateSpec.DefaultSubset = ms
	}

	if v, ok := d.GetOk("endpoint_selection"); ok && !isIntfNil(v) {

		updateSpec.EndpointSelection = ves_io_schema_cluster.EndpointSelectionPolicy(ves_io_schema_cluster.EndpointSelectionPolicy_value[v.(string)])

	}

	if v, ok := d.GetOk("endpoint_subsets"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		endpointSubsets := make([]*ves_io_schema_cluster.EndpointSubsetSelectorType, len(sl))
		updateSpec.EndpointSubsets = endpointSubsets
		for i, set := range sl {
			endpointSubsets[i] = &ves_io_schema_cluster.EndpointSubsetSelectorType{}

			endpointSubsetsMapStrToI := set.(map[string]interface{})

			if w, ok := endpointSubsetsMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				endpointSubsets[i].Keys = ls
			}

		}

	}

	if v, ok := d.GetOk("endpoints"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		endpointsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Endpoints = endpointsInt
		for i, ps := range sl {

			eMapToStrVal := ps.(map[string]interface{})
			endpointsInt[i] = &ves_io_schema.ObjectRefType{}

			endpointsInt[i].Kind = "endpoint"

			if v, ok := eMapToStrVal["name"]; ok && !isIntfNil(v) {
				endpointsInt[i].Name = v.(string)
			}

			if v, ok := eMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				endpointsInt[i].Namespace = v.(string)
			}

			if v, ok := eMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				endpointsInt[i].Tenant = v.(string)
			}

			if v, ok := eMapToStrVal["uid"]; ok && !isIntfNil(v) {
				endpointsInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("fallback_policy"); ok && !isIntfNil(v) {

		updateSpec.FallbackPolicy = ves_io_schema_cluster.SubsetFallbackPolicy(ves_io_schema_cluster.SubsetFallbackPolicy_value[v.(string)])

	}

	if v, ok := d.GetOk("health_checks"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		healthChecksInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.HealthChecks = healthChecksInt
		for i, ps := range sl {

			hcMapToStrVal := ps.(map[string]interface{})
			healthChecksInt[i] = &ves_io_schema.ObjectRefType{}

			healthChecksInt[i].Kind = "healthcheck"

			if v, ok := hcMapToStrVal["name"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Name = v.(string)
			}

			if v, ok := hcMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Namespace = v.(string)
			}

			if v, ok := hcMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Tenant = v.(string)
			}

			if v, ok := hcMapToStrVal["uid"]; ok && !isIntfNil(v) {
				healthChecksInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("http2_options"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		http2Options := &ves_io_schema_cluster.Http2ProtocolOptions{}
		updateSpec.Http2Options = http2Options
		for _, set := range sl {

			http2OptionsMapStrToI := set.(map[string]interface{})

			if w, ok := http2OptionsMapStrToI["enabled"]; ok && !isIntfNil(w) {
				http2Options.Enabled = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("http_idle_timeout"); ok && !isIntfNil(v) {

		updateSpec.HttpIdleTimeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("loadbalancer_algorithm"); ok && !isIntfNil(v) {

		updateSpec.LoadbalancerAlgorithm = ves_io_schema_cluster.LoadbalancerAlgorithm(ves_io_schema_cluster.LoadbalancerAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("outlier_detection"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		outlierDetection := &ves_io_schema_cluster.OutlierDetectionType{}
		updateSpec.OutlierDetection = outlierDetection
		for _, set := range sl {

			outlierDetectionMapStrToI := set.(map[string]interface{})

			if w, ok := outlierDetectionMapStrToI["base_ejection_time"]; ok && !isIntfNil(w) {
				outlierDetection.BaseEjectionTime = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["consecutive_5xx"]; ok && !isIntfNil(w) {
				outlierDetection.Consecutive_5Xx = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["consecutive_gateway_failure"]; ok && !isIntfNil(w) {
				outlierDetection.ConsecutiveGatewayFailure = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["interval"]; ok && !isIntfNil(w) {
				outlierDetection.Interval = w.(uint32)
			}

			if w, ok := outlierDetectionMapStrToI["max_ejection_percent"]; ok && !isIntfNil(w) {
				outlierDetection.MaxEjectionPercent = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("tls_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsParameters := &ves_io_schema.UpstreamTlsParamsType{}
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

			if w, ok := tlsParametersMapStrToI["sni"]; ok && !isIntfNil(w) {
				tlsParameters.Sni = w.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra Cluster obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_cluster.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Cluster: %s", err)
	}

	return resourceVolterraClusterRead(d, meta)
}

func resourceVolterraClusterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_cluster.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Cluster %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Cluster before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Cluster obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_cluster.ObjectType, namespace, name)
}
