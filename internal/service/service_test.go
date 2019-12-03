package service_test

import (
	"context"
	"testing"

	"github.com/syncromatics/proto-schema-registry/internal/service"
	v1 "github.com/syncromatics/proto-schema-registry/pkg/proto/schema/registry/v1"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type storageMock struct {
	get      func(context.Context, uint32) ([]byte, bool, error)
	register func(context.Context, string, []byte) (uint32, []string, bool, error)
}

func (s *storageMock) GetSchema(ctx context.Context, id uint32) (schema []byte, ok bool, err error) {
	return s.get(ctx, id)
}

func (s *storageMock) RegisterSchema(ctx context.Context, topic string, schema []byte) (id uint32, errors []string, ok bool, err error) {
	return s.register(ctx, topic, schema)
}

func Test_Service_GetSchemaSuccess(t *testing.T) {
	schema := []byte{0x0, 0x1, 0x9}

	storage := &storageMock{
		get: func(context context.Context, id uint32) ([]byte, bool, error) {
			assert.Equal(t, uint32(42), id)
			return schema, true, nil
		},
	}

	service := service.NewService(storage)

	r, err := service.GetSchema(context.Background(), &v1.GetSchemaRequest{
		Id: 42,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Exists)
	assert.Equal(t, schema, r.Schema)
}

func Test_Service_GetSchemaNotFound(t *testing.T) {
	storage := &storageMock{
		get: func(context context.Context, id uint32) ([]byte, bool, error) {
			assert.Equal(t, uint32(43), id)
			return nil, false, nil
		},
	}

	service := service.NewService(storage)

	r, err := service.GetSchema(context.Background(), &v1.GetSchemaRequest{
		Id: 43,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, r.Exists)
	assert.Nil(t, r.Schema)
}

func Test_Service_GetSchemaError(t *testing.T) {
	expectedErr := errors.Errorf("failed")
	storage := &storageMock{
		get: func(context context.Context, id uint32) ([]byte, bool, error) {
			assert.Equal(t, uint32(44), id)
			return nil, false, expectedErr
		},
	}

	service := service.NewService(storage)

	r, err := service.GetSchema(context.Background(), &v1.GetSchemaRequest{
		Id: 44,
	})
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, r)
}

func Test_Service_RegisterSchemaSuccess(t *testing.T) {
	requestSchema := []byte{0x9, 0x7, 0x5}
	storage := &storageMock{
		register: func(ctx context.Context, topic string, schema []byte) (uint32, []string, bool, error) {
			assert.Equal(t, "Test_Service_RegisterSchemaSuccess", topic)
			assert.Equal(t, requestSchema, schema)
			return 1, nil, true, nil
		},
	}

	service := service.NewService(storage)

	r, err := service.RegisterSchema(context.Background(), &v1.RegisterSchemaRequest{
		Topic:  "Test_Service_RegisterSchemaSuccess",
		Schema: requestSchema,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &v1.RegisterSchemaResponse_ResponseSuccess{
		ResponseSuccess: &v1.RegisterSchemaSuccess{
			Id: 1,
		},
	}, r.Response)
}

func Test_Service_RegisterSchemaWithSchemaErrors(t *testing.T) {
	requestSchema := []byte{0x9, 0x7, 0x5}
	schemaErrors := []string{"bad error", "even badder error"}
	storage := &storageMock{
		register: func(ctx context.Context, topic string, schema []byte) (uint32, []string, bool, error) {
			assert.Equal(t, "Test_Service_RegisterSchemaWithSchemaErrors", topic)
			assert.Equal(t, requestSchema, schema)
			return 0, schemaErrors, false, nil
		},
	}

	service := service.NewService(storage)

	r, err := service.RegisterSchema(context.Background(), &v1.RegisterSchemaRequest{
		Topic:  "Test_Service_RegisterSchemaWithSchemaErrors",
		Schema: requestSchema,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &v1.RegisterSchemaResponse_ResponseError{
		ResponseError: &v1.RegisterSchemaError{
			Errors: schemaErrors,
		},
	}, r.Response)
}

func Test_Service_RegisterSchemaError(t *testing.T) {
	requestSchema := []byte{0x9, 0x7, 0x5}
	expectedErr := errors.Errorf("failed")
	storage := &storageMock{
		register: func(ctx context.Context, topic string, schema []byte) (uint32, []string, bool, error) {
			assert.Equal(t, "Test_Service_RegisterSchemaError", topic)
			assert.Equal(t, requestSchema, schema)
			return 0, nil, false, expectedErr
		},
	}

	service := service.NewService(storage)

	r, err := service.RegisterSchema(context.Background(), &v1.RegisterSchemaRequest{
		Topic:  "Test_Service_RegisterSchemaError",
		Schema: requestSchema,
	})
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, r)
}
