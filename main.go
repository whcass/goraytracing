package main

import (
	"io/ioutil"
	"math"
	"math/rand"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RandInUnitSphere() Vector {
	p := Vector{}
	for {
		p = Vector{rand.Float64(), rand.Float64(), rand.Float64()}.MultiplyByScalar(2.0)
		p = p.Sub(Vector{1.0, 1.0, 1.0})
		if !(p.SquaredLength() >= 1.0) {
			break
		}
	}
	return p
}

func colour(r Ray, objs []Sphere) Vector {
	max := math.MaxFloat64
	min := 0.0
	hitAnything := false
	normal := Vector{}
	p := Vector{}
	for _, elem := range objs {
		hit, t, a, N := elem.hitSphere(r, min, max)
		if hit {
			hitAnything = true
			max = t
			normal = N
			p = a
			//N := r.pointAt(t).Sub(Vector{0, 0, -1}).Normalize()
			//return Vector{N.X + 1, N.Y + 1, N.Z + 1}.MultiplyByScalar(0.5)
		}
	}
	if hitAnything {
		target := p.Add(normal)
		target = target.Add(RandInUnitSphere())
		return colour(Ray{p, target.Sub(p)}, objs).MultiplyByScalar(0.5)
		//return Vector{normal.X + 1, normal.Y + 1, normal.Z + 1}.MultiplyByScalar(0.5)
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
	ns := 100
	d1 := []byte("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n" + strconv.Itoa(255) + "\n")

	lowerLeftCorner := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}
	origin := Vector{0.0, 0.0, 0.0}

	camera := Camera{lowerLeftCorner, horizontal, vertical, origin}

	floor := Sphere{Vector{0, -100.5, -1}, 100}
	main := Sphere{Vector{0, 0, -1}, 0.5}
	sphereList := []Sphere{floor, main}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vector{0, 0, 0}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := camera.GetRay(u, v)
				//p :=r.pointAt(2.0)
				col = col.Add(colour(r, sphereList))
			}
			col = col.MultiplyByScalar(0.01)

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
