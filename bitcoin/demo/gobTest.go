package main

import (
    "bytes"
    "fmt"
    "encoding/gob"
    "log"
)

type Person struct {
    Name string
    Age uint
}

func main() {
    // 定义一个结构 Person
    var xiaoMing Person
    xiaoMing.Name = "xiaoming"
    xiaoMing.Age = 8

    // 编码的数据放到bytes.Buffer
    var buffer bytes.Buffer

    // 使用gob字节化（编码)得到字节流
    // 1. 定义一个编码器
    // 2. 使用编码器进行编码
    encoder := gob.NewEncoder(&buffer)
    err := encoder.Encode(&xiaoMing)
    if err != nil {
        log.Panic("编码出错")
    }
    fmt.Printf("编码后的小明, %v\n", buffer.Bytes())

    // 使用gob 进行反序列化（解码)得到Person结构
    // 1. 定义一个解码器
    // 2. 使用解码器进行解码
    var daMing Person
    decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
    err = decoder.Decode(&daMing)
    if err != nil {
        log.Panic("解码出错")
    }

    fmt.Printf("解码后的小明 %v\n", &daMing)

}
