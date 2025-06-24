GO_LIBRARY()

TAG(ya:run_go_benchmark)
IF (NOT OPENSOURCE)
    DATA(
        sbr://3766900569=stacks
    )
ENDIF()

SRCS(
    blocks.go
    hsv.go
    render.go
    strtab.go
)
IF (NOT OPENSOURCE)
    GO_TEST_SRCS(
        blocks_test.go
        render_json_test.go
    )
    GO_XTEST_SRCS(render_test.go)
ENDIF()

GO_EMBED_PATTERN(tmpl.html)

END()

RECURSE(
    cmd
    gotest
    format
)
