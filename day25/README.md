# --- Dia 25: Cheio de Ar Quente ---

Quando a expedição finalmente chega ao ponto de extração, vários [balões de ar quente](https://pt.wikipedia.org/wiki/Bal%C3%A3o_de_ar_quente) grandes descem para encontrá-lo. As equipes rapidamente começam a descarregar o equipamento que os balões trouxeram: muitos kits de balões de ar quente, alguns tanques de combustível e uma *máquina de aquecimento de combustível*.

A máquina de aquecimento de combustível é uma nova adição ao processo. Quando esta montanha era um vulcão, a temperatura ambiente era mais razoável; agora, está tão frio que o combustível não funcionará sem ser aquecido primeiro.

Os Elfos, aparentemente em uma tentativa de fazer a nova máquina se sentir bem-vinda, já anexaram um par de [olhos malucos](https://en.wikipedia.org/wiki/Googly_eyes) e começaram a chamá-la de "Bob".

Para aquecer o combustível, Bob precisa saber a quantidade total de combustível que será processada com antecedência para que possa calibrar corretamente a saída de calor e a vazão. Essa quantidade é simplesmente a *soma* dos requisitos de combustível de todos os balões de ar quente, e esses requisitos de combustível estão listados claramente na lateral do queimador de cada balão de ar quente.

Você supõe que os elfos não terão problemas para somar alguns números e está prestes a descobrir qual balão é seu quando você recebe um toque no ombro. Aparentemente, os requisitos de combustível usam números escritos em um formato que os elfos não reconhecem; previsivelmente, eles gostariam da sua ajuda para decifrá-los.

Você faz uma lista de todos os requisitos de combustível (sua entrada do puzzle), mas também não reconhece o formato numérico. Por exemplo:

```
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122

```

Felizmente, Bob está rotulado com um número de telefone de suporte. Para não ser dissuadido, você liga e pede ajuda.

"Isso mesmo, basta fornecer a quantidade de combustível para o... ah, para mais de um queimador? Sem problema, você só precisa adicionar nossas Unidades de Combustível Numeral-Analógico Especial. Patente pendente! Eles são muito melhores do que os números normais para..."

Você menciona que está muito frio aqui e pergunta se eles podem pular adiante.

"Ok, nossas Unidades Especiais de Combustível Numeral-Analógico - SNAFU para abreviar (Special Numeral-Analogue Fuel Units) - são como números normais. Você sabe como começando à direita, números normais têm um lugar de unidades, um lugar de dezenas, um lugar de centenas, e assim por diante, onde o dígito em cada lugar diz quantos desse valor você tem?"

"SNAFU funciona da mesma maneira, exceto que usa potências de cinco em vez de dez. Começando da direita, você tem uma casa de um, uma casa de cinco, uma casa de vinte e cinco, uma casa de cento e vinte e cinco, e assim por diante. É muito fácil!"

Você pergunta por que alguns dos dígitos se parecem com `-` ou `=` em vez de "dígitos".

"Sabe, eu nunca perguntei aos engenheiros por que eles fizeram isso. Em vez de usar dígitos de quatro a zero, os dígitos são `2`, `1`, `0`, *menos* (escrito `-`) e *menos duplo* (escrito `=`). Menos vale -1 e menos duplo vale -2."

"Então, porque dez (em números normais) é dois cincos e nenhum um, em SNAFU é escrito `20`. Como oito (em números normais) é dois cincos menos dois uns, é escrito `2=`."

"Você também pode fazer isso no sentido inverso. Digamos que você tenha o número SNAFU `2=-01`. Isso é `2` na casa dos 625s, `=` (duplo menos) na casa dos 125s, `-` (menos) na casa dos 25s, `0` na casa dos 5s e `1` na casa dos 1s. (2 vezes 625) mais (-2 vezes 125) mais (-1 vezes 25) mais (0 vezes 5) mais (1 vezes 1). São 1250 mais -250 mais -25 mais 0 mais 1. *976*!"

"Vejo aqui que você está conectado através do nosso serviço de ligação premium, então vou transmitir nosso prático folheto SNAFU para você agora. Precisa vade mais alguma coisa?"

Você pergunta se ao menos o combustível irá funcionar nessas temperaturas.

"Espere, *quão* frio está? Não há *como* o combustível - ou *qualquer* combustível - funcionar nessas condições! Existem apenas alguns lugares no.. onde você disse que está mesmo?"

Nesse momento, você percebe que um dos Elfos despeja algumas gotas de um recipiente em forma de floco de neve em um dos tanques de combustível, agradece ao representante de suporte por seu tempo e desliga a ligação.

A brochura SNAFU contém mais alguns exemplos de números decimais ("normais") e suas correspondentes em SNAFU:

```
  Decimal          SNAFU
        1              1
        2              2
        3             1=
        4             1-
        5             10
        6             11
        7             12
        8             2=
        9             2-
       10             20
       15            1=0
       20            1-0
     2022         1=11-2
    12345        1-0---0
314159265  1121-1110-1=0

```

Com base nesse processo, os números SNAFU no exemplo acima podem ser convertidos em números decimais da seguinte forma:

```
 SNAFU  Decimal
1=-0-2     1747
 12111      906
  2=0=      198
    21       11
  2=01      201
   111       31
 20012     1257
   112       32
 1=-1=      353
  1-12      107
    12        7
    1=        3
   122       37

```

Em decimal, a soma desses números é `4890`.

Ao inserir esse número no painel de Bob, você descobre que alguns botões que esperava não existem. Em vez disso, você se depara com botões `=`, `-`, `0`, `1` e `2`. Bob precisa do valor de entrada como um número SNAFU, não em decimal.

Invertendo o processo, você pode determinar que para o número decimal `4890`, o número SNAFU que você precisa fornecer ao painel de Bob é `2=-1=0`.

Os Elfos estão começando a ficar com frio. *Que número SNAFU você fornece ao painel de Bob?*

