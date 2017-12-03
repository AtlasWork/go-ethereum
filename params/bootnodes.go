// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Modified Atlas Developers for Atlas.Work 2017

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Atlas.Work Go Bootnodes
	"node://3133a428679b4736ac44f7c82d01d501ef9a08b172a5123ce19979c2b304e91f83a069052cfac58a5e4381fea621047f94782fbd9d01313fa5399ed73166d414@34.232.219.29:57200", // US-East
	"node://6cbbeeb07b9146c3c3240652d6237326e02f32619fd822e099046998463f460ffa59f62cd1f6abd14890bc5543b3a3d2a3431f4a9b81bcb3196c5f7bd89d0a40@52.90.10.109:57200"
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"", // US-TX
	"", // IE
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"node://3133a428679b4736ac44f7c82d01d501ef9a08b172a5123ce19979c2b304e91f83a069052cfac58a5e4381fea621047f94782fbd9d01313fa5399ed73166d414@34.232.219.29:57200",
	"node://6cbbeeb07b9146c3c3240652d6237326e02f32619fd822e099046998463f460ffa59f62cd1f6abd14890bc5543b3a3d2a3431f4a9b81bcb3196c5f7bd89d0a40@52.90.10.109:57200"
	
}
