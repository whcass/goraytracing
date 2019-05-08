package main

import (
	"io/ioutil"
	"math"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hitSphere(center Vector, radius float64, r Ray) float64 {
	oc := r.origin().Sub(center)
	a := r.direction().Dot(r.direction())
	b := 2.0 * oc.Dot(r.direction())
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}

}

func colour(r Ray) Vector {
	t := hitSphere(Vector{0, 0, -1}, 0.5, r)
	if t > 0.0 {
		N := r.pointAt(t).Sub(Vector{0, 0, -1}).Normalize()
		return Vector{N.X + 1, N.Y + 1, N.Z + 1}.MultiplyByScalar(0.5)
	}
	unitDir := r.direction().Normalize()
	t = 0.5 * (unitDir.Y + 1.0)
	colour := Vector{1.0, 1.0, 1.0}.MultiplyByScalar(1.0 - t)
	colour = colour.Add(Vector{0.5, 0.7, 1.0}.MultiplyByScalar(t))
	return colour
}

func main() {
	nx := 400
	ny := 200
	d1 := []byte("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n" + strconv.Itoa(255) + "\n")
	lowerLeftCorner := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}
	origin := Vector{0.0, 0.0, 0.0}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			direction := lowerLeftCorner.Add(horizontal.MultiplyByScalar(u))
			direction = direction.Add(vertical.MultiplyByScalar(v))

			r := Ray{origin, direction}
			col := colour(r)

			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)
			line := []byte(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
			d1 = append(d1, line...)

		}
	}
	err := ioutil.WriteFile("image.ppm", d1, 0644)
	check(err)
}
