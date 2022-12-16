# --- Dia 16: Proboscidea Volcanium ---

Os sensores levaram você à origem do sinal de socorro: outro dispositivo portátil, igual ao que os Elfos lhe deram. No entanto, você não vê nenhum elfo por perto; em vez disso, o dispositivo está cercado por elefantes! Eles devem ter se perdido nesses túneis, e um dos elefantes aparentemente descobriu como ativar o sinal de socorro.

O chão treme novamente, muito mais forte desta vez. Que tipo de caverna é essa, exatamente? Você faz um varrimento da caverna com seu dispositivo portátil; o relatório mostra principalmente rocha ígnea, algumas cinzas, bolsas de gás pressurizado, magma... isto não é apenas uma caverna, é um vulcão!

Você precisa tirar os elefantes daqui, rápido. Seu dispositivo estima que você tem *30 minutos* antes que o vulcão entre em erupção, então você não tem tempo de voltar por onde entrou.

Você examina a caverna em busca de outras opções e descobre uma rede de canos e *válvulas* de alívio de pressão. Você não tem certeza de como tal sistema entrou dentro de um vulcão, mas não tem tempo para reclamar; seu dispositivo produz um relatório (sua entrada do puzzle) da *vazão* de cada válvula se ela fosse aberta (em pressão por minuto) e os túneis que você poderia usar para se mover entre as válvulas.

Há até uma válvula na sala em que você e os elefantes estão atualmente identificada como `AA`. Você estima que levará um minuto para abrir uma única válvula e um minuto para seguir qualquer túnel de uma válvula para outra. Qual é a maior pressão que você poderia liberar?

Por exemplo, suponha que você tenha o seguinte resultado do varrimento:

```
Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II

```

Todas as válvulas começam *fechadas*. Você começa na válvula `AA`, mas ela deve estar danificada ou *emperrada* ou algo assim: a vazão dela é `0`, então não adianta abrir. No entanto, você pode gastar um minuto movendo-se para a válvula `BB` e outro minuto abrindo-a; isso liberaria pressão durante os *28 minutos* restantes a uma vazão de `13`, uma liberação de pressão total eventual de `28 * 13 = 364`. Depois, você pode gastar seu terceiro minuto movendo-se para a válvula `CC` e seu quarto minuto abrindo-a, proporcionando *26 minutos* adicionais de eventual liberação de pressão a uma vazão de `2`, ou `52` de pressão total liberada pela válvula `CC`.

Fazendo o seu caminho através de túneis como este, você provavelmente poderia abrir muitas ou todas as válvulas no tempo de 30 minutos. No entanto, você precisa liberar o máximo de pressão possível, portanto, precisará ser metódico. Em vez disso, considere esta abordagem:

```
== Minuto 1 ==
Nenhuma válvula está aberta.
Você se move para a válvula DD.

== Minuto 2 ==
Nenhuma válvula está aberta.
Você abre a válvula DD.

== Minuto 3 ==
A válvula DD está aberta, liberando 20 de pressão.
Você se move para a válvula CC.

== Minuto 4 ==
A válvula DD está aberta, liberando 20 de pressão.
Você se move para a válvula BB.

== Minuto 5 ==
A válvula DD está aberta, liberando 20 de pressão.
Você abre a válvula BB.

== Minuto 6 ==
As válvulas BB e DD estão abertas, liberando 33 de pressão.
Você se move para a válvula AA.

== Minuto 7 ==
As válvulas BB e DD estão abertas, liberando 33 de pressão.
Você se move para a válvula II.

== Minuto 8 ==
As válvulas BB e DD estão abertas, liberando 33 de pressão.
Você se move para a válvula JJ.

== Minuto 9 ==
As válvulas BB e DD estão abertas, liberando 33 de pressão.
Você abre a válvula JJ.

== Minuto 10 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você se move para a válvula II.

== Minuto 11 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você se move para a válvula AA.

== Minuto 12 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você se move para a válvula DD.

== Minuto 13 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você se move para a válvula EE.

== Minuto 14 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você move para a válvula FF.

== Minuto 15 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você se move para a válvula GG.

== Minuto 16 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você se move para a válvula HH.

== Minuto 17 ==
As válvulas BB, DD e JJ estão abertas, liberando 54 de pressão.
Você abre a válvula HH.

== Minuto 18 ==
As válvulas BB, DD, HH e JJ estão abertas, liberando 76 de pressão.
Você se move para a válvula GG.

== Minuto 19 ==
As válvulas BB, DD, HH e JJ estão abertas, liberando 76 de pressão.
Você move para a válvula FF.

== Minuto 20 ==
As válvulas BB, DD, HH e JJ estão abertas, liberando 76 de pressão.
Você se move para a válvula EE.

== Minuto 21 ==
As válvulas BB, DD, HH e JJ estão abertas, liberando 76 de pressão.
Você abre a válvula EE.

== Minuto 22 ==
As válvulas BB, DD, EE, HH e JJ estão abertas, liberando 79 de pressão.
Você se move para a válvula DD.

== Minuto 23 ==
As válvulas BB, DD, EE, HH e JJ estão abertas, liberando 79 prescerto.
Você se move para a válvula CC.

== Minuto 24 ==
As válvulas BB, DD, EE, HH e JJ estão abertas, liberando 79 de pressão.
Você abre a válvula CC.

== Minuto 25 ==
As válvulas BB, CC, DD, EE, HH e JJ estão abertas, liberando 81 de pressão.

== Minuto 26 ==
As válvulas BB, CC, DD, EE, HH e JJ estão abertas, liberando 81 de pressão.

== Minuto 27 ==
As válvulas BB, CC, DD, EE, HH e JJ estão abertas, liberando 81 de pressão.

== Minuto 28 ==
As válvulas BB, CC, DD, EE, HH e JJ estão abertas, liberando 81 de pressão.

== Minuto 29 ==
As válvulas BB, CC, DD, EE, HH e JJ estão abertas, liberando 81 de pressão.

== Minuto 30 ==
As válvulas BB, CC, DD, EE, HH e JJ estão abertas, liberando 81 de pressão.

```

Essa abordagem permite que você libere o máximo de pressão possível em 30 minutos com esta disposição de válvulas, `1651`.

Descubra os passos para liberar o máximo de pressão em 30 minutos. *Qual é a maior pressão que você consegue liberar?*

