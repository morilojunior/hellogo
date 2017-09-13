package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	//showNames()
	showIntroduction()
	for {
		showMenu()
		command := readInputCommand()

		//O Go só permite valor boolean, se colocar uma variavel com valor inteiro dará erro
		if command == 1 {
			startMonitoring()
		} else if command == 2 {
			showLogs()
		} else if command == 0 {
			fmt.Println("Saindo do Programa")
			os.Exit(0) // Para sair do programa e retornar valor 0 para o sistema operacional
		} else {
			fmt.Println("Não conheço esse command")
			os.Exit(-1)
		}

		//Exemplo de switch case, no Go não existe break no switch, só ira executar um case
		//switch command {
		//case 1:
		//	fmt.Println("Monitoramento")
		//default:
		//	fmt.Println("Deu Ruin")
		//}
	}
}

func showIntroduction() {
	nome := "Morilo"
	idade := 26
	versao := 1.0

	fmt.Println("Olá sr.", nome, "de idade: ", idade)
	fmt.Println("primeira versão do programa ", versao)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readInputCommand() int {

	var command int

	//fmt.Scanf("%d", &command)
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := readSitesInFile()

	//Array
	//var sites [5]string

	//Slice
	//sites := []string{"https://www.caelum.com.br/", "http://www.google.com"}

	// for
	//for i := 0; i < len(sites); i++ {
	//	fmt.Println(sites[i])
	//}

	//for com range
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Minute)
	}
}

//func showNames() {
//Slice
//nomes := []string{"Morilo", "Micaele", "Joyce", "Anny"}
//fmt.Println(len(nomes))
//fmt.Println("O meu slice tem capacidade de", cap(nomes), "itens")

//nomes = append(nomes, "Ivoneide")
//fmt.Println("O meu slice tem", len(nomes), "itens")
////Quando estouramos o tamonho do slice ele dobra o valor exe: antes era 4 agora é 8
//fmt.Println("O meu slice tem capacidade de", cap(nomes), "itens")
//}

func testSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("O Site", site, "foi carregado com sucesso")
		recordLog(site, true)
	} else {
		fmt.Println("O Site", site, "está com problemas. Status code", resp.StatusCode)
		recordLog(site, false)
	}
}

func readSitesInFile() []string {
	var sites []string

	//Lê o arquivo inteiro de uma vez
	//arquivo, err := ioutil.ReadFile("sites.txt")

	//Com o os.Open() ele devolve o ponteiro do arquivo
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	//O bufio.NewReader cria um leitor
	leitor := bufio.NewReader(arquivo)

	for {
		//retorna até onde encontra o delimitador que é "\n"
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	//fecha conexão com arquivo
	arquivo.Close()

	return sites
}

func recordLog(site string, status bool) {

	// Em Go o "|" serve como "ou" nesse caso o arquivo terá a permissão de ler e escrever
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//Escreve no arquivo que criamos e converte o tipo boolean para string
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + site + " - online " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func showLogs() {
	fmt.Println("Exibindo Logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
