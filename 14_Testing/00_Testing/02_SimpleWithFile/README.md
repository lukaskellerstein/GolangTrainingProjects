

This sample showed simple app which needs File and method with CPU leaks.


# Run Tests 

```Shell
go test 
```

or more info 

```Shell
go test -v 
```

# Test coverage 

```Shell
go test -cover
```

# Test coverage - in HTML graphics

First of all you must copy this project in your GOPATH, so run

```Shell
make
```

Then, you must locate into GOAPTH copy of this project

```Shell
cd $GOPATH/src/github.com/lukaskellerstein/GolangTrainingProjects/14_Testing/00_Testing/01_Simple/
```

Test coverage to HTML file

```Shell
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```