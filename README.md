# wavy

A high-level programming language designed for the manipulation of audio files

## The Team

| Vishruth Devan | vd2461 | <vd2461@columbia.edu> |
| -------------- | ------ | --------------------- |

## Installation and Usage

1. Install Docker from <https://docs.docker.com/get-started/get-docker/>
2. Clone the repository:  

    ```bash
    git clone https://github.com/vishruthdevan/wavy.git
    ```

### Test the Lexer

- Run the `test-lexer.sh` script:  

   ```bash
   ./test-lexer.sh /wavy/<path-to-sample-file.vy> 
   ```

- If you get a permission denied error while running the script, run the following command and try again:

   ```bash
   chmod 755 test-lexer.sh
   ```

- Example Usage:

   ```bash
   ./test-lexer.sh /wavy/lexer/samples/sample_1.vy
   ```

- The expected outputs are the `.out` files in the `lexer/samples/expected_outputs/` directory. Running the script will generate `.out` files in the same directory as the input file. For example, if `/wavy/lexer/samples/sample_1.vy` was the input, the output will be written to `/wavy/lexer/samples/sample_1.vy.out`.

### Test the Parser

- Run the `test-parser.sh` script:  

   ```bash
   ./test-parser.sh /wavy/<path-to-sample-file.vy> 
   ```

- If you get a permission denied error while running the script, run the following command and try again:

   ```bash
   chmod 755 test-parser.sh
   ```

- Example Usage:

   ```bash
   ./test-parser.sh /wavy/parser/samples/sample_1.vy
   ```

- The expected outputs are the `.out` files in the `parser/samples/expected_outputs/` directory. Running the script will generate `.out` files in the same directory as the input file. For example, if `/wavy/parser/samples/sample_1.vy` was the input, the output will be written to `/wavy/parser/samples/sample_1.vy.out`.

- sample_1.vy and sample_4.vy have been intentionally modified to make the parser identify errors.


### Test the Compiler

- Run the `test-compiler.sh` script:  

   ```bash
   ./test-compiler.sh /wavy/<path-to-sample-file.vy> 
   ```

- If you get a permission denied error while running the script, run the following command and try again:

   ```bash
   chmod 755 test-compiler.sh
   ```

- Example Usage:

   ```bash
   ./test-compiler.sh /wavy/compiler/samples/sample_1.vy
   ```

- The expected outputs are the `.out` files in the `compiler/samples/expected_outputs/` directory. Running the script will generate `.out` files in the same directory as the input file. For example, if `/wavy/compiler/samples/sample_1.vy` was the input, the output will be written to `/wavy/compiler/samples/sample_1.vy.out`.

### Test the VM

- Run the `test-vm.sh` script:  

   ```bash
   ./test-vm.sh /wavy/<path-to-sample-file.vy> 
   ```

- If you get a permission denied error while running the script, run the following command and try again:

   ```bash
   chmod 755 test-vm.sh
   ```

- Example Usage:

   ```bash
   ./test-vm.sh /wavy/vm/samples/sample_1.vy
   ```

- The expected outputs are the `.out` files in the `vm/samples/expected_outputs/` directory. Running the script will generate `.out` files in the same directory as the input file. For example, if `/wavy/vm/samples/sample_1.vy` was the input, the output will be written to `/wavy/vm/samples/sample_1.vy.out`.

- sample_1.vy.incorrect has been intentionally modified to make the parser identify errors.

## Lexical Grammar Definition

### 1. Keywords (`KEYWORD`)

These are reserved words with specific meanings that cannot be used as identifiers.

**Keywords**: `function, return, if, else, true, false, null, for, in, load, export`

**Rules:**

- Exact match with the string in the `keywords` map.
- Case-sensitive.

### 2. Identifiers (`IDENTIFIER`)

An identifier represents variables, functions, or other user-defined or built-in names.

**Valid Characters:**

- Letters (`a-z`, `A-Z`)
- Digits (`0-9`)
- Underscore (`_`)
- Dollar sign (`$`)

**Rules:**

- Cannot start with a digit.
- Example: `foo`, `bar_123`, `$myVar`.

### 3. Numbers

**Types:**

- **INTEGER**: Whole numbers without a decimal point (e.g., `42`).
- **FLOAT**: Numbers with a decimal point (e.g., `3.14`).

**Rules:**

- **INTEGER**: Sequence of digits.
- **FLOAT**: Contains a decimal point and digits on both sides of the decimal point.

### 4. Strings (`STRING`)

A string is a sequence of characters enclosed in double quotes (`"`) or single quotes (`'`).

**Rules:**

