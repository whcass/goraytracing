package main

type Ray struct {
	A, B Vector
}

func (r Ray) origin() Vector {
	return r.A
}

func (r Ray) direction() Vector {
	return r.B
}

func (r Ray) pointAt(t float64) Vector {
	return r.A.Add(r.B.MultiplyByScalar(t))
}
