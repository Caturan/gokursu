package functions

func FourTrans(num1 int, num2 int) (int, int, int, float32) {
	add := num1 + num2
	subst := num1 - num2
	multi := num1 * num2
	divi := num1 / num2

	return add, subst, multi, float32(divi)
}
