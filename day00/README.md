<h3 id="ex00">Exercise 00: Anscombe</h3>


So, let's say we have a bunch of integer numbers, strictly between -100000 and 100000. It may 
probably be a large set, so let's assume our application will read it from a standard input, 
separated by newlines. Right now let's think of four major statistical metrics that we can derive
from this data, so by default we can print all of them as a result, for example, like this:

```
Mean: 8.2
Median: 9.0
Mode: 3
SD: 4.35
```

It will also make sense for user to be able to choose specifically, which of these four parameters
to print, so need to implement this as well. By default it's all of them, but there should be 
a way to specify whether print just one, two or three specific metrics out of four when running
the program (without recompilation). There is a built-in module in standard library that allows you 
to parse additional parameters.

