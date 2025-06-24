GO_LIBRARY()

SRCS(
    utils.go
)

IF (OPENSOURCE)
    SRCS(
        systemcert.go
    )
ELSE()
    SRCS(
        certifi_yandex.go
    )
ENDIF()

END()
