# 1. PACOTES 

Quando lidamos com mais de um pacote em go, criamos uma estrutura chamada módulo, que são um conjunto de pacotes que compoem o meu projeto. 
Para a criação de um módulo, utilizamos o seguinte comando:
## go mod init nome_do_modulo - Iniciando um novo módulo e que centraliza todas as dependencias do meu projeto. 

Para rodar uma aplicação no go e também criar um arquivo binário contendo a aplicação, utilizamos o seguinte comando:
## go run nome_do_arquivo.go - Rodar uma aplicação no terminal.
## go build nome_do_arquivo.go - Criar um arquivo binário executável contendo a aplicação. Quando faço a criação dele, ele terá o nome do módulo que criei.

Como o golang não é uma linguagem orientada a objetos, para saber se uma função é pública ou não, usamos identificamos ela pela letra maiúscula ou minúscula.
LETRA MAIÚSCULA - A função é pública e pode ser acessada por outros pacotes. Como uma boa prática, deve ser feito um comentário na função que será exportada.
LETRA MINÚSCULA - A função é privada e não pode ser acessada por outros pacotes. Ademais, consigo acessar funções do mesmo pacote. 

Para eu acessar um pacote externo, utilizo o seguinte comando:
## go get nome_do_pacote.com/nome_do_pacote - Acesso a um pacote externo.
## go mod tidy - Remove todas as dependencias que não estão sendo utilizadas.

# 2. VARIÁVEIS

Para variáveis no golang, que é fortemente tipada, nós conseguimos declará-las em dois tipos que são:
## EXPLICITO - var nome_da_variavel tipo - Que são declaradas com o tipo da variável explicitamente.
## IMPLICITO - nome_da_variavel := valor - Que são declaradas com o tipo da variável implicitamente.

Para fazer a declaração de mais de uma variável de uma vez, colocamos ela entre parenteses. A cada declaração deve ser feito uma quebra de linha.

var (
    var1 tipo
    var2 tipo
    var3 tipo
)

Para fazer por inferencia de tipo fazemos:
# var1, var2, var3 := valor1, valor2, valor3

# TIPOS DE DADOS
Fazendo a listagem de todos os tipos de dados presentes no golang:

## int := int8, int16, int32 (rune), int64 - Númetos inteiros com sinal.
## uint := uint8 (byte), uint16, uint32, uint64 - Números inteiros sem sinal.
## float := float32, float64 - Números reais.
## var erro error = errors.

No golang, o erro é considerado um tipo de dado. Ele pode ser declarado como variável, mas ele vem acompanhado do pacote errors, que é utilizado para atribuir valores nele.

# 3. FUNÇÕES 

Para funções, nós temos uma sintaxe onde a declaração funciona da seguinte forma:
## func nome_da_função (variavel tipo_da_variavel) retorno_tipo_da_variavel {}

Podemos também declarar funções que tem variaveis do mesmo tipo com a seguinte sintaxe:
## func nome_da_função (variavel_1, variavel_2, variavel_3 tipo_da_variavel) retorno_tipo_da_variavel {}

Também posso criar funções dentro do main como um tipo com a seguinte sintaxe:
## var := func(variavel tipo_da_variavel) {}

Podemos fazer mais de um tipo de retorno também 
## func nome_da_função (variavel_1, variavel_2, variavel_3 tipo_da_variavel) (retorno_tipo1_da_variavel, retorno_tipo2_da_variavel) { return variavel1, variavel2}
OBS: Posso utilizar o underline para não usar o retorno daquela variável especifica. 

# 4.STRUCTS

Os structs são coleções tipadas de campos usadas para agrupar variáveis relacionadas sob um único nome, funcionando como a forma primária de representar entidades do mundo real ou registros de dados.
Para a criação de um struct, nós utilizamos a seguinte sintaxe:
## type nome_do_struct struct { nome_do_campo tipo_do_campo }

Quando iremos declarar o nosso struct dentro da função main, podemos escrever de duas formas, que são as seguintes sintaxes:
## var nome_da_variavel nome_do_struct - Dessa forma iremos criar de forma avulsa uma variável do tipo struct.
## nome_da_variavel := nome_do_struct{nome_do_campo: valor_do_campo} - Dessa forma criamos a variável ja passando os valores dos campos dentro das chaves. 

Também podemos herdar as variáveis de uma struct dentro da outra sem precisar criar as mesmas variáveis novamente. Exemplo:
## type usuario struct { nome string altura float32 idade int}
## type estudante struct {usuario curso string faculdade string}

# 5. ARRAY

Arrays são uma lista de dados que podem ser criadas do mesmo tipo. A sintaxe para a criação de uma array é a seguinte:
## var arr[tamanho]tipo 

Podemos também escrever de outra forma:
## arr := [tamanho]tipo{valores} 

