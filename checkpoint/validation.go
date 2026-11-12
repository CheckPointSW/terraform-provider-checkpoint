package checkpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func objectNotFound(code string) bool {
	notFoundCode := "generic_err_object_not_found"
	return code == notFoundCode
}

func validateStringValue(optionalValues ...string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (warns []string, errs []error) {
		value := v.(string)
		ok := false
		for _, optionalValue := range optionalValues {
			if value == optionalValue {
				ok = true
				break
			}
		}
		if !ok {
			errs = append(errs, fmt.Errorf("'%q' %q is not a valid value. Optional values are: %v", value, k, optionalValues))
		}
		return
	}
}
