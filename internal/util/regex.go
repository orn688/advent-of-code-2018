package util

import (
	"fmt"
	"regexp"
)

// CaptureRegexGroups returns a mapping from regex named group names to values.
func CaptureRegexGroups(regex *regexp.Regexp, str string) (map[string]string, error) {
	names := regex.SubexpNames()
	groups := make(map[string]string, len(names))
	match := regex.FindStringSubmatch(str)

	for i := range names {
		groupName := names[i]
		if i >= len(match) {
			return groups, fmt.Errorf("no match for group %s", groupName)
		}
		groups[groupName] = match[i]
	}
	return groups, nil
}
