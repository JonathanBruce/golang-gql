package schema

// TODO:
// 	- explain why I use .graphql files to define the schema
// 	- explain why we embed the .graphql files in the binary.
// 	- explain why this file is necessary and how the method is used.
//
// Use `go generate` to pack all *.graphql files under this directory (and sub-directories) into
// a binary format.
//
//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...

import "bytes"

// String 	Stringifies graphql files so that they may be consumed by graphql-go library
func String() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames {
		bufferFileData := MustAsset(name)
		buf.Write(bufferFileData)
		dataLength := len(bufferFileData)

		if dataLength > 0 && bufferFileData[dataLength-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
