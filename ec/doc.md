# Elliptic Curve

Refer to [here for background information](https://github.com/paulwizviz/system-engineering.git)

There are two implementations of EC:

* Go standard packages (http://golang.org/pkg/crypto/elliptic)
* `secp256k1` implementation

The differences between these implementations are:

* [Standard Library](#standard-package)
* [Bindings for `libsecp256k1`](#binding-for-secp256k1)

## Standard Package

* The `crypto/elliptic` package supports several standard NIST (National Institute of Standards and Technology) elliptic curves such as P-224, P-256, P-384, and P-521.
* This package is suitable for applications requiring general ECC support with NIST-approved curves, such as securing communications (TLS), data encryption, and general-purpose cryptography.
* The package is generic and focuses on supporting a range of curves. It doesn't include specific optimizations like endomorphism that are useful in certain curves like secp256k1.
* The NIST curves are defined by specific parameters (prime fields, curve equation coefficients) chosen for general security considerations. For instance, P-256 (also known as secp256r1) uses the curve equation y^2 = x^3 - 3x + b.

[Generating key and signing message](./stdlib/stdlib_test.go)

## Binding for `secp256k1`

 * `btcsuite/btcd/secp256k1`: Part of the btcsuite project, this Go package provides a binding to `libsecp256k1` and is commonly used in Bitcoin-related Go projects.
* `decred/dcrd/dcrec/secp256k1`: Another Go binding, used by the Decred project, which also wraps around `libsecp256k1`
* `ethereum/go-ethereum/crypto/secp256k1`: The Ethereum Go client (go-ethereum) uses this binding for `secp256k1` operations.

[Working Example](./secp/secp_test.go)