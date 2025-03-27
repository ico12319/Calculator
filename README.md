# Math Calculator
This project is a basic math calculator implemented in Go that uses a recursive descent parser to evaluate arithmetic expressions.
It leverages the Composite design pattern to represent expressions as trees of operations and numeric literals, making it easy to extend and maintain.
## Features
- **Expression Parsing**


 Parse arithmetic expressions into an expression tree using a recursive descent parser, supporting operations such as addition, subtraction, multiplication, division, and exponentiation.


- **Modular Operation Components**


 Each arithmetic operation is encapsulated in its own module, enabling straightforward extension of the calculator with new operators.


- **Utility Functions**


  Implements helper functions for string manipulation and expression validation to streamline the parsing process.
  

## Design Decisions
- **Interface-Driven Architecture**


The Expression interface unifies the evaluation logic, allowing all arithmetic operations and numeric literals to be composed seamlessly.

- **Composite Design Pattern**


 The expression tree is built using the Composite design pattern, where both simple values (leaf nodes) and compound operations (composite nodes) implement the same Expression interface.
