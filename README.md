# monitordesites-golang

O projeto tem como objetivo realizar um monitoramento de status de site que verifica se os sites em questão estão online (status: 200) ou offline e registrar os resultados no arquivo "log.txt". Foi utilizada a linguagem Golang e suas respectivas bibliotecas para adaptação de um código estruturado em funções que se encaixam umas nas outras.

OBJETIVO:

- Fazer requisições periódicas em sites definidos pelo usuário.
- Verificar se o status HTTP de retorno é 200 (OK) ou não.
- Registrar o status (online/offline) de cada site com data e hora em um arquivo de log.
- Permitir que o usuário veja o histórico de monitoramento diretamente pelo terminal.

TECNOLOGIA UTILIZADA:

GOLANG (https://golang.org)

ESTRUTURA:

Hello.go - Código fonte principal.
sites.txt - arquivo com a lista de sites a serem lidos e monitorados.
log.txt - Arquivo txt com o resultado do monitoramento dos sites, incluindo a data e a hora exata do monitoramento.
README.md - Documentação do projeto.

COMO FUNCIONA:

- O programa lê os sites listados no arquivo "sites.txt"
- Para cada site, ele faz uma requisição HTTP GET
- Se o site retornar status "200", o status é considerado 'online'. Qualquer outro status é 'offline';
- A função "registraLog" guarda o site e o status em questão no arquivo de texto "log.txt"
- A função "imprimeLog" lê o arquivo.

COMO UTILIZAR:

- Insira um site por linha em cada linha do arquivo "sites.txt"
- no terminal, execute a ordem "go run Hello.go"
- Selecione a opção 1 do menu para monitorar os sites.
- Selecione a opção 2 do menu para mostrar o log atualizado do monitoramento.
- Todos os logs ficam no arquivo "log.txt", você pode alterá-lo como desejar.

AUTOR:

João Pedro Pereira Requejo
