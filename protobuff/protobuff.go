package protobuff

import (
	"github.com/golang/protobuf/proto"
	"errors"
	"log"
)

func Main () {

	m1 := Ship_CrewMember {
		Id: 1,
		Name: "Rob",
		Position: "Captain",
		SecClearance: 123,
	}
	m2 := Ship_CrewMember {
		Id: 1,
		Name: "Tyler",
		Position: "Software Engineer",
		SecClearance: 456,
	}

	ship := Ship {
		Shipname: "The last ship",
		CaptainName: "Rob",
		Crew: []*Ship_CrewMember{&m1, &m2},
	}

	log.Print("Proto Encoded")
	encoded, _ := EncodeProto(&ship)
	log.Print(encoded)
	log.Print("......")
	log.Print("Proto Decoded")
	log.Print(DecodeProto(encoded))

}

func EncodeProto(obj interface{}) ([]byte, error) {
	if v, ok := obj.(*Ship); ok {
		return proto.Marshal(v)
	}
	return nil, errors.New("Proto: Unknown message type")
}

func DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)
	return pb, proto.Unmarshal(buffer, pb)
}

