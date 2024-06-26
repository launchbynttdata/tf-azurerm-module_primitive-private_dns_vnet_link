package testimpl

import (
	"regexp"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestVnetLink(t *testing.T, ctx types.TestContext) {

	t.Run("TestVnetLink", func(t *testing.T) {
		linkId := terraform.Output(t, ctx.TerratestTerraformOptions(), "vnet_link_id")
		dnsZoneId := terraform.Output(t, ctx.TerratestTerraformOptions(), "dns_zone_id")
		resourceGroupId := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_id")
		vnetId := terraform.Output(t, ctx.TerratestTerraformOptions(), "vnet_id")
		dnsZoneName := terraform.Output(t, ctx.TerratestTerraformOptions(), "dns_zone_name")

		assert.NotEmpty(t, linkId, "ID must not be empty")
		assert.NotEmpty(t, dnsZoneId, "DNS Zone ID must not be empty")
		assert.NotEmpty(t, resourceGroupId, "Resource Group ID must not be empty")
		assert.NotEmpty(t, vnetId, "Vnet ID must not be empty")
		assert.Regexp(t, regexp.MustCompile(`^[A-Za-z0-9\.\-]+$`), dnsZoneName)
	})
}
