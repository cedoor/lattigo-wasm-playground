# Lattigo WASM Demo

A WebAssembly demonstration of homomorphic encryption using the [Lattigo](https://github.com/tuneinsight/lattigo) library's CKKS scheme. Encrypts data in Go, performs computations on encrypted values, and runs entirely in the browser.

## Quick Start

```bash
# Setup after cloning & start development server
make
```

Open `http://localhost:8080` and click "Run" to see homomorphic addition in action.

## Commands

| Command | Description |
|---------|-------------|
| `make` | Build and serve |
| `make build` | Build WASM module |
| `make serve` | Start dev server on :8080 |
| `make dev` | Build and serve |
| `make clean` | Remove build artifacts |
| `make help` | Show all commands |

## Requirements

- Go 1.25.1+
- Python 3 (for dev server)
- Modern browser with WASM support