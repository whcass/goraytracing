package main

import (
	"math"
	"math/rand"
)

type Material struct {
	colour  Vector
	matType string
	fuzz    float64
}

func (m Material) scatter(r Ray, p Vector, norm Vector) (bool, Ray, Vector) {
	if m.matType == "metal" {
		b, scatter, atten := m.metalScatter(r, p, norm)
		return b, scatter, atten
	} else if m.matType == "lambertian" {
		b, scatter, atten := m.lamScatter(r, p, norm)
		return b, scatter, atten
	} else if m.matType == "dielectric" {
		b, scatter, atten := m.dielectricScatter(r, p, norm)
		return b, scatter, atten
	}

	return false, Ray{}, Vector{}
}

func (m Material) lamScatter(ray Ray, p Vector, norm Vector) (bool, Ray, Vector) {
	target := p.Add(norm)
	target = target.Add(RandInUnitSphere())
	scattered := Ray{p, target.Sub(p)}
	attenuation := m.colour
	return true, scattered, attenuation
}

func (m Material) metalScatter(ray Ray, p Vector, norm Vector) (bool, Ray, Vector) {
	reflected := reflect(ray.direction().Normalize(), norm)
	scattered := Ray{p, reflected.Add(RandInUnitSphere().MultiplyByScalar(m.fuzz))}
	attenuation := m.colour
	return scattered.direction().Dot(norm) > 0, scattered, attenuation
}

func (m Material) dielectricScatter(ray Ray, p Vector, norm Vector) (bool, Ray, Vector) {
	reflected := reflect(ray.direction(), norm)
	atten := Vector{1.0, 1.0, 1.0}
	outwardNormal := Vector{}
	niOverNt := 0.0
	cosine := 0.0
	reflectProb := 0.0
	if ray.direction().Dot(norm) > 0 {
		outwardNormal = norm.MultiplyByScalar(-1)
		niOverNt = m.colour.X
		cosine = m.colour.X * ray.direction().Dot(norm) / ray.direction().Length()
	} else {
		outwardNormal = norm
		niOverNt = 1.0 / m.colour.X
		cosine = -ray.direction().Dot(norm) / ray.direction().Length()
	}
	scattered := Ray{}
	done, refracted := refract(ray.direction(), outwardNormal, niOverNt)
	if done {
		//
		reflectProb = shlick(cosine, m.colour.X)
	} else {
		scattered = Ray{p, reflected}
		reflectProb = 1.0
	}
	if rand.Float64() < reflectProb {
		scattered = Ray{p, reflected}
	} else {
		scattered = Ray{p, refracted}
	}
	return true, scattered, atten
}

func reflect(v Vector, n Vector) Vector {
	a := n.MultiplyByScalar(2 * v.Dot(n))
	return v.Sub(a)
}

func refract(v Vector, n Vector, niOverNt float64) (bool, Vector) {
	uv := v.Normalize()
	dt := uv.Dot(n)
	discriminat := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminat > 0 {
		refracted := uv.Sub(n.MultiplyByScalar(dt))
		refracted = refracted.MultiplyByScalar(niOverNt)
		refracted = refracted.Sub(n.MultiplyByScalar(math.Sqrt(discriminat)))
		return true, refracted
	} else {
		return false, Vector{}
	}
}

func shlick(cosine float64, idx float64) float64 {
	r0 := (1 - idx) / (1 + idx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
