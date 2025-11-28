package main

import (
	"fmt"
	"strconv"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// لو مش رقم صالح كـ float → رجّعه زي ما هو
	if _, err := strconv.ParseFloat(dec, 64); err != nil {
		return dec + "\n"
	}

	// دور على النقطة
	dot := -1
	for i := range dec {
		if dec[i] == '.' {
			dot = i
			break
		}
	}

	// مفيش نقطة → رجّع الأصل
	if dot == -1 {
		return dec + "\n"
	}

	// لو كل اللي بعد النقطة أصفار → رجّع الأصل
	zeros := true
	for i := dot + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			zeros = false
			break
		}
	}
	if zeros {
		return dec + "\n"
	}

	// شيل النقطة
	clean := dec[:dot] + dec[dot+1:]

	// ✨ هنا نستخدم فكرتك: نحوله لـ int علشان نشيل الأصفار اللي في البداية
	n, err := strconv.Atoi(clean)
	if err != nil {
		// احتياطاً: لو حصل error (overflow مثلاً) رجّع clean زي ما هو
		return clean + "\n"
	}

	// رجّعه string تاني
	return strconv.Itoa(n) + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
