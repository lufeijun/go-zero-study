package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/pkg/profile"
)

// https://www.cnblogs.com/hobbybear/p/18059425

func main() {
	// cpu profile
	// cpu()

	// mem profile
	// mem()

	// block profile
	// block()

	// mutex profile
	// mutex()

	// goroutine profile
	// goroutine()

	// thread profile
	thread()
}

// cpu profile ，查看 go tool pprof -http=:9999 cpu.pprof
func cpu() {
	p := profile.Start(profile.CPUProfile,
		profile.ProfilePath("profile"),
		profile.NoShutdownHook,
	)
	defer p.Stop()
	go busyCpu()
	go busyCpu()

	time.Sleep(1 * time.Second)
}

func busyCpu() {
	i := uint(1000000)
	for {
		log.Println("sum number", i, Add(i, i+1))
		i++
	}
}

func Add(a, b uint) uint {
	return a + b
}

func mem() {
	p := profile.Start(profile.MemProfile,
		profile.ProfilePath("profile"),
		profile.NoShutdownHook,
	)
	defer p.Stop()
	go allocMem()

	time.Sleep(5 * time.Second)
}
func allocMem() {
	buf := []byte{}
	mb := 1024 * 1024

	for {
		buf = append(buf, make([]byte, mb)...)
		log.Println("total allocated memory", len(buf))
		time.Sleep(time.Second)
	}
}

// 阻塞分析
func block() {
	runtime.SetBlockProfileRate(1)
	p := profile.Start(profile.BlockProfile,
		profile.ProfilePath("profile"),
		profile.NoShutdownHook)
	defer p.Stop()

	mLock := sync.Mutex{}
	mLock.Lock()

	go mockBlock(&mLock)
	go mockBlock(&mLock)

	time.Sleep(time.Second * 2)
	mLock.Unlock()
	time.Sleep(time.Second * 3)
	fmt.Println("End")
}
func mockBlock(l *sync.Mutex) {
	l.Lock()
	defer l.Unlock()
	fmt.Println("获取锁")
	time.Sleep(time.Second * 1)
}

// 锁
func mutex() {
	runtime.SetMutexProfileFraction(1)
	p := profile.Start(profile.MutexProfile,
		profile.ProfilePath("profile"),
		profile.NoShutdownHook)
	defer p.Stop()

	mLock := sync.Mutex{}
	mLock.Lock()

	go mockBlock(&mLock)
	go mockBlock(&mLock)

	time.Sleep(time.Second * 2)
	mLock.Unlock()
	time.Sleep(time.Second)
	fmt.Println("End")
}

// 协程
func goroutine() {
	go mockGo()
	go mockGo()

	p := profile.Start(profile.GoroutineProfile,
		profile.ProfilePath("profile"),
		profile.NoShutdownHook)
	p.Stop()

	time.Sleep(time.Second * 5)
}

func mockGo() {
	time.Sleep(time.Second * 3)
	fmt.Println("协程执行")
}

// thread 线程
func thread() {
	go mockGo()
	go mockGo()
	p := profile.Start(profile.ThreadcreationProfile,
		profile.ProfilePath("profile"),
		profile.NoShutdownHook)
	p.Stop()

	time.Sleep(time.Second * 5)

}
