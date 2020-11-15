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
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_secret_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/secret_policy"
)

// resourceVolterraSecretPolicy is implementation of Volterra's SecretPolicy resources
func resourceVolterraSecretPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSecretPolicyCreate,
		Read:   resourceVolterraSecretPolicyRead,
		Update: resourceVolterraSecretPolicyUpdate,
		Delete: resourceVolterraSecretPolicyDelete,

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

			"algo": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"allow_volterra": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"rules": {

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
	}
}

// resourceVolterraSecretPolicyCreate creates SecretPolicy resource
func resourceVolterraSecretPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_secret_policy.CreateSpecType{}
	createReq := &ves_io_schema_secret_policy.CreateRequest{
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

	if v, ok := d.GetOk("algo"); ok && !isIntfNil(v) {

		createSpec.Algo = ves_io_schema_policy.RuleCombiningAlgorithm(ves_io_schema_policy.RuleCombiningAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("allow_volterra"); ok && !isIntfNil(v) {

		createSpec.AllowVolterra =
			v.(bool)
	}

	if v, ok := d.GetOk("rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Rules = rulesInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			rulesInt[i] = &ves_io_schema.ObjectRefType{}

			rulesInt[i].Kind = "secret_policy_rule"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				rulesInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rulesInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rulesInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rulesInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra SecretPolicy object with struct: %+v", createReq)

	createSecretPolicyResp, err := client.CreateObject(context.Background(), ves_io_schema_secret_policy.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating SecretPolicy: %s", err)
	}
	d.SetId(createSecretPolicyResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraSecretPolicyRead(d, meta)
}

func resourceVolterraSecretPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_secret_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] SecretPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra SecretPolicy %q: %s", d.Id(), err)
	}
	return setSecretPolicyFields(client, d, resp)
}

func setSecretPolicyFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraSecretPolicyUpdate updates SecretPolicy resource
func resourceVolterraSecretPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_secret_policy.ReplaceSpecType{}
	updateReq := &ves_io_schema_secret_policy.ReplaceRequest{
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

	if v, ok := d.GetOk("algo"); ok && !isIntfNil(v) {

		updateSpec.Algo = ves_io_schema_policy.RuleCombiningAlgorithm(ves_io_schema_policy.RuleCombiningAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("allow_volterra"); ok && !isIntfNil(v) {

		updateSpec.AllowVolterra =
			v.(bool)
	}

	if v, ok := d.GetOk("rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Rules = rulesInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			rulesInt[i] = &ves_io_schema.ObjectRefType{}

			rulesInt[i].Kind = "secret_policy_rule"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				rulesInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rulesInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rulesInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rulesInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra SecretPolicy obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_secret_policy.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating SecretPolicy: %s", err)
	}

	return resourceVolterraSecretPolicyRead(d, meta)
}

func resourceVolterraSecretPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_secret_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] SecretPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra SecretPolicy before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra SecretPolicy obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_secret_policy.ObjectType, namespace, name)
}