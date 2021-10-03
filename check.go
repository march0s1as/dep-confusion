package main
import(
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/fatih/color"
)

func ascii() {
	fmt.Println(`
==.   .== 
'--'o'=='
   (|)
    8
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

func frescura(){
    fmt.Print("[")
    color.Set(color.FgGreen)
    fmt.Print("+")
    color.Set(color.FgWhite)
    fmt.Print("] ")
}

func error(err interface{}) {
    if err != nil{
        panic(err)
    }
}

func confusion(url string){

	var arquivo [4]string
	arquivo[0] = "package.json"
	arquivo[1] = "__init__.py"

	for i := 0; i <= 1; i++{

		teste := arquivo[i]
		site_final := fmt.Sprintf(url + "%s", teste)

		req, err := http.Get(site_final)
		error(err)

		if req.Status == "404 Not Found"{
			continue
		} else if req.Status == "404 File not found"{
			continue
		}

		frescura()
		fmt.Print(url)
		color.Set(color.FgRed)
		fmt.Println(teste)
		color.Set(color.FgWhite)

		body, err := ioutil.ReadAll(req.Body)
		_ = body
		error(err)

		f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		error(err)

		defer f.Close()
		if _, err := f.WriteString("\n" + site_final); err != nil {
  			fmt.Println(err)
  		}
	}
}

func tratar_url(eachline string){

	string := strings.HasSuffix(eachline, "/") 
	if string == false {
		url := (eachline + "/")
		confusion(url)
	} else {
		confusion(eachline)
	}
}

func main(){
	ascii()
	var wordlist string

	flag.StringVar(&wordlist, "w", "", "")
	flag.CommandLine.Usage = func() { fmt.Println("\n./dpcf -w sites.txt'") }
	flag.Parse()

	wordlist_vazia, err := os.OpenFile(wordlist, os.O_RDWR, 0000)
	_ = wordlist_vazia
	file, err := os.Open(wordlist)
	if err != nil {
		color.Red("plz specify an wordlist. ")
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	for _, eachline := range txtlines{
		tratar_url(eachline)
	}
}