package main

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -ljexl -static -lm
#cgo CFLAGS: -I${SRCDIR}/lib
#include <jexl.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.eval(C.CString("6 * 12 + 5 / 2.6"))))
}
