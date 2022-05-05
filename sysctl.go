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

/*
root@freebsd:~/projects # uname -a
FreeBSD freebsd 13.0-RELEASE FreeBSD 13.0-RELEASE #0 releng/13.0-n244733-ea31abc261f: Fri Apr  9 04:24:09 UTC 2021     root@releng1.nyi.freebsd.org:/usr/obj/usr/src/amd64.amd64/sys/GENERIC  amd64
*/
