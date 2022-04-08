// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"bytes"
	"github.com/stretchr/testify/require"
)

//nolint
func (s *azblobUnrecordedTestSuite) TestBytesWriterWriteAt() {
	_require := require.New(s.T())
	b := make([]byte, 10)
	buffer := newBytesWriter(b)

	count, err := buffer.WriteAt([]byte{1, 2}, 10)
	_require.Contains(err.Error(), "offset value is out of range")
	_require.Equal(count, 0)

	count, err = buffer.WriteAt([]byte{1, 2}, -1)
	_require.Contains(err.Error(), "offset value is out of range")
	_require.Equal(count, 0)

	count, err = buffer.WriteAt([]byte{1, 2}, 9)
	_require.Contains(err.Error(), "not enough space for all bytes")
	_require.Equal(count, 1)
	_require.Equal(bytes.Compare(b, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}), 0)

	count, err = buffer.WriteAt([]byte{1, 2}, 8)
	_require.Nil(err)
	_require.Equal(count, 2)
	_require.Equal(bytes.Compare(b, []byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 2}), 0)
}
