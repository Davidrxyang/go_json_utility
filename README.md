# go_json_utility
json parser in Go

to build the program: 

```
go build
```

you should now have the executable "go_json_utility"

the program runs in two modes, read and write:

read: reads a json file and outputs to the terminal

```
./go_json_utility FILENAME.json
```

a sample file, telecaster.json, is provided

```
./go_json_utility telecaster.json
```


write: takes CLI input from the user and writes input into a json file

```
./go_json_utility FILENAME.json
```