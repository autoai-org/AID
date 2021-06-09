package utilities

// InitFolders create all required folders under ~/.autoai/.aid
func InitConfigs() {
	initConfig := SystemConfig{RemoteReport: true}
	SaveConfig(initConfig)
}

func init() {
	InitConfigs()
}
