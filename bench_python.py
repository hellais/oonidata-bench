import io
import time
from dataformat_dacite import WebConnectivity as WebConnectivityDacite
from dataformat_mashumaro import WebConnectivity as WebConnectivityMashumaro
from dataformat_pydantic import WebConnectivity as WebConnectivityPydantic


ITER_COUNT = 10

import json
import ujson
import orjson
import dacite.core



def mashumoro_from_json(raw_msmt):
    msmt = orjson.loads(raw_msmt)
    WebConnectivityMashumaro.from_dict(msmt)


def pydantic_from_json(raw_msmt):
    msmt = orjson.loads(raw_msmt)
    WebConnectivityPydantic.parse_obj(msmt)


def dacite_from_json(raw_msmt):
    data = orjson.loads(raw_msmt)
    dacite.core.from_dict(data_class=WebConnectivityDacite, data=data)


benchmarks = {
    "json": json.loads,
    "ujson": ujson.loads,
    "orjson": orjson.loads,
    "mashumoro_from_orjson": mashumoro_from_json,
    "pydantic_from_json": pydantic_from_json,
    "dacite_from_json": dacite_from_json,
}


def bench_func(raw_msmt_list, func):
    iters = 0
    errs = 0
    t0 = time.monotonic()
    for _ in range(ITER_COUNT):
        for raw_msmt in raw_msmt_list:
            try:
                func(raw_msmt)
            except Exception as exc:
                print(exc)
                errs += 1
            iters += 1

    return time.monotonic() - t0, iters, errs


def main():

    raw_msmt_list = []
    with open("sample-file.jsonl") as in_file:
        for line in in_file:
            raw_msmt_list.append(line)

    for b_name, func in benchmarks.items():
        runtime, iters, errs = bench_func(raw_msmt_list, func)
        print(f"# {b_name}")
        print(f"  runtime: {runtime}")
        print(f"  iters/s: {iters/runtime}")
        print(f"  iters: {iters}")
        print(f"  errs: {errs}")

main()
