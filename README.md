# prova

Olá pessoal, a melhor forma de conhecer um desenvolvedor é seguir o código, então, este
é o [meu git](https://github.com/helmutkemper) e está aí para ser avaliado.

[Essa é a minha contribuição](https://github.com/helmutkemper/golang.solid.kiss.complexity.measure/wiki)
a comunidade golang e fala sobre qualidade de código e dicas.

Todos os módulos deveriam ser publicados on-line para ficar dentro do padrão Golang, e 
usar módulos, mas, achei mais fácil deixar o código em um único arquivo compactado para a 
entrega.

Como solicitado, não há frameworks, apenas códigos nativos do próprio Golang.

## Como executar o código

```shell
$ mysolution input-routes.csv
```

Faz o código testar pelo arquivo **./input-routes.csv** e se não encontrar, procurar na
árvore de diretórios. 

> A pasta **src** tem que ser adicionada como libraries.

```shell
GOROOT=__PATH_TO_GO__\go1.14.3 #gosetup
GOPATH=__PATH_TO_PROJECT__\prova #gosetup
__PATH_TO_GO__\go1.14.3\bin\go[.exe] build -o __PATH_TO_PROJECT__\prova\main\main.[bin|exe] __PATH_TO_PROJECT__\prova\main\main.go #gosetup
__PATH_TO_PROJECT__\prova\bin\main.[bin|exe] input-file.correct.csv #gosetup
```

Lembrando, que por uma questão de segurança, a árvore de diretórios só é vista do 
programa para dentro, nunca para fora.

ou 

```shell
$ mysolution
```

Faz o código perguntar pelo arquivo para começar.

**Obs:** Não fiz o texto do shell exatamente como o de vocês

## Estrutura de arquivos

**/main/main.go**: arquivo executável.

**/src**: pasta de bibliotecas;

**/src/commonData**:

O tipo comum facilita a portabilidade e a mudança do código, tipo, preço pode ser
incrementado para outros tipos de moedas. Por exemplo, todo o sistema pode passar a
trabalhar com float point simplesmente mudando Price de int para float64
Nâo achei necessário definir origin e destination como um tipo a parte apenas para uma
demonstração.

**/src/dataSourceInterface**:

a mãe da portabilidade do código orientado a objetos

a interface de dados permite que qualquer banco de dados possa ser usado de forma
fácil apenas construindo pequenas funções respeitando as regras de negócios

o grande ponto é, pequenas interfaces para grupos de funções é melhor que uma
interface gigante

**/src/factoryFlightStretch**:

**factory** ajuda a inicializar objetos complexos e serve de exemplo de código.
Também é prático pelo autocompletar do editor, pois, basta começar a escrever **factory**
e deixar a magia acontecer.

No caso desse código em particular, não sei se controle do channel deveria ter ficado lá.
Na verdade, esqueci a funcionalidade de manter o CSV atualizado e fiz de última hora.

**/src/FlightSuggestion**:

FlightSuggestion é um pacote feito para sugerir rotas de voo baratas

Pessoal, Na minha antiga empresa, e na atual, há um regra de colocar panic() nos erros
graves, assim, o servidor cai e um alarme é gerado para a equipe.
Minha base é programação de microcontroladores (80C51) em C e Assembley, por isto,
tenho consciência da necessidade de poupar memória e usar ponteiro.
Também tenho consciência de que o código pode ser colocado para rodar em um container
com pouca memória e o carregamento de uma lista de trechos de voo muito grandes pode
consumir muita memória e fazer o container travar, dependendo do limite de memória
para o mesmo.

Todos os códigos são formados por funções pequenas de responsabilidade única, para que
o código fique mais fácil de ser entendido e seguido. Nesse ponto, sigo os
trabalhos do Uncle Bob e do Thomas J. McCabe. Também sigo a filosofia do Kiss - "Keep
It Stupid Simple / Keep It Simple, Stupid!" e parte do meu trabalho pode ser
acompanhado em
https://github.com/helmutkemper/golang.solid.kiss.complexity.measure/wiki

Quanto a função flightSuggestion.FindCheapestRoute() ela ficou bem dividia e fácil de
ser seguida e entendida. A ideia básica é sempre usar nomes de métodos explicativos,
já que programadores cansados tendem a esquecer os cometários a medida que o código
muda.
Nesse ponto, detesto comentários tipo, FindCheapestRoute(), procura a rota mais
barata, embora, nem sempre tenha muito o que fazer.
Todo código deve ter nos comentários não o que o código faz, mas, porque ele faz o que
faz. Já tive de refatorar um código com constantes dentro dele que nem a equipe
original lembrava o motivo, por isto, eu sempre comento o motivo quando relevante, e
se o motivo for uma lei, minha antiga empresa era especializada em controle de ponto,
eu sempre coloco no comentário a parte relevante da lei e o link, para que isto não
se perca no tempo.

**/src/importFlightStretch**:

Este pacote importa arquivos CSV com trechos de voos e valores no seguinte formato:
  IATA code da origem;
  IATA cede do destino;
  Preço;

Exemplo:
  GRU,BRC,10
  BRC,SCL,5
  GRU,CDG,75

**/src/restServer**:

Como solicitado, foi feita uma interface restful seguindo as recomendações do site 
[**jsonapi.org**](https://jsonapi.org/) e na minha experiência prévia com desenvolvedores 
mobile.

Chaves:
* Error: array de string;
* Success: booleano, true para sucesso;
* Length: quantidade de objetos retornados;
* Object: array de objetos

Regra: todos os tipos devem ser preservados e quando uma chave não existir, ela deve ter
o seu tipo mantido com o valor padrão para **'not set'**.

Os testes, fiz só o básico para mostrar que sei fazer.
Não havia solicitação de configuração de porta, estão, está fixo na 8080.

Procurar a melhor rota:

Entrada correta de dados:
```url
[GET] http://localhost:8080/api/find?org=BRC&des=CDG
```

Resposta json:
```json
{
    "Error": [],
    "Success": true,
    "Length": 1,
    "Object": [
        {
            "Price": 30,
            "Route": "BRC - SCL - ORL - CDG"
        }
    ]
}
```

Entrada incorreta de dados:
```get
http://localhost:8080/api/find?org=BRC&des=CD
```

Resposta json:
```json
{
    "Error": [
        "destination code must be three letters long"
    ],
    "Success": false,
    "Length": 0,
    "Object": []
}
```

Adição de nova rora:

Entrada correta de dados:
```post
[POST] http://localhost:8080/api/add
```

RAW post data
```json
{
    "origin": "REC",
    "destination": "FLN",
    "price": 10
}
```

Resposta correta
```json
{
    "Error": [],
    "Success": true,
    "Length": 1,
    "Object": [
        {
            "Price": 10,
            "Route": "REC - FLN"
        }
    ]
}
```

Entrada incorreta de dados:
```post
[POST] http://localhost:8080/api/add
```

RAW post data
```json
{
    "origin": "RE",
    "destination": "FLN",
    "price": 1
}
```

Resposta incorreta:
```json
{
    "Error": [
        "origin code must be three letters long"
    ],
    "Success": false,
    "Length": 0,
    "Object": []
}
```

**/src/terminal**:

Módulo que possibilita o uso do terminal

**/src/testDataSource:**

Como não podia usar framework, implementei um banco de dados extremamente simples, o 
suficiente para fazer o código funcionar e mais nada.

**/src/tooServer**: 

Roda o servidor rest como uma thread.

**/src/toolTerminal**:

Faz o terminal rodar como uma thread.

**/src/util**:

Faz a busca de arquivo pelo nome na árvore de diretórios.

### Explicações

Todas as funções são curtas, com responsabilidade única e têm nomes coerentes. Todos os
retornos de funções são nomeados para deixar o entendimento do código mais claro. 

Os testes têm a seguinte prioridade, tudo o que tenho dúvida se estou certo, tudo o que
for crítico para o sistema, tudo o que for de valor para o cliente, coisas de menor 
importância.

Quando o código é algo que sai do dia a dia, como foi o caso desse, costumo escrever o
código de forma a funcionar, depois organizo, deixando possíveis melhorias para ser 
decidido se o esforço vale a pena.

Costumo deixar as funções separadas por arquivos, pois, facilita ver quem ainda não tem 
teste, além de ficar mais fácil encontrar.

Estou ciente de separar os commits do git em grupos de mudanças/correções para que fique 
fácil reverter um possível erro, mas, nesse código, a minha única preocupação foi 
entregar, devido ao cansaço.

Como regra dentro da Ahgora Sistemas e do Banco Neon, possíveis bugs geram panic() para
que um alerta seja gerado no servidor, mantive este hábito nesse código. 

> Não havia regra para trecho de voo gratuito, então, não me preocupei.

> A busca por arquivo pode fazer a atualização do CSV ser feita em arquivo diferente do que 
> você está olhando.

> Esqueci de colocar uma regra para rotas repetidas

```hex
43 6f 6e 74 61 74 6f 3a 0a 4e 6f 6d 65 3a 20 48 65 6c 6d 75 74 20 4b 65 6d 70 65 72 0a 54 
65 6c 65 66 6f 6e 65 3a 20 38 31 39 39 39 32 36 38 37 34 34 0a 45 2d 6d 61 69 6c 3a 20 20 
68 65 6c 6d 75 74 2e 6b 65 6d 70 65 72 40 67 6d 61 69 6c 2e 63 6f 6d
```
