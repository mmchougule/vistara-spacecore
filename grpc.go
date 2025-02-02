package main

import (
	"context"
	"log"
	"spaceco/proto"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Define the plugin map
var PluginMap = map[string]plugin.Plugin{
	"spacecore": &Spacecore{Impl: &SpacecorePlugin{}},
}

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SPACECORE_PLUGIN",
	MagicCookieValue: "v1",
}

type SpacecorePlugin struct{}

// Define the plugin implementation
type Spacecore struct {
	plugin.Plugin
	plugin.GRPCPlugin
	Impl *SpacecorePlugin
}

// GRPCServer function
type GRPCServer struct {
	Impl SpacecorePlugin
}

// GRPCServer function
func (g *Spacecore) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterScPluginServer(s, &GRPCServer{Impl: *g.Impl})
	return nil
}

func (g *GRPCServer) Start(ctx context.Context, in *proto.StartRequest) (*proto.StartResponse, error) {
	log.Printf("Starting Spacecore...\n")
	msg, err := g.Impl.Start(ctx)
	if err != nil {
		return nil, err
	}
	return &proto.StartResponse{Status: msg}, nil
}

func (g *GRPCServer) Stop(ctx context.Context, in *proto.StopRequest) (*proto.StopResponse, error) {
	msg, err := g.Impl.Stop(context.Background())
	if err != nil {
		return nil, err
	}
	return &proto.StopResponse{Status: msg}, nil
}
func (g *GRPCServer) Status(ctx context.Context, in *proto.StatusRequest) (*proto.StatusResponse, error) {
	log.Printf("Checking Spacecore status...\n")
	msg, err := g.Impl.Status(ctx)
	return &proto.StatusResponse{Status: msg}, err
}
func (g *GRPCServer) Logs(ctx context.Context, in *proto.LogsRequest) (*proto.LogsResponse, error) {
	msg, err := g.Impl.Start(context.Background())
	log.Printf("Checking Spacecore logs...\n %s", msg)
	if err != nil {
		return nil, err
	}
	return &proto.LogsResponse{}, nil
}
