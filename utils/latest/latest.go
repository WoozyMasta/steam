package latest

import (
	"fmt"
	"strconv"
	"strings"
)

/*
FindVersion returns the newest version that satisfies the specified threshold.

The threshold is a value in the range 0..99 (representing a percentage).
If threshold = 0, this function simply returns the highest version
(using CompareVersions). In that case, using threshold logic is unnecessary
because a simple comparison would suffice.

FindVersion internally calls FindVersionSafe, using the same threshold
and calculating fallback as 60% of the threshold.
*/
func FindVersion(versionMap map[string]uint32, threshold float64) (string, error) {
	// We pass threshold and threshold*0.6 as fallback to FindVersionSafe
	return FindVersionSafe(versionMap, threshold, threshold*0.6)
}

/*
FindVersionSafe returns the newest version that meets the threshold.
If no version meets the main threshold, it attempts the fallback threshold.
If it also fails to find a version at the fallback threshold, it returns
the highest version from the map.

The threshold and fallback are values in the range 0..99 (percentages).
The fallback cannot exceed 90% of the threshold. If fallback * 1.1 is greater
than threshold, the function returns an error.

The logic is as follows:
 1. Filter versions that meet the threshold (as a percentage of total count).
 2. If no versions pass the threshold, filter by fallback.
 3. If still no versions pass fallback, return the maximum version from the map.
*/
func FindVersionSafe(versionMap map[string]uint32, threshold, fallback float64) (string, error) {
	if len(versionMap) == 0 {
		return "", fmt.Errorf("empty version map")
	}
	if threshold > 99 {
		return "", fmt.Errorf("threshold %.2f%% cant be greater than 99%%", threshold)
	}
	// Make sure fallback is significantly lower than threshold (at least 10% lower).
	if fallback*1.1 > threshold {
		return "", fmt.Errorf(
			"fallback threshold %.2f%% so close to threshold %.2f%%, must be lower than %.2f%%",
			fallback, threshold, threshold*0.9,
		)
	}

	// Calculate the total count of all versions
	var totalCount uint32
	for _, c := range versionMap {
		totalCount += c
	}

	// First, try to find candidates that meet the threshold
	candidates := filterByThreshold(versionMap, float64(totalCount), threshold)
	if len(candidates) == 0 {
		// If no candidates, try the fallback threshold
		candidates = filterByThreshold(versionMap, float64(totalCount), fallback)
	}

	// If no candidates even after fallback, return the max version in the map
	if len(candidates) == 0 {
		return findMaxVersionKey(versionMap), nil
	}

	// Among all candidates, return the maximum version
	return findMaxVersionKey(candidates), nil
}

/*
CompareVersions compares two arbitrary version strings.

Returns:

	-1 if v1 < v2
	 0 if v1 == v2
	 1 if v1 > v2

Comparison logic:
 1. Both version strings are split by "."
 2. Corresponding segments are compared:
    a) If both segments are numeric, compare as numbers
    b) Otherwise, compare as strings in lexicographical order
 3. If one version has fewer segments, the missing segments are treated as empty
*/
func CompareVersions(v1, v2 string) int8 {
	v1parts := strings.Split(v1, ".")
	v2parts := strings.Split(v2, ".")

	maxLen := len(v1parts)
	if len(v2parts) > maxLen {
		maxLen = len(v2parts)
	}

	for i := 0; i < maxLen; i++ {
		var p1, p2 string
		if i < len(v1parts) {
			p1 = v1parts[i]
		}
		if i < len(v2parts) {
			p2 = v2parts[i]
		}

		// Attempt numeric comparison
		n1, err1 := strconv.Atoi(p1)
		n2, err2 := strconv.Atoi(p2)

		switch {
		case err1 == nil && err2 == nil:
			// Both segments are numeric
			if n1 < n2 {
				return -1
			} else if n1 > n2 {
				return 1
			}
		case err1 == nil && err2 != nil:
			// v1's segment is numeric, v2's is string-based
			return 1
		case err1 != nil && err2 == nil:
			// v1's segment is string-based, v2's is numeric
			return -1
		default:
			// Both segments are strings
			if p1 < p2 {
				return -1
			} else if p1 > p2 {
				return 1
			}
		}
	}

	// If all parts are equal or empty, return 0
	return 0
}

// findMaxVersionKey returns the key with the "highest" version
// according to CompareVersions.
func findMaxVersionKey(versionMap map[string]uint32) string {
	var maxVer string
	for v := range versionMap {
		if maxVer == "" {
			maxVer = v
			continue
		}

		if CompareVersions(v, maxVer) > 0 {
			maxVer = v
		}
	}

	return maxVer
}

// filterByThreshold filters the map of versions by percentage.
// It returns a new map containing only those versions whose count
// meets or exceeds the given threshold (percent of total).
func filterByThreshold(versionMap map[string]uint32, total float64, threshold float64) map[string]uint32 {
	result := make(map[string]uint32)
	// If threshold <= 0, return everything
	if threshold <= 0 {
		return versionMap
	}

	for v, c := range versionMap {
		weight := (float64(c) / total) * 100
		if weight >= threshold {
			result[v] = c
		}
	}

	return result
}
