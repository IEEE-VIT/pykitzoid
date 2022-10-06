package main

// import whatever packages you will require here
import (
  "encoding/csv"
  "strconv"
  "os"
  "log"
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

// function to read a csv and return the csv reader object
func read_csv(filename string) ([][]float32, error) {
  // Read the file into a string
  // if error occurs, return an empty slice and the error
  f, err := os.Open(filename)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  // Create a new reader to evaluate the file as a csv
  // if error occurs, return an empty slice and the error
  lines, err := csv.NewReader(f).ReadAll()
  if err != nil {
    return nil, err
  }

  //convert the string array to float array
  var lines_converted [][]float32
  for i:=0; i<len(lines); i++ {
    temp1, err := strconv.ParseFloat(lines[i][0], 32)
    if err != nil {
      log.Fatal(err)
    }
    temp2, err := strconv.ParseFloat(lines[i][1], 32)
    if err != nil {
      log.Fatal(err)
    }
    lines_converted = append(lines_converted, []float32{float32(temp1), float32(temp2)})
  }
  return lines_converted, nil
}

// function to calculate the mean of a specified column
func mean_of_col(arr [][]float32, col_number int) (float32) {
  // calculate the sum of the array
  var calculated_mean_value float32
  calculated_mean_value = 0
  for i:=0; i<len(arr); i++ {
    calculated_mean_value += arr[i][col_number]
  }
  calculated_mean_value = calculated_mean_value / float32(len(arr))
  return calculated_mean_value
}


// function to calculate slope m and intercept c using the required formula

// // FORMULA
// // m = (Σ(x - x')(y - y'))/(Σ(x-x')^2)
// // where x' and y' are the means of x and y respectively

func calculate_slope_and_intercept(csv_object [][]float32) (float32, float32) {
  var numerator, denominator float32
  numerator = 0
  denominator = 0

  // calculate the mean of x and y
  var x_mean, y_mean float32
  x_mean = mean_of_col(csv_object, 0)
  y_mean = mean_of_col(csv_object, 1)

  // calculate the numerator and denominator
  for i:=0; i<len(csv_object); i++ {
    numerator += (csv_object[i][0] - x_mean) * (csv_object[i][1] - y_mean)
    denominator += (csv_object[i][0] - x_mean) * (csv_object[i][0] - x_mean)
  }

  // calculate the slope and intercept
  var slope, intercept float32
  slope = numerator / denominator
  intercept = y_mean - slope * x_mean

  return slope, intercept
}

func predict_y(x float32, slope float32, intercept float32) float32 {
  var y_pred = slope * x + intercept
  return y_pred
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

func calculate_r_squared(csv_object [][]float32, slope float32, intercept float32) float32 {
  var y_mean = mean_of_col(csv_object, 1)
  var numerator, denominator float32

  for i:=0; i<len(csv_object); i++ {
    var y_pred = predict_y(csv_object[i][0], slope, intercept)
    numerator += (y_pred - y_mean) * (y_pred - y_mean)
    denominator += (csv_object[i][1] - y_mean) * (csv_object[i][1] - y_mean)
  }

  var r_squared = numerator / denominator
  return r_squared
}

// MAIN FUNCTION

func main() {
  // define the csv file path
  var filepath string
  filepath = "sample_data.csv"

  // read the csv file into a file reader object
  var csv_object, err = read_csv(filepath)

  // if error occurs, print the error and exit
  if err != nil {
    log.Fatal(err)
  }

  // calculate the slope and intercept
  var slope, intercept = calculate_slope_and_intercept(csv_object)
  
  // calculate r^2 value
  var r_squared = calculate_r_squared(csv_object, slope, intercept)

  // print the slope, intercept and r^2 value
  log.Println("slope: ", slope)
  log.Println("intercept: ", intercept)
  log.Println("r^2: ", r_squared)
}
