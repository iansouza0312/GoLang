# Guia de Estudos - Go Language

O Go e uma linguagem fortemente tipada, com diversas caracteristicas e sitaxes diferentes das que estamos acostumados a ver em linguagens como por exemplo Java, Rust, JS ...

### Declarando variaveis - VARIABLES

No go, temos diversas maneiras de declarar variaveis, de maneira explicita ou implicita em sua tipagem.
Uma das peculiaridades do Go e que, toda a várivael declarada no código deve necessariamente ser utilizada, caso contrário seu código não irá ser compilado.
O mesmo server para declarações de Bibliotecas/Pacotes.

**exemplos de decalração de váriavel :**

- var firstVar = "Primeira Variavel"
- secondVar := "Segunda Variavel"

### Tipos de Dados - DATATYPES

Como em qualquer linguagem, o go tem diversas maneira e sintaxes quando se trata de trabalhar com dados, alguns únicos da linguagem, como por exemplo os **Structs**.
Quando definimos o tipo de dado, temos que tomar cuidado pois por exemplo com valores numericos, declaramos int e junto o numero maximo de bits que o valor da variavel pode ter.

**exemplos de decalração de váriavel inteira :**

- var number int16 = 100

_caso declarasse uma variavel com um numero muito grande, um erro iria ocorrer_

**exemplos de erro por conta do tipo da variavel :**

- var number int8 = 100000000000

**_IMPORTANTE : Caso o vvalor em bits não seja informado, o go utiliza a arquitetura do seu computador (32, 64)_**

Alguns tipos de dados em Go acitam apelidos (alias), como por exemplo : int32 = rune; uint8 = byte.
