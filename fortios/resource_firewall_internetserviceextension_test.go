
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallInternetServiceExtension_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallInternetServiceExtension_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallInternetServiceExtensionConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallInternetServiceExtensionExists("fortios_firewall_internetserviceextension.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_internetserviceextension.trname", "comment", "EIWE"),
                    resource.TestCheckResourceAttr("fortios_firewall_internetserviceextension.trname", "fosid", "65536"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallInternetServiceExtensionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallInternetServiceExtension: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallInternetServiceExtension is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallInternetServiceExtension(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallInternetServiceExtension: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallInternetServiceExtension: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallInternetServiceExtensionDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_internetserviceextension" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallInternetServiceExtension(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallInternetServiceExtension %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallInternetServiceExtensionConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_internetserviceextension" "trname" {
  comment = "EIWE"
  fosid   = 65536
}
















`)
}