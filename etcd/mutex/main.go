package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"

	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
	"time"
)

//etcd 实现分布式锁

func main(){
	lock()
}

//阻塞等待 锁释放
func lock(){
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"192.168.56.101:2379","192.168.56.101:22379","192.168.56.101:32379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// create two separate sessions for lock competition
	//"go.etcd.io/etcd/clientv3" 使用这个clientv3会报错
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// acquire lock for s1
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// wait until s1 is locks /my-lock/
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	//5秒后释放锁
	time.Sleep(time.Second*5)
	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
}

/*
//尝试加锁，如果锁被其他程序占用 则直接返回 不会阻塞等待锁释放
func trylock(){
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"192.168.56.101:2379","192.168.56.101:22379","192.168.56.101:32379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	// create two separate sessions for lock competition
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock")

	// acquire lock for s1
	if err = m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")
	if err = m2.TryLock(context.TODO()); err == nil {
		log.Fatal("should not acquire lock")
	}
	if err == concurrency.ErrLocked {
		fmt.Println("cannot acquire lock for s2, as already locked in another session")
	}

	if err = m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")
	if err = m2.TryLock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s2")
}

 */
