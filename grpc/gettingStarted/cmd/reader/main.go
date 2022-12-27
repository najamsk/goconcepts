package main

import (
	"fmt"
	"log"
	"os"

	userpb "mygrpc/gen/go/user/v1"

	"google.golang.org/protobuf/proto"
)

func main() {
	b, err := os.ReadFile("../writer/user.bin")
	if err != nil {
		log.Fatalln("Writing error", err)
	}
	fmt.Println(string(b))
	var pu userpb.User = userpb.User{}
	err = proto.Unmarshal(b, &pu)
	if err != nil {
		log.Fatalln("unmarshalling error", err)
	}
	fmt.Printf("name: %v \n", pu.FullName)
	fmt.Printf("birthyear: %v \n", pu.BirthYear)
	fmt.Printf("uuid: %v \n", pu.Uuid)

}
