package Int

type Vector interface {
}

type Vec2 [2]int

func (v0 Vec2) Add(v1 Vec2) (v Vec2) {
	v[0] = v0[0] + v1[0]
	v[1] = v1[1] + v1[1]

	return
}

func (v0 Vec2) Sub(v1 Vec2) (v Vec2) {
	v[0] = v0[0] - v1[0]
	v[1] = v0[1] - v1[1]

	return
}

func (v0 Vec2) RightTo(v1 Vec2, v2 Vec2) bool {

	vv := v2.Sub(v1)
	vp := v0.Sub(v1)

	return vv[1]*vp[0]-vv[0]*vp[1] <= 0
}

// getDeterminant(a, b, c) {
//     const ab = new Vector(b);
//     ab.sub(a);

//     const ac = new Vector(c);
//     ac.sub(a);

//     return ab[1] * ac[0] - ab[0] * ac[1];
// }
