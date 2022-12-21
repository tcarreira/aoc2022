# --- Dia 21: Matemática de Macaco ---

Os [macacos](../day11) estão de volta! Você está preocupado que eles tentem roubar suas coisas novamente, mas parece que eles estão apenas mantendo a posição e fazendo vários barulhos de macaco para você.

Eventualmente, um dos elefantes percebe que você não fala macaquês e vem interpretar. Acontece que eles ouviram você falando sobre tentar encontrar o arvoredo; eles podem mostrar um atalho se você responder a sua *enigma*.

Cada macaco recebe uma *tarefa*: ou *gritar um número específico* ou *gritar o resultado de uma operação matemática*. Todos os macacos que gritam números sabem seu número desde o início; no entanto, os macacos da operação matemática precisam esperar que dois outros macacos gritem um número, e esses dois outros macacos podem *também* estar esperando por outros macacos.

Sua tarefa é *calcular o número que o macaco chamado `root` gritará* antes que os macacos descubram sozinhos.

Por exemplo:

```
root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32

```

Cada linha contém o nome de um macaco, dois pontos e, em seguida, a tarefa desse macaco:


  - Um número solitário significa que a tarefa do macaco é simplesmente gritar aquele número.

  - Uma tarefa como `aaaa + bbbb` significa que o macaco espera que os macacos `aaaa` e `bbbb` gritem cada um de seus números; o macaco então grita a soma desses dois números.

  - `aaaa - bbbb` significa que o macaco grita o número de `aaaa` menos o número de `bbbb`.

  - A tarefa `aaaa * bbbb` significa que o macaco grita o número de `aaaa` multiplicado pelo número de `bbbb`.

  - A tarefa `aaaa / bbbb` significa que o macaco grita o número de `aaaa` dividido pelo número de `bbbb`.


Então, no exemplo acima, o macaco `drzm` tem que esperar que os macacos `hmdt` e `zczc` gritem seus números. Felizmente, tanto `hmdt` quanto `zczc` têm tarefas que envolvem simplesmente gritar um único número, então eles fazem isso imediatamente: `32` e `2`. O macaco `drzm` pode então gritar seu número encontrando `32` menos `2`: `30`.

Depois, o macaco `sjmn` tem um de seus números (`30`, do macaco `drzm`), e também já tem seu outro número, `5`, de `dbpl`. Isso permite que ele grite seu próprio número encontrando `30` multiplicado por `5`: `150`.

Este processo continua até que `root` grite um número: `152`.

No entanto, sua situação real envolve *consideravelmente mais macacos*. *Que número o macaco chamado `root` gritará?*

