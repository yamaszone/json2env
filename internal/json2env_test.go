package json2env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToEnvWithValidData(t *testing.T) {
	err := ToEnv("../fixtures/sample.json", false)
	assert.Nil(t, err)

	err = ToEnv("../fixtures/sample.json", true)
	assert.Nil(t, err)
}

func TestToEnvWithInvalidData(t *testing.T) {
	err := ToEnv("../fixtures/invalid.json", false)
	assert.NotNil(t, err)
}
