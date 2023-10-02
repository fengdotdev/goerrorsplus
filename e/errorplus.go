package e

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type ErrorPlus struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
	Tags    []string `json:"tags"`
	FN      string `json:"fn"`
	Args    []interface{} `json:"args"`
	Trace   string `json:"trace"`
	Time    time.Time `json:"time"`
}

// TESTME
func NewErrorPlus(err error, message string, tags []string, fn interface{}, args ...interface{}) *ErrorPlus {
	ep := &ErrorPlus{
		Err:     err,
		Tags:    tags,
		Message: message,
		FN:      runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
		Args:    args,
		Trace:   runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
		Time:    time.Now(),
	}
	return ep
}

// TESTME
func (ep *ErrorPlus) Error() string {
	err := ep.Err.Error()
	return err
}
// TESTME
// short hand for VerboseError
func (ep *ErrorPlus) V() error {
	return ep.VerboseError()
}
// TESTME
func (ep *ErrorPlus) VerboseError() error {

	timestamp := ep.Time.Format("2006-01-02 15:04:05")
	err := ep.Err.Error()
	message := ep.Message
	trace := ep.Trace
	fn := ep.FN
	args := ep.Args

	//format 2006-01-02 15:04:05 error: original-error trace: /something/something/fn.errorFunc args: [1 2 3]
	return fmt.Errorf("%s error: %s  %s fn: %s trace: %s  %s", timestamp, err, message, fn, trace, args)
}
// TESTME
func (ep *ErrorPlus) Map() map[string]interface{} {
	return map[string]interface{}{
		"error":     ep.Err,
		"message": ep.Message,
		"tags":    ep.Tags,
		"fn":      ep.FN,
		"args":    ep.Args,
		"trace":   ep.Trace,
		"time":    ep.Time,
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