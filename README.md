# go-cli

## Requirements

[Go](http://golang.org) 1.22 or newer.

## Installation

```text
go get github.com/xagero/go-cli
```

## Example

```text
first := command.Construct("app:first", "First command")
first.SetCallback(func() error {
    fmt.Println("I am first command")
    return nil
})

second := command.Construct("app:second", "Second command")
second.SetCallback(func() error {
    fmt.Println("I am second command")
    return nil
})

name := "Console"
desc := "Simple console application"
version := "v0.1"

console := cli.Construct(name, desc, version)
console.AddCommand(first)
console.AddCommand(second)

if err := console.Run(context.Background(), os.Args); err != nil {
    fmt.Printf("Error encountered: %v\n", err)
}
```

## App command run

```text
./myapp app:first
```


## Todo

Project still in development

- [ ] Improve help and group command
- [ ] Console builtin options
- [ ] Improve option requirements
- [ ] Color in console
- [x] Text align