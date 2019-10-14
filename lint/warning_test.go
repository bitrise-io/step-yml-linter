package lint

import (
	"testing"
)

func TestWarning_Error(t *testing.T) {
	type fields struct {
		Messages []string
		Line     int
		Column   int
	}
	tests := []struct {
		name   string
		fields fields
		wantS  string
	}{
		{"simple test", fields{[]string{"line 1", "line 2"}, 3, 5}, "3:5: line 1\n3:5: line 2\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Warning{
				Messages: tt.fields.Messages,
				Line:     tt.fields.Line,
				Column:   tt.fields.Column,
			}
			if gotS := w.Error(); gotS != tt.wantS {
				t.Errorf("Warning.Error() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
