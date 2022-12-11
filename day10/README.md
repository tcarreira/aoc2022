# --- Dia 10: Tubo de Raios Catódicos ---

Você evita as cordas, mergulha no rio e nada até a margem.

Os Elfos gritam algo sobre se reencontrarem rio acima, mas o rio está muito alto para dizer exatamente o que eles estão falando. Eles terminam de cruzar a ponte e desaparecem de vista.

Situações como esta devem ser o motivo pelo qual os Elfos priorizaram o funcionamento do sistema de comunicações em seu dispositivo portátil. Você o puxa para fora da mochila, mas a quantidade de água que sai lentamente de uma grande rachadura em sua tela indica que provavelmente não será de muita utilidade imediata.

Quer dizer, *a menos* que você possa projetar um substituto para o sistema de vídeo do dispositivo! Parece ser algum tipo de tela de [tubo de raios catódicos](https://pt.wikipedia.org/wiki/Tubo_de_raios_cat%C3%B3dicos) e CPU simples que são acionados por um *circuito de relógio* preciso. O circuito do relógio marca a um ritmo constante; cada tique é chamado de *ciclo*.

Comece por descobrir o sinal que está sendo enviado pela CPU. A CPU possui um único registo, `X`, que começa com o valor `1`. Ele suporta apenas duas instruções:


  - `addx V` leva *dois ciclos* para ser concluído. *Após* dois ciclos, o registo `X` é incrementado pelo valor `V`. (`V` pode ser negativo.)

  - `noop` leva *um ciclo* para ser concluído. Não tem nenhum outro efeito.


A CPU usa essas instruções em um programa (sua entrada do puzzle) para, de alguma forma, dizer à tela o que desenhar.

Considere o seguinte pequeno programa:

```
noop
addx 3
addx -5

```

A execução deste programa procede da seguinte forma:


  - No início do primeiro ciclo, a instrução `noop` inicia a execução. Durante o primeiro ciclo, `X` é `1`. Após o primeiro ciclo, a instrução `noop` termina a execução, sem fazer nada.

  - No início do segundo ciclo, a instrução `addx 3` inicia a execução. Durante o segundo ciclo, `X` ainda é `1`.

  - Durante o terceiro ciclo, `X` ainda é `1`. Após o terceiro ciclo, a instrução `addx 3` finaliza a execução, configurando `X` para `4`.

  - No início do quarto ciclo, a instrução `addx -5` inicia a execução. Durante o quarto ciclo, `X` ainda é `4`.

  - Durante o quinto ciclo, `X` ainda é `4`. Após o quinto ciclo, a instrução `addx -5` finaliza a execução, configurando `X` para `-1`.


Talvez você possa aprender algo observando o valor do registo `X` durante a execução. Por enquanto, considere a *intensidade do sinal* (o número do ciclo multiplicado pelo valor do registro `X`) *durante* o 20º ciclo e a cada 40 ciclos depois disso (ou seja, durante o 20º, 60º, 100º, 140º, 180º e 220º ciclos).

Por exemplo, considere este programa maior:

```
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop

```

As intensidades de sinal interessantes podem ser determinadas da seguinte forma:


  - Durante o 20º ciclo, o registo `X` tem o valor `21`, então a força do sinal é 20 * 21 = *420*. (O 20º ciclo ocorre no meio do segundo `addx -1`, então o valor do registo `X` é o valor inicial, `1`, mais todos os outros valores `addx` até aquele ponto: 1 + 15 - 11 + 6 - 3 + 5 - 1 - 8 + 13 + 4 = 21.)

  - Durante o 60º ciclo, o registro `X` tem o valor `19`, então a intensidade do sinal é 60 * 19 = `1140`.

  - Durante o 100º ciclo, o registro `X` tem o valor `18`, então a força do sinal é 100 * 18 = `1800`.

  - Durante o 140º ciclo, o registro `X` tem o valor `21`, então a força do sinal é 140 * 21 = `2940`.

  - Durante o 180º ciclo, o registro `X` tem o valor `16`, então a força do sinal é 180 * 16 = `2880`.

  - Durante o 220º ciclo, o registro `X` tem o valor `18`, então a força do sinal é 220 * 18 = `3960`.


A soma dessas intensidades de sinal é `13140`.

Encontre a força do sinal durante o 20º, 60º, 100º, 140º, 180º e 220º ciclos. *Qual é a soma dessas seis intensidades de sinal?*

## --- Parte Dois ---

Parece que o registo `X` controla a posição horizontal de um [sprite](https://pt.wikipedia.org/wiki/Sprite_(computa%C3%A7%C3%A3o_gr%C3%A1fica)). Especificamente, o sprite tem 3 pixels de largura e o registro `X` define a posição horizontal do *meio* desse sprite. (Neste sistema, não existe "posição vertical": se a posição horizontal do sprite colocar seus pixels onde o CRT está desenhando no momento, esses pixels serão desenhados.)

Você conta os pixels no CRT: 40 de largura e 6 de altura. Essa tela CRT desenha a linha superior de pixels da esquerda para a direita, depois a linha abaixo dela e assim por diante. O pixel mais à esquerda em cada linha está na posição `0`, e o pixel mais à direita em cada linha está na posição `39`.

Tal como a CPU, o CRT está intimamente ligado ao circuito do relógio: o CRT desenha *um único pixel durante cada ciclo*. Representando cada pixel da tela como um `#`, aqui estão os ciclos durante os quais o primeiro e o último pixel de cada linha são desenhados:

```
Ciclo   1 -> ######################################## <- Ciclo  40
Ciclo  41 -> ######################################## <- Ciclo  80
Ciclo  81 -> ######################################## <- Ciclo 120
Ciclo 121 -> ######################################## <- Ciclo 160
Ciclo 161 -> ######################################## <- Ciclo 200
Ciclo 201 -> ######################################## <- Ciclo 240

```

Então, [*cronometrando*](https://www.youtube.com/watch?v=sJFnWZH5FXc) [cuidadosamente](https://en.wikipedia.org/wiki/Racing_the_Beam) as instruções da CPU e as operações de desenho do CRT, você deve ser capaz de determinar se o sprite é visível no instante em que cada pixel é desenhado. Se o sprite estiver posicionado de forma que um de seus três pixels seja o pixel que está sendo desenhado no momento, a tela produzirá um pixel *aceso* (`#`); caso contrário, a tela deixa o pixel *escuro* (`.`).
Os primeiros píxeis do exemplo maior acima são desenhados da seguinte forma:

```
Posição do Sprite: ###.....................................

Início ciclo    1: começa por executar addx 15
Durante ciclo   1: o CRT desenha o pixel na posição 0
Linha atual CRT  : #

Durante ciclo   2: o CRT desenha o pixel na posição 1
Linha atual CRT  : ##
Final do ciclo  2: termina executando addx 15 (Registo X passa a 16)
Posição do Sprite: ...............###......................

Início ciclo    3: começa por executar addx -11
Durante ciclo   3: o CRT desenha o pixel na posição 2
Linha atual CRT  : ##.

Durante ciclo   4: o CRT desenha o pixel na posição 3
Linha atual CRT  : ##..
Final do ciclo  4: termina executando addx -11 (Registo X passa a 5)
Posição do Sprite: ....###.................................

Início ciclo    5: começa por executar addx 6
Durante ciclo   5: o CRT desenha o pixel na posição 4
Linha atual CRT  : ##..#

Durante ciclo   6: o CRT desenha o pixel na posição 5
Linha atual CRT  : ##..##
Final do ciclo  6: termina executando addx 6 (Registo X passa a 11)
Posição do Sprite: ..........###...........................

Início ciclo    7: começa por executar addx -3
Durante ciclo   7: o CRT desenha o pixel na posição 6
Linha atual CRT  : ##..##.

Durante ciclo   8: o CRT desenha o pixel na posição 7
Linha atual CRT  : ##..##..
Final do ciclo  8: termina executando addx -3 (Registo X passa a 8)
Posição do Sprite: .......###..............................

Início ciclo    9: começa por executar addx 5
Durante ciclo   9: o CRT desenha o pixel na posição 8
Linha atual CRT  : ##..##..#

Durante ciclo  10: o CRT desenha o pixel na posição 9
Linha atual CRT  : ##..##..##
Final do ciclo 10: termina executando addx 5 (Registo X passa a 13)
Posição do Sprite: ............###.........................

Início ciclo   11: começa por executar addx -1
Durante ciclo  11: o CRT desenha o pixel na posição 10
Linha atual CRT  : ##..##..##.

Durante ciclo  12: o CRT desenha o pixel na posição 11
Linha atual CRT  : ##..##..##..
Final do ciclo 12: termina executando addx -1 (Registo X passa a 12)
Posição do Sprite: ...........###..........................

Início ciclo   13: começa por executar addx -8
Durante ciclo  13: o CRT desenha o pixel na posição 12
Linha atual CRT  : ##..##..##..#

Durante ciclo  14: o CRT desenha o pixel na posição 13
Linha atual CRT  : ##..##..##..##
Final do ciclo 14: termina executando addx -8 (Registo X passa a 4)
Posição do Sprite: ...###..................................

Início ciclo   15: começa por executar addx 13
Durante ciclo  15: o CRT desenha o pixel na posição 14
Linha atual CRT  : ##..##..##..##.

Durante ciclo  16: o CRT desenha o pixel na posição 15
Linha atual CRT  : ##..##..##..##..
Final do ciclo 16: termina executando addx 13 (Registo X passa a 17)
Posição do Sprite: ................###.....................

Início ciclo   17: começa por executar addx 4
Durante ciclo  17: o CRT desenha o pixel na posição 16
Linha atual CRT  : ##..##..##..##..#

Durante ciclo  18: o CRT desenha o pixel na posição 17
Linha atual CRT  : ##..##..##..##..##
Final do ciclo 18: termina executando addx 4 (Registo X passa a 21)
Posição do Sprite: ....................###.................

Início ciclo   19: começa por executar noop
Durante ciclo  19: o CRT desenha o pixel na posição 18
Linha atual CRT  : ##..##..##..##..##.
Final do ciclo 19: termina executando noop

Início ciclo   20: começa por executar addx -1
Durante ciclo  20: o CRT desenha o pixel na posição 19
Linha atual CRT  : ##..##..##..##..##..

Durante ciclo  21: o CRT desenha o pixel na posição 20
Linha atual CRT  : ##..##..##..##..##..#
Final do ciclo 21: termina executando addx -1 (Registo X passa a 20)
Posição do Sprite: ...................###..................

```

Permitir que o programa seja executado até ao final faz com que o CRT produza a seguinte imagem:

```
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....

```

Renderize a imagem fornecida pelo seu programa. *Quais são as oito letras maiúsculas que aparecem no seu CRT?*


