package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http" // Pacote para criação de ferramentas para criação de servidores e clientes web.
)

type Atacante struct {
	Nome    string
	Posicao string
}

// Contém todos os templates de arquivos que temos na nossa aplicação.
var templates *template.Template

// Função handler: Função que define o que irá acontecer quando o usuário acessa a página.

func home(w http.ResponseWriter, r *http.Request) {
	// REQUISIÇÃO: O 'r' representa o request- Tudo o que vem do usuário para o servidor.
	// RESPOSTA ESCRITA: O 'w' é o Writter ou escritor - Tudo o que vai do servidor de volta para o usuário.

	rei := Atacante{"Reinaldo", "atacante"}

	// Pego o index.html e com o 'w' já escrevo direto no responseWritter.
	// Passando os dados da variável 'rei' para a página.
	templates.ExecuteTemplate(w, "index.html", rei)
}

func main() {

	// Com o Must, se houver algum erro na leituta dos arquivos html, ele gera um panic, evitando com que o servidor suba com páginas quebradas.
	// o template.parseGlob, busca os arquivos dentro do diretório atual que tem a extensão html.
	templates = template.Must(template.ParseGlob("*.html"))

	// Router: quando o usuário acessa o /home, o servidor consulta a lista de regras que o HandleFunc criou e aciona a função home.
	http.HandleFunc("/home", home)

	fmt.Println("Executando na porta 5000")

	// O ListenAndServe faz a criação do servidor e abre a porta 5000 para receber as requisições.
	// O nil confirma que qualquer requisição que chegar nessa porta deve ser entregue ao roteador padrão (as rotas estão definidas no HandleFunc).
	log.Fatal(http.ListenAndServe(":5000", nil))
}
