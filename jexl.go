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

type Engine struct {
	engine *C.JexlEngine
}

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

func (e Engine) Run() string {
	c_res := C.run_engine(e.engine)
	return C.GoString(c_res)
}

func (e Engine) Free() {
	C.free_engine(e.engine)
}
