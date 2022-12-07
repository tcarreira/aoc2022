# --- Dia 7: Sem espaço no dispositivo ---

Você pode ouvir o chilrear dos pássaros e as gotas de chuva batendo nas folhas enquanto a expedição avança. Ocasionalmente, você pode até ouvir sons muito mais altos à distância; quão grandes ficam os animais aqui, afinal?

O dispositivo que os Elfos lhe deram tem problemas com mais do que apenas seu sistema de comunicações. Você tenta executar uma atualização do sistema:

```
$ atualizar-sistema --por-favor --vá-lá-por-favor-com-açúcar-em-cima
Erro: Sem espaço no dispositivo

```

Talvez você possa excluir alguns arquivos para liberar espaço para a atualização?

Você navega pelo sistema de arquivos para avaliar a situação e guarda a saída do terminal resultante (sua entrada do puzzle). Por exemplo:

```
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k

```

O sistema de arquivos consiste em uma árvore de arquivos (dados simples) e diretórios (que podem conter outros diretórios ou arquivos). O diretório mais externo é chamado `/`. Você pode navegar pelo sistema de arquivos, entrando ou saindo de diretórios e listar o conteúdo do diretório em que você está atualmente.

Na saída do terminal, as linhas que começam com `$` são *comandos que você executou*, muito parecidos com alguns computadores modernos:


  - `cd` significa *muda de diretório*. Isso altera qual diretório é o diretório atual, mas o resultado específico depende do argumento:

    - `cd x` *entra* um nível: ele procura no diretório atual pelo diretório denominado `x` e o torna o diretório atual.

    - `cd ..` *sai* um nível: ele encontra o diretório que contém o diretório atual, depois torna esse diretório o diretório atual.

    - `cd /` muda o diretório atual para o diretório mais externo, `/`.




  - `ls` significa *lista*. Imprime todos os arquivos e diretórios imediatamente contidos no diretório atual:
  
    - `123 abc` significa que o diretório atual contém um arquivo chamado `abc` com tamanho `123`.

    - `dir xyz` significa que o diretório atual contém um diretório chamado `xyz`.





Dados os comandos e a saída no exemplo acima, você pode determinar que o sistema de arquivos se parece visualmente com este:

```
- / (dir)
   - a (dir)
     - e (dir)
       - i (arquivo, tamanho=584)
     - f (arquivo, tamanho=29116)
     - g (arquivo, tamanho=2557)
     - h.lst (arquivo, tamanho=62596)
   - b.txt (arquivo, tamanho=14848514)
   - c.dat (arquivo, tamanho=8504156)
   - d (dir)
     - j (arquivo, tamanho=4060174)
     - d.log (arquivo, tamanho=8033020)
     - d.ext (arquivo, tamanho=5626152)
     - k (arquivo, tamanho=7214296)

```

Aqui, existem quatro diretórios: `/` (o diretório mais externo), `a` e `d` (que estão dentro de `/`), e `e` (que está dentro de `a`). Esses diretórios também contêm arquivos de vários tamanhos.

Como o disco está cheio, seu primeiro passo provavelmente deve ser encontrar diretórios que sejam bons candidatos para exclusão. Para fazer isso, você precisa determinar o *tamanho total* de cada diretório. O tamanho total de um diretório é a soma dos tamanhos dos arquivos que ele contém, direta ou indiretamente. (Os próprios diretórios não contam como tendo tamanho intrínseco.)

Os tamanhos totais dos diretórios acima podem ser encontrados da seguinte forma:


  - O tamanho total do diretório `e` é *584* porque contém um único arquivo `i` de tamanho 584 e nenhum outro diretório.

  - O diretório `a` tem tamanho total *94853* porque contém os arquivos `f` (tamanho 29116), `g` (tamanho 2557) e `h.lst` (tamanho 62596), além do arquivo `i` indiretamente (`a` contém `e` que contém `i`).

  - O diretório `d` tem tamanho total *24933642*.

  - Como diretório mais externo, `/` contém todos os arquivos. Seu tamanho total é *48381165*, a soma do tamanho de todos os arquivo.


Para começar, encontre todos os diretórios com um tamanho total de *no máximo 100000*, e depois calcule a soma de seus tamanhos totais. No exemplo acima, esses diretórios são `a` e `e`; a soma de seus tamanhos totais é `95437` (94853 + 584). (Como neste exemplo, este processo pode contar arquivos mais de uma vez!)

Encontre todos os diretórios com um tamanho total de no máximo 100.000. *Qual é a soma dos tamanhos totais desses diretórios?*

