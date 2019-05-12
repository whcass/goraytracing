package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strconv"
	"time"
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

func colour(r Ray, objs []Sphere, depth int) Vector {
	hitAnything, normal, p, obj := WorldHit(objs, r)
	if hitAnything {
		//target := p.Add(normal)
		//target = target.Add(RandInUnitSphere())
		b, scattered, attenuation := obj.mat.scatter(r, p, normal)
		if depth < 50 && b {
			return attenuation.Mult(colour(scattered, objs, depth+1))
		} else {
			return Vector{0.0, 0.0, 0.0}
		}
		//return colour(Ray{p, target.Sub(p)}, objs).MultiplyByScalar(0.5)
		//return Vector{normal.X + 1, normal.Y + 1, normal.Z + 1}.MultiplyByScalar(0.5)
	} else {
		unitDir := r.direction().Normalize()
		t := 0.5 * (unitDir.Y + 1.0)
		colour := Vector{1.0, 1.0, 1.0}.MultiplyByScalar(1.0 - t)
		colour = colour.Add(Vector{0.5, 0.7, 1.0}.MultiplyByScalar(t))
		return colour
	}

}

func WorldHit(objs []Sphere, r Ray) (bool, Vector, Vector, Sphere) {
	max := math.MaxFloat64
	min := 0.001
	hitAnything := false
	normal := Vector{}
	p := Vector{}
	obj := Sphere{}
	for _, elem := range objs {
		hit, t, a, N := elem.hitSphere(r, min, max)
		if hit {
			hitAnything = true
			max = t
			normal = N
			p = a
			obj = elem
		}
	}
	return hitAnything, normal, p, obj
}

func main() {
	nx := 1600
	ny := 800
	ns := 400
	d1 := []byte("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n" + strconv.Itoa(255) + "\n")

	//lowerLeftCorner := Vector{-2.0, -1.0, -1.0}
	//horizontal := Vector{4.0, 0.0, 0.0}
	//vertical := Vector{0.0, 2.0, 0.0}
	//origin := Vector{0.0, 0.0, 0.0}

	//camera := Camera{lowerLeftCorner, horizontal, vertical, origin}
	//camera := NewCamera(90,float64(nx)/float64(ny))
	camera := NewCamera(Vector{-2,2,1},Vector{0,0,-1},Vector{0,1,0},90,float64(nx)/float64(ny))
	floor := Sphere{Vector{0, -100.5, -1}, 100, Material{Vector{0.8, 0.8, 0.0}, "lambertian", 0}}
	//main := Sphere{Vector{1, 0, -1}, 0.5}
	//main2 := Sphere{Vector{-0.5, 0, -1}, 0.5}
	sphereList := []Sphere{
		floor,
		{Vector{1, 0, -1}, 0.5, Material{Vector{0.8, 0.6, 0.2}, "metal", 1.0}},
		{Vector{-1, 0, -1}, 0.5, Material{Vector{1.5, 0.8, 0.8}, "dielectric", 0.3}},
		{Vector{-1, 0, -1}, -0.45, Material{Vector{1.5, 0.8, 0.8}, "dielectric", 0.3}},
		{Vector{0, 0, -1}, 0.5, Material{Vector{0.8, 0.3, 0.3}, "lambertian", 0}},
	}
	start := time.Now()
	fmt.Printf("\nRendering...")
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vector{0, 0, 0}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := camera.GetRay(u, v)
				//p :=r.pointAt(2.0)
				col = col.Add(colour(r, sphereList, 0))
			}
			col = col.DivideByScalar(float64(ns))
			col = Vector{math.Sqrt(col.X), math.Sqrt(col.Y), math.Sqrt(col.Z)}
			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)
			line := []byte(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
			d1 = append(d1, line...)

		}
	}
	fmt.Printf("\nDone. Elapsed: %v", time.Since(start))
	err := ioutil.WriteFile("image.ppm", d1, 0644)
	check(err)
}
