package digilog

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBuff *bytes.Buffer

func init() {
	testBuff = &bytes.Buffer{}
	Out = &BuffOut{Out: testBuff, Err: testBuff}
}

func TestDebug(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "DEBUG"
	Debug("salutation='%s'", "hello world")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello world'\n"))
}

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "DEBUG"
	Info("salutation='%s'", "hello mother")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello mother'\n"))
}

func TestWarn(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "DEBUG"
	Warn("salutation='%s'", "hello father")
	assert.True(strings.Contains(testBuff.String(), "WARN"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello father'\n"))
}

func TestError(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "DEBUG"
	Error("salutation='%s'", "hello sister")
	assert.True(strings.Contains(testBuff.String(), "ERROR"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello sister'\n"))
}

// Can't test Critical ATM since it causes the test to exit
// func TestCritical(t *testing.T) {
// 	assert := assert.New(t)

// LogLevel = "DEBUG"
// 	Critical("salutation='%s'", "hello brother")
// 	assert.True(strings.Contains(testBuff.String(), "CRITICAL"))
// 	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello brother'\n"))
// }

func TestLogLevel(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "INFO"
	testBuff = &bytes.Buffer{}
	Debug("salutation='%s'", "hello empty void")
	assert.Empty(testBuff.String())
}
