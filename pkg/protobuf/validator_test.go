package protobuf_test

import (
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syncromatics/proto-schema-registry/pkg/protobuf"
)

func Test_Validator_ShouldFailForInvalidProto(t *testing.T) {
	schema := "poo"

	var b bytes.Buffer
	err := gzipWrite(&b, []byte(schema))
	if err != nil {
		t.Fatal(err)
	}

	err = protobuf.Validate(b.Bytes())

	assert.Equal(t,
		"invalid proto: <input>:1:1: found \"poo\" but expected [.proto element {comment|option|import|syntax|enum|service|package|message}]",
		err.Error())
}

func Test_Validator_ShouldFailForInvalidGzip(t *testing.T) {
	b := []byte{0x0, 0x1, 0x2}

	err := protobuf.Validate(b)

	assert.Equal(t,
		"invalid gzip: failed to create gzip reader: unexpected EOF",
		err.Error())
}

func Test_Validator_ShouldFailForNoRecordMessage(t *testing.T) {
	schema := `
message stuff {
	string h = 1;
}`

	var b bytes.Buffer
	err := gzipWrite(&b, []byte(schema))
	if err != nil {
		t.Fatal(err)
	}

	err = protobuf.Validate(b.Bytes())

	assert.Equal(t,
		"proto has no record message",
		err.Error())
}

func Test_Validator_ShouldPassWithValidProto(t *testing.T) {
	schema := `
message record {
	string h = 1;
}

message other {
	int32 i = 1;
}`

	var b bytes.Buffer
	err := gzipWrite(&b, []byte(schema))
	if err != nil {
		t.Fatal(err)
	}

	err = protobuf.Validate(b.Bytes())

	assert.Nil(t, err)
}

func gzipWrite(w io.Writer, data []byte) error {
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer gw.Close()
	gw.Write(data)
	return err
}
