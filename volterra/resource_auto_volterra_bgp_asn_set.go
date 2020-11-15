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
	ves_io_schema_bgp_asn_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/bgp_asn_set"
)

// resourceVolterraBgpAsnSet is implementation of Volterra's BgpAsnSet resources
func resourceVolterraBgpAsnSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraBgpAsnSetCreate,
		Read:   resourceVolterraBgpAsnSetRead,
		Update: resourceVolterraBgpAsnSetUpdate,
		Delete: resourceVolterraBgpAsnSetDelete,

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

			"as_numbers": {

				Type: schema.TypeList,

				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

// resourceVolterraBgpAsnSetCreate creates BgpAsnSet resource
func resourceVolterraBgpAsnSetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_bgp_asn_set.CreateSpecType{}
	createReq := &ves_io_schema_bgp_asn_set.CreateRequest{
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

	if v, ok := d.GetOk("as_numbers"); ok && !isIntfNil(v) {

		ls := make([]uint32, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {

			ls[i] = uint32(v.(int))
		}
		createSpec.AsNumbers = ls

	}

	log.Printf("[DEBUG] Creating Volterra BgpAsnSet object with struct: %+v", createReq)

	createBgpAsnSetResp, err := client.CreateObject(context.Background(), ves_io_schema_bgp_asn_set.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating BgpAsnSet: %s", err)
	}
	d.SetId(createBgpAsnSetResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraBgpAsnSetRead(d, meta)
}

func resourceVolterraBgpAsnSetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_bgp_asn_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] BgpAsnSet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra BgpAsnSet %q: %s", d.Id(), err)
	}
	return setBgpAsnSetFields(client, d, resp)
}

func setBgpAsnSetFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraBgpAsnSetUpdate updates BgpAsnSet resource
func resourceVolterraBgpAsnSetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_bgp_asn_set.ReplaceSpecType{}
	updateReq := &ves_io_schema_bgp_asn_set.ReplaceRequest{
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

	if v, ok := d.GetOk("as_numbers"); ok && !isIntfNil(v) {

		ls := make([]uint32, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {

			ls[i] = uint32(v.(int))
		}
		updateSpec.AsNumbers = ls

	}

	log.Printf("[DEBUG] Updating Volterra BgpAsnSet obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_bgp_asn_set.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating BgpAsnSet: %s", err)
	}

	return resourceVolterraBgpAsnSetRead(d, meta)
}

func resourceVolterraBgpAsnSetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_bgp_asn_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] BgpAsnSet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra BgpAsnSet before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra BgpAsnSet obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_bgp_asn_set.ObjectType, namespace, name)
}