package commons

func StartUp()  {
	initConfig()
	initKeys()
	createDbSession()
	addIndexesToDataBase()
}

