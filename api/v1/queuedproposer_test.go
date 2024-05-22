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

package v1_test

import (
	"encoding/json"
	"testing"

	v1 "github.com/attestantio/go-relay-client/api/v1"
	require "github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestQueuedProposer(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type v1.queuedProposerJSON",
		},
		{
			name:  "SlotMissing",
			input: []byte(`{"entry":{"message":{"fee_recipient":"0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5","gas_limit":"30000000","timestamp":"1663144444","pubkey":"0xa35e34e6aff03a0e37e0aeeeb2629ba3b503b285ddc75ff2ef8dc854653d833af289f0458cd614e3906ec5e9627b31db"},"signature":"0xb735529068b64c24c7650b08ddb09d543b79030888801176d2708f0e0c863a965fc1ba03f8fb14e5b3b486386e1f147b13848c218e143b513886a0f210c096bd03077fcac658c39402f2ca9075422a6df6b54f17f4141334239f9f9ff8137be0"}}`),
			err:   "slot missing",
		},
		{
			name:  "SlotWrongType",
			input: []byte(`{"slot":true,"entry":{"message":{"fee_recipient":"0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5","gas_limit":"30000000","timestamp":"1663144444","pubkey":"0xa35e34e6aff03a0e37e0aeeeb2629ba3b503b285ddc75ff2ef8dc854653d833af289f0458cd614e3906ec5e9627b31db"},"signature":"0xb735529068b64c24c7650b08ddb09d543b79030888801176d2708f0e0c863a965fc1ba03f8fb14e5b3b486386e1f147b13848c218e143b513886a0f210c096bd03077fcac658c39402f2ca9075422a6df6b54f17f4141334239f9f9ff8137be0"}}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field queuedProposerJSON.slot of type string",
		},
		{
			name:  "SlotInvalid",
			input: []byte(`{"slot":"true","entry":{"message":{"fee_recipient":"0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5","gas_limit":"30000000","timestamp":"1663144444","pubkey":"0xa35e34e6aff03a0e37e0aeeeb2629ba3b503b285ddc75ff2ef8dc854653d833af289f0458cd614e3906ec5e9627b31db"},"signature":"0xb735529068b64c24c7650b08ddb09d543b79030888801176d2708f0e0c863a965fc1ba03f8fb14e5b3b486386e1f147b13848c218e143b513886a0f210c096bd03077fcac658c39402f2ca9075422a6df6b54f17f4141334239f9f9ff8137be0"}}`),
			err:   "invalid value for slot: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "EntryMissing",
			input: []byte(`{"slot":"3887273"}`),
			err:   "entry missing",
		},
		{
			name:  "EntryWrongType",
			input: []byte(`{"slot":"3887273","entry":true}`),
			err:   "invalid JSON: invalid JSON: json: cannot unmarshal bool into Go value of type v1.signedValidatorRegistrationJSON",
		},
		{
			name:  "EntryInvalid",
			input: []byte(`{"slot":"3887273","entry":{}}`),
			err:   "invalid JSON: message missing",
		},
		{
			name:  "Good",
			input: []byte(`{"slot":"3887273","entry":{"message":{"fee_recipient":"0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad5","gas_limit":"30000000","timestamp":"1663144444","pubkey":"0xa35e34e6aff03a0e37e0aeeeb2629ba3b503b285ddc75ff2ef8dc854653d833af289f0458cd614e3906ec5e9627b31db"},"signature":"0xb735529068b64c24c7650b08ddb09d543b79030888801176d2708f0e0c863a965fc1ba03f8fb14e5b3b486386e1f147b13848c218e143b513886a0f210c096bd03077fcac658c39402f2ca9075422a6df6b54f17f4141334239f9f9ff8137be0"}}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res v1.QueuedProposer
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				assert.Equal(t, string(test.input), string(rt))
			}
		})
	}
}
