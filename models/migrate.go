package models

func init() {
	DB.AutoMigrate(IngMonitorConf{})
}
