GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v0.15.1)

SRCS(
    nfs.go
    parse.go
    parse_nfs.go
    parse_nfsd.go
)

GO_XTEST_SRCS(
    parse_nfs_test.go
    parse_nfsd_test.go
)

END()

RECURSE(
    # gotest
)
