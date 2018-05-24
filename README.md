# Benchmark large json parsing


## 10mb.json

### Parsing using Golang[jq]

    $ time jq '.' data/10mb.json > /dev/null
    real    0m0.567s
    user    0m0.551s
    sys     0m0.016s


### Parsing using Elixir

    Time taken [Poison]: 1218.831ms
    Time taken [Jason]: 508.461ms

## 100mb.json (https://github.com/zemirco/sf-city-lots-json/blob/master/citylots.json)

    $ time jq '.' data/citylots.json > /dev/null

    real    0m14.436s
    user    0m13.992s
    sys     0m0.420s
