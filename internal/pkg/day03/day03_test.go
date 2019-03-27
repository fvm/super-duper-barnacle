package day03

import (
	"image"
	"reflect"
	"testing"
)

func Test_newClaim(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    Claim
		wantErr bool
	}{
		{
			name: "Test",
			args: args{
				line: "#1 @ 555,891: 18x12",
			},
			want: Claim{
				claim: image.Rectangle{
					Min: image.Point{
						X: 555,
						Y: 891,
					},
					Max: image.Point{
						X: 573,
						Y: 903,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newClaim(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("newClaim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClaim() = %v, want %v", got, tt.want)
			}
		})
	}
}
