package ray

import "github.com/anirudhsundar/go-ray-tracer/src/coremath"

type Ray struct {
	origin, direction *coremath.Tuple
}

func Position(r *Ray, t float64) coremath.Tuple {
	tot_distance := r.direction.ScalarMultiply(t)
	ret := tot_distance.Add(*r.origin)
	return ret
}
