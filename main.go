package main

import (
	"fmt"
	"sync"
	"time"

	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func pushToPrometheus(status float64) {
	registry := prometheus.NewRegistry()

	labels := prometheus.Labels{
		"env":      "prod",
		"product":  "xxx",
		"project":  "grafana-server",
		"instance": "grafana.xxx.com",
	}

	gaugeOpts := prometheus.GaugeOpts{
		Name: "backup_status",
		Help: "Backup status",
	}

	backupStatus := prometheus.NewGaugeVec(
		gaugeOpts,
		[]string{"env", "product", "project", "instance"},
	)

	specificBackupStatus := backupStatus.WithLabelValues(labels["env"], labels["product"], labels["project"], labels["instance"])
	specificBackupStatus.Set(status)

	registry.MustRegister(backupStatus)

	pushGateway := push.New("http://pushgateway.xxx.com", "grafana_backup").Grouping("env", "prod").Grouping("instance", "grafana.xxx.com").Grouping("project", "grafana-server")

	err := pushGateway.Push()
	if err != nil {
		fmt.Println("Failed to push metrics to Prometheus Pushgateway:", err)
		return
	}
}

func main() {

	now := time.Now()
	date := now.Format("2006/01/02")

	const (
		ak       = ""
		sk       = ""
		endPoint = "https://obs.cn-east-3.myhuaweicloud.com"
	)

	var once sync.Once

	obsClient, err := obs.New(ak, sk, endPoint)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
	}

	once.Do(
		func() {
			input := &obs.PutFileInput{}
			input.Bucket = "sh1-pub-xxxx-backup"
			input.Key = fmt.Sprintf("grafana-backup/%s/grafana.db", date)
			input.SourceFile = "/var/lib/grafana/grafana.db"

			output, err := obsClient.PutFile(input)
			if err == nil {
				fmt.Printf("Put file(%s) under the bucket(%s) successful!\n", input.Key, input.Bucket)
				fmt.Printf("StorageClass:%s, ETag:%s\n", output.StorageClass, output.ETag)
				pushToPrometheus(1)
				pushToPrometheus(float64(1))
				return
			} else {
				pushToPrometheus(float64(0))
			}
		},
	)
}
