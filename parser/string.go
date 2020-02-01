package parser

func escape(c byte) byte {
	switch c {
		case 'a':
			return '\a'
		case 'b':
			return '\b'
		case 'f':
			return '\f'
		case 'n':
			return '\n'
		case 'r':
			return '\r'
		case 't':
			return '\t'
		case 'v':
			return '\v'
		default:
			return c
	}
}

func Parse(input string) (s string) {
	c := input[0]
	o := input


	if c != '"' {
		return "" // nota quoted string
	}
	c++

	for *c != 0 && *c != '"' {
		switch(*c) {
			case '\\':
				c++
				*o = escape(c)
			default:
				*o = *c
				break
		}
		c++
        o++
	}
    
	*o = 0
	
    return input
}
