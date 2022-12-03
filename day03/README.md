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

# --- Parte Dois ---

Assim que terminas de identificar os items fora do lugar, os Elfos se aproximam com um outro problema.

Por segurança, os Elfos estão divididos em grupos de três. Cada Elfo leva um crachá que identifica o seu grupo. Para ser eficiente, dentro de cada grupo de três Elfos, o crachá é o *único item carregado por todos os três Elfos*. Ou seja, se um crachá de um grupo é o item do tipo `B`, então todos os Elfos do grupo devem carregar um item do tipo `B` dentro do seu mochilão, e no máximo dois dos Elfos irão lear qualquer outro tipo.

O problema é que alguém se esqueceu de atualizar o selo de autenticidade dos crachás deste ano. Todos os crachás precisam ser tirados para fora dos mochilões para que os novos selos sejam colados.

Adicionalmente, ninguém anotou quais tipos de item correspondem ao crachá de cada grupo. A única forma de dizer que o tipo de item é o certo é encontrando o único tipo de item que é *comum entre todos os três Elfos* de cada grupo.

Cada conjunto de três linhas na lista corresponde a um único grupo, mas cada grupo pode ter um tipo de item de crachá diferente. Então no exemplo acima, o primeiro grupo de mochilões são as primeiras três linhas:

```
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg

```

E o segundo grupo de mochilões são as três linhas seguintes:

```
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

```

No primeiro grupo, o único tipo de item que aparece em todos os três mochilões é o `r` minúsculo; este deve ser os seus crachás. No segundo grupo, o tipo do crachá tem que ser o `Z`.

As prioridades para esses items ainda precisam ser encontradas para organizar os esforços de colagem: aqui, eles são 18 (`r`) para o primeiro grupo e 52 (`Z`) para o segundo grupo. A soma das prioridades é *70*.

Encontra o tipo de item que corresponde aos crachás de cada grupo de três Elfos. *Qual a soma das prioridades desses tipos de items?*

