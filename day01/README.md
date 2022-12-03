# --- Dia 1: Contar Calorias ---

As renas do Pai Natal costumam comer comida normal para rena, mas elas precisam de muita [energia mágica](https://adventofcode.com/2018/day/25) para entregar os presentes no Natal. Para isso, o seu snack favorito é uma espécie de fruto **estrela** que cresce apenas na selva mais profunda. Os Elfos trouxeram-te na sua expedição anual ao arvoredo onde o fruto cresce.

Para conseguir energia mágica suficiente, a expedição precisa de conseguir um mínimo de **cinquenta estrelas** até ao dia 25 de Dezembro. Apesar dos Elfos assegurarem que o arvoredo tem frutos suficientes, decides apanhar qualquer fruto que consigas ver pelo caminho, por precaução.

Apanha estrelas resolvendo quebra-cabeças. Serão disponibilizados dois quebra-cabeças a cada dia do calendário do Advento; o segundo quebra-cabeças é desbloqueado quando o primeiro for resolvido. Cada quebra-cabeças premeia com **uma estrela**. Boa sorte!

A selva deve estar demasiado crescida e de difícil navegação por veículos ou acesso pelo ar; a expedição dos Elfos é tradicionalmente feita a pé. Enquanto os vossos barcos se aproximam da terra, os Elfos começam a fazer um inventário dos seus mantimentos - mais precisamente, o número de *Calorias* que cada Elfo carrega consigo (a entrada do teu quebra-cabeças).

À vez, os Elfos anotam o número de Calorias contido em várias refeições, snacks, rações, *etc.* que eles trouxeram com eles, um item por linha. Cada Elfo separa o seu próprio inventário do Elfo anterior (se existir) por uma linha em branco.

Por exemplo, supomos que os Elfos terminaram de escrever as Calorias dos seus items e acabamos com a seguinte lista:

```
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

```

Esta lista representa as Calorias da comida transportada pelos Elfos:


 - O primeiro Elfo leva comida com  `1000`, `2000`, e `3000` Calorias, um total de `6000` Calorias.

 - O segundo Elfo leva um item de comida com `4000` Calorias.

 - O terceiro Elfo leva comida com `5000` e `6000` Calorias, um total de `11000` Calorias.

 - O quarto Elfo leva comida com `7000`, `8000`, e `9000` Calorias, um total de `24000` Calorias.

 - O quinto Elfo leva um item de comida com `10000` Calorias.


No caso dos Elfos ficarem com fome e precisarem de um snack extra, eles precisam saber a qual Elfo perguntar: eles gostariam de saber quantas Calorias leva o Elfo que carrega *mais* Calorias. No exemplo acima, isto é *`24000`* (carregado pelo quarto Elfo).

Encontra o Elfo que leva mais Calorias. *Quantas calorias esse Elfo leva no total?*

## --- Parte Dois ---

Quando terminas de calcular a resposta à pergunta dos Elfos, eles já tinham percebido que o Elfo que levava mais Calorias de comida podia eventualmente *ficar sem mais snack*.

Para evitar esta situação inaceitável, os Elfos preferiam saber o total de Calorias carregado pelo *top três* de Elfos que levam mais Calorias. Dessa forma, mesmo que um desses Elfos fique sem snacks, eles ainda têm dois suplentes.

No exemplo acima, o top três Elfos são o quarto Elfo (com `24000` Calorias), depois o terceiro Elfo (com `11000` Calorias), depois o quinto Elfo (com `10000` Calorias). A soma das Calorias levado por esses três elfos é `45000`.

Encontra o top três Elfos que leva mais Calorias. *Quantas Calorias esses três Elfos levam no total?*

