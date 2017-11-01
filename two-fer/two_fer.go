// Package twofer contains an implementation of the same "SharedWith function
package twofer

import "fmt"

// ShareWith returns a 2-fer statement.
// If
func ShareWith(name ...string) string {
	var twofer string
	if name[0] == "" {
		twofer = "One for you, one for me."
	} else {
		twofer = fmt.Sprintf("One for %s, one for me.", name[0])
	}
	return twofer
}
