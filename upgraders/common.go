package upgraders

import "strings"

// UpgradeMapsToLists is the single entry point for SDK v1→v2 state upgrades.
// For each named field in rawState it recursively converts map values (TypeMap in
// SDK v1) to []interface{}{map} (TypeList MaxItems:1 in SDK v2), handling:
//   - flat TypeMaps: the map is simply wrapped in a one-element slice
//   - nested TypeMaps: SDK v1 stored them as dot-joined keys such as
//     "office_mode.mode"; upgradeMap unflattens those back to nested objects
//   - TypeList fields: the slice is walked element-by-element so that TypeMap
//     sub-fields inside each element are also upgraded
func UpgradeMapsToLists(rawState map[string]interface{}, fields ...string) map[string]interface{} {
	for _, f := range fields {
		if v, ok := rawState[f]; ok {
			rawState[f] = upgradeValue(v)
		}
	}
	return rawState
}

// upgradeValue recurses over a single value:
//
//	map   → wrap in []interface{} (was TypeMap) and call upgradeMap on the inner map
//	slice → walk elements, calling upgradeMap on any map-typed elements (TypeList stays a list)
//	other → return unchanged
func upgradeValue(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		return []interface{}{upgradeMap(val)}
	case []interface{}:
		for i, elem := range val {
			if m, ok := elem.(map[string]interface{}); ok {
				val[i] = upgradeMap(m)
			}
		}
		return val
	default:
		return v
	}
}

// upgradeMap processes the contents of one TypeMap:
//   - keys that contain a dot are grouped by their first segment to unflatten the
//     SDK v1 nested-TypeMap dot-encoding; each group is recursively upgraded
//   - plain keys have upgradeValue called on their value (handles nested maps that
//     were already decoded as map[string]interface{} rather than dot-encoded)
func upgradeMap(flat map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	groups := make(map[string]map[string]interface{})
	for k, v := range flat {
		if idx := strings.Index(k, "."); idx >= 0 {
			prefix, rest := k[:idx], k[idx+1:]
			if groups[prefix] == nil {
				groups[prefix] = make(map[string]interface{})
			}
			groups[prefix][rest] = v
		} else {
			result[k] = upgradeValue(v)
		}
	}
	for prefix, subMap := range groups {
		result[prefix] = []interface{}{upgradeMap(subMap)}
	}
	return result
}
