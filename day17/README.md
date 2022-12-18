# --- Dia 17: Fluxo Piroclástico ---

Seu dispositivo portátil localizou uma saída alternativa da caverna para você e os elefantes. O chão está tremendo quase continuamente agora, mas as estranhas válvulas lhe deram algum tempo. Está definitivamente ficando mais quente aqui.

Os túneis eventualmente levam até uma câmara muito alta e estreita. Grandes rochas de formato estranho vindas de cima estão caindo na câmara, presumivelmente devido a todo a tremor. Se você não conseguir descobrir onde as pedras cairão em seguida, você pode ficar *esmagado*!

Os cinco tipos de rocha têm as seguintes formas peculiares, onde `#` é rocha e `.` é espaço vazio:

```
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##

```

As rochas caem na ordem mostrada acima: primeiro a forma `-`, depois a forma `+`, e assim por diante. Uma vez que o final da lista é alcançado, a mesma ordem se repete: a forma `-` cai em primeiro, sexto, 11º, 16º, etc.

As rochas não giram, mas são empurradas por jatos de gás quente que saem das próprias paredes. Uma varredura rápida revela o efeito que os jatos de gás quente terão nas rochas à medida que caem (sua entrada do puzzle).

Por exemplo, suponha que este seja o padrão de jato em sua caverna:

```
>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>

```

Em padrões de jato, `<` significa um empurrão para a esquerda, enquanto `>` significa um empurrão para a direita. O padrão acima significa que os jatos empurrarão uma rocha caindo para a direita, depois para a direita, depois para a direita, depois para a esquerda, depois para a esquerda, depois para a direita e assim por diante. Se o fim da lista for atingido, ela se repete.

A câmara alta e vertical tem exatamente *sete unidades de largura*. Cada pedra aparece de forma a que sua borda esquerda esteja a duas unidades de distância da parede esquerda e sua borda inferior esteja três unidades acima da rocha mais alta da sala (ou do chão, se não houver).

Depois que uma pedra aparece, ela alterna entre *ser empurrada por um jato de gás quente* uma unidade (na direção indicada pelo próximo símbolo no padrão do jato) e depois *cair uma unidade*. Se algum movimento fizer com que qualquer parte da rocha se mova para as paredes, chão ou uma rocha parada, o movimento não ocorre. Se um movimento *para baixo* fizer com que uma pedra em queda se mova para o chão ou uma pedra já caída, a pedra em queda para onde está (tendo pousado em algo) e imediatamente uma nova pedra começa a cair.

Desenhando pedras caindo com `@` e pedras paradas com `#`, o padrão de jato no exemplo acima se manifesta da seguinte forma:

```
A primeira pedra começa a cair:
|..@@@@.|
|.......|
|.......|
|.......|
+-------+

Jato de gás empurra a pedra para a direita:
|...@@@@|
|.......|
|.......|
|.......|
+-------+

Rocha cai 1 unidade:
|...@@@@|
|.......|
|.......|
+-------+

Jato de gás empurra a pedra para a direita, mas nada acontece:
|...@@@@|
|.......|
|.......|
+-------+

Rocha cai 1 unidade:
|...@@@@|
|.......|
+-------+

Jato de gás empurra a pedra para a direita, mas nada acontece:
|...@@@@|
|.......|
+-------+

Rocha cai 1 unidade:
|...@@@@|
+-------+

Jato de gás empurra a rocha para a esquerda:
|..@@@@.|
+-------+

A rocha cai 1 unidade, fazendo com que ela pare:
|..####.|
+-------+

Uma nova pedra começa a cair:
|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|.......|
|..####.|
+-------+

Jato de gás empurra a rocha para a esquerda:
|..@....|
|.@@@...|
|..@....|
|.......|
|.......|
|.......|
|..####.|
+-------+

Rocha cai 1 unidade:
|..@....|
|.@@@...|
|..@....|
|.......|
|.......|
|..####.|
+-------+

Jato de gás empurra a pedra para a direita:
|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|..####.|
+-------+

Rocha cai 1 unidade:
|...@...|
|..@@@..|
|...@...|
|.......|
|..####.|
+-------+

Jato de gás empurra a rocha para a esquerda:
|..@....|
|.@@@...|
|..@....|
|.......|
|..####.|
+-------+

Rocha cai 1 unidade:
|..@....|
|.@@@...|
|..@....|
|..####.|
+-------+

Jato de gás empurra a pedra para a direita:
|...@...|
|..@@@..|
|...@...|
|..####.|
+-------+

A rocha cai 1 unidade, fazendo com que ela pare:
|...#...|
|..###..|
|...#...|
|..####.|
+-------+

Uma nova pedra começa a cair:
|....@..|
|....@..|
|..@@@..|
|.......|
|.......|
|.......|
|...#...|
|..###..|
|...#...|
|..####.|
+-------+

```

No momento em que cada uma das próximas pedras começa a cair, você veria isto:

```
|..@....|
|..@....|
|..@....|
|..@....|
|.......|
|.......|
|.......|
|..#....|
|..#....|
|####...|
|..###..|
|...#...|
|..####.|
+-------+

|..@@...|
|..@@...|
|.......|
|.......|
|.......|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@@@.|
|.......|
|.......|
|.......|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|.......|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|....@..|
|....@..|
|..@@@..|
|.......|
|.......|
|.......|
|..#....|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@....|
|..@....|
|..@....|
|..@....|
|.......|
|.......|
|.......|
|.....#.|
|.....#.|
|..####.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@...|
|..@@...|
|.......|
|.......|
|.......|
|....#..|
|....#..|
|....##.|
|....##.|
|..####.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@@@.|
|.......|
|.......|
|.......|
|....#..|
|....#..|
|....##.|
|##..##.|
|######.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

```

Para provar aos elefantes que sua simulação é precisa, eles querem saber qual será a altura da torre depois de 2022 rochas pararem (mas antes que a 2023ª rocha comece a cair). Neste exemplo, a torre de rochas terá `3068` unidades de altura.

*Quantas unidades de altura terá a torre de rochas depois que 2022 rochas pararem de cair?*

## --- Parte Dois ---

Os elefantes não estão impressionados com sua simulação. Eles exigem saber qual será a altura da torre depois que `1000000000000` rochas pararem! Só então eles se sentirão confiantes o suficiente para prosseguir pela caverna.

No exemplo acima, a torre teria `1514285714288` unidades de altura!

*Qual será a altura da torre depois que `1000000000000` rochas pararem?*

