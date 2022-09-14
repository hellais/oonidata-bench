# Benchmark results for oonidata

This repo exists to demonstrate some performance benchmarks in loading and
validating OONI measurements using several available options (in go and
python).

To re-run the benchmarks you can do:
```
poetry install
./bench.sh
```

## Results

Here are the results from running it on a mac mini with the following specs:
```
CPU: 3,2 GHz 6-Core Intel Core i7
Memory: 64 GB 2667 MHz DDR4
```

```
# json
  runtime: 0.5277967529837042
  iters/s: 3164.096767475872
  iters: 1670
  errs: 0
# ujson
  runtime: 0.4175252459826879
  iters/s: 3999.758136946871
  iters: 1670
  errs: 0
# orjson
  runtime: 0.3135276829707436
  iters/s: 5326.4834038780355
  iters: 1670
  errs: 0
# mashumoro_from_orjson
  runtime: 0.6252308230032213
  iters/s: 2671.0135498092613
  iters: 1670
  errs: 0
# pydantic_from_orjson
  runtime: 1.3429232999915257
  iters/s: 1243.5557563194698
  iters: 1670
  errs: 0
# dacite_from_orjson
  runtime: 19.846701907983515
  iters/s: 84.1449631149157
  iters: 1670
  errs: 0
# go_json
  runtime: 1.159089
  iters/s: 1440.787122
  iters: 1670
  errs: 0
```
