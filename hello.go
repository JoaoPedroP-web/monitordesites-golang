package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	
	for {
		exibeMenu()

		comando := comandoLido()

		switch comando {
			case 1:
				iniciarMonitoramento()
			case 2:
				fmt.Println("Exibindo logs...")
				imprimeLog()
			case 0:
				fmt.Println("Saindo do programa...")
				os.Exit(0)
			default:
				fmt.Println("Opção indisponível.")
				os.Exit(-1)
			}
	}
}

func exibeIntroducao() {
	nome := "João"

	fmt.Println("Bem-vindo,", nome);
}

func exibeMenu() {
	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func comandoLido() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	fmt.Println("")

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leitorDeArquivoSites()

	for i := 0 ; i < monitoramentos ; i++ {
		for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testarSite(site)
		}

		time.Sleep(delay * time.Second)

		fmt.Println("")
	}

	fmt.Println("")
}

func testarSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "está ativo. Status Code:", resp.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leitorDeArquivoSites() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666);

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02-01-2006 15:04:05") + " " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLog() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}