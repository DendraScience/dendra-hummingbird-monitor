module github.com/DendraScience/dendra_hummingbird_monitor

go 1.18

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
	cloud.google.com/go/bigquery v1.46.0
	github.com/BurntSushi/toml v1.1.0
	github.com/docker/docker v23.0.1+incompatible
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/jaypipes/ghw v0.8.0
	github.com/moby/moby v23.0.1+incompatible
	github.com/prometheus/procfs v0.7.3
	github.com/sirupsen/logrus v1.9.0
	golang.org/x/sys v0.5.0
	k8s.io/api v0.26.1
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
)

require (
	cloud.google.com/go v0.107.0 // indirect
	cloud.google.com/go/compute v1.18.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v0.12.0 // indirect
	github.com/Microsoft/go-winio v0.4.17 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/apache/arrow/go/v10 v10.0.1 // indirect
	github.com/apache/thrift v0.18.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v23.1.21+incompatible // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/jaypipes/pcidb v0.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	golang.org/x/tools v0.1.12 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.110.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230216225411-c8e22ba71e44 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.0.3 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220328201542-3ee0da9b0b42 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
