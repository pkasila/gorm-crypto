package serialization

import (
	"testing"
)

func TestJSONSerializationCycle(t *testing.T) {
	jsonSerializer := NewJSON()
	value := "123"

	bytes, err := jsonSerializer.Serialize(value)
	if err != nil {
		t.Fatalf("Failed to serialize: %s\n", err.Error())
	}

	deserialized, err := jsonSerializer.Deserialize(bytes)
	if err != nil {
		t.Fatalf("Failed to serialize: %s\n", err.Error())
	}

	if (*deserialized).(string) != value {
		t.Fatalf("Source and deserialized values aren't equal: %s != %s\n", (*deserialized).(string), value)
	}
}