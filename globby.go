package globby

import (
	"os"
	"fmt"
	"regexp"
	"path/filepath"
)

type Option struct {
	baseDir string
	ignoreDot bool
	excludes []string
}

func Match(patterns []string, opt Option) []string {
	var allFiles []string
	patterns, opt, err := completeOpt(patterns, opt)
	if err != nil {
		fmt.Printf("Magth err: [%v]\n", err)
		return allFiles
	}
	for _, pattern := range patterns {
		files := find(pattern, opt);
		if files == nil || len(*files) == 0 {
			continue
		}
		allFiles = append(allFiles, *files...)
	}
	return allFiles
}

func find(pattern string, opt Option) *[] string {
	// match ./some/path/**/*
	if regexTest("\\*\\*", pattern) || 
		!regexTest("\\*", pattern) { // Dirname
		return findRecr(pattern, opt)
	}
	// match ./some/path/*
	if regexTest("\\*", pattern) {
		return findDir(pattern, opt)
	}
	return nil
}

func findDir(pattern string, opt Option) *[]string {
	var list []string
	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("err: [%v]\n", err)
		return &list
	}
	list = append(list, files...)
	return &list
}

/*
 * find recr
 */
func findRecr(pattern string, opt Option) *[]string {
	dir := strReplace(pattern, "\\*\\*.+", "");
	afterMacth := ""
	matchAfterFlag := false
	if regexTest("\\*", pattern) {
		afterMacth = strReplace(pattern, ".+\\*", "")
		matchAfterFlag = len(afterMacth) > 0;
	}

	var list []string
	err := filepath.Walk(dir, func (fullpath string, f os.FileInfo, err error) error {
		if f.IsDir() && regexTest("^\\.", f.Name()) {
        return filepath.SkipDir
    }
    if f.IsDir() {
    	return nil
    }
		path, _ := filepath.Rel(opt.baseDir, fullpath)
		if checkExclude(opt, path) {
			return nil
		}
		if !matchAfterFlag {
			list = append(list, fullpath);
			return nil
		}
		if regexTest(afterMacth + "$", path) {
			list = append(list, fullpath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("err: [%v]\n", err)
	}
	return &list
}

func completeOpt(srcPatterns []string, opt Option) ([]string, Option, error) {
	// TODO: default base dir    config dir > file dir > pwd
	// if len(opt.baseDir) == 0 {
	// 	opt.baseDir = os.
	// }

	var patterns []string
	for _, pattern := range srcPatterns {
		// TODO: check no "tmp/*", use "tmp" or "tmp/*.ext" instead

		if regexTest("^\\!", pattern) {
			opt.excludes = append(opt.excludes, strReplace(pattern, "^\\!", ""))
			continue;
		}
		if regexTest("^\\.", pattern) || // like ./dist
			 !regexTest("^\\/", pattern) { // like dist
			patterns = append(patterns, filepath.Join(opt.baseDir, pattern))
			continue;
		}
		patterns = append(patterns, pattern)
	}
	return patterns, opt, nil
}

func checkExclude(opt Option, path string) bool {
	// if exludes dirs
	for _, exclude := range opt.excludes {
		rule := exclude
		if regexTest("\\*\\*", exclude) {
			rule = strReplace(exclude, "\\*\\*/\\*+?", ".+")
		} else if regexTest("\\*", exclude) {
			rule = strReplace(exclude, "\\*", "[^/]+")
		}
		if regexTest("^" + rule, path) {
			return true // ignore
		}
	}
	return false
}

func regexTest(re string, src string) bool {
		matched, err := regexp.MatchString(re, src)
		if err != nil {
			return false;
		}
		if matched {
			return true
		}
		return false;
}

func strReplace(dest, text, repl string) string {
	re := regexp.MustCompile(text)
	return re.ReplaceAllString(dest, repl)
}
