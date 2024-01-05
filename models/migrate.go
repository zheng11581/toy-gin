package models

func init() {
	DB.AutoMigrate(Conf{})
	DB.AutoMigrate(Rule{})
	DB.AutoMigrate(SpecialRule{})
	DB.AutoMigrate(SilenceRule{})
	DB.AutoMigrate(App{})
}
