package main

type Material struct {
	colour  Vector
	matType string
}

func (m Material) scatter(r Ray, p Vector, norm Vector) (bool, Ray, Vector) {
	if m.matType == "metal" {
		b, scatter, atten := m.metalScatter(r, p, norm)
		return b, scatter, atten
	} else if m.matType == "lambertian" {
		b, scatter, atten := m.lamScatter(r, p, norm)
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
	scattered := Ray{p, reflected}
	attenuation := m.colour
	return scattered.direction().Dot(norm) > 0, scattered, attenuation
}

func reflect(v Vector, n Vector) Vector {
	a := n.MultiplyByScalar(2 * v.Dot(n))
	return v.Sub(a)
}
