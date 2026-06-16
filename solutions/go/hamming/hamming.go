package hamming

import "fmt"

func Distance(a, b string) (int, error) {
    c := 0
    len_a := len(a)
    len_b := len(b)
    if (len_a != len_b) {
        return 0, fmt.Errorf("Err")
    }
    for i := 0; i < len_a; i++ {
        if (a[i] != b[i]) {
            c++
        }
    }

    return c, nil
}
