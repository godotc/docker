# Create a image for docker run

## 1. Copy bin file and dynamic so
- Way to make home-made image:
    1. find the ldd of /bin/*
    2. cp bin to rootfs/bin/
    3. cp *.so to rootfs/path/that/corresponding/
    4. mkdir rootfs/proc (ps and mount)
     

- In root:
```shell
$ ldd /bin/bash
    linux-vdso.so.1 =>               linux-vdso.so.1 (0x00007fff38fca000)
    libreadline.so.8 =>              /usr/lib/libreadline.so.8 x00007f8b9b2ca000)
    libdl.so.2 =>                    /usr/lib/libdl.so.2 x00007f8b9b2c5000)
    libc.so.6 =>                     /usr/lib/libc.so.6 x00007f8b9b0de000)
    libncursesw.so.6 =>              /usr/lib/libncursesw.so.6 x00007f8b9b06a000)
    /lib64/ld-linux-x86-64.so.2 =>   /usr/lib64/ld-linux-x86-64.so.2 x00007f8b9b426000)
```

```shell
$ ldd /bin/mount
linux-vdso.so.1 => linux-vdso.so.1 (0x00007ffd749f8000)
    libmount.so.1 => /usr/lib/libmount.so.1 (0x00007fb9357e8000)
    libc.so.6 => /usr/lib/libc.so.6 (0x00007fb935601000)
    libblkid.so.1 => /usr/lib/libblkid.so.1 (0x00007fb9355c8000)
    /lib64/ld-linux-x86-64.so.2 => /usr/lib64/ld-linux-x86-64.so.2 (0x00007fb935848000)
```

- Need to cp to rootfs/lib/**:
    /usr/lib/libreadline.so.8  \
    /usr/lib/libdl.so.2 \
    /usr/lib/libc.so.6 \
    /usr/lib/libncursesw.so.6 \
    /usr/lib64/ld-linux-x86-64.so.2 \
    ...


## 2. Export from docek

1. From outer docker, docker pull archlinux
2. docker export -o archlinux.tar [imageId]
3. tar -xvf archlinux.tar -C base

> Need the corresponding architecture of container  
