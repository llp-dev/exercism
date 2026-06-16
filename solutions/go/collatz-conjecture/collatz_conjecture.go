package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	c := 0

    if (n <= 0) {
        return c, fmt.Errorf("Err")
    }
    
    for {
        if (n == 1) {
            return c, nil
        }
        if (n % 2 == 0) {
            n = n / 2
        } else {
            n = (n * 3) + 1
        }
        c++
    }
}
