package e

// TESTME
// short hand for NewErrorPlus	
func E(err error, message string, tags []string, fn interface{}, args ...interface{}) *ErrorPlus {
	return NewErrorPlus(err, message, tags, fn, args...)
}
