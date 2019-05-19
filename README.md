# Prometheus Folder Exporter

Folder exporter for legacy system.
Used for detect abnormal folder status to ensure other services can use that folder properly

## Example Config
```
port: 8080
targets:
  - "C:"
  - "N:"
  - "C:/windows"
```


## Example metrics
```
# HELP folder_exists_metric 
# TYPE folder_exists_metric gauge
folder_exists_metric{path="C:"} 2
folder_exists_metric{path="C:/Windows"} 2
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
