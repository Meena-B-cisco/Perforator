GO_LIBRARY()

LICENSE(BSD-3-Clause)

VERSION(v0.0.0-20240613232115-7f521ea00fb8)

SRCS(
    maps.go
)

GO_TEST_SRCS(maps_test.go)

END()

RECURSE(
    gotest
)
