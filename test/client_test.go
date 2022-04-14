package test

import (
	"os"
	"testing"

	"github.com/luxrobo/luxpay/src/client"
	"github.com/stretchr/testify/assert"
)

func setUpMockEnvVars(t *testing.T) {
	setUpTossMockEnvVars(t)
}

func TestNewClient(t *testing.T) {
	setUpMockEnvVars(t)
	tossSecret := os.Getenv("DEV_TOSS_SECRET")
	clientInst := client.NewClient("toss", tossSecret)
	assert.NotNil(t, clientInst)
}
