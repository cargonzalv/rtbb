package serializer

import (
	"google.golang.org/protobuf/proto"
)

type JsonSerializer interface {
	MarshalJson(v any) ([]byte, error)
	UnmarshalJson(data []byte, v any) error
}

type ProtobufSerializer interface {
	Marshal(m proto.Message) ([]byte, error)
	Unmarshal(b []byte, m proto.Message) error
}
