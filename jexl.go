/*
jexl is a library to evaluate the jexl scripting language in Go.
This implementation relies on the Rust library `jexl-rs`.
*/
package jexl

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -ljexl -lm
#cgo CFLAGS: -I${SRCDIR}/lib
#include <jexl.h>
*/
import "C"

// Eval evaluates an expression, whether that be an entire script or something smaller.
// This function relies on FFI to function. The expression is parsed into a C String then handed off to Rust.
// The Rust code calls to the jexl-rs library to actually evaluate the expression then hand it back to Go.
func Eval(expression string) (string, error) {
	cExpr := C.CString(expression)
	resCString := C.eval(cExpr)
	return C.GoString(resCString), nil
}
