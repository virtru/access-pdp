package main

import (
	"fmt"
	"log"
	"net"

	pbPDP "github.com/virtru/access-pdp/proto/accesspdp/v1"
	pbConv "github.com/virtru/access-pdp/protoconv"
	"google.golang.org/grpc"

	pdp "github.com/virtru/access-pdp/pdp"

	"github.com/caarlos0/env"
	"github.com/virtru/oteltracer"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var svcName = "access-pdp"

var cfg EnvConfig

var tracer = otel.Tracer("main")

//Env config
type EnvConfig struct {
	ListenPort     string `env:"LISTEN_PORT" envDefault:"50052"`
	ListenHost     string `env:"LISTEN_HOST" envDefault:"localhost"`
	Verbose        bool   `env:"VERBOSE" envDefault:"false"`
	DisableTracing bool   `env:"DISABLE_TRACING" envDefault:"false"`
}

type accessPDPServer struct {
	logger    *zap.SugaredLogger
	accessPDP *pdp.AccessPDP
	pbPDP.UnimplementedAccessPDPEndpointServer
}

func (s *accessPDPServer) DetermineAccess(req *pbPDP.DetermineAccessRequest, stream pbPDP.AccessPDPEndpoint_DetermineAccessServer) error {
	dataAttrs := pbConv.PbToAttributeInstances(req.DataAttributes)
	entityAttrSets := pbConv.PbToEntityAttrSets(req.EntityAttributeSets)
	definitions := pbConv.PbToAttributeDefinitions(req.AttributeDefinitions)

	s.logger.Debug("DetermineAccess gRPC endpoint")
	handlerCtx, handlerSpan := tracer.Start(stream.Context(), "DetermineAccess gRPC")
	defer handlerSpan.End()

	entityDecisions, err := s.accessPDP.DetermineAccess(dataAttrs, entityAttrSets, definitions, handlerCtx)
	if err != nil {
		return err
	}

	for entity, decisions := range entityDecisions {
		entityDecision := pbConv.DecisionToPbResponse(entity, decisions)
		if err := stream.Send(entityDecision); err != nil {
			return err
		}
	}
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
	pbPDP.RegisterAccessPDPEndpointServer(grpcServer, newAccessPDPSrv(logger))

	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Error on serve!", zap.Error(err))
	}
}
