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
	ves_io_schema_network_connector "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_connector"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraNetworkConnector is implementation of Volterra's NetworkConnector resources
func resourceVolterraNetworkConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNetworkConnectorCreate,
		Read:   resourceVolterraNetworkConnectorRead,
		Update: resourceVolterraNetworkConnectorUpdate,
		Delete: resourceVolterraNetworkConnectorDelete,

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

			"sli_to_global_dr": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"global_vn": {

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
					},
				},
			},

			"sli_to_global_snat": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"global_vn": {

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

						"snat_config": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_ip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"snat_pool": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"snat_pool_allocator": {

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

									"default_gw_snat": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"dynamic_routing": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"sli_to_slo_dr": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"sli_to_slo_snat": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"snat_pool": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"snat_pool_allocator": {

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

						"default_gw_snat": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"dynamic_routing": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"slo_to_global_dr": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"global_vn": {

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
					},
				},
			},

			"slo_to_global_snat": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"global_vn": {

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

						"snat_config": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_ip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"snat_pool": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"snat_pool_allocator": {

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

									"default_gw_snat": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"dynamic_routing": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"disable_forward_proxy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_forward_proxy": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"max_connect_attempts": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"white_listed_ports": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						"white_listed_prefixes": {

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
	}
}

// resourceVolterraNetworkConnectorCreate creates NetworkConnector resource
func resourceVolterraNetworkConnectorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_network_connector.CreateSpecType{}
	createReq := &ves_io_schema_network_connector.CreateRequest{
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

	connectorChoiceTypeFound := false

	if v, ok := d.GetOk("sli_to_global_dr"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.CreateSpecType_SliToGlobalDr{}
		connectorChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
		createSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SliToGlobalDr.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("sli_to_global_snat"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.CreateSpecType_SliToGlobalSnat{}
		connectorChoiceInt.SliToGlobalSnat = &ves_io_schema_network_connector.GlobalSnatConnectorType{}
		createSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SliToGlobalSnat.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

			if v, ok := cs["snat_config"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				snatConfig := &ves_io_schema_network_connector.SnatConnectorType{}
				connectorChoiceInt.SliToGlobalSnat.SnatConfig = snatConfig
				for _, set := range sl {

					snatConfigMapStrToI := set.(map[string]interface{})

					poolChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["interface_ip"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true

						if v.(bool) {
							poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_InterfaceIp{}
							poolChoiceInt.InterfaceIp = &ves_io_schema.Empty{}
							snatConfig.PoolChoice = poolChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["snat_pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPool{}

						snatConfig.PoolChoice = poolChoiceInt

						poolChoiceInt.SnatPool = v.(string)

					}

					if v, ok := snatConfigMapStrToI["snat_pool_allocator"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPoolAllocator{}
						poolChoiceInt.SnatPoolAllocator = &ves_io_schema_views.ObjectRefType{}
						snatConfig.PoolChoice = poolChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Tenant = v.(string)
							}

						}

					}

					routingChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["default_gw_snat"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DefaultGwSnat{}
							routingChoiceInt.DefaultGwSnat = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["dynamic_routing"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DynamicRouting{}
							routingChoiceInt.DynamicRouting = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("sli_to_slo_dr"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true

		if v.(bool) {
			connectorChoiceInt := &ves_io_schema_network_connector.CreateSpecType_SliToSloDr{}
			connectorChoiceInt.SliToSloDr = &ves_io_schema.Empty{}
			createSpec.ConnectorChoice = connectorChoiceInt
		}

	}

	if v, ok := d.GetOk("sli_to_slo_snat"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.CreateSpecType_SliToSloSnat{}
		connectorChoiceInt.SliToSloSnat = &ves_io_schema_network_connector.SnatConnectorType{}
		createSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			poolChoiceTypeFound := false

			if v, ok := cs["interface_ip"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true

				if v.(bool) {
					poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_InterfaceIp{}
					poolChoiceInt.InterfaceIp = &ves_io_schema.Empty{}
					connectorChoiceInt.SliToSloSnat.PoolChoice = poolChoiceInt
				}

			}

			if v, ok := cs["snat_pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPool{}

				connectorChoiceInt.SliToSloSnat.PoolChoice = poolChoiceInt

				poolChoiceInt.SnatPool = v.(string)

			}

			if v, ok := cs["snat_pool_allocator"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPoolAllocator{}
				poolChoiceInt.SnatPoolAllocator = &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SliToSloSnat.PoolChoice = poolChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						poolChoiceInt.SnatPoolAllocator.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						poolChoiceInt.SnatPoolAllocator.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						poolChoiceInt.SnatPoolAllocator.Tenant = v.(string)
					}

				}

			}

			routingChoiceTypeFound := false

			if v, ok := cs["default_gw_snat"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

				routingChoiceTypeFound = true

				if v.(bool) {
					routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DefaultGwSnat{}
					routingChoiceInt.DefaultGwSnat = &ves_io_schema.Empty{}
					connectorChoiceInt.SliToSloSnat.RoutingChoice = routingChoiceInt
				}

			}

			if v, ok := cs["dynamic_routing"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

				routingChoiceTypeFound = true

				if v.(bool) {
					routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DynamicRouting{}
					routingChoiceInt.DynamicRouting = &ves_io_schema.Empty{}
					connectorChoiceInt.SliToSloSnat.RoutingChoice = routingChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("slo_to_global_dr"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.CreateSpecType_SloToGlobalDr{}
		connectorChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
		createSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SloToGlobalDr.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("slo_to_global_snat"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.CreateSpecType_SloToGlobalSnat{}
		connectorChoiceInt.SloToGlobalSnat = &ves_io_schema_network_connector.GlobalSnatConnectorType{}
		createSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SloToGlobalSnat.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

			if v, ok := cs["snat_config"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				snatConfig := &ves_io_schema_network_connector.SnatConnectorType{}
				connectorChoiceInt.SloToGlobalSnat.SnatConfig = snatConfig
				for _, set := range sl {

					snatConfigMapStrToI := set.(map[string]interface{})

					poolChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["interface_ip"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true

						if v.(bool) {
							poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_InterfaceIp{}
							poolChoiceInt.InterfaceIp = &ves_io_schema.Empty{}
							snatConfig.PoolChoice = poolChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["snat_pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPool{}

						snatConfig.PoolChoice = poolChoiceInt

						poolChoiceInt.SnatPool = v.(string)

					}

					if v, ok := snatConfigMapStrToI["snat_pool_allocator"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPoolAllocator{}
						poolChoiceInt.SnatPoolAllocator = &ves_io_schema_views.ObjectRefType{}
						snatConfig.PoolChoice = poolChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Tenant = v.(string)
							}

						}

					}

					routingChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["default_gw_snat"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DefaultGwSnat{}
							routingChoiceInt.DefaultGwSnat = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["dynamic_routing"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DynamicRouting{}
							routingChoiceInt.DynamicRouting = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

				}

			}

		}

	}

	forwardProxyChoiceTypeFound := false

	if v, ok := d.GetOk("disable_forward_proxy"); ok && !forwardProxyChoiceTypeFound {

		forwardProxyChoiceTypeFound = true

		if v.(bool) {
			forwardProxyChoiceInt := &ves_io_schema_network_connector.CreateSpecType_DisableForwardProxy{}
			forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
			createSpec.ForwardProxyChoice = forwardProxyChoiceInt
		}

	}

	if v, ok := d.GetOk("enable_forward_proxy"); ok && !forwardProxyChoiceTypeFound {

		forwardProxyChoiceTypeFound = true
		forwardProxyChoiceInt := &ves_io_schema_network_connector.CreateSpecType_EnableForwardProxy{}
		forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
		createSpec.ForwardProxyChoice = forwardProxyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

				forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
			}

			if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

				forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
			}

			if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

				ls := make([]uint32, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

			}

			if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra NetworkConnector object with struct: %+v", createReq)

	createNetworkConnectorResp, err := client.CreateObject(context.Background(), ves_io_schema_network_connector.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating NetworkConnector: %s", err)
	}
	d.SetId(createNetworkConnectorResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraNetworkConnectorRead(d, meta)
}

func resourceVolterraNetworkConnectorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_network_connector.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkConnector %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkConnector %q: %s", d.Id(), err)
	}
	return setNetworkConnectorFields(client, d, resp)
}

func setNetworkConnectorFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraNetworkConnectorUpdate updates NetworkConnector resource
func resourceVolterraNetworkConnectorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_network_connector.ReplaceSpecType{}
	updateReq := &ves_io_schema_network_connector.ReplaceRequest{
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

	connectorChoiceTypeFound := false

	if v, ok := d.GetOk("sli_to_global_dr"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_SliToGlobalDr{}
		connectorChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
		updateSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SliToGlobalDr.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("sli_to_global_snat"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_SliToGlobalSnat{}
		connectorChoiceInt.SliToGlobalSnat = &ves_io_schema_network_connector.GlobalSnatConnectorType{}
		updateSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SliToGlobalSnat.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

			if v, ok := cs["snat_config"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				snatConfig := &ves_io_schema_network_connector.SnatConnectorType{}
				connectorChoiceInt.SliToGlobalSnat.SnatConfig = snatConfig
				for _, set := range sl {

					snatConfigMapStrToI := set.(map[string]interface{})

					poolChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["interface_ip"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true

						if v.(bool) {
							poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_InterfaceIp{}
							poolChoiceInt.InterfaceIp = &ves_io_schema.Empty{}
							snatConfig.PoolChoice = poolChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["snat_pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPool{}

						snatConfig.PoolChoice = poolChoiceInt

						poolChoiceInt.SnatPool = v.(string)

					}

					if v, ok := snatConfigMapStrToI["snat_pool_allocator"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPoolAllocator{}
						poolChoiceInt.SnatPoolAllocator = &ves_io_schema_views.ObjectRefType{}
						snatConfig.PoolChoice = poolChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Tenant = v.(string)
							}

						}

					}

					routingChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["default_gw_snat"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DefaultGwSnat{}
							routingChoiceInt.DefaultGwSnat = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["dynamic_routing"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DynamicRouting{}
							routingChoiceInt.DynamicRouting = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("sli_to_slo_dr"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true

		if v.(bool) {
			connectorChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_SliToSloDr{}
			connectorChoiceInt.SliToSloDr = &ves_io_schema.Empty{}
			updateSpec.ConnectorChoice = connectorChoiceInt
		}

	}

	if v, ok := d.GetOk("sli_to_slo_snat"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_SliToSloSnat{}
		connectorChoiceInt.SliToSloSnat = &ves_io_schema_network_connector.SnatConnectorType{}
		updateSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			poolChoiceTypeFound := false

			if v, ok := cs["interface_ip"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true

				if v.(bool) {
					poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_InterfaceIp{}
					poolChoiceInt.InterfaceIp = &ves_io_schema.Empty{}
					connectorChoiceInt.SliToSloSnat.PoolChoice = poolChoiceInt
				}

			}

			if v, ok := cs["snat_pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPool{}

				connectorChoiceInt.SliToSloSnat.PoolChoice = poolChoiceInt

				poolChoiceInt.SnatPool = v.(string)

			}

			if v, ok := cs["snat_pool_allocator"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPoolAllocator{}
				poolChoiceInt.SnatPoolAllocator = &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SliToSloSnat.PoolChoice = poolChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						poolChoiceInt.SnatPoolAllocator.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						poolChoiceInt.SnatPoolAllocator.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						poolChoiceInt.SnatPoolAllocator.Tenant = v.(string)
					}

				}

			}

			routingChoiceTypeFound := false

			if v, ok := cs["default_gw_snat"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

				routingChoiceTypeFound = true

				if v.(bool) {
					routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DefaultGwSnat{}
					routingChoiceInt.DefaultGwSnat = &ves_io_schema.Empty{}
					connectorChoiceInt.SliToSloSnat.RoutingChoice = routingChoiceInt
				}

			}

			if v, ok := cs["dynamic_routing"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

				routingChoiceTypeFound = true

				if v.(bool) {
					routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DynamicRouting{}
					routingChoiceInt.DynamicRouting = &ves_io_schema.Empty{}
					connectorChoiceInt.SliToSloSnat.RoutingChoice = routingChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("slo_to_global_dr"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_SloToGlobalDr{}
		connectorChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
		updateSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SloToGlobalDr.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("slo_to_global_snat"); ok && !connectorChoiceTypeFound {

		connectorChoiceTypeFound = true
		connectorChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_SloToGlobalSnat{}
		connectorChoiceInt.SloToGlobalSnat = &ves_io_schema_network_connector.GlobalSnatConnectorType{}
		updateSpec.ConnectorChoice = connectorChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				globalVn := &ves_io_schema_views.ObjectRefType{}
				connectorChoiceInt.SloToGlobalSnat.GlobalVn = globalVn
				for _, set := range sl {

					globalVnMapStrToI := set.(map[string]interface{})

					if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
						globalVn.Name = w.(string)
					}

					if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
						globalVn.Namespace = w.(string)
					}

					if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
						globalVn.Tenant = w.(string)
					}

				}

			}

			if v, ok := cs["snat_config"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				snatConfig := &ves_io_schema_network_connector.SnatConnectorType{}
				connectorChoiceInt.SloToGlobalSnat.SnatConfig = snatConfig
				for _, set := range sl {

					snatConfigMapStrToI := set.(map[string]interface{})

					poolChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["interface_ip"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true

						if v.(bool) {
							poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_InterfaceIp{}
							poolChoiceInt.InterfaceIp = &ves_io_schema.Empty{}
							snatConfig.PoolChoice = poolChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["snat_pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPool{}

						snatConfig.PoolChoice = poolChoiceInt

						poolChoiceInt.SnatPool = v.(string)

					}

					if v, ok := snatConfigMapStrToI["snat_pool_allocator"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

						poolChoiceTypeFound = true
						poolChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_SnatPoolAllocator{}
						poolChoiceInt.SnatPoolAllocator = &ves_io_schema_views.ObjectRefType{}
						snatConfig.PoolChoice = poolChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								poolChoiceInt.SnatPoolAllocator.Tenant = v.(string)
							}

						}

					}

					routingChoiceTypeFound := false

					if v, ok := snatConfigMapStrToI["default_gw_snat"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DefaultGwSnat{}
							routingChoiceInt.DefaultGwSnat = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

					if v, ok := snatConfigMapStrToI["dynamic_routing"]; ok && !isIntfNil(v) && !routingChoiceTypeFound {

						routingChoiceTypeFound = true

						if v.(bool) {
							routingChoiceInt := &ves_io_schema_network_connector.SnatConnectorType_DynamicRouting{}
							routingChoiceInt.DynamicRouting = &ves_io_schema.Empty{}
							snatConfig.RoutingChoice = routingChoiceInt
						}

					}

				}

			}

		}

	}

	forwardProxyChoiceTypeFound := false

	if v, ok := d.GetOk("disable_forward_proxy"); ok && !forwardProxyChoiceTypeFound {

		forwardProxyChoiceTypeFound = true

		if v.(bool) {
			forwardProxyChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_DisableForwardProxy{}
			forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
			updateSpec.ForwardProxyChoice = forwardProxyChoiceInt
		}

	}

	if v, ok := d.GetOk("enable_forward_proxy"); ok && !forwardProxyChoiceTypeFound {

		forwardProxyChoiceTypeFound = true
		forwardProxyChoiceInt := &ves_io_schema_network_connector.ReplaceSpecType_EnableForwardProxy{}
		forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
		updateSpec.ForwardProxyChoice = forwardProxyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

				forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
			}

			if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

				forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
			}

			if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

				ls := make([]uint32, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

			}

			if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra NetworkConnector obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_network_connector.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating NetworkConnector: %s", err)
	}

	return resourceVolterraNetworkConnectorRead(d, meta)
}

func resourceVolterraNetworkConnectorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_network_connector.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkConnector %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkConnector before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra NetworkConnector obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_network_connector.ObjectType, namespace, name)
}