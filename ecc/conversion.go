package ecc

//import "big"

func OctetToPoint(octet []byte) *Point {
    if octet[0] == '\x04' {
        p := NewPoint()
        p.X.SetBytes(octet[1:((len(octet)-2)>>1)])
        p.Y.SetBytes(octet[(len(octet)-1)/2:])
        return p
    }
    return nil
}
