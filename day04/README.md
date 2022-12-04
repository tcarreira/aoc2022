# --- Dia 4: Limpeza do Campo ---

O espaço precisa ser limpo antes que os últimos suprimentos possam ser descarregados dos navios, e por isso vários Elfos foram atribuídos para a tarefa de limpar secções do campo. Cada secção tem um *ID único*, e cada Elfo é atribuído a um intervalo de IDs de secção.

Contudo, como alguns dos Elfos comparam entre eles as secções que lhes foram atribuídas, eles começaram a notar que muitas das secções se sobrepunham. Para tentar rapidamente encontrar as sobreposições e reduzir o esforço duplicado, os Elfos juntam-se em pares e fazem uma *grande lista de atribuições de secções para cada par* (o input do puzzle).

Por exemplo, considere a seguinte lista de pares de atribuições de secções:

```
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8

```

Para os primeiros pares, isto significa:


 - Dentro do primeiro par de Elfos, o primeiro Elfo foi atribuído às secções `2-4` (secções `2`, `3`, e `4`),enquanto o seguindo Elfo foi atribuído às secções `6-8` (secções `6`, `7`, `8`).

 - Os Elfos no segundo par foram atribuídos a cada um duas secções.

 - Os Elfos no terceiro par foram atribuídos a três secções cada: um ficou com as secções `5`, `6`, e `7`, enquanto o outro também ficou com a `7`, mais a `8` e a `9`.


Esta lista de exemplo usa IDs de secções com apenas um dígito para ser mais fácil desenhar; o seu input real poderá conter números maiores. Visualmente, estes pares de atribuições de secções se parecem com o seguinte:

```
.234.....  2-4
.....678.  6-8

.23......  2-3
...45....  4-5

....567..  5-7
......789  7-9

.2345678.  2-8
..34567..  3-7

.....6...  6-6
...456...  4-6

.23456...  2-6
...45678.  4-8

```

Alguns dos pares notaram que uma das suas atribuições *contém completamente* a outra. Por exemplo, `2-8` contém completamente `3-7`, e `6-6` está completamente contido pelo `4-6`. Em pares onde uma atribuição contém completamente a outra, um Elfo no par iria limpar apenas secções que o seu parceiro já teria limpo, então estes parecem os casos que mais precisam de ser reconsiderados. Neste exemplo, existem `2` pares nessa situação.

*Em quantos pares de atribuições um dos intervalos contém completamente o outro?*

## --- Parte Dois ---

Parece que mesmo assim ainda existe alguma duplicação de trabalho planeado. Os Elfos *gostariam* antes de saber o número de pares em que existe *alguma sobreposição*.

No exemplo acima, os primeiros dois pares (`2-4,6-8` e `2-3,4-5`) não se sobrepõem, enquanto os outros quatro pares (`5-7,7-9`, `2-8,3-7`, `6-6,4-6`, e `2-6,4-8`) sim:


 - `5-7,7-9` se sobrepõem em apenas uma secção, `7`.

 - `2-8,3-7` se sobrepõem todas as secções da `3` até à `7`.

 - `6-6,4-6` se sobrepõem em apenas uma secção, `6`.

 - `2-6,4-8` se sobrepõem nas secções `4`, `5`, e `6`.


Então, neste exemplo, o número de pares atribuídos que se sobrepõem é `4`.

*Em quantos pares de atribuições há alguma sobreposição de intervalos?*

