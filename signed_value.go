package gormcrypto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type SignedValue[T any] struct {
	Raw   interface{} `json:"Raw"`
	Valid bool        `json:"Valid"`
}

type signedValueInternal struct {
	Raw []byte `json:"Raw"`
	Sig []byte `json:"Sig"`
}

// Scan deserializes and verifies value from DB, implements sql.Scanner interface
func (j *SignedValue[T]) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	var svi signedValueInternal
	err := json.Unmarshal(bytes, &svi)
	if err != nil {
		return err
	}

	var encValue *interface{}
	var selectionError error

	for _, serializer := range SignatureSerializers {
		encValue, selectionError = serializer.Deserialize(svi.Raw)
		if selectionError != nil {
			continue
		}
	}

	if selectionError != nil {
		return selectionError
	}

	var valid bool

	for _, algo := range SignatureAlgorithms {
		valid0, err := algo.Verify(svi.Raw, svi.Sig)
		if err != nil {
			continue
		}
		valid = valid0
	}

	j.Raw = *encValue
	j.Valid = valid

	return nil
}

// Value returns serialized value with signature, implement driver.Valuer interface
func (j SignedValue[T]) Value() (driver.Value, error) {
	serializer := SignatureSerializers[0]

	bytes, err := serializer.Serialize(j.Raw)

	if err != nil {
		return nil, err
	}

	sig, err := SignatureAlgorithms[0].Sign(bytes)
	if err != nil {
		return nil, err
	}

	s, err := json.Marshal(signedValueInternal{
		Raw: bytes,
		Sig: sig,
	})
	if err != nil {
		return nil, err
	}

	return s, nil
}
