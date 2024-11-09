package exec

import (
	"log"
	"protocols/exec/gen"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func ProtobufCase(){
	elliot := gen.Person{
		Name: "Eliot",
		Age: 24,
	}
	
	data, err := proto.Marshal(&elliot)
	if err != nil{
		log.Fatal("marshaling error", err)
	}
	fmt.Println(data)

	newElliot := &gen.Person{}
	err = proto.Unmarshal(data, newElliot)
	
	if err != nil{
		log.Fatal("UnMarshaling error", err)
	}

	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot)

}