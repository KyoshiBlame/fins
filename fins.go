package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Purchase struct {
	ID     int    `json:"ID"`
	Text   string `json:"Text"`
	Amount int    `json:"Amount"`
}

var (
	Cost     []Purchase
	dataFins = "dataFins.json"
	//Sus = make(map[string][]Purchase)
)

func main() {
	loadFins()

	if len(os.Args) < 2 {
		fmt.Println("Использование Fins")
		fmt.Println("Возможные команды Fins: add|list|delete")
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Использование: Fins add \"покупка\" стоимость")
			return
		}
		PurchText := strings.Join(os.Args[2:len(os.Args)-1], " ")
		var price int
		fmt.Sscanf(os.Args[len(os.Args)-1], "%d", &price)
		addFins(PurchText, price)
	case "list":
		list()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Использование: Fins delete [id]")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		deleteFins(id)
	default:
		fmt.Println("Неизвестная команда:", command)
		fmt.Println("Доступные команды: add, list, delete")
	}
}

func loadFins() {
	//читаю файл чтобы с ним работать
	file, err := os.ReadFile(dataFins)
	if err != nil {
		panic(err)
	}
	//проверка на то что в нём что то есть
	if len(file) == 0 {
		Cost = []Purchase{}
		return
	}
	//загрузка содержимого из файла в массив
	err = json.Unmarshal(file, &Cost)
	if err != nil {
		fmt.Println("Ошибка при расшифровке файла", err)
		Cost = []Purchase{}
	}
}

func saveFins() {

	for cnt := range Cost {
		Cost[cnt].ID = cnt + 1
	}

	data, err := json.MarshalIndent(Cost, "", " ")
	if err != nil {
		fmt.Println("Ошибка при маршлинге", err)
		os.Exit(1)
	}
	err = os.WriteFile(dataFins, data, 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранение файла", err)
		os.Exit(1)
	}
}

func addFins(text string, price int) {
	newPurch := Purchase{
		Amount: price,
		Text:   text,
	}
	Cost = append(Cost, newPurch)
	saveFins()
	fmt.Println("Новая покупка успешно добавлена😊")
}

func deleteFins(id int) {
	for cnt := range Cost {
		if cnt == id {
			Cost = append(Cost[:cnt], Cost[cnt+1:]...)
			saveFins()
			fmt.Println("Покупка удалена успешно✅")
			return
		}
	}
	fmt.Printf("Задача %d не найдена", id)
}

func list() {
	if len(Cost) == 0 {
		fmt.Println("Покупок пока что нету...")
		return
	}
	total := 0

	for _, pr := range Cost {
		total += pr.Amount
	}
	fmt.Println("Лист покупок: ")
	for _, pr := range Cost {
		fmt.Printf("[%d] | %s | price: $%d\n", pr.ID, pr.Text, pr.Amount)
	}
	fmt.Printf("Общая сумма составляет: $%d\n", total)
}
