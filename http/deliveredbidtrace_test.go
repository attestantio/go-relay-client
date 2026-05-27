// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	client "github.com/attestantio/go-relay-client"
	"github.com/attestantio/go-relay-client/http"
	"github.com/stretchr/testify/require"
)

func TestDeliveredBidTrace(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Good",
		},
	}

	service, err := http.New(context.Background(),
		http.WithTimeout(timeout),
		http.WithAddress(os.Getenv("HTTP_ADDRESS")),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bidTrace, err := service.(client.DeliveredBidTraceProvider).DeliveredBidTrace(context.Background(), 4700906)
			require.NoError(t, err)
			require.NotNil(t, bidTrace)
			fmt.Printf("%v\n", bidTrace)
		})
	}
}

func TestDeliveredBidTracesCursor(t *testing.T) {
	tests := []struct {
		name        string
		upperSlot   uint64
		numberSlots int
		expectErr   bool
	}{
		{
			name:        "SingleSlot",
			upperSlot:   4700906,
			numberSlots: 1,
		},
		{
			name:        "MultipleSlots",
			upperSlot:   4700906,
			numberSlots: 10,
		},
		{
			name:        "ZeroLimit",
			upperSlot:   4700906,
			numberSlots: 0,
			expectErr:   true,
		},
		{
			name:        "NegativeLimit",
			upperSlot:   4700906,
			numberSlots: -1,
			expectErr:   true,
		},
	}

	service, err := http.New(context.Background(),
		http.WithTimeout(timeout),
		http.WithAddress(os.Getenv("HTTP_ADDRESS")),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bidTraces, err := service.(client.DeliveredBidTraceProvider).DeliveredBidTracesCursor(context.Background(), phase0.Slot(test.upperSlot), test.numberSlots)
			if test.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.LessOrEqual(t, len(bidTraces), test.numberSlots)
			for _, trace := range bidTraces {
				require.LessOrEqual(t, uint64(trace.Slot), test.upperSlot)
			}
		})
	}
}
