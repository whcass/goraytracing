package main

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
