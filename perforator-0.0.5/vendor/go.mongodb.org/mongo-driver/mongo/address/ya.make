GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v1.17.3)

SRCS(
    addr.go
)

GO_TEST_SRCS(addr_test.go)

END()

RECURSE(
    gotest
)
