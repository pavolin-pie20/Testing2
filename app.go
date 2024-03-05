package main

import "fmt"

// Структура для представления книги
type Book struct {
Title string
Author string
Year int
}

// Метод для вывода информации о книге
func (b Book) DisplayInfo() {
fmt.Printf("Title: %s\nAuthor: %s\nYear: %d\n", b.Title, b.Author, b.Year)
}

// Структура для представления автора
type Author struct {
Name string
Country string
}

// Метод для вывода информации об авторе
func (a Author) DisplayInfo() {
fmt.Printf("Name: %s\nCountry: %s\n", a.Name, a.Country)
}

func main() {
// Создаем экземпляр книги
book := Book{
Title: "1984",
Author: "George Orwell",
Year: 1949,
}

// Создаем экземпляр автора
author := Author{
Name: "George Orwell",
Country: "United Kingdom",
}

// Вызываем методы для отображения информации о книге и авторе
fmt.Println("Book Info:")
book.DisplayInfo()

fmt.Println("\nAuthor Info:")
author.DisplayInfo()
}