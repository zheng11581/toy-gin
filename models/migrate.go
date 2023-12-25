package models

func init() {
	DB.AutoMigrate(IngMonitorConf{})
	DB.AutoMigrate(IngMonitorRule{})
	DB.AutoMigrate(IngMonitorSpecialRule{})
	DB.AutoMigrate(IngMonitorSilence{})
}
