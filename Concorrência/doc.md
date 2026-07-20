# CONCORRENCIA
Programação concorrente é um paradigma de desenvolvimento que permite a execução simultânea ou intercalada de múltiplas tarefas dentro de um programa, otimizando o uso de recursos computacionais e melhorando a eficiência do sistema. Diferente da execução sequencial tradicional, onde as tarefas ocorrem uma após a outra, a concorrência permite que o software gerencie várias operações ao mesmo tempo, seja através de múltiplos núcleos de processamento ou através do alternamento rápido de tarefas em um único núcleo.

# 1. GoRoutines
As Go routines permitem que funções ou métodos sejam executados simultaneamente com o restante do programa, mas com um custo de memória e processamento drasticamente menor do que as threads tradicionais do sistema operacional. A sintaxe para a criação de uma go routine é a seguinte:
## go execussao_da_rotina 

# 1.2 WAIT GROUP
Crio uma forma de sincronizar duas goRoutines e fazer com que o programa finalize somente quando as suas rotinas forem finalizadas. 
## var waitGroup sync.WaitGroup - Crio uma variável que representa um grupo de espera.
## waitGroup.Add(nº de goRoutines desejados) - Adiciono uma quantidade de goRoutines que eu desejo que aconteça. 
## waitGroup.done() - Retiro -1 do contador de rotinas na fila de espera.
## waitGroup.wait() - Espero a contagem de todas as goRoutines chegarem em zero.

# 1.3 CHANNELS
Canais de envio e recebimento de dados dentro de uma goRoutine.
## canal := make(chan tipo_de_dados) - Crio um canal com a palavra chave chan, e digo que ele só poderá trafegar dados com o tipo de dado especifico. 
## canal -> dado - Enviando as informações para dentro do canal.
## info, aberto := <-canal - Adiciono dentro de info a informação que está sendo recebida pelo canal, e aberto é uma variável booleana que fiz se o canal está aberto(true) ou se o canal está fechado(false) 
## close(canal) - Aqui fecho a execussão do canal.

O select em Go é uma estrutura de controle utilizada exclusivamente para operções de canais (channels), permitindo que uma goroutine aguarde múltiplas operações de comunicação simultaneamente.  Sua sintaxe assemelha-se à do switch, mas cada case deve representar uma operação de envio ou recebimento em um canal. 

O comportamento fundamental inclui:

Bloqueio: O select bloqueia a execução até que pelo menos um dos casos possa ser executado (ou seja, o canal esteja pronto para leitura ou escrita). 
Aleatoriedade: Se múltiplos casos estiverem prontos ao mesmo tempo, o Go escolhe aleatoriamente qual executar para garantir justiça e evitar starvation. 
Caso Default: A inclusão de um default torna o select não-bloqueante; se nenhum canal estiver pronto, o default é executado imediatamente, evitando que a goroutine fique presa indefinidamente.

# PADRÕES DE CONCORRENCIA

Esse padrão serve para limitar o uso de recursos. Se você tem 10.000 tarefas para fazer (como processar imagens ou fazer requisições HTTP), criar 10.000 goroutines de uma vez pode estourar a memória ou o limite de conexões do servidor.

Com um Worker Pool, você cria um número fixo de goroutines (os "trabalhadores") que ficam escutando um canal de tarefas. Eles vão pegando o trabalho conforme terminam o anterior.

O padrão Generator é uma função que retorna um canal de leitura. Quem chama a função não precisa se preocupar em como os dados estão sendo produzidos ou gerados; apenas consome o canal que recebeu. É excelente para criar fluxos contínuos de dados (streams) ou sequências.

O grande trunfo aqui é o encapsulamento: a goroutine que gera os dados nasce e morre dentro da função geradora.

Multiplexar (ou fazer um "Mux") significa juntar vários canais de entrada em um único canal de saída.

Imagine que você tem dois geradores independentes (um buscando dados do banco de dados e outro de uma API externa). Você quer centralizar todas as respostas em um único lugar para processá-las à medida que forem chegando, sem que uma trave a outra. Usamos a instrução select dentro de uma goroutine para gerenciar isso.