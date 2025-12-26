# monkey

<밑바닥부터 만드는 interpreter in Go>를 보며 따라 만드는 간단한 언어와 인터프리터

```mermaid
flowchart LR
  A["소스 코드"] --Lexing--> B["토큰"] --Parsing--> C["추상구문트리(AST)"]
```

- ASCII만 지원하며 유니코드는 지원하지 않음