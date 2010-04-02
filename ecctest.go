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
    ka := big.NewInt(687687)
    kb := big.NewInt(8798451)
    kab := big.NewInt(8798451*687687) 
    Qa := c.Multiply(ka,c.G)
    Qb := c.Multiply(kb,c.G)
    Qab := c.Multiply(kab,c.G)
    shared1 := c.Multiply(ka,Qb)
    shared2 := c.Multiply(kb,Qa)
    //pointofinf := c.Multiply(c.N,c.G)
    fmt.Printf("Qa: %s       %s\n\n", Qa.X, Qa.Y)
    
    fmt.Printf("Qb: %s       %s\n\n", Qb.X, Qb.Y)
    
    fmt.Printf("Qab: %s       %s\n\n", Qab.X, Qab.Y)
    
    fmt.Printf("shared1: %s       %s\n\n", shared1.X, shared1.Y)
    
    fmt.Printf("shared2: %s       %s\n\n", shared2.X, shared2.Y)
    
    //fmt.Printf("pointofinf: %s       %s\n\n", pointofinf.X, pointofinf.Y)
    //ki,_ := modInverse(k, c.N)
    //lol2 := c.Multiply(ki, lol)
    //fmt.Printf("1/k = %s\n", ki)
    //fmt.Printf("(k*G)/k: %s     %s\n\n", lol2.X, lol2.Y)
    
}
