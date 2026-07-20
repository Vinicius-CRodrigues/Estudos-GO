# Manipulação de JSON e Leitura de Arquivos em Go (Golang)

Este documento descreve os principais conceitos e recursos da linguagem Go utilizados para construir a aplicação de manipulação de JSON e leitura de arquivos.

---

## 1. Mapeamento de Dados com Struct Tags
* **Campos Exportados:** Em Go, campos de structs que começam com letra maiúscula (ex: `Nome`, `Email`) são públicos/exportados.
* **Struct Tags (`json:"chave"`):** Permitem definir meta-informações nos campos de uma struct. No contexto do pacote `encoding/json`, determinam como o nome da propriedade deve aparecer no formato JSON (ex: converter `Nome` em Go para `"nome"` em minúsculo no JSON).

---

## 2. Serialização de Dados (`json.Marshal`)
* **Conversão em Bytes:** O processo de conversão de dados em memória (como `structs` ou `maps`) para uma representação em texto formatada em JSON.
* **Slice de Bytes (`[]byte`):** A função `json.Marshal` devolve um slice de bytes. Para visualizá-lo como texto legível no terminal, faz-se a conversão explícita usando `string(dadosJSON)`.
* **Serialização de Mapas:** Além de structs, o Go permite converter tipos dinâmicos como mapas (`map[string]string`) diretamente para JSON.

---

## 3. Entradas e Saídas do Sistema Operacional (`os.ReadFile`)
* **Leitura de Arquivos:** Utilização da biblioteca nativa `os` para interagir com o sistema de arquivos.
* **Carregamento Direto:** O método `os.ReadFile` lê todo o conteúdo de um arquivo em disco de uma só vez e o retorna diretamente como um slice de bytes (`[]byte`), eliminando a necessidade de gerenciar ponteiros de arquivo manualmente.

---

## 4. Desserialização de Dados (`json.Unmarshal`)
* **Parsing de JSON:** O processo inverso do `Marshal`, responsável por ler o texto em JSON e popular estruturas de dados em memória do Go.
* **Ponteiros (`&`):** A função `json.Unmarshal` exige um ponteiro para a variável de destino (ex: `&dadosRaca`) para que ela possa alterar o valor da variável original no escopo da função chamadora.
* **Compatibilidade de Tipos:** A estrutura de destino precisa ser declarada como um slice (`[]Struct`) caso o arquivo JSON comece como um array (`[...]`).

---

## 5. Iteração e Controle de Fluxo (`for range`)
* **Laços de Repetição com `range`:** Sintaxe utilizada para iterar sobre coleções de dados em Go (slices, arrays, mapas).
* **Descarte de Índices (`_`):** Uso do identificador em branco (`_`) para ignorar o índice da iteração quando apenas os valores dos elementos são necessários.
* **Laços Aninhados:** Utilização de estruturas de repetição internas para percorrer slices contidos dentro dos campos de cada item de uma struct principal (ex: lista de habilidades dentro de cada raça).

---

## 6. Tratamento de Erros em Go
* **Retorno Explícito de Erros:** As funções padrão do Go retornam a interface `error` como último valor de retorno.
* **Encerramento Controlado (`log.Fatal`):** Uso da função do pacote `log` para imprimir a mensagem de erro no terminal e encerrar a execução do programa de forma segura caso ocorra algum problema (como falha na leitura do arquivo ou sintaxe inválida no JSON).