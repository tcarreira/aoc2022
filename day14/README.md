# --- Dia 14: Reservatório de Regolito ---

O sinal de socorro leva você a uma cachoeira gigante! Na verdade, espere - o sinal parece vir da própria cachoeira, e isso não faz o menor sentido. No entanto, você repara em um pequeno caminho que leva *atrás* da cachoeira.

Correção: o sinal de socorro leva você para trás de uma cachoeira gigante! Parece haver um grande sistema de cavernas aqui, e o sinal definitivamente leva mais para dentro.

Ao começar a se aprofundar no subsolo, você sente o chão tremer por um momento. A areia começa a entrar na caverna! Se você não descobrir rapidamente para onde a areia está indo, poderá ficar preso rapidamente!

Felizmente, sua [familiaridade](https://adventofcode.com/2018/day/17) com a análise do caminho de material em queda será útil aqui. Você digitaliza uma fatia vertical bidimensional da caverna acima de você (sua entrada do puzzle) e descobre que é principalmente *ar* com estruturas feitas de *rocha*.

Sua varredura traça o caminho de cada estrutura de rocha sólida e informa as coordenadas `x,y` que compõem a forma do caminho, onde `x` representa a distância à direita e `y` representa a distância para baixo. Cada caminho aparece como uma única linha de texto em sua digitalização. Após o primeiro ponto de cada caminho, cada ponto indica o fim de uma linha reta horizontal ou vertical a ser traçada a partir do ponto anterior. Por exemplo:

```
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9

```

Esta varredura significa que existem dois caminhos de rocha; o primeiro caminho consiste em duas linhas retas e o segundo caminho consiste em três linhas retas. (Especificamente, o primeiro caminho consiste em uma linha de rocha de `498,4` a `498,6` e outra linha de rocha de `498,6` a `496,6`.)

A areia está entrando na caverna a partir do ponto `500,0`.

Desenhando rocha como `#`, ar como `.`, e a fonte da areia como `+`, isso se torna:

```

  4     5  5
  9     0  0
  4     0  3
0 ......+...
1 ..........
2 ..........
3 ..........
4 ....#...##
5 ....#...#.
6 ..###...#.
7 ........#.
8 ........#.
9 #########.

```

A areia é produzida *uma unidade de cada vez*, e a próxima unidade de areia não é produzida até que a unidade de areia anterior *fique parada*. Uma unidade de areia é grande o suficiente para preencher um quadrado de ar na sua varredura.

Uma unidade de areia sempre *cai um passo* se possível. Se o quadrado imediatamente abaixo estiver bloqueado (por pedra ou areia), a unidade de areia tenta se mover diagonalmente *um passo para baixo e para a esquerda*. Se esse quadrado estiver bloqueado, a unidade de areia tenta se mover diagonalmente *um passo para baixo e para a direita*. A areia continua se movendo enquanto for capaz de fazê-lo, a cada passo tentando se mover para baixo, depois para baixo à esquerda, depois para baixo à direita. Se todos os três destinos possíveis forem bloqueados, a unidade de areia *parará* e não se moverá mais, ponto em que a próxima unidade de areia será criada de volta na fonte.

Então, desenhando a areia que parou como `o`, a primeira unidade de areia simplesmente cai e para:

```
......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
......o.#.
#########.

```

A segunda unidade de areia então cai diretamente para baixo, pousa na primeira e depois para à sua esquerda:

```
......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
.....oo.#.
#########.

```

Após um total de cinco unidades de areia terem parado, elas formam este padrão:

```
......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
......o.#.
....oooo#.
#########.

```

Após um total de 22 unidades de areia:

```
......+...
..........
......o...
.....ooo..
....#ooo##
....#ooo#.
..###ooo#.
....oooo#.
...oooo#.
#########.

```

Finalmente, apenas mais duas unidades de areia podem parar:

```
......+...
..........
......o...
.....ooo..
....#ooo##
...o#ooo#.
..###ooo#.
....oooo#.
.o.oooo#.
#########.

```

Uma vez que todas as `24` unidades de areia mostradas acima tenham parado, toda a areia restante flui para fora, caindo no vazio sem fim. Apenas por diversão, o caminho que qualquer nova areia percorre antes de cair para sempre é mostrado aqui com `~`:

```
.......+...
.......~...
......~o...
.....~ooo..
....~#ooo##
...~o#ooo#.
..~###ooo#.
..~..oooo#.
.~o.oooo#.
~#########.
~..........
~..........
~..........

```

Usando sua varredura, simule a areia caindo. *Quantas unidades de areia param antes que a areia comece a fluir para o abismo abaixo?*

## --- Parte Dois ---

Você percebe que interpretou mal a varredura. Não há um *vazio sem fim* na parte inferior da digitalização - há chão e você está de pé sobre ele!

Você não tem tempo para escanear o chão, então assuma que o chão é uma linha horizontal infinita com uma coordenada `y` igual a *dois mais a coordenada `y` mais alta* de qualquer ponto em sua varredura.

No exemplo acima, a coordenada `y` mais alta de qualquer ponto é `9` e, portanto, o piso está em `y=11`. (É como se sua varredura contivesse um caminho de rocha extra como `-infinito,11 -> infinito,11`.) Com o piso adicionado, o exemplo acima agora se parece com isto:

```
        ...........+........
        ....................
        ....................
        ....................
        .........#...##.....
        .........#...#......
        .......###...#......
        .............#......
        .............#......
        .....#########......
        ....................
<-- etc #################### etc -->

```

Para encontrar um lugar seguro para ficar, você precisará simular a queda de areia até que uma unidade de areia pare em `500,0`, bloqueando totalmente a fonte e interrompendo o fluxo de areia para dentro da caverna. No exemplo acima, a situação finalmente se parece com isso depois que `93` unidades de areia param:

```
............o............
...........ooo...........
..........ooooo..........
.........ooooooo.........
........oo#ooo##o........
.......ooo#ooo#ooo.......
......oo###ooo#oooo......
.....oooo.oooo#ooooo.....
....oooooooooo#oooooo....
...ooo#########ooooooo...
..ooooo.......ooooooooo..
#########################

```

Usando sua varredura, simule a areia caindo até que a fonte da areia fique bloqueada. *Quantas unidades de areia ficam paradas?*

