// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Filters for memory buffer.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogMemoryFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemoryFilterUpdate,
		Read:   resourceLogMemoryFilterRead,
		Update: resourceLogMemoryFilterUpdate,
		Delete: resourceLogMemoryFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sniffer_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anomaly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netscan_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netscan_vulnerability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_log_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_log_adm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_log_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldb_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_activity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_memory_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogMemoryFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogMemoryFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogMemoryFilter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogMemoryFilter")
	}

	return resourceLogMemoryFilterRead(d, m)
}

func resourceLogMemoryFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogMemoryFilter(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogMemoryFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemoryFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogMemoryFilter(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemoryFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogMemoryFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterNetscanDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterNetscanVulnerability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterSsh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterRadius(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterDhcp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterPpp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterSslvpnLogAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterSslvpnLogAdm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterSslvpnLogSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterVipSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterLdbMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterWanOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterWirelessActivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterCpuMemoryUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemoryFilterFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogMemoryFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("severity", flattenLogMemoryFilterSeverity(o["severity"], d, "severity")); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogMemoryFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic")); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogMemoryFilterLocalTraffic(o["local-traffic"], d, "local_traffic")); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogMemoryFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic")); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogMemoryFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic")); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogMemoryFilterAnomaly(o["anomaly"], d, "anomaly")); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("netscan_discovery", flattenLogMemoryFilterNetscanDiscovery(o["netscan-discovery"], d, "netscan_discovery")); err != nil {
		if !fortiAPIPatch(o["netscan-discovery"]) {
			return fmt.Errorf("Error reading netscan_discovery: %v", err)
		}
	}

	if err = d.Set("netscan_vulnerability", flattenLogMemoryFilterNetscanVulnerability(o["netscan-vulnerability"], d, "netscan_vulnerability")); err != nil {
		if !fortiAPIPatch(o["netscan-vulnerability"]) {
			return fmt.Errorf("Error reading netscan_vulnerability: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogMemoryFilterVoip(o["voip"], d, "voip")); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogMemoryFilterGtp(o["gtp"], d, "gtp")); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("dns", flattenLogMemoryFilterDns(o["dns"], d, "dns")); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("ssh", flattenLogMemoryFilterSsh(o["ssh"], d, "ssh")); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("event", flattenLogMemoryFilterEvent(o["event"], d, "event")); err != nil {
		if !fortiAPIPatch(o["event"]) {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("system", flattenLogMemoryFilterSystem(o["system"], d, "system")); err != nil {
		if !fortiAPIPatch(o["system"]) {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("radius", flattenLogMemoryFilterRadius(o["radius"], d, "radius")); err != nil {
		if !fortiAPIPatch(o["radius"]) {
			return fmt.Errorf("Error reading radius: %v", err)
		}
	}

	if err = d.Set("ipsec", flattenLogMemoryFilterIpsec(o["ipsec"], d, "ipsec")); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("Error reading ipsec: %v", err)
		}
	}

	if err = d.Set("dhcp", flattenLogMemoryFilterDhcp(o["dhcp"], d, "dhcp")); err != nil {
		if !fortiAPIPatch(o["dhcp"]) {
			return fmt.Errorf("Error reading dhcp: %v", err)
		}
	}

	if err = d.Set("ppp", flattenLogMemoryFilterPpp(o["ppp"], d, "ppp")); err != nil {
		if !fortiAPIPatch(o["ppp"]) {
			return fmt.Errorf("Error reading ppp: %v", err)
		}
	}

	if err = d.Set("admin", flattenLogMemoryFilterAdmin(o["admin"], d, "admin")); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("ha", flattenLogMemoryFilterHa(o["ha"], d, "ha")); err != nil {
		if !fortiAPIPatch(o["ha"]) {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("auth", flattenLogMemoryFilterAuth(o["auth"], d, "auth")); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("pattern", flattenLogMemoryFilterPattern(o["pattern"], d, "pattern")); err != nil {
		if !fortiAPIPatch(o["pattern"]) {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_auth", flattenLogMemoryFilterSslvpnLogAuth(o["sslvpn-log-auth"], d, "sslvpn_log_auth")); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-auth"]) {
			return fmt.Errorf("Error reading sslvpn_log_auth: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_adm", flattenLogMemoryFilterSslvpnLogAdm(o["sslvpn-log-adm"], d, "sslvpn_log_adm")); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-adm"]) {
			return fmt.Errorf("Error reading sslvpn_log_adm: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_session", flattenLogMemoryFilterSslvpnLogSession(o["sslvpn-log-session"], d, "sslvpn_log_session")); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-session"]) {
			return fmt.Errorf("Error reading sslvpn_log_session: %v", err)
		}
	}

	if err = d.Set("vip_ssl", flattenLogMemoryFilterVipSsl(o["vip-ssl"], d, "vip_ssl")); err != nil {
		if !fortiAPIPatch(o["vip-ssl"]) {
			return fmt.Errorf("Error reading vip_ssl: %v", err)
		}
	}

	if err = d.Set("ldb_monitor", flattenLogMemoryFilterLdbMonitor(o["ldb-monitor"], d, "ldb_monitor")); err != nil {
		if !fortiAPIPatch(o["ldb-monitor"]) {
			return fmt.Errorf("Error reading ldb_monitor: %v", err)
		}
	}

	if err = d.Set("wan_opt", flattenLogMemoryFilterWanOpt(o["wan-opt"], d, "wan_opt")); err != nil {
		if !fortiAPIPatch(o["wan-opt"]) {
			return fmt.Errorf("Error reading wan_opt: %v", err)
		}
	}

	if err = d.Set("wireless_activity", flattenLogMemoryFilterWirelessActivity(o["wireless-activity"], d, "wireless_activity")); err != nil {
		if !fortiAPIPatch(o["wireless-activity"]) {
			return fmt.Errorf("Error reading wireless_activity: %v", err)
		}
	}

	if err = d.Set("cpu_memory_usage", flattenLogMemoryFilterCpuMemoryUsage(o["cpu-memory-usage"], d, "cpu_memory_usage")); err != nil {
		if !fortiAPIPatch(o["cpu-memory-usage"]) {
			return fmt.Errorf("Error reading cpu_memory_usage: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogMemoryFilterFilter(o["filter"], d, "filter")); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogMemoryFilterFilterType(o["filter-type"], d, "filter_type")); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	return nil
}

func flattenLogMemoryFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogMemoryFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterNetscanDiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterNetscanVulnerability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterRadius(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterIpsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterDhcp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterPpp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSslvpnLogAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSslvpnLogAdm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSslvpnLogSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterVipSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterLdbMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterWanOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterWirelessActivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterCpuMemoryUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemoryFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandLogMemoryFilterSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok {
		t, err := expandLogMemoryFilterForwardTraffic(d, v, "forward_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-traffic"] = t
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok {
		t, err := expandLogMemoryFilterLocalTraffic(d, v, "local_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok {
		t, err := expandLogMemoryFilterMulticastTraffic(d, v, "multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok {
		t, err := expandLogMemoryFilterSnifferTraffic(d, v, "sniffer_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sniffer-traffic"] = t
		}
	}

	if v, ok := d.GetOk("anomaly"); ok {
		t, err := expandLogMemoryFilterAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("netscan_discovery"); ok {
		t, err := expandLogMemoryFilterNetscanDiscovery(d, v, "netscan_discovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netscan-discovery"] = t
		}
	}

	if v, ok := d.GetOk("netscan_vulnerability"); ok {
		t, err := expandLogMemoryFilterNetscanVulnerability(d, v, "netscan_vulnerability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netscan-vulnerability"] = t
		}
	}

	if v, ok := d.GetOk("voip"); ok {
		t, err := expandLogMemoryFilterVoip(d, v, "voip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip"] = t
		}
	}

	if v, ok := d.GetOk("gtp"); ok {
		t, err := expandLogMemoryFilterGtp(d, v, "gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp"] = t
		}
	}

	if v, ok := d.GetOk("dns"); ok {
		t, err := expandLogMemoryFilterDns(d, v, "dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		t, err := expandLogMemoryFilterSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("event"); ok {
		t, err := expandLogMemoryFilterEvent(d, v, "event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok {
		t, err := expandLogMemoryFilterSystem(d, v, "system")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("radius"); ok {
		t, err := expandLogMemoryFilterRadius(d, v, "radius")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius"] = t
		}
	}

	if v, ok := d.GetOk("ipsec"); ok {
		t, err := expandLogMemoryFilterIpsec(d, v, "ipsec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec"] = t
		}
	}

	if v, ok := d.GetOk("dhcp"); ok {
		t, err := expandLogMemoryFilterDhcp(d, v, "dhcp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp"] = t
		}
	}

	if v, ok := d.GetOk("ppp"); ok {
		t, err := expandLogMemoryFilterPpp(d, v, "ppp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp"] = t
		}
	}

	if v, ok := d.GetOk("admin"); ok {
		t, err := expandLogMemoryFilterAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok {
		t, err := expandLogMemoryFilterHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		t, err := expandLogMemoryFilterAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok {
		t, err := expandLogMemoryFilterPattern(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_log_auth"); ok {
		t, err := expandLogMemoryFilterSslvpnLogAuth(d, v, "sslvpn_log_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-log-auth"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_log_adm"); ok {
		t, err := expandLogMemoryFilterSslvpnLogAdm(d, v, "sslvpn_log_adm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-log-adm"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_log_session"); ok {
		t, err := expandLogMemoryFilterSslvpnLogSession(d, v, "sslvpn_log_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-log-session"] = t
		}
	}

	if v, ok := d.GetOk("vip_ssl"); ok {
		t, err := expandLogMemoryFilterVipSsl(d, v, "vip_ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip-ssl"] = t
		}
	}

	if v, ok := d.GetOk("ldb_monitor"); ok {
		t, err := expandLogMemoryFilterLdbMonitor(d, v, "ldb_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-monitor"] = t
		}
	}

	if v, ok := d.GetOk("wan_opt"); ok {
		t, err := expandLogMemoryFilterWanOpt(d, v, "wan_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-opt"] = t
		}
	}

	if v, ok := d.GetOk("wireless_activity"); ok {
		t, err := expandLogMemoryFilterWirelessActivity(d, v, "wireless_activity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-activity"] = t
		}
	}

	if v, ok := d.GetOk("cpu_memory_usage"); ok {
		t, err := expandLogMemoryFilterCpuMemoryUsage(d, v, "cpu_memory_usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-memory-usage"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		t, err := expandLogMemoryFilterFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok {
		t, err := expandLogMemoryFilterFilterType(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	return &obj, nil
}
