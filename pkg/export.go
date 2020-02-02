//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package pkg

// package_export_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:15
type package_export_type = int32

const (
	type_block    package_export_type = 0
	type_type                         = 1
	type_function                     = 2
	type_enum                         = 3
	type_union                        = 4
	type_struct                       = 5
	type_header                       = 6
)

// type_names - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:25
var type_names [][]byte = [][]byte{[]byte("block\x00"), []byte("type\x00"), []byte("function\x00"), []byte("enum\x00"), []byte("union\x00"), []byte("struct\x00"), []byte("header\x00")}

// package_export_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:35
type package_export_t struct {
	local_name  []byte
	export_name []byte
	declaration []byte
	symbol      []byte
	type_       package_export_type
}

// types - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:58
var types []hash_t

// package_export_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:43
func package_export_free(exp []package_export_t) {
	if len(exp) == 0 {
		return
	}
	if (int64(uintptr(unsafe.Pointer(&exp[0].local_name[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&exp[0].export_name[0])))/int64(1)) == 0 {
		_ = exp[0].local_name
	} else {
		_ = exp[0].local_name
		_ = exp[0].export_name
	}
	_ = exp[0].declaration
	_ = exp[0].symbol
	_ = exp
}

// init_types - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:59
func init_types() {
	types = kh_init_ptr()
	hash_set(types, []byte("typedef\x00"), int32((type_type)))
	hash_set(types, []byte("function\x00"), int32((type_function)))
	hash_set(types, []byte("enum\x00"), int32((type_enum)))
	hash_set(types, []byte("union\x00"), int32((type_union)))
	hash_set(types, []byte("struct\x00"), int32((type_struct)))
	hash_set(types, []byte("header\x00"), int32((type_header)))
}

// package_export_add - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:70
func package_export_add(local []byte, alias []byte, symbol []byte, type_ []byte, declaration []byte, parent []package_t) []byte {
	var export_name []byte = func() []byte {
		if len(alias) == 0 {
			return local
		}
		return alias
	}()
	if hash_has(parent[0].exports, export_name) != 0 {
		_ = local
		_ = alias
		_ = symbol
		_ = declaration
		return export_name
	}
	if len(types) == 0 {
		init_types()
	}
	var exp []package_export_t = make([]package_export_t, 1)
	exp[0].local_name = local
	exp[0].export_name = export_name
	exp[0].declaration = declaration
	exp[0].type_ = package_export_type((hash_get(types, type_)))
	exp[0].symbol = symbol
	parent[0].ordered = realloc(parent[0].ordered, 8*(uint32(parent[0].n_exports)+1)).([]interface{})
	parent[0].ordered[uint32(parent[0].n_exports)] = exp
	parent[0].n_exports += uint32(1)
	hash_set(parent[0].exports, exp[0].export_name, exp)
	hash_set(parent[0].symbols, exp[0].local_name, exp)
	return exp[0].export_name
}

// get_header_path - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:103
func get_header_path(generated []byte) []byte {
	var header []byte = strings_dup(generated)
	var len_ uint32 = strings_len(header)
	header[len_-1] = 'h'
	return header
}

// package_export_write_headers - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:112
func package_export_write_headers(pkg []package_t) {
	if pkg[0].header != nil {
		return
	}
	pkg[0].header = get_header_path(pkg[0].generated)
	if int32((pkg[0].force)) == 0 && (int32((pkg[0].silent)) != 0 || noarch.Not(utils_newer(pkg[0].source_abs, pkg[0].header))) {
		return
	}
	var header []stream_t = atomic_stream_open(pkg[0].header)
	stream_printf(header, []byte("#ifndef _package_%s_\n#define _package_%s_\n\n\x00"), pkg[0].name, pkg[0].name)
	var last_type package_export_type
	var had_newline int32 = 1
	var i int32
	for i = 0; uint32(i) < uint32(pkg[0].n_exports); i++ {
		var exp []package_export_t = pkg[0].ordered[i].([]package_export_t)
		var newline int32
		var prefix int32
		var multiline int32 = int32((noarch.BoolToInt(len(noarch.Strchr(exp[0].declaration, int32('\n'))) != 0)))
		if int32((had_newline)) == 0 && (uint32(int32((last_type))) != uint32(int32((exp[0].type_))) || int32((multiline)) != 0) {
			prefix = 1
		}
		if int32((multiline)) != 0 {
			newline = 1
		}
		stream_printf(header, []byte("%s%s\n%s\x00"), func() []byte {
			if int32((prefix)) != 0 {
				return []byte("\n\x00")
			}
			return []byte("\x00")
		}(), exp[0].declaration, func() []byte {
			if int32((newline)) != 0 {
				return []byte("\n\x00")
			}
			return []byte("\x00")
		}())
		last_type = exp[0].type_
		had_newline = newline
	}
	stream_printf(header, []byte("%s#endif\n\x00"), func() []byte {
		if int32((had_newline)) != 0 {
			return []byte("\x00")
		}
		return []byte("\n\x00")
	}())
	stream_close(header)
}

// package_export_export_headers - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.c:144
func package_export_export_headers(pkg []package_t, dep []package_t) {
	if len(pkg) == 0 || len(dep) == 0 {
		return
	}
	package_export_write_headers(dep)
	var rel []byte = utils_relative(pkg[0].source_abs, dep[0].header)
	var decl []byte
	asprintf((*[1000000][]byte)(unsafe.Pointer(&decl))[:], []byte("#include \"%s\"\x00"), rel)
	package_export_add(strings_dup(rel), nil, strings_dup([]byte("\x00")), []byte("header\x00"), decl, pkg)
	_ = rel
}
