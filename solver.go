package main

import (
  "fmt"
  "strconv"
  "strings"
)

type CellID struct{ row, col int }

type Cell struct {
  ID           CellID
  Solution     int
  solved       bool
  alternatives []int
}

func main() {
  var square [9][9]Cell;

  //square = CreateSquare();

  data := `
    _4257__89
    __6______
    53_9__6__
    _7__3_1__
    _________
    __3_4__2_
    __7__3_96
    ______5__
    61__2943_
    `;

  //data := `
  //___3__4__
  //9___1____
  //4_8_7_5_3
  //___6__9__
  //___157___
  //__5__2___
  //2_4___6_5
  //____2___8
  //__67_8___
  //  `

  square = parser(data);

  PrintSquare(square);
  //printAlternatives(square);

  solution, solved := SolveSquare(square);
  if (solved) {
    fmt.Println("solved")
  } else {
    fmt.Println("impossible")
  }
  PrintSquare(solution)
}

func CreateSquare() [9][9]Cell {
  var square [9][9]Cell;
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      square[i][j] = Cell{
        ID: CellID{i, j},
        Solution:-1,
        solved:false,
        alternatives:[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
      };
    }
  }
  return square;
}

func PrintSquare(square [9][9]Cell) {
  for i, row := range square {
    for j, cell := range row {
      if (cell.solved) {
        fmt.Printf(" %d ", cell.Solution)
      } else {
        fmt.Print(" _ ")
      }
      if j == 2 || j == 5 {
        fmt.Print("|")
      }
    }
    fmt.Print("\n");
    if i == 2 || i == 5 {
      fmt.Println("---------|---------|---------")
    }
  }
}

func SolveSquare(square [9][9]Cell) ([9][9]Cell, bool) {
  for true {
    cell, min, found := lowerCellAlternatives(square)

    if !found {
      return square, true;
    }

    if min >= 2 {
      for _, alternative := range cell.alternatives {
        squareTemp, solved := SolveSquare(solveCell(square, cell.ID.row, cell.ID.col, alternative))
        if solved {
          return squareTemp, true
        }
      }
      return square, false
    }

    if ( min == 1) {
      square = solveCell(square, cell.ID.row, cell.ID.col, cell.alternatives[0]);
    }

    if (min == 0) {
      return square, false;
    }
  }

  return square, false;
}

func printAlternatives(square [9][9]Cell) {
  for _, row := range square {
    for _, cell := range row {
      fmt.Print(cell.alternatives)
      for i := len(cell.alternatives); i < 9; i++ {
        fmt.Print("  ")
      }
      if len(cell.alternatives) != 0 {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n");
  }
}

func solveCell(square [9][9]Cell, row int, col int, solution int) [9][9]Cell {
  square[row][col].Solution = solution;
  square[row][col].solved = true;
  square[row][col].alternatives = nil;

  for j := 0; j < 9; j++ {
    square[row][j] = removeAlternative(square[row][j], solution)
  }

  for i := 0; i < 9; i++ {
    square[i][col] = removeAlternative(square[i][col], solution)
  }

  for _, cellID := range subSquare(row, col) {
    square[cellID.row][cellID.col] = removeAlternative(square[cellID.row][cellID.col], solution)
  }

  return square;
}

func removeAlternative(cell Cell, alternative int) Cell {
  var alternatives []int;
  for _, n := range cell.alternatives {
    if n != alternative {
      alternatives = append(alternatives, n);
    }
  }
  cell.alternatives = alternatives;
  return cell;
}

func lowerCellAlternatives(square [9][9]Cell) (Cell, int, bool) {
  min := 999;
  var lowerCell Cell;
  found := false;

  for _, row := range square {
    for _, cell := range row {
      if (!cell.solved && len(cell.alternatives) < min) {
        min = len(cell.alternatives);
        lowerCell = cell;
        found = true;
      }
    }
  }
  return lowerCell, min, found;
}

func between(min int, number int, max int) bool {
  return min <= number && number <= max;
}

func subSquare(row int, col int) [9]CellID {
  switch {
  //first col
  case between(0, row, 2) && between(0, col, 2):
    return [9]CellID{
      CellID{0, 0}, CellID{0, 1}, CellID{0, 2},
      CellID{1, 0}, CellID{1, 1}, CellID{1, 2},
      CellID{2, 0}, CellID{2, 1}, CellID{2, 2},
    }
  case between(3, row, 5) && between(0, col, 2):
    return [9]CellID{
      CellID{3, 0}, CellID{3, 1}, CellID{3, 2},
      CellID{4, 0}, CellID{4, 1}, CellID{4, 2},
      CellID{5, 0}, CellID{5, 1}, CellID{5, 2},
    }
  case between(6, row, 8) && between(0, col, 2):
    return [9]CellID{
      CellID{6, 0}, CellID{6, 1}, CellID{6, 2},
      CellID{7, 0}, CellID{7, 1}, CellID{7, 2},
      CellID{8, 0}, CellID{8, 1}, CellID{8, 2},
    }

  //center col
  case between(0, row, 2) && between(3, col, 5):
    return [9]CellID{
      CellID{0, 3}, CellID{0, 4}, CellID{0, 5},
      CellID{1, 3}, CellID{1, 4}, CellID{1, 5},
      CellID{2, 3}, CellID{2, 4}, CellID{2, 5},
    }
  case between(3, row, 5) && between(3, col, 5):
    return [9]CellID{
      CellID{3, 3}, CellID{3, 4}, CellID{3, 5},
      CellID{4, 3}, CellID{4, 4}, CellID{4, 5},
      CellID{5, 3}, CellID{5, 4}, CellID{5, 5},
    }
  case between(6, row, 8) && between(3, col, 5):
    return [9]CellID{
      CellID{6, 3}, CellID{6, 4}, CellID{6, 5},
      CellID{7, 3}, CellID{7, 4}, CellID{7, 5},
      CellID{8, 3}, CellID{8, 4}, CellID{8, 5},
    }

  //right col
  case between(0, row, 2) && between(6, col, 8):
    return [9]CellID{
      CellID{0, 6}, CellID{0, 7}, CellID{0, 8},
      CellID{1, 6}, CellID{1, 7}, CellID{1, 8},
      CellID{2, 6}, CellID{2, 7}, CellID{2, 8},
    }
  case between(3, row, 5) && between(6, col, 8):
    return [9]CellID{
      CellID{3, 6}, CellID{3, 7}, CellID{3, 8},
      CellID{4, 6}, CellID{4, 7}, CellID{4, 8},
      CellID{5, 6}, CellID{5, 7}, CellID{5, 8},
    }
  case between(6, row, 8) && between(6, col, 8):
    return [9]CellID{
      CellID{6, 6}, CellID{6, 7}, CellID{6, 8},
      CellID{7, 6}, CellID{7, 7}, CellID{7, 8},
      CellID{8, 6}, CellID{8, 7}, CellID{8, 8},
    }

  //default
  default:
    return [9]CellID{};
  }
}

func parser(text string) [9][9]Cell {
  square := CreateSquare();
  text = strings.TrimSpace(text);

  row, col := 0, 0;

  for _, char := range text {
    switch string(char){
    case "1", "2", "3", "4", "5", "6", "7", "8", "9":
      sol, _ := strconv.Atoi(string(char))
      square = solveCell(square, row, col, sol);
      col++

    case "\n":
      row++
      col = 0

    case "_":
      col++
    }
  }

  return square
}
