package e

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

//TESTME
// feeded with the fn to return the trace and timestamp
func TimeStampAndRuntimeFN(fn interface{}) string {
	t := time.Now()
	timestamp := t.Format("2006-01-02 15:04:05")
	s:= runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return fmt.Sprintf("trace: %s time: %s", s, timestamp)
}
//TESTME
// ERROR: feeded with the  error and fn to return a error with the trace and timestamp
func EE(err error, fn interface{})  error {
	s := TimeStampAndRuntimeFN(fn)
	return fmt.Errorf("error: %s  %s", err,s,)
}

//TESTME
// ERROR STRING:feeded with the  string error and fn to return a error with the trace and timestamp
func ES(err string, fn interface{})  error {
	s := TimeStampAndRuntimeFN(fn)
	return fmt.Errorf("error: %s  %s ", err,s)
}



