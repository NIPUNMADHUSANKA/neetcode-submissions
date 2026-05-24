func isValid(s string) bool {
    
	var stack []rune
	var x rune

	for _, i := range s{
		if i == '{' || i == '[' || i == '(' {
			stack = append(stack, i)
		}else{
			if len(stack) > 0 {
				x, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if i == '}' && x != '{'{
					return false
				}else if i == ']' && x != '['{
					return false
				}else if i == ')' && x != '('{
					return false
				}
			}else{
				return false
			}
		}

	}
	return len(stack)==0
}
