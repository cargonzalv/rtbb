//go:build tools
// +build tools

package tools

//suggested way of tracking all tool dependencies
// Refer https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	_ "go.uber.org/mock/gomock"
	_ "google.golang.org/protobuf/compiler/protogen"
	_ "google.golang.org/protobuf/proto"
)
