package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	"./party"
)

type service struct{}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9998))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	party.RegisterPartyServer(grpcServer, &service{})
	grpcServer.Serve(lis)
}

func (s *service) PartyHard(ctx context.Context, u *party.User) (*party.PartyResp, error) {
	return &party.PartyResp{Msg: fmt.Sprintf("go wants %s to party hard 🤘", u.Name)}, nil
}
