package ecc
import "big"
import "bytes"


const (
    a = "\x01\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFC";
    b = "\x00\x51\x95\x3E\xB9\x61\x8E\x1C\x9A\x1F\x92\x9A\x21\xA0\xB6\x85\x40\xEE\xA2\xDA\x72\x5B\x99\xB3\x15\xF3\xB8\xB4\x89\x91\x8E\xF1\x09\xE1\x56\x19\x39\x51\xEC\x7E\x93\x7B\x16\x52\xC0\xBD\x3B\xB1\xBF0\x73\x57\x3D\xF8\x83\xD2C\x34\xF1\xEF\x45\x1F\xD4\x6B\x50\x3F\x00";
    g = "\x04\x00\xC6\x85\x8E\x06\xB7\x04\x04\xE9\xCD\x9E\x3E\xCB\x66\x23\x95\xB4\x42\x9C\x64\x81\x39\x05\x3F\xB5\x21\xF8\x28\xAF\x60\x6B\x4D\x3D\xBA\xA1\x4B\x5E\x77\xEF\xE7\x59\x28\xFE\x1D\xC1\x27\xA2\xFF\xA8\xDE\x33\x48\xB3\xC1\x85\x6A\x42\x9B\xF9\x7E\x7E\x31\xC2\xE5\xBD\x66\x01\x18\x39\x29\x6A\x78\x9A\x3B\xC0\x04\x5C\x8A\x5F\xB4\x2C\x7D\x1B\xD9\x98\xF5\x44\x49\x57\x9B\x44\x68\x17\xAF\xBD\x17\x27\x3E\x66\x2C\x97\xEE\x72\x99\x5E\xF4\x26\x40\xC5\x50\xB9\x01\x3F\xAD\x07\x61\x35\x3C\x70\x86\xA2\x72\xC2\x40\x88\xBE\x94\x76\x9F\xD1\x66\x50";
    n = "\x01\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFA\x51\x86\x87\x83\xBF\x2F\x96\x6B\x7F\xCC\x01\x48\xF7\x09\xA5\xD0\x3B\xB5\xC9\xB8\x89\x9C\x47\xAE\xBB\x6F\xB7\x1E\x91\x38\x64\x09";
    h= 0x01;
)




func NewSecp512r1() *Curve {
    curve := new(Curve);
    curve.P = new(big.Int);
    curve.A = new(big.Int);
    curve.B = new(big.Int);
    curve.N = new(big.Int);
    curve.H = big.NewInt(h);
    curve.G = OctetToPoint(bytes.NewBufferString(g).Bytes());
    
    curve.P.Exp(big.NewInt(2),big.NewInt(521), nil);
    curve.P.Sub(curve.P, big.NewInt(1));
    curve.A.SetBytes(bytes.NewBufferString(a).Bytes());
    curve.B.SetBytes(bytes.NewBufferString(b).Bytes());
    //curve.G.SetBytes(NewBufferString(g).Bytes());
    curve.N.SetBytes(bytes.NewBufferString(n).Bytes());
    return curve;
    
}