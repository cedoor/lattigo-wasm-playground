# Lattigo WASM Demo

A WebAssembly demonstration of homomorphic encryption using the [Lattigo](https://github.com/tuneinsight/lattigo) library's CKKS scheme. Encrypts data in Go, performs computations on encrypted values, and runs entirely in the browser.

## 🚀 Try it Live

**[Live Demo: https://cedoor.github.io/lattigo-wasm-playground/](https://cedoor.github.io/lattigo-wasm-playground/)**

Click "Run" to see homomorphic encryption in action - encrypt numbers, perform addition on encrypted data, and decrypt the result, all in your browser!

## Quick Start

```bash
# Setup after cloning & start development server
make
```

Open `http://localhost:8080` and click "Run" to see homomorphic addition in action.

## What This Demo Shows

This demo demonstrates:
- **Homomorphic Encryption**: Perform calculations on encrypted data without decrypting it
- **CKKS Scheme**: Lattigo's implementation for approximate arithmetic on encrypted floating-point numbers
- **Go + WebAssembly**: Running Go cryptographic code directly in the browser
- **Real-time Computation**: Encrypt two arrays, add them while encrypted, then decrypt the result

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