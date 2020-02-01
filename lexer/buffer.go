package lexer

import (
    "github.com/bozso/cbuild/lexer/item"
)

type ItemBuffer struct {
	item.T  []items;
    cursor int
}

func New(int count) (ib ItemBuffer) {
    ib.items = make([]items, count)
    ib.cursor = 0
    
    return
}
