package global

/*
 * 需要和字典匹配
 */

const (
	RouteRootPath = "/api"
	ModelName     = "challenge"
	LoginLog      = "login_log_queue"
	OperateLog    = "operate_log_queue"
	TrafficKey    = "X-Request-Id"
	LoggerKey     = "_challenge-logger-request"

	// SysStatusOk 通用-正常
	SysStatusOk    = "1"
	SysStatusNotOk = "2"
)

const (
	DBDriverMysql    = "mysql"
	DBDriverPostgres = "postgres"
)

const (
	ModeDev  string = "dev"  //开发模式
	ModeTest string = "test" //测试模式
	ModeProd string = "prod" //生产模式
)
