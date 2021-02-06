package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		str         string
		expected    []byte
		expectedErr bool
	}{
		{
			str:         "0x01",
			expected:    []byte{0x01},
			expectedErr: false,
		},
		{
			str:         "0b1",
			expected:    []byte{0x01},
			expectedErr: false,
		},
		{
			str:         "0c1",
			expected:    []byte{0x01},
			expectedErr: true,
		},
		{
			str:         "10",
			expected:    []byte{0x0a},
			expectedErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			parse, err := Parse(test.str)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, parse)
			}
		})
	}
}
