# Prometheus Folder Exporter

Folder exporter for legacy system.
Used for detect abnormal folder status to ensure other services can process regularly


## Example metrics
```
# HELP folder_exists_metric 
# TYPE folder_exists_metric gauge
folder_exists_metric{path="C:"} 2
folder_exists_metric{path="C:/Window"} 2
folder_exists_metric{path="N:"} 1
```
* 2 = Active
* 1 = Read Error
* 0 = Inactive


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
