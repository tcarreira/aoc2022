# --- Dia 19: Não há minerais suficientes ---

Suas varreduras mostram que a lava realmente formou obsidiana!

O vento mudou de direção o suficiente para parar de enviar gotículas de lava em sua direção, então você e os elefantes saem da caverna. Ao fazer isso, você repara em uma coleção de [geodes](https://pt.wikipedia.org/wiki/Geode) ao redor da lagoa. Talvez você possa usar a obsidiana para criar alguns *robôs quebra-geodes* e quebrá-los?

Para coletar a obsidiana do fundo do lago, você precisará de *robôs coletores de obsidiana* à prova d'água. Felizmente, há uma quantidade abundante de argila por perto que você pode usar para torná-los impermeáveis.

Para colher a argila, você precisará de especializados *robôs coletores de argila*. Para fazer qualquer tipo de robô, você precisará de *minério*, que também é abundante, mas na direção oposta da argila.

A coleta de minério requer *robôs coletores de minério* com brocas grandes. Felizmente, *você tem exatamente um robô coletor de minério* em sua mochila que pode usar para *pôr em marcha* toda a operação.

Cada robô pode coletar 1 de seu tipo de recurso por minuto. Também leva um minuto para a fábrica de robôs (também convenientemente da sua mochila) construir qualquer tipo de robô, embora consuma os recursos necessários disponíveis quando a construção começa.

A fábrica de robôs tem muitos *modelos* (sua entrada do puzzle) que você pode escolher, mas depois de configurá-la com um projeto, não poderá alterá-lo. Você precisará descobrir qual modelo é o melhor.

Por exemplo:

```
Blueprint 1:
  Each ore robot costs 4 ore.
  Each clay robot costs 2 ore.
  Each obsidian robot costs 3 ore and 14 clay.
  Each geode robot costs 2 ore and 7 obsidian.

Blueprint 2:
  Each ore robot costs 2 ore.
  Each clay robot costs 3 ore.
  Each obsidian robot costs 3 ore and 8 clay.
  Each geode robot costs 3 ore and 12 obsidian.

```

(Os modelos foram quebrados em várias linhas para legibilidade. A lista real de modelos da fábrica de robôs é fornecida com um modelo por linha.)

Os elefantes estão começando a parecer famintos, então você não deve demorar muito; você precisa descobrir qual projeto maximizaria o número de geodes abertos após *24 minutos*, descobrindo quais robôs construir e quando construí-los.

Usando o modelo 1 no exemplo acima, o maior número de geodes que você pode abrir em 24 minutos é `9`. Uma maneira de conseguir isso é:

```
== Minuto 1 ==
1 robô coletor de minério coleta 1 minério; agora você tem 1 minério.

== Minuto 2 ==
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.

== Minuto 3 ==
Gaste 2 minérios para começar a construir um robô coletor de argila.
1 robô coletor de minério coleta 1 minério; agora você tem 1 minério.
O novo robô coletor de argila está pronto; agora você tem 1 deles.

== Minuto 4 ==
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.
1 robô coletor de argila coleta 1 argila; agora você tem 1 argila.

== Minuto 5 ==
Gaste 2 minérios para começar a construir um robô coletor de argila.
1 robô coletor de minério coleta 1 minério; agora você tem 1 minério.
1 robô coletor de argila coleta 1 argila; agora você tem 2 argilas.
O novo robô coletor de argila está pronto; agora você tem 2 deles.

== Minuto 6 ==
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.
2 robôs coletores de argila coletam 2 argilas; agora você tem 4 argila.

== Minuto 7 ==
Gaste 2 minérios para começar a construir um robô coletor de argila.
1 robô coletor de minério coleta 1 minério; agora você tem 1 minério.
2 robôs coletores de argila coletam 2 argilas; agora você tem 6 argila.
O novo robô coletor de argila está pronto; agora você tem 3 deles.

== Minuto 8 ==
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.
3 robôs coletores de argila coletam 3 argilas; agora você tem 9 argila.

== Minuto 9 ==
1 robô coletor de minério coleta 1 minério; agora você tem 3 minérios.
3 robôs coletores de argila coletam 3 argilas; agora você tem 12 argilas.

== Minuto 10 ==
1 robô coletor de minério coleta 1 minério; agora você tem 4 minérios.
3 robôs coletores de argila coletam 3 argilas; agora você tem 15 argilas.

== Minuto 11 ==
Gaste 3 minérios e 14 argilas para começar a construir um robô coletor de obsidiana.
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.
3 robôs coletores de argila coletam 3 argilas; agora você tem 4 argila.
O novo robô coletor de obsidiana está pronto; agora você tem 1 deles.

== Minuto 12 ==
Gaste 2 minérios para começar a construir um robô coletor de argila.
1 robô coletor de minério coleta 1 minério; agora você tem 1 minério.
3 robôs coletores de argila coletam 3 argilas; agora você tem 7 argila.
1 robô coletor de obsidiana coleta 1 obsidiana; agora você tem 1 obsidiana.
O novo robô coletor de argila está pronto; agora você tem 4 deles.

== Minuto 13 ==
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 11 argila.
1 robô coletor de obsidiana coleta 1 obsidiana; agora você tem 2 obsidianas.

== Minuto 14 ==
1 robô coletor de minério coleta 1 minério; agora você tem 3 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 15 argilas.
1 robô coletor de obsidiana coleta 1 obsidiana; agora você tem 3 obsidianas.

== Minuto 15 ==
Gaste 3 minérios e 14 argilas para começar a construir um robô coletor de obsidiana.
1 robô coletor de minério coleta 1 minério; agora você tem 1 minério.
4 robôs coletores de argila coletam 4 argilas; agora você tem 5 argila.
1 robô coletor de obsidiana coleta 1 obsidiana; agora você tem 4 obsidianas.
O novo robô coletor de obsidiana está pronto; agora você tem 2 deles.

== Minuto 16 ==
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 9 argila.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 6 obsidianas.

== Minuto 17 ==
1 robô coletor de minério coleta 1 minério; agora você tem 3 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 13 argilas.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 8 obsidianas.

== Minuto 18 ==
Gaste 2 minérios e 7 obsidianas para começar a construir um robô quebra-geodes.
1 robô coletor de minério coleta 1 minério; agora você tem 2 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 17 argilas.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 3 obsidianas.
O novo robô quebra-geoges está pronto; agora você tem 1 deles.

== Minuto 19 ==
1 robô coletor de minério coleta 1 minério; agora você tem 3 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 21 argilas.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 5 obsidianas.
1 robô quebra-geodes quebra 1 geode; agora você tem 1 geode aberto.

== Minuto 20 ==
1 robô coletor de minério coleta 1 minério; agora você tem 4 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 25 argilas.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 7 obsidianas.
1 robô quebra-geodes quebra 1 geode; agora você tem 2 geodes abertos.

== Minuto 21 ==
Gaste 2 minérios e 7 obsidianas para começar a construir um robô quebra-geodes.
1 robô coletor de minério coleta 1 minério; agora você tem 3 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 29 argilas.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 2 obsidianas.
1 robô quebra-geodes quebra 1 geode; agora você tem 3 geodes abertos.
O novo robô quebra-geoges está pronto; agora você tem 2 deles.

== Minuto 22 ==
1 robô coletor de minério coleta 1 minério; agora você tem 4 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 33 argilas.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 4 obsidianas.
2 robôs quebra-geodes quebram 2 geodes; agora você tem 5 geodes abertos.

== Minuto 23 ==
1 robô coletor de minério coleta 1 minério; agora você tem 5 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 37 argilas.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 6 obsidianas.
2 robôs quebra-geodes quebram 2 geodes; agora você tem 7 geodes abertos.

== Minuto 24 ==
1 robô coletor de minério coleta 1 minério; agora você tem 6 minérios.
4 robôs coletores de argila coletam 4 argilas; agora você tem 41 argila.
2 robôs coletores de obsidiana coletam 2 obsidianas; agora você tem 8 obsidianas.
2 robôs quebra-geodes quebram 2 geodes; agora você tem 9 geodes abertos.

```

No entanto, usando o modelo 2 no exemplo acima, você pode fazer ainda melhor: o maior número de geodes que você pode abrir em 24 minutos é `12`.

Determine o *nível de qualidade* de cada modelo *multiplicando o número de ID desse modelo* pelo maior número de geodes que podem ser abertos em 24 minutos usando esse modelo. Neste exemplo, o primeiro modelo tem ID 1 e pode abrir 9 geodes, então seu nível de qualidade é `9`. O segundo modelo tem ID 2 e pode abrir 12 geodes, então seu nível de qualidade é `24`. Por fim, se você *somar os níveis de qualidade* de todos os modelos na lista, obtém `33`.

Determine o nível de qualidade de cada modelo usando o maior número de geodes que ele poderia produzir em 24 minutos. *O que você obtém se somar o nível de qualidade de todos os modelos da sua lista?*

