# Benchmark large json parsing

A few crude benchmarks for Elixir, Golang, C++, and Ruby

    Processor: 2.5 GHz Quad-Core Intel Core i7
    Memory: 16 GB 1600 MHz DDR3
    OS: macOS 10.15.3

## 10mb.json

### CPP

    $ time _builds/rapidjson-sax < ../../data/10mb.json

    real	0m0.031s
    user	0m0.026s
    sys	0m0.003s

    $ time _builds/rapidjson-structure < ../../data/10mb.json

    real	0m0.039s
    user	0m0.032s
    sys	0m0.006s

### Golang

    $ time ./naive < ../data/10mb.json

    real	0m0.386s
    user	0m0.408s
    sys	0m0.060s

    $  time ./streamed < ../data/10mb.json

    real	0m0.182s
    user	0m0.196s
    sys	0m0.020s

    $  time ./streamed-struct < ../data/10mb.json

    real	0m0.081s
    user	0m0.074s
    sys	0m0.010s

    $  time ./streamed-struct-jsoniter < ../data/10mb.json

    real	0m0.036s
    user	0m0.030s
    sys	0m0.006s

### Elixir

    Time taken [Poison]: 857.349ms
    Time taken [Jason]: 438.837ms

### Ruby

    $ time ruby app.rb

    real	0m0.299s
    user	0m0.254s
    sys	0m0.039s

## 100mb.json (https://github.com/zemirco/sf-city-lots-json/blob/master/citylots.json)

### CPP

    $ time _builds/rapidjson-sax < ../../data/citylots.json

    real	0m0.463s
    user	0m0.435s
    sys	0m0.026s

    $ time _builds/rapidjson-structure < ../../data/citylots.json

    real	0m0.719s
    user	0m0.602s
    sys	0m0.117s

### Golang

    $ time ./naive < ../data/citylots.json

    real	0m7.613s
    user	0m9.015s
    sys	0m0.693s

    $  time ./streamed < ../data/citylots.json

    real	0m4.448s
    user	0m5.558s
    sys	0m0.420s

    $  time ./streamed-struct < ../data/citylots.json

    real	0m1.762s
    user	0m1.680s
    sys	0m0.194s

    $  time ./streamed-struct-jsoniter < ../data/citylots.json

    real	0m0.734s
    user	0m0.686s
    sys	0m0.065s

## Elixir

    Time taken [Poison]: 30483.717ms
    Time taken [Jason]: 12806.43ms

## Ruby

    $ time ruby app.rb

    real	0m6.352s
    user	0m5.977s
    sys	0m0.360s



# Building instructions:

## C++

### RapidJson

Just run the `build.sh` script from within the `cpp/rapidjson` directory and it will output two binaries into the `_builds` directory, one for sax-style parsing, the other for structure/document style parsing.

To run it the C++ programs will accept it via stdin, so don't pass it in as an argument, but instead pipe it into to it like via `time _builds/rapidjson-structure < ../../data/citylots.json` or so.

## Go

This includes four versions of a Go JSON parser.
- Three use the standard "encoding/json" library, and one uses the third-party jsoniter "github.com/json-iterator/go" (not the fastest, but easiest to swap in).
- Two decoding to generic `map[string]interface{}` and two defining the property shaped structs to hold deserialized results.
- One reading the file into memory then unmarshalling and three streaming the file through the decoder.

Build all four separately since they use the same package name.

    go build naive.go
    go build streamed.go
    go build streamd-struct.go
    GOPATH=`pwd` go get -u github.com/json-iterator/go && GOPATH=`pwd` go build streamed-struct-jsoniter.go

## Elixir

`mix test`
