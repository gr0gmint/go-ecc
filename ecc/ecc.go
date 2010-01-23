package ecc

import "big"
import "fmt"

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
    X, Y *big.Int;
}
func NewPoint() *Point {
    p := new(Point);
    p.X = new(big.Int)
    p.Y = new(big.Int)
    return p;
}

type Curve struct {
    
    P, A, B, N, H *big.Int;
    G *Point;
}

var BigTwo = big.NewInt(2);
var BigThree = big.NewInt(3);

func (curve *Curve) double(p *Point) *Point {
    lambda_numerator := new(big.Int)
    lambda_denominator := new(big.Int)
    lambda := new(big.Int)
    lambda_numerator.Exp(p.X, BigTwo, curve.P)
    lambda_numerator.Mul(lambda_numerator, BigThree)
    lambda_numerator.Sub(lambda_numerator, curve.A)
    lambda_denominator.Mul(BigTwo, p.Y)
    lambda_denominator,_ = modInverse(lambda_denominator, curve.P)
    lambda.Mul(lambda_numerator, lambda_denominator)
    lambda = lambda.Mod(lambda, curve.P)
    
    p3 := new(Point);
    p3.Y.Exp(lambda, BigTwo, nil);
    
    temp := new(big.Int);
    p3.X.Sub(p3.X, temp.Mul(BigTwo,p.X));
    p3.X = p3.X.Mod(p3.X, curve.P);
    
    p3.Y.Sub(p.X, p3.X);
    p3.Y.Mul(p3.Y, lambda);
    p3.Y = p3.Y.Mod(p3.Y, curve.P);
    return p3;
}

func (curve *Curve) Multiply(n *big.Int, p *Point) *Point {
    if p == nil {
        fmt.Printf("p == nil!?wtfbbq\n")
        return p
    }


    bytes := n.Bytes()
    
    length := len(bytes);
        bitlength := length*8
    fmt.Printf("length = %d\n", length)
    
    var leftmost int = 0x01 << 8
    for i := 0; i < bitlength; i++ {
        fmt.Printf("int(leftmost >> uint(%d mod 8)) = %d\n",i, int(leftmost >> uint(i % 8)))
        fmt.Printf("AND'ed with int(bytes[%d/8]) = %d\n",i, int(leftmost >> uint(i % 8))  & int(bytes[i/8]))
        if int(leftmost >> uint(i % 8))  & int(bytes[i/8]) != 0x00000000 {
            for j:= 0; j < bitlength - i; j++ {
                fmt.Printf("Doubling that shit\n")
                p = curve.double(p)
            }
        }
    }
    return p
}
func (curve *Curve) Add(p1, p2 *Point) *Point {
    if p1 == nil {
        return p2;
    }
    if p2 == nil {
        return p1;
    }
    lambda_numerator := new(big.Int);
    lambda_denominator := new(big.Int);
    lambda := new(big.Int);
    lambda_numerator.Sub(p2.Y, p1.Y);
    lambda_denominator.Sub(p2.X, p1.X);
    lambda_denominator, _ = modInverse(lambda_denominator, curve.P);
    lambda.Mul(lambda_numerator, lambda_denominator);
    lambda = lambda.Mod(lambda, curve.P);
    
    p3 := NewPoint();
    p3.X.Exp(lambda, BigTwo, curve.P);
    p3.X.Sub(p3.X, p1.X);
    p3.X.Sub(p3.X, p2.X);
    p3.X = p3.X.Mod(p3.X, curve.P);
    
    p3.Y.Sub(p1.X,p3.X);
    p3.Y.Mul(lambda,p3.Y);
    p3.Y.Sub(p3.Y, p1.Y);
    p3.Y = p3.Y.Mod(p3.Y, curve.P);
    
    return p3;
}
