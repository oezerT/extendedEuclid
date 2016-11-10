package main

import (
  "fmt"
  "math/big"
)

func main() {

  fmt.Println("Please enter the value for a: ")
  var a big.Int
  fmt.Scan(&a)
  fmt.Println("Please enter the value for b: ")
  var b big.Int
  fmt.Scan(&b)
  fmt.Println("Please enter the value for p: ")
  var p big.Int
  fmt.Scan(&p)
  fmt.Println("Please enter the value for q: ")
  var q big.Int
  fmt.Scan(&q)

  if p.Sign() != -1 && q.Sign() != -1 && (p.Sign() != 0 || q.Sign() != 0) {
    _, z, _ := extendedEuclid(&p, &q)
    fmt.Println("1. z = ", z)
    y := big.NewInt(0).Mod(big.NewInt(0).Mul(big.NewInt(0).Sub(&a, &b), z), &p)
    fmt.Println("2. y = ", y)
    x := big.NewInt(0).Add(big.NewInt(0).Mul(y, &q), &b)
    fmt.Println("3. x = ", x)
  } else {
    fmt.Println("invalid input")
  }
}

func extendedEuclid(a *big.Int, b *big.Int) (d *big.Int, ud *big.Int, vd *big.Int) {
  c := a
  d = b

  if a.Cmp(b) < 0 {
    c = b
    d = a
  }

  uc := big.NewInt(1)
  vc := big.NewInt(0)
  ud = big.NewInt(0)
  vd = big.NewInt(1)

  for c.Sign() != 0 {
    q := big.NewInt(0).Div(d,c)
    cOld := c
    c = big.NewInt(0).Sub(d, big.NewInt(0).Mul(q,c))
    d = cOld
    ucOld := uc
    vcOld := vc

    uc = big.NewInt(0).Sub(ud, big.NewInt(0).Mul(q, uc))
    vc = big.NewInt(0).Sub(vd, big.NewInt(0).Mul(q, vc))

    ud = ucOld
    vd = vcOld
   }
  return
}
