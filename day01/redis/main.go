package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/**
	本次使用本地虚拟机中的 docker redis
	虚拟机ip  192.168.56.101
    docker run -p  6379:6379 --privileged=true --name redis -v /data/docker/redis/conf/redis.conf:/etc/redis/redis.conf -v /data/docker/redis/data:/data -d redis-6379 redis-server /etc/redis/redis.conf --appendonly yes
 */

//go get -u github.com/go-redis/redis

var (
	rdb *redis.Client
)

func initRdb() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.56.101:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	str,err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("redis connect success:",str)
	return nil
}
//连接Redis哨兵模式
func initClient1()(err error){
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
//连接Redis集群
func initClient3()(err error){
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main(){
	if err := initRdb();err != nil{
		fmt.Println("redis connect err",err)
		return
	}

	redisExample2()
}

func redisExample1(){
	ok,err := rdb.Set("stu","tom",time.Second * 3600).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok)

	val,err := rdb.Get("user").Result()
	if err != nil {
		fmt.Println("get err:",err)
		return
	}
	fmt.Println("user:",val)
}

func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

/**

Pipeline
Pipeline 主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间（RTT）。

Pipeline 基本示例如下：

pipe := rdb.Pipeline()

incr := pipe.Incr("pipeline_counter")
pipe.Expire("pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
上面的代码相当于将以下两个命令一次发给redis server端执行，与不使用Pipeline相比能减少一次RTT。

INCR pipeline_counter
EXPIRE pipeline_counts 3600
也可以使用Pipelined：

var incr *redis.IntCmd
_, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
	incr = pipe.Incr("pipelined_counter")
	pipe.Expire("pipelined_counter", time.Hour)
	return nil
})
fmt.Println(incr.Val(), err)
在某些场景下，当我们有多条命令要执行时，就可以考虑使用pipeline来优化。

事务
Redis是单线程的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，例如在它们之间交替执行。但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。

在这种场景我们需要使用TxPipeline。TxPipeline总体上类似于上面的Pipeline，但是它内部会使用MULTI/EXEC包裹排队的命令。例如：

pipe := rdb.TxPipeline()

incr := pipe.Incr("tx_pipeline_counter")
pipe.Expire("tx_pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
上面代码相当于在一个RTT下执行了下面的redis命令：

MULTI
INCR pipeline_counter
EXPIRE pipeline_counts 3600
EXEC
还有一个与上文类似的TxPipelined方法，使用方法如下：

var incr *redis.IntCmd
_, err := rdb.TxPipelined(func(pipe redis.Pipeliner) error {
	incr = pipe.Incr("tx_pipelined_counter")
	pipe.Expire("tx_pipelined_counter", time.Hour)
	return nil
})
fmt.Println(incr.Val(), err)
Watch
在某些场景下，我们除了要使用MULTI/EXEC命令外，还需要配合使用WATCH命令。在用户使用WATCH命令监视某个键之后，直到该用户执行EXEC命令的这段时间里，如果有其他用户抢先对被监视的键进行了替换、更新、删除等操作，那么当用户尝试执行EXEC的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。

Watch(fn func(*Tx) error, keys ...string) error
Watch方法接收一个函数和一个或多个key作为参数。基本使用示例如下：

// 监视watch_count的值，并在值不变的前提下将其值+1
key := "watch_count"
err = client.Watch(func(tx *redis.Tx) error {
	n, err := tx.Get(key).Int()
	if err != nil && err != redis.Nil {
		return err
	}
	_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Set(key, n+1, 0)
		return nil
	})
	return err
}, key)
最后看一个官方文档中使用GET和SET命令以事务方式递增Key的值的示例：

const routineCount = 100

increment := func(key string) error {
	txf := func(tx *redis.Tx) error {
		// 获得当前值或零值
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		// 实际操作（乐观锁定中的本地操作）
		n++

		// 仅在监视的Key保持不变的情况下运行
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// pipe 处理错误情况
			pipe.Set(key, n, 0)
			return nil
		})
		return err
	}

	for retries := routineCount; retries > 0; retries-- {
		err := rdb.Watch(txf, key)
		if err != redis.TxFailedErr {
			return err
		}
		// 乐观锁丢失
	}
	return errors.New("increment reached maximum number of retries")
}

var wg sync.WaitGroup
wg.Add(routineCount)
for i := 0; i < routineCount; i++ {
	go func() {
		defer wg.Done()

		if err := increment("counter3"); err != nil {
			fmt.Println("increment error:", err)
		}
	}()
}
wg.Wait()

n, err := rdb.Get("counter3").Int()
fmt.Println("ended with", n, err)
 */

