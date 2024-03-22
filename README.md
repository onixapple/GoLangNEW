# GoLangNEW

Main Go Lang use cases: 

Applications runing on scaled, distributed systems.

Applications that require concurrent actions.
Running on multiple cores
Replacement for Java and C# wich have concurency mechanism built in but are at a higher cost and slower code.


Why Go Lang

Simple and readable syntax like Python
Efficient and safe at a lower level like C++

Used on server-side (microservices , dbservices)


Write GO :

1. Install GO (comes with GO CLI)

2. IDE  (VSC VS )


Commands: 
go mod init package_name

go run filename

Packages: fmt, strings, strconv

Variables defining types: var, const

Formating: Printf %v , %b, %c %X

Escape characters \n \t 

Data types: Strings, Integers, Bool, Map, Array    var userName string 

Specific data types : uint , int8 int16 int32 - ---- type check, safety, Pointers


Scan - require uesr imput - must be a pointer to scan

Arrays have fixed size  :   var books = [50]string{empty} or {"name1", "name2"}

Slices are an abstraction of arrays, more efficient to work with, resized when needed
append(slice, data) - append to the end of slice


Loops: for loop for every usecase, break - ends the loop, contiune - next iteration

Conditionals : IF ; else if , else

Switch

unused variables: _


when code is split in multiple files, you need to run all files at once in the command or with "."  .

packages need their own folders

when importing from another own package, need to specify the folder above

calling like nameOfModule.function

exporting a funciton  -> capitalize the function name

defining a map >  var map = make(map[string]string) unique data types

object :> struct : holds multimple data types. Type UserData struct { }

make the applicaiton concurrent : go 


when main thread is done , the application is done 

need to tell the main thread to wait on other threads

wg = sync.WaitGroup{

}
 
use: wg.Add(1) - adding 1 thread, depending on how many we need 

wg.Wait()

in the function that is on another thread: wg.Done() - removing the thread

What GO does better: When we create a thread he actualy creates a GoRoutine ,  (Green Thread).

Cheaper and lightweight

Less overhead - can use thousands of Goroutines





