package checkpoint

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"chkp": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("CHKP_SERVER") == "" {
		t.Fatal("CHKP_SERVER must be set for acceptance tests")
	}
	if os.Getenv("CHKP_USERNAME") == "" {
		t.Fatal("CHKP_USERNAME must be set for acceptance tests")
	}
	if os.Getenv("CHKP_PASSWORD") == "" {
		t.Fatal("CHKP_PASSWORD must be set for acceptance tests")
	}
}
