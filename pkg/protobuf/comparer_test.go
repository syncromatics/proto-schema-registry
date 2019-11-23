package protobuf_test

import (
	"testing"

	"github.com/syncromatics/proto-schema-registry/pkg/protobuf"

	"gotest.tools/assert"
)

func Test_Comparer_WithFieldTypeChange(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		string one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		int32 one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' had field type 'string' changed to 'int32' at sequence '1'",
	}, problems)
}

func Test_Comparer_WithFieldRemove(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		string one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' does not contain field sequence '1' with field type 'string'",
	}, problems)
}

func Test_Comparer_WithFieldReserved(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		string one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		reserved 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, _, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
}

func Test_Comparer_WithFieldReservedRemoved(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		reserved 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' reserved '1' was removed",
	}, problems)
}

func Test_Comparer_WithMessageRemove(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		string one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message second {
		string one = 1;
		int32 two = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"new definition does not contain the message 'first'",
	}, problems)
}

func Test_Comparer_WithFieldRepeatedChange(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		repeated string one = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		string one = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' had field repeated 'true' changed to 'false' at sequence '1'",
	}, problems)
}

func Test_Comparer_WithEnumChangedValue(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	enum first {
		ZERO = 0;
		ONE = 1;
		TWO = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	enum first {
		WHOOPS = 0;
		ONE = 1;
		TWO = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"enum 'first' had value changed from '{ZERO}' to '{WHOOPS}' at sequence '0'",
	}, problems)
}

func Test_Comparer_WithEnumMissingValue(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	enum first {
		ZERO = 0;
		ONE = 1;
		TWO = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	enum first {
		ONE = 1;
		TWO = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"enum 'first' does not contain value sequence '0' with name '{ZERO}'",
	}, problems)
}

func Test_Comparer_WithMissingEnum(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	enum first {
		ZERO = 0;
		ONE = 1;
		TWO = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	enum last {
		ZERO = 0;
		ONE = 1;
		TWO = 2;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"new definition does not contain enum 'first'",
	}, problems)
}

func Test_Comparer_WithMissingEmbeddedEnum(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		enum enum1 {
			ZERO = 0;
			ONE = 1;
			TWO = 2;
		}

		enum1 enum_value = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		enum1 enum_value = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' is missing embedded enum 'enum1'",
	}, problems)
}

func Test_Comparer_WithEmbeddedEnumValueChange(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		enum enum1 {
			ZERO = 0;
			ONE = 1;
			TWO = 2;
		}

		enum1 enum_value = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		enum enum1 {
			ZERO = 0;
			WHOOPS = 1;
			TWO = 2;
		}
		enum1 enum_value = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' embedded enum 'enum1' value changed from '{ONE}' to '{WHOOPS}' at sequence '1'",
	}, problems)
}

func Test_Comparer_WithEmbeddedEnumMissingValue(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		enum enum1 {
			ZERO = 0;
			ONE = 1;
			TWO = 2;
		}

		enum1 enum_value = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;

	message first {
		enum enum1 {
			ZERO = 0;
			TWO = 2;
		}
		enum1 enum_value = 1;
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' embedded enum 'enum1' is missing value '{ONE}' at sequence '1'",
	}, problems)
}

func Test_Comparer_WithMissingOneOf(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		oneof oneof_1 {
			string oneof_1_string = 1;
			int32 oneof_1_int32 = 2;
		}
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;
	
	message first {}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' does not contain oneOf 'oneof_1'",
	}, problems)
}

func Test_Comparer_WithOneOfMissingValue(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		oneof oneof_1 {
			string oneof_1_string = 1;
			int32 oneof_1_int32 = 2;
		}
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;
	
	message first {
		oneof oneof_1 {
			int32 oneof_1_int32 = 2;
		}
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' oneOf 'oneof_1' does not contain sequence '1' with type 'string'",
	}, problems)
}

func Test_Comparer_WithOneOfReservedValue(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		oneof oneof_1 {
			string oneof_1_string = 1;
			int32 oneof_1_int32 = 2;
		}
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;
	
	message first {
		reserved 1;
		oneof oneof_1 {
			int32 oneof_1_int32 = 2;
		}
	}

	message record {
		first first_message = 1;
	}
	`

	ok, _, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, ok)
}

func Test_Comparer_WithOneOfChangedValue(t *testing.T) {
	current := `syntax = "proto3";
	package gen;

	message first {
		oneof oneof_1 {
			string oneof_1_string = 1;
			int32 oneof_1_int32 = 2;
		}
	}

	message record {
		first first_message = 1;
	}
	`

	new := `syntax = "proto3";
	package gen;
	
	message first {
		oneof oneof_1 {
			int64 oneof_1_string = 1;
			int32 oneof_1_int32 = 2;
		}
	}

	message record {
		first first_message = 1;
	}
	`

	ok, problems, err := protobuf.CheckForBreakingChanges([]byte(current), []byte(new))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, ok)
	assert.DeepEqual(t, []string{
		"message 'first' oneOf 'oneof_1' type changed from 'string' to 'int64' at sequence '1'",
	}, problems)
}
