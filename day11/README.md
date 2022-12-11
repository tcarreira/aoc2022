# --- Dia 11: Macaco no Meio ---

Quando você finalmente começa a subir o rio, percebe que sua mochila está muito mais leve do que você se lembra. Nesse momento, um dos itens de sua mochila voa sobre você. Os macacos estão jogando [Bobinho](https://pt.wikipedia.org/wiki/Bobinho) com suas coisas perdidas!

Para recuperar suas coisas, você precisa prever onde os macacos vão jogar seus itens. Depois de uma observação cuidadosa, você percebe que os macacos operam com base em *o quão preocupado você está com cada item*.

Você faz algumas anotações (sua entrada do puzzle) sobre os itens que cada macaco possui atualmente, o quanto você está preocupado com esses itens e como o macaco toma decisões com base no seu nível de preocupação. Por exemplo:

```
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1

```

Cada macaco tem vários atributos:


  - `Starting items` (Itens iniciais) lista seu *nível de preocupação* para cada item que o macaco está segurando no momento na ordem em que serão inspecionados.

  - `Operation` (Operação) mostra como seu nível de preocupação muda conforme o macaco inspeciona um item. (Uma operação como `new = old * 5` (novo = antigo * 5) significa que seu nível de preocupação depois que o macaco inspecionou o item é cinco vezes maior que seu nível de preocupação antes da inspeção.)

  - `Test` mostra como o macaco usa seu nível de preocupação para decidir onde jogar um item a seguir.
  
  - `If true` mostra o que acontece com um item se o `Test` for verdadeiro.

  - `If false` mostra o que acontece com um item se o `Test` for falso.





Após cada macaco inspecionar um item, mas antes de testar seu nível de preocupação, seu alívio pelo fato de a inspeção do macaco não ter danificado o item faz com que seu nível de preocupação seja *dividido por três* e arredondado para o número inteiro mais próximo.

Os macacos se revezam inspecionando e jogando itens. No *turno* de um único macaco, ele inspeciona e joga todos os itens que está segurando, um de cada vez e na ordem listada. O macaco `0` vai primeiro, depois o macaco `1`, e assim por diante até que cada macaco tenha tido uma vez. O processo de todos os macacos terem tido a sua vez é chamado de *rodada*.

Quando um macaco joga um item para outro macaco, o item vai para o *final* da lista do macaco que recebe. Um macaco que começa uma rodada sem itens pode acabar inspecionando e jogando muitos itens no momento em que chega a sua vez. Se um macaco não estiver segurando nenhum item no início de seu turno, seu turno termina.

No exemplo acima, a primeira rodada procede da seguinte forma:

```
Macaco 0:
   Macaco inspeciona um item com um nível de preocupação de 79.
     O nível de preocupação é multiplicado por 19 para 1501.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 500.
     O nível de preocupação atual não é divisível por 23.
     Item com nível de preocupação 500 é jogado para macaco 3.
   Macaco inspeciona um item com um nível de preocupação de 98.
     O nível de preocupação é multiplicado por 19 para 1862.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 620.
     O nível de preocupação atual não é divisível por 23.
     O item com nível de preocupação 620 é lançado para o macaco 3.
Macaco 1:
   Macaco inspeciona um item com um nível de preocupação de 54.
     O nível de preocupação aumenta de 6 para 60.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 20.
     O nível de preocupação atual não é divisível por 19.
     O item com nível de preocupação 20 é lançado para o macaco 0.
   Macaco inspeciona um item com nível de preocupação 65.
     O nível de preocupação aumenta de 6 para 71.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 23.
     O nível de preocupação atual não é divisível por 19.
     O item com nível de preocupação 23 é lançado para o macaco 0.
   Macaco inspeciona um item com um nível de preocupação de 75.
     O nível de preocupação aumenta de 6 para 81.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 27.
     O nível de preocupação atual não é divisível por 19.
     O item com nível de preocupação 27 é lançado para o macaco 0.
   Macaco inspeciona um item com um nível de preocupação de 74.
     O nível de preocupação aumenta de 6 para 80.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 26.
     O nível de preocupação atual não é divisível por 19.
     O item com nível de preocupação 26 é lançado para o macaco 0.
Macaco 2:
   Macaco inspeciona um item com um nível de preocupação de 79.
     O nível de preocupação é multiplicado por si mesmo para 6241.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 2080.
     O nível de preocupação atual é divisível por 13.
     O item com nível de preocupação 2080 é lançado para o macaco 1.
   Macaco inspeciona um item com um nível de preocupação de 60.
     O nível de preocupação é multiplicado por si mesmo para 3600.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 1200.
     O nível de preocupação atual não é divisível por 13.
     O item com nível de preocupação 1200 é lançado para o macaco 3.
   Macaco inspeciona um item com um nível de preocupação de 97.
     O nível de preocupação é multiplicado por si mesmo para 9409.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 3136.
     O nível de preocupação atual não é divisível por 13.
     O item com nível de preocupação 3136 é lançado para o macaco 3.
Macaco 3:
   Macaco inspeciona um item com um nível de preocupação de 74.
     O nível de preocupação aumenta de 3 para 77.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 25.
     O nível de preocupação atual não é divisível por 17.
     O item com nível de preocupação 25 é jogado para o macaco 1.
   Macaco inspeciona um item com um nível de preocupação de 500.
     O nível de preocupação aumenta de 3 para 503.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 167.
     O nível de preocupação atual não é divisível por 17.
     O item com nível de preocupação 167 é lançado para o macaco 1.
   Macaco inspeciona um item com um nível de preocupação de 620.
     O nível de preocupação aumenta em 3 para 623.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 207.
     O nível de preocupação atual não é divisível por 17.
     O item com nível de preocupação 207 é lançado para o macaco 1.
   Macaco inspeciona um item com um nível de preocupação de 1200.
     O nível de preocupação aumenta em 3 para 1203.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 401.
     O nível de preocupação atual não é divisível por 17.
     O item com nível de preocupação 401 é lançado para o macaco 1.
   Macaco inspeciona um item com um nível de preocupação de 3136.
     O nível de preocupação aumenta em 3 para 3139.
     O macaco fica entediado com o item. O nível de preocupação é dividido por 3 para 1046.
     O nível de preocupação atual não é divisível por 17.
     O item com nível de preocupação 1046 é lançado para o macaco 1.

```

Após a rodada 1, os macacos estão segurando itens com estes níveis de preocupação:

```
Macaco 0: 20, 23, 27, 26
Macaco 1: 2080, 25, 167, 207, 401, 1046
Macaco 2:
Macaco 3:

```

Os macacos 2 e 3 não estão segurando nenhum item no final da rodada; ambos inspecionaram itens durante a rodada e jogaram todos antes do final da rodada.

Este processo continua por mais algumas rodadas:

```
Após a rodada 2, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 695, 10, 71, 135, 350
Macaco 1: 43, 49, 58, 55, 362
Macaco 2:
Macaco 3:

Após a rodada 3, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 16, 18, 21, 20, 122
Macaco 1: 1468, 22, 150, 286, 739
Macaco 2:
Macaco 3:

Após a rodada 4, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 491, 9, 52, 97, 248, 34
Macaco 1: 39, 45, 43, 258
Macaco 2:
Macaco 3:

Após a rodada 5, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 15, 17, 16, 88, 1037
Macaco 1: 20, 110, 205, 524, 72
Macaco 2:
Macaco 3:

Após a rodada 6, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 8, 70, 176, 26, 34
Macaco 1: 481, 32, 36, 186, 2190
Macaco 2:
Macaco 3:

Após a rodada 7, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 162, 12, 14, 64, 732, 17
Macaco 1: 148, 372, 55, 72
Macaco 2:
Macaco 3:

Após a rodada 8, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 51, 126, 20, 26, 136
Macaco 1: 343, 26, 30, 1546, 36
Macaco 2:
Macaco 3:

Após a rodada 9, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 116, 10, 12, 517, 14
Macaco 1: 108, 267, 43, 55, 288
Macaco 2:
Macaco 3:

Após a rodada 10, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 91, 16, 20, 98
Macaco 1: 481, 245, 22, 26, 1092, 30
Macaco 2:
Macaco 3:

...

Após a rodada 15, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 83, 44, 8, 184, 9, 20, 26, 102
Macaco 1: 110, 36
Macaco 2:
Macaco 3:

...

Após a rodada 20, os macacos estão segurando itens com estes níveis de preocupação:
Macaco 0: 10, 12, 14, 26, 34
Macaco 1: 245, 93, 53, 199, 115
Macaco 2:
Macaco 3:

```

Perseguir todos os macacos ao mesmo tempo é impossível; você terá que se concentrar nos *dois macacos mais ativos* se quiser alguma esperança de recuperar suas coisas. Conte o *número total de vezes que cada macaco inspeciona os itens* em 20 rodadas:

```
Macaco 0 inspecionou itens 101 vezes.
Macaco 1 inspecionou itens 95 vezes.
Macaco 2 inspecionou itens 7 vezes.
Macaco 3 inspecionou itens 105 vezes.

```

Neste exemplo, os dois macacos mais ativos inspecionaram os itens 101 e 105 vezes. O nível de *macacadas* nesta situação pode ser encontrado multiplicando-os: `10605`.

Descubra quais macacos perseguir contando quantos itens eles inspecionam em 20 rodadas. *Qual é o nível de macacadas depois de 20 rodadas de travessuras de símios?*

