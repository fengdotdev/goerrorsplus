package e

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// transitive version of error plus
type TErrorPlus struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
	Tags    []string `json:"tags"`
	FN      string `json:"fn"`
	Args    []interface{} `json:"args"`
	Trace   string `json:"trace"`
	Time    time.Time `json:"time"`
	RuntimeGoVer string `json:"runtimeGoVer"`
}


type ErrorPlus struct {
	runtimeGoVer string
	err     error  
	message string 
	tags    []string 
	fN      string 
	args    []interface{} 
	trace   string 
	time    time.Time 
}



// TESTME
func NewErrorPlus(err error, message string, tags []string, fn interface{}, args ...interface{}) *ErrorPlus {

	_,caller,line,_:= runtime.Caller(2)

	ep := &ErrorPlus{
		runtimeGoVer: runtime.Version(),
		err:     err,
		tags:    tags,
		message: message,
		fN:      runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
		args:    args,
		trace:   fmt.Sprintf("[caller: %s line: %d]",	caller, line ),
		time:    time.Now(),
	}
	return ep
}




// TESTME
// short hand for VerboseError
func (ep *ErrorPlus) V() error {
	return ep.VerboseError()
}
// TESTME
// get a verbose error with version, timestamp, message, error, tags, function name, trace, and args
func (ep *ErrorPlus) VerboseError() error {

	timestamp := ep.time.Format("2006-01-02 15:04:05")
	version:= ep.runtimeGoVer
	err := ep.err.Error()
	message := ep.message
	tags := ep.tags
	trace := ep.trace
	fn := ep.fN
	args := ep.args
	

	s:= fmt.Sprintf("go ver.: %s | at: %s | msg: %s | error: %s | tags: %s | fn: %s | trace: %s | args: %s",version, timestamp  ,err, message,tags, fn, trace, args)

// example: go ver.: go1.21.1 | at: 2023-10-03 16:13:19 | msg: recover: div by zero | error: div failed | tags: [math-err] | fn: main.division | trace: [caller: G:/fengdotdev/github/goerrorsplus/example/main.go line: 31] | args: [%!s(int=1) %!s(int=0)]
	return errors.New(s)

}


// data output methods

// TESTME
func (ep *ErrorPlus) Map() map[string]interface{} {
	return map[string]interface{}{
		"error":     ep.err,
		"message": ep.message,
		"tags":    ep.tags,
		"fn":      ep.fN,
		"args":    ep.args,
		"trace":   ep.trace,
		"time":    ep.time,
		"runtimeGoVer": ep.runtimeGoVer,
	}		
}


// TESTME
func (ep *ErrorPlus) String() string {
	return fmt.Sprintf("%s", ep.Map())
}
// TESTME
func (ep *ErrorPlus) JSON() string {
	return fmt.Sprintf("%s", ep.Map())
}
// TESTME
func (ep *ErrorPlus) JSONDATA() []byte {

	b, err := json.Marshal(ep.Map())
	if err != nil {
		return []byte{}
	}
	return b
}



//getter methods

// return a transitive version of error plus with no methods and json annotations
func (ep *ErrorPlus) ErrorPlus() *TErrorPlus{
	return &TErrorPlus{
		Err:     ep.err,
		Message: ep.message,
		Tags:    ep.tags,
		FN:      ep.fN,
		Args:    ep.args,
		Trace:   ep.trace,
		Time:    ep.time,
		RuntimeGoVer: ep.runtimeGoVer,
	}
}


// TESTME

// get the original error without the error plus properties complain with the error interface
func (ep *ErrorPlus) Error() string {
	err := ep.err.Error()
	return err
}

// TESTME
func (ep *ErrorPlus) GetError() error{
	return ep.err
}

// TESTME
// get the message associated with the error plus object  
func (ep *ErrorPlus) GetMessage() string{
	return ep.message
}

// TESTME
// get the tags associated with the error plus object
func (ep *ErrorPlus) GetTags() []string{
	return ep.tags
}

// TESTME
func (ep *ErrorPlus) GetFN() string{
	return ep.fN
}


// TESTME
func (ep *ErrorPlus) GetArgs() []interface{}{
	return ep.args
}

// TESTME
func (ep *ErrorPlus) GetTrace() string{
	return ep.trace
}

// TESTME
func (ep *ErrorPlus) GetTime() time.Time{
	return ep.time
}


// TESTME
func (ep *ErrorPlus) GetRuntimeGoVer() string{
return ep.runtimeGoVer
}