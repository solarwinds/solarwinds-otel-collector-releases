// Copyright 2025 SolarWinds Worldwide, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The goal of this connection-check is to be able to perform GRPC endpoint check.
// This feature is mainly intended for kubernetes use
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	otellog "go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func sendTestMessage(logger *zap.Logger, endpoint, apiToken, clusterUid string, insecure bool) {
	ctx := context.Background()
	otel.SetErrorHandler(&OtelErrorHandler{logger: logger})

	exporterOptions := []otlploggrpc.Option{
		otlploggrpc.WithEndpoint(endpoint),
		otlploggrpc.WithHeaders(map[string]string{"Authorization": "Bearer " + apiToken}),
		otlploggrpc.WithCompressor("gzip"),
	}

	if insecure {
		exporterOptions = append(exporterOptions, otlploggrpc.WithInsecure())
	}

	exporter, err := otlploggrpc.New(ctx, exporterOptions...)
	if err != nil {
		logger.Fatal("Failed to create log exporter", zap.Error(err))
	}

	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewSimpleProcessor(exporter)),
		sdklog.WithResource(resource.NewWithAttributes("", attribute.String("sw.k8s.cluster.uid", clusterUid))),
	)
	defer loggerProvider.Shutdown(ctx)

	otelLogger := loggerProvider.Logger("connection-check")

	record := otellog.Record{}
	record.SetSeverityText("INFO")
	record.SetBody(otellog.StringValue("otel-endpoint-check successful"))
	record.SetTimestamp(time.Now())

	otelLogger.Emit(ctx, record)
	logger.Info("Connection check was successful")
}

type OtelErrorHandler struct {
	logger *zap.Logger
}

func (d *OtelErrorHandler) Handle(err error) {
	switch status.Code(err) {
	case codes.Unauthenticated:
		d.logger.Fatal("ERROR: A valid token is not set", zap.Error(err))
	case codes.Unavailable:
		d.logger.Fatal("ERROR: The target endpoint is not available", zap.Error(err))
	default:
		d.logger.Fatal("ERROR", zap.Error(err))
	}
}

func main() {
	var clusterUid, endpoint, apiToken string
	var insecure bool

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	flag.StringVar(&clusterUid, "clusteruid", "", "Cluster UID")
	flag.StringVar(&endpoint, "endpoint", "", "Endpoint")
	flag.StringVar(&apiToken, "apitoken", "", "Api token")
	flag.BoolVar(&insecure, "insecure", false, "Insecure")

	flag.Parse()
	sendTestMessage(logger, endpoint, apiToken, clusterUid, insecure)
}
