package main

import (
	"os"
	"io/ioutil"
	"github.com/prometheus/client_golang/prometheus"
)
//FolderStatus ...
type FolderStatus float64

const (
	//ReadError ...
	ReadError	FolderStatus = 1
	//Active ...
	Active 		FolderStatus = 2
)

//FolderMetric ...
type FolderMetric struct {
	Path string
	Reader FolderReader
}

//FolderReader ...
type FolderReader func (path string) ([]os.FileInfo, error)

//Desc ... 
func (fm *FolderMetric) Desc() *prometheus.Desc {
	return prometheus.NewDesc("folder_exists_metric", "",[]string{}, prometheus.Labels{ "path": fm.Path })
}

//Collect ...
func (fm *FolderMetric) Collect(ch chan<- prometheus.Metric) {
	_, err := fm.Reader(fm.Path)
	var status FolderStatus
	switch {
		case err != nil:
			status = ReadError
		default:
			status = Active
	}
	ch <- prometheus.MustNewConstMetric(fm.Desc(), prometheus.GaugeValue, float64(status))
}

//FolderCollector ...
type FolderCollector struct {
	Metrics []FolderMetric
}

//Describe ...
func (f *FolderCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range f.Metrics {
		ch <- metric.Desc()
	}
}

//Collect ...
func (f *FolderCollector) Collect(ch chan<- prometheus.Metric) {
	for _, metric := range f.Metrics {
		metric.Collect(ch)
	}
}

//NewFolderCollector ...
func NewFolderCollector(targets []string) *FolderCollector {
	metrics := []FolderMetric{}
	for _, path := range targets {
		metrics = append(metrics, FolderMetric{Path: path, Reader: ioutil.ReadDir})
	}
	return &FolderCollector{ metrics }
}