package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/virtru/access-pdp/proto/accesspdp/v1"
	pdp "github.com/virtru/access-pdp/pdp"


	"github.com/virtru/oteltracer"
	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

var svcName = "access-pdp"

var cfg EnvConfig

//Env config
type EnvConfig struct {
	ListenPort                        string `env:"LISTEN_PORT" envDefault:"50052"`
	ListenHost                        string `env:"LISTEN_HOST" envDefault:"localhost"`
	Verbose                           bool   `env:"VERBOSE" envDefault:"false"`
	DisableTracing                    bool   `env:"DISABLE_TRACING" envDefault:"false"`
}

type accessPDPServer struct {
	logger *zap.SugaredLogger
	accessPDP *pdp.AccessPDP
	pb.UnimplementedAccessPDPEndpointServer
	// savedFeatures []*pb.Feature // read-only after initialized

	// mu         sync.Mutex // protects routeNotes
	// routeNotes map[string][]*pb.RouteNote
}

func (s *accessPDPServer) DetermineAccess(req *pb.DetermineAccessRequest, stream pb.AccessPDPEndpoint_DetermineAccessServer) error {

	s.accessPDP.DetermineAccess(req.DataAttributes, req.EntityAttributeSets, req.AttributeDefinitions)
		// pdpCtx, pdpSpan := tracer.Start(handlerCtx, "DetermineAccess")
		//1. Hit entitlements PDP first, to get entity attributes
		//2. Then hit Attribute Authority, to get attribute definitions for all data attributes
		//3. Then call PDP.
		// result, err := s.accessPDP.DetermineAccess(
		// 	dataAttributes,
		// 	entityAttributes,
		// 	definitions,
		// 	pdpCtx)
		// pdpSpan.End()
		// if err != nil {
		// 	s.logger.Errorf("Access PDP returned error! Error was %s", err)
		// 	// w.WriteHeader(http.StatusInternalServerError)
		// 	return err
		// }




	// for _, feature := range s.savedFeatures {
	// 	if inRange(feature.Location, rect) {
	// 		if err := stream.Send(feature); err != nil {
	// 			return err
	// 		}
	// 	}
	// }
	return nil
}

func newAccessPDPSrv(logger *zap.SugaredLogger) *accessPDPServer {
	accessPDP := pdp.NewAccessPDP(logger)
	s := &accessPDPServer{logger: logger, accessPDP: accessPDP}
	return s
}

func main() {

	var zapLog *zap.Logger
	var logErr error

	// Parse env
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err.Error())
	}

	if cfg.Verbose {
		fmt.Print("Enabling verbose logging")
		zapLog, logErr = zap.NewDevelopment() // or NewProduction, or NewDevelopment
	} else {
		fmt.Print("Enabling production logging")
		zapLog, logErr = zap.NewProduction()
	}

	if logErr != nil {
		log.Fatalf("Logger initialization failed!")
	}

	defer func() {
		err := zapLog.Sync()
		if err != nil {
			log.Fatal("Error flushing zap log!")
		}
	}()

	logger := zapLog.Sugar()

	logger.Infof("%s init", svcName)

	if !cfg.DisableTracing {
		tracerCancel, err := oteltracer.InitTracer(svcName)
		if err != nil {
			logger.Errorf("Error initializing tracer: %v", err)
		}
		defer tracerCancel()
	}

	///test
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.ListenHost, cfg.ListenPort))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAccessPDPEndpointServer(grpcServer, newAccessPDPSrv(logger))
	grpcServer.Serve(lis)
}
