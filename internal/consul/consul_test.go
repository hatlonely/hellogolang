package consul

import (
	"net"
	"strconv"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestConsul(t *testing.T) {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	{
		agent := client.Agent()
		err := agent.ServiceRegister(&api.AgentServiceRegistration{
			ID:      "add2",
			Name:    "add2",
			Tags:    []string{"hatlonely"},
			Port:    3001,
			Address: "127.0.0.1",
		})
		if err != nil {
			t.Error(err)
		}
	}

	{
		kv := client.KV()
		_, err := kv.Put(&api.KVPair{
			Key:   "name",
			Value: []byte("hatlonely"),
		}, nil)
		if err != nil {
			t.Error(err)
		}

		pair, _, err := kv.Get("name", nil)
		if err != nil {
			t.Error(err)
		}
		t.Log(string(pair.Value))
	}

	{
		health := client.Health()
		check, _, err := health.Checks("add", nil)
		if err != nil {
			t.Error(err)
		}
		t.Log(check.AggregatedStatus())

		services, _, err := health.Service("grpc.health.v1.addservice", "", true, &api.QueryOptions{
			WaitIndex: 0,
		})
		if err != nil {
			t.Error(err)
		}
		for _, service := range services {
			t.Error(net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port)))
		}
	}
}
