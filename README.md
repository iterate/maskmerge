# Merge utility for Protobuf using FieldMask

`MaskMerge` implements merging of two protocol buffer messages as specified by the [FieldMask documentation].

## Installation

```bash
go get github.com/iterate/maskmerge
```

## Usage

Feed `MaskMerge` with a `proto.Message` to update, a `proto.Message` to read values from, and a `fieldmaskpb.FieldMask`.
See [example_test.go](./example_test.go)
for a full example.

```go
maskmerge.MaskMerge(dst, src, mask)
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](./LICENSE)

[FieldMask documentation]: https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask