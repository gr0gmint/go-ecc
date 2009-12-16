package "ecc"

import "./secp512r1.go"

import "big"



type Point struct {
    x, y *big.Int;
}
type Curve struct {
    
    P, A, B, N, H *big.Int;
    G *Point;
}

func (curve *Curve) Multiply(n *big.Int, p *Point) *Point {
}
func (curve *Curve) Add(p1, p2 *Point) *Point {
    
}
