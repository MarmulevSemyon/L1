package main

import (
	"fmt"
	"unsafe"
)

var justString string // "", len = 0, ptr = 0x0, *justString = 0xe4326(условно) — строка объявлена, но не инициализирована

func someFunc() {
	v := string(make([]byte, 1024)) // "", len = 1024, prt = 0xdf326 (условно)
	fmt.Printf("\tv in func\t|%s\t|%d\t|%p\t|%p\n", v, len(v), &v, (unsafe.StringData(v)))
	// v сразу инициализировали, поэтому её поля:
	// len = 1024, prt = 0xdf326 (условно), ""

	justString = v[:100] // "", len = 100, prt = 0xdf326
	fmt.Printf("justString in func\t|%s\t|%d\t|%p\t|%p\n", justString, len(justString), &justString, (unsafe.StringData(justString)))
	// justString теперь ссылается на те же данные, что и v
	// len = 100, ptr = тот же ptr, что и у v
	// не смотря на то, что сама переменная v "локальная" — её буфер не освобождается

	// корректно можно создавать так:
	justStringGood := string([]byte(v[:100]))
	fmt.Printf("justStringGood in func\t|%s\t|%d\t|%p\t|%p\n", justStringGood, len(justStringGood), &justStringGood, (unsafe.StringData(justStringGood)))

}

func main() {
	// string — структура из:
	// * длины массива байт (len)
	// * указателя на этот массив (ptr *byte)

	// justString только объявили:
	// len = 0, ptr = 0x0
	fmt.Printf("justString in main\t|%s\t|%d\t|%p\t|%p\n", justString, len(justString), &justString, (unsafe.StringData(justString)))
	someFunc()
	// После вызова someFunc():
	// justString содержит первые 100 байт строки v
	// len = 100, ptr = тот же самый буфер, который был создан внутри someFunc()
	fmt.Printf("justString in main\t|%s\t|%d\t|%p\t|%p\n", justString, len(justString), &justString, (unsafe.StringData(justString)))
	// Хотя v выходит из области видимости, её буфер не освобождается сборщиком мусора,
	// потому что justString продолжает ссылаться на этот буфер.
	// полчучается, что маленькая подстрока (justString) удерживает в памяти весь большой массив v.

}
