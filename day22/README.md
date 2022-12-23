# --- Dia 22: Mapa de Macacos ---

Os macacos levam você a uma trilha surpreendentemente fácil pela selva. Eles estão até indo mais ou menos na direção certa de acordo com o Sistema de Posicionamento do Arvoredo do seu dispositivo portátil.

Enquanto você caminha, os macacos explicam que o arvoredo é protegido por um *campo de forças*. Para passar pelo campo de forças, você deve fornecer uma senha; fazer isso envolve traçar um *caminho* específico em um tabuleiro de formato estranho.

Pelo menos, você tem quase a certeza de que é isso que deve fazer; os elefantes não são exatamente fluentes em macaquês.

Os macacos dão a você anotações que tiraram quando viram a senha fornecida pela última vez (sua entrada do puzzle).

Por exemplo:

```
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5

```

A primeira metade das anotações dos macacos é um *mapa do tabuleiro*. É composto por um conjunto de *quadrados em aberto* (nos quais você pode se mover, desenhados com `.`) e *paredes sólidas* (quadrados nos quais você não pode entrar, desenhados com `#`).

A segunda metade é uma descrição do *caminho que você deve seguir*. Consiste em números e letras alternados:


  - Um *número* indica o *número de quadrados a serem movidos* na direção que você está. Se você bater em uma parede, você para de avançar e continua com a próxima instrução.

  - Uma *letra* indica se deve rodar 90 graus *no sentido horário* (`R`) ou *no sentido anti-horário* (`L`). A rotação acontece no local; isso não altera o quadrado em que você está no momento.


Portanto, um caminho como `10R5` significa "vá em frente 10 quadrados, gire 90 graus no sentido horário e depois avance 5 quadrados".

Você começa o caminho no quadrado em aberto mais à esquerda da linha superior de quadrados. Inicialmente, você está olhando *para a direita* (da perspectiva de como o mapa é desenhado).

Se uma instrução de movimento tirar você do mapa, você *dá a volta* para o outro lado do tabuleiro. Em outras palavras, se o seu próximo quadrado estiver fora do tabuleiro, você deve olhar o máximo que puder na direção oposta à que você tem atualmente até encontrar a borda oposta do tabuleiro e reaparecer lá.

Por exemplo, se você estiver em `A` e virado para a direita, o quadrado à sua frente está marcada como `B`; se você estiver em `C` e voltado para baixo, o quadrado à sua frente está marcado como `D`:

```
        ...#
        .#..
        #...
        ....
...#.D.....#
........#...
B.#....#...A
.....C....#.
        ...#....
        .....#..
        .#......
        ......#.

```

É possível que o próximo quadrado (depois de dar a volta) seja uma *parede*; isso ainda conta como se houvesse uma parede à sua frente e, portanto, o movimento para antes de você realmente passar para o outro lado do tabuleiro.

Ao desenhar a *última direção que você teve* com uma seta em cada quadrado que você visita, o caminho completo percorrido pelo exemplo acima fica assim:

```
        >>v#    
        .#v.    
        #.v.    
        ..v.    
...#...v..v#    
>>>v...>#.>>    
..#v...#....    
...>>>>v..#.    
        ...#....
        .....#..
        .#......
        ......#.

```

Para concluir o fornecimento da senha para este estranho dispositivo de entrada, você precisa determinar os números para sua *linha*, *coluna* e *direção* finais conforme sua posição final aparece da perspectiva do mapa original. As linhas começam em `1` no topo e contam para baixo; as colunas começam em `1` à esquerda e contam para a direita. (No exemplo acima, linha 1, coluna 1 refere-se ao espaço vazio sem quadrados no canto superior esquerdo.) A direção é `0` para direita (`>`), `1` para baixo (`v`), `2` para esquerda (`<`) e `3` para cima (`^`). A *senha final* é a soma de 1000 vezes a linha, 4 vezes a coluna e a direção.

No exemplo acima, a linha final é `6`, a coluna final é `8` e a direção final é `0`. Portanto, a senha final é 1000 * 6 + 4 * 8 + 0: `6032`.

Siga o caminho dado nas notas dos macacos. *Qual é a senha final?*

## --- Parte Dois ---

Ao aproximar o campo de força, você pensa ouvir alguns Elfos à distância. Talvez eles já tenham chegado?

Você se aproxima do estranho *dispositivo de entrada*, mas não é exatamente o que os macacos desenharam em suas anotações. Em vez disso, você se depara com um grande *cubo*; cada uma de suas seis faces é um quadrado de quadrados com 50x50.

Para ser justo, o mapa dos macacos *tem* seis regiões 50x50. Se você *dobrar o mapa com cuidado*, deverá conseguir moldá-lo na forma de um cubo!

No exemplo acima, as seis (menores, 4x4) faces do cubo são:

```
        1111
        1111
        1111
        1111
222233334444
222233334444
222233334444
222233334444
        55556666
        55556666
        55556666
        55556666

```

Você ainda começa na mesma posição e com a mesma face de antes, mas as regras de *dar a volta* são diferentes. Agora, se você sair do tabuleiro, em vez disso você *continuará ao redor do cubo*. Da perspectiva do mapa, isso pode parecer um pouco estranho. No exemplo acima, se você estiver em A e mover para a direita, chegará em B virado para baixo; se você estiver em C e descer, chegará em D voltado para cima:

```
        ...#
        .#..
        #...
        ....
...#.......#
........#..A
..#....#....
.D........#.
        ...#..B.
        .....#..
        .#......
        ..C...#.

```

As paredes ainda bloqueiam seu caminho, mesmo que estejam em uma face diferente do cubo. Se você estiver em E voltado para cima, seu movimento é bloqueado pela parede marcada pela seta:

```
        ...#
        .#..
     -->#...
        ....
...#..E....#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

```

Usando o mesmo método de desenhar a *última face que você tinha* com uma seta em cada quadrado que você visita, o caminho completo percorrido pelo exemplo acima agora se parece com isto:

```
        >>v#    
        .#v.    
        #.v.    
        ..v.    
...#..^...v#    
.>>>>>^.#.>>    
.^#....#....    
.^........#.    
        ...#..v.
        .....#v.
        .#v<<<<.
        ..v...#.

```

A senha final ainda é calculada a partir de sua posição final e direção da perspectiva do mapa. Neste exemplo, a linha final é `5`, a coluna final é `7` e a direção final é `3`, então a senha final é 1000 * 5 + 4 * 7 + 3 = `5031`.

Dobre o mapa em um cubo, *depois* siga o caminho dado nas notas dos macacos. *Qual é a senha final?*

