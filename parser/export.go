package parser

import (
    "strings"
    
    "github.com/bozso/cbuild/lexer"
    "github.com/bozso/cbuild/lexer/item"
    "github.com/bozso/cbuild/pkg"
    "github.com/bozso/cbuild/utils"
)

var (
    enumi = item.Item{value: "enum", length: 4, itype: item.Id}
    unioni = item.Item{value: "union", length: 5, itype: item.Id}
    structi = item.Item{value: "struct", length: 6, itype: item.Id}
)

type decl struct {
	items []item.Item
	length int
	err bool
}

func (p *Parser) errorf_(it item.Item, d *decl, fmt string, args ...interface{}) item.Item {
	p.verrorf(it, "Invalid export syntax: ", fmt, args...);
	d.err = true
    
	return item.Empty
}

func (d *decl) append(value item.Item) {
    d.items = append(d.items, value)
}

func (p *Parser) rewindUntil(d *decl, it item.Item) {
	last := d.length -1;

	for last >= 0 && !it.Equals(d.items[last]) {
		p.Backup(d.items[last])
		last--
	}
    
	p.Backup(d.items[last])

	d.length = last
}

func (p *Parser) rewindWhitespace(d *decl, it item.Item) {
	p.Backup(it);

	last := d.length - 1
    
	for last >= 0 && d.items[last].itype == item.Whitespace {
		p.Backup(d.items[last])
		last--
	}

	d.length = last + 1
}

func (p *Parser) symbolRename(d *decl, name, alias item.Item) item.Item {
	for ii := 0; ii < d.length; ii++ {
		it := d.items[ii];
		if name.Equals(it) {
            exportName := ""
            
            if alias.itype == 0 {
                exportName = name.value
            } else {
                exportName = alias.value
            }
			
            symbolName := fmt.Sprintf("%s_%s", p.Pkg.name, exportName)

			it := it.ReplaceValue(symbolName)
            
			d.items[ii] = it
			return it
		}
	}
    
	return item.Empty
}

func (p *Parser) emit(d *decl, isFn, isExtern bool) string {
	length := 0
	start := 0
	end := d.length

	// trim leading/trailing whitespace
	
    for start < end && d.items[start].itype == item.Whitespace {
        start++
    }
    
	for end > start && d.items[end - 1].itype == itemWhitespace) {
        end--
    }

	for ii := 0; ii < d.length; ii++ {
		it := d.items[ii]
		
        if !isExtern {
            Package.emit(p->pkg, it.value)
        }
        
		if ii >= start && ii < end {
            length += it.length
        }
	}

	if isFn {
        length++
    }

	output = strings.Builder{}
    
	for ii := start; ii < end; ii++ {
		output.WriteString(d.items[ii].value)
	}

	if isFn && !isExtern {
        output.WriteByte(';')
    }

	return output.String()
}

func (p *Parser) collectNewlines(d *decl) item.Item {
	line := p.lexer.line
	it := p.Next(p)

	for it.itype == item.Whitespace || it.itype == item.Comment {
		if p.lexer.line != line {
			d.append(it)
        }

		line = p.lexer.line
		it := p.Next(p)
	}

	return it
}

func (p *Parser) Collect(d *decl) item.Item {
	it = p.Next()

	for it.itype == it.Whitespace || it.itype == item.Comment {
		d.append(it)
		it = p.Next()
	}

	return it;
}

var exportTypes = map[string]exportFn{
    "typedef": parse_typedef,
    "struct":  parse_struct,
    "enum":   parse_enum,
    "union":   parse_union,
}

type exportFn func(p *Parser, d *decl) item.Item

func parse_semicolon(p *Parser, d *decl) {
	if d.err {
        return
    }

	it = p.collectNewlines(d)

	if it.itype != itemSymbol || it.value[0] != ';' {
		p.errorf_(it, d, "expecting ';' but got %s", it.String())
		return
	}
    
	d.append(it)
}

