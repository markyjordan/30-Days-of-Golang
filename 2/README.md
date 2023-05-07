# Day 2

The goal for Day 2 was to achieve proficiency in Go's fundamental data types and control flow constructs, encompassing loops and function signatures.

## Notes

### Bool
- The bool data type represents boolean values, which can be either `true` or `false`.
- The zero value of bool is `false`.
- The `&&` operator performs logical AND, while the `||` operator performs logical OR. 
- The `!` operator performs logical NOT.
- In Go, boolean values are used extensively for control flow and conditional statements.

### String
- The string data type represents a sequence of Unicode characters.
- Strings are immutable in Go, which means that once a string is created, it cannot be changed.
- String literals are enclosed in double quotes " or backticks `.
- Go provides a rich set of built-in functions for manipulating strings, such as len, concatenation, substring, etc.
- Go also supports Unicode and provides built-in functions for handling Unicode code points and runes.

### Integers
- The int data type represents signed integers with a minimum size of 32 bits.
- The `int8`, `int16`, `int32`, and `int64` data types represent signed integers with widths of 8, 16, 32, and 64 bits respectively.
- The range of values that can be stored in an integer data type depends on its width. For example, an `int8` can store values between -128 and 127, while an `int64` can store values between -9223372036854775808 and 9223372036854775807.

### Rune
- The `rune` data type is an alias for `int32`.
- It represents a Unicode code point.
- In Go, the `rune` data type is often used to represent characters in a `string`.

### Unsigned integers
- The uint data type represents unsigned integers with a minimum size of 32
  bits.
- The `uint8`, `uint16`, `uint32`, and `uint64` data types represent unsigned integers with widths of 8, 16, 32, and 64 bits respectively.
- The range of values that can be stored in an unsigned integer data type
  depends on its width. For example, a `uint8` can store values between 0 and
  255, while a `uint64` can store values between 0 and 18446744073709551615.

### Byte
- The `byte` data type is an alias for `uint8`.
- It represents an 8-bit unsigned integer that can store values between 0 and 255.
- In Go, the `byte` data type is often used to represent individual bytes of data, such as those found in a file or network packet.

### Floats
- The `float32` and `float64` data types represent floating-point numbers with widths of 32 and 64 bits respectively.
- The `float32` data type has a precision of about 7 decimal digits, while the `float64` data type has a precision of about 15 decimal digits.
- Go follows the IEEE-754 standard for floating-point arithmetic, which means that floating-point calculations may not always be exact.

### Complex
- The `complex64` and `complex128` data types represent complex numbers with widths of 64 and 128 bits respectively.
- A complex number is a number that can be expressed in the form a + bi, where a
and b are real numbers and i is the imaginary unit, which is defined as the square root of -1.
- In Go, complex numbers are represented as a pair of floating-point numbers, with the real part represented by a `float32` or `float64`, and the imaginary part represented by a `float32` or `float64`.
- Go provides built-in functions for performing arithmetic operations on complex numbers, such as addition, subtraction, multiplication, and division.
- Like floating-point numbers, complex numbers in Go may not always be exact due to limitations in floating-point arithmetic.

---
Date Completed: 2023-05-07

<br>

[Return to the Top](#day-2)
