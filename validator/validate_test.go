/*
 * CDDL - A CDDL implementation for Go
 *
 * Copyright 2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.org/onflow/cddl/parser"
)

func TestValidate(t *testing.T) {

	t.Parallel()

	t.Run("valid", func(t *testing.T) {

		t.Parallel()

		cddl, err := parser.ParseCDDL(`
          x = 1
        `)
		require.NoError(t, err)

		err = Validate(cddl)
		require.NoError(t, err)

	})

	t.Run("use undefined", func(t *testing.T) {

		t.Parallel()

		cddl, err := parser.ParseCDDL(`
          x = y
        `)
		require.NoError(t, err)

		err = Validate(cddl)
		var usageErr UndefinedRuleUsageError
		require.ErrorAs(t, err, &usageErr)
		assert.Equal(t, "y", usageErr.ID)
	})
}
