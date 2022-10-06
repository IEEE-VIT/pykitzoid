package main

// import whatever packages you will require here
import (
  "encoding/csv"
  "strconv"
  "fmt"
  "os"
)

// use the following sample dataset
/*
| x     | y     |
|-------|-------|
| 4512  | 1530  |
| 3738  | 1297  |
| 4261  | 1335  |
| 3777  | 1282  |
| 4177  | 1590  |
*/

// expect the dataset to be in csv form

// FUNCTION DEFINITION
// add parameters and change the return type as necessary


// function to calculate the mean of x and y

func mean(csv_file_path string) (float32, float32) {

  var x, y float32
  x = 0
  y = 0
  
  //Check if file path is valid
  fd, error := os.Open(csv_file_path)
  if error != nil {
    fmt.Println(error)
  }
  //fmt.Println("File path valid!")
  defer fd.Close()

  //Read the contents of the CSV file
  fileReader := csv.NewReader(fd)
  records, error := fileReader.ReadAll()
  if error != nil {
    fmt.Println(error)
  }

  //fmt.Println(records)
  for row:=0; row<len(records); row++ {
      x_val, _ := strconv.ParseFloat(records[row][0], 32)
      y_val, _ := strconv.ParseFloat(records[row][1], 32)
      x += float32(x_val)
      y += float32(y_val)
  }
  
  return x/float32(len(records)) , y/float32(len(records))
}

// function to calculate slope m and intercept c using the required formula

// // FORMULA
// // m = (Σ(x - x')(y - y'))/(Σ(x-x')^2)
// // where x' and y' are the means of x and y respectively

func calculate_slope_and_intercept() (float32, float32) {
  // insert code here
}

// function to plot regression line

func plot_regression_line() {
  // insert code here
}

// function to plot the data points

func plot_data_points() {
  // insert code here
}

// function to calculate R-squared value to measure accuracy of model

// // FORMULA
// R^2 = (Σ(y_pred - y')^2)/(Σ(y-y')^2)
// where y_pred is the predicted value for y, y' is the mean

func calculate_r_squared() float32 {
  // insert code here
}

// MAIN FUNCTION

func main() {
  var x_mean, y_mean float32
  x_mean, y_mean = mean("sample_data.csv")
  fmt.Println(x_mean, y_mean)
}