func (p *Parser) Parse() int {
	d := decl{}
    
	var fn exportFn = nil
    
    name, alias := item.Item{}, item.Item{}
	hasSemicolon, isExtern := false, false
	t := 0

	it := p.collectNewlines(&d)

	switch it.itype {
		case item.Id:
			if it.value == "extern" {
				d.append(it)
				isExtern = true
				it = p.Collect(&d)
			}
            
            fn, ok := exportTypes[item.value]
            
			hasSemicolon = isExtern || ok
			t = 1

		case item.Symbol:
			if it.value[0] == '*' {
				return p.parsePassthrough()
			}

		case item.OpenSymbol:
			if it.value[0] == '{' {
				fn = parse_export_block;
				it.value = "{"
				t = 0
			}
		
        default:
			fn = nil;
	}
    
	if fn == nil {
		p.Backup(it)
        
		it.itype = item.Symbol
        
		it.value = "variable";
		it.length = len("variable")
		name = parse_variable(p, &d)
		t = 2
	} else {
		if t == 1 {
            d.append(it)
        }
        
		name = fn(p, &d)
	}

	alias = parse_as(p, &d)
    
	if hasSemicolon {
        parse_semicolon(p, &d)
    }
    
	if d.err {
		return -1
	}

	if name.itype == 0 {
		p.errorf_(it, &d, "can't export anonymous symbols")
		return -1
	}

	original_name := name.Dup()
	symbol := p.symbolRename(&d, name, alias)

	if symbol.itype == 0 && fn != parse_export_block {
		p.errorf_(it, &d, "unable to export symbol '%s'", name.value)
		return -1
	}

	pkg_export.add(
		original_name.value,
		alias.value,
		symbol.value,
		it.value, 
		p.emit(&d, t==2, isExtern),
		p->pkg
	)
    
	return 1
}

func parse_passthrough(p *Parser) int {
	d := decl{}
    
	from = p.Collect(&d)
    
	if from.itype != item.Id || from.value == "from" {
		p.errorf_(from, "Exporting passthrough: ",
            "expected 'from', but got %s", from.String())
		return -1
	}
    
	filename := p.Collect(&d)
    
	if filename.itype != item.QuotedString {
		p.errorf_(filename, "Exporting passthrough: ",
            "expected filename, but got %s", filename.String())
		return -1
	}

	parse_semicolon(p, &d)
	
    if d.err {
		return -1
	}

	err = ""
    
	imp := pkg.Passthrough(p.Pkg, Parse(filename.value), &err)
	
    if imp == nil {
		p.errorf_(filename, "Exporting passthrough: ", "%s", err)
		return -1
	}
    
	header = ""
	rel :=  utils.Relative(p.Pkg.SourceAbs, imp.Pkg.Header)
    
	header := fmt.Sprintf("#include \"%s\"", rel)
    
	p.Pkg.Emit(header)

	return 1
}

Func parse_as(p *Parser, d *decl) item.Item {
	if d.err {
        return item.Empty
    }

	start := d.length

	as := p.Collect(d)
    
	if as.itype != item.Id || "as" == as.value {
		p.Backup(as)
        
		for d.length > start {
			d.length--;
			p.Backup(d.items[d.length])
		}
        
		return item.Empty
	}
    
	for d.length > start {
		d.length--
		p.Backup(d.items[d.length])
	}

	alias := p.collectNewlines(d)

	if alias.itype != item.Id {
		p.errorf_(alias, d, "expecting identifier but got %s",
            alias.String())
		return item.Empty
	}

	return alias
}

