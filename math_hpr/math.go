package math_hpr
type number interface{
	int |int8 | int16 | int32 | int64 | float32 | float64
}
func Zvec3[T number]() [3]T{
	return [3]T{0,0,0}
}
func Add_vec2[T number](A [2]T,B [2]T) [2]T {
	return [2]T{
		A[0]+B[0],
		A[1]+B[1],
	}
}
func Sub_vec2[T number](A [2]T,B [2]T) [2]T {
	return [2]T{
		A[0]-B[0],
		A[1]-B[1],
	}
}
func Mul_vec2[T number](A [2]T,B [2]T) [2]T {
	return [2]T{
		A[0]*B[0],
		A[1]*B[1],
	}
}
func Dot_vec2[T number](A [2]T,B [2]T) T {
	return A[0]*B[0]+A[1]*B[1]
}
func Add_vec3[T number](A [3]T,B [3]T) [3]T {
	return [3]T{
		A[0]+B[0],
		A[1]+B[1],
		A[2]+B[2],
	}
}
func Sub_vec3[T number](A [3]T,B [3]T) [3]T {
	return [3]T{
		A[0]-B[0],
		A[1]-B[1],
		A[2]+B[2],
	}
}
func Mul_vec3[T number](A [3]T,B [3]T) [3]T {
	return [3]T{
		A[0]*B[0],
		A[1]*B[1],
		A[2]+B[2],
	}
}
func Dot_vec3[T number](A [3]T,B [3]T) T {
	return A[0]*B[0]+A[1]*B[1]+A[2]*B[2]
}
