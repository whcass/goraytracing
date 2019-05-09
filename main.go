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

func colour(r Ray, objs []Sphere) Vector {
	max := math.MaxFloat64
	min := 0.0
	hitAnything := false
	normal := Vector{}
	//pSomething := Vector{}
	for _, elem := range objs {
		hit, t, _, N := elem.hitSphere(r, min, max)
		if hit {
			hitAnything = true
			max = t
			normal = N
			//pSomething = p
			//N := r.pointAt(t).Sub(Vector{0, 0, -1}).Normalize()
			//return Vector{N.X + 1, N.Y + 1, N.Z + 1}.MultiplyByScalar(0.5)
		}
	}
	if hitAnything {
		return Vector{normal.X + 1, normal.Y + 1, normal.Z + 1}.MultiplyByScalar(0.5)
	} else {
		unitDir := r.direction().Normalize()
		t := 0.5 * (unitDir.Y + 1.0)
		colour := Vector{1.0, 1.0, 1.0}.MultiplyByScalar(1.0 - t)
		colour = colour.Add(Vector{0.5, 0.7, 1.0}.MultiplyByScalar(t))
		return colour
	}

}

func main() {
	nx := 800
	ny := 400
	d1 := []byte("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n" + strconv.Itoa(255) + "\n")
	lowerLeftCorner := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}
	origin := Vector{0.0, 0.0, 0.0}
	floor := Sphere{Vector{0, -100.5, -1}, 100}
	main := Sphere{Vector{0, 0, -1}, 0.5}
	sphereList := []Sphere{floor, main}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			direction := lowerLeftCorner.Add(horizontal.MultiplyByScalar(u))
			direction = direction.Add(vertical.MultiplyByScalar(v))

			r := Ray{origin, direction}
			col := colour(r, sphereList)

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
