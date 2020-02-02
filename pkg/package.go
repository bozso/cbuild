//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

/* AST Error :
cannot parse line: `UnusedAttr 0x7953930 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x7953930 <line:15:451> Inherited unused

*/

package pkg

// package_var_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:11
type package_var_type = int32

const (
	build_var_set         package_var_type = 0
	build_var_set_default                  = 1
	build_var_append                       = 2
)

// package_var_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:17
type package_var_t struct {
	name      []byte
	value     []byte
	operation package_var_type
}

// package_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:23
type package_t struct {
	deps        []hash_t
	exports     []hash_t
	symbols     []hash_t
	ordered     []interface{}
	n_exports   uint32
	variables   []package_var_t
	n_variables uint32
	name        []byte
	source_abs  []byte
	generated   []byte
	header      []byte
	errors      uint32
	exported    int32
	c_file      int32
	force       int32
	silent      int32
	out         []stream_t
}

// package_path_cache - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:45
var package_path_cache []hash_t

// package_id_cache - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:46
var package_id_cache []hash_t

// package_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:50
var package_new func([]byte, [][]byte) []package_t

// package_emit - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:52
func package_emit(pkg []package_t, value []byte) {
	if pkg[0].out != nil {
		// TODO: fix the circrular dependency issue
		stream_write(pkg[0].out, value, uint32(noarch.Strlen(value)))
	}
}

// package_c_file - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.c:56
func package_c_file(abs_path []byte, error_ [][]byte) []package_t {
	var cached []package_t = hash_get(package_path_cache, abs_path).([]package_t)
	if len(cached) != 0 {
		return cached
	}
	var pkg []package_t = make([]package_t, 1)
	pkg[0].name = abs_path
	pkg[0].source_abs = abs_path
	pkg[0].generated = abs_path
	pkg[0].c_file = 1
	hash_set(package_path_cache, abs_path, pkg)
	return pkg
}
