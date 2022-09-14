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

package client

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	v1 "github.com/attestantio/go-relay-client/api/v1"
)

// Service is the service providing a connection to an MEV relay.
type Service interface {
	// Name returns the name of the relay implementation.
	Name() string

	// Address returns the address of the relay.
	Address() string

	// Pubkey returns the public key of the relay (if any).
	Pubkey() *phase0.BLSPubKey
}

// QueuedProposersProviders is the interface for providing queued proposer information.
type QueuedProposersProvider interface {
	Service

	// QueuedProposers provides information on the proposers queued to obtain a blinded block.
	QueuedProposers(ctx context.Context) ([]*v1.QueuedProposer, error)
}
