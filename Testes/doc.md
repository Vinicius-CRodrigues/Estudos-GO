# TESTES UNITÁRIOS 

Testes unitários em Go são scripts automatizados que verificam se unidades individuais de código, como funções ou métodos, funcionam conforme o esperado e isoladamente de dependências externas. Eles são fundamentais para garantir a qualidade do software, prevenir regressões e detectar erros precocemente, sendo a espinha dorsal da garantia de qualidade na linguagem. 

Em Go, a prática é facilitada pelo pacote nativo **testing**. A convenção exige que os testes residam em arquivos com o sufixo _test.go e sigam a nomenclatura TestNomeDaFunc, onde NomeDaFunc é o nome da função sendo testada. A execução é simples, utilizando o comando go test, que descobre e roda automaticamente os testes definidos.

Para criarmos a função de teste utilizamos a seguinte sintaxe:
## func TestNomeDaFunc(t *testing.T) {}

# COMANDOS COM O GO TEST

## go test ./... - Ele procura funções de teste em todos os pacotes do seu projeto e roda os testes que foram criados. 
## go test -v - Demonstra com mais detalhes quais funções rodaram testes e quais passaram ou não nos testes. 
Ao contrário da concorrência, que alterna rapidamente entre tarefas para criar a ilusão de simultaneidade em um único núcleo, o paralelismo exige hardware com capacidade de processamento independente para que as tarefas ocorram efetivamente ao mesmo tempo, sem a necessidade de ceder a vez.
## t.parallel() - Rodo as funções de teste em paralelo. 
## go test --cover - executa os testes unitários do pacote Go e calcula a cobertura de código, indicando a porcentagem de instruções do código-fonte que foram executadas durante os testes.
## go test --coverprofile nome_arquivo.txt - Cria um relatório demonstrando as coberturas de teste. 
## go tool cover --func=nome_arquivo.txt - Ele vai ler o arquivo que foi gerado, e irá detalhar melhor via terminal o que foi coberto pela ferramenta de teste. 
## go tool cover --html=nome_arquivo.txt - Ele informa através de uma página html, o que está sendo cobrido pela função de testes e o que não está. 