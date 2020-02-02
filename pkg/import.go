//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//


package pkg

// package_import_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/import.c:15
type package_import_t struct {
	alias    []byte
	filename []byte
	c_file   int32
	pkg      []package_t
}
// package_import_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/import.c:22
func package_import_free(imp []package_import_t) []package_t {
	if len(imp) == 0 {
		return nil
	}
	if (int64(uintptr(unsafe.Pointer(&imp[0].alias[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&imp[0].filename[0])))/int64(1)) == 0 {
		_ = imp[0].alias
	} else {
		_ = imp[0].alias
		_ = imp[0].filename
	}
	var pkg []package_t = imp[0].pkg
	_ = imp
	return pkg
}
// package_import_add - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/import.c:37
func package_import_add(alias []byte, filename []byte, parent []package_t, error_ [][]byte) []package_import_t {
	var imp []package_import_t = make([]package_import_t, 1)
	hash_set(parent[0].deps, alias, imp)
	imp[0].alias = alias
	imp[0].filename = filename
	imp[0].c_file = 0
	imp[0].pkg = package_new(filename, error_, parent[0].force, parent[0].silent)
	if len(imp[0].pkg) == 0 {
		return nil
	}
	package_export_write_headers(imp[0].pkg)
	return imp
}
// package_import_add_c_file - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/import.c:53
func package_import_add_c_file(parent []package_t, filename []byte, error_ [][]byte) []package_import_t {
	var alias []byte = realpath(filename, nil)
	if len(alias) == 0 {
		error_[0] = noarch.Strerror((noarch.ErrnoLocation())[0])
		return nil
	}
	var imp []package_import_t = make([]package_import_t, 1)
	hash_set(parent[0].deps, alias, imp)
	imp[0].alias = alias
	imp[0].filename = filename
	imp[0].c_file = 1
	imp[0].pkg = package_c_file(alias, error_)
	return imp
}
// package_import_passthrough - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/import.c:72
func package_import_passthrough(parent []package_t, filename []byte, error_ [][]byte) []package_import_t {
	var imp []package_import_t = package_import_add(filename, filename, parent, error_)
	if len(error_[0]) != 0 {
		return nil
	}
	if len(imp) == 0 || len(imp[0].pkg) == 0 {
		asprintf(error_, []byte("Could not import '%s'\x00"), filename)
		return nil
	}
	{
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((imp[0].pkg[0].exports)[0].n_buckets)); k++ {
				if !noarch.Not((imp[0].pkg[0].exports)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/package/import.c:80 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/package/import.c:80 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/package/import.c:80 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (imp[0].pkg[0].exports)[0].vals[k]
				{
					var exp []package_export_t = val.([]package_export_t)
					hash_set(parent[0].exports, exp[0].export_name, exp)
				}
			}
		}
	}
	package_export_export_headers(parent, imp[0].pkg)
	return imp
}
