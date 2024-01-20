package serialize

import (
	"github.com/tomoya_kamaji/protobuf/src/proto/gen"
	"google.golang.org/protobuf/proto"
)

type ProtocolBuffersSerializer struct{}

func NewProtocolBuffersSerializer () *ProtocolBuffersSerializer {
		return &ProtocolBuffersSerializer{}
}

func (s *ProtocolBuffersSerializer) Marshal(p *gen.Person) ([]byte, error) {
    data, err := proto.Marshal(p)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func (s *ProtocolBuffersSerializer) Unmarshal(data []byte) (*gen.Person, error) {
    p := new(gen.Person)
    err := proto.Unmarshal(data, p)
    if err != nil {
        return nil, err
    }
    return p, nil
}