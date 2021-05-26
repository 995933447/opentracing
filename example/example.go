package main

import (
	"fmt"
	opentracingUtil "github.com/995933447/opentracing"
	"github.com/995933447/opentracing/tracer_config"
	zipkinTracerConfig "github.com/995933447/opentracing/tracer_config/zipkin"
	"github.com/opentracing/opentracing-go"
	"os"
	"time"
)

func main()  {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	var tracerConfig tracer_config.TracerConfig
	theZipkinTracerConfig := zipkinTracerConfig.NewTracerConfig(
		"http://localhost:9411/api/v2/spans",
		"test-opentracing",
		hostName,
		zipkinTracerConfig.AlwaysSampleType,
		)
	tracerConfig = *theZipkinTracerConfig

	err = opentracingUtil.BuildGlobalTracer(opentracingUtil.ZipKinTracerDriver, &tracerConfig)
	if err != nil {
		panic(err)
	}

	span := opentracing.StartSpan("test-span")
	time.Sleep(time.Second)
	fmt.Println("ok.")
	span.Finish()
	// 底层reporter不会立即发送数据给broker, 给zipkin等待足够的时间发送span
	time.Sleep(time.Second * 5)
}