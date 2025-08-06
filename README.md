# GO GAME MIN DOCKER 

## 📝 Descrição
Jogo derivado de damas escrito em linguagem GO para exemplo de aplicação em uma imagem otimizada em docker. <br>
A imagem docker otimizada realiza a compilação dos codigos em go para um binário e empacota em uma imagem docker Scrach. Minimizando o uso de recursos de computação na execução da aplicação.


## 💻 Tecnologias Utilizadas
- Docker version 28.3.2
- Golang 1.21.3 linux/amd64

## 🚀 Recursos e Funcionalidades
Jogo damas 3x3 em linha de comando.

### Regras do jogo: 
- Tabuleiro: Altura e largura são definidos ao inicio da partida e devem ser entre 5 a 8 espaços.
	- Exemplos de formatos aceitos: 5x5 , 8x5, 5x8 e 6x7
- Peças: 3 peças para cada jogador. 
  - Posicionamento inicial: <br>
  	- **Pretas**: na esquerda e topo <br>
        - **Brancas**: na direita e fundo  <br>
  - Movimento: As peças podem se movimentar para um espaço vazio adjacente ou capturar uma peça oponente pulando o espaço dela. <br>
- Jogo: Os jogadores determinam o tamanho do tabuleiro a ser utilizado.  
	- As peças brancas iniciam o primeiro turno fazendo uma movimentação e alternando a vez para o outro jogador até que ocorra uma condição de fim do jogo. <br>
	- Condições de fim do jogo: <br>
		- Vitória por eliminação: O jogador que eliminar todas as peças do oponente vence. <br>
		- Vitória por pontuação:  O jogador com mais peças ao final de 10 turnos vence. <br>
		- Empate: Caso os jogadores tenham o mesmo numero de peças ao final de 10 turnos, o jogo termina em empate.<br>


## ⚙️ Como Rodar o Jogo

**Pré-requisitos:**
- Docker

**Passos:**
1. Clone o repositório:
```bash
git clone https://github.com/IkeSalmonson/go_game_min_docker
```

2. Navegue até o diretório do projeto:
``` bash
cd go_game_min_docker
```

3. Crie a  imagem Docker Prod (otimizada) : 
``` bash
docker build -f Dockerfile.prod -t ikesalmonson/go-game-min-docker:prod .
```
4. Execute a imagem criada em modo iterativo:
``` bash
docker run --rm -it --name go-game-min-docker-prod  ikesalmonson/go-game-min-docker:prod 
``` 

### ⚙️ Como Rodar para desenvolvimento
Criar a imagem Docker Dev:
``` 
docker build -t ikesalmonson/go-game-min-docker:dev .   
```

Executar a imagem com o codigo Go incluso via volume e execute a imagem dev em modo iterativo: 
``` 
docker run --rm -it --name go-game-min-docker-dev -v $(pwd):/go/src ikesalmonson/go-game-min-docker:dev bash 
```
Para iniciar o jogo, execute o comando:
``` 
go run cmd/main.go
``` 


## 📊 Resultados

---

### Comparativo de Tamanho das Imagens Docker

Esta tabela demonstra a diferença no tamanho das imagens de desenvolvimento e produção, conforme listado pelo comando `docker images`. A imagem de produção, construída com `scratch`, é significativamente menor.

| Repositório / Imagem | Tag    | ID da Imagem | Criada Há | Tamanho |
| :------------------- | :----- | :----------- | :-------- | :------ |
| ikesalmonson/go-game-min-docker | prod | bb315d8b50cb | 22 horas  | **1.89MB** |
| ikesalmonson/go-game-min-docker | dev  | cbf45488a510 | 27 horas  | **815MB** |

---

### Comparativo de Consumo de Recursos em Tempo Real

Esta tabela mostra o uso de recursos (CPU, Memória, I/O de Disco) dos containers de desenvolvimento e produção em estado ocioso (aguardando as dimensões do tabuleiro), obtido via `docker stats`. Note a economia de memória e I/O na versão de produção.

| ID do Container | Nome do Container       | CPU %  | Uso de Memória / Limite | Memória % | I/O de Rede | I/O de Bloco | PIDs |
| :-------------- | :---------------------- | :----- | :---------------------- | :-------- | :---------- | :----------- | :--- |
| 672966a094ad  | go-game-min-docker-dev | 0.00% | 72.89MiB / 7.757GiB   | 0.92%   | 992B / 126B | 51.3MB / 56.3MB | 17 |
| 6a50fe30014a  | go-game-min-docker-prod| 0.00% | 916KiB / 7.757GiB     | 0.01%   | 796B / 126B | 0B / 0B    | 5  |

---


## ✨ Demonstração

[Insira aqui um GIF animado ou algumas screenshots que mostrem seu projeto em ação. Ferramentas como ScreenToGif (Windows)]
