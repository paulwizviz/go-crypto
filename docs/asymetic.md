# Asymmetric Cryptography

Asymmetric cryptography, also know as public key cryptography, uses a pair of keys: private and public for encryption and decryption. In this section we discuss asymmetric keys.

## Elliptic Curve

Elliptic Curve Cryptography (ECC) is a key-based technique for encrypting data. ECC focuses on pairs of public and private keys for decryption and encryption (see [here](https://avinetworks.com/glossary/elliptic-curve-cryptography/)).

### Working examples

* [Generating Elliptic curve keys using Go standard library](../examples/ec/ec_test.go)

### References

* [Elliptic Curves - Computerphile](https://www.youtube.com/watch?v=NF1pwjL9-DE)

## Rivest-Shamir-Adleman (RSA)

Please refe to [What Is RSA Algorithm and How Does It Work in Cryptography? by simplilearn](https://www.simplilearn.com/tutorials/cryptography-tutorial/rsa-algorithm).

There are two broad components when it comes to RSA cryptography, they are: 

* Key Generation for encrypting and decrypting the data to be exchanged.
* Encryption/Decryption Function.

### Working examples

You will find the following working examples:

* [Ex 1 - Generating key](../examples/rsa/ex1/main.go)

### References

* [The Mathematics of RSA and the Diffie-Hellman Protocol](https://www.youtube.com/watch?v=xmwxDHX6xUc)
* [RSA Algorithm - How does it work? - I'll PROVE it with an Example! -- Cryptography - Practical TLS](https://www.youtube.com/watch?v=Pq8gNbvfaoM)