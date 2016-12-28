// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/omarvides/goalearn/design
// --out=$(GOPATH)/src/github.com/omarvides/goalearn
// --version=v1.1.0-dirty
//
// API "user": Application Resource Href Factories
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"fmt"
	"strings"
)

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/users/%v", paramuserID)
}
