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
