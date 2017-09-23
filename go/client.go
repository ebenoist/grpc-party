package main

import (
	"fmt"

	context "golang.org/x/net/context"

	"./party"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := party.NewPartyClient(conn)
	resp, err := client.PartyHard(context.Background(), &party.User{Name: "Go"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Msg)
}
