GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v1.17.3)

SRCS(
    tag.go
)

GO_TEST_SRCS(tag_test.go)

END()

RECURSE(
    gotest
)
