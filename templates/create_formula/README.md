# Como adicionar uma nova linguagem ao rit create formula

Para adicionar uma nova linguagem ao rit create formula, 
você apenas precisa adicionar uma nova pasta no caminho templates/create_formula/languages.
todos os nomes de pastas adicionadas nesse caminho serão listadas para o usuário quando ele executar rit create formula.

Dentro dessa pasta você deve adicionar:

- /src                  (obrigatório)
- Makefile              (caso rode no linux)
- build.bat             (caso rode no windows)
- config.json           (obrigatório)
- Dockerfile            (caso rode com docker)
- README.md             (documentação github)
- set_umask.sh          (caso rode com docker)
- metadata.json         (documentação portal)

Vamos entender para que server cada componente:

### Source 

- **/src** nessa pasta você vai colocar uma formula de exemplo na linguagem que você quer adicionar.

- **config.json** nesse arquivo você deve adicionar os inputs da formula que está na pasta src. alem disso você pode adicionar o campo *dockerImageBuilder* que é a imagem que deve ser utiliza para fazer o build do seu código. Lembrando que esse campo é opcional caso ele não exista o ritchie sempre vai fazer o build local

### Build Local

Todo build deve gerar uma pasta bin com no minimo os seguintes arquivos:
- run.sh
- run.bat

Esses arquivos são os **arquivos de execução** que serão executados ao chamar a formula logo esses arquivos devem saber como rodar o código.

- **Makefile** é responsável por fazer o build do código em maquinas linux.
- **build.bat** é responsável por fazer o build do código em maquinas windows.

alem de gerar os **arquivos de execução** o build também deve copiar os arquivos necessários para a pasta bin para que os  **arquivos de execução** consigam funcionar.

### Build com Docker

Caso você adicione o campo *dockerImageBuilder* no **config.json** o ritchie vai tentar fazer o build utilizando o docker. com isso ele vai rodar o arquivo 
**Makefile** dentro de um docker com a imagem do campo *dockerImageBuilder*
o build com docker deve gerar a mesma pasta bin que um build local, a grande vantagem do build com docker é que o usuário não precisar ter as ferramentas necessárias para o build instaladas na maquina.

### Run com Docker

Caso você adicione o arquivo **Dockerfile** o ritchie vai rodar o **run.sh** dentro de um docker. Para isso ele vai utilizar os arquivos:

- **Dockerfile** utilizaremos esse arquivo para fazer o build do docker que vai rodar o **run.sh**
- **set_umask.sh** é o entrypoint do docker, normalmente esse arquivo utiliza o comando umask para que o volume dentro do docker tenha uma melhor compatibilidade.

**Lembrando que você deve copiar esses arquivos para a pasta bin ao fazer o build.**

### Documentação

- **metadata.json** arquivo utilizado pelo portal de formulas do ritchie para fazer a indexação de formulas.
- **README.md** arquivo para explicar como utilizar a formula, quando alguém abrir a pasta da sua formula pelo github ele vai ver o conteúdo desse arquivo

### Pasta root

Quando o usuário cria uma formula pela primeira vez em um workspace ele copia os arquivo da pasta root para o workspace do usuário, caso a linguagem tenha alguma regra nova de gitignore você pode adicionar essa regra no arquivo .gitignore na pasta root.
