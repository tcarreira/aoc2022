# --- Dia 8: Casa no Topo da Árvore ---

A expedição se depara com um trecho peculiar de árvores altas, todas cuidadosamente plantadas em uma grelha. Os Elfos explicam que uma expedição anterior plantou essas árvores como um esforço de reflorestamento. Agora, eles estão curiosos para saber se este seria um bom local para uma [casa da árvore](https://pt.wikipedia.org/wiki/Casa_da_%C3%A1rvore).

Primeiro, determine se há cobertura de árvores suficiente aqui para manter uma casa na árvore *oculta*. Para fazer isso, você precisa contar o número de árvores que são *visíveis de fora da grelha* ao olhar diretamente ao longo de uma linha ou coluna.

Os Elfos já lançaram um [quadricóptero](https://pt.wikipedia.org/wiki/Quadrotor) para gerar um mapa com a altura de cada árvore (*sua entrada do quebra-cabeça*). Por exemplo:

```
30373
25512
65332
33549
35390

```

Cada árvore é representada por um único dígito cujo valor é a sua altura, onde `0` é o mais baixo e `9` é o mais alto.

Uma árvore é *visível* se todas as outras árvores entre ela e uma borda da grelha são *mais curtas* do que ela. Considere apenas árvores na mesma linha ou coluna; isto é, apenas olhe para cima, para baixo, para a esquerda ou para a direita de qualquer árvore.

Todas as árvores ao redor da grelha são *visíveis* - como já estão na borda, não há árvores para bloquear a visão. Neste exemplo, isso deixa apenas as *nove árvores internas* para considerar:


  - O `5` superior esquerdo é *visível* da esquerda e do topo. (Não é visível da direita ou de baixo, pois outras árvores de altura `5` estão no caminho.)

  - O `5` superior central é *visível* do topo e da direita.

  - O `1` superior direito não é visível de nenhuma direção; para que ela fosse visível, precisaria haver apenas árvores de altura *0* entre ela e uma borda.

  - O `5` do meio esquerdo é *visível*, mas apenas da direita.

  - O centro `3` não é visível de nenhuma direção; para que fosse visível, precisaria haver apenas árvores de altura máxima `2` entre ele e uma borda.

  - O `3` do meio à direita é *visível* da direita.

  - Na linha inferior, o `5` do meio é *visível*, mas o `3` e o `4` não são.


Com 16 árvores visíveis na borda e outras 5 visíveis no interior, um total de `21` árvores são visíveis neste arranjo.

Considere seu mapa; *quantas árvores são visíveis de fora da grelha?*

## --- Parte dois ---

Satisfeitos com a quantidade de cobertura arbórea disponível, os Elfos só precisam saber o melhor local para construir sua casa na árvore: eles gostariam de poder ver muitas *árvores*.

Para medir a distância de visibilidade de uma determinada árvore, olhe para cima, para baixo, para a esquerda e para a direita dessa árvore; pare se chegar a uma borda ou na primeira árvore que tenha a mesma altura ou seja mais alta que a árvore em questão. (Se uma árvore estiver bem na borda, pelo menos uma de suas distâncias de visualização será zero.)

Os elfos não se importam com árvores distantes mais altas do que aquelas encontradas pelas regras acima; a casa da árvore proposta tem [beirais] grandes (https://pt.wikipedia.org/wiki/Beiral) para mantê-la seca, então eles não seriam capazes de ver mais alto do que a casa da árvore de qualquer maneira.

No exemplo acima, considere o `5` do meio na segunda linha:

```
30373
25512
65332
33549
35390

```


  - Olhando para cima, sua visão não é bloqueada; ele pode ver a árvore `1` (de altura `3`).

  - Olhando para a esquerda, sua visão é bloqueada imediatamente; ele pode ver apenas `1` árvore (de altura `5`, logo ao lado dela).

  - Olhando para a direita, sua visão não é bloqueada; ele pode ver `2` árvores.

  - Olhando para baixo, sua visão é bloqueada eventualmente; ele pode ver `2` árvores (uma de altura `3`, então a árvore de altura `5` que bloqueia sua visão).


A *pontuação cênica* de uma árvore é encontrada *multiplicando* sua distância de visualização em cada uma das quatro direções. Para esta árvore, isso é `4` (encontrado multiplicando `1 * 1 * 2 * 2`).

No entanto, você pode fazer ainda melhor: considere a árvore de altura `5` no meio da quarta linha:

```
30373
25512
65332
33549
35390

```


  - Olhando para cima, sua visão é bloqueada em `2` árvores (por outra árvore com uma altura de `5`).

  - Olhando para a esquerda, sua visão não é bloqueada; ele pode ver `2` árvores.

  - Olhando para baixo, sua visão também não é bloqueada; ele pode ver a árvore `1`.

  - Olhando para a direita, sua visão é bloqueada em `2` árvores (por uma enorme árvore de altura `9`).


A pontuação cênica desta árvore é `8` (`2 * 2 * 1 * 2`); este é o local ideal para a casa da árvore.

Considere cada árvore em seu mapa. *Qual é a maior pontuação cênica possível para qualquer árvore?*

