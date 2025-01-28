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

// The goal of this k8sconnectioncheck is be able to perform GRPC endpoint check.
// This feature is mainly intended for kubernetes use
package k8sconnectioncheck

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	otellog "go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/solarwinds/solarwinds-otel-collector/pkg/version"
)

func sendTestMessage(endpoint, apiToken, clusterUid string, insecure bool) {
	ctx := context.Background()
	otel.SetErrorHandler(new(OtelErrorHandler))

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
		log.Fatalf("ERROR: Failed to create log exporter\nDETAILS: %s", err)
	}

	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewSimpleProcessor(exporter)),
		sdklog.WithResource(resource.NewWithAttributes("", attribute.String("sw.k8s.cluster.uid", clusterUid))),
	)
	defer loggerProvider.Shutdown(ctx)

	logger := loggerProvider.Logger("solarwinds-otel-collector", otellog.WithInstrumentationVersion(version.Version))

	record := otellog.Record{}
	record.SetSeverityText("INFO")
	record.SetBody(otellog.StringValue("otel-endpoint-check successful"))
	record.SetTimestamp(time.Now())

	logger.Emit(ctx, record)
	log.Print("Connection check was successful")
}

type OtelErrorHandler struct{}

func (d *OtelErrorHandler) Handle(err error) {
	switch status.Code(err) {
	case codes.Unauthenticated:
		log.Fatalf("ERROR: A valid token is not set\nDETAILS: %s", err)
	case codes.Unavailable:
		log.Fatalf("ERROR: The target endpoint is not available\nDETAILS: %s", err)
	default:
		log.Fatalf("ERROR: %s", err)
	}
}

func NewCommand() *cobra.Command {
	var clusterUid, endpoint, apiToken string
	var insecure bool

	testCommand := &cobra.Command{
		Use:   "test-connection",
		Short: "Sends a single log to the provided endpoint",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			sendTestMessage(endpoint, apiToken, clusterUid, insecure)
		},
	}
	testCommand.Flags().StringVar(&clusterUid, "clusteruid", "", "")
	testCommand.Flags().StringVar(&endpoint, "endpoint", "", "")
	testCommand.Flags().StringVar(&apiToken, "apitoken", "", "")
	testCommand.Flags().BoolVar(&insecure, "insecure", false, "")
	testCommand.MarkFlagRequired("clusteruid")
	testCommand.MarkFlagRequired("endpoint")
	testCommand.MarkFlagRequired("apitoken")

	return testCommand
}
