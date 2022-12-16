package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

type Note struct {
  Name string
  Surname string
  NoteText string
}

var NoteStorage = []Note{}
var reader = bufio.NewReader(os.Stdin)

func main() {
  CreateNewNote()
  for {
    fmt.Println("что делать дальше? \n c - создать новую заметку \t l - вывести все записи \t q - завершить исполнение")
    Command := ReadLine("Введите команду: ")
    switch Command {
    case "c":
      CreateNewNote()
      break
    case "l":
      ListAllNotes()
      break
    case "q":
      os.Exit(0)
    default:
      fmt.Println("Команда не найдена. Попробуйте ещё раз")
    }
  }
}

func ReadLine(helpText string) string {
  fmt.Print(helpText)
  input, err := reader.ReadString('\n')

  if err != nil {
    log.Fatal(err)
  }

  return strings.TrimSuffix(input, "\n")
}

func CreateNewNote() (name, surname, note string) {
  NewNote := Note{}
  NewNote.Name = ReadLine("Введите имя: ")
  NewNote.Surname = ReadLine("Введите фамилию: ")
  NewNote.NoteText = ReadLine("Введите заметку: ")

  NoteStorage = append(NoteStorage, NewNote)
  fmt.Printf("Введённые данные: \n  имя: %s\n  фамилия: %s\n  заметка: %s\n", NewNote.Name, NewNote.Surname, NewNote.NoteText)
  return name, surname, note
}

func ListAllNotes() {
  for index, note := range NoteStorage {
    fmt.Printf("\nЗаписка №%d\n  имя: %s\n  фамилия: %s\n  заметка: %s\n", index, note.Name, note.Surname, note.NoteText)
  }
}