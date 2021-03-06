# go-atcoder
for AtCoder user golang

## Requirement

- [github.com/Tatamo/atcoder-cli](https://github.com/Tatamo/atcoder-cli)
- [github.com/online-judge-tools/oj](https://github.com/online-judge-tools/oj)

```shell
$ pip3 install online-judge-tools
$ npm install -g atcoder-cli
$ exec $SHELL -l
$ oj login https://atcoder.jp/
$ acc login
$ acc check-oj
```

## Installation

```shell
$ cp -r go-template `acc config-dir`
# confirm templates
$ acc templates
NAME         SUBMIT-PROGRAM
go-template  main.go
$ acc config default-template go-template
```

## Usage

```shell
# example. if contest name is abc051
$ acc new abc051
$ cd abc051/a
$ vim main.go # write solution
$ oj t -c "go run ./main.go" -d tests # recommended to set an alias
$ acc submit main.go
```
