# GetDolarAWS
[![Go Reference](https://pkg.go.dev/badge/mitvix/cloudcost.svg)](https://pkg.go.dev/mitvix/cloudcost)

Utilitário de extração do valor do dólar em reais de invoices AWS no formato PDF.

## Overview

Este utilitário utiliza o pacote popper-utils (pdftotext) disponível em sistemas *nix para processar via GO a extração do valor do Dólar em Reais de arquivos de invoices AWS no formato .pdf.


## [ESTUDO DE CASO]
Software criado _(nas raras horas vagas)_ para estudo e análise da línguagem Go (Golang) disponível em [go.dev](https://go.dev). Línguagem de programação opensource criada por [Rob Pike](https://pt.wikipedia.org/wiki/Rob_Pike), [Robert Griesemer](https://en.wikipedia.org/wiki/Robert_Griesemer) e [Ken Thompson](https://pt.wikipedia.org/wiki/Ken_Thompson) nos laboratórios do Google em meados de 2007 e liberado sob licença opensource BSD em 2009.

<details>

<summary>Sobre Go e por que essa linguagem</summary>

### Golang

Go foi projetado inicialmente com o objetivo de substituir projetos em C e C++ dentro do Google, por isso possui características simílares a essas línguagens, incluindo sua síntaxe, mas com abstrações voltadas a simplicidade e legibilidade, além de uma forte combinação de suporte a concorrência e desempenho. Sua estrutura automática de gerenciamento de memória (Garbage Collector) facilita a vida do desenvolvedor, mas gera overhead que a deixa pouco atrás em performance quando comparada a C, C++ e Rust, porém, muito a frente em desempenho em relação a Python, Java, PHP e etc. 

E mesmo perdendo em performance para Rust e C++, Go se tornou uma línguagem equilibrada que combina estruturas de baixo nível de C com a usabilidade do mundo moderno e sem o pesadelo da Orientação a Objetos, fazendo dela uma línguagem de programação simples, completa e perfeita para o uso em APIs, Micro serviços, Web Dev, Cloud e etc. 

Dentre os principais projetos escritos em Go, temos Kubernetes, kubectl, Minikube, Docker e outros. Veja mais em [https://go.dev/solutions/cloud#use-case](https://go.dev/solutions/cloud#use-case)

### Por que Go ?

_"Queria que fosse em que? C, Rust ???? mal sei fazer em Go, não sai da 5º série..."_


</details>

## Instalação

Instalação (Linux)
```
git clone https://github.com/mitvix/getdolaraws
cd getdolaraws
go build -o getdolaraws main.go
sudo mv getdolaraws /usr/local/bin/getdolaraws

# Debian, Ubuntu etc
sudo apt install popper-utils
# Arch, Manjaro etc
sudo pacman -Syu popper-utils

getdolaraws -invoice Arquivo_da_InvoiceBR.pdf

```

Argumentos:

* -invoice `Arquivo da invoice em formato .pdf`


## Nota do Autor e Contato

* _IMPORTANTE Programar é uma arte, um passa tempo e faço por hobby, não sou programador profissional e não pretendo ser. Se der em algum momento atualizo para usar a Standard Lib ou um pacote nativo para não depender do pdftotext, mas funciona tão bem... que passa ;-)_

Manfrin <mitvix@hotmail.com>
