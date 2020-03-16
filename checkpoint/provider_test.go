package checkpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"checkpoint": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("CHECKPOINT_SERVER") == "" {
		t.Fatal("CHECKPOINT_SERVER must be set for acceptance tests")
	}
	if os.Getenv("CHECKPOINT_USERNAME") == "" {
		t.Fatal("CHECKPOINT_USERNAME must be set for acceptance tests")
	}
	if os.Getenv("CHECKPOINT_PASSWORD") == "" {
		t.Fatal("CHECKPOINT_PASSWORD must be set for acceptance tests")
	}
	if os.Getenv("CHECKPOINT_CONTEXT") == "" {
		t.Fatal("CHECKPOINT_CONTEXT must be set for acceptance tests")
	}
}
