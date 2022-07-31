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
	//NacosNamespace    string `mapstructure:"nacos.namespace"`
	//NacosHost         string `mapstructure:"nacos.host"`
	//NacosLogDir       string `mapstructure:"nacos.logDir"`
	//NacosCacheDir     string `mapstructure:"nacos.cacheDir"`
	//NacosLogLevel     string `mapstructure:"nacos.logLevel"`
	//NacosPort         uint64 `mapstructure:"nacos.port"`
	DataId string `mapstructure:"nacos.dataId"`
	Group  string `mapstructure:"nacos.group"`
	//NacosContextPath  string `mapstructure:"nacos.contextPath"`
	//NacosScheme       string `mapstructure:"nacos.scheme"`
	//NacosTimeout      uint64 `mapstructure:"nacos.timeout"`
	//NacosBeatInterval int64  `mapstructure:"nacos.beatInterval"`
	constant.ClientConfig
	constant.ServerConfig
}

// InitRemoteConfig 初始化远程配置
func InitRemoteConfig(config NacosConfigparams) (*viper.Viper, error) {
	vl := viper.New()
	clientConfig := constant.ClientConfig{
		TimeoutMs:            config.TimeoutMs,
		BeatInterval:         config.BeatInterval,
		NamespaceId:          config.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		AppName:              config.AppName,
		Endpoint:             config.Endpoint,
		RegionId:             config.RegionId,
		AccessKey:            config.AccessKey,
		SecretKey:            config.SecretKey,
		OpenKMS:              config.OpenKMS,
		CacheDir:             config.CacheDir,
		UpdateThreadNum:      config.UpdateThreadNum,
		NotLoadCacheAtStart:  config.NotLoadCacheAtStart,
		UpdateCacheWhenEmpty: config.UpdateCacheWhenEmpty,
		Username:             config.Username,
		Password:             config.Password,
		LogDir:               config.LogDir,
		LogLevel:             config.LogLevel,
		LogSampling:          config.LogSampling,
		ContextPath:          config.ServerConfig.ContextPath,
		LogRollingConfig:     config.LogRollingConfig,
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
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		return nil, err
	}
	// 获取配置
	remoteConfig, err := client.GetConfig(vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
	})
	if err != nil {
		return nil, err
	}
	// 解析配置
	parseConfig(vl, remoteConfig)
	err = client.ListenConfig(vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
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
