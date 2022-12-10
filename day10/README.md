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

