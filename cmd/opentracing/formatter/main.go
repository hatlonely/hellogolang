package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	tracer, closer := initJaeger("formatter")
	defer closer.Close()

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		for k, v := range r.Header {
			fmt.Println(k, ":", v)
		}

		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		greeting := span.BaggageItem("greeting")
		if greeting == "" {
			greeting = "Hello"
		}

		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("%s, %s!", greeting, helloTo)
		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", helloStr),
		)
		w.Write([]byte(helloStr))
	})

	log.Fatal(http.ListenAndServe(":8083", nil))
}

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
