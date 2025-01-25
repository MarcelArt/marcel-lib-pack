package goment_test

import (
	"testing"

	"github.com/MarcelArt/marcel-lib-pack/pkg/goment"
	"github.com/stretchr/testify/assert"
)

func TestFormats(t *testing.T) {
	t.Run("should return formatted date DD/MM/YYYY", func(t *testing.T) {
		date, err := goment.New("2021-01-02", "YYYY-MM-DD")
		assert.NoError(t, err)

		dateString := date.Formats("DD/MM/YYYY")
		assert.Equal(t, "02/01/2021", dateString)
	})

	t.Run("should return formatted date dddd, MMMM D YYYY", func(t *testing.T) {
		date, err := goment.New("2021-01-02", "YYYY-MM-DD")
		assert.NoError(t, err)

		dateString := date.Formats("dddd, MMMM D YYYY")
		assert.Equal(t, "Saturday, January 2 2021", dateString)
	})
}
