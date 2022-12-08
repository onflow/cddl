package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCDDL(t *testing.T) {

	t.Parallel()

	t.Run("valid", func(t *testing.T) {

		t.Parallel()

		_, err := ParseCDDL(`
          x = 1
        `)
		require.NoError(t, err)
	})

	t.Run("valid, complex", func(t *testing.T) {

		t.Parallel()

		// Appendix H. Examples, RFC7071 reputon

		_, err := ParseCDDL(`
          reputation-object = {
            reputation-context,
            reputon-list
          }

          reputation-context = (
            application: text
          )

          reputon-list = (
            reputons: reputon-array
          )

          reputon-array = [* reputon]

          reputon = {
            rater-value,
            assertion-value,
            rated-value,
            rating-value,
            ? conf-value,
            ? normal-value,
            ? sample-value,
            ? gen-value,
            ? expire-value,
            * ext-value,
          }
        `)
		require.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {

		t.Parallel()

		_, err := ParseCDDL(`
          x =
        `)
		require.Error(t, err)

		var parseErr ParseError
		require.ErrorAs(t, err, &parseErr)
	})
}
