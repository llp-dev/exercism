package isbnverifier

func ToDigit(ch byte) int {
    if (ch == 'X') {
        return 10
    }
    return int(ch - '0')
}

func Validate(v int) bool {
    return v >= 0 && v <= 9
}

func IsValidISBN(isbn string) bool {
    var d1, d2, d3, d4, d5, d6, d7, d8, d9, d10 int
    
    if (len(isbn) == 13) {
    d1 = ToDigit(isbn[0])
    d2 = ToDigit(isbn[2])
    d3 = ToDigit(isbn[3])
    d4 = ToDigit(isbn[4])
    d5 = ToDigit(isbn[6])
    d6 = ToDigit(isbn[7])
    d7 = ToDigit(isbn[8])
    d8 = ToDigit(isbn[9])
    d9 = ToDigit(isbn[10])
    d10 = ToDigit(isbn[12])
    } else if (len(isbn) == 10) {
    d1 = ToDigit(isbn[0])
    d2 = ToDigit(isbn[1])
    d3 = ToDigit(isbn[2])
    d4 = ToDigit(isbn[3])
    d5 = ToDigit(isbn[4])
    d6 = ToDigit(isbn[5])
    d7 = ToDigit(isbn[6])
    d8 = ToDigit(isbn[7])
    d9 = ToDigit(isbn[8])
    d10 = ToDigit(isbn[9])
    } else {
        return false
    }

    if (Validate(d1) == false || Validate(d2) == false || Validate(d3) == false || Validate(d4) == false || Validate(d5) == false || Validate(d6) == false || Validate(d7) == false || Validate(d8) == false || Validate(d9) == false) {
        return false
    }
    
    return (d1 * 10 + d2 * 9 + d3 * 8 + d4 * 7 + d5 * 6 + d6 * 5 + d7 * 4 + d8 * 3 + d9 * 2 + d10 * 1) % 11 == 0
}
