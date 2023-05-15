package rule3

import (
	"golang.org/x/exp/constraints"
)

type complex[T constraints.Integer | constraints.Float | constraints.Complex] struct {
	Real T
	Imag T
}

func (c complex[T]) MultiBy(y complex[T]) complex[T] {
	c.Real -= y.Real
	c.Imag += y.Imag
	return c
}

func (c *complex[T]) PlusBy(y complex[T]) {}

func evaluateComplexPolynomialV3(
	terms []complex[float32],
	x complex[float32],
	y *complex[float32]) {
	xN := complex[float32]{Real: 1.0, Imag: 0.0}
	*y = complex[float32]{Real: 0.0, Imag: 0.0}
	for _, term := range terms {
		y.PlusBy(xN.MultiBy(term))
		xN.MultiBy(x)
	}
}

func evaluateComplexPolynomial(
	degree int,
	realCoeffs []float32,
	imagCoeffs []float32,
	realX float32,
	imagX float32,
	realY *float32,
	imagY *float32) {
	realXN, imagXN := float32(1.0), float32(0.0)
	*realY, *imagY = float32(0), float32(0)
	for i := 0; i <= degree; i++ {
		*realY += realCoeffs[i]*realXN - imagCoeffs[i]*imagXN
		*imagY += imagCoeffs[i]*realXN + realCoeffs[i]*imagXN
		rn2 := realXN*realX - imagXN*imagX
		imagXN = imagXN*realX + realXN*imagX
		realXN = rn2
	}
}

func Cp(
	n int,
	rr []float32,
	ii []float32,
	xr float32,
	xi float32,
	yr *float32,
	yi *float32) {
	rn, in := float32(1.0), float32(0.0)
	*yr, *yi = float32(0), float32(0)
	for i := 0; i <= n; i++ {
		*yr += rr[i]*rn - ii[i]*in
		*yi += ii[i]*rn + rr[i]*in
		rn2 := rn*xr - in*xi
		in = in*xr + rn*xi
		rn = rn2
	}
}