Para que eu consiga fixar o tamanho do array de acordo com os elementos que eu passei, utilizo três pontos dentro dos colchetes:
## arr := [...]tipo{valores}
OBS: Ele ainda continua com uma quantidade de posições fixas. 

# 6. SLICE

Slices são estruturas, que mesmo tendo um tipo de dado imutável, ela deixa adicionar quantos elementos eu quiser, desde que seja do mesmo tipo de dado. A sixtaxe para a criação de um slice é a seguinte:
## var slice []tipo 

Podemos também escrever de outra forma:
## slice := []tipo{valores} 

Para adicionar mais um elemento ao final do slice utilizamos:
## slice = append(slice, valor)

Posso também adicionar uma fatia dos arrays dentro de um slice da seguinte forma: 
## novoSlice = arr[inicio:final]

Ademais, consigo alterar o elemento que está na posição que foi passada para o slice.
## slice[indice] = novoValor

Para a criação de arrays internos, nós utilizamos a função make para passar os parâmetros necessarios com a seguinte sintaxe:
## sliceInterno := make([]tipo, tamanho, capacidade)

Para eu ver o tamanho e capacidade de um slice usamos a seguinte sintaxe:
## tamanho = len(slice)
## capacidade = cap(slice)

Caso ele estoure a capacidade do slice, o go dobra essa capacidade para não estourar.

# 7. PONTEIROS

Ponteiros são variáveis que armazenam endereços de memória de outras variáveis, em vez de conterem valores diretamente. Quando um ponteiro contém o endereço de uma variável, diz-se que ele "aponta" para essa variável. Para a criação de um ponteiro utilizamos a seguinte sintaxe:
## var ponteiro *tipo - Criamos uma variável ponteiro que aponta para o uma variável do tipo passado. 

Para acessar o endereço de memória de uma variável do tipo especificado, eu utilizo a seguinte sintaxe:
## ponteiro = &variavel

Para acessar o valor que está presente dentro daquele endereço de memória utilizamos a seguinte sintaxe:
## *ponteiro

# 8. MAPAS

São estruturas do tipo chave valor. Para a criação de um map, utilizamos a seguinte sitaxe:
## var mapa map[tipo_chave]tipo_valor
## mapa := map[tipo_chave]tipo_valor {chave: valor}

Também posso criar um map, onde a chave é um tipo e o valor é outro map, com a seguinte sintaxe:
## mapa := map[tipo_chave]map[tipo_chave]tipo_valor {chave: {chave: valor}}

Adicionando novas chaves e valores no mapa e também deletar :
## nome_map[chave] = valor
## delete(nome_do_mapa, chave)

# 9. CONDICIONAIS

Podemos criar variáveis dentro do próprio if com a estrutura if init.
## if nome_variavel := valor; condição {} - Inicio primeiramente a variável e finalizo com o ponto e virgula, e após isso passo a condicional que desejo. 

# 10. SWITCH CASE

A estrutura switch case é uma estrutura de controle de fluxo utilizada para tomar decisões com base no valor de uma variável ou expressão, sendo uma alternativa mais legível e organizada ao uso de múltiplas declarações. Para a declaração fazemos:

	switch valor {
	case valor:
        Adicionar a operação a ser executada caso caia nesse caso
	default:
		Adicionar quando a operação não encontra nenhum valor que está presente nos casos.
	}

# 11. LOOPS 

As estruturas de laços de repetição em golang só são construídas usando o for. As sintaxes de for que são utilizadas:
## for condição {} - Tem a característica de um while.
## for i:=0; i < tamanho_desejado; i++ {} - For com quantidade definida de laços. 
## for index, data := range elemento {} - Consigo acessar cada dado de cada indice do elemento. 

# 12. DEFER 

O defer é uma palavra-chave usada para adiar a execução de uma função até que a função que a contém retorne. Com o defer na frente, a variável será adiada e será executada somente após a execussão completa da função, ou antes de acontecer o retorno de uma função.
## defer fmt.Println - Essa instrução será adiada e executada após o termino da função ou anteriormente a algum retorno. 

# 13. PANIC RECOVER

panic (O Impacto): É um comando que você usa para interromper o programa imediatamente quando acontece um erro gravíssimo que o código não sabe como resolver (ex: tentar ler um arquivo vital que não existe ou dividir por zero).

recover (O Airbag): É uma função que captura o pânico e impede o programa de quebrar (fazer o crash). Ela avisa o Go: "Pode parar o desespero, eu limpo a bagunça e o programa pode continuar rodando".

A Regra Única
O recover só funciona se estiver dentro de um defer.

Como o panic para tudo na hora, o defer é o único bloco de código que o Go garante que vai rodar "no último segundo" antes da função morrer.