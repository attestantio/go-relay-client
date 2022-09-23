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
)

// ReceivedBidTraces provides all bid traces received for a given slot.
func (s *Service) ReceivedBidTraces(ctx context.Context, slot phase0.Slot) ([]*v1.BidTraceWithTimestamp, error) {
	started := time.Now()

	url := fmt.Sprintf("/relay/v1/data/bidtraces/builder_blocks_received?slot=%d", slot)

	contentType, respBodyReader, err := s.get(ctx, url)
	if err != nil {
		log.Trace().Str("url", url).Err(err).Msg("Request failed")
		monitorOperation(s.Address(), "received bid traces", false, time.Since(started))
		return nil, errors.Wrap(err, "failed to request received bid traces")
	}
	if respBodyReader == nil {
		monitorOperation(s.Address(), "received bid traces", false, time.Since(started))
		return nil, errors.New("failed to obtain received bid traces")
	}

	res := make([]*v1.BidTraceWithTimestamp, 0)
	switch contentType {
	case ContentTypeJSON:
		if err := json.NewDecoder(respBodyReader).Decode(&res); err != nil {
			monitorOperation(s.Address(), "received bid traces", false, time.Since(started))
			return nil, errors.Wrap(err, "failed to parse received bid traces")
		}
	default:
		return nil, fmt.Errorf("unsupported content type %v", contentType)
	}

	monitorOperation(s.Address(), "received bid traces", true, time.Since(started))
	return res, nil
}
