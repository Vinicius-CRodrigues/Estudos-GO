# Funcionamento do Protocolo HTTP

O **HTTP** (*Hypertext Transfer Protocol*) funciona estritamente no modelo **Requisição-Resposta (Request(pergunta)-Response(resposta))**. O cliente (navegador, aplicativo, script) faz o pedido, o servidor processa e devolve o resultado.

Mas para entender a engenharia por trás dessa comunicação, é preciso olhar para a anatomia dessas mensagens e o comportamento do protocolo na rede.

## A Natureza do Protocolo

*   **É "Stateless" (Sem estado):** O protocolo HTTP por si só não tem memória. O servidor trata cada requisição como uma transação totalmente nova e independente, sem lembrar da requisição anterior. É por isso que, para manter você logado em um sistema, a aplicação precisa usar mecanismos extras (como *Cookies* ou *Tokens JWT*) que são enviados junto com todas as requisições para provar quem você é.
*   **Depende da Camada de Transporte:** O HTTP é um protocolo de aplicação (Camada 7 do modelo OSI). Ele não transporta os pacotes fisicamente. Para isso, ele geralmente roda sobre o protocolo TCP (na porta 80). Antes do cliente enviar a requisição HTTP de fato, ocorre o *Handshake* do TCP para garantir que a conexão entre cliente e servidor foi estabelecida de forma confiável.

---

## A Anatomia da Comunicação

A conversa entre cliente e servidor é feita basicamente através de textos estruturados.

### 1. A Requisição (O que o cliente manda)

Quando você acessa um site ou consome uma API, a requisição não é apenas um "me envie os dados". Ela é dividida em:

*   **Método (Verbo):** Define a intenção da ação.
    *   `GET`: Pede para buscar um dado (ex: carregar uma página web).
    *   `POST`: Envia um dado novo para o servidor (ex: enviar um formulário de cadastro).
    *   `PUT` / `PATCH`: Atualiza um dado existente.
    *   `DELETE`: Apaga um dado.
*   **Caminho (URI):** O recurso específico que está sendo acessado (ex: `/api/usuarios/123`).
*   **Headers (Cabeçalhos):** Metadados cruciais da requisição. Eles dizem ao servidor qual navegador está sendo usado (`User-Agent`), qual formato de resposta o cliente entende (ex: `Accept: application/json`), etc.
*   **Body (Corpo):** Opcional. É a "carga útil" (*payload*) da requisição, usada no `POST` ou `PUT` para enviar os dados reais, como um arquivo JSON com as suas credenciais de login.

### 2. A Resposta (O que o servidor devolve)

Após o processamento (que pode envolver consultar um banco de dados ou acionar outras APIs internas), o servidor devolve:

*   **Status Code:** Um número de 3 dígitos que resume o resultado da transação.
    *   **2xx (Sucesso):** Tudo certo (ex: `200 OK`).
    *   **3xx (Redirecionamento):** O recurso mudou de lugar (ex: `301 Moved Permanently`).
    *   **4xx (Erro do Cliente):** Você pediu algo errado ou não tem permissão (ex: `404 Not Found` ou `403 Forbidden`).
    *   **5xx (Erro do Servidor):** O cliente fez tudo certo, mas a aplicação do servidor falhou (ex: `500 Internal Server Error`).
*   **Headers:** Metadados da resposta, como o tamanho do arquivo (`Content-Length`) e o tipo de dado que está sendo devolvido (`Content-Type: text/html`).
*   **Body:** O conteúdo em si que foi solicitado (o código HTML da página, um arquivo de imagem, ou um JSON com os dados).

# Execução do projeto

Para iniciar o servidor, rode o comando abaixo no seu terminal, garantindo que você está dentro do diretório do projeto:

```bash
go run nome_do_seu_projeto.go
```

> **OBSERVAÇÕES IMPORTANTES:**
> * **Terminal travado:** Ao rodar o comando, o terminal ficará "travado". Isso é normal e indica que o servidor está no ar, aguardando acessos.
> * **Como parar a aplicação:** Pressione `Ctrl + C` no terminal para encerrar a execução.
> * **Erro fatal (Porta em uso):** Se o programa retornar um erro "fatal" ao iniciar, a porta definida no código já está sendo utilizada por outra aplicação.

## Como acessar a página

Abra o seu navegador e digite o seguinte endereço na barra de navegação:

**localhost:numero_da_porta/home**

Ao acessar essa URL, a página HTML passada como template será carregada com sucesso.