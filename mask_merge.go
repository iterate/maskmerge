/*
Package maskmerge deals with merging protocol buffer messages using field masks.

See google.golang.org/protobuf/types/known/fieldmaskpb for more information on
how field masks are interpreted.
*/
package maskmerge

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// MaskMerge merges the fields in mask in src into dst.
// Following the specification field mask specification, lists and messages are
// merged, not replaced.
func MaskMerge(dst, src proto.Message, mask *fieldmaskpb.FieldMask) error {
	if !dst.ProtoReflect().IsValid() {
		return errors.New("invalid destination message")
	}
	dstr := dst.ProtoReflect()
	srcr := src.ProtoReflect()

	// Is this the best way to check if this is the same message?
	if stn, dtn := srcr.Descriptor().FullName(), dstr.Descriptor().FullName(); stn != dtn {
		return fmt.Errorf("incompatible types %s and %s", stn, dtn)
	}

	for _, fullPath := range mask.Paths {
		path, childPath, hasChild := cut(fullPath, ".")
		if len(path) == 0 {
			return fmt.Errorf("invalid field name %s", fullPath)
		}

		// get the field descriptor from this update mask path
		fd := dstr.Descriptor().Fields().ByName(protoreflect.Name(path))

		if fd == nil {
			return fmt.Errorf("invalid field name %s for %s", fullPath, dstr.Descriptor().Name())
		}
		_, _ = childPath, hasChild

		if hasChild {
			sm := &fieldmaskpb.FieldMask{Paths: []string{childPath}}

			// check if the target field is valid before we try to mutate it
			if !dstr.Has(fd) {
				// create it
				dstr.Set(fd, dstr.NewField(fd))
			}

			if fd.Kind() != protoreflect.MessageKind {
				return errors.New("nested field is not a message")
			} else if !sm.IsValid(dstr.Get(fd).Message().Interface()) {
				return fmt.Errorf("nested field %s is not valid for %s (%s)", childPath, path, fullPath)
			}

			if err := MaskMerge(dstr.Get(fd).Message().Interface(), srcr.Mutable(fd).Message().Interface(), sm); err != nil {
				return fmt.Errorf("applying to nested field %s: %w", fd.FullName(), err)
			}
			continue
		}

		switch {
		case fd.IsMap() || fd.IsList() || fd.Kind() == protoreflect.MessageKind:
			dstr.Set(fd, mergeMessages(fd, dstr.Get(fd), srcr.Get(fd)))
		default:
			dstr.Set(fd, srcr.Get(fd))
		}
	}

	return nil
}

func mergeMessages(fd protoreflect.FieldDescriptor, a protoreflect.Value, b protoreflect.Value) protoreflect.Value {
	switch {
	case fd.IsList():
		return mergeList(a, b)
	case fd.IsMap():
		return mergeMap(a, b)
	case fd.Kind() == protoreflect.MessageKind:
		return mergeMessage(a, b)
	}

	return a
}

func mergeMap(a protoreflect.Value, b protoreflect.Value) protoreflect.Value {
	dst := a.Map()
	src := b.Map()
	src.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		dst.Set(key, value)
		return true
	})
	return protoreflect.ValueOfMap(dst)
}

func mergeMessage(a protoreflect.Value, b protoreflect.Value) protoreflect.Value {
	dst := a.Message()
	src := b.Message()
	src.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		dst.Set(fd, v)
		return true
	})
	return protoreflect.ValueOfMessage(dst)
}

func mergeList(a protoreflect.Value, b protoreflect.Value) protoreflect.Value {
	dst := a.List()
	src := b.List()
	for i := 0; i < src.Len(); i++ {
		dst.Append(src.Get(i))
	}
	return protoreflect.ValueOfList(dst)
}

// cut cuts s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
//
// todo: replace with strings.Cut in go1.18
func cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
