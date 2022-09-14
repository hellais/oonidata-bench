import inspect
import dataformat_dacite as dataformat
import typing
from typing import get_type_hints, get_origin, Union

def is_optional(t):
    if typing.get_origin(t) == Union and len(typing.get_args(t)) == 2:
        return any([arg == type(None) for arg in typing.get_args(t)])
    return False

def python_type_to_go(t):
    if t == dict[str, str]:
        return "map[string]string", "", str(t)
    elif t == str:
        return "string", "", str(t)
    elif t == dict:
        return "map[string]interface{}", "", str(t)
    elif t == bytes:
        return "[]byte", "", str(t)
    elif t == int:
        return "int", "", str(t)
    elif t == float:
        return "float32", "", str(t)
    elif inspect.isclass(t) and issubclass(obj, dataformat.BaseModel):
        return f"{t.__name__}", "", str(t)
    elif typing.get_origin(t) == list:
        list_entry_type = python_type_to_go(typing.get_args(t)[0])[0]
        return f"[]{list_entry_type}", "", str(t)
    elif is_optional(t):
        optional_type = list(filter(lambda x: x != type(None), typing.get_args(t)))[0]
        nested_type = python_type_to_go(optional_type)[0]
        return nested_type, ",omitempty", str(t)
    elif get_origin(t) == Union:
        return "interface{}", ",omitempty", str(t)
    return "interface{}", "", str(t)

def make_go_struct(name, obj):
    type_hints = get_type_hints(obj)
    print("type " + name + " struct {")
    for th_name, th in type_hints.items():
        golang_type, json_extra, orig_type = python_type_to_go(th)
        json_key = th_name + json_extra
        print(f"\t {th_name} {golang_type} `json:\"{json_key}\"` // {orig_type}")
    print("}")

for name, obj in inspect.getmembers(dataformat):
    if inspect.isclass(obj) and issubclass(obj, dataformat.BaseModel):
        make_go_struct(name, obj)

"""
get_type_hints(dataformat.BaseMeasurement)
{'annotations': dict[str, str], 'input': typing.Union[str, typing.List[str], NoneType], 'report_id': <class 'str'>, 'measurement_start_time': <class 'str'>, 'test_start_time': <class 'str'>, 'probe_asn': <class 'str'>, 'probe_network_name': typing.Optional[str], 'probe_cc': <class 'str'>, 'probe_ip': typing.Optional[str], 'resolver_asn': typing.Optional[str], 'resolver_ip': typing.Optional[str], 'resolver_network_name': typing.Optional[str], 'test_name': <class 'str'>, 'test_version': <class 'str'>, 'test_runtime': <class 'float'>, 'software_name': <class 'str'>, 'software_version': <class 'str'>, 'test_helpers': typing.Optional[dict], 'test_keys': <class 'oonidata.dataformat.BaseTestKeys'>, 'data_format_version': typing.Optional[str], 'measurement_uid': typing.Optional[str]}
isclass(dataformat.BaseModel, get_type_hints(oonidata.dataformat.BaseMeasurement)['test_keys'])
"""
