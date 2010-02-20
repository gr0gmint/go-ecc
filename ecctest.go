package main


import (
    "fmt"
  "big"
     "crypto/ecc"
     )
     
//FROM crypto/rsa.go ... We don't want to link the entire .a file.
var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)
// modInverse returns ia, the inverse of a in the multiplicative group of prime
// order n. It requires that a be a member of the group (i.e. less than n).
func modInverse(a, n *big.Int) (ia *big.Int, ok bool) {
    g := new(big.Int);
    x := new(big.Int);
    y := new(big.Int);
    big.GcdInt(g, x, y, a, n);
    if g.Cmp(bigOne) != 0 {
        // In this case, a and n aren't coprime and we cannot calculate
        // the inverse. This happens because the values of n are nearly
        // prime (being the product of two primes) rather than truly
        // prime.
        return
    }

    if x.Cmp(bigOne) < 0 {
        // 0 is not the multiplicative inverse of any element so, if x
        // < 1, then x is negative.
        x.Add(x, n)
    }

    return x, true;
}




func main() {
    c := ecc.NewSecp512r1()
    k := big.NewInt(255) 
    lol := c.Multiply(k,c.G)
    fmt.Printf("G: %s       %s\n\n", c.G.X, c.G.Y)
    fmt.Printf("k*G: %s       %s\n\n", lol.X, lol.Y)
    //ki,_ := modInverse(k, c.N)
    //lol2 := c.Multiply(ki, lol)
    //fmt.Printf("1/k = %s\n", ki)
    //fmt.Printf("(k*G)/k: %s     %s\n\n", lol2.X, lol2.Y)
    
}
