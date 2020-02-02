//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals

package main

// cli_flag_options - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:10
type cli_flag_options struct {
	description []byte
	long_name   []byte
	short_name  []byte
	required    int32
}

// flag_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:17
type flag_type = int32

const (
	flag_type_bool   flag_type = 0
	flag_type_int              = 1
	flag_type_string           = 2
)

// flag_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:23
type flag_t struct {
	dest           interface{}
	type_          flag_type
	description    []byte
	long_name      []byte
	short_name     []byte
	short_name_buf [3]byte
	required       int32
	has_value      int32
}

// cli_cmd_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:39
type cli_cmd_t struct {
	cb          interface{}
	ctx         interface{}
	name        []byte
	description []byte
}

// cli_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:51
type cli_t struct {
	flags           []hash_t
	commands        []hash_t
	default_command []cli_cmd_t
	has_commands    int32
	usage           []byte
	name            []byte
	argc            int32
	argv            [][]byte
}

// cli_cmd_cb - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:65
type cli_cmd_cb = func([]cli_t, []byte, interface{}) int32

// flag_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:34
func flag_free(f []flag_t) {
	_ = f[0].long_name
	_ = f
}

// cmd_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:46
func cmd_free(c []cli_cmd_t) {
	_ = c[0].name
	_ = c
}

// cli_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:67
func cli_new(usage []byte) []cli_t {
	var cli []cli_t = make([]cli_t, 1)
	cli[0].flags = kh_init_ptr()
	cli[0].commands = kh_init_ptr()
	cli[0].usage = usage
	return cli
}

// flag - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:76
func flag(cli []cli_t, dest interface{}, type_ flag_type, options cli_flag_options) {
	var f []flag_t = make([]flag_t, 1)
	f[0].type_ = type_
	f[0].dest = dest
	f[0].required = options.required
	f[0].description = options.description
	f[0].long_name = make([]byte, noarch.Strlen(options.long_name)+int32(3))
	noarch.Strcpy(f[0].long_name, []byte("--\x00"))
	noarch.Strcat(f[0].long_name, options.long_name)
	hash_set(cli[0].flags, f[0].long_name, f)
	if len(options.short_name) == 0 {
		f[0].short_name = nil
	} else {
		f[0].short_name_buf[:][0] = '-'
		f[0].short_name_buf[:][1] = options.short_name[0]
		f[0].short_name_buf[:][2] = byte(0)
		f[0].short_name = f[0].short_name_buf[:]
		hash_set(cli[0].flags, f[0].short_name, f)
	}
}

// cli_flag_bool - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:102
func cli_flag_bool(cli []cli_t, out []int32, options cli_flag_options) {
	options.required = 0
	flag(cli, out, flag_type_bool, options)
}

// cli_flag_int - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:107
func cli_flag_int(cli []cli_t, out []int32, options cli_flag_options) {
	flag(cli, out, flag_type_string, options)
}

// cli_flag_string - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:111
func cli_flag_string(cli []cli_t, out [][]byte, options cli_flag_options) {
	flag(cli, out, flag_type_string, options)
}

// cli_command - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:115
func cli_command(cli []cli_t, name []byte, cb cli_cmd_cb, description []byte, is_default int32, ctx interface{}) {
	cli[0].has_commands = 1
	var cmd []cli_cmd_t = make([]cli_cmd_t, 1)
	cmd[0].name = noarch.Strdup(name)
	cmd[0].cb = cli_cmd_cb(cb)
	cmd[0].ctx = ctx
	cmd[0].description = description
	if int32((is_default)) != 0 {
		cli[0].default_command = cmd
	}
	hash_set(cli[0].commands, cmd[0].name, cmd)
}

// parse_flag - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:128
func parse_flag(flag []flag_t, arg []byte) {
	switch uint32(int32((flag[0].type_))) {
	case uint32(int32((flag_type_bool))):
		if flag[0].dest != nil {
			(flag[0].dest.([]int32))[0] = 1
		}
	case uint32(int32((flag_type_int))):
		if flag[0].dest != nil {
			(flag[0].dest.([]int32))[0] = noarch.Strtol(arg, nil, 0)
		}
	case uint32(int32((flag_type_string))):
		if flag[0].dest != nil {
			(flag[0].dest.([][]byte))[0] = arg
		}
	}
	flag[0].has_value = 1
}

