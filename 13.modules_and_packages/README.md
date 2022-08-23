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
