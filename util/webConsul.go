/********************************************************************************
* @author: Yakult
* @date: 2023/8/7 16:02
* @description:
********************************************************************************/

package util

import (
	"bytedanceCamp/dao/global"
	"fmt"
	"github.com/hashicorp/consul/api"
)

type WebRegistry struct {
	Host string
	Port int
}

type RegistryWebClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistryWebClient(host string, port int) RegistryWebClient {
	return &WebRegistry{
		Host: host,
		Port: port,
	}
}

// Register 服务注册
func (r *WebRegistry) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ProjectConfig.Consul.Host, global.ProjectConfig.Consul.Port)

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
	// 生成http对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	return err
}

func (r *WebRegistry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ProjectConfig.Consul.Host, global.ProjectConfig.Consul.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}
