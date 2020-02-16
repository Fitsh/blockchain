package main 

import (
    "../bolt"
    "log"
)

func main() {

    // 打开数据库
    db, err := bolt.Open("test.db", 0600, nil)
    if err != nil {
        log.Panic("打开数据库失败")
    }

    // 将要操作数据库（改写）
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

    // 写数据

    // 读数据
}
