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

 Dia | Parte 1    ( time ms ) | Parte 2        ( time ms ) |
-----+------------------------+----------------------------+
  01 | 70116      (      0.1) | 206582         (      0.2) | primeiro problema
  02 | 9759       (      0.2) | 12429          (      0.2) | seguir instruções: pedra-papel-tesoura
  03 | 7727       (      0.1) | 2609           (      1.5) | interseção de mapas (maps/dicts)
  04 | 485        (      0.4) | 857            (      0.4) | sobreposição de intervalos
  05 | WSFTMRHPP  (      0.1) | GSLCMFBRP      (      0.1) | pilhas (stacks)
  06 | 1855       (      0.0) | 3256           (      0.1) | substrings (uniq chars)
  07 | 1118405    (      0.2) | 12545514       (      0.2) | árvores sistema arquivos (fs trees)
  08 | 1779       (      0.4) | 172224         (      0.7) | programação dinâmica
  09 | 6391       (      1.3) | 2593           (      1.6) | map/dict/set coord visitadas
  10 | 16480      (      0.0) | PLEFULPB       (      0.0) | desenhando caracteres (env: AOC_DAY10)
  11 | 51075      (      0.1) | 11741456163    (     35.3) | big integers / overflow
  12 | 520        (      0.2) | 508            (      0.1) | Caminho mais curto
  13 | 5013       (      2.1) | 25038          (      3.9) | ler e ordenar campos não-estruturados
  14 | 1133       (      1.4) | 27566          (      6.8) | simul ampulheta: areia que cai
  15 | 6275922    (   1671.9) | 11747175442119 (    493.8) | ponto fora de n-áreas (via perímetro)
  16 | 2330       (    154.4) | 2675           (  10133.5) | teoria de grafos, dfs
  17 | 3133       (      0.7) | 1547953216393  (     11.7) | tetris + cache de estado
  18 | 4390       (      4.7) | 2534           (      4.3) | labirinto em mundo minecraft
  19 | -          (      0.0) | -              (      0.0) | 
```
<!-- ci:result:end -->
