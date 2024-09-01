package stdlib

import "fmt"

func Example_pass() {

	msg := "Hello world"

	privkey, pubkey, err := generateKey()
	if err != nil {
		fmt.Printf("Generating key failed: %v", err)
		return
	}

	// s value is the x component of the curve y^2 = x^3 - 3x + b
	r, s, signmsg, err := signMessage(msg, privkey)
	if err != nil {
		fmt.Printf("Signed message: %v", err)
		return
	}

	status := verifyMessage(r, s, pubkey, signmsg)
	fmt.Println(status)

	// Output:
	// true

}

func Example_failed() {
	msg := "Hello world"

	privkey, _, err := generateKey()
	if err != nil {
		fmt.Printf("Generating private key failed: %v", err)
		return
	}

	r, s, signmsg, err := signMessage(msg, privkey)
	if err != nil {
		fmt.Printf("Signed message: %v", err)
		return
	}

	// Public key not from private key
	_, pubkey, err := generateKey()
	if err != nil {
		fmt.Printf("Generating public key failed: %v", err)
		return
	}

	status := verifyMessage(r, s, pubkey, signmsg)
	fmt.Println(status)

	// Output:
	// false
}