// cli_usage - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:142
func cli_usage(cli []cli_t) int32 {
	noarch.Fprintf(noarch.Stderr, []byte("usage %s: [options]%s %s\n\n\x00"), cli[0].name, func() []byte {
		if int32((cli[0].has_commands)) != 0 {
			return []byte(" command\x00")
		}
		return []byte("\x00")
	}(), cli[0].usage)
	{
		// flags
		noarch.Fprintf(noarch.Stderr, []byte("FLAGS: \n\n\x00"))
		var width int32 = noarch.Strlen([]byte("--help\x00"))
		{
			var key []byte
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].flags)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].flags)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].flags)[0].keys[k]))
					{
						width = int32(func() uint32 {
							if noarch.Strlen(key) > int32(uint32(width)) {
								return uint32(noarch.Strlen(key))
							}
							return uint32(width)
						}())
					}
				}
			}
		}
		var w int32 = int32(uint32(width) - uint32(noarch.Strlen([]byte("--help\x00"))))
		noarch.Fprintf(noarch.Stderr, []byte("          %s%*.*s %s\t%s\n\x00"), []byte("--help\x00"), w, w, []byte(" \x00"), []byte("-h\x00"), []byte("prints this message\x00"))
		{
			var key []byte
			var val interface{}
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].flags)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].flags)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].flags)[0].keys[k]))
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
					val = (cli[0].flags)[0].vals[k]
					{
						if int32(key[1]) != int32('-') {
							continue
						}
						var flag []flag_t = val.([]flag_t)
						var w int32 = int32(uint32(width) - uint32(noarch.Strlen(key)))
						noarch.Fprintf(noarch.Stderr, []byte("          %s%*.*s %s\t%s\n\x00"), key, w, w, []byte(" \x00"), func() []byte {
							if flag[0].short_name != nil {
								return flag[0].short_name
							}
							return []byte("  \x00")
						}(), flag[0].description)
					}
				}
			}
		}
		noarch.Fprintf(noarch.Stderr, []byte("\n\x00"))
	}
	if int32((cli[0].has_commands)) != 0 {
		noarch.Fprintf(noarch.Stderr, []byte("COMMANDS: \n\n\x00"))
		var width int32
		{
			var key []byte
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].commands)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].commands)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].commands)[0].keys[k]))
					{
						width = int32(func() uint32 {
							if noarch.Strlen(key) > int32(uint32(width)) {
								return uint32(noarch.Strlen(key))
							}
							return uint32(width)
						}())
					}
				}
			}
		}
		{
			var key []byte
			var val interface{}
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].commands)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].commands)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].commands)[0].keys[k]))
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
					val = (cli[0].commands)[0].vals[k]
					{
						var cmd []cli_cmd_t = val.([]cli_cmd_t)
						var w int32 = int32(uint32(width) - uint32(noarch.Strlen(key)))
						noarch.Fprintf(noarch.Stderr, []byte("%s%s%*.*s  %s\n\x00"), func() []byte {
							if (int64(uintptr(unsafe.Pointer(&cmd[0])))/int64(64) - int64(uintptr(unsafe.Pointer(&cli[0].default_command[0])))/int64(64)) == 0 {
								return []byte("(default) \x00")
							}
							return []byte("          \x00")
						}(), cmd[0].name, w, w, []byte(" \x00"), cmd[0].description)
					}
				}
			}
		}
		noarch.Fprintf(noarch.Stderr, []byte("\n\x00"))
	}
	return -1
}

// cli_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:185
func cli_parse(cli []cli_t, argc int32, argv [][]byte) int32 {
	cli[0].argv = make([][]byte, 1*int32(uint32(argc-1)))
	cli[0].name = argv[0]
	var cmd []cli_cmd_t
	var i int32
	{
		// collect args, commands, flags
		for i = 1; i < argc; i++ {
			var arg []byte = argv[i]
			if noarch.Strcmp([]byte("--\x00"), arg) == 0 {
				break
			}
			if noarch.Strcmp([]byte("-h\x00"), arg) == 0 || noarch.Strcmp([]byte("--help\x00"), arg) == 0 {
				return cli_usage(cli)
			}
			var arg_allocated []byte = noarch.Strdup(arg)
			var flag []flag_t = hash_get(cli[0].flags, arg_allocated).([]flag_t)
			if len(flag) != 0 {
				parse_flag(flag, arg)
				_ = arg_allocated
				continue
			}
			if int32(arg[0]) == int32('-') {
				noarch.Fprintf(noarch.Stderr, []byte("Unrecognized flag '%s'\n\n\x00"), arg)
				_ = arg_allocated
				return cli_usage(cli)
			}
			if len(cmd) == 0 {
				cmd = hash_get(cli[0].commands, arg_allocated).([]cli_cmd_t)
				if cmd != nil {
					_ = arg_allocated
					continue
				}
			}
			cli[0].argv[cli[0].argc] = arg
			cli[0].argc++
			_ = arg_allocated
		}
	}
	for ; i < argc; i++ {
		// collect args after a --
		cli[0].argv[cli[0].argc] = argv[i]
		cli[0].argc++
	}
	if len(cmd) == 0 {
		cmd = cli[0].default_command
	}
	if len(cmd) == 0 {
		return func() int32 {
			if int32((cli[0].has_commands)) != 0 {
				return cli_usage(cli)
			}
			return 0
		}()
	}
	var cb cli_cmd_cb = cmd[0].cb.(cli_cmd_cb)
	var result int32 = cb(cli, cmd[0].name, cmd[0].ctx)
	if result == 0 {
		return result
	}
	return cli_usage(cli)
}

// cli_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:241
func cli_free(cli []cli_t) {
	{
		var key []byte
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((cli[0].flags)[0].n_buckets)); k++ {
				if !noarch.Not((cli[0].flags)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
				key = []byte(((cli[0].flags)[0].keys[k]))
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (cli[0].flags)[0].vals[k]
				{
					if int32(key[0]) == int32('-') && int32(key[1]) == int32('-') {
						flag_free(val.([]flag_t))
					}
				}
			}
		}
	}
	kh_destroy_ptr(cli[0].flags)
	{
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((cli[0].commands)[0].n_buckets)); k++ {
				if !noarch.Not((cli[0].commands)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (cli[0].commands)[0].vals[k]
				{
					cmd_free(val.([]cli_cmd_t))
				}
			}
		}
	}
	kh_destroy_ptr(cli[0].commands)
	_ = cli[0].argv
	_ = cli
}
