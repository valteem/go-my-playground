package containers

type JSONSerializer interface {
	ToJSON() ([]byte, error)      // JSON representation of the content of the container
	MarshalJSON() ([]byte, error) // implements json.Marshaler
}

type JSONDeserializer interface {
	FromJSON([]byte) error      // populates container from input JSON
	UnmarshalJSON([]byte) error // implements json.Unmarshaler
}
