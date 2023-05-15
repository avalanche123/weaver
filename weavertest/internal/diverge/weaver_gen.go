// go:build !ignoreWeaverGen

package diverge

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)
var _ codegen.LatestVersion = codegen.Version[[0][10]struct{}]("You used 'weaver generate' codegen version 0.10.0, but you built your code with an incompatible weaver module version. Try upgrading 'weaver generate' and re-running it.")

func init() {
	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Errer",
		Iface:       reflect.TypeOf((*Errer)(nil)).Elem(),
		Impl:        reflect.TypeOf(errer{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return errer_local_stub{impl: impl.(Errer), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return errer_client_stub{stub: stub, errMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Errer", Method: "Err"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return errer_server_stub{impl: impl.(Errer), addLoad: addLoad}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Failer",
		Iface:       reflect.TypeOf((*Failer)(nil)).Elem(),
		Impl:        reflect.TypeOf(failer{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return failer_local_stub{impl: impl.(Failer), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return failer_client_stub{stub: stub, imJustHereSoWeaverGenerateDoesntComplainMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Failer", Method: "ImJustHereSoWeaverGenerateDoesntComplain"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return failer_server_stub{impl: impl.(Failer), addLoad: addLoad}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Pointer",
		Iface: reflect.TypeOf((*Pointer)(nil)).Elem(),
		Impl:  reflect.TypeOf(pointer{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return pointer_local_stub{impl: impl.(Pointer), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return pointer_client_stub{stub: stub, getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Pointer", Method: "Get"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return pointer_server_stub{impl: impl.(Pointer), addLoad: addLoad}
		},
		RefData: "",
	})
}

// Local stub implementations.

type errer_local_stub struct {
	impl   Errer
	tracer trace.Tracer
}

// Check that errer_local_stub implements the Errer interface.
var _ Errer = &errer_local_stub{}

func (s errer_local_stub) Err(ctx context.Context, a0 int) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "diverge.Errer.Err", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Err(ctx, a0)
}

type failer_local_stub struct {
	impl   Failer
	tracer trace.Tracer
}

// Check that failer_local_stub implements the Failer interface.
var _ Failer = &failer_local_stub{}

func (s failer_local_stub) ImJustHereSoWeaverGenerateDoesntComplain(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "diverge.Failer.ImJustHereSoWeaverGenerateDoesntComplain", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ImJustHereSoWeaverGenerateDoesntComplain(ctx)
}

type pointer_local_stub struct {
	impl   Pointer
	tracer trace.Tracer
}

// Check that pointer_local_stub implements the Pointer interface.
var _ Pointer = &pointer_local_stub{}

func (s pointer_local_stub) Get(ctx context.Context) (r0 Pair, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "diverge.Pointer.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx)
}

// Client stub implementations.

type errer_client_stub struct {
	stub       codegen.Stub
	errMetrics *codegen.MethodMetrics
}

// Check that errer_client_stub implements the Errer interface.
var _ Errer = &errer_client_stub{}

func (s errer_client_stub) Err(ctx context.Context, a0 int) (err error) {
	// Update metrics.
	start := time.Now()
	s.errMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "diverge.Errer.Err", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.errMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.errMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	s.errMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.errMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type failer_client_stub struct {
	stub                                            codegen.Stub
	imJustHereSoWeaverGenerateDoesntComplainMetrics *codegen.MethodMetrics
}

// Check that failer_client_stub implements the Failer interface.
var _ Failer = &failer_client_stub{}

func (s failer_client_stub) ImJustHereSoWeaverGenerateDoesntComplain(ctx context.Context) (err error) {
	// Update metrics.
	start := time.Now()
	s.imJustHereSoWeaverGenerateDoesntComplainMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "diverge.Failer.ImJustHereSoWeaverGenerateDoesntComplain", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.imJustHereSoWeaverGenerateDoesntComplainMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.imJustHereSoWeaverGenerateDoesntComplainMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.imJustHereSoWeaverGenerateDoesntComplainMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.imJustHereSoWeaverGenerateDoesntComplainMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type pointer_client_stub struct {
	stub       codegen.Stub
	getMetrics *codegen.MethodMetrics
}

// Check that pointer_client_stub implements the Pointer interface.
var _ Pointer = &pointer_client_stub{}

func (s pointer_client_stub) Get(ctx context.Context) (r0 Pair, err error) {
	// Update metrics.
	start := time.Now()
	s.getMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "diverge.Pointer.Get", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.getMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.getMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type errer_server_stub struct {
	impl    Errer
	addLoad func(key uint64, load float64)
}

// Check that errer_server_stub implements the codegen.Server interface.
var _ codegen.Server = &errer_server_stub{}

// GetStubFn implements the stub.Server interface.
func (s errer_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Err":
		return s.err
	default:
		return nil
	}
}

func (s errer_server_stub) err(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Err(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type failer_server_stub struct {
	impl    Failer
	addLoad func(key uint64, load float64)
}

// Check that failer_server_stub implements the codegen.Server interface.
var _ codegen.Server = &failer_server_stub{}

// GetStubFn implements the stub.Server interface.
func (s failer_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "ImJustHereSoWeaverGenerateDoesntComplain":
		return s.imJustHereSoWeaverGenerateDoesntComplain
	default:
		return nil
	}
}

func (s failer_server_stub) imJustHereSoWeaverGenerateDoesntComplain(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.ImJustHereSoWeaverGenerateDoesntComplain(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type pointer_server_stub struct {
	impl    Pointer
	addLoad func(key uint64, load float64)
}

// Check that pointer_server_stub implements the codegen.Server interface.
var _ codegen.Server = &pointer_server_stub{}

// GetStubFn implements the stub.Server interface.
func (s pointer_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Get":
		return s.get
	default:
		return nil
	}
}

func (s pointer_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = &Pair{}

func (x *Pair) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Pair.WeaverMarshal: nil receiver"))
	}
	serviceweaver_enc_ptr_int_98a2a745(enc, x.X)
	serviceweaver_enc_ptr_int_98a2a745(enc, x.Y)
}

func (x *Pair) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Pair.WeaverUnmarshal: nil receiver"))
	}
	x.X = serviceweaver_dec_ptr_int_98a2a745(dec)
	x.Y = serviceweaver_dec_ptr_int_98a2a745(dec)
}

func serviceweaver_enc_ptr_int_98a2a745(enc *codegen.Encoder, arg *int) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.Int(*arg)
	}
}

func serviceweaver_dec_ptr_int_98a2a745(dec *codegen.Decoder) *int {
	if !dec.Bool() {
		return nil
	}
	var res int
	res = dec.Int()
	return &res
}
