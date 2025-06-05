package basepath_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-api-cms/internal/basepath"
)

func TestDir(t *testing.T) {
	basePath := basepath.Dir()
	fmt.Println("Base Path Dir:", basePath)
	assert.NotEmpty(t, basePath)
}
