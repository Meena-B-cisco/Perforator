GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v1.17.3)

SRCS(
    bson.go
    decoder.go
    doc.go
    encoder.go
    marshal.go
    primitive_codecs.go
    raw.go
    raw_element.go
    raw_value.go
    registry.go
    types.go
    unmarshal.go
)

GO_TEST_SRCS(
    benchmark_test.go
    bson_corpus_spec_test.go
    bson_test.go
    decoder_test.go
    encoder_test.go
    extjson_prose_test.go
    fuzz_test.go
    marshal_test.go
    marshal_value_cases_test.go
    marshal_value_test.go
    marshaling_cases_test.go
    primitive_codecs_test.go
    raw_test.go
    raw_value_test.go
    truncation_test.go
    unmarshal_test.go
    unmarshal_value_test.go
    unmarshaling_cases_test.go
)

GO_XTEST_SRCS(
    decoder_example_test.go
    encoder_example_test.go
    example_test.go
)

END()

RECURSE(
    bsoncodec
    bsonoptions
    bsonrw
    bsontype
    gotest
    mgocompat
    primitive
)
