package main

import (
	"os"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
    "testing"
	"gotest.tools/assert"
)

func Test_FolderMetric_Collect(t *testing.T) {
	testCases := map[string]struct {
		err error
		expected float64
    }{
		"folder not found": {fmt.Errorf("error"), 1},
		"folder exists": {nil, 2},
    }

    for name, it := range testCases {
		t.Run(name, func(t *testing.T) {
			fm := FolderMetric{ Path: ".", Reader: func(f string)([]os.FileInfo, error) { return nil, it.err }  }
			ch := make(chan prometheus.Metric, 1)
			defer close(ch)
			fm.Collect(ch)
			found := <-ch
			metric := &dto.Metric{}
			found.Write(metric)
			assert.Equal(t, metric.GetGauge().GetValue(), it.expected)
        })
    }
}