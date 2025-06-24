GO_LIBRARY()

LICENSE(MIT)

VERSION(v5.7.2)

SRCS(
    pgmock.go
)

GO_XTEST_SRCS(pgmock_test.go)

END()

RECURSE(
    gotest
)
