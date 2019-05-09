package main

import "math"

type Sphere struct {
	Center Vector
	r      float64
}

func (s Sphere) hitSphere(r Ray, min float64, max float64) (bool, float64, Vector, Vector) {
	oc := r.origin().Sub(s.Center)
	a := r.direction().Dot(r.direction())
	b := oc.Dot(r.direction())
	c := oc.Dot(oc) - s.r*s.r
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < max && temp > min {
			t := temp
			p := r.pointAt(t)
			N := p.Sub(s.Center).MultiplyByScalar(1 / s.r)
			return true, t, p, N
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < max && temp > min {
			t := temp
			p := r.pointAt(t)
			N := p.Sub(s.Center).MultiplyByScalar(1 / s.r)
			return true, t, p, N
		}
	}
	return false, -1.0, Vector{}, Vector{}

}
