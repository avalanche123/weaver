// go:build !ignoreWeaverGen

package weavertest

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)
var _ codegen.LatestVersion = codegen.Version[[0][10]struct{}]("You used 'weaver generate' codegen version 0.10.0, but you built your code with an incompatible weaver module version. Try upgrading 'weaver generate' and re-running it.")

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/weavertest/testMainInterface",
		Iface: reflect.TypeOf((*testMainInterface)(nil)).Elem(),
		Impl:  reflect.TypeOf(testMain{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return testMainInterface_local_stub{impl: impl.(testMainInterface), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return testMainInterface_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return testMainInterface_server_stub{impl: impl.(testMainInterface), addLoad: addLoad}
		},
		RefData: "",
	})
}

// weaver.Instance checks.
var _ weaver.Instance = &testMain{}

// Local stub implementations.

type testMainInterface_local_stub struct {
	impl   testMainInterface
	tracer trace.Tracer
}

// Check that testMainInterface_local_stub implements the testMainInterface interface.
var _ testMainInterface = &testMainInterface_local_stub{}

// Client stub implementations.

type testMainInterface_client_stub struct {
	stub codegen.Stub
}

// Check that testMainInterface_client_stub implements the testMainInterface interface.
var _ testMainInterface = &testMainInterface_client_stub{}

// Server stub implementations.

type testMainInterface_server_stub struct {
	impl    testMainInterface
	addLoad func(key uint64, load float64)
}

// Check that testMainInterface_server_stub implements the codegen.Server interface.
var _ codegen.Server = &testMainInterface_server_stub{}

// GetStubFn implements the codegen.Server interface.
func (s testMainInterface_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}
