package functions_test

import (
	"fmt"
	"github.com/tae2089/terraform-provider-crypto/internal/util"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/tae2089/terraform-provider-crypto/internal/provider"
)

func TestExampleFunction_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDecryptFunctionConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "hello, world!"),
				),
			},
		},
	})
}

func testDecryptFunctionConfig() string {
	projectDir := util.GetProjectDir("../..")
	projectDir = projectDir + "/examples/functions/decrypt/private.pem"
	return fmt.Sprintf(`
	output "test" {
		value = provider::crypto::decrypt("%s", "a6f2a0df08bd86f5a49d0a079b7799158674f7eebb56e4a9f2636279e91455dffcc8672d944e37f9697f210d4e02c61f3812198ee8fda794569efad86ef968371fb285e70321ef45c61db64d329fc262e7b29dea22279ea760e4f7a2c8ab1e83a20a1418e8a354aeb7e3994c6d31656945d5afc79ff574acbc831414e78ce5236a1ee5ee7601e118ee995ca96063238bbaf622d56530f5ce1b3980f65bba3d21c76ff22aff473d32db7f94b939ceec09a2573d43a78a8be3f4e09b9c91b91d32c7be8ee6ae31b8c0ee39b128bb4b697b68abb2ff555b99db698bb39e84ad2b9c3d88a3a0e1398a7202154fd153419f98e2b8eea45782845565479a8adc0a9a94")
	}`, projectDir,
	)
}

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"crypto": providerserver.NewProtocol6WithError(provider.New()()),
}
