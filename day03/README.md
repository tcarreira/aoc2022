# --- Dia 3: Reorganização do Mochilão ---

Um Elfo tem a importante tarefa de carregar todos os [mochilões](https://pt.wikipedia.org/wiki/Mochila) com suprimentos para a viagem pela *selva*. Infelizmente, esse Elfo não seguiu exatamente as instruções de embalagem, e alguns itens precisam ser rearranjados.

Cada mochilão tem dois grandes *compartimentos*. Todos os itens de um determinado tipo devem ir para exatamente um dos dois compartimentos. O Elfo que fez o empacotamento não seguiu essa regra para exatamente um tipo de item por mochilão.

Os Elfos fizeram uma lista de todos os itens atualmente em cada mochilão (seu input do puzzle), mas precisam da sua ajuda para encontrar os erros. Cada tipo de item é identificado por uma única letra minúscula ou maiúscula (ou seja, `a` e `A` se referem a tipos de itens diferentes).

A lista de itens para cada mochilão é dada como caracteres em uma única linha. Dado um mochilão, sempre há o mesmo número de itens em cada um de seus dois compartimentos, então a primeira metade dos caracteres representa os itens no primeiro compartimento, enquanto a segunda metade dos caracteres representa os itens no segundo compartimento.

Por exemplo, suponha que você tenha a seguinte lista de conteúdo de seis mochilões:

```
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

```


 - O primeiro mochilão contém os items `vJrwpWtwJgWrhcsFMMfFFhFp`, o que significa que o primeiro compartimento contém os items `vJrwpWtwJgWr`, enquanto o segundo compartimento contém os items `hcsFMMfFFhFp`. O único tipo que aparece em ambos compartimentos é o `p` minúsculo.

 - O segundo mochilão contem os items `jqHRNqRjqzjGDLGL` e `rsFMfFZSrLrFZsSL`. O único tipo que aparece em ambos compartimentos é o `L` maiúsculo.

 - O terceiro mochilão contém `PmmdzqPrV` e `vPwwTWBwg`; o único tipo que aparece em ambos compartimentos é o `P` maiúsculo.

 - Os compartimentos do quarto mochilão compartilham apenas o tipo `v`.

 - Os compartimentos do quinto mochilão compartilham apenas o tipo `t`.

 - os compartimentos do sexto mochilão compartilham apenas o tipo `s`.


Para ajudar à reorganização de items, cada tipo de items precisa ser convertido em uma *prioridade*.


 - Os tipos de item minúsculos de `a` a `z` têm prioridades 1 até 26.

 - Os tipos de item maiúsculos de `A` a `Z` têm prioridades 27 até 52.


No exemplo acima, a prioridade dos items que aparece em ambos os compartimentos de cada mochilão é 16 (`p`), 38 (`L`), 42 (`P`), 22 (`v`), 20 (`t`), e 19 (`s`); a soma delas é `157`.

Encontra o tipo de items que aparece em ambos os compartimentos de cada mochilão. *Qual a soma das prioridades desses tipos de items?*

