package main

import (
	GT "../goTools"
	"fmt"
)


func main() {

	srcOMao := make(map[string]map[string]map[string]interface{})

	srcMap := make(map[string]map[string]interface{})
	subMap := make(map[string]interface{})
	subMap["a"] = 1
	subMap["b"] = "qwe"
	subMap["c"] = map[string]int{
		"q" : 10,
		"w" : 11,
	}
	srcMap["sub1"] = subMap
	srcOMao["ub1"] = srcMap

	var dste map[string]map[string]map[string]interface{}
	thisByte, theErr := GT.FGobEncode(srcOMao, 0)
	fmt.Println("gob 编码后的结果", thisByte, theErr)

	thisErr := GT.FGobDecodeMSSSItf(&dste)

	fmt.Println("gob 解码后的结果", dste, thisErr)

	secondMap := make(map[string]int)
	secondMap["qwe"] = 1

	retMap := make(map[string]int)
	_, theErr = GT.FGobEncode(secondMap)
	thisErr = GT.FGobDecodeMSInt(&retMap)
	fmt.Println("注册类型后的 gob 转换结果", retMap, theErr, thisErr)

	thisStr, thisErr := GT.FJsonDumps(srcMap)
	fmt.Println("json dumps 后的结果", thisStr, thisErr)

	var tempItf interface{}
	secondErr := GT.FJsonLoads(thisStr, &tempItf)
	fmt.Println("字符串解析 json 的结果", tempItf, secondErr)

	ret1, ok1 := tempItf.(map[string]map[string]interface{})
	fmt.Println("interface 全部自动推断的结果", ret1, ok1)

	ret, ok := tempItf.(map[string]interface{})
	fmt.Println("interface 自动推断一层的结果", ret, ok)


}

