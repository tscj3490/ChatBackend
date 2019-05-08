package config

/////////////////////////////
/*    Mysql Production     */
/////////////////////////////

// MysqlHost ...
var MysqlHost = "localhost"

// MysqlPort ...
var MysqlPort = "3306"

// MysqlUser ...
var MysqlUser = "root" //root //skyclean

// MysqlPassword ...
var MysqlPassword = "Terry123!" //root // 7iD6u%0g

// MysqlDatabase ...
var MysqlDatabase = "chat"

// Constants for database.
const (
	MysqlProtocol = "tcp"
	// for DEVELOPMENT
	MysqlHostDev     = "localhost"
	MysqlPortDev     = "3306"
	MysqlUserDev     = "root"      //root  //skyclean
	MysqlPasswordDev = "Terry123!" //root  //7iD6u%0g
	MysqlDatabaseDev = "chat"
	MysqlOptionsDev  = "charset=utf8&parseTime=True"
	// for PRODUCTION
	MysqlOptions = "charset=utf8&parseTime=True"
)

// MysqlDSL return mysql DSL.
func MysqlDSL() string {
	var mysqlDSL string
	switch Environment {
	case "DEVELOPMENT":
		mysqlDSL = MysqlUserDev + ":" + MysqlPasswordDev + "@" + MysqlProtocol + "(" + MysqlHostDev + ":" + MysqlPortDev + ")/" + MysqlDatabaseDev + "?" + MysqlOptionsDev
	default:
		mysqlDSL = MysqlUser + ":" + MysqlPassword + "@" + MysqlProtocol + "(" + MysqlHost + ":" + MysqlPort + ")/" + MysqlDatabase + "?" + MysqlOptions
	}
	return mysqlDSL
}
