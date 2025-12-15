package main

import (
	"fmt"
)

// клиент (ф-я) ожидает NowTime
func Vospitanie(n NowTime) {
	n.Otrugaly()
}

// что ожидает клиент
type NowTime interface {
	Otrugaly()
}

// что-то, что клиент не ожидает
// но мы ходим ему это передать
type OldClass struct {
	rozgi int
}

func (o *OldClass) Otlupily() {
	o.rozgi--
	fmt.Printf("Хлестанули, розг осталось %d\n", o.rozgi)
}

// Адапер, который подменит то, что ожидает клиент
type SanctionsAdapter struct {
	oldClass *OldClass
}

func (s *SanctionsAdapter) Otrugaly() {
	s.oldClass.Otlupily()
}

func main() {
	old := &OldClass{rozgi: 10}
	s := &SanctionsAdapter{oldClass: old}

	fmt.Println("Send him to Dagestan for 2–3 years and forget")
	Vospitanie(s)
	Vospitanie(s)
	Vospitanie(s)
}
