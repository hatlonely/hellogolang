package grpcsr

import (
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
)

// NewConsulRegister create a new consul register
func NewConsulRegister() *ConsulRegister {
	return &ConsulRegister{
		Address: "127.0.0.1:8500",
		Service: "unknown",
		Tag:     []string{},
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

	IP := localIP()
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.Service, IP, r.Port),
		Name:    fmt.Sprintf("grpc.health.v1.%v", r.Service),
		Tags:    r.Tag,
		Port:    r.Port,
		Address: IP,
		Check: &api.AgentServiceCheck{
			Interval: r.Interval.String(),
			GRPC:     fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Service),
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
