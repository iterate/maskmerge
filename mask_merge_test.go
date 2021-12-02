package maskmerge

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/iterate/maskmerge/internal/prototest"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestMaskMerge(t *testing.T) {
	type args struct {
		dst  proto.Message
		src  proto.Message
		mask *fieldmaskpb.FieldMask
	}
	tests := []struct {
		name    string
		args    args
		want    proto.Message
		wantErr bool
	}{
		{
			name: "Simple",
			args: args{
				dst:  &prototest.Simple{FieldA: "Hi", FieldB: 23},
				src:  &prototest.Simple{FieldA: "Hello", FieldB: 2},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"field_a"}},
			},
			want: &prototest.Simple{FieldA: "Hello", FieldB: 23},
		},
		{
			name: "Nested",
			args: args{
				dst:  new(prototest.Nested),
				src:  &prototest.Nested{FieldOne: "Ignored", SubSimple: &prototest.Simple{FieldA: "ToUpdate", FieldB: 123}},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"sub_simple.field_a"}},
			},
			want: &prototest.Nested{SubSimple: &prototest.Simple{FieldA: "ToUpdate"}},
		},
		{
			name: "Invalid dst",
			args: args{
				dst:  new(prototest.Simple),
				src:  &prototest.Nested{FieldOne: "Ignored"},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"field_one"}},
			},
			wantErr: true,
		},
		{
			name: "Invalid src",
			args: args{
				dst:  &prototest.Nested{FieldOne: "Ignored"},
				src:  new(prototest.Simple),
				mask: &fieldmaskpb.FieldMask{Paths: []string{"field_one"}},
			},
			wantErr: true,
		},
		{
			name: "Canonical example",
			args: args{
				dst:  &prototest.R1{F: &prototest.A1{B: &prototest.B1{D: 1, X: 2}, C: []int32{1}}},
				src:  &prototest.R1{F: &prototest.A1{B: &prototest.B1{D: 10}, C: []int32{2}}},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"f.b", "f.c"}},
			},
			wantErr: false,
			want:    &prototest.R1{F: &prototest.A1{B: &prototest.B1{D: 10, X: 2}, C: []int32{1, 2}}},
		},
		{
			name: "Canonical example with multiple subfields",
			args: args{
				dst:  &prototest.R1{F: &prototest.A1{B: &prototest.B1{D: 1, X: 2}, C: []int32{1}}},
				src:  &prototest.R1{F: &prototest.A1{B: &prototest.B1{D: 10}, C: []int32{2}}},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"f.b.d", "f.b.x", "f.c"}},
			},
			wantErr: false,
			want:    &prototest.R1{F: &prototest.A1{B: &prototest.B1{D: 10}, C: []int32{1, 2}}},
		},
		{
			name: "Does not reset values",
			args: args{
				dst:  &prototest.R2{F: &prototest.B1{D: 1, X: 2}},
				src:  &prototest.R2{F: &prototest.B1{D: 3}},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"f"}},
			},
			wantErr: false,
			want:    &prototest.R2{F: &prototest.B1{D: 3, X: 2}},
		},
		{
			name: "Does reset values",
			args: args{
				dst:  &prototest.R2{F: &prototest.B1{D: 1, X: 2}},
				src:  &prototest.R2{F: &prototest.B1{D: 3}},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"f.x"}},
			},
			wantErr: false,
			want:    &prototest.R2{F: &prototest.B1{D: 1}},
		},
		{
			name: "Simple list merge",
			args: args{
				dst:  &prototest.A1{C: []int32{1, 2, 3}},
				src:  &prototest.A1{C: []int32{55}},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"c"}},
			},
			wantErr: false,
			want:    &prototest.A1{C: []int32{1, 2, 3, 55}},
		},
		{
			name: "Map overwrite",
			args: args{
				dst:  &prototest.A2{E: map[string]int32{"age": 32}},
				src:  &prototest.A2{E: map[string]int32{"age": 33}},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"e"}},
			},
			wantErr: false,
			want:    &prototest.A2{E: map[string]int32{"age": 33}},
		},
		{
			name: "Nested on non-message",
			args: args{
				dst:  new(prototest.Simple),
				src:  new(prototest.Simple),
				mask: &fieldmaskpb.FieldMask{Paths: []string{"field_a.a.b"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateTarget := proto.Clone(tt.args.dst)
			oldData := proto.Clone(tt.args.src)

			if err := MaskMerge(updateTarget, oldData, tt.args.mask); (err != nil) != tt.wantErr {
				t.Errorf("MaskMerge() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !proto.Equal(tt.want, updateTarget) {
				t.Errorf("MaskMerge() mismatch (-want +got):\n%s", cmp.Diff(tt.want, updateTarget, protocmp.Transform()))
			}
		})
	}
}
