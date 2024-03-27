# Tic Tac Toe

A console version of the classic game written in Go.

Two players, each assigned a unique mark (X or O) alternate placing their mark
on the board. The game ends when a player places their mark in three consecutive
squares forming a horizontal, vertical, or diagonal line or when there are no
moves left.

# Dependencies

Go must be installed on your system since the game requires building from
source.

# Running the Game

1. Clone the repository

   ```text
   git clone https://github.com/freddiehaddad/tictactoe.git
   ```

1. Navigate to the repository directory

   ```text
   cd tictactoe
   ```

1. Launch

   ```text
   go run ./bin/tictactoe.go
   ```

## Example Play

```text
Welcome to Tic Tac Toe!
     0   1   2
   ┏━━━┯━━━┯━━━┓
 A ┃   │   │   ┃
   ┠───┼───┼───┨
 B ┃   │   │   ┃
   ┠───┼───┼───┨
 C ┃   │   │   ┃
   ┗━━━┷━━━┷━━━┛
X: Type a row and column for your mark (e.g. B2): a0
     0   1   2
   ┏━━━┯━━━┯━━━┓
 A ┃ X │   │   ┃
   ┠───┼───┼───┨
 B ┃   │   │   ┃
   ┠───┼───┼───┨
 C ┃   │   │   ┃
   ┗━━━┷━━━┷━━━┛
O: Type a row and column for your mark (e.g. B2): b0
     0   1   2
   ┏━━━┯━━━┯━━━┓
 A ┃ X │   │   ┃
   ┠───┼───┼───┨
 B ┃ O │   │   ┃
   ┠───┼───┼───┨
 C ┃   │   │   ┃
   ┗━━━┷━━━┷━━━┛
X: Type a row and column for your mark (e.g. B2): b1
     0   1   2
   ┏━━━┯━━━┯━━━┓
 A ┃ X │   │   ┃
   ┠───┼───┼───┨
 B ┃ O │ X │   ┃
   ┠───┼───┼───┨
 C ┃   │   │   ┃
   ┗━━━┷━━━┷━━━┛
O: Type a row and column for your mark (e.g. B2): c0
     0   1   2
   ┏━━━┯━━━┯━━━┓
 A ┃ X │   │   ┃
   ┠───┼───┼───┨
 B ┃ O │ X │   ┃
   ┠───┼───┼───┨
 C ┃ O │   │   ┃
   ┗━━━┷━━━┷━━━┛
X: Type a row and column for your mark (e.g. B2): c2
     0   1   2
   ┏━━━┯━━━┯━━━┓
 A ┃ X │   │   ┃
   ┠───┼───┼───┨
 B ┃   │ X │   ┃
   ┠───┼───┼───┨
 C ┃   │   │ X ┃
   ┗━━━┷━━━┷━━━┛
X wins!
     0   1   2
   ┏━━━┯━━━┯━━━┓
 A ┃   │   │   ┃
   ┠───┼───┼───┨
 B ┃   │   │   ┃
   ┠───┼───┼───┨
 C ┃   │   │   ┃
   ┗━━━┷━━━┷━━━┛
X: Type a row and column for your mark (e.g. B2): q

Exiting!
```
