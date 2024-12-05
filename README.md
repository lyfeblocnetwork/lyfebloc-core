# lyfebloc-core

[![Go Reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/lyfeblocnetwork/lyfebloc-core)
[![GitHub Release](https://img.shields.io/github/v/release/lyfeblocnetwork/lyfebloc-core)](https://github.com/lyfeblocnetwork/lyfebloc-core/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/lyfeblocnetwork/lyfebloc-core)](https://goreportcard.com/report/github.com/lyfeblocnetwork/lyfebloc-core)
[![Build](https://github.com/lyfeblocnetwork/lyfebloc-core/actions/workflows/build.yml/badge.svg)](https://github.com/lyfeblocnetwork/lyfebloc-core/actions/workflows/build.yml)
[![Lint](https://github.com/lyfeblocnetwork/lyfebloc-core/actions/workflows/lint.yml/badge.svg)](https://github.com/lyfeblocnetwork/lyfebloc-core/actions/workflows/lint.yml)
[![Tests](https://github.com/lyfeblocnetwork/lyfebloc-core/actions/workflows/tests.yml/badge.svg)](https://github.com/lyfeblocnetwork/lyfebloc-core/actions/workflows/tests.yml)

lyfebloc-core is a fork of [cometbft/cometbft](https://github.com/cometbft/cometbft), an implementation of the Tendermint protocol, with the following changes:

1. Modifications to how `DataHash` in the block header is determined. In CometBFT, `DataHash` is based on the transactions included in a block. In Lyfebloc, block data (including transactions) are erasure coded into a data square to enable data availability sampling. In order for the header to contain a commitment to this data square, `DataHash` was modified to be the Merkle root of the row and column roots of the erasure coded data square. See [lyfebloc-app/pkg/da/data_availability_header.go](https://github.com/lyfeblocnetwork/lyfebloc-app/blob/main/pkg/da/data_availability_header.go) for the implementation. Note on the implementation: lyfebloc-app computes the hash in prepare_proposal and adds it to txs via [abci_prepare_proposal.go](https://github.com/lyfeblocnetwork/lyfebloc-app/blob/ff16dfa025678fd5e87193ef31b82c76e4e50cad/app/abci_prepare_proposal.go#L122) so it is taken from txs and applied to [Data.Hash](https://github.com/lyfeblocnetwork/lyfebloc-core/blob/316e1051f14557805c8fc377fc1ee58c0664dc69/state/execution.go#L184) in lyfebloc-core.

<!-- See [./docs/celestia-architecture](./docs/celestia-architecture/) for architecture decision records (ADRs) on Celestia modifications. -->

## Diagram

```ascii
                ^  +-------------------------------+  ^
                |  |                               |  |
                |  |  State-machine = Application  |  |
                |  |                               |  |   lyfebloc-app (built with Cosmos SDK)
                |  |            ^      +           |  |
                |  +----------- | ABCI | ----------+  v
Lyfebloc         |  |            +      v           |  ^
validator or    |  |                               |  |
full consensus  |  |           Consensus           |  |
node            |  |                               |  |
                |  +-------------------------------+  |   lyfebloc-core (fork of CometBFT)
                |  |                               |  |
                |  |           Networking          |  |
                |  |                               |  |
                v  +-------------------------------+  v
```

## Install

See <https://github.com/lyfeblocnetwork/lyfebloc#install>

## Usage

See <https://github.com/lyfeblocnetwork/lyfebloc#usage>

## Contributing

This repo intends on preserving the minimal possible diff with [cometbft/cometbft](https://github.com/cometbft/cometbft) to make fetching upstream changes easy. If the proposed contribution is

- **specific to Lyfebloc**: consider if [lyfebloc-app](https://github.com/lyfeblocnetwork/lyfebloc-app) is a better target
- **not specific to Lyfebloc**: consider making the contribution upstream in CometBFT

1. [Install Go](https://go.dev/doc/install) 1.22+
2. Fork this repo
3. Clone your fork
4. Find an issue to work
5. Work on a change in a branch on your fork
6. When your change is ready, push your branch and create a PR that targets this repo

### Helpful Commands

```sh
# Build a new CometBFT binary and output to build/comet
make build

# Install CometBFT binary
make install

# Run tests
make test

# If you modified any protobuf definitions in a `*.proto` file then
# you may need to lint, format, and generate updated `*.pb.go` files
make proto-lint
make proto-format
make proto-gen
```

## Branches

The canonical branches in this repo are based on CometBFT releases. For example: [`v0.0.1-cmt-v0.38.2`](https://github.com/lyfeblocnetwork/lyfebloc-core/tree/v0.0.1-cmt-v0.38.2) is based on the CometBFT `v0.38.2` release branch and contains Lyfebloc-specific changes.

## Versioning

Releases are formatted: `v<LYFEBLOC_CORE_VERSION>-tm-v<COMETBFT_VERSION>`
For example: [`v0.0.1-cmt-v0.38.2`](https://github.com/lyfeblocnetwork/lyfebloc-core/releases/tag/v0.0.1-cmt-v0.38.2) is lyfebloc-core version `0.0.1` based on CometBFT `0.38.2`.
`LYFEBLOC_CORE_VERSION` strives to adhere to [Semantic Versioning](http://semver.org/).

## Careers

We are hiring Go engineers! Join us in building the future of blockchain scaling and interoperability. [Apply here](https://cauchye.notion.site/Backend-engineer-4822ce9ed78b456baa7e4dde41888434).
