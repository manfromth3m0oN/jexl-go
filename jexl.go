package jexl

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -ljexl -lm
#cgo CFLAGS: -I${SRCDIR}/lib
#include <jexl.h>
*/
import "C"

func Eval(expression string) (string, error) {
	cExpr := C.CString(expression)
	resCString := C.eval(cExpr)
	return C.GoString(resCString), nil
}
