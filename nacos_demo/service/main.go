package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/util"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	// server config
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			"10.211.55.4",
			8848,
			constant.WithContextPath("/nacos")),
	}

	// client config
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("dev"),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	// Register service instance
	//registerServiceInstance(client, vo.RegisterInstanceParam{
	//	Ip:          "127.0.0.1",
	//	Port:        8848,
	//	ServiceName: "demo.go",
	//	GroupName:   "group-a",
	//	ClusterName: "cluster-a",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   true,
	//	Metadata:    map[string]string{"idc": "shanghai"},
	//})

	//DeRegister
	deRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8848,
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Cluster:     "cluster-a",
		Ephemeral:   true, //it must be true
	})

	//Get service with serviceName, groupName , clusters
	//getService(client, vo.GetServiceParam{
	//	ServiceName: "demo.go",
	//	GroupName:   "group-a",
	//	Clusters:    []string{"cluster-a"},
	//})

	//SelectAllInstance
	//GroupName=DEFAULT_GROUP
	//selectAllInstances(client, vo.SelectAllInstancesParam{
	//	ServiceName: "demo.go",
	//	GroupName:   "group-a",
	//	Clusters:    []string{"cluster-a"},
	//})

	//SelectInstances only return the instances of healthy=${HealthyOnly},enable=true and weight>0
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	selectInstances(client, vo.SelectInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Clusters:    []string{"cluster-a"},
		HealthyOnly: true,
	})

	//Subscribe key=serviceName+groupName+cluster
	//Note:We call add multiple SubscribeCallback with the same key.
	subscribeParam := &vo.SubscribeParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		SubscribeCallback: func(services []model.Instance, err error) {
			fmt.Printf("callback return services:%s \n\n", util.ToJsonString(services))
		},
	}
	subscribe(client, subscribeParam)

	//wait for client pull change from server
	//time.Sleep(3 * time.Second)
	//
	//updateServiceInstance(client, vo.UpdateInstanceParam{
	//	Ip:          "127.0.0.1",
	//	Port:        8848,
	//	ServiceName: "demo.go",
	//	GroupName:   "group-a",
	//	ClusterName: "cluster-a",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   true,
	//	Metadata:    map[string]string{"idc": "beijing1"}, //update metadata
	//})

	//selectInstances(client, vo.SelectInstancesParam{
	//	ServiceName: "demo.go",
	//	GroupName:   "group-a",
	//	Clusters:    []string{"cluster-a"},
	//})

}
