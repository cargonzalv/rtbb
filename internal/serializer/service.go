package serializer

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// This service is using protojson and json for marshalling and un-matrhsalling
// Similar functionality can also be implemented with other package like jsonpb
type service struct {
}

var _ JsonSerializer = (*service)(nil)

func (s *service) Marshal(m proto.Message) ([]byte, error) {
	return protojson.Marshal(m)
}

func (s *service) MarshalJson(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (s *service) Unmarshal(data []byte, m proto.Message) error {
	return protojson.Unmarshal(data, m)
}

func (s *service) UnmarshalJson(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
