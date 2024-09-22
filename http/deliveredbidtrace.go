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

package http

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	v1 "github.com/attestantio/go-relay-client/api/v1"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// DeliveredBidTrace provides a bid trace of a delivered payload for a given slot.
// Will return nil if the relay did not deliver a bid for the slot.
func (s *Service) DeliveredBidTrace(ctx context.Context, slot phase0.Slot) (*v1.BidTrace, error) {
	ctx, span := otel.Tracer("attestantio.go-relay-client.http").Start(ctx, "DeliveredBidTrace", trace.WithAttributes(
		//nolint:gosec
		attribute.Int64("slot", int64(slot)),
	))
	defer span.End()
	started := time.Now()

	url := fmt.Sprintf("/relay/v1/data/bidtraces/proposer_payload_delivered?slot=%d", slot)

	contentType, respBodyReader, err := s.get(ctx, url)
	if err != nil {
		log.Trace().Str("url", url).Err(err).Msg("Request failed")
		monitorOperation(s.Address(), "delivered bid trace", false, time.Since(started))
		return nil, errors.Wrap(err, "failed to request delivered bid trace")
	}
	if respBodyReader == nil {
		monitorOperation(s.Address(), "delivered bid trace", false, time.Since(started))
		return nil, errors.New("failed to obtain delivered bid trace")
	}

	res := make([]*v1.BidTrace, 0)
	switch contentType {
	case ContentTypeJSON:
		if err := json.NewDecoder(respBodyReader).Decode(&res); err != nil {
			return nil, errors.Wrap(err, "failed to parse delivered bid trace")
		}
	default:
		return nil, fmt.Errorf("unsupported content type %v", contentType)
	}

	if len(res) == 0 {
		// This means there was no delivered bid trace, but that's an acceptable response.
		monitorOperation(s.Address(), "delivered bid trace", true, time.Since(started))
		return nil, nil
	}

	monitorOperation(s.Address(), "delivered bid trace", true, time.Since(started))
	return res[0], nil
}
