package digilog

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBuff *bytes.Buffer
var Out *BuffOut

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
	l.SetOutput(Out)
	l.Debug("test_event", "salutation='", "hello world", "'")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello world'\n"), "failed asserting that %s ends with %s", testBuff.String(), "event_id=test_event salutation='hello world'\n")
}

func TestDebug_AddTagStructAndMap(t *testing.T) {
	type TestStruct struct {
		One   string
		Two   int
		Three bool
	}

	s := TestStruct{
		One:   "one",
		Two:   2,
		Three: true,
	}

	m := map[string]interface{}{
		"Four": "Four",
		"Five": 5,
		"Six":  false,
	}
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "DEBUG"
	l := New()
	l.SetOutput(Out)
	l.AddTag("test_struct", s)
	l.AddTag("test_map", m)
	l.AddMeta("meta_string", "salutations")
	l.AddMeta("meta_bool", true)

	// Test that Tags and Meta are added in an expected format
	l.Debug("test_event", "salutation='", "hello world", "'")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	e := "event_id=test_event test_struct=\"{One:one Two:2 Three:true}\" test_map=\"map[Five:5 Four:Four Six:false]\" meta_string=\"salutations\" meta_bool=\"true\" salutation='hello world'\n"
	assert.True(strings.HasSuffix(testBuff.String(), e), "failed asserting that %s ends with %s", testBuff.String(), e)

	// Test that meta is removed after the log is written
	testBuff.Reset()
	l.Debug("test_event", "salutation='", "hello world", "'")
	e = "event_id=test_event test_struct=\"{One:one Two:2 Three:true}\" test_map=\"map[Five:5 Four:Four Six:false]\" salutation='hello world'\n"
	assert.True(strings.HasSuffix(testBuff.String(), e), "failed asserting that %s ends with %s", testBuff.String(), e)
}

func TestDebug_AddTagsMetas(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "DEBUG"
	l := New()
	l.SetOutput(Out)

	l.AddTags(map[string]interface{}{
		"salutation": "hello world",
	})
	l.AddMetas(map[string]interface{}{
		"next_salutation": "goodbye world",
	})

	l.Debug("test_event")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	e := "event_id=test_event salutation=\"hello world\" next_salutation=\"goodbye world\" \n"
	assert.True(strings.HasSuffix(testBuff.String(), e), "failed asserting that %s ends with %s", testBuff.String(), e)

	testBuff.Reset()
	l.Debug("test_event")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	e = "event_id=test_event salutation=\"hello world\" \n"
	assert.True(strings.HasSuffix(testBuff.String(), e), "failed asserting that %s ends with %s", testBuff.String(), e)
}

func TestDebugf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "DEBUG"
	l := New()
	l.SetOutput(Out)
	l.Debugf("test_event", "salutation='%s'", "hello world")
	assert.True(strings.Contains(testBuff.String(), "DEBUG"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello world'\n"))
}

func TestInfo(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "INFO"
	l := New()
	l.SetOutput(Out)
	l.Info("test_event", "salutation='", "hello mother", "'")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello mother'\n"))
}

func TestInfoAddTag(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "INFO"
	l := New()
	l.SetOutput(Out)
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
	l.SetOutput(Out)
	l.Infof("test_event", "salutation='%s'", "hello mother")
	assert.True(strings.Contains(testBuff.String(), "INFO"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello mother'\n"))
}

func TestWarn(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "WARN"
	l := New()
	l.SetOutput(Out)
	l.Warn("test_event", "salutation='", "hello father", "'")
	assert.True(strings.Contains(testBuff.String(), "WARN"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello father'\n"))
}

func TestWarnf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "WARN"
	l := New()
	l.SetOutput(Out)
	l.Warnf("test_event", "salutation='%s'", "hello father")
	assert.True(strings.Contains(testBuff.String(), "WARN"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello father'\n"))
}

func TestError(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "ERROR"
	l := New()
	l.SetOutput(Out)
	l.Error("test_event", "salutation='", "hello sister", "'")
	assert.True(strings.Contains(testBuff.String(), "ERROR"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello sister'\n"))
}

func TestErrorf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "ERROR"
	l := New()
	l.SetOutput(Out)
	l.Errorf("test_event", "salutation='%s'", "hello sister")
	assert.True(strings.Contains(testBuff.String(), "ERROR"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello sister'\n"))
}

func TestCritical(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "CRITICAL"
	l := New()
	l.SetOutput(Out)
	l.Critical("test_event", fmt.Errorf("salutation='%s'", "hello brother"))
	assert.True(strings.Contains(testBuff.String(), "CRITICAL"))
	assert.True(strings.HasSuffix(testBuff.String(), "event_id=test_event salutation='hello brother'\n"))
}

func TestCriticalf(t *testing.T) {
	testBuff.Reset()
	assert := assert.New(t)

	LogLevel = "CRITICAL"
	l := New()
	l.SetOutput(Out)
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
	l.SetOutput(Out)
	l.Debugf("test_event", "salutation='%s'", "hello empty void")
	assert.Empty(testBuff.String())
}
