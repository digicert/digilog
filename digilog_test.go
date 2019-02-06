package digilog

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBuff *bytes.Buffer

func init() {
	testBuff = &bytes.Buffer{}
	CriticalExit = false
	Out = &BuffOut{Out: testBuff, Err: testBuff}
}

func TestDebug(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "DEBUG"
	Debug("salutation='", "hello world", "'")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello world'\n"))
}
func TestDebugf(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "DEBUG"
	Debugf("salutation='%s'", "hello world")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello world'\n"))
}

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "INFO"
	Info("salutation='", "hello mother", "'")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello mother'\n"))
}

func TestInfof(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "INFO"
	Infof("salutation='%s'", "hello mother")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello mother'\n"))
}

func TestWarn(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "WARN"
	Warn("salutation='", "hello father", "'")
	assert.True(strings.Contains(testBuff.String(), "WARN"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello father'\n"))
}

func TestWarnf(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "WARN"
	Warnf("salutation='%s'", "hello father")
	assert.True(strings.Contains(testBuff.String(), "WARN"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello father'\n"))
}

func TestError(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "ERROR"
	Error("salutation='", "hello sister", "'")
	assert.True(strings.Contains(testBuff.String(), "ERROR"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello sister'\n"))
}

func TestErrorf(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "ERROR"
	Errorf("salutation='%s'", "hello sister")
	assert.True(strings.Contains(testBuff.String(), "ERROR"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello sister'\n"))
}

func TestCritical(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "CRITICAL"
	Critical(fmt.Errorf("salutation='%s'", "hello brother"))
	assert.True(strings.Contains(testBuff.String(), "CRITICAL"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello brother'\n"))
}

func TestCriticalf(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "CRITICAL"
	Criticalf("salutation='%s'", "hello brother")
	assert.True(strings.Contains(testBuff.String(), "CRITICAL"))
	assert.True(strings.HasSuffix(testBuff.String(), "salutation='hello brother'\n"))
}

func TestLogLevel(t *testing.T) {
	assert := assert.New(t)

	LogLevel = "INFO"
	testBuff = &bytes.Buffer{}
	Debugf("salutation='%s'", "hello empty void")
	assert.Empty(testBuff.String())
}
