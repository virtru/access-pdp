package protoconv_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/virtru/access-pdp/pdp"
	"github.com/virtru/access-pdp/protoconv"
)

func TestValueFailureToPbWithNilDataAttribute(t *testing.T) {
	failure := &pdp.ValueFailure{
		DataAttribute: nil,
		Message:       "test message",
	}

	pbFailure := protoconv.ValueFailureToPb(failure)

	require.NotNil(t, pbFailure, "The resulting pbFailure should not be nil")
	require.Nil(t, pbFailure.DataAttribute, "The DataAttribute in pbFailure should be nil")
	require.Equal(t, "test message", pbFailure.Message, "The Message in pbFailure should match the input")
}
