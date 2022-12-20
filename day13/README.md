# --- Dia 13: Sinal de Socorro ---

Você sobe a colina e novamente tenta entrar em contato com os Elfos. No entanto, você recebe um sinal que não esperava: um *sinal de socorro*.

Seu dispositivo portátil ainda não deve estar funcionando corretamente; os pacotes do sinal de socorro foram decodificados *fora de ordem*. Você precisará reordenar a lista de pacotes recebidos (sua entrada do puzzle) para decodificar a mensagem.

Sua lista consiste em pares de pacotes; pares são separados por uma linha em branco. Você precisa identificar *quantos pares de pacotes estão na ordem correta*.

Por exemplo:

```
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]

```

*Os dados do pacote consistem em listas e números inteiros.* Cada lista começa com `[`, termina com `]` e contém zero ou mais valores separados por vírgula (inteiros ou outras listas). Cada pacote é sempre uma lista e aparece em sua própria linha.

Ao comparar dois valores, o primeiro valor é chamado de *esquerda* e o segundo valor é chamado de *direita*. Então:


  - Se *ambos os valores forem inteiros*, o *inteiro inferior* deve vir primeiro. Se o inteiro esquerdo for menor que o inteiro direito, as entradas estarão na ordem correta. Se o inteiro esquerdo for maior que o inteiro direito, as entradas não estão na ordem correta. Caso contrário, as entradas são o mesmo número inteiro; continue verificando a próxima parte da entrada.

  - Se *ambos os valores forem listas*, compare o primeiro valor de cada lista, depois o segundo valor e assim por diante. Se a lista da esquerda ficar sem itens primeiro, as entradas estarão na ordem correta. Se a lista direita ficar sem itens primeiro, as entradas não estarão na ordem correta. Se as listas tiverem o mesmo tamanho e nenhuma comparação tomar uma decisão sobre a ordem, continue verificando a próxima parte da entrada.

  - Se *exatamente um valor for um número inteiro*, converta o número inteiro em uma lista que contenha esse número inteiro como seu único valor e repita a comparação. Por exemplo, se comparar `[0,0,0]` e `2`, converta o valor direito para `[2]` (uma lista contendo `2`); o resultado é então encontrado comparando `[0,0,0]` e `[2]`.


Usando essas regras, você pode determinar quais pares no exemplo estão na ordem correta:

```
== Par 1 ==
- Compare [1,1,3,1,1] vs [1,1,5,1,1]
   - Compare 1 vs 1
   - Compare 1 vs 1
   - Compare 3 vs 5
     - O lado esquerdo é menor, então as entradas estão na ordem certa

== Par 2 ==
- Compare [[1],[2,3,4]] vs [[1],4]
   - Compare [1] vs [1]
     - Compare 1 vs 1
   - Compare [2,3,4] vs 4
     - Tipos mistos; converter a direita para [4] e repetir a comparação
     - Compare [2,3,4] vs [4]
       - Compare 2 vs 4
         - O lado esquerdo é menor, então as entradas estão na ordem correta

== Par 3 ==
- Compare [9] vs [[8,7,6]]
   - Compare 9 vs [8,7,6]
     - Tipos mistos; converter à esquerda para [9] e repetir a comparação
     - Compare [9] vs [8,7,6]
       - Compare 9 vs 8
         - O lado direito é menor, então as entradas não estão na ordem correta

== Par 4 ==
- Compare [[4,4],4,4] vs [[4,4],4,4,4]
   - Compare [4,4] vs [4,4]
     - Compare 4 vs 4
     - Compare 4 vs 4
   - Compare 4 vs 4
   - Compare 4 vs 4
   - O lado esquerdo ficou sem itens, então as entradas estão na ordem certa

== Par 5 ==
- Compare [7,7,7,7] vs [7,7,7]
   - Compare 7 vs 7
   - Compare 7 vs 7
   - Compare 7 vs 7
   - O lado direito ficou sem itens, então as entradas não estão na ordem correta

== Par 6 ==
- Compare [] vs [3]
   - O lado esquerdo ficou sem itens, então as entradas estão na ordem certa

== Par 7 ==
- Compare [[[]]] vs [[]]
   - Compare [[]] vs []
     - O lado direito ficou sem itens, então as entradas não estão na ordem correta

== Par 8 ==
- Compare [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]] ]],8,9]
   - Compare 1 vs 1
   - Compare [2,[3,[4,[5,6,7]]]] vs [2,[3,[4,[5,6,0]]]]
     - Compare 2 vs 2
     - Compare [3,[4,[5,6,7]]] vs [3,[4,[5,6,0]]]
       - Compare 3 vs 3
       - Compare [4,[5,6,7]] vs [4,[5,6,0]]
         - Compare 4 vs 4
         - Compare [5,6,7] vs [5,6,0]
           - Compare 5 vs 5
           - Compare 6 vs 6
           - Compare 7 vs 0
             - O lado direito é menor, então as entradas não estão na ordem correta

```

Quais são os índices dos pares que já estão *na ordem correta*? (O primeiro par tem índice 1, o segundo par tem índice 2 e assim por diante.) No exemplo acima, os pares na ordem correta são 1, 2, 4 e 6; a soma destes índices é `13`.

Determine quais pares de pacotes já estão na ordem correta. *Qual é a soma dos índices desses pares?*

## --- Parte Dois ---

Agora, você só precisa colocar *todos* os pacotes na ordem certa. Desconsidere as linhas em branco em sua lista de pacotes recebidos.

O protocolo de sinal de socorro também requer que você inclua dois *pacotes divisores* adicionais:

```
[[2]]
[[6]]

```

Usando as mesmas regras de antes, organize todos os pacotes - os da sua lista de pacotes recebidos, bem como os dois pacotes divisores - na ordem correta.

Para o exemplo acima, o resultado de colocar os pacotes na ordem correta é:

```
[]
[[]]
[[[]]]
[1,1,3,1,1]
[1,1,5,1,1]
[[1],[2,3,4]]
[1,[2,[3,[4,[5,6,0]]]],8,9]
[1,[2,[3,[4,[5,6,7]]]],8,9]
[[1],4]
[[2]]
[3]
[[4,4],4,4]
[[4,4],4,4,4]
[[6]]
[7,7,7]
[7,7,7,7]
[[8,7,6]]
[9]

```

No final, localize os pacotes divisores. Para encontrar a *chave decodificadora* para este sinal de socorro, você precisa determinar os índices dos dois pacotes divisores e multiplicá-los. (O primeiro pacote está no índice 1, o segundo pacote está no índice 2 e assim por diante.) Neste exemplo, os pacotes divisores são *10º* e *14º* e, portanto, a chave do decodificador é `140`.

Organize todos os pacotes na ordem correta. *Qual é a chave decodificadora para o sinal de socorro?*

