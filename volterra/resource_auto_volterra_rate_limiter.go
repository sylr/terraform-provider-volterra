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
	ves_io_schema_rate_limiter "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/rate_limiter"
)

// resourceVolterraRateLimiter is implementation of Volterra's RateLimiter resources
func resourceVolterraRateLimiter() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraRateLimiterCreate,
		Read:   resourceVolterraRateLimiterRead,
		Update: resourceVolterraRateLimiterUpdate,
		Delete: resourceVolterraRateLimiterDelete,

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

			"limits": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"total_number": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"unit": {
							Type:     schema.TypeString,
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
		},
	}
}

// resourceVolterraRateLimiterCreate creates RateLimiter resource
func resourceVolterraRateLimiterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_rate_limiter.CreateSpecType{}
	createReq := &ves_io_schema_rate_limiter.CreateRequest{
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

	if v, ok := d.GetOk("limits"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		limits := make([]*ves_io_schema_rate_limiter.RateLimitValue, len(sl))
		createSpec.Limits = limits
		for i, set := range sl {
			limits[i] = &ves_io_schema_rate_limiter.RateLimitValue{}

			limitsMapStrToI := set.(map[string]interface{})

			if w, ok := limitsMapStrToI["total_number"]; ok && !isIntfNil(w) {
				limits[i].TotalNumber = w.(uint32)
			}

			if v, ok := limitsMapStrToI["unit"]; ok && !isIntfNil(v) {

				limits[i].Unit = ves_io_schema_rate_limiter.RateLimitPeriodUnit(ves_io_schema_rate_limiter.RateLimitPeriodUnit_value[v.(string)])

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

	log.Printf("[DEBUG] Creating Volterra RateLimiter object with struct: %+v", createReq)

	createRateLimiterResp, err := client.CreateObject(context.Background(), ves_io_schema_rate_limiter.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating RateLimiter: %s", err)
	}
	d.SetId(createRateLimiterResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraRateLimiterRead(d, meta)
}

func resourceVolterraRateLimiterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_rate_limiter.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] RateLimiter %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra RateLimiter %q: %s", d.Id(), err)
	}
	return setRateLimiterFields(client, d, resp)
}

func setRateLimiterFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraRateLimiterUpdate updates RateLimiter resource
func resourceVolterraRateLimiterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_rate_limiter.ReplaceSpecType{}
	updateReq := &ves_io_schema_rate_limiter.ReplaceRequest{
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

	if v, ok := d.GetOk("limits"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		limits := make([]*ves_io_schema_rate_limiter.RateLimitValue, len(sl))
		updateSpec.Limits = limits
		for i, set := range sl {
			limits[i] = &ves_io_schema_rate_limiter.RateLimitValue{}

			limitsMapStrToI := set.(map[string]interface{})

			if w, ok := limitsMapStrToI["total_number"]; ok && !isIntfNil(w) {
				limits[i].TotalNumber = w.(uint32)
			}

			if v, ok := limitsMapStrToI["unit"]; ok && !isIntfNil(v) {

				limits[i].Unit = ves_io_schema_rate_limiter.RateLimitPeriodUnit(ves_io_schema_rate_limiter.RateLimitPeriodUnit_value[v.(string)])

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

	log.Printf("[DEBUG] Updating Volterra RateLimiter obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_rate_limiter.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating RateLimiter: %s", err)
	}

	return resourceVolterraRateLimiterRead(d, meta)
}

func resourceVolterraRateLimiterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_rate_limiter.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] RateLimiter %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra RateLimiter before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra RateLimiter obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_rate_limiter.ObjectType, namespace, name)
}