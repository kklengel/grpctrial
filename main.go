package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	Do()
}

func Do() error {

	chatMessage := &ChatMessage{
		Text: "hallo welt",
		Done: true,
	}

	fmt.Println(proto.MarshalTextString(chatMessage))
	return nil
}
