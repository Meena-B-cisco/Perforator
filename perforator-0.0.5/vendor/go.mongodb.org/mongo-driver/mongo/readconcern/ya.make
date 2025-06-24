GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v1.17.3)

SRCS(
    readconcern.go
)

GO_XTEST_SRCS(readconcern_test.go)

END()

RECURSE(
    gotest
)
