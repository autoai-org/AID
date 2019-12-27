package utilities

// CheckError checks if error is not nil, and if it is nil it will logger the information
func CheckError(err error, errorMessage string) {
	if err != nil {
		logger.Error(errorMessage)
		logger.Fatal(err.Error())
	}
}
