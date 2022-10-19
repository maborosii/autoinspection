package prom

import (
	"context"
	"fmt"
	"node_metrics_go/global"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"go.uber.org/zap"
)

func ClientForProm(address string) v1.API {
	client, err := api.NewClient(api.Config{
		Address: address,
	})
	if err != nil {
		global.Logger.Fatal("Error creating client: ", zap.Error(err))
	}
	v1api := v1.NewAPI(client)
	return v1api
}

func QueryFromPromDemo(promql string, api v1.API) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(promql)
	result, _, _ := api.Query(ctx, promql, time.Now())

	fmt.Printf("%v\n", result.String())
}
