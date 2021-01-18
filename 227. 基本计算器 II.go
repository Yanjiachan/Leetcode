func calculate(s string) int {
	var helper func(*int) int
	helper = func(i *int) int {
		stack := []int{0}
		num := 0
		sign := "+"
		for ; *i < len(s); *i++ {
			val := string(s[*i])
			if val <= "9" && val >= "0" {
				temp, _ := strconv.Atoi(val)
				num = num*10 + temp
			}
			if val == "(" {
				*i++
				num = helper(i)
			}
			if !(val <= "9" && val >= "0") && val != " " || *i == len(s)-1 {
				//非数字
				switch sign {
				case "+":
					stack = append(stack, num)
				case "-":
					stack = append(stack, -num)
				case "*":
					stack[len(stack)-1] *= num
				case "/":
					stack[len(stack)-1] /= num
				default:
				}
				sign = val
				num = 0
			}
			if val == ")" {
				*i++
				break
			}
		}
		for len(stack) > 1 {
			stack[0] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		return stack[0]
	}
	i := 0
	return helper(&i)
}