- The string must be enclosed in a matching pair of quotes.
- Can contain any unicode character.
- Example: `"Hello, World!"`

### 5. Operators and Symbols

| **Operator** | **Token**    | **Description**       |
| ------------ | ------------ | --------------------- |
| `=`          | `ASSIGN`     | Assignment            |
| `+`          | `PLUS`       | Addition              |
| `-`          | `MINUS`      | Subtraction/Negation  |
| `*`          | `ASTERISK`   | Multiplication        |
| `/`          | `SLASH`      | Division              |
| `!`          | `BANG`       | Logical NOT           |
| `<`          | `LT`         | Less than             |
| `>`          | `GT`         | Greater than          |
| `==`         | `EQUALS`     | Equality comparison   |
| `!=`         | `NOT_EQUALS` | Inequality comparison |

### 6. Punctuation

| **Symbol** | **Token**   | **Description**      |
| ---------- | ----------- | -------------------- |
| `,`        | `COMMA`     | Separator            |
| `;`        | `SEMICOLON` | Statement terminator |
| `:`        | `COLON`     | Type separator       |
| `(`        | `LPR`       | Left parenthesis     |
| `)`        | `RPR`       | Right parenthesis    |
| `{`        | `LBRACE`    | Left brace           |
| `}`        | `RBRACE`    | Right brace          |
| `[`        | `LBRACKET`  | Left bracket         |
| `]`        | `RBRACKET`  | Right bracket        |

### 7. End of File (`EOF`)

- **Token:** `EOF`  
- **Description:** Signals the end of the input stream.

### 8. Illegal Characters (`ILLEGAL`)

- **Token:** `ILLEGAL`  
- **Description:** Any unrecognized or invalid character.

## Lexer Sequence

1. **Initialization**: The lexer starts with the given input and sets up the necessary positions (line, column, etc.).

2. **Reading Characters**: It moves one character at a time, advancing through the input.

3. **Skipping Whitespace**: If it encounters spaces, tabs, or newlines, it skips them until it finds a meaningful character.

4. **Identifying Tokens**:  

   - For each character, it checks if it matches a known token.
   - For complex tokens (like `==`), it looks ahead (peeks) to decide if it’s a multi-character token.

5. **Handling Identifiers and Keywords**:  

   - If it encounters a letter or valid character (like `_`), it reads an entire word.
   - It then checks if the word is a keyword or just a regular identifier.

6. **Handling Numbers**:  

   - If it finds a digit, it reads a full numeric sequence, supporting both integers and floats.
   - If it detects an invalid number, it raises an error.

7. **Reading Strings**:  

   - For strings enclosed in `"` or `'`, it reads until it finds the closing quote or raises an error for an unterminated string.

8. **Error Handling**:  

   - An error is thrown for situations where an invalid number, unterminated string or illegal character is detected.
   - The lexer only reports the error and continues with scanning the rest of the input.

9. **End of Input**:  

   - When the lexer reaches the end of the input, it emits an `EOF` (End of File) token.

10. **Returning Tokens**:  

    - For each token identified, the lexer advances and returns it to the caller for further processing.

## Lexical Error Specification

| **Error Type**          | **Description**                                                | **Example Input**  | **Error Message**                                            |
| ----------------------- | -------------------------------------------------------------- | ------------------ | ------------------------------------------------------------ |
| **Illegal Character**   | Encountered an unrecognized or invalid character.              | `^foo = 10`        | `Lexical error at line 1, position 1: Illegal character "@"` |
| **Unterminated String** | A string literal is not properly closed with a matching quote. | `"hello`           | `Lexical error at line 1, position 7: Unterminated string`   |
| **Invalid Number**      | Incorrect number format detected (e.g., multiple dots).        | `12.34.`, `123abc` | `Lexical error at line 1, position 6: Invalid number`        |

## Context Free Grammar

### Non-Terminals

1. Program: Represents the full program consisting of statements.
2. Statement: Represents a general statement (e.g., assignment, return, expression).
3. Expression: Represents expressions including arithmetic operations, boolean values, function calls, etc.
4. Block: Represents a block of statements, often used in function bodies, if statements, and for loops.
5. Literal: Represents literals like integers, floats, strings, and boolean values.

### Terminals

1. IDENTIFIER: Identifier (variable/function name).
2. ASSIGN: =
3. RETURN: return
4. IF, ELSE, FOR: Keywords for control flow.
5. PLUS, MINUS, ASTERISK, SLASH: Arithmetic operators (+, -, *, /).
6. LPAREN, RPAREN: Parentheses ((, )).
7. LBRACE, RBRACE: Braces ({, }).
8. SEMICOLON: ;
9. INT_LITERAL, FLOAT_LITERAL, STRING_LITERAL, BOOL_LITERAL: Various literal types.

