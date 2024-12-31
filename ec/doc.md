# Elliptic Curve

Refer to [here for background information](https://github.com/paulwizviz/system-engineering.git)

There are two implementations of EC:

* Go standard packages (http://golang.org/pkg/crypto/elliptic)
* `secp256k1` implementation

The differences between these implementations are:

* [Standard Library](#standard-package)
* [Bindings for `libsecp256k1`](#binding-for-secp256k1)

## Standard Library vs `secp256k1`

| | Standard Lib | `secp256k1`  |
| --- | ---- | --- |
| **Feature** | `crypto/ecdsa` (Standard Go)  | `secp256k1` Binding (e.g., `decred/dcrd` or `go-ethereum/crypto`)   |
|  **Curve Support** | Supports standard NIST curves (P224, P256, P384, P521) but not `secp256k1` directly.  | Specifically designed for and optimized for the `secp256k1` curve. |
| **Performance** | General-purpose ECDSA implementation. Performance can be good but not optimized for `secp256k1` | Highly optimized for `secp256k1` operations. Usually significantly faster for this specific curve. |
| **Security** | Uses Go's standard cryptographic primitives. Considered secure for supported curves. | Often uses highly optimized assembly code for performance, which can have security implications if not thoroughly vetted. However, well-maintained libraries are generally considered secure. |
| **API** | Provides a standard `ecdsa.PrivateKey` and `ecdsa.PublicKey` structure and functions for signing and verifying. | Often provides more specialized functions for `secp256k1` (e.g., direct access to scalar multiplication, point operations). |
| **Dependencies** | Part of the Go standard library. No external dependencies. | Requires an external dependency. |
| **ASN.1 Encoding** | Supports standard ASN.1 encoding for private keys (PKCS#8) for its supported curves. Does not directly support ASN.1 encoding for `secp256k1`. | May provide custom encoding/decoding functions or handle ASN.1 in a way specific to `secp256k1` (though often just raw bytes are used). |
| **Use Cases** | General ECDSA operations with standard NIST curves. Suitable when `secp256k1` is not required or performance is not critical. | Essential for applications that heavily use `secp256k1`, like cryptocurrencies (Bitcoin, Ethereum) and other blockchain technologies. |
| **Example Usage** | ecdsa.GenerateKey(`elliptic.P256`, `rand.Reader`), `ecdsa.Sign`, `ecdsa.Verify` | `secp256k1.GeneratePrivateKey()`, `secp256k1.Sign`, `secp256k1.Verify` (or similar, depending on the library) |

## Working Examples

* Standard lib
    * [Example 1 - Generate Private key](../ec/stdlib/ex1/main.go)
    * [Example 2 - Verify signed message](../ec/stdlib/ex2/main.go)
    * [Example 3 - Verify with wrong public key](../ec/stdlib/ex3/main.go)
    * [Example 4 - Encode private key in PEM](../ec/stdlib/ex4/main.go)
    * [Example 5 - Encode and encrypt private key in PEM](../ec/stdlib/ex5/main.go)
