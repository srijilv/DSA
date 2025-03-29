package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type User struct {
	name string
}

func (u User) create(name string, ch chan string) {
	ch <- fmt.Sprintf("hello world - %s", name)
}

func main() {
	count := 3
	wg := sync.WaitGroup{}
	ch := make(chan string, 1)
	u := User{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			u.create(fmt.Sprintf("name%d", i), ch)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}

	wg.Wait()

	os.Exit(1)
}

type Pallindrome struct {
	counter int
	mt      sync.Mutex
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano()) // Ensure different results on each run
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func mainMultiplePallindrome() {
	words := []string{"Aka", "toMato", "malAyalam", "bbbbbb"}
	pall := &Pallindrome{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		words = append(words, randomString(5))
	}

	fmt.Printf("words: %+v\n", words)

	wg := sync.WaitGroup{}

	for _, w := range words {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pall.checkPallindrome(strings.ToLower(w))
		}()
	}
	wg.Wait()
	fmt.Println("total pallindrome count", pall.counter)

}

func (p *Pallindrome) checkPallindrome(str string) {
	strSlices := []byte(str)
	compSlices := []byte{}
	fmt.Println("processing input string:", str)

	for i := 0; i < len(strSlices); i++ {
		compSlices = append(compSlices, strSlices[len(strSlices)-i-1])
	}
	if str == string(compSlices) {
		fmt.Println(str, " is pallindrome")
		p.mt.Lock()
		defer p.mt.Unlock()
		p.counter++
	}
}
