# equation-solver
This simple CLI application can be used to solve quadratic equations

![image](https://user-images.githubusercontent.com/24240873/152986109-30bcc0f1-aaa6-4881-8a94-82b83ccd0f11.png)

## Installation

Using this package requires a working Go environment. See the install instructions for Go.

Install using the following command:

```
$ go get -u github.com/JenyaFTW/equation-solver
```

Make sure your `PATH` includes `$GOPATH/bin` or else `equation-solver` command will not work

## Build using source

Make sure to have this repository cloned and open and run `go install`

## Usage

This application includes two modes: interactive and file modes

To use interactive mode, simply call `equation-solver` command and import the parameters of your equation

```
$ equation-solver
a = 1
b = 2
c = 3
```

To use file mode, make sure you have a file containing real numbers with a space between each, like this:

```
1 2 3
```

Then run the following command:

```
$ equation-solver file.txt
```

## Test revert commit

[74032c2e13ed9a8370a6d3e6487a3ff766ab9ce7](https://github.com/JenyaFTW/equation-solver/commit/74032c2e13ed9a8370a6d3e6487a3ff766ab9ce7)
