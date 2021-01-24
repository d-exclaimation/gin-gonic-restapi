package database

func Handle(err error) {
	// If error exist, panic to the logs and stop the process
	if err != nil {
		panic(err)
	}
}