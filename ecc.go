package ecc

import "./secp512r1.go"
import "./conversion.go"

import "big"


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





type Point struct {
    x, y *big.Int;
}
func NewPoint() *Point {
    p := new(Point);
    p.x,p.y = new(big.Int);
}

type Curve struct {
    
    P, A, B, N, H *big.Int;
    G *Point;
}

var BigTwo = big.NewInt(2);
var BigThree = big.NewInt(3);

func (curve *Curve) double(p *Point) *Point {
    lamda_numerator = new(big.Int);
    lambda_denominator = new(big.Int);
    lambda = new(big.Int);
    lambda_numerator.Exp(p.x, BigTwo, curve.P);
    lambda_numerator.Mul(lambda_numerator, BigThree);
    lambda_numerator.Sub(lambda_numerator, curve.A);
    lambda_denominator.Mul(BigTwo, p.y);
    lambda_denominator = modInverse(lambda_denominator, curve.P);
    lambda.Mul(lambda_numerator, lambda_denominator);
    lambda = lambda.Mod(lambda, curve.P);
    
    p3 := new(Point);
    p3.x.Exp(lambda, BigTwo);
    
    temp := new(big.Int);
    p3.x.Sub(p3.x, temp.Mul(BigTwo,p.x));
    p3.x = p3.x.Mod(p3.x, curve.P);
    
    p3.y.Sub(p.x, p3.x);
    p3.y.Mul(p3.y, lambda);
    p3.y = p3.y.Mod(p3.y, curve.P);
    return p3;
}

func (curve *Curve) Multiply(n *big.Int, p *Point) *Point {
    if p == nil {
        return p
    }
    bitlength := n.Len()
    bytes := n.Bytes()
    length := len(bytes);
    leftmost := 0x01 << 7
    
    for i := 0; i < length; i++ {
        if ((leftmost >> (i % 8) ) & int(bytes[i/8])) {
            for j:= 0; j < bitlength - i; j++ {
                
            }
        }
    }
}
func (curve *Curve) Add(p1, p2 *Point) *Point {
    if p1 == nil {
        return p2;
    }
    if p2 == nil {
        return p1;
    }
    lamda_numerator = new(big.Int);
    lambda_denominator = new(big.Int);
    lambda = new(big.Int);
    lambda_numerator.Sub(p2.y, p1.y);
    lambda_denominator.Sub(p2.x, p1.x);
    lambda_denominator, _ = modInverse(lambda_denominator, curve.P);
    lambda.Mul(lambda_numerator, lambda_denominator);
    lambda = lambda.Mod(lambda, curve.P);
    
    p3 := NewPoint();
    p3.x.Exp(lambda, BigTwo, curve.P);
    p3.x.Sub(p3.x, p1.x);
    p3.x.Sub(p3.x, p2.x);
    p3.x = p3.x.Mod(p3.x, curve.P);
    
    p3.y.Sub(p1.x,p3.x);
    p3.y.Mul(lambda,p3.y);
    p3.y.Sub(p3.y, p1.y);
    p3.y = p3.y.Mod(p3.y, curve.P);
    
    return p3;
}