### Production Rules

```plaintext
<Program> → <StatementList>

<StatementList> → <Statement> <StatementList>
                  | ε

<Statement> → <ExpressionStatement>
              | <AssignmentStatement>
              | <ReturnStatement>
              | <IfStatement>
              | <ForLoopStatement>
              | <FunctionDeclaration>

<ExpressionStatement> → <Expression> SEMICOLON

<AssignmentStatement> → IDENTIFIER ASSIGN <Expression> SEMICOLON

<ReturnStatement> → RETURN <Expression> SEMICOLON

<IfStatement> → IF LPAREN <Expression> RPAREN LBRACE <Block> RBRACE
               | IF LPAREN <Expression> RPAREN LBRACE <Block> RBRACE ELSE LBRACE <Block> RBRACE

<ForLoopStatement> → FOR LPAREN <Expression> RPAREN LBRACE <Block> RBRACE

<FunctionDeclaration> → IDENTIFIER LPAREN <ParameterList> RPAREN LBRACE <Block> RBRACE

<ParameterList> → IDENTIFIER <ParameterListTail>
                 | ε

<ParameterListTail> → COMMA IDENTIFIER <ParameterListTail>
                      | ε

<Block> → <StatementList>

<Expression> → <Literal>
               | IDENTIFIER
               | <PrefixExpression>
               | <InfixExpression>
               | <FunctionCall>
               | <GroupedExpression>
               | <ArrayLiteral>
               | <IndexExpression>

<PrefixExpression> → (BANG | MINUS) <Expression>

<InfixExpression> → <Expression> (PLUS | MINUS | ASTERISK | SLASH) <Expression>

<FunctionCall> → IDENTIFIER LPAREN <ArgumentList> RPAREN

<ArgumentList> → <Expression> <ArgumentListTail>
                 | ε

<ArgumentListTail> → COMMA <Expression> <ArgumentListTail>
                     | ε

<GroupedExpression> → LPAREN <Expression> RPAREN

<ArrayLiteral> → LBRACKET <ExpressionList> RBRACKET

<ExpressionList> → <Expression> <ExpressionListTail>
                   | ε

<ExpressionListTail> → COMMA <Expression> <ExpressionListTail>
                       | ε

<IndexExpression> → <Expression> LBRACKET <Expression> RBRACKET

<Literal> → INT_LITERAL
            | FLOAT_LITERAL
            | STRING_LITERAL
            | BOOL_LITERAL
```

## Parsing Algorithm

- The Wavy programming language employs recursive descent parsing combined with Pratt parsing specifically for expression evaluation. This design choice leverages recursive descent parsing to provide a clear and modular approach to syntax analysis, where each grammar rule is represented by a function, enabling easy readability and maintainability of the parser. 
- For expressions, Wavy uses Pratt parsing, which allows flexible handling of operator precedence and associativity, making it well-suited for parsing complex expressions efficiently. This hybrid approach ensures that Wavy's syntax and expression parsing are both intuitive and powerful, facilitating robust language processing.

## Demo Video about Parsing

URL: [https://youtu.be/WfligR-tuQg](https://youtu.be/WfligR-tuQg)

## Compiler and VM Specification

### Intermediate Representation (IR)

- After the source code is parsed, the output is passed to the compiler, which generates an **Intermediate Representation (IR)** of the code. The IR consists of instructions formatted as:
  - `<StackPosition> <OpCode> <Argument>`
  - Example: `0025 OpConstant 10` or `0046 OpArray 13`.
- Each instruction operates on a single argument, making the IR simple and optimizing-friendly.

### Virtual Machine Execution

- A **virtual machine (VM)** reads the IR instructions and processes them sequentially. It uses a virtual stack to evaluate each instruction, managing the call stack during execution.
- The VM continues until all instructions are processed, and the stack returns to its initial state.

### Variable Scoping

- In the **Wavy programming language**, variables are scoped within braces `{}` and can only be accessed within the scope they are defined.
- A **symbol table** is used to manage variables, and they are "popped" from the stack when their scope ends.

### Error Handling

- The compiler includes error handling to manage issues like:
  - **Unknown operators**
  - **Undefined variables**
- These errors are detected during compilation and flagged accordingly.

This structure ensures that the code is executed efficiently, supports variable scoping, and allows for optimized compilation processes.

## Demo Video about Compiling

URL: [https://youtu.be/yAqygMO68q8](https://youtu.be/yAqygMO68q8)
