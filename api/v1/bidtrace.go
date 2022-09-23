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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// BidTrace represents a bid trace.
type BidTrace struct {
	Slot                 phase0.Slot
	ParentHash           phase0.Hash32
	BlockHash            phase0.Hash32
	BuilderPubkey        phase0.BLSPubKey
	ProposerPubkey       phase0.BLSPubKey
	ProposerFeeRecipient bellatrix.ExecutionAddress
	GasLimit             uint64
	GasUsed              uint64
	Value                *big.Int
}

// bidTraceJSON is the spec representation of the struct.
type bidTraceJSON struct {
	Slot                 string `json:"slot"`
	ParentHash           string `json:"parent_hash"`
	BlockHash            string `json:"block_hash"`
	BuilderPubkey        string `json:"builder_pubkey"`
	ProposerPubkey       string `json:"proposer_pubkey"`
	ProposerFeeRecipient string `json:"proposer_fee_recipient"`
	GasLimit             string `json:"gas_limit"`
	GasUsed              string `json:"gas_used"`
	Value                string `json:"value"`
}

// MarshalJSON implements json.Marshaler.
func (b *BidTrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(&bidTraceJSON{
		Slot:                 fmt.Sprintf("%d", b.Slot),
		ParentHash:           fmt.Sprintf("%#x", b.ParentHash),
		BlockHash:            fmt.Sprintf("%#x", b.BlockHash),
		BuilderPubkey:        fmt.Sprintf("%#x", b.BuilderPubkey),
		ProposerPubkey:       fmt.Sprintf("%#x", b.ProposerPubkey),
		ProposerFeeRecipient: fmt.Sprintf("%#x", b.ProposerFeeRecipient),
		GasLimit:             fmt.Sprintf("%d", b.GasLimit),
		GasUsed:              fmt.Sprintf("%d", b.GasUsed),
		Value:                b.Value.String(),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *BidTrace) UnmarshalJSON(input []byte) error {
	var data bidTraceJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	return b.unpack(&data)
}

func (b *BidTrace) unpack(data *bidTraceJSON) error {
	if data.Slot == "" {
		return errors.New("slot missing")
	}
	slot, err := strconv.ParseUint(data.Slot, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for slot")
	}
	b.Slot = phase0.Slot(slot)

	if data.ParentHash == "" {
		return errors.New("parent hash missing")
	}
	parentHash, err := hex.DecodeString(strings.TrimPrefix(data.ParentHash, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for parent hash")
	}
	if len(parentHash) != phase0.RootLength {
		return errors.New("incorrect length for parent hash")
	}
	copy(b.ParentHash[:], parentHash)

	if data.BlockHash == "" {
		return errors.New("block hash missing")
	}
	blockHash, err := hex.DecodeString(strings.TrimPrefix(data.BlockHash, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for block hash")
	}
	if len(blockHash) != phase0.RootLength {
		return errors.New("incorrect length for block hash")
	}
	copy(b.BlockHash[:], blockHash)

	if data.BuilderPubkey == "" {
		return errors.New("builder pubkey missing")
	}
	builderPubkey, err := hex.DecodeString(strings.TrimPrefix(data.BuilderPubkey, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for builder pubkey")
	}
	if len(builderPubkey) != phase0.PublicKeyLength {
		return errors.New("incorrect length for builder pubkey")
	}
	copy(b.BuilderPubkey[:], builderPubkey)

	if data.ProposerPubkey == "" {
		return errors.New("proposer pubkey missing")
	}
	proposerPubkey, err := hex.DecodeString(strings.TrimPrefix(data.ProposerPubkey, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for proposer pubkey")
	}
	if len(proposerPubkey) != phase0.PublicKeyLength {
		return errors.New("incorrect length for proposer pubkey")
	}
	copy(b.ProposerPubkey[:], proposerPubkey)

	if data.ProposerFeeRecipient == "" {
		return errors.New("proposer fee recipient missing")
	}
	proposerFeeRecipient, err := hex.DecodeString(strings.TrimPrefix(data.ProposerFeeRecipient, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for proposer fee recipient")
	}
	if len(proposerFeeRecipient) != bellatrix.FeeRecipientLength {
		return errors.New("incorrect length for proposer fee recipient")
	}
	copy(b.ProposerFeeRecipient[:], proposerFeeRecipient)

	if data.GasLimit == "" {
		return errors.New("gas limit missing")
	}
	gasLimit, err := strconv.ParseUint(data.GasLimit, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for gas limit")
	}
	b.GasLimit = gasLimit

	if data.GasUsed == "" {
		return errors.New("gas used missing")
	}
	gasUsed, err := strconv.ParseUint(data.GasUsed, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for gas used")
	}
	b.GasUsed = gasUsed

	if data.Value == "" {
		return errors.New("value missing")
	}
	value, success := new(big.Int).SetString(data.Value, 10)
	if !success {
		return errors.New("value invalid")
	}
	b.Value = value

	return nil
}

// String returns a string version of the structure.
func (b *BidTrace) String() string {
	data, err := json.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
