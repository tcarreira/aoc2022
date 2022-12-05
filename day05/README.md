# --- Dia 5: Pilhas de Suprimentos ---

A expedição pode partir assim que os últimos suprimentos tenham sido descarregados dos navios. Os suprimentos são guardados em pilhas de *caixotes* marcados, mas porque os suprimentos necessários estão enterrados debaixo de muitos outros caixotes, os caixotes precisam ser reorganizados.

O navio tem um *guindaste de carga gigante*, capaz de mover caixotes entre pilhas. Para garantir que nenhum caixote seja esmagado ou caia, a operadora do guindaste vai reordená-los seguindo um conjunto de passos cuidadosamente planeados. Depois dos caixotes serem reorganizados, o caixotes desejados estarão no topo de cada pilha.

Os Elfos não querem interromper a operadora do guindaste durante este delicado procedimento, mas eles esqueceram perguntar-lhe *qual* caixote irá ficar onde, e eles querem estar preparados para descarrega-los assim que possível, para que possam embarcar.

Contudo, eles têm um desenho das pilhas de caixotes inicial, *e* o procedimento de reorganização (o input do puzzle). Por exemplo:

```
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

```

Neste exemplo, existem três pilhas de caixotes. A pilha 1 contém dois caixotes: caixote `Z` está no fundo, e o caixote `N` está no topo. A pilha 2 contém três caixotes; do fundo para o topo, eles são caixotes `M`, `C`, e `D`. Finalmente, a pilha 3 contém um único caixote, `P`.

Então, o procedimento de reorganização é dado. Em cada passo do procedimento, uma quantidade de caixotes é movida de uma pilha para uma pilha diferente. No primeiro passo do procedimento de reorganização acima, um caixote é movido da pilha 2 para a pilha 1, resultando nesta configuração:

```
[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 

```

No segundo passo, três caixotes são movidos da pilha 1 para a pilha 3. Os caixotes são movidos *um de cada vez*, então o primeiro caixote a ser movido (`D`) acaba abaixo do segundo e terceiro caixotes:

```
        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3

```

Depois, ambos os caixotes são movidos da pilha 2 para a pilha 1. Novamente, porque os caixotes são movidos *um de cada vez*, o caixote `C` acaba abaixo do caixote `M`:

```
        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3

```

Finalmente, um caixote é movido da pilha 1 para a pilha 2:

```
        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3

```

Os Elfos só precisam saber *qual caixote irá acabar no topo de cada pilha*; neste exemplo, os caixotes no topo são `C` na pilha 1, `M` na pilha 2, e `Z` na pilha 3, então é preciso juntar estes três e dar aos Elfos a mensagem `CMZ`.

*Depois do procedimento de reorganização terminar, qual caixote acaba no topo de cada pilha?*

## --- Parte Dois ---

Enquanto observa a operadora do guindaste reorganizar habilmente os caixotes, você nota que o processo não segue as previsões.

Alguma lama estava cobrindo a escrita no lado do guindaste, e você rapidamente limpa-a. O guindaste não é um MoveCaixotes 9000 - é um **MoveCaixotes 9001**.

O MoveCaixotes 9001 é notável por várias novas funcionalidades: ar condicionado, assentos em pele, um porta-copos extra, e *a capacidade de apanhar e mover vários caixotes de uma só vez*./

Considerando novamente o exemplo acima, os caixotes começam na mesma configuração:

```
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

```

Mover apenas um caixote da pilha 2 para a pilha 1 tem o mesmo comportamento de antes:

```
[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 

```

Contudo, a ação de mover três caixotes da pilha 1 para a pilha 3 significa que esses três caixotes *ficam na mesma ordem*, resultando nesta nova configuração:

```
        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3

```

Depois, quando ambos os caixotes são movidos da pilha 2 para a pilha 1, eles também *mantêm a ordem*:

```
        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3

```

FInalmente, um único caixote ainda é movido da pilha 1 para a pilha 2, mas é o caixote `C` que é movido:

```
        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3

```

Neste exemplo, o MoveCaixotes 9001 deixou os caixotes em uma ordem totalmente diferente: `MCD`.

Antes que o processo de reorganização termine, atualize o seu simulador para que os ELfos saibam onde devem esperar para descarregar os últimos suprimentos. *Depois do processo de reorganização estar completo, qual caixote termina em qual pilha?*

