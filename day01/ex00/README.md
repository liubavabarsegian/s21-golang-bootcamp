<h3 id="ex00">Exercise 00: Reading</h3>

First things first, he had to learn how to read the database. The owner already had a CLI, so he decided that reading the file should be straightforward, so both these should work (files can be distinguished by an extension, for simplicity):

`~$ ./readDB -f original_database.xml`
`~$ ./readDB -f stolen_database.json`

Not only that, he also decided that reading both files shouldn't be that difficult to do through the same interface, which he called `DBReader`. That means, reading different formats means that we have different *implementations* of the same interface `DBReader`, which should spit out the same object types as a result, whether it's reading from the original database or the stolen one. Right, his idea is to choose the appropriate implementation based on a file extension.

So, you'll need to help him with that. Think of which kinds of objects are there in these databases and how they can be represented in code. Then, write an interface `DBReader` and two implementations of it - one for reading JSON and one for reading XML. Both of them should return the object of the same type as a result.

To check that his idea works, make the code print JSON version of the database when it's reading from XML and vice versa. Both XML and JSON fields should be indented with 4 spaces ("pretty-printing").

