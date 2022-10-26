package types

type Serializer interface {
	Marshal() ([]byte, error)
}

type Deserializer interface {
	Unmarshal(bin []byte) error
}
