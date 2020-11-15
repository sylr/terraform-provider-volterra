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
	ves_io_schema_discovery "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/discovery"
)

// resourceVolterraDiscovery is implementation of Volterra's Discovery resources
func resourceVolterraDiscovery() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraDiscoveryCreate,
		Read:   resourceVolterraDiscoveryRead,
		Update: resourceVolterraDiscoveryUpdate,
		Delete: resourceVolterraDiscoveryDelete,

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

			"discovery_consul": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connection_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"api_server": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"tls_info": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ca_certificate_url": {

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

															"certificate": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"certificate_url": {

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

															"key_url": {

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

															"server_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"trusted_ca_url": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"http_basic_auth_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"passwd_url": {

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

												"user_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"scheme": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"publish_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"publish": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"discovery_k8s": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connection_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"api_server": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"tls_info": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ca_certificate_url": {

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

															"certificate": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"certificate_url": {

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

															"key_url": {

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

															"server_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"trusted_ca_url": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"in_cluster": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"kubeconfig_url": {

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

									"isolated": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"reachable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"publish_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"dns_delegation": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dns_mode": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"subdomain": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"publish": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"publish_fqdns": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"where": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"site": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"ref": {

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

						"virtual_network": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ref": {

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

						"virtual_site": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"ref": {

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

// resourceVolterraDiscoveryCreate creates Discovery resource
func resourceVolterraDiscoveryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_discovery.CreateSpecType{}
	createReq := &ves_io_schema_discovery.CreateRequest{
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

	discoveryChoiceTypeFound := false

	if v, ok := d.GetOk("discovery_consul"); ok && !discoveryChoiceTypeFound {

		discoveryChoiceTypeFound = true
		discoveryChoiceInt := &ves_io_schema_discovery.CreateSpecType_DiscoveryConsul{}
		discoveryChoiceInt.DiscoveryConsul = &ves_io_schema_discovery.ConsulDiscoveryType{}
		createSpec.DiscoveryChoice = discoveryChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["access_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				accessInfo := &ves_io_schema_discovery.ConsulAccessInfo{}
				discoveryChoiceInt.DiscoveryConsul.AccessInfo = accessInfo
				for _, set := range sl {

					accessInfoMapStrToI := set.(map[string]interface{})

					if v, ok := accessInfoMapStrToI["connection_info"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						connectionInfo := &ves_io_schema_discovery.RestConfigType{}
						accessInfo.ConnectionInfo = connectionInfo
						for _, set := range sl {

							connectionInfoMapStrToI := set.(map[string]interface{})

							if w, ok := connectionInfoMapStrToI["api_server"]; ok && !isIntfNil(w) {
								connectionInfo.ApiServer = w.(string)
							}

							if v, ok := connectionInfoMapStrToI["tls_info"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								tlsInfo := &ves_io_schema_discovery.TLSClientConfigType{}
								connectionInfo.TlsInfo = tlsInfo
								for _, set := range sl {

									tlsInfoMapStrToI := set.(map[string]interface{})

									if v, ok := tlsInfoMapStrToI["ca_certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										caCertificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CaCertificateUrl = caCertificateUrl
										for _, set := range sl {

											caCertificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := caCertificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												caCertificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := caCertificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["certificate"]; ok && !isIntfNil(w) {
										tlsInfo.Certificate = w.(string)
									}

									if v, ok := tlsInfoMapStrToI["certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										certificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CertificateUrl = certificateUrl
										for _, set := range sl {

											certificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := certificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												certificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := certificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := tlsInfoMapStrToI["key_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										keyUrl := &ves_io_schema.SecretType{}
										tlsInfo.KeyUrl = keyUrl
										for _, set := range sl {

											keyUrlMapStrToI := set.(map[string]interface{})

											if v, ok := keyUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												keyUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := keyUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["server_name"]; ok && !isIntfNil(w) {
										tlsInfo.ServerName = w.(string)
									}

									if w, ok := tlsInfoMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
										tlsInfo.TrustedCaUrl = w.(string)
									}

								}

							}

						}

					}

					if v, ok := accessInfoMapStrToI["http_basic_auth_info"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						httpBasicAuthInfo := &ves_io_schema_discovery.ConsulHttpBasicAuthInfoType{}
						accessInfo.HttpBasicAuthInfo = httpBasicAuthInfo
						for _, set := range sl {

							httpBasicAuthInfoMapStrToI := set.(map[string]interface{})

							if v, ok := httpBasicAuthInfoMapStrToI["passwd_url"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								passwdUrl := &ves_io_schema.SecretType{}
								httpBasicAuthInfo.PasswdUrl = passwdUrl
								for _, set := range sl {

									passwdUrlMapStrToI := set.(map[string]interface{})

									if v, ok := passwdUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										passwdUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := passwdUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwdUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwdUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwdUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

							if w, ok := httpBasicAuthInfoMapStrToI["user_name"]; ok && !isIntfNil(w) {
								httpBasicAuthInfo.UserName = w.(string)
							}

						}

					}

					if v, ok := accessInfoMapStrToI["scheme"]; ok && !isIntfNil(v) {

						accessInfo.Scheme = ves_io_schema_discovery.SchemeType(ves_io_schema_discovery.SchemeType_value[v.(string)])

					}

				}

			}

			if v, ok := cs["publish_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				publishInfo := &ves_io_schema_discovery.ConsulVipDiscoveryInfoType{}
				discoveryChoiceInt.DiscoveryConsul.PublishInfo = publishInfo
				for _, set := range sl {

					publishInfoMapStrToI := set.(map[string]interface{})

					publishChoiceTypeFound := false

					if v, ok := publishInfoMapStrToI["disable"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.ConsulVipDiscoveryInfoType_Disable{}
							publishChoiceInt.Disable = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

					if v, ok := publishInfoMapStrToI["publish"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.ConsulVipDiscoveryInfoType_Publish{}
							publishChoiceInt.Publish = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("discovery_k8s"); ok && !discoveryChoiceTypeFound {

		discoveryChoiceTypeFound = true
		discoveryChoiceInt := &ves_io_schema_discovery.CreateSpecType_DiscoveryK8S{}
		discoveryChoiceInt.DiscoveryK8S = &ves_io_schema_discovery.K8SDiscoveryType{}
		createSpec.DiscoveryChoice = discoveryChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["access_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				accessInfo := &ves_io_schema_discovery.K8SAccessInfo{}
				discoveryChoiceInt.DiscoveryK8S.AccessInfo = accessInfo
				for _, set := range sl {

					accessInfoMapStrToI := set.(map[string]interface{})

					configTypeTypeFound := false

					if v, ok := accessInfoMapStrToI["connection_info"]; ok && !isIntfNil(v) && !configTypeTypeFound {

						configTypeTypeFound = true
						configTypeInt := &ves_io_schema_discovery.K8SAccessInfo_ConnectionInfo{}
						configTypeInt.ConnectionInfo = &ves_io_schema_discovery.RestConfigType{}
						accessInfo.ConfigType = configTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["api_server"]; ok && !isIntfNil(v) {

								configTypeInt.ConnectionInfo.ApiServer = v.(string)
							}

							if v, ok := cs["tls_info"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								tlsInfo := &ves_io_schema_discovery.TLSClientConfigType{}
								configTypeInt.ConnectionInfo.TlsInfo = tlsInfo
								for _, set := range sl {

									tlsInfoMapStrToI := set.(map[string]interface{})

									if v, ok := tlsInfoMapStrToI["ca_certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										caCertificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CaCertificateUrl = caCertificateUrl
										for _, set := range sl {

											caCertificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := caCertificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												caCertificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := caCertificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["certificate"]; ok && !isIntfNil(w) {
										tlsInfo.Certificate = w.(string)
									}

									if v, ok := tlsInfoMapStrToI["certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										certificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CertificateUrl = certificateUrl
										for _, set := range sl {

											certificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := certificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												certificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := certificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := tlsInfoMapStrToI["key_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										keyUrl := &ves_io_schema.SecretType{}
										tlsInfo.KeyUrl = keyUrl
										for _, set := range sl {

											keyUrlMapStrToI := set.(map[string]interface{})

											if v, ok := keyUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												keyUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := keyUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["server_name"]; ok && !isIntfNil(w) {
										tlsInfo.ServerName = w.(string)
									}

									if w, ok := tlsInfoMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
										tlsInfo.TrustedCaUrl = w.(string)
									}

								}

							}

						}

					}

					if v, ok := accessInfoMapStrToI["in_cluster"]; ok && !isIntfNil(v) && !configTypeTypeFound {

						configTypeTypeFound = true
						configTypeInt := &ves_io_schema_discovery.K8SAccessInfo_InCluster{}

						accessInfo.ConfigType = configTypeInt

						configTypeInt.InCluster =
							v.(bool)

					}

					if v, ok := accessInfoMapStrToI["kubeconfig_url"]; ok && !isIntfNil(v) && !configTypeTypeFound {

						configTypeTypeFound = true
						configTypeInt := &ves_io_schema_discovery.K8SAccessInfo_KubeconfigUrl{}
						configTypeInt.KubeconfigUrl = &ves_io_schema.SecretType{}
						accessInfo.ConfigType = configTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["secret_encoding_type"]; ok && !isIntfNil(v) {

								configTypeInt.KubeconfigUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

							}

							secretInfoOneofTypeFound := false

							if v, ok := cs["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
								secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
								secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
								secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
								secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

					k8SPodNetworkChoiceTypeFound := false

					if v, ok := accessInfoMapStrToI["isolated"]; ok && !isIntfNil(v) && !k8SPodNetworkChoiceTypeFound {

						k8SPodNetworkChoiceTypeFound = true

						if v.(bool) {
							k8SPodNetworkChoiceInt := &ves_io_schema_discovery.K8SAccessInfo_Isolated{}
							k8SPodNetworkChoiceInt.Isolated = &ves_io_schema.Empty{}
							accessInfo.K8SPodNetworkChoice = k8SPodNetworkChoiceInt
						}

					}

					if v, ok := accessInfoMapStrToI["reachable"]; ok && !isIntfNil(v) && !k8SPodNetworkChoiceTypeFound {

						k8SPodNetworkChoiceTypeFound = true

						if v.(bool) {
							k8SPodNetworkChoiceInt := &ves_io_schema_discovery.K8SAccessInfo_Reachable{}
							k8SPodNetworkChoiceInt.Reachable = &ves_io_schema.Empty{}
							accessInfo.K8SPodNetworkChoice = k8SPodNetworkChoiceInt
						}

					}

				}

			}

			if v, ok := cs["publish_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				publishInfo := &ves_io_schema_discovery.K8SVipDiscoveryInfoType{}
				discoveryChoiceInt.DiscoveryK8S.PublishInfo = publishInfo
				for _, set := range sl {

					publishInfoMapStrToI := set.(map[string]interface{})

					publishChoiceTypeFound := false

					if v, ok := publishInfoMapStrToI["disable"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_Disable{}
							publishChoiceInt.Disable = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

					if v, ok := publishInfoMapStrToI["dns_delegation"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true
						publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_DnsDelegation{}
						publishChoiceInt.DnsDelegation = &ves_io_schema_discovery.K8SDelegationType{}
						publishInfo.PublishChoice = publishChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["dns_mode"]; ok && !isIntfNil(v) {

								publishChoiceInt.DnsDelegation.DnsMode = ves_io_schema_discovery.K8SDNSMode(ves_io_schema_discovery.K8SDNSMode_value[v.(string)])

							}

							if v, ok := cs["subdomain"]; ok && !isIntfNil(v) {

								publishChoiceInt.DnsDelegation.Subdomain = v.(string)
							}

						}

					}

					if v, ok := publishInfoMapStrToI["publish"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true
						publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_Publish{}
						publishChoiceInt.Publish = &ves_io_schema_discovery.K8SPublishType{}
						publishInfo.PublishChoice = publishChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								publishChoiceInt.Publish.Namespace = v.(string)
							}

						}

					}

					if v, ok := publishInfoMapStrToI["publish_fqdns"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_PublishFqdns{}
							publishChoiceInt.PublishFqdns = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.NetworkSiteRefSelector{}
		createSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_Site{}
				refOrSelectorInt.Site = &ves_io_schema.SiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.Site.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.Site.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_network"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualNetwork{}
				refOrSelectorInt.VirtualNetwork = &ves_io_schema.NetworkRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualNetwork.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_network"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualSite{}
				refOrSelectorInt.VirtualSite = &ves_io_schema.VSiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.VirtualSite.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualSite.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra Discovery object with struct: %+v", createReq)

	createDiscoveryResp, err := client.CreateObject(context.Background(), ves_io_schema_discovery.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Discovery: %s", err)
	}
	d.SetId(createDiscoveryResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraDiscoveryRead(d, meta)
}

func resourceVolterraDiscoveryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_discovery.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Discovery %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Discovery %q: %s", d.Id(), err)
	}
	return setDiscoveryFields(client, d, resp)
}

func setDiscoveryFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraDiscoveryUpdate updates Discovery resource
func resourceVolterraDiscoveryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_discovery.ReplaceSpecType{}
	updateReq := &ves_io_schema_discovery.ReplaceRequest{
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

	discoveryChoiceTypeFound := false

	if v, ok := d.GetOk("discovery_consul"); ok && !discoveryChoiceTypeFound {

		discoveryChoiceTypeFound = true
		discoveryChoiceInt := &ves_io_schema_discovery.ReplaceSpecType_DiscoveryConsul{}
		discoveryChoiceInt.DiscoveryConsul = &ves_io_schema_discovery.ConsulDiscoveryType{}
		updateSpec.DiscoveryChoice = discoveryChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["access_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				accessInfo := &ves_io_schema_discovery.ConsulAccessInfo{}
				discoveryChoiceInt.DiscoveryConsul.AccessInfo = accessInfo
				for _, set := range sl {

					accessInfoMapStrToI := set.(map[string]interface{})

					if v, ok := accessInfoMapStrToI["connection_info"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						connectionInfo := &ves_io_schema_discovery.RestConfigType{}
						accessInfo.ConnectionInfo = connectionInfo
						for _, set := range sl {

							connectionInfoMapStrToI := set.(map[string]interface{})

							if w, ok := connectionInfoMapStrToI["api_server"]; ok && !isIntfNil(w) {
								connectionInfo.ApiServer = w.(string)
							}

							if v, ok := connectionInfoMapStrToI["tls_info"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								tlsInfo := &ves_io_schema_discovery.TLSClientConfigType{}
								connectionInfo.TlsInfo = tlsInfo
								for _, set := range sl {

									tlsInfoMapStrToI := set.(map[string]interface{})

									if v, ok := tlsInfoMapStrToI["ca_certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										caCertificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CaCertificateUrl = caCertificateUrl
										for _, set := range sl {

											caCertificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := caCertificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												caCertificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := caCertificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["certificate"]; ok && !isIntfNil(w) {
										tlsInfo.Certificate = w.(string)
									}

									if v, ok := tlsInfoMapStrToI["certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										certificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CertificateUrl = certificateUrl
										for _, set := range sl {

											certificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := certificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												certificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := certificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := tlsInfoMapStrToI["key_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										keyUrl := &ves_io_schema.SecretType{}
										tlsInfo.KeyUrl = keyUrl
										for _, set := range sl {

											keyUrlMapStrToI := set.(map[string]interface{})

											if v, ok := keyUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												keyUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := keyUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["server_name"]; ok && !isIntfNil(w) {
										tlsInfo.ServerName = w.(string)
									}

									if w, ok := tlsInfoMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
										tlsInfo.TrustedCaUrl = w.(string)
									}

								}

							}

						}

					}

					if v, ok := accessInfoMapStrToI["http_basic_auth_info"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						httpBasicAuthInfo := &ves_io_schema_discovery.ConsulHttpBasicAuthInfoType{}
						accessInfo.HttpBasicAuthInfo = httpBasicAuthInfo
						for _, set := range sl {

							httpBasicAuthInfoMapStrToI := set.(map[string]interface{})

							if v, ok := httpBasicAuthInfoMapStrToI["passwd_url"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								passwdUrl := &ves_io_schema.SecretType{}
								httpBasicAuthInfo.PasswdUrl = passwdUrl
								for _, set := range sl {

									passwdUrlMapStrToI := set.(map[string]interface{})

									if v, ok := passwdUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										passwdUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := passwdUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwdUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwdUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwdUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										passwdUrl.SecretInfoOneof = secretInfoOneofInt

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

							if w, ok := httpBasicAuthInfoMapStrToI["user_name"]; ok && !isIntfNil(w) {
								httpBasicAuthInfo.UserName = w.(string)
							}

						}

					}

					if v, ok := accessInfoMapStrToI["scheme"]; ok && !isIntfNil(v) {

						accessInfo.Scheme = ves_io_schema_discovery.SchemeType(ves_io_schema_discovery.SchemeType_value[v.(string)])

					}

				}

			}

			if v, ok := cs["publish_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				publishInfo := &ves_io_schema_discovery.ConsulVipDiscoveryInfoType{}
				discoveryChoiceInt.DiscoveryConsul.PublishInfo = publishInfo
				for _, set := range sl {

					publishInfoMapStrToI := set.(map[string]interface{})

					publishChoiceTypeFound := false

					if v, ok := publishInfoMapStrToI["disable"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.ConsulVipDiscoveryInfoType_Disable{}
							publishChoiceInt.Disable = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

					if v, ok := publishInfoMapStrToI["publish"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.ConsulVipDiscoveryInfoType_Publish{}
							publishChoiceInt.Publish = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("discovery_k8s"); ok && !discoveryChoiceTypeFound {

		discoveryChoiceTypeFound = true
		discoveryChoiceInt := &ves_io_schema_discovery.ReplaceSpecType_DiscoveryK8S{}
		discoveryChoiceInt.DiscoveryK8S = &ves_io_schema_discovery.K8SDiscoveryType{}
		updateSpec.DiscoveryChoice = discoveryChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["access_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				accessInfo := &ves_io_schema_discovery.K8SAccessInfo{}
				discoveryChoiceInt.DiscoveryK8S.AccessInfo = accessInfo
				for _, set := range sl {

					accessInfoMapStrToI := set.(map[string]interface{})

					configTypeTypeFound := false

					if v, ok := accessInfoMapStrToI["connection_info"]; ok && !isIntfNil(v) && !configTypeTypeFound {

						configTypeTypeFound = true
						configTypeInt := &ves_io_schema_discovery.K8SAccessInfo_ConnectionInfo{}
						configTypeInt.ConnectionInfo = &ves_io_schema_discovery.RestConfigType{}
						accessInfo.ConfigType = configTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["api_server"]; ok && !isIntfNil(v) {

								configTypeInt.ConnectionInfo.ApiServer = v.(string)
							}

							if v, ok := cs["tls_info"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								tlsInfo := &ves_io_schema_discovery.TLSClientConfigType{}
								configTypeInt.ConnectionInfo.TlsInfo = tlsInfo
								for _, set := range sl {

									tlsInfoMapStrToI := set.(map[string]interface{})

									if v, ok := tlsInfoMapStrToI["ca_certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										caCertificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CaCertificateUrl = caCertificateUrl
										for _, set := range sl {

											caCertificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := caCertificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												caCertificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := caCertificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := caCertificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												caCertificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["certificate"]; ok && !isIntfNil(w) {
										tlsInfo.Certificate = w.(string)
									}

									if v, ok := tlsInfoMapStrToI["certificate_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										certificateUrl := &ves_io_schema.SecretType{}
										tlsInfo.CertificateUrl = certificateUrl
										for _, set := range sl {

											certificateUrlMapStrToI := set.(map[string]interface{})

											if v, ok := certificateUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												certificateUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := certificateUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := certificateUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												certificateUrl.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := tlsInfoMapStrToI["key_url"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										keyUrl := &ves_io_schema.SecretType{}
										tlsInfo.KeyUrl = keyUrl
										for _, set := range sl {

											keyUrlMapStrToI := set.(map[string]interface{})

											if v, ok := keyUrlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												keyUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := keyUrlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := keyUrlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												keyUrl.SecretInfoOneof = secretInfoOneofInt

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

									if w, ok := tlsInfoMapStrToI["server_name"]; ok && !isIntfNil(w) {
										tlsInfo.ServerName = w.(string)
									}

									if w, ok := tlsInfoMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
										tlsInfo.TrustedCaUrl = w.(string)
									}

								}

							}

						}

					}

					if v, ok := accessInfoMapStrToI["in_cluster"]; ok && !isIntfNil(v) && !configTypeTypeFound {

						configTypeTypeFound = true
						configTypeInt := &ves_io_schema_discovery.K8SAccessInfo_InCluster{}

						accessInfo.ConfigType = configTypeInt

						configTypeInt.InCluster =
							v.(bool)

					}

					if v, ok := accessInfoMapStrToI["kubeconfig_url"]; ok && !isIntfNil(v) && !configTypeTypeFound {

						configTypeTypeFound = true
						configTypeInt := &ves_io_schema_discovery.K8SAccessInfo_KubeconfigUrl{}
						configTypeInt.KubeconfigUrl = &ves_io_schema.SecretType{}
						accessInfo.ConfigType = configTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["secret_encoding_type"]; ok && !isIntfNil(v) {

								configTypeInt.KubeconfigUrl.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

							}

							secretInfoOneofTypeFound := false

							if v, ok := cs["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
								secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
								secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
								secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

								secretInfoOneofTypeFound = true
								secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
								secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
								configTypeInt.KubeconfigUrl.SecretInfoOneof = secretInfoOneofInt

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

					k8SPodNetworkChoiceTypeFound := false

					if v, ok := accessInfoMapStrToI["isolated"]; ok && !isIntfNil(v) && !k8SPodNetworkChoiceTypeFound {

						k8SPodNetworkChoiceTypeFound = true

						if v.(bool) {
							k8SPodNetworkChoiceInt := &ves_io_schema_discovery.K8SAccessInfo_Isolated{}
							k8SPodNetworkChoiceInt.Isolated = &ves_io_schema.Empty{}
							accessInfo.K8SPodNetworkChoice = k8SPodNetworkChoiceInt
						}

					}

					if v, ok := accessInfoMapStrToI["reachable"]; ok && !isIntfNil(v) && !k8SPodNetworkChoiceTypeFound {

						k8SPodNetworkChoiceTypeFound = true

						if v.(bool) {
							k8SPodNetworkChoiceInt := &ves_io_schema_discovery.K8SAccessInfo_Reachable{}
							k8SPodNetworkChoiceInt.Reachable = &ves_io_schema.Empty{}
							accessInfo.K8SPodNetworkChoice = k8SPodNetworkChoiceInt
						}

					}

				}

			}

			if v, ok := cs["publish_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				publishInfo := &ves_io_schema_discovery.K8SVipDiscoveryInfoType{}
				discoveryChoiceInt.DiscoveryK8S.PublishInfo = publishInfo
				for _, set := range sl {

					publishInfoMapStrToI := set.(map[string]interface{})

					publishChoiceTypeFound := false

					if v, ok := publishInfoMapStrToI["disable"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_Disable{}
							publishChoiceInt.Disable = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

					if v, ok := publishInfoMapStrToI["dns_delegation"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true
						publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_DnsDelegation{}
						publishChoiceInt.DnsDelegation = &ves_io_schema_discovery.K8SDelegationType{}
						publishInfo.PublishChoice = publishChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["dns_mode"]; ok && !isIntfNil(v) {

								publishChoiceInt.DnsDelegation.DnsMode = ves_io_schema_discovery.K8SDNSMode(ves_io_schema_discovery.K8SDNSMode_value[v.(string)])

							}

							if v, ok := cs["subdomain"]; ok && !isIntfNil(v) {

								publishChoiceInt.DnsDelegation.Subdomain = v.(string)
							}

						}

					}

					if v, ok := publishInfoMapStrToI["publish"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true
						publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_Publish{}
						publishChoiceInt.Publish = &ves_io_schema_discovery.K8SPublishType{}
						publishInfo.PublishChoice = publishChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								publishChoiceInt.Publish.Namespace = v.(string)
							}

						}

					}

					if v, ok := publishInfoMapStrToI["publish_fqdns"]; ok && !isIntfNil(v) && !publishChoiceTypeFound {

						publishChoiceTypeFound = true

						if v.(bool) {
							publishChoiceInt := &ves_io_schema_discovery.K8SVipDiscoveryInfoType_PublishFqdns{}
							publishChoiceInt.PublishFqdns = &ves_io_schema.Empty{}
							publishInfo.PublishChoice = publishChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.NetworkSiteRefSelector{}
		updateSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_Site{}
				refOrSelectorInt.Site = &ves_io_schema.SiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.Site.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.Site.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_network"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualNetwork{}
				refOrSelectorInt.VirtualNetwork = &ves_io_schema.NetworkRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualNetwork.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_network"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualSite{}
				refOrSelectorInt.VirtualSite = &ves_io_schema.VSiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.VirtualSite.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualSite.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra Discovery obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_discovery.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Discovery: %s", err)
	}

	return resourceVolterraDiscoveryRead(d, meta)
}

func resourceVolterraDiscoveryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_discovery.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Discovery %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Discovery before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Discovery obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_discovery.ObjectType, namespace, name)
}