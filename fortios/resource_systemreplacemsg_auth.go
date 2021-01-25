// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Replacement messages.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemReplacemsgAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgAuthCreate,
		Read:   resourceSystemReplacemsgAuthRead,
		Update: resourceSystemReplacemsgAuthUpdate,
		Delete: resourceSystemReplacemsgAuthDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"msg_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 28),
				ForceNew:     true,
				Required:     true,
			},
			"buffer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),
				Optional:     true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemReplacemsgAuthCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemReplacemsgAuth(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgAuth resource while getting object: %v", err)
	}

	o, err := c.CreateSystemReplacemsgAuth(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgAuth resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgAuth")
	}

	return resourceSystemReplacemsgAuthRead(d, m)
}

func resourceSystemReplacemsgAuthUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemReplacemsgAuth(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgAuth resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemReplacemsgAuth(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgAuth resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgAuth")
	}

	return resourceSystemReplacemsgAuthRead(d, m)
}

func resourceSystemReplacemsgAuthDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemReplacemsgAuth(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgAuth resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgAuthRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemReplacemsgAuth(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgAuth resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgAuth(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgAuth resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgAuthMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgAuthBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgAuthHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgAuthFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgAuth(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("msg_type", flattenSystemReplacemsgAuthMsgType(o["msg-type"], d, "msg_type")); err != nil {
		if !fortiAPIPatch(o["msg-type"]) {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	if err = d.Set("buffer", flattenSystemReplacemsgAuthBuffer(o["buffer"], d, "buffer")); err != nil {
		if !fortiAPIPatch(o["buffer"]) {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgAuthHeader(o["header"], d, "header")); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgAuthFormat(o["format"], d, "format")); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgAuthFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReplacemsgAuthMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgAuthBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgAuthHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgAuthFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgAuth(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("msg_type"); ok {
		t, err := expandSystemReplacemsgAuthMsgType(d, v, "msg_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	if v, ok := d.GetOk("buffer"); ok {
		t, err := expandSystemReplacemsgAuthBuffer(d, v, "buffer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok {
		t, err := expandSystemReplacemsgAuthHeader(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok {
		t, err := expandSystemReplacemsgAuthFormat(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	return &obj, nil
}
