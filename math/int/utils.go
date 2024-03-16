package Int

import "fmt"

// Lerp performs a linear interpolation between start and end, using progress as a
// percentage of the total range. If progress is outside the range [0, range_], the
// function will panic. If range_ is zero, the function will panic.
//
// The result is rounded to the nearest integer.
//
// Example usage:
//
//	Lerp(10, 20, 50, 100) = 15
//
// The above will interpolate between 10 and 20, using 50% of the total range of
// 100. The result is 15.
func Lerp(start, end, progress, range_ int) (result int) {
	if range_ == 0 {
		panic("Integer division by zero in lerp")
	}

	if progress < 0 || progress > range_ {
		panic(fmt.Sprintf("Progress value %d out of range [0, %d] in lerp", progress, range_))
	}

	result = start + (end-start)*progress/range_
	return
}
