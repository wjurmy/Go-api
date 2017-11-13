package common

func StartUp() {

	//Initialize AppConfig variables
	initConfig()

	// Initialize public/private auth keys for JWT Authentication
	initKeys()

	// Start MySql database session
	createDbSession()
}

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}
