package opentracing

import (
	"github.com/995933447/opentracing/tracer_config"
	zipkinTracerConfig "github.com/995933447/opentracing/tracer_config/zipkin"
	"github.com/opentracing/opentracing-go"
	opentracingWrapper "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinHttpReporter "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	ZipKinTracerDriver = 0
)

// Build a global tracer by specific opentracing driver.such as zipkin,jeager.
func BuildGlobalTracer(tracerDriverType int, tracerConfig *tracer_config.TracerConfig) error {
	var err error
	switch tracerDriverType {
		case ZipKinTracerDriver:
			theZipkinTracerConfig := (*tracerConfig).(zipkinTracerConfig.TracerConfig)
			err = buildZipkinGlobalTracer(&theZipkinTracerConfig)
		default:
			panic("Do not support tracer driver.")
	}
	return err
}

func buildZipkinGlobalTracer(config *zipkinTracerConfig.TracerConfig) error {
	reporter := zipkinHttpReporter.NewReporter(config.BrokerAddress)

	localEndpoint, err := zipkin.NewEndpoint(
		config.LocalEndPointer.Name,
		config.LocalEndPointer.Address,
	)

	if err != nil {
		return err
	}

	tracer, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(config.Sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)

	if err != nil {
		return err
	}

	globalTracer := opentracingWrapper.Wrap(tracer)
	opentracing.SetGlobalTracer(globalTracer)

	return nil
}

// Build a zikin deriver gloabl tracer.
func BuildDefaultGlobalTracer(hostPort, brokerAddress, serviceName string, sampleType zipkinTracerConfig.SampleType) error {
	var tracerConfig tracer_config.TracerConfig
	theZipkinTracerConfig := zipkinTracerConfig.NewTracerConfig(brokerAddress, serviceName, hostPort, sampleType)
	tracerConfig = *theZipkinTracerConfig
	return BuildGlobalTracer(ZipKinTracerDriver, &tracerConfig)
}