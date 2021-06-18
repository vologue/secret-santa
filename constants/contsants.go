package constants

import "os"

// Mysql Constants
var MysqlHost string = os.Getenv("MYSQL_HOST")
var MysqlPort string = os.Getenv("MYSQL_PORT")
var DbName string = os.Getenv("DB_NAME")
var MysqlUser string = os.Getenv("MYSQL_USERNAME")
var MysqlPassword string = os.Getenv("MYSQL_PASSWORD")

// Application constants
var Port string = os.Getenv("PORT")
