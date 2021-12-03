module github.com/DendraScience/dendra_hummingbird_monitor

go 1.17

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
	google.golang.org/grpc v1.33.2 // indirect
)

require (
	cloud.google.com/go v0.51.0 // indirect
	github.com/Microsoft/go-winio v0.4.17 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/containerd/containerd v1.5.7 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/jaypipes/pcidb v0.6.0 // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.opencensus.io v0.22.3 // indirect
	golang.org/x/exp v0.0.0-20191227195350-da58074b4299 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	golang.org/x/tools v0.1.1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/api v0.15.0 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	honnef.co/go/tools v0.0.1-2019.2.3 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
)
