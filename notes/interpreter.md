- `Interpreter` is a component that processes structured text data. It does so by turning it into separate lexical token (`lexing`) and then interpreting sequences of said tokens (`parsing)
- Interpreter has two parts -
  - `Lexing` - Converting them into a sequence tokens which can be parsed
  - `Parsing` - Parsing the tokens as an Abstract Syntax Tree (AST) to do business logic onto them.

## Lexing

- Take an example where we want to get the sum of `(13 + 4) - (12 + 1)`
- First, we will need to convert this into tokens like `Plus`, `Left Parantheses`, `Right Parantheses`, `Subtract` and `Integer` so that later, we can parse these tokens to apply business logic.

```go
type TokenType int

const (
 Int TokenType = iota
 Plus
 Minus
 LParen
 RParen
)

type Token struct {
 Type TokenType
 Text string
}

func Lex(input string) []Token {
 var result []Token
 for i := 0; i < len(input); i++ {
  switch input[i] {
  // We convert the +, -, (, ) into their own tokens with text as them only

  default:
   if unicode.IsDigit(rune(input[i])) {
    sb := strings.Builder{}
    for j := i; j < len(input); j++ {
     if unicode.IsDigit(rune(input[j])) {
      sb.WriteRune(rune(input[j]))
      i++
     } else {
      result = append(result, Token{Type: Int, Text: sb.String()})
      i--
      break
     }
    }
    continue
   }
  }
 }

 return result
}
```

- `Token` contains text just so that we can put integers correctly, otherwise text is not required for others.
- It's also useful if we want to reconstruct tokens into the input string.
- Now, the input string has been converted to a list of tokens, we can parse them and apply business logic based on the symbols

## Parsing

- Parsing is working on the tokens to get the required output.

```go
type Element interface {
 Value() int
}

//  implements Element
type Integer struct {
 value int
}

type Operation int

const (
 Addition Operation = iota
 Subtraction
)

//  Implement Element and is the business logic here
type BinaryOperation struct {
 Type        Operation
 Left, Right Element
}
```

- So, we define a `Element` interface which is basically either the input integers or result of a binary operation.
- We have done this because we know that there could be nested Binary Operations
- The `Parse()` function provides logic on how to divide the whole tokens array into meaningful binary operations so that we can get the result.

```go
func Parse(tokens []Token) Element {
 result := BinaryOperation{}
 haveLhs := false
 for i := 0; i < len(tokens); i++ {
  token := &tokens[i]
  switch token.Type {
  case Int:
   n, _ := strconv.Atoi(token.Text)
   integer := NewInteger(n)
   if !haveLhs {
    result.Left = integer
    haveLhs = true
   } else {
    result.Right = integer
   }
   continue

  case Plus:
   result.Type = Addition
   continue

  case Minus:
   result.Type = Subtraction
   continue

  case LParen:
   j := i
   numParantheses := -1
   for ; j < len(tokens); j++ {
    if tokens[j].Type == RParen {
     if numParantheses == 0 {
      break
     } else {
      numParantheses--
     }
    }
    if tokens[j].Type == LParen {
     numParantheses++
    }
   }
   var subexp []Token
   for k := i + 1; k < j; k++ {
    subexp = append(subexp, tokens[k])
   }
   element := Parse(subexp)
   if !haveLhs {
    result.Left = element
    haveLhs = true
   } else {
    result.Right = element
   }
   i = j
  }
 }

 return &result
}
```

- In case of parantheses, we recursively solve the expression inside the parantheses as `BinaryOperation{}` to arrive at a value.
- All the other tokens are either to set the type or the position of the element in a binary operation.
