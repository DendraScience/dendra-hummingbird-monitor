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
	cloud.google.com/go/bigquery v1.32.0
	github.com/BurntSushi/toml v1.1.0
	github.com/docker/docker v20.10.17+incompatible
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/jaypipes/ghw v0.8.0
	github.com/moby/moby v20.10.17+incompatible
	github.com/prometheus/procfs v0.7.3
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/sys v0.0.0-20220615213510-4f61da869c0c
	k8s.io/kubernetes v1.24.2
)

require (
	cloud.google.com/go/compute v1.6.1 // indirect
	cloud.google.com/go/iam v0.3.0 // indirect
	github.com/JeffAshton/win_pdh v0.0.0-20161109143554-76bb4ee9f0ab // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/aws/aws-sdk-go v1.38.49 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.3.0 // indirect
	github.com/cilium/ebpf v0.7.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/containerd/containerd v1.4.12 // indirect
	github.com/containerd/ttrpc v1.0.2 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/godbus/dbus/v5 v5.0.6 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/cadvisor v0.44.1 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/jaypipes/pcidb v0.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mindprince/gonvml v0.0.0-20190828220739-9ebdce4bb989 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/moby/sys/mountinfo v0.6.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mrunalp/fileutils v0.5.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/runc v1.1.1 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417 // indirect
	github.com/opencontainers/selinux v1.10.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/seccomp/libseccomp-golang v0.9.2-0.20210429002308-3879420cc921 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/vishvananda/netlink v1.1.0 // indirect
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/oauth2 v0.0.0-20220608161450-d0670ef3b1eb // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
	k8s.io/api v0.24.2 // indirect
	k8s.io/apiserver v0.24.2 // indirect
	k8s.io/component-base v0.24.2 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220328201542-3ee0da9b0b42 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

require (
	cloud.google.com/go v0.102.0 // indirect
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20211202183452-c5a74bcca799 // indirect
	golang.org/x/net v0.0.0-20220615171555-694bf12d69de // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/api v0.84.0 // indirect
	google.golang.org/genproto v0.0.0-20220615141314-f1464d18c36b // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
	k8s.io/metrics v0.24.2
)
