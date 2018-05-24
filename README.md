# Benchmark large json parsing

A few crude benchmarks for Elixir, Golang and Ruby

    Processor Intel(R) Core(TM) i7-5820K CPU @ 3.30GHz
    OS: Ubuntu 16.04

## 10mb.json

### Golang [jq]

    $ time jq '.' data/10mb.json > /dev/null
    real    0m0.567s
    user    0m0.551s
    sys     0m0.016s


### Elixir

    Time taken [Poison]: 1218.831ms
    Time taken [Jason]: 508.461ms

### Ruby

    time ruby app.rb

    real    0m0.220s
    user    0m0.203s
    sys     0m0.017s

## 100mb.json (https://github.com/zemirco/sf-city-lots-json/blob/master/citylots.json)

### Golang [jq]

    $ time jq '.' data/citylots.json > /dev/null

    real    0m14.436s
    user    0m13.992s
    sys     0m0.420s

## Elixir

      Time taken [Poison]: 32_640.87ms
      Time taken [Jason]: 11_602.128ms

## Ruby

      $ time ruby app.rb

      real    0m4.738s
      user    0m4.498s
      sys     0m0.240s
