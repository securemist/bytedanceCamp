/********************************************************************************
* @author: Yakult
* @date: 2023/8/4 20:57
* @description:
********************************************************************************/

package util

import (
	"bytedanceCamp/config"
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

// Register 服务注册
func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", config.GetConfig().Consul.Host, config.GetConfig().Consul.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	// 生成对应grpc的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	return err
}

func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", config.GetConfig().Consul.Host, config.GetConfig().Consul.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}
