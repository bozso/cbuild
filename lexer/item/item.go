package item

type Type int

const(
	Error Type = iota
	Eof
	Whitespace
	CCode
	Id
	Number
	CharLiteral
	QuotedString
	Preprocessor
	Comment

	/* symbols */
	Symbol,
	OpenSymbol,
	CloseSymbol,
	Arrow,

    totalSymbols
)

var TypeNames = string[totalSymbols]{
	"Error",
	"eof",
	"white space",
	"c code",
	"identifier",
	"number",
	"character literal",
	"string",
	"preprocessor",
	"comment",
	"symbol",
	"open symbol",
	"close symbol",
	"arrow symbol",
}

type Item struct {
	itype Type
	value string
    length, line, linePos, start, index int
}

var Empty = Item{}

func New(value string, itype Type, line, line_pos, start int) Item {
	return Item{
		value    : value,
		length   : len(value),
		itype    : itype,
		line     : line,
		line_pos : line_pos,
		start    : start,
	}
}

func (it Item) String() (s string) {
    name := TypeNames[it.itype]
    
    switch it.itype {
    case Error:
    if v := it.value; len(v) == 0 {
        s = "(null)"
    } else {
        s = v
    }
    
    case Eof:
        s = "<eof>"
    
    default:
    v := it.value
    
    if it.length > 20 || firstOccurrance(v, '\n') != 0 {
        return fmt.Sprintf("%s '%.*s...'", name, 20, v)
    }
    s = fmt.Sprintf("%s '%s'", name, v)
    
    }
    return
}

func (a Item) Equals(b Item) bool {
	return (
			a.itype    == b.itype    &&
			a.value    == b.value    &&
			a.length   == b.length   &&
			a.line     == b.line     &&
			a.line_pos == b.line_pos &&
			a.start    == b.start
	)
}

func (it Item) Dup() Item {
    return New(
        it.value,
        it.itype,
        it.line,
        it.linePos,
        it.start
    )
}

func (it Item) ReplaceValue(value string) Item {
	return New(
		value,
		it.itype,
		it.line,
		it.linePos,
		it.start
	)
}

type Buffer struct {
	items    []Item
}

func NewBuffer() (b T) {
	b.items = make([]Item)
	return
}

func (b Buffer) Len() int {
    return len(b.items)
}

func (b *Buffer) Push(it Item) *Buffer {
    b.items = append(b.items, it)
	return b
}

// TODO: implement this properly
// lex_buffer_next - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:65
//func (b *T) Next() item.T {
	//if uint32(b.length) == 0 {
		//return lex_item_new([]byte("No more items\x00"), item_error, 0, 0, 0)
	//}
	
    //it := b.items[uint32(b.cursor)]
	
    //b.cursor = func() uint32 {
		//if uint32(b.cursor)+1 < uint32(b.capacity) {
			//return uint32(b.cursor) + 1
		//}
		//return 0
	//}()
    
	//b.length -= uint32(1)
	//return it
//}
