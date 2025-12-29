# conversor-homebank
Script que lê uma planilha e gera os arquivos necessários para importar informações para o homebank

## QIF

Um dos formatos aceitos pelo HomeBank é o QIF. Baseado em um arquivo de texto.

A estrutura do QIF ustilizada pelo HomeBank é a seguinte:

### Conta

```qif
!Account
NPoupanca
^
```

| Identificador | Descrição     | Exemplo     |
| ------------- | ------------- | ----------- |
| N             | Nome da Conta | `NPoupanca` |
|               |               |             |

### Transações

```qif
!Type:Bank
D2024/02/02
T-120.00
C
PLucas
MPizzaria
LLazer:Comida
^
```

Antes da primeira transação, deve-se estar descrito sua tipagem ex.: `!Type:Bank`

| Identificador | Descrição                      | Exemplo         |
| ------------- | ------------------------------ | --------------- |
| D             | Data                           | `D2024/02/02`   |
| T             | Valor                          | `T-120.00`      |
| C             | Situação: "" "c" "R"           | `C`             |
| P             | Favorecido                     | `PLucas`        |
| M             | Observações                    | `MPizzaria`     |
| L             | Categoria                      | `LLazer:Comida` |
| L[]           | Transações: [conta favorecida] | `L[Conta2]`     |
|               |                                |                 |

Um arquivo qif podem ter várias contas e transações em sequencia


## Arquivo de entrada 

### Colunas tabela

Construir uma tabela que descreve as transaçoes por meio de relações `Destino`X`Origem`, sendo o `Destino` a categoria ou conta que será destinado aquele valor, e `Origem` a conta de onde o valor será retirado. Abaixo um exemplo

| Transação | Destino | Conta_1 | Conta_2 |
| --------- | ------- | ------- | ------- |
| 1         | Viagem  | 200.90  | 3000.00 |
| 0         | Lazer   | 542.27  |         |
| 1         | Moradia | 1300.00 | 700.15  |

Note que os valores sempre serão tratados como removidos (negativos) da `Origem`.
Caso o valor esteja vazio, o conversor irá ignorar a transação para essa relação.

A coluna de transação serve para auxiliar a distinguir o `Destino` entre categorias e transações para outra conta

## TODO

- [X] Modelar arquivo de excel
- [ ] Ler arquivo de excel e extrair infromações necessárias
- [X] Parser para QIF