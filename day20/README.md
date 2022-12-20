# --- Dia 20: Sistema de Posicionamento do Arvoredo ---

Finalmente chegou a hora de nos encontrarmos novamente com os Elfos. Quando você tenta contatá-los, no entanto, você não obtém resposta. Talvez você esteja fora de alcance?

Você sabe que eles estão indo para o arvoredo onde a fruta **estrela** cresce, então se você descobrir onde fica, você deve ser capaz de se encontrar com eles.

Felizmente, seu dispositivo portátil possui um arquivo (sua entrada do puzzle) que contém as coordenadas do arvoredo! Infelizmente, o arquivo é *criptografado* - só para o caso de o dispositivo cair em mãos erradas.

Talvez você possa *descriptografa-lo*?

Quando você ainda estava no acampamento, você ouviu alguns elfos falando sobre criptografia de arquivos de coordenadas. A principal operação envolvida na descriptografia do arquivo é chamada de *mistura*.

O arquivo criptografado é uma lista de números. Para *misturar* o arquivo, mova cada número para frente ou para trás no arquivo um número de posições igual ao valor do número que está sendo movido. A lista é *circular*, portanto, um número movido para lá de uma extremidade da lista aparece no início da outra extremidade, como se as extremidades estivessem conectadas.

Por exemplo, para mover o `1` em uma sequência como `4, 5, 6, 1, 7, 8, 9`, o `1` move uma posição para frente: `4, 5, 6, 7, 1, 8 , 9`. Para mover o `-2` em uma sequência como `4, -2, 5, 6, 7, 8, 9`, o `-2` move duas posições para trás, aparecendo do outro lado: `4, 5, 6, 7, 8, -2, 9`.

Os números devem ser movidos *na ordem em que aparecem originalmente* no arquivo criptografado. Os números que se movem durante o processo de mistura não alteram a ordem em que os números são movidos.

Considere este arquivo criptografado:

```
1
2
-3
3
-2
0
4

```

A mistura desse arquivo ocorre da seguinte maneira:

```
Arranjo inicial:
1, 2, -3, 3, -2, 0, 4

1 é movido para entre 2 e -3:
2, 1, -3, 3, -2, 0, 4

2 é movido para entre -3 e 3:
1, -3, 2, 3, -2, 0, 4

-3 é movido para entre -2 e 0:
1, 2, 3, -2, -3, 0, 4

3 é movido para entre 0 e 4:
1, 2, -2, -3, 0, 3, 4

-2 é movido para entre 4 e 1:
1, 2, -3, 0, 3, 4, -2

0 não se move:
1, 2, -3, 0, 3, 4, -2

4 é movido para entre -3 e 0:
1, 2, -3, 4, 0, 3, -2

```

Em seguida, as coordenadas do arvoredo podem ser encontradas observando as posições 1000º, 2000º e 3000º após o valor `0`, voltando ao início da lista conforme necessário. No exemplo acima, o 1000º número após `0` é `4`, o 2000º é `-3` e o 3000º é `2`; somando-os produz `3`.

Misture seu arquivo criptografado exatamente uma vez. *Qual é a soma dos três números que formam as coordenadas do arvoredo?*

## --- Parte Dois ---

Os valores das coordenadas do arvoredo parecem absurdos. Enquanto você pondera sobre os mistérios da criptografia dos Elfos, de repente você se lembra do resto da rotina de descriptografia que ouviu no acampamento.

Primeiro, você precisa aplicar a *chave de descriptografia*, `811589153`. Multiplique cada número pela chave de descriptografia antes de começar; isso produzirá a lista real de números a serem misturados.

Em segundo lugar, você precisa misturar a lista de números *dez vezes*. A ordem em que os números são misturados não muda durante a mistura; os números ainda são movidos na ordem em que apareceram na lista original pré-misturada. (Portanto, se -3 aparecer em quarto lugar na lista original de números a serem misturados, -3 será o quarto número a ser movido durante cada rodada de mistura.)

Usando o mesmo exemplo acima:

```
Arranjo inicial:
811589153, 1623178306, -2434767459, 2434767459, -1623178306, 0, 3246356612

Após 1 rodada de mistura:
0, -2434767459, 3246356612, -1623178306, 2434767459, 1623178306, 811589153

Após 2 rodadas de mistura:
0, 2434767459, 1623178306, 3246356612, -2434767459, -1623178306, 811589153

Após 3 rodadas de mistura:
0, 811589153, 2434767459, 3246356612, 1623178306, -1623178306, -2434767459

Após 4 rodadas de mistura:
0, 1623178306, -2434767459, 811589153, 2434767459, 3246356612, -1623178306

Após 5 rodadas de mistura:
0, 811589153, -1623178306, 1623178306, -2434767459, 3246356612, 2434767459

Após 6 rodadas de mistura:
0, 811589153, -1623178306, 3246356612, -2434767459, 1623178306, 2434767459

Após 7 rodadas de mistura:
0, -2434767459, 2434767459, 1623178306, -1623178306, 811589153, 3246356612

Após 8 rodadas de mistura:
0, 1623178306, 3246356612, 811589153, -2434767459, 2434767459, -1623178306

Após 9 rodadas de mistura:
0, 811589153, 1623178306, -2434767459, 3246356612, 2434767459, -1623178306

Após 10 rodadas de mistura:
0, -2434767459, 1623178306, 3246356612, -1623178306, 2434767459, 811589153

```

As coordenadas do arvoredo ainda podem ser encontradas da mesma maneira. Aqui, o 1000º número após `0` é `811589153`, o 2000º é `2434767459` e o 3000º é `-1623178306`; somando-os produz `1623178306`.

Aplique a chave de descriptografia e misture seu arquivo criptografado dez vezes. *Qual é a soma dos três números que formam as coordenadas do arvoredo?*
