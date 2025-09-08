package main

import "fmt"

// Целевой интерфейс, который ожидает клиент
type Charger interface {
	Charge() string
}

// Существующая структура с несовместимым интерфейсом
type USPlug struct{}

func (u *USPlug) PlugInUS() string {
	return "Charging with US plug"
}

// Адаптер, который преобразует USPlug в Charger
type USPlugAdapter struct {
	usPlug *USPlug
}

func (a *USPlugAdapter) Charge() string {
	return a.usPlug.PlugInUS() // Делегируем вызов
}

// Клиентский код, который работает с Charger
func ChargeDevice(charger Charger) {
	fmt.Println(charger.Charge())
}

func main() {
	usPlug := &USPlug{}
	adapter := &USPlugAdapter{usPlug: usPlug}

	ChargeDevice(adapter)
}
