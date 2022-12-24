# --- Dia 24: Bacia de Nevasca ---

Com tudo replantado para o próximo ano (e com elefantes e macacos para cuidar do arvoredo), você e os Elfos partem para o ponto de extração.

No meio da montanha que protege o arvoredo há uma área plana e aberta que serve como ponto de extração. É um pouco difícil de subir, mas nada que a expedição não aguente.

Pelo menos, isso normalmente seria verdade; agora que a montanha está coberta de neve, as coisas se tornaram mais difíceis do que os Elfos estão acostumados.

À medida que a expedição chega a um vale que deve ser atravessado para chegar ao local de extração, você descobre que ventos fortes e turbulentos estão empurrando pequenas *nevascas* de neve e gelo afiado ao redor do vale. Ainda bem que todos levaram roupas quentes! Para atravessar com segurança, você precisará encontrar uma maneira de evitá-las.

Felizmente, é fácil ver tudo isso da entrada do vale, então você faz um mapa do vale e das nevascas (sua entrada do puzzle). Por exemplo:

```
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#

```

As paredes do vale são desenhadas como `#`; tudo o mais é chão. Terreno limpo - onde atualmente não há nevasca - é desenhado como `.`. Caso contrário, as nevascas são desenhadas com uma seta indicando a direção do movimento: para cima (`^`), para baixo (`v`), para a esquerda (`<`) ou para a direita (`>`).

O mapa acima inclui duas nevascas, uma se movendo para a direita (`>`) e outra se movendo para baixo (`v`). Em um minuto, cada nevasca se move uma posição na direção que aponta:

```
#.#####
#.....#
#.>...#
#.....#
#.....#
#...v.#
#####.#

```

Devido à *conservação da energia da nevasca*, quando uma nevasca atinge a parede do vale, uma nova nevasca se forma no lado oposto do vale, movendo-se na mesma direção. Depois de mais um minuto, a nevasca que se move para baixo foi substituída por uma nova nevasca que se move para baixo no topo do vale:

```
#.#####
#...v.#
#..>..#
#.....#
#.....#
#.....#
#####.#

```

Como as nevascas são feitas de minúsculos flocos de neve, elas passam umas pelas outras. Depois de mais um minuto, ambas as nevascas ocupam temporariamente a mesma posição, marcada como `2`:

```
#.#####
#.....#
#...2.#
#.....#
#.....#
#.....#
#####.#

```

Depois de mais um minuto, a situação se resolve, devolvendo a cada nevasca seu espaço pessoal:

```
#.#####
#.....#
#....>#
#...v.#
#.....#
#.....#
#####.#

```

Finalmente, depois de mais um minuto, a nevasca à direita voltada para a direita é substituída por uma nova à esquerda voltada para a mesma direção:

```
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#

```

Esse processo se repete pelo menos enquanto você o observa, mas provavelmente para sempre.

Aqui está um exemplo mais complexo:

```
#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#

```

Sua expedição começa na única posição fora da parede na linha superior e precisa alcançar a única posição fora da parede na linha inferior. A cada minuto, você pode *mover* para cima, baixo, esquerda ou direita, ou pode *esperar* no lugar. Você e as nevascas agem *simultaneamente* e você não pode compartilhar uma posição com uma nevasca.

No exemplo acima, a maneira mais rápida de alcançar seu objetivo requer `18` passos. Desenhando a posição da expedição como `E`, uma maneira de conseguir isso é:

```
Estado inicial:
#E######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#

Minuto 1, mover para baixo:
#.######
#E>3.<.#
#<..<<.#
#>2.22.#
#>v..^<#
######.#

Minuto 2, mover para baixo:
#.######
#.2>2..#
#E^22^<#
#.>2.^>#
#.>..<.#
######.#

Minuto 3, aguardar:
#.######
#<^<22.#
#E2<.2.#
#><2>..#
#..><..#
######.#

Minuto 4, mover para cima:
#.######
#E<..22#
#<<.<..#
#<2.>>.#
#.^22^.#
######.#

Minuto 5, mover para a direita:
#.######
#2Ev.<>#
#<.<..<#
#.^>^22#
#.2..2.#
######.#

Minuto 6, mover para a direita:
#.######
#>2E<.<#
#.2v^2<#
#>..>2>#
#<....>#
######.#

Minuto 7, mover para baixo:
#.######
#.22^2.#
#<vE<2.#
#>>v<>.#
#>....<#
######.#

Minuto 8, mover para a esquerda:
#.######
#.<>2^.#
#.E<<.<#
#.22..>#
#.2v^2.#
######.#

Minuto 9, mover para cima:
#.######
#<E2>>.#
#.<<.<.#
#>2>2^.#
#.v><^.#
######.#

Minuto 10, mover para a direita:
#.######
#.2E.>2#
#<2v2^.#
#<>.>2.#
#..<>..#
######.#

Minuto 11, aguardar:
#.######
#2^E^2>#
#<v<.^<#
#..2.>2#
#.<..>.#
######.#

Minuto 12, mover para baixo:
#.######
#>>.<^<#
#.<E.<<#
#>v.><>#
#<^v^^>#
######.#

Minuto 13, mover para baixo:
#.######
#.>3.<.#
#<..<<.#
#>2E22.#
#>v..^<#
######.#

Minuto 14, mover para a direita:
#.######
#.2>2..#
#.^22^<#
#.>2E^>#
#.>..<.#
######.#

Minuto 15, mover para a direita:
#.######
#<^<22.#
#.2<.2.#
#><2>E.#
#..><..#
######.#

Minuto 16, mover para a direita:
#.######
#.<..22#
#<<.<..#
#<2.>>E#
#.^22^.#
######.#

Minuto 17, mover para baixo:
#.######
#2.v.<>#
#<.<..<#
#.^>^22#
#.2..2E#
######.#

Minuto 18, mover para baixo:
#.######
#>2.<.<#
#.2v^2<#
#>..>2>#
#<....>#
######E#

```

*Qual é o menor número de minutos necessários para evitar as nevascas e atingir a meta?*

