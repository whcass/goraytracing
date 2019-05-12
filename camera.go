package main

import "math"

type Camera struct {
	lowerLeftCorner Vector
	horizontal      Vector
	vertical        Vector
	origin          Vector
}

func (c Camera) GetRay(u float64, v float64) Ray {
	direction := c.lowerLeftCorner.Add(c.horizontal.MultiplyByScalar(u))
	direction = direction.Add(c.vertical.MultiplyByScalar(v))
	direction = direction.Sub(c.origin)
	return Ray{c.origin, direction}
}

func NewCamera(lookFrom Vector, lookat Vector, vup Vector,vfov float64, aspect float64) Camera{
	theta := vfov*math.Pi/180
	halfHeight := math.Tan(theta/2.0)
	halfWidth := aspect * halfHeight
	w := lookFrom.Sub(lookat)
	w = w.Normalize()
	u := vup.Cross(w)
	u = u.Normalize()
	v := w.Cross(u)
	lowerLeft := lookFrom.Sub(u.MultiplyByScalar(halfWidth))
	lowerLeft = lowerLeft.Sub(v.MultiplyByScalar(halfHeight))
	lowerLeft = lowerLeft.Sub(w)

	horizontal := u.MultiplyByScalar(2*halfWidth)
	vertical := v.MultiplyByScalar(2*halfHeight)
	cam := Camera{
		lowerLeft,
		horizontal,
		vertical,
		lookFrom,
	}
	return cam
}
