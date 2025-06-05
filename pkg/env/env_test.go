package env_test

import (
	"os"
	"testing"
	"time"

	"bitbucket.org/apps_master_iglobal/magneto-pkg/env"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	os.Setenv("TEST_STRING", "hello")
	assert.Equal(t, "hello", env.Get("TEST_STRING"))
	assert.Equal(t, "default", env.Get("NOT_SET", "default"))
	assert.Equal(t, "", env.Get("ALSO_NOT_SET"))
}

func TestGetInt(t *testing.T) {
	os.Setenv("TEST_INT", "42")
	assert.Equal(t, 42, env.GetInt("TEST_INT"))
	assert.Equal(t, 10, env.GetInt("INVALID_INT", 10)) // not set yet, fallback
	os.Setenv("INVALID_INT", "abc")
	assert.Equal(t, 10, env.GetInt("INVALID_INT", 10))
}

func TestGetFloat64(t *testing.T) {
	os.Setenv("TEST_FLOAT", "3.14")
	assert.Equal(t, 3.14, env.GetFloat64("TEST_FLOAT"))
	assert.Equal(t, 1.23, env.GetFloat64("INVALID_FLOAT", 1.23)) // fallback
	os.Setenv("INVALID_FLOAT", "abc")
	assert.Equal(t, 1.23, env.GetFloat64("INVALID_FLOAT", 1.23))
}

func TestGetBool(t *testing.T) {
	os.Setenv("BOOL_TRUE", "true")
	os.Setenv("BOOL_FALSE", "false")
	os.Setenv("INVALID_BOOL", "notabool")

	assert.Equal(t, true, env.GetBool("BOOL_TRUE"))
	assert.Equal(t, false, env.GetBool("BOOL_FALSE"))
	assert.Equal(t, true, env.GetBool("INVALID_BOOL", true))
	assert.Equal(t, false, env.GetBool("NOT_SET_BOOL"))
}

func TestGetDuration(t *testing.T) {
	os.Setenv("TEST_DURATION", "5s")
	os.Setenv("INVALID_DURATION", "nonsense")

	assert.Equal(t, 5*time.Second, env.GetDuration("TEST_DURATION"))
	assert.Equal(t, 10*time.Second, env.GetDuration("INVALID_DURATION", 10*time.Second))
	assert.Equal(t, time.Duration(0), env.GetDuration("NOT_SET_DURATION"))
}

func TestGetSlice(t *testing.T) {
	os.Setenv("TEST_SLICE", "a,b,c")
	assert.Equal(t, []string{"a", "b", "c"}, env.GetSlice("TEST_SLICE", ","))

	assert.Equal(t, []string{"x", "y"}, env.GetSlice("NOT_SET_SLICE", ",", "x", "y"))
	assert.Equal(t, []string{}, env.GetSlice("ALSO_NOT_SET", ","))
}

func TestGetSliceInt(t *testing.T) {
	os.Setenv("TEST_SLICE_INT", "1,2,3")
	assert.Equal(t, []int{1, 2, 3}, env.GetSliceInt("TEST_SLICE_INT", ","))

	os.Setenv("INVALID_SLICE_INT", "1,abc,3")
	assert.Equal(t, []int{9, 8}, env.GetSliceInt("INVALID_SLICE_INT", ",", []int{9, 8}))

	assert.Equal(t, []int{5, 6}, env.GetSliceInt("NOT_SET_INT", ",", []int{5, 6}))
	assert.Equal(t, []int{}, env.GetSliceInt("ALSO_NOT_SET", ","))
}
