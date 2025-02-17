module github.com/solarwinds/solarwinds-otel-collector/exporter/solarwindsexporter

go 1.23.4

require (
	github.com/solarwinds/solarwinds-otel-collector/extension/solarwindsextension v0.113.8
	github.com/solarwinds/solarwinds-otel-collector/pkg/testutil v0.113.8
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.119.0
	go.opentelemetry.io/collector/component/componenttest v0.119.0
	go.opentelemetry.io/collector/config/configgrpc v0.119.0
	go.opentelemetry.io/collector/config/configopaque v1.25.0
	go.opentelemetry.io/collector/config/configretry v1.25.0
	go.opentelemetry.io/collector/config/configtls v1.25.0
	go.opentelemetry.io/collector/confmap v1.25.0
	go.opentelemetry.io/collector/exporter v0.119.0
	go.opentelemetry.io/collector/exporter/exportertest v0.119.0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.119.0
	go.opentelemetry.io/collector/pdata v1.25.0
	go.uber.org/goleak v1.3.0
)

require (
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/solarwinds/solarwinds-otel-collector/pkg/version v0.113.8 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/client v1.25.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.119.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.25.0 // indirect
	go.opentelemetry.io/collector/config/confignet v1.25.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.119.0 // indirect
	go.opentelemetry.io/collector/consumer v1.25.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.119.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror/xconsumererror v0.119.0 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.119.0 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.119.0 // indirect
	go.opentelemetry.io/collector/exporter/exporterhelper/xexporterhelper v0.119.0 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.119.0 // indirect
	go.opentelemetry.io/collector/extension v0.119.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.119.0 // indirect
	go.opentelemetry.io/collector/extension/xextension v0.119.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.25.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.119.0 // indirect
	go.opentelemetry.io/collector/pipeline v0.119.0 // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.119.0 // indirect
	go.opentelemetry.io/collector/receiver v0.119.0 // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.119.0 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.119.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.59.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/solarwinds/solarwinds-otel-collector/extension/solarwindsextension => ../../extension/solarwindsextension

replace github.com/solarwinds/solarwinds-otel-collector/pkg/testutil => ../../pkg/testutil

replace github.com/solarwinds/solarwinds-otel-collector/pkg/version => ../../pkg/version
