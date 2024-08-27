# Elliptic Curve

Refer to [here for background information](https://github.com/paulwizviz/system-engineering.git)

There are two implementations of EC:

* Go standard packages (http://golang.org/pkg/crypto/elliptic)
* `secp256k1` implementation

The differences between these implementations are:

* `Go Standard Library`:
    * The `crypto/elliptic` package supports several standard elliptic curves such as P-224, P-256, P-384, and P-521. These curves are NIST (National Institute of Standards and Technology) recommended and are different from secp256k1 in terms of their parameters and security considerations.
    * This package is suitable for applications requiring general ECC support with NIST-approved curves, such as securing communications (TLS), data encryption, and general-purpose cryptography.
    * The package is generic and focuses on supporting a range of curves. It doesn't include specific optimizations like endomorphism that are useful in certain curves like secp256k1.
    * The NIST curves are defined by specific parameters (prime fields, curve equation coefficients) chosen for general security considerations. For instance, P-256 (also known as secp256r1) uses the curve equation y^2 = x^3 - 3x + b.
* `secp256k1`: 
    * This curve is specifically used in cryptocurrency applications like Bitcoin, Ethereum, and others. The specific curve and its implementations are chosen for their compatibility and performance benefits in blockchain-related operations.
    * Implementations often include specific optimizations such as the GLV endomorphism, which can speed up scalar multiplication. Additionally, implementations like libsecp256k1 include safeguards against side-channel attacks and provide efficient batch verification of signatures.
    * This curve is specifically used in cryptocurrency applications like Bitcoin, Ethereum, and others. The specific curve and its implementations are chosen for their compatibility and performance benefits in blockchain-related operations.
    * This curve has a different set of parameters and uses the equation y^2 = x^3 + 7, which is known as a Koblitz curve. It has specific properties that make it more efficient for certain operations (e.g., ECDSA signature verification) but potentially less secure against specific theoretical attacks.

* Go Bindings for libsecp256k1:
    * `btcsuite/btcd/secp256k1`: Part of the btcsuite project, this Go package provides a binding to `libsecp256k1` and is commonly used in Bitcoin-related Go projects.
    * `decred/dcrd/dcrec/secp256k1`: Another Go binding, used by the Decred project, which also wraps around `libsecp256k1`
    * `ethereum/go-ethereum/crypto/secp256k1`: The Ethereum Go client (go-ethereum) uses this binding for `secp256k1` operations.

## Key Generation

* [Example 1](./ex1/main.go)

## Digital Signature

Refer to [Example 2](./ex2/main.go)