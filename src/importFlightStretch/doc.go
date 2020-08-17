// Este código foi feito para o test de Helmut Kemper (helmut.kemper@gmail.com) que está
// concorrendo a vaga de desenvolvedor Golang.
//
// Este pacote importa arquivos CSV com trechos de voos e valores no seguinte formato:
//   IATA code da origem;
//   IATA cede do destino;
//   Preço;
//
// Exemplo:
//   GRU,BRC,10
//   BRC,SCL,5
//   GRU,CDG,75
//
// Atenção:
//   Não foi solicitado que o preço fosse em ponto flutuante e nem levasse em conta a
//   conversão de moedas
package importFlightStretch
