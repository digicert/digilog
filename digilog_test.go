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
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "DEBUG"
	l := New()
	l.Debug("test_event", "salutation='", "hello world", "'")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello world'\n"), "failed asserting that %s ends with %s", testBuff.String(), "event_id=test_event salutation='hello world'\n")
}
func TestDebugf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "DEBUG"
	l := New()
	l.Debugf("test_event", "salutation='%s'", "hello world")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello world'\n"))
}

func TestInfo(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "INFO"
	l := New()
	l.Info("test_event", "salutation='", "hello mother", "'")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello mother'\n"))
}

func TestInfoAddTag(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "INFO"
	l := New()
	l.AddTag("foo", "bar")
	l.Info("test_event", "salutation='", "hello mother", "'")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event foo=\"bar\" salutation='hello mother'\n"), "failed asserting that %s ends with %s", testBuff.String(), "event_id=test_event salutation='hello world'\n")
}

func TestInfof(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "INFO"
	l := New()
	l.Infof("test_event", "salutation='%s'", "hello mother")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello mother'\n"))
}

func TestWarn(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "WARN"
	l := New()
	l.Warn("test_event", "salutation='", "hello father", "'")
	assert.True(strings.Contains(testBuff.String(), "WARN"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello father'\n"))
}

func TestWarnf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "WARN"
	l := New()
	l.Warnf("test_event", "salutation='%s'", "hello father")
	assert.True(strings.Contains(testBuff.String(), "WARN"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello father'\n"))
}

func TestError(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "ERROR"
	l := New()
	l.Error("test_event", "salutation='", "hello sister", "'")
	assert.True(strings.Contains(testBuff.String(), "ERROR"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello sister'\n"))
}

func TestErrorf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "ERROR"
	l := New()
	l.Errorf("test_event", "salutation='%s'", "hello sister")
	assert.True(strings.Contains(testBuff.String(), "ERROR"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello sister'\n"))
}

func TestCritical(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "CRITICAL"
	l := New()
	l.Critical("test_event", fmt.Errorf("salutation='%s'", "hello brother"))
	assert.True(strings.Contains(testBuff.String(), "CRITICAL"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello brother'\n"))
}

func TestCriticalf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "CRITICAL"
	l := New()
	l.Criticalf("test_event", "salutation='%s'", "hello brother")
	assert.True(strings.Contains(testBuff.String(), "CRITICAL"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello brother'\n"))
}

func TestLogLevel(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "INFO"
	testBuff = &bytes.Buffer{}
	l := New()
	l.Debugf("test_event", "salutation='%s'", "hello empty void")
	assert.Empty(testBuff.String())
}
