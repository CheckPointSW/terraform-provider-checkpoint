package checkpoint

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// gaiaResourceImporter restores the lookup fields used by Read. Gaia state IDs
// created by the provider are synthetic, so a passthrough importer is insufficient.
// A single required lookup field can be supplied directly; composite identifiers
// and optional context fields use a JSON object with Terraform attribute names.
func gaiaResourceImporter(fields map[string]schema.ValueType, required ...string) *schema.ResourceImporter {
	return &schema.ResourceImporter{
		State: func(d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
			id := strings.TrimSpace(d.Id())
			if id == "" {
				return nil, fmt.Errorf("Gaia import ID must not be empty")
			}
			values := make(map[string]json.RawMessage)
			if strings.HasPrefix(id, "{") {
				if err := json.Unmarshal([]byte(id), &values); err != nil {
					return nil, fmt.Errorf("Gaia import ID must be a valid JSON object: %w", err)
				}
			} else if len(required) == 1 {
				if fields[required[0]] == schema.TypeString {
					values[required[0]], _ = json.Marshal(d.Id())
				} else {
					values[required[0]] = json.RawMessage(id)
				}
			} else if len(required) != 0 || id != "settings" {
				return nil, fmt.Errorf("use a JSON object with the resource lookup fields, or settings for a resource with no required lookup fields")
			}

			for _, name := range required {
				if _, ok := values[name]; !ok {
					return nil, fmt.Errorf("Gaia import ID is missing required lookup field %q", name)
				}
			}
			// Validate everything before changing state. Never accept configuration
			// attributes (for example passwords) as part of an import identifier.
			decoded := make(map[string]interface{}, len(values))
			for name, raw := range values {
				fieldType, ok := fields[name]
				if !ok {
					return nil, fmt.Errorf("unknown Gaia import lookup field %q", name)
				}
				if string(raw) == "null" {
					return nil, fmt.Errorf("Gaia import lookup field %q must not be null", name)
				}
				var value interface{}
				var err error
				switch fieldType {
				case schema.TypeString:
					var v string
					err = json.Unmarshal(raw, &v)
					if err == nil && strings.TrimSpace(v) == "" {
						err = fmt.Errorf("must not be empty")
					}
					value = v
				case schema.TypeInt:
					var v int
					err = json.Unmarshal(raw, &v)
					value = v
				case schema.TypeBool:
					var v bool
					err = json.Unmarshal(raw, &v)
					value = v
				default:
					return nil, fmt.Errorf("unsupported Gaia import lookup type for field %q", name)
				}
				if err != nil {
					return nil, fmt.Errorf("invalid Gaia import lookup field %q: %w", name, err)
				}
				decoded[name] = value
			}
			for name, value := range decoded {
				if err := d.Set(name, value); err != nil {
					return nil, fmt.Errorf("cannot set Gaia import lookup field %q: %w", name, err)
				}
			}
			return []*schema.ResourceData{d}, nil
		},
	}
}
