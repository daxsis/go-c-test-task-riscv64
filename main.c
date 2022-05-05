#include <sys/types.h>
#include <sys/sysctl.h>
#include <stdio.h>

int main() {

        int mib[2];
        size_t len, memory;

        mib[0] = CTL_HW;
        mib[1] = HW_REALMEM;

        sysctl(mib, 2, &memory, &len, NULL, 0);
        printf("%lu", memory);

        return 0;
}

/*
 * root@freebsd:~/projects # uname -a
 * FreeBSD freebsd 14.0-CURRENT FreeBSD 14.0-CURRENT #0 main-n254961-b91a48693a5: Thu Apr 21 05:45:02 UTC 2022     root@releng1.nyi.freebsd.org:/usr/obj/usr/src/riscv.riscv64/sys/GENERIC riscv
*/
