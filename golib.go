//
// test module path with different package name
// go-lib vs golib
//

package golib

import "fmt"

const _moduleVersion = "v0.0.1"

const _packageVersion = "v0.0.0"

// Version test
func Version() string {
	return _moduleVersion
}

// FuncTest test
func FuncTest() string {
	return fmt.Sprintf("golib/FuncTest() Module Version %s, _packageVersion %s ", _moduleVersion, _packageVersion)
}
