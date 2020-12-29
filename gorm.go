package gorm_crypto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pkosilo/gorm-crypto/encryption"
)

type EncryptedValue struct {
	Raw	interface{}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *EncryptedValue) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	bytes, err := encryption.DecryptWithPrivateKey(bytes, PrivateKey)
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

	return encryption.EncryptWithPublicKey(bytes, PublicKey)
}