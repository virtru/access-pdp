package main

import (
	ctx "context"
	"fmt"
	"log"
	"net"

	pbPDP "github.com/virtru/access-pdp/proto/accesspdp/v1"
	pbConv "github.com/virtru/access-pdp/protoconv"

	pdp "github.com/virtru/access-pdp/pdp"

	//This allows clients to introspect the server
	//and list operations - can be removed if that is not desired.
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
	"google.golang.org/grpc/examples/data"

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
	ListenPort      string `env:"LISTEN_PORT" envDefault:"50052"`
	ListenHost      string `env:"LISTEN_HOST" envDefault:"localhost"`
	Verbose         bool   `env:"VERBOSE" envDefault:"false"`
	DisableTracing  bool   `env:"DISABLE_TRACING" envDefault:"false"`
	EnableGRPCTLS   bool   `env:"ENABLE_GRPC_TLS" envDefault:"false"`
	GRPCTLSCertFile string `env:"GRPC_TLS_CERTFILE" envDefault:"x509/server_cert.pem"`
	GRPCTLSKeyFile  string `env:"GRPC_TLS_KEYFILE" envDefault:"x509/server_key.pem"`
}

type accessPDPServer struct {
	logger    *zap.SugaredLogger
	accessPDP *pdp.AccessPDP
	pbPDP.UnimplementedAccessPDPEndpointServer
	pbPDP.UnimplementedHealthServer
}

func (s *accessPDPServer) Check(parentCtx ctx.Context, req *pbPDP.HealthCheckRequest) (*pbPDP.HealthCheckResponse, error) {
	s.logger.Debug("Health check endpoint hit")
	return &pbPDP.HealthCheckResponse{
		Status: 1,
	}, nil
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

	var opts []grpc.ServerOption
	if cfg.EnableGRPCTLS {
		logger.Debug("gRPC TLS mode enabled, checking for server TLS keys...")
		var certFile, keyFile string
		if cfg.GRPCTLSCertFile == "" {
			certFile = data.Path(cfg.GRPCTLSCertFile)
			logger.Debugf("Found server certfile at %s", cfg.GRPCTLSCertFile)
		}
		if cfg.GRPCTLSKeyFile == "" {
			keyFile = data.Path(cfg.GRPCTLSKeyFile)
			logger.Debugf("Found server keyfile at %s", cfg.GRPCTLSKeyFile)
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logger.Fatalf("Failed to generate credentials %v", err)
		}
		logger.Info("Starting gRPC server with TLS")
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	} else {
		logger.Info("Starting gRPC server without TLS")
	}

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	srv := newAccessPDPSrv(logger)
	pbPDP.RegisterAccessPDPEndpointServer(grpcServer, srv)
	pbPDP.RegisterHealthServer(grpcServer, srv)

	logger.Info("Serving gRPC endpoint")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Error on serve!", zap.Error(err))
	}
}
