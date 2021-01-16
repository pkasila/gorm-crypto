package gormcrypto

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type EncryptedValue struct {
	Raw interface{} `json:"Raw"`
}

// Scan decrypts and deserializes value from DB, implements sql.Scanner interface
func (j *EncryptedValue) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	var encValue *interface{}
	var selectionError error

	for _, algo := range Algorithms {
		var dBytes []byte
		dBytes, selectionError = algo.Decrypt(bytes)
		if selectionError != nil {
			continue
		}

		for _, serializer := range Serializers {
			encValue, selectionError = serializer.Deserialize(dBytes)
			if selectionError != nil {
				continue
			}
		}
	}

	if selectionError != nil {
		return selectionError
	}

	j.Raw = (*encValue).(map[string]interface{})["Raw"]

	return nil
}

// Value returns serialized and encrypted value, implement driver.Valuer interface
func (j EncryptedValue) Value() (driver.Value, error) {
	bytes, err := Serializers[0].Serialize(j)

	if err != nil {
		return nil, err
	}

	return Algorithms[0].Encrypt(bytes)
}
