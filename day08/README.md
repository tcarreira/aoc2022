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

