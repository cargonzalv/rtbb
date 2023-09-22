package serializer

import (
	"encoding/json"
)

// This service is using `json` package for marshalling and un-matrhsalling data.
type jsonSerializer struct {
}

var _ JsonSerializer = (*jsonSerializer)(nil)

func (s *jsonSerializer) MarshalJson(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (s *jsonSerializer) UnmarshalJson(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
