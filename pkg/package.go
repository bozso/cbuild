package pkg

import (
    "io"
    "github.com/bozso/cbuild/package/buildvar"    
)

type T struct {
	exported, cFile, force, silent bool
    nExports, nVar, errors int
    name sourceAbs, generated, header string
    out io.Write
    variables []buildvar.T
	
    //hash_t   * deps;
	//hash_t   * exports;
	//hash_t   * symbols;
	//void    ** ordered;
}

/* TODO: fix the circrular dependency issue */
type New func(relativePath string, err []string, force, silent bool) T
type new func(relativePath string, err []string) T


func (p *T) Emit(value string) (n int, err error) {
    return p.out.WriteString(value)
}

var (
    pathCache = make(map[string]T)
    pkgNew new = new_
)

func CFile(absPath string, err []string) (p T) {
	p, ok := pathCache[absPath]
    
	if ok {
        return
    }

	p.name       = absPath
	p.sourceAbs = absPath
	p.generated  = absPath
	p.cFile     = true
    
    pathCache[absPath] = p
	return
}
