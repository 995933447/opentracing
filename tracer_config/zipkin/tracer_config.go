package zipkin

import (
	"github.com/995933447/opentracing/tracer_config"
	zipkinTracer "github.com/openzipkin/zipkin-go"
	"strconv"
)

type TracerConfig struct {
	BrokerAddress string
	LocalEndPointer LocalEndPointer
	Sampler zipkinTracer.Sampler
}

type LocalEndPointer struct {
	Name string
	Address string
}

type SampleType int

const (
	AlwaysSampleType = 0
)

const (
	BrokerAddressOptionName = "broker_address"
	LocalEndPointerOptionName = "local_end_pointer"
	SamplerTypeOptionName = "sampler"
	)

// New a zipkin deriver config.
func NewTracerConfig(
	brokerAddress,
	localEndPointerName,
	localEndPointerAddress string,
	sampleType SampleType,
	) *TracerConfig {
	tracerConfig := new(TracerConfig)
	tracerConfig.SetBrokerAddress(brokerAddress)
	tracerConfig.setLocalEndPointer(LocalEndPointer{Name: localEndPointerName, Address: localEndPointerAddress})
	switch sampleType {
		case AlwaysSampleType:
			tracerConfig.setSampler(zipkinTracer.AlwaysSample)
		default:
			panic("Do not support sampler type " + strconv.Itoa(int(sampleType)))
	}

	return tracerConfig
}

func (tracerConfig TracerConfig) SetOption(optionName string, optionValue interface{}) *tracer_config.TracerConfig {
	switch optionName {
		case BrokerAddressOptionName:
			tracerConfig.SetBrokerAddress(optionValue.(string))
		case LocalEndPointerOptionName:
			tracerConfig.setLocalEndPointer(optionValue.(LocalEndPointer))
		case SamplerTypeOptionName:
			tracerConfig.setSampler(optionValue.(zipkinTracer.Sampler))
	}

	var openTracingTracerConfig tracer_config.TracerConfig
	openTracingTracerConfig = tracerConfig

	return &openTracingTracerConfig
}

func (tracerConfig *TracerConfig) setSampler(sampler zipkinTracer.Sampler) {
	tracerConfig.Sampler = sampler
}

func (tracerConfig *TracerConfig) setLocalEndPointer(localEndPointer LocalEndPointer) {
	tracerConfig.LocalEndPointer = localEndPointer
}

func (tracerConfig *TracerConfig) SetBrokerAddress(brokerAddress string) {
	tracerConfig.BrokerAddress = brokerAddress
}