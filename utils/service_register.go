package utils

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func getClientServerConfig(config NacosConfigparams) (constant.ClientConfig, []constant.ServerConfig) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              config.LogDir,
		CacheDir:            config.CacheDir,
		LogLevel:            config.LogLevel,
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      config.IpAddr,
			ContextPath: config.ServerConfig.ContextPath,
			Port:        config.Port,
			Scheme:      config.Scheme,
		},
	}
	return clientConfig, serverConfigs
}

// RegisterInstance 注册服务实例
func RegisterInstance(config NacosConfigparams, params vo.RegisterInstanceParam) (bool, error) {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return false, err
	}
	success, err := namingClient.RegisterInstance(params)
	if err != nil {
		return false, err
	}
	return success, nil
}

// DeregisterInstance 取消注册服务实例
func DeregisterInstance(config NacosConfigparams, params vo.DeregisterInstanceParam) (bool, error) {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return false, err
	}
	success, err := namingClient.DeregisterInstance(params)
	if err != nil {
		return false, err
	}
	return success, nil
}

// GetService GetService 获取服务实例
func GetService(config NacosConfigparams, params vo.GetServiceParam) (model.Service, error) {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return model.Service{}, err
	}
	service, err := namingClient.GetService(params)
	if err != nil {
		return model.Service{}, err
	}
	return service, nil
}

// SelectAllInstances 获取所有服务实例
func SelectAllInstances(config NacosConfigparams, params vo.SelectAllInstancesParam) ([]model.Instance, error) {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}
	services, err := namingClient.SelectAllInstances(params)
	if err != nil {
		return nil, err
	}
	return services, nil
}

// SelectServices 获取服务实例
func SelectServices(config NacosConfigparams, params vo.SelectInstancesParam) ([]model.Instance, error) {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}
	services, err := namingClient.SelectInstances(params)
	if err != nil {
		return nil, err
	}
	return services, nil
}

// SelectOneHealthyInstance 获取一个健康的服务实例
func SelectOneHealthyInstance(config NacosConfigparams, params vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return &model.Instance{}, err
	}
	instance, err := namingClient.SelectOneHealthyInstance(params)
	if err != nil {
		return &model.Instance{}, err
	}
	return instance, nil
}

// Subscribe 订阅服务
func Subscribe(config NacosConfigparams, params *vo.SubscribeParam) error {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil
	}
	err = namingClient.Subscribe(params)
	if err != nil {
		return err
	}
	return nil
}

// Unsubscribe 取消订阅服务
func Unsubscribe(config NacosConfigparams, params *vo.SubscribeParam) error {
	clientConfig, serverConfigs := getClientServerConfig(config)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil
	}
	err = namingClient.Unsubscribe(params)
	if err != nil {
		return err
	}
	return nil
}
