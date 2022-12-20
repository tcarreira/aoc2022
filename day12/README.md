# --- Dia 12: Algoritmo de Escalada ---

Você tenta entrar em contato com os Elfos usando seu *dispositivo portátil*, mas o rio que você está seguindo deve estar muito baixo para obter um sinal decente.

Você pede ao dispositivo um mapa de altura da área circundante (sua entrada do puzzle). O mapa de altura mostra a área local de cima em forma de uma grelha; a elevação de cada quadrado da grelha é dada por uma única letra minúscula, onde `a` é a elevação mais baixa, `b` é a próxima mais baixa, e assim por diante até a elevação mais alta, `z`.

Também estão incluídas no mapa de altura as marcas para sua posição atual (`S`) e o local que deve obter o melhor sinal (`E`). Sua posição atual (`S`) tem elevação `a`, e o local que deve receber o melhor sinal (`E`) tem elevação `z`.

Você gostaria de chegar a `E`, mas para economizar energia, você deve fazê-lo no *menor número de passos possível*. Durante cada passo, você pode mover exatamente um quadrado para cima, para baixo, para a esquerda ou para a direita. Para evitar a necessidade de retirar seu equipamento de escalada, a elevação do quadrado de destino pode ser *no máximo um ponto mais alto* do que a elevação do seu quadrado atual; ou seja, se a sua elevação atual for `m`, você pode ir para a elevação `n`, mas não para a elevação `o`. (Isso também significa que a elevação do quadrado de destino pode ser muito menor do que a elevação do seu quadrado atual.)

Por exemplo:

```
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi

```

Aqui, você começa no canto superior esquerdo; seu objetivo está perto do meio. Você pode começar movendo para baixo ou para a direita, mas eventualmente precisará ir em direção ao `e` na parte inferior. A partir daí, você pode espiralar até o objetivo:

```
v..v<<<<
>v.vv<<^
.>vv>E^^
..v>>>^^
..>>>>>^

```

No diagrama acima, os símbolos indicam se o caminho sai de cada quadrado movendo-se para cima (`^`), para baixo (`v`), para a esquerda (`<`) ou para a direita (`>`). O local que deve obter o melhor sinal ainda é `E`, e `.` marca os quadrados não visitados.

Este caminho atinge a meta em `31` passos, o menor possível.

*Qual é o menor número de passos necessários para mover desde a sua posição atual até o local que deve receber o melhor sinal?*

## --- Parte Dois ---

Ao subir a colina, você suspeita que os elfos vão querer transformar isso em uma trilha para caminhadas. O começo não é muito cênico; talvez você possa encontrar um ponto de partida melhor.

Para maximizar o exercício durante a caminhada, a trilha deve começar o mais baixo possível: elevação `a`. O objetivo ainda é o quadrado marcado com `E`. No entanto, a trilha ainda deve ser direta, dando o menor número de passos para atingir seu objetivo. Portanto, você precisará encontrar o caminho mais curto de *qualquer quadrado na elevação `a`* até o quadrado marcado como `E`.

Considere novamente o exemplo acima:

```
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi

```

Agora, existem seis opções para a posição inicial (cinco marcadas com `a`, mais o quadrado marcado com `S` que conta como estando na elevação `a`). Se você começar no quadrado inferior esquerdo, poderá atingir a meta mais rapidamente:

```
...v<<<<
...vv<<^
...v>E^^
.>v>>>^^
>^>>>>>^

```

Este caminho atinge a meta em apenas `29` passos, o menor número possível.

*Qual é o menor número de passos necessários para mover a partir de qualquer quadrado com elevação `a` até o local que deve receber o melhor sinal?*

