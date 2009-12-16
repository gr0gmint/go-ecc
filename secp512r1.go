package "ecc"
import "big"
import "math"
import "./ecc.go"

const (
    a = '\x01FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC";
    b = 0x0051953EB9618E1C9A1F929A21A0B68540EEA2DA725B99B315F3B8B489918EF109E156193951EC7E937B1652C0BD3BB1BF073573DF883D2C34F1EF451FD46B503F00;
    G = 0x0200c6858e06b70404e9cd9e3ecb662395b4429c648139053fb521f828af606b4d3dbaa14b5e77efe75928fe1dc127a2ffa8de3348b3c1856a429bf97e7e31c2e5bd66;
    n = 0x01FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFA51868783BF2F966B7FCC0148F709A5D03BB5C9B8899C47AEBB6FB71E91386409;
    h= 0x01;
)



func NewSecp512r1() *Curve {
    curve := new(Secp512r1);
    curve.p = new(big.Int);
    curve.a = new(big.Int);
    curve.b = new(big.Int);
    curve.xG = new(big.Int);
    curve.yG = new(big.Int);
    curve.n = new(big.Int);
    curve.h = big.NewInt(1);
    
    curve.p.Exp(NewInt(2),NewInt(521)); //subtract with 1
    curve.a.SetBytes(a);
    curve.b.SetBytes(b);
    curve.G.SetBytes(G);
    curve.n.SetBytes(n);
    return curve;
    
}

/*
    p = 

*/
