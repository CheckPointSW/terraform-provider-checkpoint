package checkpoint

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// FMW-12074: mask_length = 0 must not be treated as "unset" for this Required field.
func TestGaiaStaticRouteMaskLengthZero(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceGaiaStaticRoute().Schema, map[string]interface{}{
		"address":     "0.0.0.0",
		"mask_length": 0,
		"type":        "blackhole",
	})

	if v, ok := d.GetOk("mask_length"); ok {
		t.Fatalf("GetOk(\"mask_length\") = (%v, %v); expected ok=false for value 0 - this is the bug GetOk cannot distinguish an explicit 0 from an unset Required int field", v, ok)
	}

	v, ok := d.GetOkExists("mask_length")
	if !ok {
		t.Fatalf("GetOkExists(\"mask_length\") = (%v, %v); want ok=true for explicit mask_length=0", v, ok)
	}
	if v.(int) != 0 {
		t.Fatalf("GetOkExists(\"mask_length\") value = %v; want 0", v)
	}
}
