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

3. Run the `test-lexer.sh` script:  

    ```bash
    ./test-lexer.sh /wavy/<path-to-sample-file.vy> 
    ```

    If you get a permission denied error while running the script, run the following command and try again:

   ```bash
   chmod 755 test-lexer.sh
   ```

4. Example Usage:

    ```bash
    ./test-lexer.sh /wavy/samples/sample_1.vy
    ```

5. The expected outputs are the `.out` files in the `samples/expected_outputs/` directory. Running the script will generate `.out` files in the same directory as the input file. For example, if `/wavy/samples/sample_1.vy` was the input, the output will be written to `/wavy/samples/sample_1.vy.out`.

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
