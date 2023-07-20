package core

import (
	"regexp"
)

// replaces multiple spaces with a single space
// adds space after a full stop if it's missing

func CleanupDescription(description string) string {
	// Replace multiple spaces with a single one
	re := regexp.MustCompile(`\s+`)
	description = re.ReplaceAllString(description, " ")

	// Add space after punctuation if it's missing
	re = regexp.MustCompile(`([.,!?])(\S)`)
	description = re.ReplaceAllString(description, "$1 $2")

	return description
}
