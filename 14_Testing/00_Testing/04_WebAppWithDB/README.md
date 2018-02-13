
# Real MongoDB

Run docker image - `docker run -d -p 27017:27017 mongo`

Create database - `test`

Create collection - `texts`

Insert one item 

```JSON
{
    "_id" : ObjectId("5a7db3a3e0f8067329415b10"),
    "text" : "some text from database"
}
```

# Run App

`go run *.go`

and open webpage `http://localhost:8086/home`

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
cd $GOPATH/src/github.com/lukaskellerstein/GolangTrainingProjects/14_Testing/00_Testing/04_WebAppWithDB/
```

Test coverage to HTML file

```Shell
go test -coverprofile=coverage.out && \
go tool cover -html=coverage.out
```




