package Int

func Lerp(a, b, t, l int) (r int) {
	r = a + (b-a)*t/l
	return
}
