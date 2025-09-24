// CKKS homomorphic encryption demo using Lattigo in WebAssembly
package main

import (
	"encoding/base64"
	"syscall/js"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

var (
	params    ckks.Parameters
	encoder   *ckks.Encoder
	sk        *rlwe.SecretKey
	pk        *rlwe.PublicKey
	encryptor *rlwe.Encryptor
	decryptor *rlwe.Decryptor
	evaluator *ckks.Evaluator
)

func initCKKS(this js.Value, args []js.Value) any {
	// CKKS parameters: small ring for demo
	params, _ = ckks.NewParametersFromLiteral(ckks.ParametersLiteral{
		LogN:            12,             // 2^12 = 4096 slots
		LogQ:            []int{55, 45, 45},
		LogP:            []int{61},
		LogDefaultScale: 45,
	})

	// Setup encoding and keys
	encoder = ckks.NewEncoder(params)
	keygen := rlwe.NewKeyGenerator(params)
	sk = keygen.GenSecretKeyNew()
	pk = keygen.GenPublicKeyNew(sk)
	rlk := keygen.GenRelinearizationKeyNew(sk)

	// Setup crypto operations
	encryptor = rlwe.NewEncryptor(params, pk)
	decryptor = rlwe.NewDecryptor(params, sk)
	evaluator = ckks.NewEvaluator(params, rlwe.NewMemEvaluationKeySet(rlk))

	return js.Undefined()
}

func encrypt(this js.Value, args []js.Value) any {
	// Convert JS Float64Array to complex numbers
	jsArr := args[0]
	n := jsArr.Length()
	vals := make([]complex128, n)
	for i := 0; i < n; i++ {
		vals[i] = complex(jsArr.Index(i).Float(), 0)
	}

	// Encode and encrypt
	pt := ckks.NewPlaintext(params, params.MaxLevel())
	encoder.Encode(vals, pt)
	ct, _ := encryptor.EncryptNew(pt)

	// Return as base64 for JS
	data, _ := ct.MarshalBinary()
	return js.ValueOf(base64.StdEncoding.EncodeToString(data))
}

func evalAdd(this js.Value, args []js.Value) any {
	// Deserialize base64 ciphertexts
	aBytes, _ := base64.StdEncoding.DecodeString(args[0].String())
	bBytes, _ := base64.StdEncoding.DecodeString(args[1].String())

	ctA := new(rlwe.Ciphertext)
	ctB := new(rlwe.Ciphertext)
	ctA.UnmarshalBinary(aBytes)
	ctB.UnmarshalBinary(bBytes)

	// Homomorphic addition
	ctC, _ := evaluator.AddNew(ctA, ctB)

	// Serialize result
	out, _ := ctC.MarshalBinary()
	return js.ValueOf(base64.StdEncoding.EncodeToString(out))
}

func decrypt(this js.Value, args []js.Value) any {
	// Deserialize ciphertext
	raw, _ := base64.StdEncoding.DecodeString(args[0].String())
	ct := new(rlwe.Ciphertext)
	ct.UnmarshalBinary(raw)

	// Decrypt and decode
	pt := decryptor.DecryptNew(ct)
	res := make([]complex128, 1<<params.LogMaxSlots())
	encoder.Decode(pt, res)

	// Return first 8 real values as Float64Array
	out := js.Global().Get("Float64Array").New(8)
	for i := 0; i < 8; i++ {
		out.SetIndex(i, real(res[i]))
	}
	return out
}

func main() {
	// Export functions to JavaScript
	js.Global().Set("ckksInit", js.FuncOf(initCKKS))
	js.Global().Set("ckksEncrypt", js.FuncOf(encrypt))
	js.Global().Set("ckksEvalAdd", js.FuncOf(evalAdd))
	js.Global().Set("ckksDecrypt", js.FuncOf(decrypt))
	select {}
}
