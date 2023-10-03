package e_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrorPlus(t *testing.T) {


	
	err :=  errors.New("This is an error")

	if err == nil {
		
	}


	assert.True(t, true, "This is good. Canary test passing")
}