package main

import (
	"fmt"
	"math/rand"
	"time"

	context "golang.org/x/net/context"

	"./party"
	"google.golang.org/grpc"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := 3000 + r.Intn(3)
	server := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Go connecting to %s\n", server)
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
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
