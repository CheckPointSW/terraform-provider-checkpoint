package checkpoint

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider

// Legacy acceptance tests still reference TestCase.Providers.
// Prefer testAccProviderFactories for new tests.
var testAccProviders map[string]*schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"checkpoint": testAccProvider,
	}
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"checkpoint": func() (*schema.Provider, error) { return Provider(), nil },
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
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