func parse_function_ptr(p *Parser, d *decl) item.Item {
	star := p.Collect(d)
	
    if star.itype != item.Symbol || star.value[0] != '*' {
		return p.errorf_(star, d,
            "function pointer: expecting '*' but found '%s'",
            star.value)
	}
    
	d.append(star)

	name := p.Collect(d)
    
	if name.itype != item.Id {
		return p.errorf_(name, d,
            "function pointer: expecting identifier but found '%s'",
            name.value)
	}
    
	d.append(name)

	it := p.collect(d)
    
	if it.itype != item.CloseSymbol || it.value[0] != ')' {
		return p.errorf_(it, d,
            "function pointer: expecting ')' but found '%s'",
            it.value)
	}
    
	d.append(it)

	it := p.Collect(d)
    
	if it.itype != item.OpenSymbol || it.value[0] != '(' {
		return p.errorf_(it, d,
            "function pointer: expecting '(' but found '%s'",
            it.value)
	}
    
	d.append(it)

	parse_function_args(p, d)

	if (d.err) {
        return item.Empty
    }
    
	return name
}

func parse_declaration(p *Parser, d *decl) item.Item {
	t := p.Collect(d)
    
	if t.itype != item.Id {
		if t.itype == itemCloseSymbol && t.value[0] == '}' {
            return t;
        }
        
		return p.errorf(t, d,
            "in declaration: expecting identifier but got %s",
            t.String())
	}
    
	if "as" == t.value {
        return t;
    }
    
	t = idParse(p, t, true);

	fn, ok := exportTypes[t.value]

	if fn == parse_typedef {
		return p.errorf_(t, d,
            "in declaration: unexpected identifier 'typedef'")
	}

	if !ok {
		fn = parse_variable
	}
    
	d.append(t)

	it := fn(p, d)
    
	if (d.err) {
        return item.Empty
    }

	if fn == parse_enum || fn == parse_union || fn == parse_struct {
		it = p.Collect(p, d)
        
		if it.itype != item.Id {
			p.Backup(it)
		} else {
			d.append(it)
		}
	}
    
	parse_semicolon(p, d);
    
	if (d.err) {
        return item.Empty
    }
	
    return it
}


