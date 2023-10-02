package e

func E(err error, message string, tags []string, fn interface{}, args ...interface{}) *ErrorPlus {
	return NewErrorPlus(err, message, tags, fn, args...)
}
