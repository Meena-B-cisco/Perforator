GO_LIBRARY()

LICENSE(MIT)

VERSION(v5.7.2)

SRCS(
    sanitize.go
)

GO_XTEST_SRCS(sanitize_test.go)

END()

RECURSE(
    gotest
)