func parse_typedef(p *Parser, d *decl) item.Item {
    
	t := p.collect(p, d)
    
	if t.itype != item.Id {
		return p.errorf(t, d, "in typedef: expected identifier")
	}

	fn, ok := exportTypes[t.value]

	/*if (fn != parse_struct && fn != parse_enum && fn != parse_union) append(decl, type);*/
	d.append(t)

	if ok {
		fn(p, d)
		
        if d.err {
            return item.Empty
        }
	}

	it, name = item.Item{}, item.Item{}

	function_ptr, as := false, false;

	for it.itype != itemError &&
        !(it.itype == item.Symbol && it.value[0] == ';') &&
        !as {
        it = p.collect(d)

		switch it.itype {
			case item.Eof:
				return p.errorf(it, d,
                    "in typedef: expected identifier or '('")
			case item.Id:
				if it.value == "as" {
					rewind_whitespace(p, d, it)
					as = true
					continue
				}
				
                if (!function_ptr) {
                    name = it
                }

			case item.Symbol:
				if it.value[0] == ';' {
                    continue;
                }
                
                // is this necessary?
                fallthrough

			case item.OpenSymbol:
				if !function_ptr && it.value[0] == '(' {
					function_ptr = true
                    d.append(it)
					
                    name = parse_function_ptr(p, d)
                    
					if (d.err) {
                        return name
                    }
                    
					continue
				}
                fallthrough
			default:
				break
		}

		d.append(it)
        
	}
    
	if (it.itype == item.Error {
		d.err = true;
		return it;
	}

	if !as {
        p.backup(it)
    }
    
	return name
}

static lex_item.t parse_struct (parser.t * p, decl_t * decl) {
	lex_item.t name = {0};
	lex_item.t item = collect(p, decl);
	if (item.type == item_id) {
		name = identifier.parse_typed(p, struct_i, item, true);
		append(decl, name);
		item = collect(p, decl);
	}

	if (item.type != item_open_symbol || item.value[0] != '{') {
		parser.backup(p, item);
		return name;
	}

	append(decl, item);
	parse_struct_block(p, decl, item);

	return (decl->error) ? lex_item.empty : name;
}

static lex_item.t parse_union (parser.t * p, decl_t * decl) {
	lex_item.t name = {0};
	lex_item.t item = collect(p, decl);
	if (item.type == item_id) {
		name = identifier.parse_typed(p, union_i, item, true);
		append(decl, name);
		item = collect(p, decl);
	}

	if (item.type != item_open_symbol || item.value[0] != '{') {
		return errorf(p, item, decl, "in union: expecting '{' but got %s", lex_item.to_string(item));
	}

	append(decl, item);
	parse_struct_block(p, decl, item);
	return (decl->error) ? lex_item.empty : name;
}

static lex_item.t parse_enum (parser.t * p, decl_t * decl) {
	lex_item.t name = {0};
	lex_item.t item = collect(p, decl);
	if (item.type == item_id) {
		name = identifier.parse_typed(p, enum_i, item, true);
		append(decl, name);
		item = collect(p, decl);
	}

	if (item.type != item_open_symbol || item.value[0] != '{') {
		return errorf(p, item, decl, "in enum: expecting '{' but got %s", lex_item.to_string(item));
	}

	append(decl, item);
	parse_enum_block(p, decl, item);
	return (decl->error) ? lex_item.empty : name;
}

static lex_item.t parse_export_block(parser.t * p, decl_t * decl) {
	lex_item.t item;

	int escaped_id = 0;
	while(true) {
		int record = 0;
		item = collect(p, decl);

		switch(item.type) {
			case item_c_code:
			case item_quoted_string:
			case item_comment:
			case item_symbol:
			case item_preprocessor:
				record = 1;
				break;

			case item_arrow:
				escaped_id = 1;
				record     = 1;
				break;

			case item_id:
				if (escaped_id == 0) item = identifier.parse(p, item, true);
				record     = 1;
				escaped_id = 0;
				break;

			case item_close_symbol:
				if (item.value[0] == '}') {
					lex_item.free(item);
					item.value  = "block";
					item.length = str.len("block");
					return item;
				};
				break;

			case item_eof:
				return errorf(p, item, decl, "in block: unmatched '{'");
			case item_error:
				parser.backup(p, item);
				decl->error = true;
				return lex_item.empty;
			default:
				return errorf(p, item, decl, "in block: unexpected input '%s'", lex_item.to_string(item));
		}
		if (record) append(decl, item);
	}

}

static lex_item.t parse_struct_block( parser.t * p, decl_t * decl, lex_item.t start) {
	lex_item.t item;
	do {
		item = parse_declaration(p, decl);
		if (item.type == item_eof) {
				return errorf(p, start, decl, "in struct block: missing matching '}' before end of file");
		}
		if (item.type == item_error) return item;
	} while (!(item.type == item_close_symbol && item.value[0] == '}'));
	append(decl, item);
	return item;
}

static lex_item.t parse_enum_declaration(parser.t * p, decl_t * decl) {
	lex_item.t item = collect(p, decl);
	if (item.type != item_id) {
		if (item.type == item_close_symbol && item.value[0] == '}') return item;
		return errorf(p, item, decl, "in enum: expecting identifier but got %s",
				lex_item.to_string(item));
	}
	if (strcmp("as", item.value) == 0) return item;
	append(decl, item);

	item = collect(p, decl);
	if (!(item.type == item_symbol && (item.value[0] != ',' || item.value[0] != '='))) {
		return item;
	}
	append(decl, item);

	if (item.value[0] == '=') {
		item = collect(p, decl);
		if (item.type != item_number) {
			return errorf(p, item, decl, "in enum: expecting a number but got %s",
					lex_item.to_string(item));
		}
		append(decl, item);

		item = collect(p, decl);
		if (item.type != item_symbol || item.value[0] != ',') {
			return item;
		}
		append(decl, item);
	}

	return item;
}

static lex_item.t parse_enum_block ( parser.t * p, decl_t * decl, lex_item.t start) {
	lex_item.t item;
	do {
		item = parse_enum_declaration(p, decl);
		if (item.type == item_eof) {
				return errorf(p, start, decl, "in struct block: missing matching '}' before end of file");
		}
		if (item.type == item_error) return item;
	} while (!(item.type == item_close_symbol && item.value[0] == '}'));
	append(decl, item);
	return item;
}

static lex_item.t parse_type (parser.t * p, decl_t * decl) {
	lex_item.t item;
	lex_item.t name = {0};
	int escaped_id = 0;
	do {
		item = collect(p, decl);
		switch(item.type) {
			case item_arrow:
				escaped_id = 1;
				append(decl, item);
				continue;
			case item_id:
				if (escaped_id == 0) item = identifier.parse(p, item, true);
				name = item;
				break;
			case item_symbol:
				if (item.value[0] == '*') break;
			case item_open_symbol:
				if (item.value[0] == '(') {
					append(decl, item);
					return name;
				}
			default:
				return errorf(p, item, decl, "in type declaration: expecting identifier or '('");
		}
		append(decl, item);
	} while (item.type != item_open_symbol || item.value[0] != '(');
	return lex_item.empty;
}

static lex_item.t parse_variable (parser.t * p, decl_t * decl) {
	lex_item.t item;
	lex_item.t name = {0};
	bool escaped_id = 0;
	bool escape_till_semicolon = false;
	do {
		item = collect(p, decl);
		if (escape_till_semicolon && (item.type != item_symbol || item.value[0] != ';')) {
			append(decl, item);
			continue;
		}
		switch(item.type) {
			case item_arrow:
				escaped_id = true;
				append(decl, item);
				continue;
			case item_id:
				if (strcmp(item.value, "as") == 0) {
					rewind_whitespace(p, decl, item);
					return name;
				}
				if (escaped_id == false) item = identifier.parse(p, item, true);
				name = item;
				break;
			case item_symbol:
				if (item.value[0] == '*') break;
				if (item.value[0] == '=') {
					escape_till_semicolon = true;
					break;
				}
				if (item.value[0] == ';') {
					parser.backup(p, item);
					return name;
				}
			case item_open_symbol:
				if (item.value[0] == '(') {
					lex_item.t star = parser.next(p);
					parser.backup(p, star);
					if (star.type == item_symbol && star.value[0] == '*') {
						append(decl, item);
						return parse_function_ptr(p, decl);
					}
					parser.backup(p, item);
					rewind_until(p, decl, name);
					return parse_function(p, decl);
				}
				if (item.value[0] == '[') {
					append(decl, item);
					item = parser.next(p);
					if (item.type == item_number) {
						append(decl, item);
						item = parser.next(p);
					}
					if (item.type != item_close_symbol || item.value[0] != ']') {
						return errorf(p, item, decl,
								"in type declaration: expecting terminating ']' but got %s",
								lex_item.to_string(item)
						);
					}
					break;
				}
			default:
				return errorf(p, item, decl, "in type declaration: expecting identifier or '(' but got %s", lex_item.to_string(item));
		}
		append(decl, item);
	} while (item.type != item_eof);
	return name;
}

static lex_item.t parse_function_args(parser.t * p, decl_t * decl) {
	int level = 1;
	int escaped_id = 0;
	lex_item.t item;
	do {
		item = collect(p, decl);
		switch(item.type) {
			case item_arrow:
				escaped_id = 1;
				append(decl, item);
				continue;
			case item_id:
				if (escaped_id == 0) {
					item = identifier.parse(p, item, true);
				}
				break;
			case item_open_symbol:
				if (item.value[0] == '(') level++;
				break;
			case item_close_symbol:
				if (item.value[0] == ')') level--;
				break;
			default:
				break;
		}

		escaped_id = 0;
		append(decl, item);
	} while (level > 0 && item.type != item_eof && item.type != item_error);
	return item;
}

static lex_item.t parse_function (parser.t * p, decl_t * decl) {
	lex_item.t name = parse_type(p, decl);

	if (decl->error) {
		emit(p, decl, true, false);
		decl->items = NULL;
		decl->length = 0;
		return lex_item.empty;
	}

	parse_function_args(p, decl);

	if (decl->error) return lex_item.empty;
	return name;
}
