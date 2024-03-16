package Int

type Vec2 struct {
	X, Y int
}

func (v Vec2) Add(v1 Vec2) Vec2 {
	return Vec2{
		X: v.X + v1.X,
		Y: v.Y + v1.Y,
	}
}

func (v Vec2) Sub(v1 Vec2) Vec2 {
	return Vec2{
		X: v.X - v1.X,
		Y: v.Y - v1.Y,
	}
}

func (v Vec2) Mul(v1 Vec2) Vec2 {
	return Vec2{
		X: v.X * v1.X,
		Y: v.Y * v1.Y,
	}
}

func (v Vec2) Div(v1 Vec2) Vec2 {
	return Vec2{
		X: v.X / v1.X,
		Y: v.Y / v1.Y,
	}
}

func (v Vec2) Scale(s int) Vec2 {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v Vec2) Neg() Vec2 {
	return Vec2{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vec2) Abs() Vec2 {
	return v
}

func (v Vec2) Floor() Vec2 {
	return v
}

func (v Vec2) Ceil() Vec2 {
	return v
}

func (v Vec2) Round() Vec2 {
	return v
}

func (v Vec2) RoundTo(places int) Vec2 {
	return v
}

func (v0 Vec2) RightTo(v1 Vec2, v2 Vec2) bool {

	vv := v2.Sub(v1)
	vp := v0.Sub(v1)

	return vv.Y*vp.X-vv.X*vp.Y <= 0

}
