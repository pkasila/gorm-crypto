package serialization

import (
	"encoding/json"
)

type JSON struct {
	Serializer
}

func NewJSON() JSON {
	return JSON{}
}

func (JSON) Serialize(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (JSON) Deserialize(bytes []byte) (*interface{}, error) {
	var value interface{}
	err := json.Unmarshal(bytes, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}