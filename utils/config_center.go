package utils

import (
	"fmt"
	"github.com/Gentleelephant/common/consts"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"strings"
)

type NacosConfigparams struct {
	NacosNamespace   string `mapstructure:"nacos.namespace"`
	NacosHost        string `mapstructure:"nacos.host"`
	NacosLogDir      string `mapstructure:"nacos.logDir"`
	NacosCacheDir    string `mapstructure:"nacos.cacheDir"`
	NacosLogLevel    string `mapstructure:"nacos.logLevel"`
	NacosPort        uint64 `mapstructure:"nacos.port"`
	NacosDataId      string `mapstructure:"nacos.dataId"`
	NacosGroup       string `mapstructure:"nacos.group"`
	NacosContextPath string `mapstructure:"nacos.contextPath"`
	NacosScheme      string `mapstructure:"nacos.scheme"`
	NacosTimeout     uint64 `mapstructure:"nacos.timeout"`
}

// InitRemoteConfig 初始化远程配置
func InitRemoteConfig(config NacosConfigparams) (*viper.Viper, error) {
	vl := viper.New()
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.NacosNamespace, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           config.NacosTimeout,
		NotLoadCacheAtStart: true,
		LogDir:              config.NacosLogDir,
		CacheDir:            config.NacosCacheDir,
		LogLevel:            config.NacosLogLevel,
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      config.NacosHost,
			ContextPath: config.NacosContextPath,
			Port:        config.NacosPort,
			Scheme:      config.NacosScheme,
		},
	}
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		return nil, err
	}
	// 获取配置
	remoteConfig, err := client.GetConfig(vo.ConfigParam{
		DataId: config.NacosDataId,
		Group:  config.NacosGroup,
	})
	if err != nil {
		return nil, err
	}
	// 解析配置
	parseConfig(vl, remoteConfig)
	err = client.ListenConfig(vo.ConfigParam{
		DataId: config.NacosDataId,
		Group:  config.NacosGroup,
		OnChange: func(namespace, group, dataId, data string) {
			// 刷新配置
			parseConfig(vl, data)
		},
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(remoteConfig)
	return vl, nil
}

func parseConfig(viper *viper.Viper, data string) {
	// 解析配置
	reader := strings.NewReader(data)
	viper.SetConfigType(consts.ConfigFileType)
	err := viper.ReadConfig(reader)
	if err != nil {
		panic(err)
	}
}

// GetConfig 根据配置文件获得配置
func GetConfig(filePah string, fileType string) (*viper.Viper, error) {
	vl := viper.New()
	vl.SetConfigFile(filePah)
	vl.SetConfigType(fileType)
	err := vl.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return vl, nil
}
