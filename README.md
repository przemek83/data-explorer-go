[![Build & test](https://github.com/przemek83/data-explorer-go/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/przemek83/data-explorer-go/actions/workflows/build-and-test.yml)
[![CodeQL](https://github.com/przemek83/data-explorer-go/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/przemek83/data-explorer-go/actions/workflows/github-code-scanning/codeql)
[![codecov](https://codecov.io/github/przemek83/data-explorer-go/graph/badge.svg?token=C49CYPL3LL)](https://codecov.io/github/przemek83/data-explorer-go)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=przemek83_data-explorer-go&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=przemek83_data-explorer-go)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=przemek83_data-explorer-go&metric=bugs)](https://sonarcloud.io/summary/new_code?id=przemek83_data-explorer-go)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=przemek83_data-explorer-go&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=przemek83_data-explorer-go)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=przemek83_data-explorer-go&metric=coverage)](https://sonarcloud.io/summary/new_code?id=przemek83_data-explorer-go)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=przemek83_data-explorer-go&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=przemek83_data-explorer-go)

# Table of content
- [Problem description](#problem-description)
- [Usage](#usage)
- [Input data format](#input-data-format)
- [License](#license)

# About Project
Small tool for aggregating and grouping data. Written in Go, mimicking functionality of my older data-explorer project which was written in C++. Created to learn Go better, exercise TDD and have some fun.

# Problem description
For given input data allow calculating average, minimum and maximum taking into consideration grouping column.

# Building
First you need to download repo to you machine. Make sure you have go installed and version is greater or equal to `1.22`. Having that done use following command to build:
```
go build cmd/data-explorer/data-explorer.go
```
When executed in root directory of repository, it should create binary called `data-explorer`.

# Usage 
`data-explorer file {avg,min,max} aggregation grouping`  
Where:  
+ `file` - name of file with data to load,  
+ `{avg,min,max}` - type of operation, use one of those,  
+ `aggregation` - name of column used for aggregating data,  
+ `grouping` - name of column used for grouping data.

Example usage:  
`data-explorer sample.txt avg score first_name`  

Example output:
```
Data loaded in 0.000000s
Operation completed in 0.000000s
Results:
map[dave:8 tamas:5.5 tim:8]
```

# Input data format
Input data need to have following structure:  
```
<column 1 name>;<column 2 name>;<column 3 name>  
<column 1 type>;<column 2 type>;<column 3 type>  
<data 1 1>;<data 2 1>;<data 3 1> 
...  
<data 1 n>;<data 2 n>;<data 3 n> 
```
Where column type can be `string` or `integer`.  

Example data:
```
first_name;age;movie_name;score
string;integer;string;integer
tim;26;inception;8
tim;26;pulp_fiction;8
tamas;44;inception;7
tamas;44;pulp_fiction;4
dave;0;inception;8
dave;0;ender's_game;8
```
Not too flexible format of data was used for simplicity of parsing.

# Testing
To executre tests manually, open root directory of repository and run command:
```
go test -v ./...
```
It should generate similar output:

    $ go test -v ./...
    ?   	data-explorer/cmd/data-explorer	[no test files]
    === RUN   TestCalculatorExecute
    === RUN   TestCalculatorExecute/Max_age_grouped_by_movie_name
    === RUN   TestCalculatorExecute/Max_score_grouped_by_movie_name
    
    (...)

    --- PASS: TestMakeQuery/Args_list_too_short,_avg_operation,_proper_columns. (0.00s)
    --- PASS: TestMakeQuery/Proper_number_of_args,_avg_operation,_wrong_aggregate_column. (0.00s)
    --- PASS: TestMakeQuery/Proper_number_of_args,_avg_operation,_wrong_grouping_column. (0.00s)

    PASS
    ok  	data-explorer/internal	0.004s

# License
Project is distributed under the MIT License. See `LICENSE` for more information.