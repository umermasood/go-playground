# 13. Modules & Packages

## Modules
Modules in Go are just a collection of packages with a go.mod file at the root level of your project.

- `go get`
    - downloads the source code to module cache in `GOPATH/pkg/mod`
    - [Dependencies] updates the `go.mod` to require those packages
    - doesn't build and install packages anymore
- `go build`
  - builds an executable of the program in your root directory
  - `-o` flag can be used to specify custom build directory
- `go install`
  - installs your program executable in `GOPATH/bin`
  - recommended way to build and install your packages
- List all dependencies using `go list -m all`
- Remove unused dependencies using `go mod tidy`

### Vendoring
Create copies of the third party packages being used in the project.

- Vendoring is done using `go mod vendor`


## Traditional Dependency Management in Go
Before *Go Modules* were supported by Go, we had to work with packages and for that we needed our source code to be
present inside the `GOPATH`.

Drawbacks
- This approach forced the gophers to relocate all projects inside one directory or having to
constantly change `GOPATH` variable to switch between different projects.
- You can't install multiple versions of the same package
- Traditional package management makes the package backward incompatible

Things to remember
- `go get` downloads the dependency source code inside `GOPATH`
- `go build` creates binary executables and package archives inside `GOPATH`

Go doesn't have central repo for downloading packages like npm

`go get`
- Clones the latest commit of the master branch of the package repo
- Cannot install the package at a specific commit / (version)
---

## What we want to do

- Work on Go projects in any directory -> Gives flexibility
- To be able to install specific / multiple versions of a dependency -> To avoid breaking changes
- To be able to import specific / multiple versions of a dependency -> To avoid breaking changes
- A file that lists the dependencies for Go project -> like `package.json`

With Go Modules, we are able to do all these.

## What is a Module?
A module is just a collection of packages:
- Distribution Packages
- Executable Packages

Module can also be called a package that contains nested packages.

Module:
- Must be a VCS Repo
- Should contain one or more packages
- Package should contain one or more `.go` files in a single directory

### Publishing a Module
- We can create alias for a specific commit hash, this is called tagging.
- Go supports **SemVer (Semantic Versioning)** for version management of modules
- we can use `git tag vX.Y.Z` to tag the latest commit in a branch.
- X is major (introduces breaking changes), Y is minor (Backwards compatible additional functionalities), and Z is patch
version (implements bug fixes, optimizations, etc).


### Update Module / Dependencies
- Use `go get -u` to update dependency modules
- Use `go get -u=patch` to update to patch fixes only, even if the minor updates are available
- Use `go get module-path@version` to download and use a specific module version


### Minimal Version Selection


#### Diamond Dependency Problem
- Multiple versions of a dependency with a common major version should be backwards compatible
- i.e. in v1.0.1 vs v1.0.2, version 2 is backwards compatible with v1
- If we use v1.0.2 for our final build, everything should work fine
- Go calls this minimal version selection (which basically means to select maximal version between dependencies)
