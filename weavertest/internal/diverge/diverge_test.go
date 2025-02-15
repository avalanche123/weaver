// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package diverge

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
)

//go:generate ../../../cmd/weaver/weaver generate ./...

// errorRecorder is passed to weavertest.Run to capture error during
// weaver initialization.
type errorRecorder struct {
	testing.TB
	err error
}

func (rec *errorRecorder) Fatal(args ...any) {
	if len(args) == 1 {
		if err, ok := args[0].(error); ok {
			rec.err = err
			return
		}
	}
	rec.TB.Fatal(args...)
}

// TestFailer demonstrates that we only propagate constructor errors in
// singleprocess mode. In multiprocess mode, if a constructor returns an error,
// the entire process crashes. This causes us to get a non-nil network error.
func TestFailer(t *testing.T) {
	for _, single := range []bool{true, false} {
		t.Run(fmt.Sprintf("Single=%v", single), func(t *testing.T) {
			recorder := &errorRecorder{t, nil}
			weavertest.Run(recorder, weavertest.Options{SingleProcess: single}, func(f Failer) {})
			if want, got := single, errors.Is(recorder.err, ErrFailed); want != got {
				t.Fatalf("expecting Is(ErrFailed) = %v, got %v for error %v", want, got, recorder.err)
			}
		})
	}
}

// TestDealiasing demonstrates that pointers are only de-aliased in multiprocess mode.
func TestDealiasing(t *testing.T) {
	for _, single := range []bool{true, false} {
		t.Run(fmt.Sprintf("Single=%v", single), func(t *testing.T) {
			weavertest.Run(t, weavertest.Options{SingleProcess: single}, func(p Pointer) {
				pair, err := p.Get(context.Background())
				if err != nil {
					t.Fatal(err)
				}
				if want, got := single, (pair.X == pair.Y); want != got {
					t.Fatalf("expecting aliasing = %v, got %v", want, got)
				}
			})
		})
	}
}

// TestCustomErrors* demonstrates that custom Is methods are ignored in multiprocess mode.
func TestCustomErrors(t *testing.T) {
	for _, single := range []bool{true, false} {
		t.Run(fmt.Sprintf("Single=%v", single), func(t *testing.T) {
			weavertest.Run(t, weavertest.Options{SingleProcess: single}, func(e Errer) {
				err := e.Err(context.Background(), 1)
				if want, got := single, errors.Is(err, IntError{2}); want != got {
					t.Fatalf("expecting Is(IntError{2}) = %v, got %v for error %v", want, got, err)
				}
			})
		})
	}
}
