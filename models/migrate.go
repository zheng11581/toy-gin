package models

func init() {
	DB.AutoMigrate(Conf{})
	DB.AutoMigrate(Rule{})
	DB.AutoMigrate(Special{})
	DB.AutoMigrate(Silence{})
	DB.AutoMigrate(App{})
}
