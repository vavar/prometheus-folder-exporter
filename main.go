package main

import (
	"fmt"
    "io/ioutil"
	"net/http"
	"gopkg.in/yaml.v2"
	"github.com/prometheus/common/log"
  	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Config ...
type Config struct {
	Port string 		`yml:"port"`
	Targets []string 	`yml:"targets"`
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
