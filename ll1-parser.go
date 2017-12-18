package ll1

import s "strings"
import "fmt"

func id(str string) (bool, string) {
	fmt.Println("in id str = ", str)
	if (str[0] >= 'a' && str[0] <= 'z') || (str[0] >= 'A' && str[0] <= 'Z') {
		str = str[1:]
		if len(str) == 0 {
			return true, str
		}
		for (str[0] >= 'a' && str[0] <= 'z') || (str[0] >= 'A' && str[0] <= 'Z') || (str[0] >= '0' && str[0] <= '9') {
			str = str[1:]
			if len(str) == 0 {
				break
			}
		}
	} else {
		return false, str
	}
	return true, str
}

func real(str string) (bool, string) {
	fmt.Println("in real str = ", str)
	for str[0] >= '0' && str[0] <= '9' {
		str = str[1:]
		if len(str) == 0 {
			return true, str
		}
	}
	if str[0] == '.' {
		str = str[1:]
		if len(str) == 0 {
			return false, str
		}
	}
	for str[0] >= '0' && str[0] <= '9' {
		str = str[1:]
		if len(str) == 0 {
			break
		}
	}
	return true, str
}

func A(str string) (bool, string) {
	fmt.Println("in A str = ", str)
	if len(str) == 0 {
		return true, str
	}
	if str[0] == '(' {
		str = str[1:]
		a, r := E(str)
		if !a || len(r) == 0 || r[0] != ')' {
			return false, r
		}
		r = r[1:]
		return true, r
	}
	return true, str
}

func F(str string) (bool, string) {
	if len(str) == 0 {
		return false, str
	}
	fmt.Println("in F str = ", str)
	if (str[0] >= 'a' && str[0] <= 'z') || (str[0] >= 'A' && str[0] <= 'Z') {
		a, r := id(str)
		if !a {
			return false, r
		}
		return A(r)
	}

	if str[0] >= '0' && str[0] <= '9' {
		return real(str)
	}

	if str[0] == '(' {
		str = str[1:]
		a, r := E(str)
		if !a || len(r) == 0 || r[0] != ')' {
			return false, r
		}
		r = r[1:]
		return true, r
	}

	return false, str
}

func TT(str string) (bool, string) {
	fmt.Println("in TT str = ", str)
	if len(str) == 0 {
		return true, str
	}
	if str[0] == '*' || str[0] == '/' {
		str = str[1:]
		a, r := F(str)
		if !a {
			return false, r
		}
		return TT(r)
	}

	return true, str
}

func T(str string) (bool, string) {
	fmt.Println("in T str = ", str)
	a, r := F(str)
	if !a {
		return false, r
	}
	return TT(r)
}

func EE(str string) (bool, string) {
	fmt.Println("in EE str = ", str)
	if len(str) == 0 {
		return true, str
	}
	if str[0] == '+' || str[0] == '-' {

		str = str[1:]
		a, r := T(str)
		if !a {
			return false, r
		}
		return EE(r)
	}

	return true, str
}

func E(str string) (bool, string) {
	fmt.Println("in E str = ", str)
	a, r := T(str)
	if !a {
		return false, r
	}
	return EE(r)
}

func Z(str string) (bool, string) {
	fmt.Println("in Z str = ", str)
	a, r := id(str)
	if !a || len(r) == 0 {
		return false, r
	}
	if r[0] == '=' {
		r = r[1:]
	} else {
		return false, r
	}
	a, r = E(r)
	if !a || len(r) == 0 || r[0] != ';' {
		return false, r
	}
	r = r[1:]
	return true, r
}

func S(str string) (bool, string) {
	fmt.Println("in S str = ", str)
	if len(str) == 0 {
		return true, str
	}

	a, r := Z(str)
	if !a {
		return false, r
	}
	return S(r)
}

func Parse(str string) bool {
	str = s.Replace(str, " ", "", -1)

	a, r := S(str)
	if len(r) == 0 && a {
		return true
	}
	return false
}
