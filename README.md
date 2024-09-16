# Golang learning

my Golang challenge journey repository.

Getting started with <https://roadmap.sh/golang> here, and to follow a few simple tutorial to master the basic, and get some confidence about it.

## tips

```sh

# this is compile and test run, need the dot (current folder)
go run .
# this would help adding the module
go mod tidy
```



## Source for Continuous learning
1. to focus on the Go Module part. <https://go.dev/doc/tutorial/create-module> to understand how to organise and import.
2. In Go . code executed as an application must be in a `main` package. the primary one is the **hello** and the referenced one is **greetings** folder.
3. adaptation required for local development, such that `go mod edit -replace divineforge.com/greetings=../greetings` then as usual the `go mod tidy` to get things up. See the **go.mod** file afterwards for its updated content.
4. Error Handlings; <https://go.dev/doc/tutorial/handle-errors>
