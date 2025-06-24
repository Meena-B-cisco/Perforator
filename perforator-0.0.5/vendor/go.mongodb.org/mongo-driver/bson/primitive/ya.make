GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v1.17.3)

SRCS(
    decimal.go
    objectid.go
    primitive.go
)

GO_TEST_SRCS(
    decimal_test.go
    objectid_test.go
    primitive_test.go
)

END()

RECURSE(
    gotest
)
