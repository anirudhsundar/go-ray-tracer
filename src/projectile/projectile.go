package projectile

import "github.com/anirudhsundar/go-ray-tracer/src/coremath"

type Projectile struct {
	Position coremath.Tuple
	Velocity coremath.Tuple
}

type Environment struct {
	Gravity coremath.Tuple
	Wind    coremath.Tuple
}

func Tick(env Environment, proj Projectile) Projectile {
	position := proj.Position.Add(proj.Velocity)
	velocity := proj.Velocity.Add(env.Gravity).Add(env.Wind)
	return Projectile{position, velocity}
}
