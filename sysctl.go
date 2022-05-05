package main

// #include <sys/types.h>
// #include <sys/sysctl.h>
// #include <stdlib.h>
import "C"
import (
        "fmt"
        "unsafe"
)

func main() {
        var mem C.size_t
        var len C.size_t = C.size_t(4)
        var mib[2] C.int
        cs := C.CString("hw.realmem")
        defer C.free(unsafe.Pointer(cs))
        C.sysctlnametomib(cs, (*C.int)(&mib[0]), (*C.ulong)(&len)) // get integer representation

        var length C.ulong = 16 // make sure it fits
        C.sysctl((*C.int)(&mib[0]),
                2,
                unsafe.Pointer(&mem),
                &length,
                nil,
                0)
        fmt.Print(mem)
}
