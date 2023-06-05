/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/06/02 21:49
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3}
	wg := sync.WaitGroup{}
	wg.Add(3)
	for _, v := range nums {
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()
}
