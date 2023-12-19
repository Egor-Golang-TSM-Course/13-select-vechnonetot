package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура Chef (Повар)
type Chef struct {
	Name           string
	CookingTimeSec int
}

// Функция приготовления блюда
func (c *Chef) Cook(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s начинает готовить блюдо.\n", c.Name)
	time.Sleep(time.Duration(c.CookingTimeSec) * time.Second)
	fmt.Printf("%s заканчивает готовить блюдо.\n", c.Name)
}

func main() {
	// Создание поваров
	chefs := []Chef{
		{Name: "Повар 1", CookingTimeSec: 3},
		{Name: "Повар 2", CookingTimeSec: 5},
		{Name: "Повар 3", CookingTimeSec: 2},
	}

	// Запуск горутин для каждого повара
	var wg sync.WaitGroup
	for _, chef := range chefs {
		wg.Add(1)
		go chef.Cook(&wg)
	}

	// Время работы ресторана
	time.Sleep(10 * time.Second)

	// Завершения работы всех поваров
	wg.Wait()

	fmt.Println("Ресторан закрывается.")
}
