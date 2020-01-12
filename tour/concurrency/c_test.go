//
// mix of examples and exercises from 'A Tour of Go': section 'Concurrency'
//

package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
	"sort"
	"sync"
	"testing"
)

func TestExplicitChan(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	for i := 0; i < 100; i++ {
		assert.Equal(t, i, <-ch)
	}
}

func TestRangeChan(t *testing.T) {
	var ch chan int = make(chan int)
	var j int = 0

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}

		close(ch)
	}()

	for i := range ch {
		assert.Equal(t, j, i)
		j++
	}

}

func TestBufferedChan(t *testing.T) {
	var ch chan int = make(chan int, 1)

	ch <- 0
	assert.Equal(t, 0, <-ch)

	ch <- 1
	assert.Equal(t, 1, <-ch)
}

func TestSelectChan(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}

		c2 <- 100
	}()

	go func() {
		v := <-c3
		c3 <- v
	}()

	n := 0

	for {
		select {
		case v := <-c1:
			assert.Equal(t, n, v)
			n++
		case v := <-c2:
			assert.Equal(t, 100, v)
			c3 <- v
		case v := <-c3:
			assert.Equal(t, 100, v)
			return
		}
	}
}

// Exercise: Equivalent Binary Trees

func Walk(t *tree.Tree, ch chan int, root bool) {
	if t == nil {
		return
	}

	Walk(t.Left, ch, false)
	ch <- t.Value
	Walk(t.Right, ch, false)

	if root {
		close(ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	var c1 chan int = make(chan int)
	var b1 [10]int
	var i1 = 0

	go Walk(t1, c1, true)

	for v := range c1 {
		b1[i1], i1 = v, i1+1
	}

	var c2 chan int = make(chan int)
	var b2 [10]int
	var i2 = 0

	go Walk(t2, c2, true)

	for v := range c2 {
		b2[i2], i2 = v, i2+1
	}

	return b1 == b2
}

func TestBinaryTrees(t *testing.T) {
	assert.Equal(t, true, Same(tree.New(1), tree.New(1)))
	assert.Equal(t, true, Same(tree.New(2), tree.New(2)))
	assert.Equal(t, true, Same(tree.New(3), tree.New(3)))
	assert.Equal(t, true, Same(tree.New(4), tree.New(4)))

	assert.Equal(t, false, Same(tree.New(1), tree.New(2)))
	assert.Equal(t, false, Same(tree.New(1), tree.New(3)))
	assert.Equal(t, false, Same(tree.New(1), tree.New(4)))
	assert.Equal(t, false, Same(tree.New(2), tree.New(3)))
	assert.Equal(t, false, Same(tree.New(2), tree.New(4)))
	assert.Equal(t, false, Same(tree.New(3), tree.New(4)))
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string, inc int) {
	for i := 0; i < inc; i++ {
		c.mux.Lock()
		c.v[key]++
		c.mux.Unlock()
	}
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func TestMutex(t *testing.T) {
	var a = SafeCounter{v: make(map[string]int)}
	var b = SafeCounter{v: make(map[string]int)}
	var c = SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			a.Inc("test1key", 1)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			b.Inc("test2key", 10)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Inc("test3key", 100)
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, 100, a.Value("test1key"))
	assert.Equal(t, 1000, b.Value("test2key"))
	assert.Equal(t, 10000, c.Value("test3key"))
}

// Exercise: Web Crawler

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type Visited struct {
	v   map[string]bool
	mux sync.Mutex
}

func (c *Visited) visit(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	if u := c.v[url]; u == false {
		c.v[url] = true
		return false
	}

	return true
}

func Crawl(url string, depth int, fetcher Fetcher, visited *Visited, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if visited.visit(url) {
		return
	}

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, visited, wg)
	}

	return
}

func TestWebCrawler(t *testing.T) {
	var results = map[int][]string{
		0: []string{},
		1: []string{"https://golang.org/"},
		2: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/"},
		3: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
		4: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
		5: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
		6: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
		7: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
		8: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
		9: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
	}

	for i := 0; i < len(results); i++ {
		var visited Visited = Visited{v: make(map[string]bool)}
		var wg sync.WaitGroup

		wg.Add(1)
		go Crawl("https://golang.org/", i, fetcher, &visited, &wg)
		wg.Wait()

		keys := make([]string, 0, len(visited.v))
		for k := range visited.v {
			keys = append(keys, k)
		}

		sort.Strings(keys)
		assert.Equal(t, results[i], keys)
	}
}
