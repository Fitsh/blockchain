package main 

import (
    "../bolt"
    "log"
    "fmt"
)

func main() {

    // 打开数据库
    db, err := bolt.Open("test.db", 0600, nil)
    if err != nil {
        log.Panic("打开数据库失败")
    }

    defer db.Close()

    // 将要操作数据库（改写）
    // 写数据
    db.Update(func(tx *bolt.Tx) error {

        // 找到bucket(如果没有则创建)
        bucket := tx.Bucket([]byte("b1"))
        if bucket == nil {
            // 没有抽屉，需要创建
            bucket, err = tx.CreateBucket([]byte("b1"))
            if err != nil {
                log.Panic("创建bucket(b1)失败")
            }
        }
        
        bucket.Put([]byte("111"), []byte("hello"))
        bucket.Put([]byte("222"), []byte("world"))
        return nil
    })

    // 读数据
    db.View(func(tx *bolt.Tx) error {
        // 找到抽屉，没有直接报错
        bucket := tx.Bucket([]byte("b1"))
        if bucket == nil {
            log.Panic("bucket b1 不应该为空")
        }

        // 直接读取数据
        v1 := bucket.Get([]byte("111"))
        v2 := bucket.Get([]byte("222"))
        fmt.Printf("v1: %s\n", v1)
        fmt.Printf("v2: %s\n", v2)
        return nil
    })
}
