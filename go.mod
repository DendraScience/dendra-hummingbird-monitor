module github.com/DendraScience/dendra_hummingbird_monitor

go 1.21.0

replace k8s.io/client-go => k8s.io/client-go v0.24.2

//replace k8s.io/metrics => k8s.io/metrics v0.24.2

replace k8s.io/api => k8s.io/api v0.24.2

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.24.2

replace k8s.io/apiserver => k8s.io/apiserver v0.24.2

replace k8s.io/apimachinery => k8s.io/apimachinery v0.24.2

replace github.com/jaypipes/ghw => github.com/taigrr/ghw v0.8.0-1

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.24.2

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.24.2

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.24.2

replace k8s.io/code-generator => k8s.io/code-generator v0.24.2

replace k8s.io/component-base => k8s.io/component-base v0.24.2

replace k8s.io/component-helpers => k8s.io/component-helpers v0.24.2

replace k8s.io/controller-manager => k8s.io/controller-manager v0.24.2

replace k8s.io/cri-api => k8s.io/cri-api v0.24.2

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.24.2

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.24.2

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.24.2

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.24.2

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.24.2

replace k8s.io/kubectl => k8s.io/kubectl v0.24.2

replace k8s.io/kubelet => k8s.io/kubelet v0.24.2

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.24.2

replace k8s.io/metrics => k8s.io/metrics v0.24.2

replace k8s.io/mount-utils => k8s.io/mount-utils v0.24.2

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.24.2

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.24.2

require (
	cloud.google.com/go/bigquery v1.61.0
	github.com/BurntSushi/toml v1.1.0
	github.com/docker/docker v27.0.0+incompatible
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/jaypipes/ghw v0.8.0
	github.com/moby/moby v27.0.0+incompatible
	github.com/prometheus/procfs v0.7.3
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/sys v0.21.0
	k8s.io/api v0.30.2
	k8s.io/apimachinery v0.30.2
	k8s.io/client-go v1.5.2
)

require (
	cloud.google.com/go v0.115.0 // indirect
	cloud.google.com/go/auth v0.5.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/compute v1.27.0 // indirect
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	cloud.google.com/go/iam v1.1.8 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/apache/arrow/go/v10 v10.0.1 // indirect
	github.com/apache/arrow/go/v15 v15.0.2 // indirect
	github.com/apache/thrift v0.20.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/emicklei/go-restful/v3 v3.12.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/google/gnostic v0.7.0 // indirect
	github.com/google/gnostic-models v0.6.9-0.20230804172637-c7be7c783f49 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.4 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/jaypipes/pcidb v0.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.52.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240604190554-fc45aab8b7f8 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/term v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.184.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240610135401-a8a62080eff3 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240610135401-a8a62080eff3 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240610135401-a8a62080eff3 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.0.3 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240521193020-835d969ad83a // indirect
	k8s.io/utils v0.0.0-20240502163921-fe8a2dddb1d0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
