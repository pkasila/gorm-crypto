package gormcrypto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type EncryptedValue struct {
	Raw interface{}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *EncryptedValue) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	bytes, err := Algorithm.Decrypt(bytes)
	if err != nil {
		return err
	}

	var encValue EncryptedValue
	err = json.Unmarshal(bytes, &encValue)
	if err != nil {
		return err
	}
	*j = encValue

	return nil
}

// Value return json value, implement driver.Valuer interface
func (j EncryptedValue) Value() (driver.Value, error) {
	bytes, err := json.Marshal(j)

	if err != nil {
		return nil, err
	}

	return Algorithm.Encrypt(bytes)
}
