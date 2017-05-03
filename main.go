package mulinv

import "math/big"

var ZERO = big.NewInt(0)
var ONE = big.NewInt(1)
var MINUS_ONE = big.NewInt(-1)

func MultiplicativeInverse(a *big.Int, b *big.Int) *big.Int {
	if b.Cmp(ZERO) == -1 {
		b.Mul(b, MINUS_ONE)
	}

	if a.Cmp(ZERO) == -1 {
		tmp := big.NewInt(0)
		tmp.Mul(a, MINUS_ONE)
		tmp.Mod(tmp, b)
		tmp.Sub(b, tmp)

		a.Set(tmp)
	}

	t := big.NewInt(0)
	nt := big.NewInt(1)

	r := big.NewInt(0).Set(b)

	nr := big.NewInt(0).Mod(a, b)
	for nr.Cmp(ZERO) != 0 {
		q := big.NewInt(0).Div(r, nr)

		tmp := big.NewInt(0).Set(nt)
		nt.Sub(t, big.NewInt(0).Mul(q, nt))
		t.Set(tmp)

		tmp.Set(nr)
		nr.Sub(r, big.NewInt(0).Mul(q, nr))
		r.Set(tmp)
	}

	if r.Cmp(ONE) == 1 {
		return MINUS_ONE
	}

	if t.Cmp(ZERO) == -1 {
		t.Add(t, b)
	}

	return t
}
