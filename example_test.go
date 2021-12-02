package maskmerge_test

import (
	"fmt"
	"log"

	"github.com/iterate/maskmerge"
	"github.com/iterate/maskmerge/internal/prototest"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func ExampleMaskMerge() {
	// Original data to update
	stored := &prototest.B1{D: 12, X: 15}
	// New data from the user
	userData := &prototest.B1{X: 42}
	// Only update path x
	mask, err := fieldmaskpb.New(userData, "x")
	if err != nil {
		log.Fatalf("invalid field mask: %v", err)
	}

	if err := maskmerge.MaskMerge(stored, userData, mask); err != nil {
		log.Fatalf("merging messages: %v", err)
	}

	// not using prototext because the output is not stable.
	fmt.Printf("d:%d x:%d", stored.D, stored.X)
	// Output: d:12 x:42
}
