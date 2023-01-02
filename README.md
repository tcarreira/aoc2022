# Advent of Code 2022

Traduções para Português e implementação das soluções do AoC2022

![Advent Of Code](./aoc.jpeg)

## Descrição

Soluções do https://adventofcode.com/2022 feitas em Go.
As traduções dos enunciados estão disponíveis dentro de cada dia.

As soluções são desenvolvidas ao vivo e transmitidas no
[youtube](https://youtube.com/@tcarreira) e no
[twitch](https://twitch.tv/tcarreira).

Existem playlists para seguir no Youtube:
- Gravações do streaming das resoluções ao vivo: https://youtube.com/playlist?list=PLJgtTDTisWI4Vtuc8pjCLcPUqaUgvzXVn
- Resumos curtos das soluções: https://youtube.com/playlist?list=PLJgtTDTisWI5K05p0SMu9XaIUYC6JzEPQ

Qualquer um pode entrar no [LeaderBoard privado](https://adventofcode.com/2022/leaderboard/private) usando o código: `827447-e95d42f1` (até um limite de 200)


## Soluções e tempos

`go run .`

<!-- ci:result:start -->
```
######################################################
Resolvendo os puzzles do https://adventofcode.com/2022
######################################################


*****************     Soluções     *****************
 Dia | Parte 1              | Parte 2              |
-----+----------------------+----------------------+
  01 | 70116                | 206582               | primeiro problema
  02 | 9759                 | 12429                | seguir instruções: pedra-papel-tesoura
  03 | 7727                 | 2609                 | interseção de mapas (maps/dicts)
  04 | 485                  | 857                  | sobreposição de intervalos
  05 | WSFTMRHPP            | GSLCMFBRP            | pilhas (stacks)
  06 | 1855                 | 3256                 | substrings (uniq chars)
  07 | 1118405              | 12545514             | árvores sistema arquivos (fs trees)
  08 | 1779                 | 172224               | programação dinâmica
  09 | 6391                 | 2593                 | map/dict/set coord visitadas
  10 | 16480                | PLEFULPB             | desenhando caracteres (env: AOC_DAY10)
  11 | 51075                | 11741456163          | big integers / overflow
  12 | 520                  | 508                  | Caminho mais curto
  13 | 5013                 | 25038                | ler e ordenar campos não-estruturados
  14 | 1133                 | 27566                | simul ampulheta: areia que cai
  15 | 6275922              | 11747175442119       | ponto fora de n-áreas (via perímetro)
  16 | 2330                 | 2675                 | * teoria de grafos, dfs
  17 | 3133                 | 1547953216393        | * tetris + cache de estado
  18 | 4390                 | 2534                 | labirinto em mundo minecraft
  19 | 1616                 | 8990                 | ** recursos, prune, optimizs, no-cache
  20 | 5498                 | 3390007892081        | Listas (duplamente) ligadas
  21 | 21120928600114       | 3453748220116        | árvore binária & backtrack
  22 | 27436                | 15426                | caminhar na face do cubo (planif hardc)
  23 | 4162                 | 986                  | jogo da vida
  24 | 281                  | 807                  | mapa 2d no tempo, bfs
  25 | 2011-=2=-1020-1===-1 | -                    | converter número base estranha


*****************      Tempos      *****************
 Dia | Parte 1              | Parte 2              |
-----+----------------------+----------------------+
  01 |              0.16 ms |              0.18 ms | primeiro problema
  02 |              0.13 ms |              0.12 ms | seguir instruções: pedra-papel-tesoura
  03 |              0.06 ms |              0.97 ms | interseção de mapas (maps/dicts)
  04 |              0.21 ms |              0.20 ms | sobreposição de intervalos
  05 |              0.10 ms |              0.10 ms | pilhas (stacks)
  06 |              0.01 ms |              0.05 ms | substrings (uniq chars)
  07 |              0.17 ms |              0.16 ms | árvores sistema arquivos (fs trees)
  08 |              2.89 ms |              0.66 ms | programação dinâmica
  09 |              1.82 ms |              1.08 ms | map/dict/set coord visitadas
  10 |              0.03 ms |              0.03 ms | desenhando caracteres (env: AOC_DAY10)
  11 |              0.12 ms |             17.29 ms | big integers / overflow
  12 |              0.20 ms |              0.39 ms | Caminho mais curto
  13 |              1.82 ms |              2.90 ms | ler e ordenar campos não-estruturados
  14 |              1.16 ms |              5.40 ms | simul ampulheta: areia que cai
  15 |           1094.00 ms |            233.10 ms | ponto fora de n-áreas (via perímetro)
  16 |            117.58 ms |           6561.80 ms | * teoria de grafos, dfs
  17 |              0.49 ms |              6.97 ms | * tetris + cache de estado
  18 |              2.40 ms |              2.46 ms | labirinto em mundo minecraft
  19 |            641.53 ms |          11370.56 ms | ** recursos, prune, optimizs, no-cache
  20 |            257.85 ms |           1519.47 ms | Listas (duplamente) ligadas
  21 |              0.63 ms |              0.62 ms | árvore binária & backtrack
  22 |              1.77 ms |              2.01 ms | caminhar na face do cubo (planif hardc)
  23 |             15.52 ms |           1106.77 ms | jogo da vida
  24 |             28.69 ms |             93.58 ms | mapa 2d no tempo, bfs
  25 |              0.03 ms |              0.00 ms | converter número base estranha


**************     Memória Alocada    **************
 Dia | Parte 1              | Parte 2              |
-----+----------------------+----------------------+
  01 |                99 KB |               107 KB | primeiro problema
  02 |               295 KB |               295 KB | seguir instruções: pedra-papel-tesoura
  03 |                54 KB |               770 KB | interseção de mapas (maps/dicts)
  04 |               293 KB |               294 KB | sobreposição de intervalos
  05 |                94 KB |                93 KB | pilhas (stacks)
  06 |                 1 KB |                 1 KB | substrings (uniq chars)
  07 |               123 KB |               125 KB | árvores sistema arquivos (fs trees)
  08 |               274 KB |               244 KB | programação dinâmica
  09 |               570 KB |               390 KB | map/dict/set coord visitadas
  10 |                37 KB |                36 KB | desenhando caracteres (env: AOC_DAY10)
  11 |                35 KB |             13623 KB | big integers / overflow
  12 |               279 KB |               202 KB | Caminho mais curto
  13 |               635 KB |              1474 KB | ler e ordenar campos não-estruturados
  14 |              1138 KB |              1139 KB | simul ampulheta: areia que cai
  15 |            228668 KB |                 7 KB | ponto fora de n-áreas (via perímetro)
  16 |             36953 KB |             15555 KB | * teoria de grafos, dfs
  17 |                80 KB |              3921 KB | * tetris + cache de estado
  18 |               763 KB |               508 KB | labirinto em mundo minecraft
  19 |                15 KB |                14 KB | ** recursos, prune, optimizs, no-cache
  20 |               715 KB |               715 KB | Listas (duplamente) ligadas
  21 |               674 KB |               673 KB | árvore binária & backtrack
  22 |              2408 KB |              2523 KB | caminhar na face do cubo (planif hardc)
  23 |              4796 KB |            381151 KB | jogo da vida
  24 |             19711 KB |             60433 KB | mapa 2d no tempo, bfs
  25 |                 2 KB |                 0 KB | converter número base estranha

```
<!-- ci:result:end -->
