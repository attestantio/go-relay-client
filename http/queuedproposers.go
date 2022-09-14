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

	v1 "github.com/attestantio/go-relay-client/api/v1"
	"github.com/pkg/errors"
)

// QueuedProposers provides information on the proposers queued to obtain a blinded block.
func (s *Service) QueuedProposers(ctx context.Context) ([]*v1.QueuedProposer, error) {
	started := time.Now()

	url := "/relay/v1/builder/validators"

	contentType, respBodyReader, err := s.get(ctx, url)
	if err != nil {
		log.Trace().Str("url", url).Err(err).Msg("Request failed")
		monitorOperation(s.Address(), "builder bid", false, time.Since(started))
		return nil, errors.Wrap(err, "failed to request queued proposers")
	}
	if respBodyReader == nil {
		monitorOperation(s.Address(), "builder bid", false, time.Since(started))
		return nil, errors.New("failed to obtain queued proposers")
	}

	res := make([]*v1.QueuedProposer, 0)
	switch contentType {
	case ContentTypeJSON:
		if err := json.NewDecoder(respBodyReader).Decode(&res); err != nil {
			return nil, errors.Wrap(err, "failed to parse queued proposers")
		}
	default:
		return nil, fmt.Errorf("unsupported content type %v", contentType)
	}

	monitorOperation(s.Address(), "queued proposers", true, time.Since(started))
	return res, nil
}
