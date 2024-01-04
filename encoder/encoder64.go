package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"
)

func encodeAndSave(inputPath string, outputPath string) (string, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Неверный путь входного файла")
		return "", err
	}
	encodedData := b64.StdEncoding.EncodeToString([]byte(data))
	err = os.WriteFile(outputPath, []byte(encodedData), 0644)
	if err != nil {
		fmt.Println("Не удалось записать файл")
		return "", err
	} else {
		fmt.Println("Файл успешно записан")
	}
	return encodedData, nil
}

func decodeAndSave(inputPath string, outputPath string) (string, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Неверный путь входного файла")
		return "", err
	}
	decodedData, err := b64.StdEncoding.DecodeString(string(data))
	if err != nil {
		fmt.Println("Ошибка при изменении файла")
		return "", err
	}
	err = os.WriteFile(outputPath, decodedData, 0644)
	if err != nil {
		fmt.Println("Не удалось записать файл")
		return "", err
	} else {
		fmt.Println("Файл успешно записан файл")
	}
	return string(decodedData), nil
}

func encodeOrDecode(input string, output string, command string) (string, error) {
	commands := os.Args[1:]

	if command == "" {
		command = commands[0]
	}
	if len(commands) < 2 {
		fmt.Println("Введите путь входного файла")
	}
	if input == "" && output == "" {
		input = commands[1]
		output = ""
	}

	if input == "-i" {
		if len(commands) > 2 {
			input = commands[2]
			if len(commands) > 3 {
				output = commands[3]
				if output == "-o" && len(commands) > 4 {
					output = commands[4]
				} else {
					fmt.Println("Введите путь выходного файла")
				}
			} else {
				output = input + ".out"
			}
		} else {
			fmt.Println("Введите путь входного файла")
		}
	} else {
		output = input + ".out"
	}

	if command == "encode" {
		return encodeAndSave(input, output)
	}

	if command == "decode" {
		return decodeAndSave(input, output)
	}

	if command == "" {
		fmt.Println("Введите команду")
	}

	return "", nil
}

func main() {
	encodeOrDecode("", "", "")
	tests()
}
