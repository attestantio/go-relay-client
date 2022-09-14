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

package v1

import (
	"encoding/json"
	"fmt"
	"strconv"

	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// QueuedProposer represents a queued proposer.
type QueuedProposer struct {
	Slot  phase0.Slot
	Entry *v1.SignedValidatorRegistration
}

// queuedProposerJSON is the spec representation of the struct.
type queuedProposerJSON struct {
	Slot  string                          `json:"slot"`
	Entry *v1.SignedValidatorRegistration `json:"entry"`
}

// MarshalJSON implements json.Marshaler.
func (q *QueuedProposer) MarshalJSON() ([]byte, error) {
	return json.Marshal(&queuedProposerJSON{
		Slot:  fmt.Sprintf("%d", q.Slot),
		Entry: q.Entry,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (q *QueuedProposer) UnmarshalJSON(input []byte) error {
	var data queuedProposerJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	return q.unpack(&data)
}

func (q *QueuedProposer) unpack(data *queuedProposerJSON) error {
	if data.Slot == "" {
		return errors.New("slot missing")
	}
	slot, err := strconv.ParseUint(data.Slot, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for slot")
	}
	q.Slot = phase0.Slot(slot)

	if data.Entry == nil {
		return errors.New("entry missing")
	}
	q.Entry = data.Entry

	return nil
}

// String returns a string version of the structure.
func (q *QueuedProposer) String() string {
	data, err := json.Marshal(q)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
