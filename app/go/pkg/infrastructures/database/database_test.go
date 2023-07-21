package database

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientImpl(t *testing.T) {
	tests := []struct {
		name string
		want *ClientImpl
	}{
		{
			name: "success",
			want: &ClientImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClientImpl()

			assert.Equal(t, reflect.TypeOf(got), reflect.TypeOf(tt.want))
		})
	}
}
