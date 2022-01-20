package main

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

func BatchInsert(size int, db int) {
	//start := time.Now()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.33.xx.xx:6379",
		Password: "123456", // no password set
		DB:       db,       // use default DB
	})

	//for i := 0; i < 100000; i++ {
	//  err := rdb.Set(ctx, randStr(20), randStr(20), 0).Err()
	//  if err != nil {
	//      println(err)
	//      return
	//  }
	//}
	var wg sync.WaitGroup
	wg.Add(size)
	worker := NewDataWorker()
	go func() {
		for i := 0; i < size; i++ {
			go addData(worker, &wg)
		}
		// 等待生成完成，并且写入空对象
		wg.Wait()
		worker.JobChannel <- &DataProduction{}
	}()
	worker.Start(rdb)
	//end := time.Now()
	//useTime := end.Sub(start)
	//fmt.Println(useTime.String())

}

// addData 填充测试数据
func addData(worker DataWorker, wg *sync.WaitGroup) {

	// Generate a snowflake ID.
	for i := 0; i < 100000; i++ {
		worker.JobChannel <- NewDataProduction()
	}
	wg.Done()
}

// DataProduction 测试数据结构
type DataProduction struct {
	key   string
	value string
}

// NewDataProduction 生成测试数据方法
func NewDataProduction() *DataProduction {
	dp := DataProduction{
		key:   randStr(25),
		value: randStr(25),
	}
	return &dp
}

// ConsumptionData 消费数据
var ctx = context.Background()

func (p DataProduction) ConsumptionData(rdb *redis.Client) {
	err := rdb.Set(ctx, p.key, p.value, 0).Err()
	if err != nil {
		panic(err)
	}
}

// DataWorker 工作对象
type DataWorker struct {
	JobChannel chan *DataProduction
}

// NewDataWorker 初始化工作对象
func NewDataWorker() DataWorker {
	return DataWorker{JobChannel: make(chan *DataProduction)}
}

func (w DataWorker) Start(rdb *redis.Client) {
	var wg sync.WaitGroup
	wg.Add(1)
	// 四协程
	//for i := 0; i < 4; i++ {
	go func(rdb *redis.Client, group *sync.WaitGroup) {
		for {
			select {
			case job := <-w.JobChannel:
				if job.key == "" {
					//参考  如果为空，则说明已经生成完毕
					group.Done()
					return
				}
				job.ConsumptionData(rdb)

			}
		}
	}(rdb, &wg)
	//}
	// 等待数据消费完毕再退出
	wg.Wait()
}

// 随机字符串
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func randStr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
