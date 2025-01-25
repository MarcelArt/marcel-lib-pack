package try_test

import (
	"testing"

	"github.com/MarcelArt/marcel-lib-pack/pkg/try"
	"github.com/stretchr/testify/assert"
)

func throw() (err error) {
	defer try.Catch(&err)
	panic("error")
}

func TestCatch(t *testing.T) {
	t.Run("should catch error", func(t *testing.T) {
		err := throw()

		assert.Error(t, err)
	})
}
