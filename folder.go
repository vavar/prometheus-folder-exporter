package main

import (
	"io/ioutil"
	"github.com/prometheus/client_golang/prometheus"
)

type FolderStatus float64

const (
	Inactive 	FolderStatus = 0
	ReadError	FolderStatus = 1
	Active 		FolderStatus = 2
)

type FolderMetric struct {
	Path string
}

func (fm *FolderMetric) Desc() *prometheus.Desc {
	return prometheus.NewDesc("folder_exists_metric", "",[]string{}, prometheus.Labels{ "path": fm.Path })
}

func (fm *FolderMetric) Collect(ch *chan<- prometheus.Metric) {
	_, err := ioutil.ReadDir(fm.Path)
	status := Inactive
	switch {
		case err != nil:
			status = ReadError
		default:
			status = Active
	}
	*ch <- prometheus.MustNewConstMetric(fm.Desc(), prometheus.GaugeValue, float64(status))
}

type FolderCollector struct {
	Metrics []FolderMetric
}

func (f *FolderCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range f.Metrics {
		ch <- metric.Desc()
	}
}

func (f *FolderCollector) Collect(ch chan<- prometheus.Metric) {
	for _, metric := range f.Metrics {
		metric.Collect(&ch)
	}
}

func NewFolderCollector(targets []string) *FolderCollector {
	metrics := []FolderMetric{}
	for _, path := range targets {
		metrics = append(metrics, FolderMetric{ Path: path})
	}
	return &FolderCollector{ metrics }
}