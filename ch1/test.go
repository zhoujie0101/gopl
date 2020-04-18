package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
    AbsoluteZeroC Celsius = -273.15 // 绝对零度
	    FreezingC     Celsius = 0       // 结冰点温度
		    BoilingC      Celsius = 100     // 沸水温度
			)
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	var f float64
	fmt.Println(f)
//	fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
}
