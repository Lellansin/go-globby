package globby

import (
	"os"
	"testing"
	"path/filepath"
	"sort"
	"reflect"
)

func TestMatch(t *testing.T) {
	opt := Option{}
	opt.ignoreDot = false
	opt.baseDir = "/Users/lellansin/github/static"
	arr := []string{ "!tmp/online-code", "tmp/*.sh" }
	Match(arr, opt)
}

func TestSignleStarFiles(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir := filepath.Join(curDir, "./tmp");
	defer os.RemoveAll(tmpDir)
	makeTmpFiles(tmpDir, []string {
		".git",
		"app.js",
		"package.json",
		"src/router.js",
		"src/store.js",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	})

	opt := Option{}
	opt.baseDir = tmpDir

	patterns := []string{
		"src/*.js",
	}

	files := Match(patterns, opt)
	expected := []string {
		"src/router.js",
		"src/store.js",
	}
	if checkFiles(tmpDir, files, expected) {
		t.Errorf("files not match, expected %v, but got %v", expected, files)
	}
}

/*
 *  Match "src/api"
 */
func TestDirMatch(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir := filepath.Join(curDir, "./tmp");
	defer os.RemoveAll(tmpDir)
	makeTmpFiles(tmpDir, []string {
		".git",
		"app.js",
		"package.json",
		"src/router.js",
		"src/store.js",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	})

	opt := Option{}
	opt.baseDir = tmpDir
	patterns := []string{
		"src/api",
	}

	files := Match(patterns, opt)
	expected := []string {
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	}
	if checkFiles(tmpDir, files, expected) {
		t.Errorf("files not match, expected %v, but got %v", expected, files)
	}
}

/*
 *  Match "/**" + "/*"
 */
func TestDirStar(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir := filepath.Join(curDir, "./tmp");
	defer os.RemoveAll(tmpDir)
	makeTmpFiles(tmpDir, []string {
		".git",
		"app.js",
		"package.json",
		"src/router.js",
		"src/store.js",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	})

	opt := Option{}
	opt.baseDir = tmpDir
	patterns := []string{
		"src/**/*",
	}

	files := Match(patterns, opt)
	expected := []string {
		"src/router.js",
		"src/store.js",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	}
	if checkFiles(tmpDir, files, expected) {
		t.Errorf("files not match, expected %v, but got %v", expected, files)
	}
}

/*
 *  Match "/**" + "/*.js"
 */
func TestDirStar2(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir := filepath.Join(curDir, "./tmp");
	defer os.RemoveAll(tmpDir)
	makeTmpFiles(tmpDir, []string {
		".git",
		"app.js",
		"package.json",
		"src/router.js",
		"src/store.js",
		"src/store.ts",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	})

	opt := Option{}
	opt.baseDir = tmpDir
	patterns := []string{
		"src/**/*.js",
	}

	files := Match(patterns, opt)
	expected := []string {
		"src/router.js",
		"src/store.js",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	}
	if checkFiles(tmpDir, files, expected) {
		t.Errorf("files not match, expected %v, but got %v", expected, files)
	}
}

/*
 * Match "/**" + "/*.js"
 * ignore files in the match items 
 */
func TestDirIgnoreFile(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir := filepath.Join(curDir, "./tmp");
	defer os.RemoveAll(tmpDir)
	makeTmpFiles(tmpDir, []string {
		".git",
		"app.js",
		"package.json",
		"src/router.js",
		"src/store.js",
		"src/store.ts",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
		"src/service/home.js",
		"src/service/user.js",
		"src/service/test.js",
	})

	opt := Option{}
	opt.baseDir = tmpDir
	patterns := []string{
		"src/**/*.js",
		"!src/service/home.js",
	}

	files := Match(patterns, opt)
	expected := []string {
		"src/router.js",
		"src/store.js",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
		"src/service/user.js",
		"src/service/test.js",
	}
	if checkFiles(tmpDir, files, expected) {
		t.Errorf("files not match, expected %v, but got %v", expected, files)
	}
}

/*
 * Match "/**" + "/*.js"
 * ignore dir in the match items 
 */
func TestDirIgnoreDir(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir := filepath.Join(curDir, "./tmp");
	defer os.RemoveAll(tmpDir)
	makeTmpFiles(tmpDir, []string {
		".git",
		"app.js",
		"package.json",
		"src/router.js",
		"src/store.js",
		"src/store.ts",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
		"src/service/home.js",
		"src/service/user.js",
		"src/service/test.js",
	})

	opt := Option{}
	opt.baseDir = tmpDir
	patterns := []string{
		"src/**/*.js",
		"!src/service",
	}

	files := Match(patterns, opt)
	expected := []string {
		"src/router.js",
		"src/store.js",
		"src/api/home.js",
		"src/api/user.js",
		"src/api/test.js",
	}
	if checkFiles(tmpDir, files, expected) {
		t.Errorf("files not match, expected %v, but got %v", expected, files)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func makeTmpFiles(baseDir string, files []string) {
	for _, file := range files {
		file = filepath.Join(baseDir, file)
		dir, _ := filepath.Split(file)
		os.MkdirAll(dir, os.ModePerm)
		os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0666)
	}
}

func checkFiles(baseDir string, resultFiles []string, expectedFiles []string) bool {
	var expected []string
	for _, file := range expectedFiles {
		expected = append(expected, filepath.Join(baseDir, file))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(resultFiles)))
	sort.Sort(sort.Reverse(sort.StringSlice(expected)))
	return !reflect.DeepEqual(resultFiles, expected);
}
