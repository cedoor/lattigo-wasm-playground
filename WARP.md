# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

This is a WebAssembly (WASM) demonstration of the Lattigo homomorphic encryption library. The project shows how to use Lattigo's CKKS scheme (for approximate arithmetic over encrypted data) in a web browser through Go's WebAssembly compilation target.

### Architecture

The project follows a simple client-server architecture where:

- **Go WASM Module** (`go/main.go`): Compiles to WebAssembly and exposes JavaScript functions for CKKS operations
- **Web Frontend** (`web/index.html`): Browser interface that loads and interacts with the WASM module
- **Bridge Layer** (`web/wasm_exec.js`): Go's standard WebAssembly bridge for JavaScript interop

### Core Components

**WASM Module Functions (exposed to JavaScript):**
- `ckksInit()`: Initializes CKKS parameters and cryptographic keys
- `ckksEncrypt(Float64Array)`: Encrypts an array of numbers, returns base64-encoded ciphertext
- `ckksEvalAdd(ciphertext1, ciphertext2)`: Performs homomorphic addition on encrypted data
- `ckksDecrypt(ciphertext)`: Decrypts and returns first 8 slots as Float64Array

**Key Architecture Notes:**
- Uses Lattigo v6 for homomorphic encryption operations
- CKKS parameters: LogN=12 (4096 slots), 3 moduli, 45-bit default scale
- Ciphertexts are serialized to base64 for JavaScript transport
- Go's `syscall/js` package bridges Go functions to JavaScript
- Uses unsafe operations for efficient float64/uint64 conversions

## Development Commands

### Build WASM Module
```bash
GOOS=js GOARCH=wasm go build -o web/main.wasm go/main.go
```

### Get Go WASM Support Files
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
```

### Development Server
```bash
# Using Python 3
cd web && python3 -m http.server 8080

# Using Go (alternative)
go run -tags dev ./cmd/server  # if you add a dev server later

# Using npx (if Node.js available)
cd web && npx http-server -p 8080
```

### Module Management
```bash
# Update dependencies
go mod tidy

# Verify dependencies
go mod verify

# Download dependencies
go mod download
```

### Testing Build
```bash
# Verify WASM builds successfully
GOOS=js GOARCH=wasm go build -o /tmp/test.wasm go/main.go && echo "Build OK" || echo "Build failed"
```

## Development Workflow

1. **Making Changes to Go Code**: Edit `go/main.go`, then rebuild with the WASM build command above
2. **Testing**: Start development server and navigate to `http://localhost:8080` to test in browser
3. **Debugging WASM**: Use browser developer tools console to see Go runtime output and JavaScript interactions
4. **Adding New Functions**: Expose new Go functions via `js.Global().Set("functionName", js.FuncOf(yourFunction))`

## Technical Constraints

- **Memory Management**: WASM module has limited memory; large ciphertext operations may cause issues
- **Serialization**: All data exchange between Go and JavaScript happens through JSON-compatible types or binary data
- **Browser Compatibility**: Requires modern browsers with WebAssembly and crypto.getRandomValues support
- **CKKS Precision**: Approximate arithmetic scheme means results have small numerical errors
- **No File I/O**: WASM environment cannot access local files directly

## Dependencies

- **Lattigo v6**: Homomorphic encryption library - handles all cryptographic operations
- **Go 1.25.1+**: Required for the WebAssembly compilation target used
- **Modern Web Browser**: Chrome 69+, Firefox 62+, Safari 14+ for full WASM support

## Key Files

- `go/main.go`: Main WASM application logic and JavaScript function bindings
- `web/index.html`: Demo interface showing encrypt-add-decrypt workflow
- `web/wasm_exec.js`: Standard Go WebAssembly runtime (from Go toolchain)
- `go.mod`: Module dependencies, pinned to Lattigo v6.1.1