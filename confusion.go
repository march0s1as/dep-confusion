package main
import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/fatih/color"
)

func frescura(){
    fmt.Print("[")
    color.Set(color.FgGreen)
    fmt.Print("OK")
    color.Set(color.FgWhite)
    fmt.Print("] ")
}
func frescura2(){
    fmt.Print("[")
    color.Set(color.FgRed)
    fmt.Print("MAYBE VULNERABLE")
    color.Set(color.FgWhite)
    fmt.Print("] ")
}

func ascii(){
	fmt.Println(`
       ./~
(=#####{>==================-
       °\_

`)
}
func Between(str, starting, ending string) string {
    s := strings.Index(str, starting)
    if s < 0 {
        return ""
    }
    s += len(starting)
    e := strings.Index(str[s:], ending)
    if e < 0 {
        return ""
    }
    return str[s : s+e]
}

func error(err interface{}) {
    if err != nil{
        panic(err)
    }
}

func checar_npm(wordlist string){
	url := fmt.Sprintf("https://www.npmjs.com/search?q=%s", wordlist)
	req, err := http.Get(url)
	error(err)

	body, err := ioutil.ReadAll(req.Body)
	error(err)

	resultado := Between(string(body),`hover-black">`,"</h3")
	if resultado == ""{
		frescura2()
		fmt.Println(wordlist)
	}
}

func checar_pip(wordlist string){
	url := fmt.Sprintf("https://pypi.org/search/?q=%s&o=", wordlist)
	req, err := http.Get(url)
	error(err)

	body, err := ioutil.ReadAll(req.Body)
	error(err)

	resultado := Between(string(body),`package-snippet__name">`,"</span>")
	if resultado == ""{
		frescura2()
		fmt.Println(wordlist)
	}
}

func checar_gem(wordlist string){
	url := fmt.Sprintf("https://rubygems.org/search?query=%s", wordlist)
	req, err := http.Get(url)
	error(err)

	body, err := ioutil.ReadAll(req.Body)
	error(err)

	resultado := Between(string(body),`<h2 class="gems__gem__name">`,`<span class="gems__gem__version">`)
	if resultado == ""{
		frescura2()
		fmt.Println(wordlist)
	}
}

func main(){
	ascii()
	var tipo string
	var wordlist string

	flag.StringVar(&tipo, "t", "", "")
	flag.StringVar(&wordlist, "w", "", "")
	flag.CommandLine.Usage = func() { fmt.Println(" > ./confusion.go -w WORDLIST -t TYPE \n > example: ./confusion.go -w libs.txt -t pip\n > example: ./confusion.go -w libs.txt -t gem\n > example: ./confusion.go -w libs.txt -t npm ") }
	flag.Parse()

	wordlist_vazia, err := os.OpenFile(wordlist, os.O_RDWR, 0000)
	_ = wordlist_vazia
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println(`
 > ./confusion.go -w WORDLIST -t TYPE 
 > example: ./confusion.go -w libs.txt -t pip
 > example: ./confusion.go -w libs.txt -t gem
 > example: ./confusion.go -w libs.txt -t npm`)
		return
	}

	if tipo == "pip"{
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string
		 
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
		file.Close()

		for _, eachline := range txtlines{
			checar_pip(eachline)
		}
	} else if tipo == "npm"{
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string
		 
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
		file.Close()

		for _, eachline := range txtlines{
			checar_npm(eachline)
		}
	} else if tipo == "gem" {

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string
		 
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
		file.Close()

		for _, eachline := range txtlines{
			checar_gem(eachline)
		}
	}
}