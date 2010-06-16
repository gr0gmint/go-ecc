package ecc

import "big"
import "fmt"

//FROM crypto/rsa.go ... We don't want to link the entire .a file.
var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)

// modInverse returns ia, the inverse of a in the multiplicative group of prime
// order n. It requires that a be a member of the group (i.e. less than n).
func modInverse(a, n *big.Int) (ia *big.Int, ok bool) {
    g := new(big.Int)
    x := new(big.Int)
    y := new(big.Int)
    big.GcdInt(g, x, y, a, n)
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

    return x, true
}





type Point struct {
    X, Y *big.Int
}
func NewPoint() *Point {
    p := new(Point)
    p.X = new(big.Int)
    p.Y = new(big.Int)
    return p
}

type Curve struct {
    
    P, A, B, N, H *big.Int
    G *Point
}

var BigZero = big.NewInt(0)
var BigTwo = big.NewInt(2)
var BigThree = big.NewInt(3)

func (curve *Curve) double(p *Point) *Point {
    fmt.Printf("d")
    lambda_numerator := new(big.Int)
    lambda_denominator := new(big.Int)
    lambda := new(big.Int)
    lambda_numerator.Exp(p.X, BigTwo, curve.P)
    lambda_numerator.Mul(lambda_numerator, BigThree)
    lambda_numerator.Add(lambda_numerator, curve.A)
    lambda_denominator.Mul(BigTwo, p.Y)
    lambda_denominator,ok := modInverse(lambda_denominator, curve.P)
    if !ok {
        fmt.Printf("Its not OKAY\n")
        return nil
    }
    lambda.Mul(lambda_numerator, lambda_denominator)
    lambda = lambda.Mod(lambda, curve.P)
    
    p3 := NewPoint()
   
    temp := new(big.Int)
    temp2 := new(big.Int)
    //temp3 := new(big.Int)
    //temp4 := new(big.Int)
    //temp5 := new(big.Int)
    //temp6 := new(big.Int)
    p3.X.Sub(temp.Exp(lambda, BigTwo, curve.P), temp2.Mul(BigTwo,p.X))
    p3.X = p3.X.Mod(p3.X, curve.P)
    
    p3.Y.Sub(p.X, p3.X)
    p3.Y.Mul(p3.Y, lambda)
    p3.Y = p3.Y.Sub(p3.Y, p.Y)
    p3.Y = p3.Y.Mod(p3.Y, curve.P)
    if p3.X.Cmp(BigZero) == -1 { //if X is negative
        p3.X.Neg(p3.X)
        p3.X.Sub(curve.P, p3.X)
    }
    if p3.Y.Cmp(BigZero) == -1 { //if Y is negative
        p3.Y.Neg(p3.Y)
        p3.Y.Sub(curve.P, p3.Y)
    }
    return p3
}

func (curve *Curve) Multiply(n *big.Int, p *Point) *Point {
    if p == nil {
        //fmt.Printf("p == nil!?wtfbbq\n")
        return p
    }


    bytes := n.Bytes()
    length := len(bytes)
    bitlength := length*8
        
    fmt.Printf("length = %d\n", bitlength)
    
    var rightmost uint = 0x01
    //fmt.Printf("leftmost = %d\n", leftmost)
    p2 := p
    last_i := bitlength -1
    var ptotal *Point
    ptotal=nil
    for i := bitlength - 1; i >= 0; i-- {
        //fmt.Printf("\n(i mod 8) = %d \n", 7-(i%8))
        if uint(rightmost << uint(7-(i % 8)))  & uint(bytes[i/8]) != 0 {
            for j:= last_i; j > i; j-- {
                //fmt.Printf("Doubling! i=%d\n",i)
                p2 = curve.double(p2)
            }
            last_i = i
            fmt.Printf("last_i = %d\n", last_i)
            ptotal = curve.Add(p2, ptotal)
        }
    }
    return ptotal
}
func (curve *Curve) Add(p1, p2 *Point) *Point {
    fmt.Printf("a")
    if p1 == nil {
        return p2
    }
    if p2 == nil {
        return p1
    }
    lambda_numerator := new(big.Int)
    lambda_denominator := new(big.Int)
    lambda := new(big.Int)
    lambda_numerator.Sub(p2.Y, p1.Y)
    lambda_denominator.Sub(p2.X, p1.X)
    if lambda_denominator.Cmp(BigZero) == -1 { //if Y is negative
        lambda_denominator.Neg(lambda_denominator)
        lambda_denominator.Sub(curve.P, lambda_denominator)
    }
    lambda_denominator, ok := modInverse(lambda_denominator, curve.P)
    if !ok {
        fmt.Printf("Add : Not ok\n")
        return nil
    }
    lambda.Mul(lambda_numerator, lambda_denominator)
    lambda = lambda.Mod(lambda, curve.P)
    
    p3 := NewPoint()
    p3.X.Exp(lambda, BigTwo, curve.P)
    p3.X.Sub(p3.X, p1.X)
    p3.X.Sub(p3.X, p2.X)
    p3.X = p3.X.Mod(p3.X, curve.P)
    
    p3.Y.Sub(p1.X,p3.X)
    p3.Y.Mul(lambda,p3.Y)
    p3.Y.Sub(p3.Y, p1.Y)
    p3.Y = p3.Y.Mod(p3.Y, curve.P)
    
    if p3.X.Cmp(BigZero) == -1 { //if X is negative
        p3.X.Neg(p3.X)
        p3.X.Sub(curve.P, p3.X)
    }
    if p3.Y.Cmp(BigZero) == -1 { //if Y is negative
        p3.Y.Neg(p3.Y)
        p3.Y.Sub(curve.P, p3.Y)
    }
    return p3
}
