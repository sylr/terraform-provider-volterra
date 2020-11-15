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
	ves_io_schema_alert_receiver "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_receiver"
)

// resourceVolterraAlertReceiver is implementation of Volterra's AlertReceiver resources
func resourceVolterraAlertReceiver() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAlertReceiverCreate,
		Read:   resourceVolterraAlertReceiverRead,
		Update: resourceVolterraAlertReceiverUpdate,
		Delete: resourceVolterraAlertReceiverDelete,

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

			"email": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"email": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"opsgenie": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_key": {

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

						"url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"pagerduty": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"routing_key": {

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

						"url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"slack": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"channel": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"url": {

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

			"sms": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"contact_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraAlertReceiverCreate creates AlertReceiver resource
func resourceVolterraAlertReceiverCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_alert_receiver.CreateSpecType{}
	createReq := &ves_io_schema_alert_receiver.CreateRequest{
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

	receiverTypeFound := false

	if v, ok := d.GetOk("email"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.CreateSpecType_Email{}
		receiverInt.Email = &ves_io_schema_alert_receiver.EmailConfig{}
		createSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["email"]; ok && !isIntfNil(v) {

				receiverInt.Email.Email = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("opsgenie"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.CreateSpecType_Opsgenie{}
		receiverInt.Opsgenie = &ves_io_schema_alert_receiver.OpsGenieConfig{}
		createSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["api_key"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				apiKey := &ves_io_schema.SecretType{}
				receiverInt.Opsgenie.ApiKey = apiKey
				for _, set := range sl {

					apiKeyMapStrToI := set.(map[string]interface{})

					if v, ok := apiKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						apiKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := apiKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := apiKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := apiKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := apiKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

			if v, ok := cs["url"]; ok && !isIntfNil(v) {

				receiverInt.Opsgenie.Url = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("pagerduty"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.CreateSpecType_Pagerduty{}
		receiverInt.Pagerduty = &ves_io_schema_alert_receiver.PagerDutyConfig{}
		createSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["routing_key"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				routingKey := &ves_io_schema.SecretType{}
				receiverInt.Pagerduty.RoutingKey = routingKey
				for _, set := range sl {

					routingKeyMapStrToI := set.(map[string]interface{})

					if v, ok := routingKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						routingKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := routingKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := routingKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := routingKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := routingKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

			if v, ok := cs["url"]; ok && !isIntfNil(v) {

				receiverInt.Pagerduty.Url = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("slack"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.CreateSpecType_Slack{}
		receiverInt.Slack = &ves_io_schema_alert_receiver.SlackConfig{}
		createSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["channel"]; ok && !isIntfNil(v) {

				receiverInt.Slack.Channel = v.(string)
			}

			if v, ok := cs["url"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				url := &ves_io_schema.SecretType{}
				receiverInt.Slack.Url = url
				for _, set := range sl {

					urlMapStrToI := set.(map[string]interface{})

					if v, ok := urlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						url.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := urlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := urlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := urlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := urlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

	if v, ok := d.GetOk("sms"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.CreateSpecType_Sms{}
		receiverInt.Sms = &ves_io_schema_alert_receiver.SMSConfig{}
		createSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["contact_number"]; ok && !isIntfNil(v) {

				receiverInt.Sms.ContactNumber = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra AlertReceiver object with struct: %+v", createReq)

	createAlertReceiverResp, err := client.CreateObject(context.Background(), ves_io_schema_alert_receiver.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AlertReceiver: %s", err)
	}
	d.SetId(createAlertReceiverResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAlertReceiverRead(d, meta)
}

func resourceVolterraAlertReceiverRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_alert_receiver.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AlertReceiver %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AlertReceiver %q: %s", d.Id(), err)
	}
	return setAlertReceiverFields(client, d, resp)
}

func setAlertReceiverFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAlertReceiverUpdate updates AlertReceiver resource
func resourceVolterraAlertReceiverUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_alert_receiver.ReplaceSpecType{}
	updateReq := &ves_io_schema_alert_receiver.ReplaceRequest{
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

	receiverTypeFound := false

	if v, ok := d.GetOk("email"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.ReplaceSpecType_Email{}
		receiverInt.Email = &ves_io_schema_alert_receiver.EmailConfig{}
		updateSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["email"]; ok && !isIntfNil(v) {

				receiverInt.Email.Email = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("opsgenie"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.ReplaceSpecType_Opsgenie{}
		receiverInt.Opsgenie = &ves_io_schema_alert_receiver.OpsGenieConfig{}
		updateSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["api_key"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				apiKey := &ves_io_schema.SecretType{}
				receiverInt.Opsgenie.ApiKey = apiKey
				for _, set := range sl {

					apiKeyMapStrToI := set.(map[string]interface{})

					if v, ok := apiKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						apiKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := apiKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := apiKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := apiKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := apiKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						apiKey.SecretInfoOneof = secretInfoOneofInt

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

			if v, ok := cs["url"]; ok && !isIntfNil(v) {

				receiverInt.Opsgenie.Url = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("pagerduty"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.ReplaceSpecType_Pagerduty{}
		receiverInt.Pagerduty = &ves_io_schema_alert_receiver.PagerDutyConfig{}
		updateSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["routing_key"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				routingKey := &ves_io_schema.SecretType{}
				receiverInt.Pagerduty.RoutingKey = routingKey
				for _, set := range sl {

					routingKeyMapStrToI := set.(map[string]interface{})

					if v, ok := routingKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						routingKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := routingKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := routingKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := routingKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := routingKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						routingKey.SecretInfoOneof = secretInfoOneofInt

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

			if v, ok := cs["url"]; ok && !isIntfNil(v) {

				receiverInt.Pagerduty.Url = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("slack"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.ReplaceSpecType_Slack{}
		receiverInt.Slack = &ves_io_schema_alert_receiver.SlackConfig{}
		updateSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["channel"]; ok && !isIntfNil(v) {

				receiverInt.Slack.Channel = v.(string)
			}

			if v, ok := cs["url"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				url := &ves_io_schema.SecretType{}
				receiverInt.Slack.Url = url
				for _, set := range sl {

					urlMapStrToI := set.(map[string]interface{})

					if v, ok := urlMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						url.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := urlMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := urlMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := urlMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

					if v, ok := urlMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						url.SecretInfoOneof = secretInfoOneofInt

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

	if v, ok := d.GetOk("sms"); ok && !receiverTypeFound {

		receiverTypeFound = true
		receiverInt := &ves_io_schema_alert_receiver.ReplaceSpecType_Sms{}
		receiverInt.Sms = &ves_io_schema_alert_receiver.SMSConfig{}
		updateSpec.Receiver = receiverInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["contact_number"]; ok && !isIntfNil(v) {

				receiverInt.Sms.ContactNumber = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra AlertReceiver obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_alert_receiver.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AlertReceiver: %s", err)
	}

	return resourceVolterraAlertReceiverRead(d, meta)
}

func resourceVolterraAlertReceiverDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_alert_receiver.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AlertReceiver %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AlertReceiver before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AlertReceiver obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_alert_receiver.ObjectType, namespace, name)
}