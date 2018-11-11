package numbers

func gcdTwoNumbers(a int, b int) int {
	if b == 1 {
		return 1
	}

	mod := a % b
	if mod == 0 {
		return b
	} else {
		return gcdTwoNumbers(b, mod)
	}
}

func Gcd(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return gcdTwoNumbers(nums[0], nums[1])
	} else {
		gcd_ := gcdTwoNumbers(nums[0], nums[1])
		elems := append(nums[2:], gcd_)

		return Gcd(elems...)
	}
}

func Lcm(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		a := nums[0]
		b := nums[1]

		return a * b / Gcd(a, b)
	} else {
		lcm_ := Lcm(nums[0], nums[1])
		elems := append(nums[2:], lcm_)

		return Lcm(elems...)
	}
}
