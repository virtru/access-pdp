package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pbPDP "github.com/virtru/access-pdp/proto/accesspdp/v1"
	pbAttr "github.com/virtru/access-pdp/proto/attributes/v1"


	pdp "github.com/virtru/access-pdp/pdp"
	attrs "github.com/virtru/access-pdp/attributes"


	"github.com/virtru/oteltracer"
	"go.opentelemetry.io/otel"
	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

var svcName = "access-pdp"

var cfg EnvConfig

var tracer = otel.Tracer("main")

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

func PbToAttributeInstances(pbinst []*pbAttr.AttributeInstance) []attrs.AttributeInstance {
	var instances []attrs.AttributeInstance


	if pbinst != nil {
		for _, v := range pbinst {
			instances = append(instances, attrs.AttributeInstance{Authority: v.Authority, Name: v.Name, Value: v.Value})
		}
	}

	return instances
}

func PbToEntityAttrSets(pbsets map[string]*pbPDP.ListOfAttributeInstances) map[string][]attrs.AttributeInstance {
	entitySets := make(map[string][]attrs.AttributeInstance)

	if pbsets != nil {
		for entity, instances := range pbsets {

			var convAttrs []attrs.AttributeInstance
			if instances != nil {
				convAttrs = PbToAttributeInstances(instances.AttributeInstances)
			}

			entitySets[entity] = convAttrs
		}
	}
	return entitySets
}

func PbToAttributeDefinitions(pbdefs []*pbAttr.AttributeDefinition) []attrs.AttributeDefinition {
	var defs []attrs.AttributeDefinition


	if pbdefs != nil {
		for _, v := range pbdefs {
			convAttr := attrs.AttributeDefinition{
				Authority: v.Authority,
				Name: v.Name,
				Rule: v.Rule,
				State: *v.State,
				Order: v.Order,
			}

			//GroupBy is optional - if it is present, it is just represented as another AttributeInstance
			if v.GroupBy != nil {
				convAttr.GroupBy = &attrs.AttributeInstance{Authority: v.GroupBy.Authority, Name: v.GroupBy.Name, Value: v.GroupBy.Value}
			}

			defs = append(defs, convAttr)
		}
	}

	return defs
}

func DataRuleResultsToPb(results []*pdp.DataRuleResult) []*pbPDP.DataRuleResult {
	var pbresults []*pbPDP.DataRuleResult


	if results != nil {
		for _, v := range results {
			pbresults = append(pbresults, &pbPDP.DataRuleResult{Passed: v.Passed, RuleDefinition: v.RuleDefinition, ValueFailures: v.ValueFailures})
		}
	}

	return instances
}

func DecisionToPbResponse(entity string, decision *pdp.Decision) pbPDP.DetermineAccessResponse {

	return pbPDP.DetermineAccessResponse{
		Entity: entity,
		Access: decision.Access,
	}
}
func (s *accessPDPServer) DetermineAccess(req *pbPDP.DetermineAccessRequest, stream pbPDP.AccessPDPEndpoint_DetermineAccessServer) error {


	dataAttrs := PbToAttributeInstances(req.DataAttributes)
	entityAttrSets := PbToEntityAttrSets(req.EntityAttributeSets)
	definitions := PbToAttributeDefinitions(req.AttributeDefinitions)

	s.logger.Debug("DetermineAccess gRPC endpoint")
	handlerCtx, handlerSpan := tracer.Start(stream.Context(), "DetermineAccess gRPC")
	defer handlerSpan.End()

	entityDecisions, err := s.accessPDP.DetermineAccess(dataAttrs, entityAttrSets, definitions, handlerCtx)
	if err != nil {
		return err
	}

	for entity, decisions := range entityDecisions {
		DecisionsToPb(decisions)
	}


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
