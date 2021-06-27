package projectile

import "github.com/anirudhsundar/go-ray-tracer/src/tuple"

type Projectile struct {
	Position tuple.Tuple
	Velocity tuple.Tuple
}

type Environment struct {
	Gravity tuple.Tuple
	Wind    tuple.Tuple
}

func Tick(env Environment, proj Projectile) Projectile {
	position := proj.Position.Add(proj.Velocity)
	velocity := proj.Velocity.Add(env.Gravity).Add(env.Wind)
	return Projectile{position, velocity}
}
