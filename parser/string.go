//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package parser

// escape - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/string.c:5
func escape(c []byte) (c4goDefaultReturn byte) {
	switch int32(c[0]) {
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
		return c[0]
	}
	return
}

// string_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/string.c:26
func string_parse(input []byte) []byte {
	var c []byte = input
	var o []byte = input
	if int32(c[0]) != int32('"') {
		// nota quoted string
		return nil
	}
	func() []byte {
		tempVarUnary := c
		defer func() {
			c = c[0+1:]
		}()
		return tempVarUnary
	}()
	for int32(c[0]) != 0 && int32(c[0]) != int32('"') {
		switch int32(c[0]) {
		case '\\':
			func() []byte {
				tempVarUnary := c
				defer func() {
					c = c[0+1:]
				}()
				return tempVarUnary
			}()
			o[0] = escape(c)
		default:
			o[0] = c[0]
			break
		}
		func() []byte {
			tempVarUnary := c
			defer func() {
				c = c[0+1:]
			}()
			return tempVarUnary
		}()
		func() []byte {
			tempVarUnary := o
			defer func() {
				o = o[0+1:]
			}()
			return tempVarUnary
		}()
	}
	o[0] = byte(0)
	return input
}
