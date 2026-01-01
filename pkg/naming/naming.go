package naming

func ToSnake(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	var out []rune
	for i, r := range runes {
		if isUpper(r) {
			if i > 0 {
				prev := runes[i-1]
				next := rune(0)
				if i+1 < len(runes) {
					next = runes[i+1]
				}
				if isLower(prev) || isDigit(prev) || (isUpper(prev) && next != 0 && isLower(next)) {
					out = append(out, '_')
				}
			}
			out = append(out, toLower(r))
			continue
		}
		if r == '-' || r == ' ' {
			out = append(out, '_')
			continue
		}
		out = append(out, toLower(r))
	}
	return string(out)
}

func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = toLower(runes[0])
	return string(runes)
}

func ToUpperCamel(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = toUpper(runes[0])
	return string(runes)
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r - 'A' + 'a'
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 'a' + 'A'
	}
	return r
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
