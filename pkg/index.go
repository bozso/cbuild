//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package pkg


// init_cache - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:21
func init_cache() {
	package_path_cache = kh_init_ptr()
	package_id_cache = kh_init_ptr()
}
// abs_path - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:26
func abs_path(rel_path []byte) []byte {
	return realpath(rel_path, nil)
}
// package_name - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:30
func package_name(rel_path []byte) []byte {
	var buffer []byte = noarch.Strdup(rel_path)
	var filename []byte = __xpg_basename(buffer)
	// trim the .c extension
	filename[noarch.Strlen(filename)-int32(2)] = byte(0)
	var c []byte
	for c = filename; int32(c[0]) != 0; func() []byte {
		tempVarUnary := c
		defer func() {
			c = c[0+1:]
		}()
		return tempVarUnary
	}() {
		switch int32(c[0]) {
		case '.':
			fallthrough
		case '-':
			fallthrough
		case ' ':
			c[0] = '_'
		default:
			break
		}
	}
	var name []byte = noarch.Strdup(filename)
	_ = buffer
	return name
}
// assert_name - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:54
func assert_name(relative_path []byte, error_ [][]byte) int32 {
	var len_ noarch.SsizeT = noarch.SsizeT(noarch.Strlen(relative_path))
	var suffix_l uint32 = uint32(noarch.Strlen([]byte(".module.c\x00")))
	if uint32(len_) < suffix_l || noarch.Strcmp([]byte(".module.c\x00"), c4goPointerArithByteSlice(relative_path, int(0+len_-suffix_l))) != 0 {
		if error_ != nil {
			asprintf(error_, []byte("Unsupported input filename '%s', Expecting '<file>.module.c'\x00"), relative_path)
		}
		return 1
	}
	return 0
}
// index_generated_name - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:65
func index_generated_name(path []byte) []byte {
	var len_ noarch.SsizeT = noarch.SsizeT(noarch.Strlen(path))
	var suffix_l uint32 = uint32(noarch.Strlen([]byte(".module.c\x00")))
	var buffer []byte = make([]byte, uint32(len_)-suffix_l+uint32(noarch.Strlen([]byte(".c\x00")))+1)
	noarch.Strncpy(buffer, path, int32(uint32(len_)-suffix_l))
	noarch.Strcpy(c4goPointerArithByteSlice(buffer, int(0+(len_-suffix_l))), []byte(".c\x00"))
	return buffer
}
// index_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:76
func index_parse(input []stream_t, out []stream_t, rel []byte, key []byte, generated []byte, error_ [][]byte, force int32, silent int32) []package_t {
	if len(package_path_cache) == 0 {
		init_cache()
	}
	if package_new == nil {
		package_new = index_new
	}
	var p []package_t = make([]package_t, 1)
	p[0].deps = kh_init_ptr()
	p[0].exports = kh_init_ptr()
	p[0].ordered = nil
	p[0].n_exports = 0
	p[0].symbols = kh_init_ptr()
	p[0].source_abs = key
	p[0].generated = generated
	p[0].out = out
	p[0].name = package_name(p[0].generated)
	p[0].force = force
	p[0].silent = silent
	hash_set(package_path_cache, key, p)
	p[0].errors = uint32(grammer_parse(input, rel, p, error_))
	return p
}
// index_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:109
func index_new(relative_path []byte, error_ [][]byte, force int32, silent int32) []package_t {
	if len(package_path_cache) == 0 {
		init_cache()
	}
	if package_new == nil {
		package_new = index_new
	}
	if int32((assert_name(relative_path, error_))) != 0 {
		return nil
	}
	var key []byte = abs_path(relative_path)
	if len(key) == 0 {
		error_[0] = noarch.Strerror((noarch.ErrnoLocation())[0])
		return nil
	}
	var cached []package_t = hash_get(package_path_cache, key).([]package_t)
	if len(cached) != 0 {
		_ = key
		return cached
	}
	var input []stream_t = file_open(relative_path, 0)
	if input[0].error_.code != 0 {
		error_[0] = noarch.Strdup(input[0].error_.message)
		_ = key
		return nil
	}
	var generated []byte = index_generated_name(key)
	var out []stream_t
	if int32((force)) != 0 || noarch.Not(silent) && int32((utils_newer(relative_path, generated))) != 0 {
		out = atomic_stream_open(generated)
		if out[0].error_.code != 0 {
			_ = key
			noarch.Fprintf(noarch.Stderr, []byte("ERROR: '%s'\n\x00"), out[0].error_.message)
			if error_ != nil {
				error_[0] = noarch.Strdup(out[0].error_.message)
			}
			atomic_stream_abort(out)
			return nil
		}
	}
	var p []package_t = index_parse(input, out, relative_path, key, generated, error_, force, silent)
	if len(p) == 0 || len(error_[0]) != 0 {
		_ = key
		if out != nil {
			atomic_stream_abort(out)
		}
		return nil
	}
	if out != nil {
		stream_close(out)
	}
	return p
}
// index_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/index.c:159
func index_free(pkg []package_t) {
	if len(pkg) == 0 {
		return
	}
	if package_path_cache != nil {
		hash_del(package_path_cache, pkg[0].source_abs)
	}
	var i int32
	{
		// exports
		for i = 0; uint32(i) < uint32(pkg[0].n_exports); i++ {
			package_export_free(pkg[0].ordered[i].([]package_export_t))
		}
	}
	_ = pkg[0].ordered
	kh_destroy_ptr(pkg[0].exports)
	kh_destroy_ptr(pkg[0].symbols)
	_ = pkg[0].name
	_ = pkg[0].source_abs
	_ = pkg[0].generated
	_ = pkg[0].header
	{
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((pkg[0].deps)[0].n_buckets)); k++ {
				if !noarch.Not((pkg[0].deps)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// imports
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/package/index.c:181 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/package/index.c:181 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/package/index.c:181 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (pkg[0].deps)[0].vals[k]
				{
					index_free(package_import_free(val.([]package_import_t)))
				}
			}
		}
	}
	kh_destroy_ptr(pkg[0].deps)
	_ = pkg
}
