# prova

Olá pessoal, a melhor forma de conhecer um desenvolvedor é seguir o código, então, este
é o [meu git](https://github.com/helmutkemper) está aí para ser avaliado.

[Essa é a minha contribuição](https://github.com/helmutkemper/golang.solid.kiss.complexity.measure/wiki)
a comunidade golang e fala sobre qualidade de código e dicas.

Pessoal, todos os módulos deveriam ser publicados on-line para ficar dentro do padrão 
Golang, e usar módulos, mas, achei mais fácil deixar o código em um único arquivo 
compactado para a entrega.

Como solicitado, não frameworks, apenas códigos nativos do próprio Golang.

Todas as funções são curtas, com responsabilidade única e têm nomes coerentes. Todos os
retornos de funções são nomeados para deixar o entendimento do código mais claro. 

Os testes têm a seguinte prioridade, tudo o que tenho dúvida se estou certo, tudo o que
for crítico para o sistema, tudo o que for de valor para o cliente, coisas de menor 
importância.

Quando o código é algo que sai do dia a dia, como foi o caso desse, costumo escrever o
código de forma a funcionar, depois organizo, deixando possíveis melhorias para ser 
decidido se o esforço vale a pena.

Costumo deixar as funções separadas por arquivos, pois, facilita ver quem ainda não tem 
teste e fica mais fácil de encontrar.

Estou ciente de separar os commits do git em grupos de mudanças/correções para que fique 
fácil reverter um possível erro, mas, nesse código, a minha única preocupação foi 
entregar, devido ao cansaço.

Como regra dentro da Ahgora Sistemas e do Banco Neon, possíveis bugs geram panic() para
que um alerta seja gerado no servidor, mantive este hábito nesse código. 

GOROOT=C:\Users\helmut.kemper.CS\go\go1.14.3 #gosetup
GOPATH=C:\Users\helmut.kemper.CS\Desktop\trabalho\prova;C:\Users\helmut.kemper.CS\GolandProjects\Libraries #gosetup
C:\Users\helmut.kemper.CS\go\go1.14.3\bin\go.exe test -json ./... -run "Test.*" #gosetup






















