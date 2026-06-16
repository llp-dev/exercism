package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    res := make(map[string]int)
    res["quarter_of_a_dozen"] = 3
    res["half_of_a_dozen"] = 6
    res["dozen"] = 12
    res["small_gross"] = 120
    res["gross"] = 144
    res["great_gross"] = 1728
    return res
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    v := units[unit]
    if (v == 0) {
        return false
    }
    b := bill[item]
    bill[item] = b + v
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    b := bill[item]
    if (b == 0) {
        return false
    }
    v := units[unit]
    if (v == 0) {
        return false
    }
    if (b < v) {
        return false
    }
    if (b == v) {
        delete(bill, item)
        return true
    }
    bill[item] = b - v
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    b := bill[item]
    if (b == 0) {
        return b, false
    }
    return b, true
}
