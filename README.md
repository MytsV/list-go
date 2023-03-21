# List-go

## Table of contents
* [General description](#general-description)
* [Variant](#variant)
* [Setup and usage](#setup-and-usage)
* [Commit with failing tests](#commit-with-failing-tests)
* [Example](#example)
* [Conclusion](#conclusion)

## General description
A Go package which implements List collection.\
Виконаний як лабораторна робота №2 з предмету "МТРПЗ" Миць Вікторією, група ІМ-12.

## Variant
Номер моєї залікової книжки: 1221\
`1221 % 4 = 1`
| Остача | Початкова реалізація      | Друга реалізація                          |
|--------|---------------------------|-------------------------------------------|
| 1      | Двобічно зв'язаний список | Список на базі вбудованих масивів/списків |

## Setup and usage
To run this project, first clone it onto your local machine.\
Alternatively, you can use [repository's codespace](https://mytsv-super-space-fiesta-7w4rx47qqjwhx497.github.dev/).

Make sure you have [Go](https://go.dev/doc/install) and [Make](https://linuxhint.com/install-use-make-windows//) latest versions installed.

Run the example program:\
`go run ./cmd/example`\
or\
`make out/example`

Run tests:\
`go test`\
or\
`make test`

## Commit with failing tests
[Click](https://github.com/MytsV/list-go/commit/08a013a44aa9a91b40b54ff8431e8320e2d93908)

## Example
Golang by default has example fuction feature. These functions serve both as documentation and tests. Per task, програма повинна містити демонстрацію використання усіх методів класу. Instead of main.go, this demonstration is in [ExampleList() function](https://github.com/MytsV/list-go/blob/d05bf96a1464266c68c91030ea13baff368e6b7e/list_test.go#L548).

## Conclusion
I've decided to try Test First approach. It definitely made the proccess of developing the first implementation much less worrisome, although quite more time-consuming. Unit tests have proved to be more safe and convenient than manual testing of every function even on this stage of development. Test First approach let me account for every exceptional behaviour ahead.

I saw the real power of unit tests when I refactored and subsequently reimplemented List structure. No longer I had to worry about breaking everything with little changes. Even though I spent more time on code initially, **unit tests paid for themselves**. Moreover, their use definitely will save even more time with later refactorings if I work hard to make unit tests themselves easily maintainable.