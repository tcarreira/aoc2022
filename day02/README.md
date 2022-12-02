# --- Dia 2: Pedra Papel Tesoura ---

Os Elfos começam a montar acampamento na praia. Para decidir qual tenda fica mais próxima do armazém de petiscos, um gigantesco [Pedra Papel Tesoura](https://pt.wikipedia.org/wiki/Pedra-papel-tesoura) já está em andamento.

Pedra Papel Tesoura é um jogo entre dois jogadores. Cada jogo contém muitas rodadas; em cada rodada, os jogadores escolhem simultaneamente uma das formas de Pedra, Papel ou Tesoura. Então, um vencedor para aquela rodada é selecionado: Pedra vence Tesoura, Tesoura vence Papel, e Papel vence Pedra. Se ambos os jogadores escolherem a mesma forma, a rodada termina em um empate.

Agradecido pela sua ajuda ontem, um Elfo te dá uma *estratégia criptografada* (seu *input* do puzzle) que diz que vai te ajudar a ganhar. "A primeira coluna é o que seu oponente vai jogar: `A` para Pedra, `B` para Papel, e `C` para Tesoura. A segunda coluna--" De repente, o Elfo é chamado para ajudar com a tenda de alguém.

A segunda coluna, *você deduz*, deve ser o que você deve jogar em resposta: `X` para Pedra, `Y` para Papel, e `Z` para Tesoura. Ganhar todas as vezes seria suspeito, então as respostas devem ter sido escolhidas com cuidado.

O vencedor do torneio inteiro é o jogador com a maior pontuação. Sua *pontuação total* é a soma de suas pontuações para cada rodada. A pontuação para uma única rodada é a pontuação para a *forma que você escolheu* (1 para Pedra, 2 para Papel, e 3 para Tesoura) mais a pontuação para o *resultado da rodada* (0 se você perdeu, 3 se a rodada foi um empate, e 6 se você ganhou).

Como você não pode ter certeza se o Elfo está tentando te ajudar ou te enganar, você deve calcular a pontuação que você teria se seguisse a estratégia.

Por exemplo, suponha que você tenha recebido a seguinte estratégia:

```
A Y
B X
C Z

```

Essa estratégia prevê e recomenda o seguinte:


 - Na primeira rodada, seu oponente vai escolher Pedra (`A`), e você deve escolher Papel (`Y`). Isso termina em uma vitória para você com uma pontuação de *8* (2 porque você escolheu Papel + 6 porque você ganhou).

 - Na segunda rodada, seu oponente vai escolher Papel (`B`), e você deve escolher Pedra (`X`). Isso termina em uma derrota para você com uma pontuação de *1* (1 + 0).

 - A terceira rodada é um empate com ambos os jogadores escolhendo Tesoura, dando a você uma pontuação de 3 + 3 = *6*.


No final, se você seguisse a estratégia, você teria uma pontuação total de  `15` (8 + 1 + 6).

*Qual é a pontuação total que você teria se seguisse a estratégia do Elfo?*

## --- Parte Dois ---

O Elfo acabou de ajudar com a tenda e volta para perto de você. "De qualquer forma, a segunda coluna diz como a rodada deve terminar: `X` se você deve perder, `Y` se você deve empatar, e `Z` se você deve ganhar. Bom sorte!"

A pontuação total ainda é calculada da mesma forma, mas agora você precisa escolher a forma que você vai jogar para que a rodada termine como indicado. O exemplo acima decorrrria da seguinte forma:


 - Na primeira rodada, seu oponente vai escolher Pedra (`A`), e você precisa que a rorada termine em empate (`Y`), então você também escolhe Pedra. Isso dá uma pontuação de 1 + 3 = *4*.

 - Na segunda rodada, seu oponente vai escolher Papel (`B`), e você escolhe Pedra para perder (`X`) com uma pontuação de 1 + 0 = *1*.

 - Na terceira rodada, você ganha à Tesoura do seu oponente com Pedra, para uma pontuação de 1 + 6 = *7*.


Agora que você decifrou corretamente a estratégia ultra secreta, você teria uma pontuação total de `12`.

Seguindo as instruções do Elfo para a segunda coluna, *qual é a pontuação total que você teria se seguisse a estratégia?*

