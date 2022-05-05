# go-c-test-task-riscv64

#Installation on Arch Linux:
- FreeBSD/riscv64 image
`$ wget https://download.freebsd.org/ftp/snapshots/VM-IMAGES/14.0-CURRENT/riscv64/Latest/FreeBSD-14.0-CURRENT-riscv-riscv64.raw.xz`
- decompress:
`$ xz --decompress FreeBSD-14.0-CURRENT-riscv-riscv64.raw.xz`
- extende the image ( filesystem will pickup the changes on the first boot):
`$ truncate -s 10G FreeBSD-14.0-CURRENT-riscv-riscv64.raw`
- install cross-compilatiol tools:
`$ sudo pacman -S riscv64-linux-gnu-binutils riscv64-linux-gnu-gcc riscv64-linux-gnu-gdb riscv64-linux-gnu-glibc`
- download u-boot
`
$ git clone https://github.com/u-boot/u-boot.git
$ cd u-boot
$ git checkout v2022.04
`
- build u-boot:
`
$ export ARCH=riscv
$ export CROSS_COMPILE=riscv64-linux-gnu-
$ make qemu-riscv64_smode_defconfig
$ make
`

- docs: https://github.com/riscv-software-src/opensbi/blob/master/docs/platform/qemu_virt.md
- dowload opnsbi
`$ git clone https://github.com/riscv/opensbi.git`
`$ make PLATFORM=generic FW_PAYLOAD_PATH=../u-boot/u-boot.bin`
- start the emulator:
`$ qemu-system-riscv64 -machine virt -m 2048M -smp 2 -nographic \
    -bios /usr/local/share/opensbi/lp64/generic/firmware/fw_jump.elf \
    -kernel /usr/local/share/u-boot/u-boot-qemu-riscv64/u-boot.bin \
    -drive file=FreeBSD-14.0-CURRENT-riscv-riscv64.raw,format=raw,id=hd0 -device virtio-blk-device,drive=hd0 \
    -netdev user,id=net0,ipv6=off,hostfwd=tcp::8022-:22 -device virtio-net-device,netdev=net0`
- to exit from QEMU (ctrl+a x)
- docker container with Debian/riscv on QEMU

# Additional Information
- U-boot guide:
https://github.com/u-boot/u-boot/blob/master/doc/board/emulation/qemu-riscv.rst

- Debug:
https://github.com/cyrus-and/gdb-dashboard

- Cross Debug FreeBSD kernel on Linux with QEMU
https://github.com/chenshuo/notes/wiki/Cross-debug-FreeBSD-kernel-on-Linux-host-with-QEMU

- Debian RISC-V
https://wiki.debian.org/RISC-V

- QEMU docs
https://wiki.qemu.org/Documentation/Platforms/RISCV

# Start the image
`
qemu-system-riscv64 -machine virt \
-m 4096M \
-smp 2 \
-nographic \
-bios ./opensbi/build/platform/generic/firmware/fw_jump.elf \
-kernel ./u-boot/u-boot.bin \
-drive file=FreeBSD-14.0-CURRENT-riscv-riscv64.raw,format=raw,id=hd0 \
-device virtio-blk-device,drive=hd0 \
-netdev user,id=net0,ipv6=off,hostfwd=tcp::8022-:22 \
-device virtio-net-device,netdev=net0
`
