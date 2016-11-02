





package main

import (
  "fmt"
  "math/big"
)

func main() {

  fmt.Println("Please enter the first number: ")
  var a big.Int
  fmt.Scan(&a)
  fmt.Println("Please enter the second number: ")
  var b big.Int
  fmt.Scan(&b)

  if a.Sign() != -1 && b.Sign() != -1 && (a.Sign() != 0 || b.Sign() != 0) {
  //if a >= 0 && b >= 0 && ( a != 0 || b != 0) {
    d, u, v := extendedEslid(a,b)
    fmt.Println("The gcd is: ")
    fmt.Println(d)
    fmt.Println(" with u: ")
    fmt.Println(u)
    fmt.Println(" and v: ")
    fmt.Println(v)
  } else {
    fmt.Println("invalid input")
  }
}

func extendedEslid(a big.Int, b big.Int) (d *big.Int, ud *big.Int, vd *big.Int) {
  c := &a
  d = &b

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
