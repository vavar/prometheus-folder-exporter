package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v3"
)

// Config ...
type Config struct {
	Port    string   `yml:"port"`
	Targets []string `yml:"targets"`
}

func main() {

	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
		return
	}

	config := Config{}
	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Start http server : %s ...\n", config.Port)
	prometheus.MustRegister(NewFolderCollector(config.Targets))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
