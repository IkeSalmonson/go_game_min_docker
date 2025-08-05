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
    - O tabuleiro: Altura e largura são definidos ao inicio da partida e devem ser entre 5 a 8 espaços. Exemplos de formatos aceitos: 5x5 , 8x5, 5x8 e 6x7.
    - As peças:
        São posicionadas 
            Pretas: na esquerda e topo
            Brancas: na direita e fundo  
        As peças podem se movimentar para um espaço vazio adjacente ou capturar uma peça oponente pulando o espaço dela.
    - Jogo 
        As peças brancas iniciam o primeiro turno.
        Para ganhar o jogo os jogadores devem eliminar todas as peças do oponente ou ter mais peças ao final de 10 turnos. Caso os jogadores tenham o mesmo numero de peças ao final de 10 turnos, o jogo termina em empate.      


## ⚙️ Como Rodar

**Pré-requisitos:**
- Docker

**Passos:**
1. Clone o repositório:
   ```bash
   git clone https://github.com/IkeSalmonson/go_game_min_docker
   ```
Navegue até o diretório do projeto:
``` bash
cd go_game_min_docker
```

	Docker image Prod / otimizada : 
	
	docker build -f Dockerfile.prod -t ikesalmonson/go-game-min-docker:prod .
	docker run --rm -it   ikesalmonson/go-game-min-docker:prod 

[Próximo passo, ex: Crie um ambiente virtual...]
[Próximo passo, ex: Instale as dependências...]
[Próximo passo, ex: Inicie os containers Docker...]

### ⚙️ Como Rodar para desenvolvimento
	Criar a imagem Docker Dev:
   ``` 
    docker build -t ikesalmonson/go-game-min-docker:dev .   
   ```
Executar a imagem com o codigo Go incluso via volume e executar a imagem dev em modo iterativo: 
``` 
    docker run --rm -it -v $(pwd):/go/src ikesalmonson/go-game-min-docker:dev bash 
```


## 📊 Resultados
[inserir tabela comparativa do tamanho das imagens de dev e prod]


## ✨ Demonstração

[Insira aqui um GIF animado ou algumas screenshots que mostrem seu projeto em ação. Ferramentas como ScreenToGif (Windows)]
