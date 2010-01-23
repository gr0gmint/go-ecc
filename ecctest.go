package main
import "crypto/ecc"
import "fmt"
import "big"

func main() {
    c := ecc.NewSecp512r1()
    sixteen := big.NewInt(15)
    lol := c.Multiply(sixteen,c.G)
    fmt.Printf("%s\n\n%s\n\n", lol.X, lol.Y)
    lol = c.Multiply(big.NewInt(3), lol)
    fmt.Printf("%s\n\n%s\n\n", lol.X, lol.Y)
}
