# pty-go
Repository so as to build Pseudo-Terminals with GoLang.

The first step was to create a loop that handled input management, then it was necessary to create a
command selector and finally an executor. The selector could choose between a series of functions 
defined in the package commands, and we created a special command to exit the application. 
With some refactoring, we went from functions to structs containing both the name and action.


We saw how to improve the application in various ways. First, we created a support for multiline input 
(using a custom split function for a scanner) that supported quoted strings, with new lines. You can see this work in scanner package.


Then, we created some tools to add colored output to our functions and used them in one of the commands
previously defined. See color package!! 

We also used Levenshtein distance to suggest similar commands when the user specifies a non-existing one.

Finally, we separated commands from the main application and created a way of registering 
new commands from the outside. We used an interface because this allows better extension and 
customization, together with a basic implementation of the interface.

<p>
    <img src="https://raw.githubusercontent.com/zucchinidev/pty-go/master/img/visualization.png" />
</p>