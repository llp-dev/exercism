package phonenumber

import "fmt"

func IsDigit(r byte) bool {
    return '0' <= r && r <= '9'
}

func Number(phoneNumber string) (string, error) {
    n := ""

    for i := 0; i < len(phoneNumber); i++ {
        if IsDigit(phoneNumber[i]) {
            n = n + string(phoneNumber[i])
        }
    }

    if (n[0] == '1') {
        n = n[1:]
    }

    if (n[0] == '0' || n[0] == '1' || n[3] == '0' || n[3] == '1' || len(n) != 10) {
        return "", fmt.Errorf("Err")
    }

    return n, nil
}

func AreaCode(phoneNumber string) (string, error) {
    n, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    
    return n[:3], nil
}

func Format(phoneNumber string) (string, error) {
    n, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("(%s) %s-%s", n[0:3], n[3:6], n[6:10]), nil
}
