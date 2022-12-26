//go:build tools
// +build tools

package tools

import (
	_ "github.com/UnnoTed/fileb0x"
	_ "golang.org/x/net/webdav" // Used inside generated code.
)
