# --- Dia 18: Pedregulhos Ferventes ---

Você e os elefantes finalmente alcançam ar fresco. Você emergiu perto da base de um grande vulcão que parece estar em erupção ativa! Felizmente, a lava parece estar fluindo para longe de você e em direção ao oceano.

Mesmo assim, pedaços de lava ainda estão sendo ejetados em sua direção, então você fica se protegendo um pouco mais na saída da caverna. Fora da caverna, você pode ver a lava caindo em um lago e ouvi-la sibilar alto enquanto se solidifica.

Dependendo dos compostos específicos da lava e da velocidade com que esfria, pode estar se formando [obsidiana](https://pt.wikipedia.org/wiki/Obsidiana)! A taxa de resfriamento deve ser baseada na área de superfície das gotículas de lava, então você faz uma varredura rápida de uma gotícula enquanto ela passa voando por você (sua entrada do puzzle).

Devido à rapidez com que a lava se move, a varredura não é muito boa; sua resolução é bastante baixa e, como resultado, ela aproxima a forma da gota de lava com *cubos 1x1x1 em uma matriz 3D*, cada um dado como sua posição `x,y,z`.

Para aproximar a área da superfície, conte o número de lados de cada cubo que não estão imediatamente conectados a outro cubo. Então, se sua varredura fosse apenas dois cubos adjacentes como `1,1,1` e `2,1,1`, cada cubo teria um único lado coberto e cinco lados expostos, uma área de superfície total de `10` lados.

Aqui está um exemplo maior:

```
2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5

```

No exemplo acima, depois de contar todos os lados que não estão conectados a outro cubo, a área total da superfície é `64`.

*Qual é a área de superfície das gotícula de lava do seu varrimento?*

