# --- Dia 6: Problema de ajuste ---

Os preparativos estão finalmente completos; você e os elfos deixam o acampamento a pé e começam a caminhar em direção ao arvoredo do fruto **estrela**.

Conforme você se move pela densa vegetação rasteira, um dos elfos lhe dá um *dispositivo* portátil. Ele diz que tem muitos recursos sofisticados, mas o mais importante para configurar agora é o *sistema de comunicações*.

No entanto, porque ele ouviu falar que você tem [experiência](https://adventofcode.com/2016/day/25) [significativa](https://adventofcode.com/2016/day/6) [trabalhando](https://adventofcode.com/2019/day/7) [com](https://adventofcode.com/2019/day/9) [sistemas](https://adventofcode.com/2021/dia/25) [baseados em sinal](https://adventofcode.com/2019/day/16), ele convenceu os outros Elfos de que não haveria problema em dar a você o único dispositivo com defeito - certamente você não terá problemas para consertá-lo.

Como que inspirado pelo timing cômico, o dispositivo emite algumas *faíscas coloridas*.

Para poder se comunicar com os Elfos, o dispositivo precisa *travar no sinal* deles. O sinal é uma série de caracteres aparentemente aleatórios que o dispositivo recebe um de cada vez.

Para corrigir o sistema de comunicações, você precisa adicionar uma sub-rotina ao dispositivo que deteta um *marcador de início de pacote* no fluxo de dados. No protocolo usado pelos Elfos, o início de um pacote é indicado por uma sequência de *quatro caracteres todos diferentes*.

O dispositivo enviará à sua sub-rotina um buffer de fluxo de dados (sua entrada do puzzle); sua sub-rotina precisa identificar a primeira posição onde os quatro caracteres recebidos mais recentemente eram todos diferentes. Especificamente, ele precisa relatar o número de caracteres desde o início do buffer até o final do primeiro marcador de quatro caracteres.

Por exemplo, suponha que você receba o seguinte buffer de fluxo de dados:

`mjqjpqmgbljsphdztnvjfqwrcgsmlb`

Depois que os três primeiros caracteres (`mjq`) foram recebidos, ainda não foram recebidos caracteres suficientes para localizar o marcador. A primeira vez que um marcador pode ocorrer é após o quarto caractere ser recebido, tornando os quatro caracteres mais recentes `mjqj`. Como `j` é repetido, isso não é um marcador.

A primeira vez que um marcador aparece é após a chegada do *sétimo* caractere. Depois disso, os últimos quatro caracteres recebidos são `jpqm`, que são todos diferentes. Neste caso, sua sub-rotina deve informar o valor `7`, porque o primeiro marcador de início de pacote é concluído após o processamento de 7 caracteres.

Aqui estão mais alguns exemplos:


 - `bvwbjplbgvbhsrlpgdmjqwftvncz`: primeiro marcador após o caractere `5`

 - `nppdvjthqldpwncqszvftbrmjlhg`: primeiro marcador após o caractere `6`

 - `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`: primeiro marcador após o caractere `10`

 - `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`: primeiro marcador após o caractere `11`


*Quantos caracteres precisam ser processados antes que o primeiro marcador de início de pacote seja detetado?*

## --- Parte dois ---

O sistema de comunicações do seu dispositivo está detectando corretamente os pacotes, mas ainda não está funcionando. Parece que também precisa procurar por *mensagens*.

Um *marcador de início de mensagem* é como um marcador de início de pacote, exceto que consiste em *14 caracteres distintos* em vez de 4.

Aqui estão as primeiras posições dos marcadores de início de mensagem para todos os exemplos acima:


  - `mjqjpqmgbljsphdztnvjfqwrcgsmlb`: primeiro marcador após o caractere `19`

  - `bvwbjplbgvbhsrlpgdmjqwftvncz`: primeiro marcador após o caractere `23`

  - `nppdvjthqldpwncqszvftbrmjlhg`: primeiro marcador após o caractere `23`

  - `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`: primeiro marcador após o caractere `29`

  - `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`: primeiro marcador após o caractere `26`


*Quantos caracteres precisam ser processados antes que o primeiro marcador de início de mensagem seja detectado?*

