package consts

const (
	ServiceDynamicPort = "service.dynamic.port"
	ConfigFileType     = "yaml"
)

// Nacos
const (
	NacosIpAddr = "nacos.IpAddr"
	NacosPort   = "nacos.port"
	NacosGroup  = "nacos.group"
	NacosDataId = "nacos.dataId"

	//  client config
	NacosTimeoutMs            = "nacos.timeoutMs"
	NacosBeatInterval         = "nacos.beatInterval"
	NacosNamespaceId          = "nacos.namespaceId"
	NacosAppName              = "nacos.appName"
	NacosEndPoint             = "nacos.endPoint"
	NacosRegionId             = "nacos.regionId"
	NacosAccessKey            = "nacos.accessKey"
	NacosSecretKey            = "nacos.secretKey"
	NacosOpenKMS              = "nacos.openKMS"
	NacosCacheDir             = "nacos.cacheDir"
	NacosUpdateThreadNum      = "nacos.updateThreadNum"
	NacosNotLoadCacheAtStart  = "nacos.notLoadCacheAtStart"
	NacosUpdateCacheWhenEmpty = "nacos.updateCacheWhenEmpty"
	NacosUsername             = "nacos.username"
	NacosPassword             = "nacos.password"
	NacosLogDir               = "nacos.logDir"
	NacosLogLevel             = "nacos.logLevel"
	NacosContextPath          = "nacos.contextPath"
	NacosScheme               = "nacos.scheme"
)

// Sql
const (
	SqlHost     = "sql.host"
	SqlPort     = "sql.port"
	SqlUsername = "sql.username"
	SqlPassword = "sql.password"
	SqlDatabase = "sql.database"
)

// Redis
const (
	RedisHost     = "redis.host"
	RedisPort     = "redis.port"
	RedisPassword = "redis.password"
	RedisDB       = "redis.db"
)
