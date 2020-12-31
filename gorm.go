package gormcrypto

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type EncryptedValue struct {
	Raw interface{}	`json:"Raw"`
}

// Scan decrypts and deserializes value from DB, implements sql.Scanner interface
func (j *EncryptedValue) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	bytes, err := Algorithm.Decrypt(bytes)
	if err != nil {
		return err
	}

	encValue, err := Serializer.Deserialize(bytes)
	if err != nil {
		return err
	}
	j.Raw = (*encValue).(map[string]interface{})["Raw"]

	return nil
}

// Value returns serialized and encrypted value, implement driver.Valuer interface
func (j EncryptedValue) Value() (driver.Value, error) {
	bytes, err := Serializer.Serialize(j)

	if err != nil {
		return nil, err
	}

	return Algorithm.Encrypt(bytes)
}
