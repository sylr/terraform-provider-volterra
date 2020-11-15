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
	ves_io_schema_fleet "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fleet"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraFleet is implementation of Volterra's Fleet resources
func resourceVolterraFleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraFleetCreate,
		Read:   resourceVolterraFleetRead,
		Update: resourceVolterraFleetUpdate,
		Delete: resourceVolterraFleetDelete,

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

			"bond_device_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bond_devices": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"devices": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"active_backup": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"lacp": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rate": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"link_polling_interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"link_up_delay": {
										Type:     schema.TypeInt,
										Optional: true,
									},

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

			"no_bond_devices": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"dc_cluster_group": {

				Type:     schema.TypeSet,
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

			"dc_cluster_group_inside": {

				Type:     schema.TypeSet,
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

			"no_dc_cluster_group": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_default_fleet_config_download": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"fleet_label": {
				Type:     schema.TypeString,
				Required: true,
			},

			"disable_gpu": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_gpu": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"inside_virtual_network": {

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

			"default_config": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"device_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"devices": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_device": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": {

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

												"use": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"owner": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"interface_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interfaces": {

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

			"network_connectors": {

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

			"network_firewall": {

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

			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"outside_virtual_network": {

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

			"default_storage_class": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"storage_class_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"storage_classes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_storage_parameters": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"default_storage_class": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"dell_emc_isilon_f800": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"base_path": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"az_service_ip_address": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"az_service_name": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"iscsi_access_zone": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"root_client_enable": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"hpe_nimbus_storage_af40": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"limit_iops": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"limit_mbps": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"perf_policy": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"netapp_trident": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selector": {
													Type:     schema.TypeMap,
													Optional: true,
												},
											},
										},
									},

									"pure_service_orchestrator": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"backend": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"storage_class_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"storage_device": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"no_storage_device": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"storage_device_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"storage_devices": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_advanced_parameters": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"dell_emc_isilon_f800": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"api_server_ip_address": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"api_server_name": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"api_server_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"base_path": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"secure_network": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"trusted_ca_url": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"iscsi_access_zone": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"password": {

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

												"username": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"volume_prefix": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"hpe_nimbus_storage_af40": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"api_server_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"limit_iops": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"limit_mbps": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"password": {

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

												"perf_policy": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"storage_server_ip_address": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"storage_server_name": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"username": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"netapp_trident": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"netapp_backend_ontap_nas": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"auto_export_cidrs": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"prefixes": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"auto_export_policy": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"backend_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"data_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"data_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"limit_aggregate_usage": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"limit_volume_size": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"management_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"management_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"nfs_mount_options": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"password": {

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

															"region": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"storage": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"labels": {
																			Type:     schema.TypeMap,
																			Optional: true,
																		},

																		"volume_defaults": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"encryption": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"export_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"security_style": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_dir": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"snapshot_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"space_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"split_on_clone": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"tiering_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"unix_permissions": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"zone": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"storage_driver_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"storage_prefix": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"svm": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"username": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"volume_defaults": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"encryption": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"export_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"security_style": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_dir": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"snapshot_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"space_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"split_on_clone": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"tiering_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"unix_permissions": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"netapp_backend_ontap_san": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"no_chap": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"use_chap": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"chap_initiator_secret": {

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

																		"chap_target_initiator_secret": {

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

																		"chap_target_username": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"chap_username": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"data_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"data_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"igroup_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"limit_aggregate_usage": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"limit_volume_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"management_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"management_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"password": {

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

															"region": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"storage": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"labels": {
																			Type:     schema.TypeMap,
																			Optional: true,
																		},

																		"volume_defaults": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"encryption": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"export_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"security_style": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_dir": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"snapshot_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"space_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"split_on_clone": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"tiering_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"unix_permissions": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"zone": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"storage_driver_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"storage_prefix": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"svm": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"username": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"volume_defaults": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"encryption": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"export_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"security_style": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_dir": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"snapshot_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"space_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"split_on_clone": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"tiering_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"unix_permissions": {
																			Type:     schema.TypeInt,
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

									"pure_service_orchestrator": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"arrays": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"flash_array": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"default_fs_opt": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"default_fs_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"default_mount_opts": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"disable_preempt_attachments": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"flash_arrays": {

																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"api_token": {

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

																					"labels": {
																						Type:     schema.TypeMap,
																						Optional: true,
																					},

																					"mgmt_dns_name": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"mgmt_ip": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"iscsi_login_timeout": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},

																		"san_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"flash_blade": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enable_snapshot_directory": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"export_rules": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"flash_blades": {

																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"api_token": {

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

																					"lables": {
																						Type:     schema.TypeMap,
																						Optional: true,
																					},

																					"mgmt_dns_name": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"mgmt_ip": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"nfs_endpoint_dns_name": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"nfs_endpoint_ip": {

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

												"cluster_id": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"enable_storage_topology": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"enable_strict_topology": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"storage_device": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"no_storage_interfaces": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"storage_interface_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interfaces": {

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

			"no_storage_static_routes": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"storage_static_routes": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"storage_routes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attrs": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"labels": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"nexthop": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": {

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

												"nexthop_address": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"addr": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"ipv6": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"addr": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"type": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"subnets": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ipv4": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"plen": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"prefix": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"ipv6": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"plen": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"prefix": {
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
				},
			},

			"volterra_software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraFleetCreate creates Fleet resource
func resourceVolterraFleetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_fleet.CreateSpecType{}
	createReq := &ves_io_schema_fleet.CreateRequest{
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

	bondChoiceTypeFound := false

	if v, ok := d.GetOk("bond_device_list"); ok && !bondChoiceTypeFound {

		bondChoiceTypeFound = true
		bondChoiceInt := &ves_io_schema_fleet.CreateSpecType_BondDeviceList{}
		bondChoiceInt.BondDeviceList = &ves_io_schema_fleet.FleetBondDevicesListType{}
		createSpec.BondChoice = bondChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["bond_devices"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				bondDevices := make([]*ves_io_schema_fleet.FleetBondDeviceType, len(sl))
				bondChoiceInt.BondDeviceList.BondDevices = bondDevices
				for i, set := range sl {
					bondDevices[i] = &ves_io_schema_fleet.FleetBondDeviceType{}

					bondDevicesMapStrToI := set.(map[string]interface{})

					if w, ok := bondDevicesMapStrToI["devices"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						bondDevices[i].Devices = ls
					}

					lacpChoiceTypeFound := false

					if v, ok := bondDevicesMapStrToI["active_backup"]; ok && !isIntfNil(v) && !lacpChoiceTypeFound {

						lacpChoiceTypeFound = true

						if v.(bool) {
							lacpChoiceInt := &ves_io_schema_fleet.FleetBondDeviceType_ActiveBackup{}
							lacpChoiceInt.ActiveBackup = &ves_io_schema.Empty{}
							bondDevices[i].LacpChoice = lacpChoiceInt
						}

					}

					if v, ok := bondDevicesMapStrToI["lacp"]; ok && !isIntfNil(v) && !lacpChoiceTypeFound {

						lacpChoiceTypeFound = true
						lacpChoiceInt := &ves_io_schema_fleet.FleetBondDeviceType_Lacp{}
						lacpChoiceInt.Lacp = &ves_io_schema_fleet.BondLacpType{}
						bondDevices[i].LacpChoice = lacpChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["rate"]; ok && !isIntfNil(v) {

								lacpChoiceInt.Lacp.Rate = uint32(v.(int))
							}

						}

					}

					if w, ok := bondDevicesMapStrToI["link_polling_interval"]; ok && !isIntfNil(w) {
						bondDevices[i].LinkPollingInterval = w.(uint32)
					}

					if w, ok := bondDevicesMapStrToI["link_up_delay"]; ok && !isIntfNil(w) {
						bondDevices[i].LinkUpDelay = w.(uint32)
					}

					if w, ok := bondDevicesMapStrToI["name"]; ok && !isIntfNil(w) {
						bondDevices[i].Name = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("no_bond_devices"); ok && !bondChoiceTypeFound {

		bondChoiceTypeFound = true

		if v.(bool) {
			bondChoiceInt := &ves_io_schema_fleet.CreateSpecType_NoBondDevices{}
			bondChoiceInt.NoBondDevices = &ves_io_schema.Empty{}
			createSpec.BondChoice = bondChoiceInt
		}

	}

	dcClusterGroupChoiceTypeFound := false

	if v, ok := d.GetOk("dc_cluster_group"); ok && !dcClusterGroupChoiceTypeFound {

		dcClusterGroupChoiceTypeFound = true
		dcClusterGroupChoiceInt := &ves_io_schema_fleet.CreateSpecType_DcClusterGroup{}
		dcClusterGroupChoiceInt.DcClusterGroup = &ves_io_schema_views.ObjectRefType{}
		createSpec.DcClusterGroupChoice = dcClusterGroupChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroup.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroup.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroup.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("dc_cluster_group_inside"); ok && !dcClusterGroupChoiceTypeFound {

		dcClusterGroupChoiceTypeFound = true
		dcClusterGroupChoiceInt := &ves_io_schema_fleet.CreateSpecType_DcClusterGroupInside{}
		dcClusterGroupChoiceInt.DcClusterGroupInside = &ves_io_schema_views.ObjectRefType{}
		createSpec.DcClusterGroupChoice = dcClusterGroupChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroupInside.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroupInside.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroupInside.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("no_dc_cluster_group"); ok && !dcClusterGroupChoiceTypeFound {

		dcClusterGroupChoiceTypeFound = true

		if v.(bool) {
			dcClusterGroupChoiceInt := &ves_io_schema_fleet.CreateSpecType_NoDcClusterGroup{}
			dcClusterGroupChoiceInt.NoDcClusterGroup = &ves_io_schema.Empty{}
			createSpec.DcClusterGroupChoice = dcClusterGroupChoiceInt
		}

	}

	if v, ok := d.GetOk("enable_default_fleet_config_download"); ok && !isIntfNil(v) {

		createSpec.EnableDefaultFleetConfigDownload =
			v.(bool)
	}

	if v, ok := d.GetOk("fleet_label"); ok && !isIntfNil(v) {

		createSpec.FleetLabel =
			v.(string)
	}

	gpuChoiceTypeFound := false

	if v, ok := d.GetOk("disable_gpu"); ok && !gpuChoiceTypeFound {

		gpuChoiceTypeFound = true

		if v.(bool) {
			gpuChoiceInt := &ves_io_schema_fleet.CreateSpecType_DisableGpu{}
			gpuChoiceInt.DisableGpu = &ves_io_schema.Empty{}
			createSpec.GpuChoice = gpuChoiceInt
		}

	}

	if v, ok := d.GetOk("enable_gpu"); ok && !gpuChoiceTypeFound {

		gpuChoiceTypeFound = true

		if v.(bool) {
			gpuChoiceInt := &ves_io_schema_fleet.CreateSpecType_EnableGpu{}
			gpuChoiceInt.EnableGpu = &ves_io_schema.Empty{}
			createSpec.GpuChoice = gpuChoiceInt
		}

	}

	if v, ok := d.GetOk("inside_virtual_network"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		insideVirtualNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.InsideVirtualNetwork = insideVirtualNetworkInt
		for i, ps := range sl {

			ivnMapToStrVal := ps.(map[string]interface{})
			insideVirtualNetworkInt[i] = &ves_io_schema.ObjectRefType{}

			insideVirtualNetworkInt[i].Kind = "virtual_network"

			if v, ok := ivnMapToStrVal["name"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Name = v.(string)
			}

			if v, ok := ivnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Namespace = v.(string)
			}

			if v, ok := ivnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Tenant = v.(string)
			}

			if v, ok := ivnMapToStrVal["uid"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Uid = v.(string)
			}

		}

	}

	interfaceChoiceTypeFound := false

	if v, ok := d.GetOk("default_config"); ok && !interfaceChoiceTypeFound {

		interfaceChoiceTypeFound = true

		if v.(bool) {
			interfaceChoiceInt := &ves_io_schema_fleet.CreateSpecType_DefaultConfig{}
			interfaceChoiceInt.DefaultConfig = &ves_io_schema.Empty{}
			createSpec.InterfaceChoice = interfaceChoiceInt
		}

	}

	if v, ok := d.GetOk("device_list"); ok && !interfaceChoiceTypeFound {

		interfaceChoiceTypeFound = true
		interfaceChoiceInt := &ves_io_schema_fleet.CreateSpecType_DeviceList{}
		interfaceChoiceInt.DeviceList = &ves_io_schema_fleet.FleetDeviceListType{}
		createSpec.InterfaceChoice = interfaceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["devices"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				devices := make([]*ves_io_schema_fleet.DeviceInstanceType, len(sl))
				interfaceChoiceInt.DeviceList.Devices = devices
				for i, set := range sl {
					devices[i] = &ves_io_schema_fleet.DeviceInstanceType{}

					devicesMapStrToI := set.(map[string]interface{})

					deviceInstanceTypeFound := false

					if v, ok := devicesMapStrToI["network_device"]; ok && !isIntfNil(v) && !deviceInstanceTypeFound {

						deviceInstanceTypeFound = true
						deviceInstanceInt := &ves_io_schema_fleet.DeviceInstanceType_NetworkDevice{}
						deviceInstanceInt.NetworkDevice = &ves_io_schema_fleet.NetworkingDeviceInstanceType{}
						devices[i].DeviceInstance = deviceInstanceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["interface"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								deviceInstanceInt.NetworkDevice.Interface = intfInt
								for i, ps := range sl {

									iMapToStrVal := ps.(map[string]interface{})
									intfInt[i] = &ves_io_schema.ObjectRefType{}

									intfInt[i].Kind = "network_interface"

									if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
										intfInt[i].Name = v.(string)
									}

									if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										intfInt[i].Namespace = v.(string)
									}

									if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										intfInt[i].Tenant = v.(string)
									}

									if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
										intfInt[i].Uid = v.(string)
									}

								}

							}

							if v, ok := cs["use"]; ok && !isIntfNil(v) {

								deviceInstanceInt.NetworkDevice.Use = ves_io_schema_fleet.NetworkingDeviceInstanceUseType(ves_io_schema_fleet.NetworkingDeviceInstanceUseType_value[v.(string)])

							}

						}

					}

					if w, ok := devicesMapStrToI["name"]; ok && !isIntfNil(w) {
						devices[i].Name = w.(string)
					}

					if v, ok := devicesMapStrToI["owner"]; ok && !isIntfNil(v) {

						devices[i].Owner = ves_io_schema_fleet.DeviceOwnerType(ves_io_schema_fleet.DeviceOwnerType_value[v.(string)])

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("interface_list"); ok && !interfaceChoiceTypeFound {

		interfaceChoiceTypeFound = true
		interfaceChoiceInt := &ves_io_schema_fleet.CreateSpecType_InterfaceList{}
		interfaceChoiceInt.InterfaceList = &ves_io_schema_fleet.FleetInterfaceListType{}
		createSpec.InterfaceChoice = interfaceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["interfaces"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				interfacesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				interfaceChoiceInt.InterfaceList.Interfaces = interfacesInt
				for i, ps := range sl {

					iMapToStrVal := ps.(map[string]interface{})
					interfacesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
						interfacesInt[i].Name = v.(string)
					}

					if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						interfacesInt[i].Namespace = v.(string)
					}

					if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						interfacesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("network_connectors"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		networkConnectorsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.NetworkConnectors = networkConnectorsInt
		for i, ps := range sl {

			ncMapToStrVal := ps.(map[string]interface{})
			networkConnectorsInt[i] = &ves_io_schema.ObjectRefType{}

			networkConnectorsInt[i].Kind = "network_connector"

			if v, ok := ncMapToStrVal["name"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Name = v.(string)
			}

			if v, ok := ncMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Namespace = v.(string)
			}

			if v, ok := ncMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Tenant = v.(string)
			}

			if v, ok := ncMapToStrVal["uid"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("network_firewall"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		networkFirewallInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.NetworkFirewall = networkFirewallInt
		for i, ps := range sl {

			nfMapToStrVal := ps.(map[string]interface{})
			networkFirewallInt[i] = &ves_io_schema.ObjectRefType{}

			networkFirewallInt[i].Kind = "network_firewall"

			if v, ok := nfMapToStrVal["name"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Name = v.(string)
			}

			if v, ok := nfMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Namespace = v.(string)
			}

			if v, ok := nfMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Tenant = v.(string)
			}

			if v, ok := nfMapToStrVal["uid"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		createSpec.OperatingSystemVersion =
			v.(string)
	}

	if v, ok := d.GetOk("outside_virtual_network"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		outsideVirtualNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.OutsideVirtualNetwork = outsideVirtualNetworkInt
		for i, ps := range sl {

			ovnMapToStrVal := ps.(map[string]interface{})
			outsideVirtualNetworkInt[i] = &ves_io_schema.ObjectRefType{}

			outsideVirtualNetworkInt[i].Kind = "virtual_network"

			if v, ok := ovnMapToStrVal["name"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Name = v.(string)
			}

			if v, ok := ovnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Namespace = v.(string)
			}

			if v, ok := ovnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Tenant = v.(string)
			}

			if v, ok := ovnMapToStrVal["uid"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Uid = v.(string)
			}

		}

	}

	storageClassChoiceTypeFound := false

	if v, ok := d.GetOk("default_storage_class"); ok && !storageClassChoiceTypeFound {

		storageClassChoiceTypeFound = true

		if v.(bool) {
			storageClassChoiceInt := &ves_io_schema_fleet.CreateSpecType_DefaultStorageClass{}
			storageClassChoiceInt.DefaultStorageClass = &ves_io_schema.Empty{}
			createSpec.StorageClassChoice = storageClassChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_class_list"); ok && !storageClassChoiceTypeFound {

		storageClassChoiceTypeFound = true
		storageClassChoiceInt := &ves_io_schema_fleet.CreateSpecType_StorageClassList{}
		storageClassChoiceInt.StorageClassList = &ves_io_schema_fleet.FleetStorageClassListType{}
		createSpec.StorageClassChoice = storageClassChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["storage_classes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				storageClasses := make([]*ves_io_schema_fleet.FleetStorageClassType, len(sl))
				storageClassChoiceInt.StorageClassList.StorageClasses = storageClasses
				for i, set := range sl {
					storageClasses[i] = &ves_io_schema_fleet.FleetStorageClassType{}

					storageClassesMapStrToI := set.(map[string]interface{})

					if w, ok := storageClassesMapStrToI["advanced_storage_parameters"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						storageClasses[i].AdvancedStorageParameters = ms
					}

					if w, ok := storageClassesMapStrToI["default_storage_class"]; ok && !isIntfNil(w) {
						storageClasses[i].DefaultStorageClass = w.(bool)
					}

					if w, ok := storageClassesMapStrToI["description"]; ok && !isIntfNil(w) {
						storageClasses[i].Description = w.(string)
					}

					deviceChoiceTypeFound := false

					if v, ok := storageClassesMapStrToI["dell_emc_isilon_f800"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_DellEmcIsilonF800{}
						deviceChoiceInt.DellEmcIsilonF800 = &ves_io_schema_fleet.StorageClassDellIsilonF800Type{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["base_path"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.BasePath = v.(string)
							}

							httpsChoiceTypeFound := false

							if v, ok := cs["az_service_ip_address"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true
								httpsChoiceInt := &ves_io_schema_fleet.StorageClassDellIsilonF800Type_AzServiceIpAddress{}

								deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt

								httpsChoiceInt.AzServiceIpAddress = v.(string)

							}

							if v, ok := cs["az_service_name"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true
								httpsChoiceInt := &ves_io_schema_fleet.StorageClassDellIsilonF800Type_AzServiceName{}

								deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt

								httpsChoiceInt.AzServiceName = v.(string)

							}

							if v, ok := cs["iscsi_access_zone"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.IscsiAccessZone = v.(string)
							}

							if v, ok := cs["root_client_enable"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.RootClientEnable = v.(bool)
							}

						}

					}

					if v, ok := storageClassesMapStrToI["hpe_nimbus_storage_af40"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_HpeNimbusStorageAf40{}
						deviceChoiceInt.HpeNimbusStorageAf40 = &ves_io_schema_fleet.StorageClassHPENimbusStorageAf40Type{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["limit_iops"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitIops = uint32(v.(int))
							}

							if v, ok := cs["limit_mbps"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitMbps = uint32(v.(int))
							}

							if v, ok := cs["perf_policy"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.PerfPolicy = v.(string)
							}

						}

					}

					if v, ok := storageClassesMapStrToI["netapp_trident"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_NetappTrident{}
						deviceChoiceInt.NetappTrident = &ves_io_schema_fleet.StorageClassNetappTridentType{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["selector"]; ok && !isIntfNil(v) {

								ms := map[string]string{}
								for k, v := range v.(map[string]interface{}) {
									ms[k] = v.(string)
								}
								deviceChoiceInt.NetappTrident.Selector = ms
							}

						}

					}

					if v, ok := storageClassesMapStrToI["pure_service_orchestrator"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_PureServiceOrchestrator{}
						deviceChoiceInt.PureServiceOrchestrator = &ves_io_schema_fleet.StorageClassPureServiceOrchestratorType{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["backend"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.Backend = v.(string)
							}

						}

					}

					if w, ok := storageClassesMapStrToI["storage_class_name"]; ok && !isIntfNil(w) {
						storageClasses[i].StorageClassName = w.(string)
					}

					if w, ok := storageClassesMapStrToI["storage_device"]; ok && !isIntfNil(w) {
						storageClasses[i].StorageDevice = w.(string)
					}

				}

			}

		}

	}

	storageDeviceChoiceTypeFound := false

	if v, ok := d.GetOk("no_storage_device"); ok && !storageDeviceChoiceTypeFound {

		storageDeviceChoiceTypeFound = true

		if v.(bool) {
			storageDeviceChoiceInt := &ves_io_schema_fleet.CreateSpecType_NoStorageDevice{}
			storageDeviceChoiceInt.NoStorageDevice = &ves_io_schema.Empty{}
			createSpec.StorageDeviceChoice = storageDeviceChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_device_list"); ok && !storageDeviceChoiceTypeFound {

		storageDeviceChoiceTypeFound = true
		storageDeviceChoiceInt := &ves_io_schema_fleet.CreateSpecType_StorageDeviceList{}
		storageDeviceChoiceInt.StorageDeviceList = &ves_io_schema_fleet.FleetStorageDeviceListType{}
		createSpec.StorageDeviceChoice = storageDeviceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["storage_devices"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				storageDevices := make([]*ves_io_schema_fleet.FleetStorageDeviceType, len(sl))
				storageDeviceChoiceInt.StorageDeviceList.StorageDevices = storageDevices
				for i, set := range sl {
					storageDevices[i] = &ves_io_schema_fleet.FleetStorageDeviceType{}

					storageDevicesMapStrToI := set.(map[string]interface{})

					if w, ok := storageDevicesMapStrToI["advanced_advanced_parameters"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						storageDevices[i].AdvancedAdvancedParameters = ms
					}

					deviceChoiceTypeFound := false

					if v, ok := storageDevicesMapStrToI["dell_emc_isilon_f800"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_DellEmcIsilonF800{}
						deviceChoiceInt.DellEmcIsilonF800 = &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							addressChoiceTypeFound := false

							if v, ok := cs["api_server_ip_address"]; ok && !isIntfNil(v) && !addressChoiceTypeFound {

								addressChoiceTypeFound = true
								addressChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_ApiServerIpAddress{}

								deviceChoiceInt.DellEmcIsilonF800.AddressChoice = addressChoiceInt

								addressChoiceInt.ApiServerIpAddress = v.(string)

							}

							if v, ok := cs["api_server_name"]; ok && !isIntfNil(v) && !addressChoiceTypeFound {

								addressChoiceTypeFound = true
								addressChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_ApiServerName{}

								deviceChoiceInt.DellEmcIsilonF800.AddressChoice = addressChoiceInt

								addressChoiceInt.ApiServerName = v.(string)

							}

							if v, ok := cs["api_server_port"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.ApiServerPort = uint32(v.(int))
							}

							if v, ok := cs["base_path"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.BasePath = v.(string)
							}

							httpsChoiceTypeFound := false

							if v, ok := cs["secure_network"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true

								if v.(bool) {
									httpsChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_SecureNetwork{}
									httpsChoiceInt.SecureNetwork = &ves_io_schema.Empty{}
									deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt
								}

							}

							if v, ok := cs["trusted_ca_url"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true
								httpsChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_TrustedCaUrl{}

								deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt

								httpsChoiceInt.TrustedCaUrl = v.(string)

							}

							if v, ok := cs["iscsi_access_zone"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.IscsiAccessZone = v.(string)
							}

							if v, ok := cs["password"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								password := &ves_io_schema.SecretType{}
								deviceChoiceInt.DellEmcIsilonF800.Password = password
								for _, set := range sl {

									passwordMapStrToI := set.(map[string]interface{})

									if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["username"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.Username = v.(string)
							}

							if v, ok := cs["volume_prefix"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.VolumePrefix = v.(string)
							}

						}

					}

					if v, ok := storageDevicesMapStrToI["hpe_nimbus_storage_af40"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_HpeNimbusStorageAf40{}
						deviceChoiceInt.HpeNimbusStorageAf40 = &ves_io_schema_fleet.StorageDeviceHPENimbusStorageAf40Type{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["api_server_port"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.ApiServerPort = uint32(v.(int))
							}

							if v, ok := cs["limit_iops"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitIops = uint32(v.(int))
							}

							if v, ok := cs["limit_mbps"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitMbps = uint32(v.(int))
							}

							if v, ok := cs["password"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								password := &ves_io_schema.SecretType{}
								deviceChoiceInt.HpeNimbusStorageAf40.Password = password
								for _, set := range sl {

									passwordMapStrToI := set.(map[string]interface{})

									if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["perf_policy"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.PerfPolicy = v.(string)
							}

							if v, ok := cs["storage_server_ip_address"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.StorageServerIpAddress = v.(string)
							}

							if v, ok := cs["storage_server_name"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.StorageServerName = v.(string)
							}

							if v, ok := cs["username"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.Username = v.(string)
							}

						}

					}

					if v, ok := storageDevicesMapStrToI["netapp_trident"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_NetappTrident{}
						deviceChoiceInt.NetappTrident = &ves_io_schema_fleet.StorageDeviceNetappTridentType{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							backendChoiceTypeFound := false

							if v, ok := cs["netapp_backend_ontap_nas"]; ok && !isIntfNil(v) && !backendChoiceTypeFound {

								backendChoiceTypeFound = true
								backendChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappTridentType_NetappBackendOntapNas{}
								backendChoiceInt.NetappBackendOntapNas = &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType{}
								deviceChoiceInt.NetappTrident.BackendChoice = backendChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["auto_export_cidrs"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										autoExportCidrs := &ves_io_schema_views.PrefixStringListType{}
										backendChoiceInt.NetappBackendOntapNas.AutoExportCidrs = autoExportCidrs
										for _, set := range sl {

											autoExportCidrsMapStrToI := set.(map[string]interface{})

											if w, ok := autoExportCidrsMapStrToI["prefixes"]; ok && !isIntfNil(w) {
												ls := make([]string, len(w.([]interface{})))
												for i, v := range w.([]interface{}) {
													ls[i] = v.(string)
												}
												autoExportCidrs.Prefixes = ls
											}

										}

									}

									if v, ok := cs["auto_export_policy"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.AutoExportPolicy = v.(bool)
									}

									if v, ok := cs["backend_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.BackendName = v.(string)
									}

									dataLifTypeFound := false

									if v, ok := cs["data_lif_dns_name"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_DataLifDnsName{}

										backendChoiceInt.NetappBackendOntapNas.DataLif = dataLifInt

										dataLifInt.DataLifDnsName = v.(string)

									}

									if v, ok := cs["data_lif_ip"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_DataLifIp{}

										backendChoiceInt.NetappBackendOntapNas.DataLif = dataLifInt

										dataLifInt.DataLifIp = v.(string)

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										backendChoiceInt.NetappBackendOntapNas.Labels = ms
									}

									if v, ok := cs["limit_aggregate_usage"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.LimitAggregateUsage = v.(string)
									}

									if v, ok := cs["limit_volume_size"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.LimitVolumeSize = v.(string)
									}

									managementLifTypeFound := false

									if v, ok := cs["management_lif_dns_name"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_ManagementLifDnsName{}

										backendChoiceInt.NetappBackendOntapNas.ManagementLif = managementLifInt

										managementLifInt.ManagementLifDnsName = v.(string)

									}

									if v, ok := cs["management_lif_ip"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_ManagementLifIp{}

										backendChoiceInt.NetappBackendOntapNas.ManagementLif = managementLifInt

										managementLifInt.ManagementLifIp = v.(string)

									}

									if v, ok := cs["nfs_mount_options"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.NfsMountOptions = v.(string)
									}

									if v, ok := cs["password"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										password := &ves_io_schema.SecretType{}
										backendChoiceInt.NetappBackendOntapNas.Password = password
										for _, set := range sl {

											passwordMapStrToI := set.(map[string]interface{})

											if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := cs["region"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.Region = v.(string)
									}

									if v, ok := cs["storage"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										storage := make([]*ves_io_schema_fleet.OntapVirtualStoragePoolType, len(sl))
										backendChoiceInt.NetappBackendOntapNas.Storage = storage
										for i, set := range sl {
											storage[i] = &ves_io_schema_fleet.OntapVirtualStoragePoolType{}

											storageMapStrToI := set.(map[string]interface{})

											if w, ok := storageMapStrToI["labels"]; ok && !isIntfNil(w) {
												ms := map[string]string{}
												for k, v := range w.(map[string]interface{}) {
													ms[k] = v.(string)
												}
												storage[i].Labels = ms
											}

											if v, ok := storageMapStrToI["volume_defaults"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
												storage[i].VolumeDefaults = volumeDefaults
												for _, set := range sl {

													volumeDefaultsMapStrToI := set.(map[string]interface{})

													if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
														volumeDefaults.Encryption = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.ExportPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
														volumeDefaults.SecurityStyle = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotDir = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SpaceReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
														volumeDefaults.SplitOnClone = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.TieringPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
														volumeDefaults.UnixPermissions = w.(int32)
													}

												}

											}

											if w, ok := storageMapStrToI["zone"]; ok && !isIntfNil(w) {
												storage[i].Zone = w.(string)
											}

										}

									}

									if v, ok := cs["storage_driver_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.StorageDriverName = v.(string)
									}

									if v, ok := cs["storage_prefix"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.StoragePrefix = v.(string)
									}

									if v, ok := cs["svm"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.Svm = v.(string)
									}

									if v, ok := cs["username"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.Username = v.(string)
									}

									if v, ok := cs["volume_defaults"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
										backendChoiceInt.NetappBackendOntapNas.VolumeDefaults = volumeDefaults
										for _, set := range sl {

											volumeDefaultsMapStrToI := set.(map[string]interface{})

											if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
												volumeDefaults.Encryption = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.ExportPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
												volumeDefaults.SecurityStyle = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotDir = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SpaceReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
												volumeDefaults.SplitOnClone = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.TieringPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
												volumeDefaults.UnixPermissions = w.(int32)
											}

										}

									}

								}

							}

							if v, ok := cs["netapp_backend_ontap_san"]; ok && !isIntfNil(v) && !backendChoiceTypeFound {

								backendChoiceTypeFound = true
								backendChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappTridentType_NetappBackendOntapSan{}
								backendChoiceInt.NetappBackendOntapSan = &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType{}
								deviceChoiceInt.NetappTrident.BackendChoice = backendChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									chapChoiceTypeFound := false

									if v, ok := cs["no_chap"]; ok && !isIntfNil(v) && !chapChoiceTypeFound {

										chapChoiceTypeFound = true

										if v.(bool) {
											chapChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_NoChap{}
											chapChoiceInt.NoChap = &ves_io_schema.Empty{}
											backendChoiceInt.NetappBackendOntapSan.ChapChoice = chapChoiceInt
										}

									}

									if v, ok := cs["use_chap"]; ok && !isIntfNil(v) && !chapChoiceTypeFound {

										chapChoiceTypeFound = true
										chapChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_UseChap{}
										chapChoiceInt.UseChap = &ves_io_schema_fleet.DeviceNetappBackendOntapSanChapType{}
										backendChoiceInt.NetappBackendOntapSan.ChapChoice = chapChoiceInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["chap_initiator_secret"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												chapInitiatorSecret := &ves_io_schema.SecretType{}
												chapChoiceInt.UseChap.ChapInitiatorSecret = chapInitiatorSecret
												for _, set := range sl {

													chapInitiatorSecretMapStrToI := set.(map[string]interface{})

													if v, ok := chapInitiatorSecretMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

														chapInitiatorSecret.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

													}

													secretInfoOneofTypeFound := false

													if v, ok := chapInitiatorSecretMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
														secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapInitiatorSecretMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
														secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapInitiatorSecretMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
														secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapInitiatorSecretMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
														secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := cs["chap_target_initiator_secret"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												chapTargetInitiatorSecret := &ves_io_schema.SecretType{}
												chapChoiceInt.UseChap.ChapTargetInitiatorSecret = chapTargetInitiatorSecret
												for _, set := range sl {

													chapTargetInitiatorSecretMapStrToI := set.(map[string]interface{})

													if v, ok := chapTargetInitiatorSecretMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

														chapTargetInitiatorSecret.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

													}

													secretInfoOneofTypeFound := false

													if v, ok := chapTargetInitiatorSecretMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
														secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapTargetInitiatorSecretMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
														secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapTargetInitiatorSecretMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
														secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapTargetInitiatorSecretMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
														secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := cs["chap_target_username"]; ok && !isIntfNil(v) {

												chapChoiceInt.UseChap.ChapTargetUsername = v.(string)
											}

											if v, ok := cs["chap_username"]; ok && !isIntfNil(v) {

												chapChoiceInt.UseChap.ChapUsername = v.(string)
											}

										}

									}

									dataLifTypeFound := false

									if v, ok := cs["data_lif_dns_name"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_DataLifDnsName{}

										backendChoiceInt.NetappBackendOntapSan.DataLif = dataLifInt

										dataLifInt.DataLifDnsName = v.(string)

									}

									if v, ok := cs["data_lif_ip"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_DataLifIp{}

										backendChoiceInt.NetappBackendOntapSan.DataLif = dataLifInt

										dataLifInt.DataLifIp = v.(string)

									}

									if v, ok := cs["igroup_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.IgroupName = v.(string)
									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										backendChoiceInt.NetappBackendOntapSan.Labels = ms
									}

									if v, ok := cs["limit_aggregate_usage"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.LimitAggregateUsage = uint32(v.(int))
									}

									if v, ok := cs["limit_volume_size"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.LimitVolumeSize = v.(int32)
									}

									managementLifTypeFound := false

									if v, ok := cs["management_lif_dns_name"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_ManagementLifDnsName{}

										backendChoiceInt.NetappBackendOntapSan.ManagementLif = managementLifInt

										managementLifInt.ManagementLifDnsName = v.(string)

									}

									if v, ok := cs["management_lif_ip"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_ManagementLifIp{}

										backendChoiceInt.NetappBackendOntapSan.ManagementLif = managementLifInt

										managementLifInt.ManagementLifIp = v.(string)

									}

									if v, ok := cs["password"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										password := &ves_io_schema.SecretType{}
										backendChoiceInt.NetappBackendOntapSan.Password = password
										for _, set := range sl {

											passwordMapStrToI := set.(map[string]interface{})

											if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := cs["region"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.Region = v.(string)
									}

									if v, ok := cs["storage"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										storage := make([]*ves_io_schema_fleet.OntapVirtualStoragePoolType, len(sl))
										backendChoiceInt.NetappBackendOntapSan.Storage = storage
										for i, set := range sl {
											storage[i] = &ves_io_schema_fleet.OntapVirtualStoragePoolType{}

											storageMapStrToI := set.(map[string]interface{})

											if w, ok := storageMapStrToI["labels"]; ok && !isIntfNil(w) {
												ms := map[string]string{}
												for k, v := range w.(map[string]interface{}) {
													ms[k] = v.(string)
												}
												storage[i].Labels = ms
											}

											if v, ok := storageMapStrToI["volume_defaults"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
												storage[i].VolumeDefaults = volumeDefaults
												for _, set := range sl {

													volumeDefaultsMapStrToI := set.(map[string]interface{})

													if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
														volumeDefaults.Encryption = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.ExportPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
														volumeDefaults.SecurityStyle = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotDir = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SpaceReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
														volumeDefaults.SplitOnClone = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.TieringPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
														volumeDefaults.UnixPermissions = w.(int32)
													}

												}

											}

											if w, ok := storageMapStrToI["zone"]; ok && !isIntfNil(w) {
												storage[i].Zone = w.(string)
											}

										}

									}

									if v, ok := cs["storage_driver_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.StorageDriverName = v.(string)
									}

									if v, ok := cs["storage_prefix"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.StoragePrefix = v.(string)
									}

									if v, ok := cs["svm"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.Svm = v.(string)
									}

									if v, ok := cs["username"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.Username = v.(string)
									}

									if v, ok := cs["volume_defaults"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
										backendChoiceInt.NetappBackendOntapSan.VolumeDefaults = volumeDefaults
										for _, set := range sl {

											volumeDefaultsMapStrToI := set.(map[string]interface{})

											if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
												volumeDefaults.Encryption = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.ExportPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
												volumeDefaults.SecurityStyle = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotDir = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SpaceReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
												volumeDefaults.SplitOnClone = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.TieringPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
												volumeDefaults.UnixPermissions = w.(int32)
											}

										}

									}

								}

							}

						}

					}

					if v, ok := storageDevicesMapStrToI["pure_service_orchestrator"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_PureServiceOrchestrator{}
						deviceChoiceInt.PureServiceOrchestrator = &ves_io_schema_fleet.StorageDevicePureStorageServiceOrchestratorType{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["arrays"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								arrays := &ves_io_schema_fleet.PsoArrayConfiguration{}
								deviceChoiceInt.PureServiceOrchestrator.Arrays = arrays
								for _, set := range sl {

									arraysMapStrToI := set.(map[string]interface{})

									if v, ok := arraysMapStrToI["flash_array"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										flashArray := &ves_io_schema_fleet.FlashArrayType{}
										arrays.FlashArray = flashArray
										for _, set := range sl {

											flashArrayMapStrToI := set.(map[string]interface{})

											if w, ok := flashArrayMapStrToI["default_fs_opt"]; ok && !isIntfNil(w) {
												flashArray.DefaultFsOpt = w.(string)
											}

											if w, ok := flashArrayMapStrToI["default_fs_type"]; ok && !isIntfNil(w) {
												flashArray.DefaultFsType = w.(string)
											}

											if w, ok := flashArrayMapStrToI["default_mount_opts"]; ok && !isIntfNil(w) {
												ls := make([]string, len(w.([]interface{})))
												for i, v := range w.([]interface{}) {
													ls[i] = v.(string)
												}
												flashArray.DefaultMountOpts = ls
											}

											if w, ok := flashArrayMapStrToI["disable_preempt_attachments"]; ok && !isIntfNil(w) {
												flashArray.DisablePreemptAttachments = w.(bool)
											}

											if v, ok := flashArrayMapStrToI["flash_arrays"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												flashArrays := make([]*ves_io_schema_fleet.FlashArrayEndpoint, len(sl))
												flashArray.FlashArrays = flashArrays
												for i, set := range sl {
													flashArrays[i] = &ves_io_schema_fleet.FlashArrayEndpoint{}

													flashArraysMapStrToI := set.(map[string]interface{})

													if v, ok := flashArraysMapStrToI["api_token"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														apiToken := &ves_io_schema.SecretType{}
														flashArrays[i].ApiToken = apiToken
														for _, set := range sl {

															apiTokenMapStrToI := set.(map[string]interface{})

															if v, ok := apiTokenMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

																apiToken.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

															}

															secretInfoOneofTypeFound := false

															if v, ok := apiTokenMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
																secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
																secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
																secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
																secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

													if w, ok := flashArraysMapStrToI["labels"]; ok && !isIntfNil(w) {
														ms := map[string]string{}
														for k, v := range w.(map[string]interface{}) {
															ms[k] = v.(string)
														}
														flashArrays[i].Labels = ms
													}

													mgmtEndpointTypeFound := false

													if v, ok := flashArraysMapStrToI["mgmt_dns_name"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashArrayEndpoint_MgmtDnsName{}

														flashArrays[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtDnsName = v.(string)

													}

													if v, ok := flashArraysMapStrToI["mgmt_ip"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashArrayEndpoint_MgmtIp{}

														flashArrays[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtIp = v.(string)

													}

												}

											}

											if w, ok := flashArrayMapStrToI["iscsi_login_timeout"]; ok && !isIntfNil(w) {
												flashArray.IscsiLoginTimeout = w.(int32)
											}

											if w, ok := flashArrayMapStrToI["san_type"]; ok && !isIntfNil(w) {
												flashArray.SanType = w.(string)
											}

										}

									}

									if v, ok := arraysMapStrToI["flash_blade"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										flashBlade := &ves_io_schema_fleet.FlashBladeType{}
										arrays.FlashBlade = flashBlade
										for _, set := range sl {

											flashBladeMapStrToI := set.(map[string]interface{})

											if w, ok := flashBladeMapStrToI["enable_snapshot_directory"]; ok && !isIntfNil(w) {
												flashBlade.EnableSnapshotDirectory = w.(bool)
											}

											if w, ok := flashBladeMapStrToI["export_rules"]; ok && !isIntfNil(w) {
												flashBlade.ExportRules = w.(string)
											}

											if v, ok := flashBladeMapStrToI["flash_blades"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												flashBlades := make([]*ves_io_schema_fleet.FlashBladeEndpoint, len(sl))
												flashBlade.FlashBlades = flashBlades
												for i, set := range sl {
													flashBlades[i] = &ves_io_schema_fleet.FlashBladeEndpoint{}

													flashBladesMapStrToI := set.(map[string]interface{})

													if v, ok := flashBladesMapStrToI["api_token"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														apiToken := &ves_io_schema.SecretType{}
														flashBlades[i].ApiToken = apiToken
														for _, set := range sl {

															apiTokenMapStrToI := set.(map[string]interface{})

															if v, ok := apiTokenMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

																apiToken.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

															}

															secretInfoOneofTypeFound := false

															if v, ok := apiTokenMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
																secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
																secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
																secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
																secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

													if w, ok := flashBladesMapStrToI["lables"]; ok && !isIntfNil(w) {
														ms := map[string]string{}
														for k, v := range w.(map[string]interface{}) {
															ms[k] = v.(string)
														}
														flashBlades[i].Lables = ms
													}

													mgmtEndpointTypeFound := false

													if v, ok := flashBladesMapStrToI["mgmt_dns_name"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_MgmtDnsName{}

														flashBlades[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtDnsName = v.(string)

													}

													if v, ok := flashBladesMapStrToI["mgmt_ip"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_MgmtIp{}

														flashBlades[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtIp = v.(string)

													}

													nfsEndpointTypeFound := false

													if v, ok := flashBladesMapStrToI["nfs_endpoint_dns_name"]; ok && !isIntfNil(v) && !nfsEndpointTypeFound {

														nfsEndpointTypeFound = true
														nfsEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_NfsEndpointDnsName{}

														flashBlades[i].NfsEndpoint = nfsEndpointInt

														nfsEndpointInt.NfsEndpointDnsName = v.(string)

													}

													if v, ok := flashBladesMapStrToI["nfs_endpoint_ip"]; ok && !isIntfNil(v) && !nfsEndpointTypeFound {

														nfsEndpointTypeFound = true
														nfsEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_NfsEndpointIp{}

														flashBlades[i].NfsEndpoint = nfsEndpointInt

														nfsEndpointInt.NfsEndpointIp = v.(string)

													}

												}

											}

										}

									}

								}

							}

							if v, ok := cs["cluster_id"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.ClusterId = v.(string)
							}

							if v, ok := cs["enable_storage_topology"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.EnableStorageTopology = v.(bool)
							}

							if v, ok := cs["enable_strict_topology"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.EnableStrictTopology = v.(bool)
							}

						}

					}

					if w, ok := storageDevicesMapStrToI["storage_device"]; ok && !isIntfNil(w) {
						storageDevices[i].StorageDevice = w.(string)
					}

				}

			}

		}

	}

	storageInterfaceChoiceTypeFound := false

	if v, ok := d.GetOk("no_storage_interfaces"); ok && !storageInterfaceChoiceTypeFound {

		storageInterfaceChoiceTypeFound = true

		if v.(bool) {
			storageInterfaceChoiceInt := &ves_io_schema_fleet.CreateSpecType_NoStorageInterfaces{}
			storageInterfaceChoiceInt.NoStorageInterfaces = &ves_io_schema.Empty{}
			createSpec.StorageInterfaceChoice = storageInterfaceChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_interface_list"); ok && !storageInterfaceChoiceTypeFound {

		storageInterfaceChoiceTypeFound = true
		storageInterfaceChoiceInt := &ves_io_schema_fleet.CreateSpecType_StorageInterfaceList{}
		storageInterfaceChoiceInt.StorageInterfaceList = &ves_io_schema_fleet.FleetInterfaceListType{}
		createSpec.StorageInterfaceChoice = storageInterfaceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["interfaces"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				interfacesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				storageInterfaceChoiceInt.StorageInterfaceList.Interfaces = interfacesInt
				for i, ps := range sl {

					iMapToStrVal := ps.(map[string]interface{})
					interfacesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
						interfacesInt[i].Name = v.(string)
					}

					if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						interfacesInt[i].Namespace = v.(string)
					}

					if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						interfacesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	storageStaticRoutesChoiceTypeFound := false

	if v, ok := d.GetOk("no_storage_static_routes"); ok && !storageStaticRoutesChoiceTypeFound {

		storageStaticRoutesChoiceTypeFound = true

		if v.(bool) {
			storageStaticRoutesChoiceInt := &ves_io_schema_fleet.CreateSpecType_NoStorageStaticRoutes{}
			storageStaticRoutesChoiceInt.NoStorageStaticRoutes = &ves_io_schema.Empty{}
			createSpec.StorageStaticRoutesChoice = storageStaticRoutesChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_static_routes"); ok && !storageStaticRoutesChoiceTypeFound {

		storageStaticRoutesChoiceTypeFound = true
		storageStaticRoutesChoiceInt := &ves_io_schema_fleet.CreateSpecType_StorageStaticRoutes{}
		storageStaticRoutesChoiceInt.StorageStaticRoutes = &ves_io_schema_fleet.FleetStorageStaticRoutesListType{}
		createSpec.StorageStaticRoutesChoice = storageStaticRoutesChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["storage_routes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				storageRoutes := make([]*ves_io_schema.StaticRouteType, len(sl))
				storageStaticRoutesChoiceInt.StorageStaticRoutes.StorageRoutes = storageRoutes
				for i, set := range sl {
					storageRoutes[i] = &ves_io_schema.StaticRouteType{}

					storageRoutesMapStrToI := set.(map[string]interface{})

					if v, ok := storageRoutesMapStrToI["attrs"]; ok && !isIntfNil(v) {

						attrsList := []ves_io_schema.RouteAttrType{}
						for _, j := range v.([]interface{}) {
							attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
						}
						storageRoutes[i].Attrs = attrsList

					}

					if w, ok := storageRoutesMapStrToI["labels"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						storageRoutes[i].Labels = ms
					}

					if v, ok := storageRoutesMapStrToI["nexthop"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						nexthop := &ves_io_schema.NextHopType{}
						storageRoutes[i].Nexthop = nexthop
						for _, set := range sl {

							nexthopMapStrToI := set.(map[string]interface{})

							if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								nexthop.Interface = intfInt
								for i, ps := range sl {

									iMapToStrVal := ps.(map[string]interface{})
									intfInt[i] = &ves_io_schema.ObjectRefType{}

									intfInt[i].Kind = "network_interface"

									if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
										intfInt[i].Name = v.(string)
									}

									if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										intfInt[i].Namespace = v.(string)
									}

									if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										intfInt[i].Tenant = v.(string)
									}

									if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
										intfInt[i].Uid = v.(string)
									}

								}

							}

							if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								nexthopAddress := &ves_io_schema.IpAddressType{}
								nexthop.NexthopAddress = nexthopAddress
								for _, set := range sl {

									nexthopAddressMapStrToI := set.(map[string]interface{})

									verTypeFound := false

									if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv4{}
										verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
										nexthopAddress.Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv4.Addr = v.(string)
											}

										}

									}

									if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv6{}
										verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
										nexthopAddress.Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv6.Addr = v.(string)
											}

										}

									}

								}

							}

							if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

								nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

							}

						}

					}

					if v, ok := storageRoutesMapStrToI["subnets"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
						storageRoutes[i].Subnets = subnets
						for i, set := range sl {
							subnets[i] = &ves_io_schema.IpSubnetType{}

							subnetsMapStrToI := set.(map[string]interface{})

							verTypeFound := false

							if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpSubnetType_Ipv4{}
								verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
								subnets[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["plen"]; ok && !isIntfNil(v) {

										verInt.Ipv4.Plen = uint32(v.(int))
									}

									if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

										verInt.Ipv4.Prefix = v.(string)
									}

								}

							}

							if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpSubnetType_Ipv6{}
								verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
								subnets[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["plen"]; ok && !isIntfNil(v) {

										verInt.Ipv6.Plen = uint32(v.(int))
									}

									if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

										verInt.Ipv6.Prefix = v.(string)
									}

								}

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		createSpec.VolterraSoftwareVersion =
			v.(string)
	}

	log.Printf("[DEBUG] Creating Volterra Fleet object with struct: %+v", createReq)

	createFleetResp, err := client.CreateObject(context.Background(), ves_io_schema_fleet.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Fleet: %s", err)
	}
	d.SetId(createFleetResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraFleetRead(d, meta)
}

func resourceVolterraFleetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_fleet.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Fleet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Fleet %q: %s", d.Id(), err)
	}
	return setFleetFields(client, d, resp)
}

func setFleetFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraFleetUpdate updates Fleet resource
func resourceVolterraFleetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_fleet.ReplaceSpecType{}
	updateReq := &ves_io_schema_fleet.ReplaceRequest{
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

	bondChoiceTypeFound := false

	if v, ok := d.GetOk("bond_device_list"); ok && !bondChoiceTypeFound {

		bondChoiceTypeFound = true
		bondChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_BondDeviceList{}
		bondChoiceInt.BondDeviceList = &ves_io_schema_fleet.FleetBondDevicesListType{}
		updateSpec.BondChoice = bondChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["bond_devices"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				bondDevices := make([]*ves_io_schema_fleet.FleetBondDeviceType, len(sl))
				bondChoiceInt.BondDeviceList.BondDevices = bondDevices
				for i, set := range sl {
					bondDevices[i] = &ves_io_schema_fleet.FleetBondDeviceType{}

					bondDevicesMapStrToI := set.(map[string]interface{})

					if w, ok := bondDevicesMapStrToI["devices"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						bondDevices[i].Devices = ls
					}

					lacpChoiceTypeFound := false

					if v, ok := bondDevicesMapStrToI["active_backup"]; ok && !isIntfNil(v) && !lacpChoiceTypeFound {

						lacpChoiceTypeFound = true

						if v.(bool) {
							lacpChoiceInt := &ves_io_schema_fleet.FleetBondDeviceType_ActiveBackup{}
							lacpChoiceInt.ActiveBackup = &ves_io_schema.Empty{}
							bondDevices[i].LacpChoice = lacpChoiceInt
						}

					}

					if v, ok := bondDevicesMapStrToI["lacp"]; ok && !isIntfNil(v) && !lacpChoiceTypeFound {

						lacpChoiceTypeFound = true
						lacpChoiceInt := &ves_io_schema_fleet.FleetBondDeviceType_Lacp{}
						lacpChoiceInt.Lacp = &ves_io_schema_fleet.BondLacpType{}
						bondDevices[i].LacpChoice = lacpChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["rate"]; ok && !isIntfNil(v) {

								lacpChoiceInt.Lacp.Rate = uint32(v.(int))
							}

						}

					}

					if w, ok := bondDevicesMapStrToI["link_polling_interval"]; ok && !isIntfNil(w) {
						bondDevices[i].LinkPollingInterval = w.(uint32)
					}

					if w, ok := bondDevicesMapStrToI["link_up_delay"]; ok && !isIntfNil(w) {
						bondDevices[i].LinkUpDelay = w.(uint32)
					}

					if w, ok := bondDevicesMapStrToI["name"]; ok && !isIntfNil(w) {
						bondDevices[i].Name = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("no_bond_devices"); ok && !bondChoiceTypeFound {

		bondChoiceTypeFound = true

		if v.(bool) {
			bondChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_NoBondDevices{}
			bondChoiceInt.NoBondDevices = &ves_io_schema.Empty{}
			updateSpec.BondChoice = bondChoiceInt
		}

	}

	dcClusterGroupChoiceTypeFound := false

	if v, ok := d.GetOk("dc_cluster_group"); ok && !dcClusterGroupChoiceTypeFound {

		dcClusterGroupChoiceTypeFound = true
		dcClusterGroupChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_DcClusterGroup{}
		dcClusterGroupChoiceInt.DcClusterGroup = &ves_io_schema_views.ObjectRefType{}
		updateSpec.DcClusterGroupChoice = dcClusterGroupChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroup.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroup.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroup.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("dc_cluster_group_inside"); ok && !dcClusterGroupChoiceTypeFound {

		dcClusterGroupChoiceTypeFound = true
		dcClusterGroupChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_DcClusterGroupInside{}
		dcClusterGroupChoiceInt.DcClusterGroupInside = &ves_io_schema_views.ObjectRefType{}
		updateSpec.DcClusterGroupChoice = dcClusterGroupChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroupInside.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroupInside.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				dcClusterGroupChoiceInt.DcClusterGroupInside.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("no_dc_cluster_group"); ok && !dcClusterGroupChoiceTypeFound {

		dcClusterGroupChoiceTypeFound = true

		if v.(bool) {
			dcClusterGroupChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_NoDcClusterGroup{}
			dcClusterGroupChoiceInt.NoDcClusterGroup = &ves_io_schema.Empty{}
			updateSpec.DcClusterGroupChoice = dcClusterGroupChoiceInt
		}

	}

	if v, ok := d.GetOk("enable_default_fleet_config_download"); ok && !isIntfNil(v) {

		updateSpec.EnableDefaultFleetConfigDownload =
			v.(bool)
	}

	gpuChoiceTypeFound := false

	if v, ok := d.GetOk("disable_gpu"); ok && !gpuChoiceTypeFound {

		gpuChoiceTypeFound = true

		if v.(bool) {
			gpuChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_DisableGpu{}
			gpuChoiceInt.DisableGpu = &ves_io_schema.Empty{}
			updateSpec.GpuChoice = gpuChoiceInt
		}

	}

	if v, ok := d.GetOk("enable_gpu"); ok && !gpuChoiceTypeFound {

		gpuChoiceTypeFound = true

		if v.(bool) {
			gpuChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_EnableGpu{}
			gpuChoiceInt.EnableGpu = &ves_io_schema.Empty{}
			updateSpec.GpuChoice = gpuChoiceInt
		}

	}

	if v, ok := d.GetOk("inside_virtual_network"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		insideVirtualNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.InsideVirtualNetwork = insideVirtualNetworkInt
		for i, ps := range sl {

			ivnMapToStrVal := ps.(map[string]interface{})
			insideVirtualNetworkInt[i] = &ves_io_schema.ObjectRefType{}

			insideVirtualNetworkInt[i].Kind = "virtual_network"

			if v, ok := ivnMapToStrVal["name"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Name = v.(string)
			}

			if v, ok := ivnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Namespace = v.(string)
			}

			if v, ok := ivnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Tenant = v.(string)
			}

			if v, ok := ivnMapToStrVal["uid"]; ok && !isIntfNil(v) {
				insideVirtualNetworkInt[i].Uid = v.(string)
			}

		}

	}

	interfaceChoiceTypeFound := false

	if v, ok := d.GetOk("default_config"); ok && !interfaceChoiceTypeFound {

		interfaceChoiceTypeFound = true

		if v.(bool) {
			interfaceChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_DefaultConfig{}
			interfaceChoiceInt.DefaultConfig = &ves_io_schema.Empty{}
			updateSpec.InterfaceChoice = interfaceChoiceInt
		}

	}

	if v, ok := d.GetOk("device_list"); ok && !interfaceChoiceTypeFound {

		interfaceChoiceTypeFound = true
		interfaceChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_DeviceList{}
		interfaceChoiceInt.DeviceList = &ves_io_schema_fleet.FleetDeviceListType{}
		updateSpec.InterfaceChoice = interfaceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["devices"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				devices := make([]*ves_io_schema_fleet.DeviceInstanceType, len(sl))
				interfaceChoiceInt.DeviceList.Devices = devices
				for i, set := range sl {
					devices[i] = &ves_io_schema_fleet.DeviceInstanceType{}

					devicesMapStrToI := set.(map[string]interface{})

					deviceInstanceTypeFound := false

					if v, ok := devicesMapStrToI["network_device"]; ok && !isIntfNil(v) && !deviceInstanceTypeFound {

						deviceInstanceTypeFound = true
						deviceInstanceInt := &ves_io_schema_fleet.DeviceInstanceType_NetworkDevice{}
						deviceInstanceInt.NetworkDevice = &ves_io_schema_fleet.NetworkingDeviceInstanceType{}
						devices[i].DeviceInstance = deviceInstanceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["interface"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								deviceInstanceInt.NetworkDevice.Interface = intfInt
								for i, ps := range sl {

									iMapToStrVal := ps.(map[string]interface{})
									intfInt[i] = &ves_io_schema.ObjectRefType{}

									intfInt[i].Kind = "network_interface"

									if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
										intfInt[i].Name = v.(string)
									}

									if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										intfInt[i].Namespace = v.(string)
									}

									if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										intfInt[i].Tenant = v.(string)
									}

									if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
										intfInt[i].Uid = v.(string)
									}

								}

							}

							if v, ok := cs["use"]; ok && !isIntfNil(v) {

								deviceInstanceInt.NetworkDevice.Use = ves_io_schema_fleet.NetworkingDeviceInstanceUseType(ves_io_schema_fleet.NetworkingDeviceInstanceUseType_value[v.(string)])

							}

						}

					}

					if w, ok := devicesMapStrToI["name"]; ok && !isIntfNil(w) {
						devices[i].Name = w.(string)
					}

					if v, ok := devicesMapStrToI["owner"]; ok && !isIntfNil(v) {

						devices[i].Owner = ves_io_schema_fleet.DeviceOwnerType(ves_io_schema_fleet.DeviceOwnerType_value[v.(string)])

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("interface_list"); ok && !interfaceChoiceTypeFound {

		interfaceChoiceTypeFound = true
		interfaceChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_InterfaceList{}
		interfaceChoiceInt.InterfaceList = &ves_io_schema_fleet.FleetInterfaceListType{}
		updateSpec.InterfaceChoice = interfaceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["interfaces"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				interfacesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				interfaceChoiceInt.InterfaceList.Interfaces = interfacesInt
				for i, ps := range sl {

					iMapToStrVal := ps.(map[string]interface{})
					interfacesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
						interfacesInt[i].Name = v.(string)
					}

					if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						interfacesInt[i].Namespace = v.(string)
					}

					if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						interfacesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("network_connectors"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		networkConnectorsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.NetworkConnectors = networkConnectorsInt
		for i, ps := range sl {

			ncMapToStrVal := ps.(map[string]interface{})
			networkConnectorsInt[i] = &ves_io_schema.ObjectRefType{}

			networkConnectorsInt[i].Kind = "network_connector"

			if v, ok := ncMapToStrVal["name"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Name = v.(string)
			}

			if v, ok := ncMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Namespace = v.(string)
			}

			if v, ok := ncMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Tenant = v.(string)
			}

			if v, ok := ncMapToStrVal["uid"]; ok && !isIntfNil(v) {
				networkConnectorsInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("network_firewall"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		networkFirewallInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.NetworkFirewall = networkFirewallInt
		for i, ps := range sl {

			nfMapToStrVal := ps.(map[string]interface{})
			networkFirewallInt[i] = &ves_io_schema.ObjectRefType{}

			networkFirewallInt[i].Kind = "network_firewall"

			if v, ok := nfMapToStrVal["name"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Name = v.(string)
			}

			if v, ok := nfMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Namespace = v.(string)
			}

			if v, ok := nfMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Tenant = v.(string)
			}

			if v, ok := nfMapToStrVal["uid"]; ok && !isIntfNil(v) {
				networkFirewallInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		updateSpec.OperatingSystemVersion =
			v.(string)
	}

	if v, ok := d.GetOk("outside_virtual_network"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		outsideVirtualNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.OutsideVirtualNetwork = outsideVirtualNetworkInt
		for i, ps := range sl {

			ovnMapToStrVal := ps.(map[string]interface{})
			outsideVirtualNetworkInt[i] = &ves_io_schema.ObjectRefType{}

			outsideVirtualNetworkInt[i].Kind = "virtual_network"

			if v, ok := ovnMapToStrVal["name"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Name = v.(string)
			}

			if v, ok := ovnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Namespace = v.(string)
			}

			if v, ok := ovnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Tenant = v.(string)
			}

			if v, ok := ovnMapToStrVal["uid"]; ok && !isIntfNil(v) {
				outsideVirtualNetworkInt[i].Uid = v.(string)
			}

		}

	}

	storageClassChoiceTypeFound := false

	if v, ok := d.GetOk("default_storage_class"); ok && !storageClassChoiceTypeFound {

		storageClassChoiceTypeFound = true

		if v.(bool) {
			storageClassChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_DefaultStorageClass{}
			storageClassChoiceInt.DefaultStorageClass = &ves_io_schema.Empty{}
			updateSpec.StorageClassChoice = storageClassChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_class_list"); ok && !storageClassChoiceTypeFound {

		storageClassChoiceTypeFound = true
		storageClassChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_StorageClassList{}
		storageClassChoiceInt.StorageClassList = &ves_io_schema_fleet.FleetStorageClassListType{}
		updateSpec.StorageClassChoice = storageClassChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["storage_classes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				storageClasses := make([]*ves_io_schema_fleet.FleetStorageClassType, len(sl))
				storageClassChoiceInt.StorageClassList.StorageClasses = storageClasses
				for i, set := range sl {
					storageClasses[i] = &ves_io_schema_fleet.FleetStorageClassType{}

					storageClassesMapStrToI := set.(map[string]interface{})

					if w, ok := storageClassesMapStrToI["advanced_storage_parameters"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						storageClasses[i].AdvancedStorageParameters = ms
					}

					if w, ok := storageClassesMapStrToI["default_storage_class"]; ok && !isIntfNil(w) {
						storageClasses[i].DefaultStorageClass = w.(bool)
					}

					if w, ok := storageClassesMapStrToI["description"]; ok && !isIntfNil(w) {
						storageClasses[i].Description = w.(string)
					}

					deviceChoiceTypeFound := false

					if v, ok := storageClassesMapStrToI["dell_emc_isilon_f800"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_DellEmcIsilonF800{}
						deviceChoiceInt.DellEmcIsilonF800 = &ves_io_schema_fleet.StorageClassDellIsilonF800Type{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["base_path"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.BasePath = v.(string)
							}

							httpsChoiceTypeFound := false

							if v, ok := cs["az_service_ip_address"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true
								httpsChoiceInt := &ves_io_schema_fleet.StorageClassDellIsilonF800Type_AzServiceIpAddress{}

								deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt

								httpsChoiceInt.AzServiceIpAddress = v.(string)

							}

							if v, ok := cs["az_service_name"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true
								httpsChoiceInt := &ves_io_schema_fleet.StorageClassDellIsilonF800Type_AzServiceName{}

								deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt

								httpsChoiceInt.AzServiceName = v.(string)

							}

							if v, ok := cs["iscsi_access_zone"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.IscsiAccessZone = v.(string)
							}

							if v, ok := cs["root_client_enable"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.RootClientEnable = v.(bool)
							}

						}

					}

					if v, ok := storageClassesMapStrToI["hpe_nimbus_storage_af40"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_HpeNimbusStorageAf40{}
						deviceChoiceInt.HpeNimbusStorageAf40 = &ves_io_schema_fleet.StorageClassHPENimbusStorageAf40Type{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["limit_iops"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitIops = uint32(v.(int))
							}

							if v, ok := cs["limit_mbps"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitMbps = uint32(v.(int))
							}

							if v, ok := cs["perf_policy"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.PerfPolicy = v.(string)
							}

						}

					}

					if v, ok := storageClassesMapStrToI["netapp_trident"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_NetappTrident{}
						deviceChoiceInt.NetappTrident = &ves_io_schema_fleet.StorageClassNetappTridentType{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["selector"]; ok && !isIntfNil(v) {

								ms := map[string]string{}
								for k, v := range v.(map[string]interface{}) {
									ms[k] = v.(string)
								}
								deviceChoiceInt.NetappTrident.Selector = ms
							}

						}

					}

					if v, ok := storageClassesMapStrToI["pure_service_orchestrator"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageClassType_PureServiceOrchestrator{}
						deviceChoiceInt.PureServiceOrchestrator = &ves_io_schema_fleet.StorageClassPureServiceOrchestratorType{}
						storageClasses[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["backend"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.Backend = v.(string)
							}

						}

					}

					if w, ok := storageClassesMapStrToI["storage_class_name"]; ok && !isIntfNil(w) {
						storageClasses[i].StorageClassName = w.(string)
					}

					if w, ok := storageClassesMapStrToI["storage_device"]; ok && !isIntfNil(w) {
						storageClasses[i].StorageDevice = w.(string)
					}

				}

			}

		}

	}

	storageDeviceChoiceTypeFound := false

	if v, ok := d.GetOk("no_storage_device"); ok && !storageDeviceChoiceTypeFound {

		storageDeviceChoiceTypeFound = true

		if v.(bool) {
			storageDeviceChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_NoStorageDevice{}
			storageDeviceChoiceInt.NoStorageDevice = &ves_io_schema.Empty{}
			updateSpec.StorageDeviceChoice = storageDeviceChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_device_list"); ok && !storageDeviceChoiceTypeFound {

		storageDeviceChoiceTypeFound = true
		storageDeviceChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_StorageDeviceList{}
		storageDeviceChoiceInt.StorageDeviceList = &ves_io_schema_fleet.FleetStorageDeviceListType{}
		updateSpec.StorageDeviceChoice = storageDeviceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["storage_devices"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				storageDevices := make([]*ves_io_schema_fleet.FleetStorageDeviceType, len(sl))
				storageDeviceChoiceInt.StorageDeviceList.StorageDevices = storageDevices
				for i, set := range sl {
					storageDevices[i] = &ves_io_schema_fleet.FleetStorageDeviceType{}

					storageDevicesMapStrToI := set.(map[string]interface{})

					if w, ok := storageDevicesMapStrToI["advanced_advanced_parameters"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						storageDevices[i].AdvancedAdvancedParameters = ms
					}

					deviceChoiceTypeFound := false

					if v, ok := storageDevicesMapStrToI["dell_emc_isilon_f800"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_DellEmcIsilonF800{}
						deviceChoiceInt.DellEmcIsilonF800 = &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							addressChoiceTypeFound := false

							if v, ok := cs["api_server_ip_address"]; ok && !isIntfNil(v) && !addressChoiceTypeFound {

								addressChoiceTypeFound = true
								addressChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_ApiServerIpAddress{}

								deviceChoiceInt.DellEmcIsilonF800.AddressChoice = addressChoiceInt

								addressChoiceInt.ApiServerIpAddress = v.(string)

							}

							if v, ok := cs["api_server_name"]; ok && !isIntfNil(v) && !addressChoiceTypeFound {

								addressChoiceTypeFound = true
								addressChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_ApiServerName{}

								deviceChoiceInt.DellEmcIsilonF800.AddressChoice = addressChoiceInt

								addressChoiceInt.ApiServerName = v.(string)

							}

							if v, ok := cs["api_server_port"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.ApiServerPort = uint32(v.(int))
							}

							if v, ok := cs["base_path"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.BasePath = v.(string)
							}

							httpsChoiceTypeFound := false

							if v, ok := cs["secure_network"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true

								if v.(bool) {
									httpsChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_SecureNetwork{}
									httpsChoiceInt.SecureNetwork = &ves_io_schema.Empty{}
									deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt
								}

							}

							if v, ok := cs["trusted_ca_url"]; ok && !isIntfNil(v) && !httpsChoiceTypeFound {

								httpsChoiceTypeFound = true
								httpsChoiceInt := &ves_io_schema_fleet.StorageDeviceDellIsilonF800Type_TrustedCaUrl{}

								deviceChoiceInt.DellEmcIsilonF800.HttpsChoice = httpsChoiceInt

								httpsChoiceInt.TrustedCaUrl = v.(string)

							}

							if v, ok := cs["iscsi_access_zone"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.IscsiAccessZone = v.(string)
							}

							if v, ok := cs["password"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								password := &ves_io_schema.SecretType{}
								deviceChoiceInt.DellEmcIsilonF800.Password = password
								for _, set := range sl {

									passwordMapStrToI := set.(map[string]interface{})

									if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["username"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.Username = v.(string)
							}

							if v, ok := cs["volume_prefix"]; ok && !isIntfNil(v) {

								deviceChoiceInt.DellEmcIsilonF800.VolumePrefix = v.(string)
							}

						}

					}

					if v, ok := storageDevicesMapStrToI["hpe_nimbus_storage_af40"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_HpeNimbusStorageAf40{}
						deviceChoiceInt.HpeNimbusStorageAf40 = &ves_io_schema_fleet.StorageDeviceHPENimbusStorageAf40Type{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["api_server_port"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.ApiServerPort = uint32(v.(int))
							}

							if v, ok := cs["limit_iops"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitIops = uint32(v.(int))
							}

							if v, ok := cs["limit_mbps"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.LimitMbps = uint32(v.(int))
							}

							if v, ok := cs["password"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								password := &ves_io_schema.SecretType{}
								deviceChoiceInt.HpeNimbusStorageAf40.Password = password
								for _, set := range sl {

									passwordMapStrToI := set.(map[string]interface{})

									if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										password.SecretInfoOneof = secretInfoOneofInt

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

							if v, ok := cs["perf_policy"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.PerfPolicy = v.(string)
							}

							if v, ok := cs["storage_server_ip_address"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.StorageServerIpAddress = v.(string)
							}

							if v, ok := cs["storage_server_name"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.StorageServerName = v.(string)
							}

							if v, ok := cs["username"]; ok && !isIntfNil(v) {

								deviceChoiceInt.HpeNimbusStorageAf40.Username = v.(string)
							}

						}

					}

					if v, ok := storageDevicesMapStrToI["netapp_trident"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_NetappTrident{}
						deviceChoiceInt.NetappTrident = &ves_io_schema_fleet.StorageDeviceNetappTridentType{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							backendChoiceTypeFound := false

							if v, ok := cs["netapp_backend_ontap_nas"]; ok && !isIntfNil(v) && !backendChoiceTypeFound {

								backendChoiceTypeFound = true
								backendChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappTridentType_NetappBackendOntapNas{}
								backendChoiceInt.NetappBackendOntapNas = &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType{}
								deviceChoiceInt.NetappTrident.BackendChoice = backendChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["auto_export_cidrs"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										autoExportCidrs := &ves_io_schema_views.PrefixStringListType{}
										backendChoiceInt.NetappBackendOntapNas.AutoExportCidrs = autoExportCidrs
										for _, set := range sl {

											autoExportCidrsMapStrToI := set.(map[string]interface{})

											if w, ok := autoExportCidrsMapStrToI["prefixes"]; ok && !isIntfNil(w) {
												ls := make([]string, len(w.([]interface{})))
												for i, v := range w.([]interface{}) {
													ls[i] = v.(string)
												}
												autoExportCidrs.Prefixes = ls
											}

										}

									}

									if v, ok := cs["auto_export_policy"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.AutoExportPolicy = v.(bool)
									}

									if v, ok := cs["backend_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.BackendName = v.(string)
									}

									dataLifTypeFound := false

									if v, ok := cs["data_lif_dns_name"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_DataLifDnsName{}

										backendChoiceInt.NetappBackendOntapNas.DataLif = dataLifInt

										dataLifInt.DataLifDnsName = v.(string)

									}

									if v, ok := cs["data_lif_ip"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_DataLifIp{}

										backendChoiceInt.NetappBackendOntapNas.DataLif = dataLifInt

										dataLifInt.DataLifIp = v.(string)

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										backendChoiceInt.NetappBackendOntapNas.Labels = ms
									}

									if v, ok := cs["limit_aggregate_usage"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.LimitAggregateUsage = v.(string)
									}

									if v, ok := cs["limit_volume_size"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.LimitVolumeSize = v.(string)
									}

									managementLifTypeFound := false

									if v, ok := cs["management_lif_dns_name"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_ManagementLifDnsName{}

										backendChoiceInt.NetappBackendOntapNas.ManagementLif = managementLifInt

										managementLifInt.ManagementLifDnsName = v.(string)

									}

									if v, ok := cs["management_lif_ip"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapNasType_ManagementLifIp{}

										backendChoiceInt.NetappBackendOntapNas.ManagementLif = managementLifInt

										managementLifInt.ManagementLifIp = v.(string)

									}

									if v, ok := cs["nfs_mount_options"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.NfsMountOptions = v.(string)
									}

									if v, ok := cs["password"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										password := &ves_io_schema.SecretType{}
										backendChoiceInt.NetappBackendOntapNas.Password = password
										for _, set := range sl {

											passwordMapStrToI := set.(map[string]interface{})

											if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := cs["region"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.Region = v.(string)
									}

									if v, ok := cs["storage"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										storage := make([]*ves_io_schema_fleet.OntapVirtualStoragePoolType, len(sl))
										backendChoiceInt.NetappBackendOntapNas.Storage = storage
										for i, set := range sl {
											storage[i] = &ves_io_schema_fleet.OntapVirtualStoragePoolType{}

											storageMapStrToI := set.(map[string]interface{})

											if w, ok := storageMapStrToI["labels"]; ok && !isIntfNil(w) {
												ms := map[string]string{}
												for k, v := range w.(map[string]interface{}) {
													ms[k] = v.(string)
												}
												storage[i].Labels = ms
											}

											if v, ok := storageMapStrToI["volume_defaults"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
												storage[i].VolumeDefaults = volumeDefaults
												for _, set := range sl {

													volumeDefaultsMapStrToI := set.(map[string]interface{})

													if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
														volumeDefaults.Encryption = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.ExportPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
														volumeDefaults.SecurityStyle = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotDir = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SpaceReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
														volumeDefaults.SplitOnClone = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.TieringPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
														volumeDefaults.UnixPermissions = w.(int32)
													}

												}

											}

											if w, ok := storageMapStrToI["zone"]; ok && !isIntfNil(w) {
												storage[i].Zone = w.(string)
											}

										}

									}

									if v, ok := cs["storage_driver_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.StorageDriverName = v.(string)
									}

									if v, ok := cs["storage_prefix"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.StoragePrefix = v.(string)
									}

									if v, ok := cs["svm"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.Svm = v.(string)
									}

									if v, ok := cs["username"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapNas.Username = v.(string)
									}

									if v, ok := cs["volume_defaults"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
										backendChoiceInt.NetappBackendOntapNas.VolumeDefaults = volumeDefaults
										for _, set := range sl {

											volumeDefaultsMapStrToI := set.(map[string]interface{})

											if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
												volumeDefaults.Encryption = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.ExportPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
												volumeDefaults.SecurityStyle = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotDir = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SpaceReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
												volumeDefaults.SplitOnClone = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.TieringPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
												volumeDefaults.UnixPermissions = w.(int32)
											}

										}

									}

								}

							}

							if v, ok := cs["netapp_backend_ontap_san"]; ok && !isIntfNil(v) && !backendChoiceTypeFound {

								backendChoiceTypeFound = true
								backendChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappTridentType_NetappBackendOntapSan{}
								backendChoiceInt.NetappBackendOntapSan = &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType{}
								deviceChoiceInt.NetappTrident.BackendChoice = backendChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									chapChoiceTypeFound := false

									if v, ok := cs["no_chap"]; ok && !isIntfNil(v) && !chapChoiceTypeFound {

										chapChoiceTypeFound = true

										if v.(bool) {
											chapChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_NoChap{}
											chapChoiceInt.NoChap = &ves_io_schema.Empty{}
											backendChoiceInt.NetappBackendOntapSan.ChapChoice = chapChoiceInt
										}

									}

									if v, ok := cs["use_chap"]; ok && !isIntfNil(v) && !chapChoiceTypeFound {

										chapChoiceTypeFound = true
										chapChoiceInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_UseChap{}
										chapChoiceInt.UseChap = &ves_io_schema_fleet.DeviceNetappBackendOntapSanChapType{}
										backendChoiceInt.NetappBackendOntapSan.ChapChoice = chapChoiceInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["chap_initiator_secret"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												chapInitiatorSecret := &ves_io_schema.SecretType{}
												chapChoiceInt.UseChap.ChapInitiatorSecret = chapInitiatorSecret
												for _, set := range sl {

													chapInitiatorSecretMapStrToI := set.(map[string]interface{})

													if v, ok := chapInitiatorSecretMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

														chapInitiatorSecret.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

													}

													secretInfoOneofTypeFound := false

													if v, ok := chapInitiatorSecretMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
														secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapInitiatorSecretMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
														secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapInitiatorSecretMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
														secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapInitiatorSecretMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
														secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
														chapInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := cs["chap_target_initiator_secret"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												chapTargetInitiatorSecret := &ves_io_schema.SecretType{}
												chapChoiceInt.UseChap.ChapTargetInitiatorSecret = chapTargetInitiatorSecret
												for _, set := range sl {

													chapTargetInitiatorSecretMapStrToI := set.(map[string]interface{})

													if v, ok := chapTargetInitiatorSecretMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

														chapTargetInitiatorSecret.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

													}

													secretInfoOneofTypeFound := false

													if v, ok := chapTargetInitiatorSecretMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
														secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapTargetInitiatorSecretMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
														secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapTargetInitiatorSecretMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
														secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

													if v, ok := chapTargetInitiatorSecretMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

														secretInfoOneofTypeFound = true
														secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
														secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
														chapTargetInitiatorSecret.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := cs["chap_target_username"]; ok && !isIntfNil(v) {

												chapChoiceInt.UseChap.ChapTargetUsername = v.(string)
											}

											if v, ok := cs["chap_username"]; ok && !isIntfNil(v) {

												chapChoiceInt.UseChap.ChapUsername = v.(string)
											}

										}

									}

									dataLifTypeFound := false

									if v, ok := cs["data_lif_dns_name"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_DataLifDnsName{}

										backendChoiceInt.NetappBackendOntapSan.DataLif = dataLifInt

										dataLifInt.DataLifDnsName = v.(string)

									}

									if v, ok := cs["data_lif_ip"]; ok && !isIntfNil(v) && !dataLifTypeFound {

										dataLifTypeFound = true
										dataLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_DataLifIp{}

										backendChoiceInt.NetappBackendOntapSan.DataLif = dataLifInt

										dataLifInt.DataLifIp = v.(string)

									}

									if v, ok := cs["igroup_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.IgroupName = v.(string)
									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										backendChoiceInt.NetappBackendOntapSan.Labels = ms
									}

									if v, ok := cs["limit_aggregate_usage"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.LimitAggregateUsage = uint32(v.(int))
									}

									if v, ok := cs["limit_volume_size"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.LimitVolumeSize = v.(int32)
									}

									managementLifTypeFound := false

									if v, ok := cs["management_lif_dns_name"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_ManagementLifDnsName{}

										backendChoiceInt.NetappBackendOntapSan.ManagementLif = managementLifInt

										managementLifInt.ManagementLifDnsName = v.(string)

									}

									if v, ok := cs["management_lif_ip"]; ok && !isIntfNil(v) && !managementLifTypeFound {

										managementLifTypeFound = true
										managementLifInt := &ves_io_schema_fleet.StorageDeviceNetappBackendOntapSanType_ManagementLifIp{}

										backendChoiceInt.NetappBackendOntapSan.ManagementLif = managementLifInt

										managementLifInt.ManagementLifIp = v.(string)

									}

									if v, ok := cs["password"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										password := &ves_io_schema.SecretType{}
										backendChoiceInt.NetappBackendOntapSan.Password = password
										for _, set := range sl {

											passwordMapStrToI := set.(map[string]interface{})

											if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

												password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

											}

											secretInfoOneofTypeFound := false

											if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
												secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
												secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
												secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

											if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

												secretInfoOneofTypeFound = true
												secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
												secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
												password.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := cs["region"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.Region = v.(string)
									}

									if v, ok := cs["storage"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										storage := make([]*ves_io_schema_fleet.OntapVirtualStoragePoolType, len(sl))
										backendChoiceInt.NetappBackendOntapSan.Storage = storage
										for i, set := range sl {
											storage[i] = &ves_io_schema_fleet.OntapVirtualStoragePoolType{}

											storageMapStrToI := set.(map[string]interface{})

											if w, ok := storageMapStrToI["labels"]; ok && !isIntfNil(w) {
												ms := map[string]string{}
												for k, v := range w.(map[string]interface{}) {
													ms[k] = v.(string)
												}
												storage[i].Labels = ms
											}

											if v, ok := storageMapStrToI["volume_defaults"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
												storage[i].VolumeDefaults = volumeDefaults
												for _, set := range sl {

													volumeDefaultsMapStrToI := set.(map[string]interface{})

													if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
														volumeDefaults.Encryption = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.ExportPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
														volumeDefaults.SecurityStyle = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotDir = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SnapshotReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
														volumeDefaults.SpaceReserve = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
														volumeDefaults.SplitOnClone = w.(bool)
													}

													if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
														volumeDefaults.TieringPolicy = w.(string)
													}

													if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
														volumeDefaults.UnixPermissions = w.(int32)
													}

												}

											}

											if w, ok := storageMapStrToI["zone"]; ok && !isIntfNil(w) {
												storage[i].Zone = w.(string)
											}

										}

									}

									if v, ok := cs["storage_driver_name"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.StorageDriverName = v.(string)
									}

									if v, ok := cs["storage_prefix"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.StoragePrefix = v.(string)
									}

									if v, ok := cs["svm"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.Svm = v.(string)
									}

									if v, ok := cs["username"]; ok && !isIntfNil(v) {

										backendChoiceInt.NetappBackendOntapSan.Username = v.(string)
									}

									if v, ok := cs["volume_defaults"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										volumeDefaults := &ves_io_schema_fleet.OntapVolumeDefaults{}
										backendChoiceInt.NetappBackendOntapSan.VolumeDefaults = volumeDefaults
										for _, set := range sl {

											volumeDefaultsMapStrToI := set.(map[string]interface{})

											if w, ok := volumeDefaultsMapStrToI["encryption"]; ok && !isIntfNil(w) {
												volumeDefaults.Encryption = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["export_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.ExportPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["security_style"]; ok && !isIntfNil(w) {
												volumeDefaults.SecurityStyle = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_dir"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotDir = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["snapshot_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SnapshotReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["space_reserve"]; ok && !isIntfNil(w) {
												volumeDefaults.SpaceReserve = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["split_on_clone"]; ok && !isIntfNil(w) {
												volumeDefaults.SplitOnClone = w.(bool)
											}

											if w, ok := volumeDefaultsMapStrToI["tiering_policy"]; ok && !isIntfNil(w) {
												volumeDefaults.TieringPolicy = w.(string)
											}

											if w, ok := volumeDefaultsMapStrToI["unix_permissions"]; ok && !isIntfNil(w) {
												volumeDefaults.UnixPermissions = w.(int32)
											}

										}

									}

								}

							}

						}

					}

					if v, ok := storageDevicesMapStrToI["pure_service_orchestrator"]; ok && !isIntfNil(v) && !deviceChoiceTypeFound {

						deviceChoiceTypeFound = true
						deviceChoiceInt := &ves_io_schema_fleet.FleetStorageDeviceType_PureServiceOrchestrator{}
						deviceChoiceInt.PureServiceOrchestrator = &ves_io_schema_fleet.StorageDevicePureStorageServiceOrchestratorType{}
						storageDevices[i].DeviceChoice = deviceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["arrays"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								arrays := &ves_io_schema_fleet.PsoArrayConfiguration{}
								deviceChoiceInt.PureServiceOrchestrator.Arrays = arrays
								for _, set := range sl {

									arraysMapStrToI := set.(map[string]interface{})

									if v, ok := arraysMapStrToI["flash_array"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										flashArray := &ves_io_schema_fleet.FlashArrayType{}
										arrays.FlashArray = flashArray
										for _, set := range sl {

											flashArrayMapStrToI := set.(map[string]interface{})

											if w, ok := flashArrayMapStrToI["default_fs_opt"]; ok && !isIntfNil(w) {
												flashArray.DefaultFsOpt = w.(string)
											}

											if w, ok := flashArrayMapStrToI["default_fs_type"]; ok && !isIntfNil(w) {
												flashArray.DefaultFsType = w.(string)
											}

											if w, ok := flashArrayMapStrToI["default_mount_opts"]; ok && !isIntfNil(w) {
												ls := make([]string, len(w.([]interface{})))
												for i, v := range w.([]interface{}) {
													ls[i] = v.(string)
												}
												flashArray.DefaultMountOpts = ls
											}

											if w, ok := flashArrayMapStrToI["disable_preempt_attachments"]; ok && !isIntfNil(w) {
												flashArray.DisablePreemptAttachments = w.(bool)
											}

											if v, ok := flashArrayMapStrToI["flash_arrays"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												flashArrays := make([]*ves_io_schema_fleet.FlashArrayEndpoint, len(sl))
												flashArray.FlashArrays = flashArrays
												for i, set := range sl {
													flashArrays[i] = &ves_io_schema_fleet.FlashArrayEndpoint{}

													flashArraysMapStrToI := set.(map[string]interface{})

													if v, ok := flashArraysMapStrToI["api_token"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														apiToken := &ves_io_schema.SecretType{}
														flashArrays[i].ApiToken = apiToken
														for _, set := range sl {

															apiTokenMapStrToI := set.(map[string]interface{})

															if v, ok := apiTokenMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

																apiToken.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

															}

															secretInfoOneofTypeFound := false

															if v, ok := apiTokenMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
																secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
																secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
																secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
																secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

													if w, ok := flashArraysMapStrToI["labels"]; ok && !isIntfNil(w) {
														ms := map[string]string{}
														for k, v := range w.(map[string]interface{}) {
															ms[k] = v.(string)
														}
														flashArrays[i].Labels = ms
													}

													mgmtEndpointTypeFound := false

													if v, ok := flashArraysMapStrToI["mgmt_dns_name"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashArrayEndpoint_MgmtDnsName{}

														flashArrays[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtDnsName = v.(string)

													}

													if v, ok := flashArraysMapStrToI["mgmt_ip"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashArrayEndpoint_MgmtIp{}

														flashArrays[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtIp = v.(string)

													}

												}

											}

											if w, ok := flashArrayMapStrToI["iscsi_login_timeout"]; ok && !isIntfNil(w) {
												flashArray.IscsiLoginTimeout = w.(int32)
											}

											if w, ok := flashArrayMapStrToI["san_type"]; ok && !isIntfNil(w) {
												flashArray.SanType = w.(string)
											}

										}

									}

									if v, ok := arraysMapStrToI["flash_blade"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										flashBlade := &ves_io_schema_fleet.FlashBladeType{}
										arrays.FlashBlade = flashBlade
										for _, set := range sl {

											flashBladeMapStrToI := set.(map[string]interface{})

											if w, ok := flashBladeMapStrToI["enable_snapshot_directory"]; ok && !isIntfNil(w) {
												flashBlade.EnableSnapshotDirectory = w.(bool)
											}

											if w, ok := flashBladeMapStrToI["export_rules"]; ok && !isIntfNil(w) {
												flashBlade.ExportRules = w.(string)
											}

											if v, ok := flashBladeMapStrToI["flash_blades"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												flashBlades := make([]*ves_io_schema_fleet.FlashBladeEndpoint, len(sl))
												flashBlade.FlashBlades = flashBlades
												for i, set := range sl {
													flashBlades[i] = &ves_io_schema_fleet.FlashBladeEndpoint{}

													flashBladesMapStrToI := set.(map[string]interface{})

													if v, ok := flashBladesMapStrToI["api_token"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														apiToken := &ves_io_schema.SecretType{}
														flashBlades[i].ApiToken = apiToken
														for _, set := range sl {

															apiTokenMapStrToI := set.(map[string]interface{})

															if v, ok := apiTokenMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

																apiToken.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

															}

															secretInfoOneofTypeFound := false

															if v, ok := apiTokenMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
																secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
																secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
																secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

															if v, ok := apiTokenMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

																secretInfoOneofTypeFound = true
																secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
																secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
																apiToken.SecretInfoOneof = secretInfoOneofInt

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

													if w, ok := flashBladesMapStrToI["lables"]; ok && !isIntfNil(w) {
														ms := map[string]string{}
														for k, v := range w.(map[string]interface{}) {
															ms[k] = v.(string)
														}
														flashBlades[i].Lables = ms
													}

													mgmtEndpointTypeFound := false

													if v, ok := flashBladesMapStrToI["mgmt_dns_name"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_MgmtDnsName{}

														flashBlades[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtDnsName = v.(string)

													}

													if v, ok := flashBladesMapStrToI["mgmt_ip"]; ok && !isIntfNil(v) && !mgmtEndpointTypeFound {

														mgmtEndpointTypeFound = true
														mgmtEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_MgmtIp{}

														flashBlades[i].MgmtEndpoint = mgmtEndpointInt

														mgmtEndpointInt.MgmtIp = v.(string)

													}

													nfsEndpointTypeFound := false

													if v, ok := flashBladesMapStrToI["nfs_endpoint_dns_name"]; ok && !isIntfNil(v) && !nfsEndpointTypeFound {

														nfsEndpointTypeFound = true
														nfsEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_NfsEndpointDnsName{}

														flashBlades[i].NfsEndpoint = nfsEndpointInt

														nfsEndpointInt.NfsEndpointDnsName = v.(string)

													}

													if v, ok := flashBladesMapStrToI["nfs_endpoint_ip"]; ok && !isIntfNil(v) && !nfsEndpointTypeFound {

														nfsEndpointTypeFound = true
														nfsEndpointInt := &ves_io_schema_fleet.FlashBladeEndpoint_NfsEndpointIp{}

														flashBlades[i].NfsEndpoint = nfsEndpointInt

														nfsEndpointInt.NfsEndpointIp = v.(string)

													}

												}

											}

										}

									}

								}

							}

							if v, ok := cs["cluster_id"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.ClusterId = v.(string)
							}

							if v, ok := cs["enable_storage_topology"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.EnableStorageTopology = v.(bool)
							}

							if v, ok := cs["enable_strict_topology"]; ok && !isIntfNil(v) {

								deviceChoiceInt.PureServiceOrchestrator.EnableStrictTopology = v.(bool)
							}

						}

					}

					if w, ok := storageDevicesMapStrToI["storage_device"]; ok && !isIntfNil(w) {
						storageDevices[i].StorageDevice = w.(string)
					}

				}

			}

		}

	}

	storageInterfaceChoiceTypeFound := false

	if v, ok := d.GetOk("no_storage_interfaces"); ok && !storageInterfaceChoiceTypeFound {

		storageInterfaceChoiceTypeFound = true

		if v.(bool) {
			storageInterfaceChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_NoStorageInterfaces{}
			storageInterfaceChoiceInt.NoStorageInterfaces = &ves_io_schema.Empty{}
			updateSpec.StorageInterfaceChoice = storageInterfaceChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_interface_list"); ok && !storageInterfaceChoiceTypeFound {

		storageInterfaceChoiceTypeFound = true
		storageInterfaceChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_StorageInterfaceList{}
		storageInterfaceChoiceInt.StorageInterfaceList = &ves_io_schema_fleet.FleetInterfaceListType{}
		updateSpec.StorageInterfaceChoice = storageInterfaceChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["interfaces"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				interfacesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				storageInterfaceChoiceInt.StorageInterfaceList.Interfaces = interfacesInt
				for i, ps := range sl {

					iMapToStrVal := ps.(map[string]interface{})
					interfacesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
						interfacesInt[i].Name = v.(string)
					}

					if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						interfacesInt[i].Namespace = v.(string)
					}

					if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						interfacesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	storageStaticRoutesChoiceTypeFound := false

	if v, ok := d.GetOk("no_storage_static_routes"); ok && !storageStaticRoutesChoiceTypeFound {

		storageStaticRoutesChoiceTypeFound = true

		if v.(bool) {
			storageStaticRoutesChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_NoStorageStaticRoutes{}
			storageStaticRoutesChoiceInt.NoStorageStaticRoutes = &ves_io_schema.Empty{}
			updateSpec.StorageStaticRoutesChoice = storageStaticRoutesChoiceInt
		}

	}

	if v, ok := d.GetOk("storage_static_routes"); ok && !storageStaticRoutesChoiceTypeFound {

		storageStaticRoutesChoiceTypeFound = true
		storageStaticRoutesChoiceInt := &ves_io_schema_fleet.ReplaceSpecType_StorageStaticRoutes{}
		storageStaticRoutesChoiceInt.StorageStaticRoutes = &ves_io_schema_fleet.FleetStorageStaticRoutesListType{}
		updateSpec.StorageStaticRoutesChoice = storageStaticRoutesChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["storage_routes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				storageRoutes := make([]*ves_io_schema.StaticRouteType, len(sl))
				storageStaticRoutesChoiceInt.StorageStaticRoutes.StorageRoutes = storageRoutes
				for i, set := range sl {
					storageRoutes[i] = &ves_io_schema.StaticRouteType{}

					storageRoutesMapStrToI := set.(map[string]interface{})

					if v, ok := storageRoutesMapStrToI["attrs"]; ok && !isIntfNil(v) {

						attrsList := []ves_io_schema.RouteAttrType{}
						for _, j := range v.([]interface{}) {
							attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
						}
						storageRoutes[i].Attrs = attrsList

					}

					if w, ok := storageRoutesMapStrToI["labels"]; ok && !isIntfNil(w) {
						ms := map[string]string{}
						for k, v := range w.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						storageRoutes[i].Labels = ms
					}

					if v, ok := storageRoutesMapStrToI["nexthop"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						nexthop := &ves_io_schema.NextHopType{}
						storageRoutes[i].Nexthop = nexthop
						for _, set := range sl {

							nexthopMapStrToI := set.(map[string]interface{})

							if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								nexthop.Interface = intfInt
								for i, ps := range sl {

									iMapToStrVal := ps.(map[string]interface{})
									intfInt[i] = &ves_io_schema.ObjectRefType{}

									intfInt[i].Kind = "network_interface"

									if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
										intfInt[i].Name = v.(string)
									}

									if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										intfInt[i].Namespace = v.(string)
									}

									if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										intfInt[i].Tenant = v.(string)
									}

									if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
										intfInt[i].Uid = v.(string)
									}

								}

							}

							if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								nexthopAddress := &ves_io_schema.IpAddressType{}
								nexthop.NexthopAddress = nexthopAddress
								for _, set := range sl {

									nexthopAddressMapStrToI := set.(map[string]interface{})

									verTypeFound := false

									if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv4{}
										verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
										nexthopAddress.Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv4.Addr = v.(string)
											}

										}

									}

									if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv6{}
										verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
										nexthopAddress.Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv6.Addr = v.(string)
											}

										}

									}

								}

							}

							if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

								nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

							}

						}

					}

					if v, ok := storageRoutesMapStrToI["subnets"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
						storageRoutes[i].Subnets = subnets
						for i, set := range sl {
							subnets[i] = &ves_io_schema.IpSubnetType{}

							subnetsMapStrToI := set.(map[string]interface{})

							verTypeFound := false

							if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpSubnetType_Ipv4{}
								verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
								subnets[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["plen"]; ok && !isIntfNil(v) {

										verInt.Ipv4.Plen = uint32(v.(int))
									}

									if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

										verInt.Ipv4.Prefix = v.(string)
									}

								}

							}

							if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpSubnetType_Ipv6{}
								verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
								subnets[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["plen"]; ok && !isIntfNil(v) {

										verInt.Ipv6.Plen = uint32(v.(int))
									}

									if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

										verInt.Ipv6.Prefix = v.(string)
									}

								}

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		updateSpec.VolterraSoftwareVersion =
			v.(string)
	}

	log.Printf("[DEBUG] Updating Volterra Fleet obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_fleet.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Fleet: %s", err)
	}

	return resourceVolterraFleetRead(d, meta)
}

func resourceVolterraFleetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_fleet.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Fleet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Fleet before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Fleet obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_fleet.ObjectType, namespace, name)
}
