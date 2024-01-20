package serialize

import (
	"fmt"
	"testing"

	"github.com/tomoya_kamaji/protobuf/src/proto/gen"
	"github.com/tomoya_kamaji/protobuf/src/redis"
)

func TestProtocolBuffersSerializer(t *testing.T) {
    client :=redis.NewRedisClient()
    s := NewProtocolBuffersSerializer()
    testCases := []struct {
        name    string
        p       *gen.Person
        wantErr bool
    }{
        {"Test 1", &gen.Person{Name: "hoge",Email:"Email"}, false},
        {"Test 2", &gen.Person{Name: "fuga"}, false},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // シリアライズ
            data, err := s.Marshal(tc.p)
            if (err != nil) != tc.wantErr {
                t.Fatalf("Marshal() error = %v, wantErr %v", err, tc.wantErr)
            }


            client.Set("key",data)

            v,err := client.Get("key")
            if err != nil {
                t.Fatalf("Failed to get: %v", err)
                return
            }

            // デシリアライズ
            p2, err := s.Unmarshal(v)
            if err != nil {
                t.Fatalf("Failed to unmarshal: %v", err)
            }
            fmt.Printf("p2: %v\n", p2)

            // 元のPerson構造体とデシリアライズしたPerson構造体が一致することを確認
            if tc.p.Name != p2.Name {
                t.Errorf("Unmarshaled Person does not match original. got: %v, want: %v", p2.Name, tc.p.Name)
            }
        })
    }
}