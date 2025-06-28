package geo

import (
	"reflect"
	"testing"
)

func TestConvertRawPolygon(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		want      PolygonGeometry
		expectErr bool
	}{
		{
			name: "valid single ring",
			input: []interface{}{
				[]interface{}{
					[]interface{}{1.0, 2.0},
					[]interface{}{3.0, 4.0},
					[]interface{}{5.0, 6.0},
					[]interface{}{1.0, 2.0}, // closed ring
				},
			},
			want: PolygonGeometry{
				{
					{1.0, 2.0},
					{3.0, 4.0},
					{5.0, 6.0},
					{1.0, 2.0},
				},
			},
			expectErr: false,
		},
		{
			name: "invalid type at top level",
			input: map[string]interface{}{
				"bad": "data",
			},
			want:      nil,
			expectErr: true,
		},
		{
			name: "invalid ring type",
			input: []interface{}{
				"not a ring",
			},
			want:      nil,
			expectErr: true,
		},
		{
			name: "invalid point structure",
			input: []interface{}{
				[]interface{}{
					"invalid",
				},
			},
			want:      nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertRawPolygon(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %+v, got %+v", tt.want, got)
			}
		})
	}
}
