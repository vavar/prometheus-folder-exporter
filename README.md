# Prometheus Folder Exporter

Folder exporter for legacy system
Use for detect folder status to ensure other services can process regularly


## Build

```
go build
```

## Run

```
go run .
```

## Dependencies
```
    "gopkg.in/yaml.v2"
	"github.com/prometheus/common/log"
  	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
```