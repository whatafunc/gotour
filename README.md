### Intro
Hey, I am having fun going throught the Go Tour and practicing some CI/CD with git flow and prior local tests for building and running the code against it's _test file

### How to use locally
chmod +x test.sh 
./test.sh 


### Expected output:
🔍 Testing GoTour Exercises Repository
======================================

📦 Step 1: Checking all Go files compile...
-------------------------------------------
✅ ./exercises/ex2-wordcount/use_strings_Fields compiles as package
✅ ./exercises/ex2-wordcount/used_IsLetter compiles as package
✅ ./exercises/ex2-wordcount/used_IsLetter compiles as package
✅ ./exercises/ex2-wordcount/use_strings_split compiles as package
✅ ./exercises/ex1-pic compiles as package
✅ ./practice/tic-tac-toe compiles as package
✅ ./practice/fibonaci compiles as package
✅ ./practice/fibonaci compiles as package

🔨 Step 2: Building all main packages...
---------------------------------------
Building ./exercises/ex2-wordcount/use_strings_Fields... ✅ ./exercises/ex2-wordcount/use_strings_Fields builds successfully
Building ./exercises/ex2-wordcount/used_IsLetter... ✅ ./exercises/ex2-wordcount/used_IsLetter builds successfully
Building ./exercises/ex2-wordcount/use_strings_split... ✅ ./exercises/ex2-wordcount/use_strings_split builds successfully
Building ./exercises/ex1-pic... ✅ ./exercises/ex1-pic builds successfully
Building ./practice/tic-tac-toe... ✅ ./practice/tic-tac-toe builds successfully
Building ./practice/fibonaci... ✅ ./practice/fibonaci builds successfully

🧪 Step 3: Running available tests...
------------------------------------
Testing ./exercises/ex2-wordcount/used_IsLetter... === RUN   Test_main
Text: "I am learning Go!"
Count: map[am:1 go:1 i:1 learning:1]

Text: "Hello world hello"
Count: map[hello:2 world:1]

Text: "Multiple   spaces    here"
Count: map[here:1 multiple:1 spaces:1]

Text: ""
Count: map[]

Text: "Punctuation! Should, work? Right."
Count: map[punctuation:1 right:1 should:1 work:1]

--- PASS: Test_main (0.00s)
PASS
ok      github.com/whatafunc/gotour/exercises/ex2-wordcount/used_IsLetter       0.467s
✅ ./exercises/ex2-wordcount/used_IsLetter tests passed
Testing ./practice/fibonaci... === RUN   Test_main
--- PASS: Test_main (0.00s)
PASS
ok      github.com/whatafunc/gotour/practice/fibonaci   0.462s
✅ ./practice/fibonaci tests passed

📊 Step 4: Summary
------------------
Total: 16 passed, 0 failed, 0 skipped
🎉 All checks passed!