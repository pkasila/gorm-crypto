package serialization

type Serializer interface {
	Serialize(value interface{}) ([]byte, error)
	Deserialize([]byte) (*interface{}, error)
}
