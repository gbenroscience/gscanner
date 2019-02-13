# gscanner

This is a simple Golang library that helps you quickly scan a string and split it into 
substrings based on an array of supplied tokens.
 
This library benchmarks very fast and is stable.

There are times when you either can do without the overhead of regular expressions, or the tokens required to 
split a string are a finite number.

There is no need to resort to regular expressions in this case.

This simple library lends itself as an hi-speed scanner/splitter and returns an array containing the substrings of the original
string. Whether you would love to retain the splitting tokens in the scanner's output is totally up to you! 

Simply set the ```IncludeTokensInOutput``` property of your ```GScanner``` to true to retain the splitting tokens.
Else set it to false.



# Example

A beautiful example usage would be for scanning an arithmetic expression or other expression into a form wherein the input tokens are very visible:

```golang

input := "(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)" 
	
scanner := &scanner.GScanner{ Input: input, 
                              Tokens: []string{"-", "sin", "sinh", "+", "(", ")", "cos"}, 
			      IncludeTokensInOutput: true, 
			      }
	arr:= scanner.Scan()
```
	



  The output would be:
  
  
  ```golang
   [(, 28, +, 32, +, 11, -, 9E12, +, sin, (, 3.2E9/, cos, (, -, 3, ), ), -, sin, sinh, (, 5, ), +, sinh, (, 8, )]
   ```



 

The square braces and the commas are just for formatting.

If you set 
```golang 
IncludeTokensInOutput
``` 
to false by doing this:

	```golang
	 scanner := &scanner.GScanner{
		Input:                 input,
		Tokens:                []string{"-", "sin", "sinh", "+", "(", ")", "cos"},
		IncludeTokensInOutput: false,
	}```
  
  Then the output is:
  
 ```golang
[28, 32, 11, 9E12, 3.2E9/, 3, 5, 8]
```

Enjoy!

