# go-globby

Quickly get the all file names which matched pattern, ignores file system errors such as I/O errors reading directories.

## Quickly start

```go
package main

import (
	"fmt"
	"github.com/Lellansin/go-globby"
)

/*
/path/to/your/project
└── dist
    ├── assert
    │   ├── images
    │   │   ├── empty.jpg
    │   │   └── none.jpg
    │   ├── index.css
    │   └── index.d.ts
    ├── docs
    │   ├── 1.md
    │   ├── 2.md
    │   └── test
    │       └── log
    └── scripts
        ├── cp.sh
        └── test.sh
 */

func main() {
	patterns := []string{
		"dist/assert/**/*",        // any file below "dist/assert" (recursive)
		"!dist/assert/**/*.d.ts",  // exclude any *.d.ts under "dist/assert" (recursive)
		"dist/scripts",            // same as "dist/scripts/**/*" (recursive)
		"dist/docs/*.md",          // only *.md under "dist/image" (no recursive)
		"!dist/docs/2.md",         // only *.md under "dist/image" (no recursive)
	}
	// opt := globby.Option{ baseDir: "/fullpath/to/your/project" }
	opt := globby.Option{
		BaseDir: "/path/to/your/project", // default is os.Getwd()
		RelativeReturn: true,             // default is false
	}

	matches := globby.Match(patterns, opt)
  for _, file := range matches {
    fmt.Println(file)
  }
  /*
	Output:
		dist/assert/images/empty.jpg
		dist/assert/images/none.jpg
		dist/assert/index.css
		dist/scripts/cp.sh
		dist/scripts/test.sh
		dist/docs/1.md
	*/
}

```

## Syntax

### path

* Absolutely path: `/path` 
* Relative path: `./path`, `path`, `path/`

### include

* Both `path/**/*` and `path/` got all the child files recursively.
* Get all the `*.js` file recursively: `path/**/*.js`
* Get the `*.js` file under path/: `path/*.js` (not recursively)

### exclude

* Both `!path/**/*` and `!path/` exclude all the child files recursively.
* Exclude all the `*.js` file recursively: `!path/**/*.js`
* Exclude the `*.js` file under path/: `!path/*.js` (not recursively)

## Option

* `Option.BaseDir` string, to specialfy the base directory, default to os.Getwd().
* `Option.IgnoreDot` bool, if true, items like `.git/`, `.gitignore` will be ignore, default is false.
* `Option.RelativeReturn` bool, if true Match() will return relative filepath string array, default is false (return absolutely filepath)

## Test

see [globby_test.go](https://github.com/Lellansin/go-globby/blob/master/globby_test.go) for more.
