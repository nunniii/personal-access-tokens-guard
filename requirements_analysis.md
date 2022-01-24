
# Guia de engenharia e arquitetura 

# Requisitos

## Funcionais

- [X] Módulo de interação com o usuário.
- [x] Módulo de autenticação; 
    * A autenticação será feita através de uma senha mestre.
- [] Módulo de criptografia (simétrica);
- [X] Módulo responsável por ler e escrever os dados (arquivos criptografados contendo os tokens e hash da senha mestre).


**Os tokens serão salvos em um ou vários arquivos criptografados, e organizado, como no exemplo:**

|Apelido / nome do arquivo |Token / conteúdo do arquivo|
|---|---|
|Token do projeto exemplo| `asdfqwer1234`
|Token trabalho| `123456ytrewq`
