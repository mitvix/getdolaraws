package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"unicode"
)

func main() {

	// controla args[...] para buscar o arquivo pdf na linha de comando	
	fileInvoice := flag.String("invoice", "", "Exemplo ./getdolaraws -invoice Invoice.pdf")
	flag.Parse()
	filePath := *fileInvoice
	file, err := filepath.Glob(CaseInsensitive(&filePath))
	if err != nil || file == nil {
		flag.Usage()
	}

	// sudo pacman -Syu popper-utils (pdftotext --help)
	args := []string{
		"-layout",  // Maintain (as best as possible) the original physical layout of the text.
		"-nopgbrk", // Don't insert page breaks (form feed characters) between pages.
		filePath,   // The input file.
		"-",        // Send the output to stdout.
	}
	// executa o comando de terminal (somente *nix, sorry)
	cmd := exec.CommandContext(context.Background(), "pdftotext", args...)

	var buf bytes.Buffer
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		fmt.Println("\nNão esqueça de instalar o popper-utils (sudo apt install popper-utils), se der atualizo para Std lib :-)")
		fmt.Println("Mais informações em https://github.com/mitvix/getdolaraws")
		fmt.Println(err)
		return
	}

	// extrai do buffer via regex o valor do dólar
	getDolar(buf.String())
}

func getDolar(text string) {

	// Expressão regular para encontrar o número float antes de "BRL"
	re := regexp.MustCompile(`\b(\d+\.\d+)\s*BRL`)

	// Encontrar a primeira correspondência
	match := re.FindStringSubmatch(text)

	if len(match) > 1 {
		usd, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			fmt.Println("Erro ao converter para float:", err)
			return
		}
		// Printa na tela o valor do dólar encontrado
		fmt.Println(usd)
	} else {
		fmt.Println("Não foi possível encontrar o valor do Dólar em Reais")
	}
}

// case insensitive to file name (named return)
func CaseInsensitive(path *string) (p string) {
	if runtime.GOOS == "windows" {
		return *path
	}

	for _, r := range *path {
		if unicode.IsLetter(r) {
			p += fmt.Sprintf("[%c%c]", unicode.ToLower(r), unicode.ToUpper(r))
		} else {
			p += string(r)
		}
	}
	return
}
