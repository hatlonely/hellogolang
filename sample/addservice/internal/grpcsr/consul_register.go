package grpcsr

import (
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
)

// NewConsulRegister create a new consul register
func NewConsulRegister() *ConsulRegister {
	return &ConsulRegister{
		Address: "127.0.0.1:8500",
		Service: "addservice",
		Tag:     []string{"hatlonely"},
		Port:    3000,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
}

// ConsulRegister consul service register
type ConsulRegister struct {
	Address                        string
	Service                        string
	Tag                            []string
	Port                           int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

// Register register service
func (r *ConsulRegister) Register() error {
	config := api.DefaultConfig()
	config.Address = r.Address
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	agent := client.Agent()

	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v", r.Service, r.Port),
		Name:    fmt.Sprintf("grpc.health.v1.%v", r.Service),
		Tags:    r.Tag,
		Port:    r.Port,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			Interval: r.Interval.String(),
			GRPC:     fmt.Sprintf("127.0.0.1:%v/%v", r.Port, r.Service),
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}
