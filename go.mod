module github.com/DendraScience/dendra_hummingbird_monitor

go 1.16

replace k8s.io/client-go => k8s.io/client-go v0.19.2

replace k8s.io/api => k8s.io/api v0.19.2

replace google.golang.org/grpc => google.golang.org/grpc v1.23.1

replace github.com/jaypipes/ghw => github.com/taigrr/ghw v0.8.0-1

require (
	cloud.google.com/go/bigquery v1.0.1
	github.com/BurntSushi/toml v0.3.1
	github.com/docker/docker v20.10.10+incompatible
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/jaypipes/ghw v0.8.0
	github.com/moby/moby v20.10.10+incompatible
	github.com/prometheus/procfs v0.6.0
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881
)

require (
	github.com/containerd/containerd v1.5.7 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	golang.org/x/tools v0.1.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
