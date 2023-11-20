# Game of life

## Base rules

The following rules comes from [wikipedia](https://en.wikipedia.org/wiki/Conway's_Game_of_Life)

- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction

## Mutations

Cells can mutate when they are created.
The behaviour of mutating can be described this way:
- A cell is created ( by any means )
- A cell have x percent chance ( pMutate ) to undergo a genetic modification
- If the cell undergo a genetic modification, the same cell have x percent chance ( Attribute.Probability ) to gain the modification
- A cell create from one or several mutatede cells have more chances to undergo a genetic modification

Mutations does not obey exactly to base rules, moreover they aren't stables, they can die ( Attribute.Stability )