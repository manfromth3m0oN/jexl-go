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
import "encoding/json"

// Eval evaluates an expression, whether that be an entire script or something smaller.
// This function relies on FFI to function. The expression is parsed into a C String then handed off to Rust.
// The Rust code calls to the jexl-rs library to actually evaluate the expression then hand it back to Go.
func Eval(expression string) (string, error) {
	cExpr := C.CString(expression)
	resCString := C.eval(cExpr)
	return C.GoString(resCString), nil
}

// Engine holds a reference to the Rust jexl engine object to allow for long running FFI
type Engine struct {
	engine *C.JexlEngine
}

// NewEngine creates a new engine with a context and script.
// This means you should use one engine per script.
// Don't forget to call free on your engine, you'll leak memory otherwise.
func NewEngine(context interface{}, script string) (Engine, error) {
	b, err := json.Marshal(context)
	if err != nil {
		return Engine{}, err
	}
	engine := C.new_engine(C.CString(string(b)), C.CString(script))
	return Engine{
		engine,
	}, nil
}

// Run will run the script provided at Initialization in the context also provided.
func (e Engine) Run() string {
	c_res := C.run_engine(e.engine)
	return C.GoString(c_res)
}

// Free instructs the Rust code to drop the memory associated with this frontend
func (e Engine) Free() {
	C.free_engine(e.engine)
}
