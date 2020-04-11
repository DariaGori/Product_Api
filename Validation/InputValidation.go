package InputValidation

import "strconv"

// Check whether the string is fully numeric
func IsNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}