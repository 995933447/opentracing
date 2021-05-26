package main

import (
	opentracingUtil "github.com/995933447/opentracing"
	zipkinTracerConfig "github.com/995933447/opentracing/tracer_config/zipkin"
	"github.com/opentracing/opentracing-go"
	"os"
	"time"
)

func main()  {
	hostName, _ := os.Hostname()
	err := opentracingUtil.BuildDefaultGlobalTracer(
		hostName,
		"http://localhost:9411/api/v2/spans",
		"hi-drone-app-gateway",
		zipkinTracerConfig.AlwaysSampleType,
	)
	if err != nil {
		panic(err)
	}
	span := opentracing.StartSpan("test-span")
	defer span.Finish()
	time.Sleep(time.Second)
}