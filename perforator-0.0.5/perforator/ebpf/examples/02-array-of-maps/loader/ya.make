GO_LIBRARY()

PEERDIR(
    perforator/ebpf/examples/02-array-of-maps/prog
    vendor/github.com/cilium/ebpf
    ${GOSTD}/errors
)

RUN_PROGRAM(
    perforator/ebpf/tools/btf2go
    -elf
    perforator/ebpf/examples/02-array-of-maps/prog/prog.debug.elf
    -package
    loader
    -output
    prog.go
    IN
    perforator/ebpf/examples/02-array-of-maps/prog/prog.debug.elf
    OUT
    prog.go
)

RESOURCE(
    perforator/ebpf/examples/02-array-of-maps/prog/prog.release.elf ebpf/prog.release.elf
    perforator/ebpf/examples/02-array-of-maps/prog/prog.debug.elf ebpf/prog.debug.elf
)

SRCS(
    loader.go
)

END()
