// FlightSuggestion é um pacote feito para sugerir rotas de voo baratas
//
// Pessoal, Na minha antiga empresa, e na atual, há um regra de colocar panic() nos erros
// graves, assim, o servidor cai e um alarme é gerado para a equipe.
// Minha base é programação de microcontroladores (80C51) em C e Assembley, por isto, eu
// tenho consciência da necessidade de poupar memória e usar ponteiro.
// Também tenho consciência de que o código pode ser colocado para rodar em um container
// com pouca memória e o carregamento de uma lista de trechos de voo muito grandes pode
// consumir muita memória e fazer o container travar, dependendo do limite de memória
// para o mesmo.
//
// Todos os códigos são formados por funções pequenas de responsabilidade única, para que
// o código fique mais fácil de ser entendido e seguido. Nesse ponto, eu sigo os
// trabalhos do Uncle Bob e do Thomas J. McCabe. Também sigo a filosofia do Kiss - "Keep
// It Stupid Simple / Keep It Simple, Stupid!" e parte do meu trabalho pode ser
// acompanhado em
// https://github.com/helmutkemper/golang.solid.kiss.complexity.measure/wiki
//
// Quanto a função flightSuggestion.FindCheapestRoute() ela ficou bem dividia e fácil de
// ser seguida e entendida. A ideia básica é sempre usar nomes de métodos explicativos,
// já que programadores cansados tendem a esquecer os cometários a medida que o código
// muda.
// Nesse ponto, detesto comentários tipo, FindCheapestRoute(), procura a rota mais
// barata, embora, nem sempre tenha muito o que fazer.
// Todo código deve ter nos comentários não o que o código faz, mas, porque ele faz o que
// faz. Já tive de refatorar um código com constantes dentro dele que nem a equipe
// original lembrava o motivo, por isto, eu sempre comento o motivo quando relevante, e
// se o motivo for uma lei, minha antiga empresa era especializada em controle de ponto,
// eu sempre coloco no comentário a parte relevante da lei e o link, para que isto não
// se perca no tempo.
package flightSuggestion
