package utils

import (
    "os"
)

func countLevels(path string) int {
    return strings.Count(path, os.PathSeparator)
}

func Relative(from, to string) (s string) {
	// determine shared prefix
    
    ii, lastStep = 0, 0
	
    l1, l2 := len(from), len(to)
    
    ll := 0
    
    if l1 > l2 {
        ll = l2
    } else {
        ll = l1
    }
    
    for ii = 0; ii < ll; ii++ {
        c1, c2 := from[ii], to[ii]
        
        if c1 != c2 {
            break
        }
        
        if IsPathSeparator(c1) {
            lastStep = ii
        }
    }
    
	dots := countLevels(from[last_sep + 1: l1])
    
    if dots == 0 {
        dots = 2
    } else {
        dots *= 3
    }
    
    // TODO: check to make sure this is right
    s := make(byte[], dots + l2 - lastSep + 1)
	s[0] = 0;
    
    for dots-- {
        s += "../"
    }
	
    s += to[lastSep + 1: l2]
    
    return
}

func Newer(a, b string) bool {
	struct stat sta;
	struct stat stb;

	int resa = lstat(a, &sta);
	if (resa == -1) return false;

	int resb = lstat(b, &stb);
	if (resb == -1) return true;


#ifdef __MACH__
	if (sta.st_mtimespec.tv_sec == stb.st_mtimespec.tv_sec) {
		return sta.st_mtimespec.tv_nsec > stb.st_mtimespec.tv_nsec;
	}
	return sta.st_mtimespec.tv_sec > stb.st_mtimespec.tv_sec;
#else
	if (sta.st_mtim.tv_sec == stb.st_mtim.tv_sec) {
		return sta.st_mtim.tv_nsec > stb.st_mtim.tv_nsec;
	}
	return sta.st_mtim.tv_sec > stb.st_mtim.tv_sec;
#endif
}
