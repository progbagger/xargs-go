# xargs

`xargs` is a Unix command-like utility that transforms standart input into sequence of arguments for the command specified in `xargs` arguments.

## How to build?

Just execute `make` command in the **src** project directory.

If the `package <name> is not in src` error occures then execute in the root if the project:

```shell
go work init
go work use src/xargs
```

## How to use?

There are no flags in this utility. Just pass the output of the previous executed command to this utility via pipe or using any other method of attaching **stdin** and type needed command in `xargs` arguments.

## Example

This example is executed in the **src** directory after executing `make` command.

```shell
-> % echo ".. ../build ." | ../build/xargs ls
.:
Makefile  xargs

..:
LICENSE  README.md  build  go.work  src

../build:
xargs
```
