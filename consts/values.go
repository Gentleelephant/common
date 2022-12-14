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

const (
	LoggerMessageKey        = "logger.messageKey"
	LoggerLevelKey          = "logger.levelKey"
	LoggerTimeKey           = "logger.timeKey"
	LoggerCallerKey         = "logger.callerKey"
	LoggerNameKey           = "logger.nameKey"
	LoggerFunctionKey       = "logger.functionKey"
	LoggerStacktraceKey     = "logger.stacktraceKey"
	LoggerLineEnding        = "logger.lineEnding"
	LoggerSeparator         = "logger.separator"
	LoggerEncoding          = "logger.encoding" //json or console
	LoggerDevelopmentMode   = "logger.developmentMode"
	LoggerDisableCaller     = "logger.disableCaller"
	LoggerDisableStacktrace = "logger.disableStacktrace"
	LoggerOutputPath        = "logger.outputPath"
	LoggerErrorOutputPath   = "logger.errorOutputPath"
	LoggerMaxSize           = "logger.maxSize"
	LoggerMaxAge            = "logger.maxAge"
	LoggerMaxBackups        = "logger.maxBackups"
	LoggerLevel             = "logger.level"
)
