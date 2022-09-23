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

func TestBidTrace(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type v1.bidTraceJSON",
		},
		{
			name:  "SlotMissing",
			input: []byte(`{"parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "slot missing",
		},
		{
			name:  "SlotWrongType",
			input: []byte(`{"slot":true,"parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.slot of type string",
		},
		{
			name:  "SlotInvalid",
			input: []byte(`{"slot":"-1","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid value for slot: strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:  "ParentHashMissing",
			input: []byte(`{"slot":"3939006","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "parent hash missing",
		},
		{
			name:  "ParentHashWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":true,"block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.parent_hash of type string",
		},
		{
			name:  "ParentHashInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"invalid","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid value for parent hash: encoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "ParentHashWrongLength",
			input: []byte(`{"slot":"3939006","parent_hash":"0xd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "incorrect length for parent hash",
		},
		{
			name:  "BlockHashMissing",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "block hash missing",
		},
		{
			name:  "BlockHashWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":true,"builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.block_hash of type string",
		},
		{
			name:  "BlockHashInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"invalid","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid value for block hash: encoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "BlockHashWrongLength",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "incorrect length for block hash",
		},
		{
			name:  "BuilderPubkeyMissing",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "builder pubkey missing",
		},
		{
			name:  "BuilderPubkeyWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":true,"proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.builder_pubkey of type string",
		},
		{
			name:  "BuilderPubkeyInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"invalid","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid value for builder pubkey: encoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "BuilderPubkeyWrongLength",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xdead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "incorrect length for builder pubkey",
		},
		{
			name:  "ProposerPubkeyMissing",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "proposer pubkey missing",
		},
		{
			name:  "ProposerPubkeyWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":true,"proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.proposer_pubkey of type string",
		},
		{
			name:  "ProposerPubkeyInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"invalid","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid value for proposer pubkey: encoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "ProposerPubkeyWrongLength",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x7d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "incorrect length for proposer pubkey",
		},
		{
			name:  "ProposerFeeRecipientMissing",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "proposer fee recipient missing",
		},
		{
			name:  "ProposerFeeRecipientWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":true,"gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.proposer_fee_recipient of type string",
		},
		{
			name:  "ProposerFeeRecipientInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"invalid","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid value for proposer fee recipient: encoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "ProposerFeeRecipientWrongLength",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0xa6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "incorrect length for proposer fee recipient",
		},
		{
			name:  "GasLimitMissing",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "gas limit missing",
		},
		{
			name:  "GasLimitWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":true,"gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.gas_limit of type string",
		},
		{
			name:  "GasLimitInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"-1","gas_used":"12077817","value":"34682404831419603"}`),
			err:   "invalid value for gas limit: strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:  "GasUsedMissing",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","value":"34682404831419603"}`),
			err:   "gas used missing",
		},
		{
			name:  "GasUsedWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":true,"value":"34682404831419603"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.gas_used of type string",
		},
		{
			name:  "GasUsedInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"-1","value":"34682404831419603"}`),
			err:   "invalid value for gas used: strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:  "ValueMissing",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817"}`),
			err:   "value missing",
		},
		{
			name:  "ValueWrongType",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field bidTraceJSON.value of type string",
		},
		{
			name:  "ValueInvalid",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"invalid"}`),
			err:   "value invalid",
		},
		{
			name:  "Good",
			input: []byte(`{"slot":"3939006","parent_hash":"0x6cd0618e3e13b751506264263b09979e461e35dec0dfbac20d81ece99a43b9dc","block_hash":"0x4c4f7e0a46a4f8b010bc7f899c949b5b9c0c58d510b6a5b46eda48d796a469ed","builder_pubkey":"0xa1dead01e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc","proposer_pubkey":"0x897d53adc5f6993166720dd365f924c0400a61be59cb53589009b8c3ba571032ca319de34e0459f6fcc8734e35a84fd0","proposer_fee_recipient":"0x32a6bcae2dd28f85555467d85600f4ecc8172808","gas_limit":"30000000","gas_used":"12077817","value":"34682404831419603"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res v1.BidTrace
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
