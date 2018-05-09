# Go 工具集说明

## 1、RPC - gob 

位置：代码在 `gobTools.go` 文件中。

使用说明：
> 调用 `FGobEncode()` 将对象进行编码。
> 调用 `FGobDecodeMSInt()` 等方法将对应类型的对象解码。


## 2、RPC - json

位置：代码在 `jsonTools.go` 文件中。

使用说明：
> 调用 `FJsonDumps(src interface{})` 将 json 对象进行编码。
> 调用 `FJsonLoads(src string, dst *interface{})` 将字符串解码为 json。


