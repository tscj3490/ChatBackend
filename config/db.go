package config

/////////////////////////////
/*    Mysql Production     */
/////////////////////////////

// MysqlHost ...
var MysqlHost = "localhost"

// MysqlPort ...
var MysqlPort = "3306"

// MysqlUser ...
var MysqlUser = "chatuser" //root //chatuser

// MysqlPassword ...
var MysqlPassword = "Terry123!" //root // Terry123!

// MysqlDatabase ...
var MysqlDatabase = "chat"

// Constants for database.
const (
	MysqlProtocol = "tcp"
	// for DEVELOPMENT
	MysqlHostDev     = "localhost"
	MysqlPortDev     = "3306"
	MysqlUserDev     = "chatuser"  //root  //chatuser
	MysqlPasswordDev = "Terry123!" //root  //Terry123!
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
