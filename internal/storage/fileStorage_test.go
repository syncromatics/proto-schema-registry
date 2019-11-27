package storage_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"path"
	"testing"

	"github.com/syncromatics/proto-schema-registry/internal/storage"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

var (
	schema1 string = `syntax = "proto3";
	package gen;

	message first {
		string one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}`

	breakingSchema1 string = `syntax = "proto3";
	package gen;

	message first {
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}`

	schema1Version2 string = `syntax = "proto3";
	package gen;

	message first {
		string one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
		int32 two = 2;
	}`

	invalidProto string = `syntax = "proto3";
	whatever
	`
)

func Test_FileStorage_ShouldOnlyAcceptValidGzippedSchema(t *testing.T) {
	tID := uuid.New()

	p := path.Join("/tmp", tID.String())
	file, err := storage.NewFileStorage(kafkaBroker, 1, p, fmt.Sprintf("_proto_test_schemas__%s", tID))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)
	defer cancel()

	grp.Go(file.Run(ctx))

	schema := []byte{0, 0, 0}
	_, validationErrors, ok, err := file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchema_1", schema)
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, ok)
	assert.Equal(t, []string{
		"invalid gzip: failed to create gzip reader: unexpected EOF",
	}, validationErrors)
}

func Test_FileStorage_ShouldOnlyAcceptValidProtobuf(t *testing.T) {
	var b bytes.Buffer
	err := gzipWrite(&b, []byte(invalidProto))
	if err != nil {
		t.Fatal(err)
	}

	schemaBytes := b.Bytes()

	tID := uuid.New()

	p := path.Join("/tmp", tID.String())
	file, err := storage.NewFileStorage(kafkaBroker, 1, p, fmt.Sprintf("_proto_test_schemas__%s", tID))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)
	defer cancel()

	grp.Go(file.Run(ctx))

	_, validationErrors, ok, err := file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchema_1", schemaBytes)
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, ok)
	assert.Equal(t, []string{
		"invalid proto: <input>:2:2: found \"whatever\" but expected [.proto element {comment|option|import|syntax|enum|service|package|message}]",
	}, validationErrors)
}

func Test_FileStorage_AddSchema(t *testing.T) {
	var b bytes.Buffer
	err := gzipWrite(&b, []byte(schema1))
	if err != nil {
		t.Fatal(err)
	}

	schemaBytes := b.Bytes()

	tID := uuid.New()

	p := path.Join("/tmp", tID.String())
	file, err := storage.NewFileStorage(kafkaBroker, 1, p, fmt.Sprintf("_proto_test_schemas__%s", tID))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)
	defer cancel()

	grp.Go(file.Run(ctx))

	id, _, ok, err := file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchema_1", schemaBytes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, int64(1), id)

	gSchema, ok, err := file.GetSchema(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, schemaBytes, gSchema)

	id, _, ok, err = file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchema_2", schemaBytes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, int64(2), id)

	gSchema, ok, err = file.GetSchema(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, schemaBytes, gSchema)
}

func Test_FileStorage_AddSchemaShouldReturnIdIfItAlreadyExists(t *testing.T) {
	var b bytes.Buffer
	err := gzipWrite(&b, []byte(schema1))
	if err != nil {
		t.Fatal(err)
	}

	schemaBytes := b.Bytes()

	tID := uuid.New()

	p := path.Join("/tmp", tID.String())
	file, err := storage.NewFileStorage(kafkaBroker, 1, p, fmt.Sprintf("_proto_test_schemas__%s", tID))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)
	defer cancel()

	grp.Go(file.Run(ctx))

	id, _, ok, err := file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchemaShouldReturnIdIfItAlreadyExists", schemaBytes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, true, id == 1)

	gSchema, ok, err := file.GetSchema(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, schemaBytes, gSchema)

	id, _, ok, err = file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchemaShouldReturnIdIfItAlreadyExists", schemaBytes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, true, id == 1)
}

func Test_FileStorage_AddSchemaShouldCheckForBreakingChanges(t *testing.T) {
	tID := uuid.New()

	p := path.Join("/tmp", tID.String())
	file, err := storage.NewFileStorage(kafkaBroker, 1, p, fmt.Sprintf("_proto_test_schemas__%s", tID))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)
	defer cancel()

	grp.Go(file.Run(ctx))

	var b1 bytes.Buffer
	err = gzipWrite(&b1, []byte(schema1))
	if err != nil {
		t.Fatal(err)
	}

	id, _, ok, err := file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchemaShouldCheckForBreakingChanges", b1.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, int64(1), id)

	var b2 bytes.Buffer
	err = gzipWrite(&b2, []byte(breakingSchema1))
	if err != nil {
		t.Fatal(err)
	}

	_, errors, ok, err := file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchemaShouldCheckForBreakingChanges", b2.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.Equal(t, 1, len(errors))
}

func Test_FileStorage_AddSchemaShouldAllowGoodChanges(t *testing.T) {
	tID := uuid.New()

	p := path.Join("/tmp", tID.String())
	file, err := storage.NewFileStorage(kafkaBroker, 1, p, fmt.Sprintf("_proto_test_schemas__%s", tID))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)
	defer cancel()

	grp.Go(file.Run(ctx))

	var b1 bytes.Buffer
	err = gzipWrite(&b1, []byte(schema1))
	if err != nil {
		t.Fatal(err)
	}

	id, _, ok, err := file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchemaShouldCheckForBreakingChanges", b1.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, true, id == 1)

	var b2 bytes.Buffer
	err = gzipWrite(&b2, []byte(schema1Version2))
	if err != nil {
		t.Fatal(err)
	}

	id, _, ok, err = file.RegisterSchema(context.Background(), "Test_FileStorage_AddSchemaShouldCheckForBreakingChanges", b2.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
	assert.Equal(t, int64(2), id)
}

func gzipWrite(w io.Writer, data []byte) error {
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer gw.Close()
	gw.Write(data)
	return err
}
