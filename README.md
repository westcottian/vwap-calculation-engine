# vwap-calculation-engine
vwap-calculation-engine is a real-time VWAP (Volume-Weighted Average Price) calculation engine.

It uses [Coinbase's WS API](https://docs.cloud.coinbase.com/exchange/docs/websocket-overview) as data feed.

# Design

Communicates with [Coinbase Websocket Matches Channel](https://docs.cloud.coinbase.com/exchange/docs/websocket-channels#match), calculates VWAP of a number of Trading Pairs using set sliding window of trades.

Please refer to the [Design Doc](https://miro.com/app/board/uXjVOyTkRcE=/)

# Dependencies

- `Go` >= 1.17

- `Make` (optional for a better build experience)

- `Docker` (optional for a better deploy experience)

# Config
You can customize config properties located in `.env` according to the environment:

```
# Application
LOG_LEVEL=debug
SLIDING_WINDOW_SIZE=200
# Coinbase
COINBASE_SERVICE_ADDRESS=ws-feed.exchange.coinbase.com
COINBASE_PRODUCT_IDS=BTC-USD|ETH-USD|ETH-BTC
COINBASE_CHANNELS=matches
```

To create your own `.env` file:
```
make create-env
```

# Run

You can either:

```
make run
```

Or even (if you wish to run it as a standalone container):

```
make docker-run
```

You can also:

*Run in your favorite IDE or straight up: `go run cmd/main.go`*

# Test

Generate mocks:

```
make mock-generate
```

Run tests and generate coverage report:

```
make test
```

# Lint

Run lint:

```
make lint
```

# Tools

Download build tools:

```
make tools
```
Organize and format code:

```
make fmt
```

# CI

There are 2 GitHub Actions:
 - **Verify**: runs on every `push/pull request` and asserts `tests` and `lint`
 - **Release**: manual release using `semver`

Check them out: [CI Workflow](https://github.com/westcottian/vwap-calculation-engine/actions)

# Docs
You can find a more detailed overview regarding assumptions and decisions over [here](docs/).
