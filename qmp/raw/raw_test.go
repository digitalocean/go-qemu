package raw

import (
	"bytes"
	"testing"
)

func TestKeyValueQcode(t *testing.T) {
	key := QKeyCodeShift
	code := KeyValueQcode(key)
	json, err := code.MarshalJSON()
	expectedJSON := []byte(`{"data":"shift","type":"qcode"}`)
	if err != nil {
		t.Fatalf("Failed to marshall %v: %s", code, err)
	}
	if !bytes.Equal(json, expectedJSON) {
		t.Fatalf("Expected %s, but output was %s", expectedJSON, json)
	}
}
