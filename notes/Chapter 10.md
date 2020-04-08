# Packages and the Go Tool

## Introduction
Point:
- modularity allows packages to be shared and reused by different projects, within an org or made available to world
- Packages also provide encapsulation by controlling which names are visible or exported outside the package

3 main reasons for Go's compiler speed:
- all imports are explicitly listed at the beginning of each source file so compiler does not have to read and process an entire file to determine its dependencies
- dependencies of a package form a directed acyclic graphc meaning packages can be compiled separately/parallel
- the object file for a compiled Go package records export information for dependencies and package (when compiling the compierl must read one object file for each import but need not look beyond)

## Import Paths
For packages you intend to share or publish - import pahts should be globally unique.
Import paths (outside of the standard library) should start with the Internet domain name of the org that owns or hosts the package. 

## The Package Declaration
A package declaration is required at the start of every Go source file. Main purpose: determine the default identifier for that package.

3 major exceptions to "last segment" convention:
- a package defining a command always have the name main, regardless of the package's import path
- some files in the directory may have the suffic _test on their pacakge name if the file name ends with _test.go
- some tools for dependency management append version number suffixes to package import paths (ex: "gopkg.in/yaml.v2")

## Import Declaration
If need to import 2 packages with same names must specify an alt name for at least one of them to avoid a conflict:
```
import (
    "crypto/rand"
    mrand "math/rand" //alt name mrand avoid conflict 
)
```

Each import declaration establishes a dependency from the current package to the imported package. The go build tool reports an error if these dependencies form a cycle. 

## Blank Imports
To suppress an "unused import" error - rename import in whichj the alt name is _, the blank identifier. 
```
import _ "image/png" //register PNG decoder 
```

## Packages and Naming
- keep names short
- usually take singular form
- be descriptive and unambiguous where possible
- avoid names that already have other connotations
- consider how 2 parts of a qualified identifier work together, not he member name alone (ex: bytes.Equal or flag.Int)

## The Go Tool
The go tool is used for downloading, querying, formatting, building, testing and installing packages of Go code.

It is a package manager (like apt or rpm) that answers queries about its inventory of packages, computes their dependencies and downloads them from remote version-control systems. 

- go build: compile packages and dependencies
- go clean: remove object files
- go doc: show documentation for package or symbol
- go env: print Go environment information
- go fmt: run gofmt on package sources
- go get: download and install packages and dependencies
- go install: compile and install packages and dependencies
- go list: list packages
- go run: compile and run Go program
- go test: test packages
- go version: print Go version
- go vet: run go tool vet on packages

## Workspace Organization
3 subdirectories
- src : source code
- bin: executable programs
- pkg: build tools store compiled packages

## Downloading Packages
go get -u command generally retrieves latest version of each package
## Building Packages
go build command compiles each arg package

## Documenting Packages
Go style encourages good documentation of package APIs.

Go document comments are always complete sentences.

the go doc tool prints the declaration and doc comment of the entity specified on the command line

the godoc tool serves cross-linked HTML pages that provide the same info as go doc and more.

## Internal Packages
An internal package may be imported only by another package that is inside the tree rooted at the parent of the internal directory.

## Querying Packages
Go list tests whether a package is present in the workspace and prints its import path if so. 
-f flag customize the output format using the template language of package text/template. This command prints the transitive dependencies of the strconv package seperated by spaces. 
