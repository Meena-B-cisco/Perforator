GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v0.15.1)

SRCS(
    stats.go
)

GO_TEST_SRCS(stats_test.go)

END()

RECURSE(
    # gotest
)
