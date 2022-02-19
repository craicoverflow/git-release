package version

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
)

func TestParse(t *testing.T) {
	type args struct {
		ref *plumbing.Reference
	}
	tests := []struct {
		name    string
		args    args
		want    *Version
		wantErr bool
	}{
		{
			name: "patch version should match",
			args: args{
				ref: plumbing.NewHashReference("refs/tags/v0.0.1", plumbing.NewHash("b1deda75c671cf6b1c8818bfe2e5ff86d36a78ff")),
			},
			want: &Version{
				Prefix: "v",
				Major:  0,
				Minor:  0,
				Patch:  1,
				Hash:   Hash(plumbing.NewHash("b1deda75c671cf6b1c8818bfe2e5ff86d36a78ff")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
			if !bytes.Equal(got.Hash[:], tt.want.Hash[:]) {
				t.Errorf("Parse() hash = %v, want %v", got.Hash, tt.want.Hash)
			}
		})
	}
}
