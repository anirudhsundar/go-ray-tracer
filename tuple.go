package ray_tracer

type Tuple struct {
	x, y, z, w float64
}

func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}
