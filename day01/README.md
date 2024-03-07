<h3 id="ex00">Exercise 00: Reading</h3>

First things first, he had to learn how to read the database. The owner already had a CLI, so he decided that reading the file should be straightforward, so both these should work (files can be distinguished by an extension, for simplicity):

`~$ ./readDB -f original_database.xml`
`~$ ./readDB -f stolen_database.json`

Not only that, he also decided that reading both files shouldn't be that difficult to do through the same interface, which he called `DBReader`. That means, reading different formats means that we have different *implementations* of the same interface `DBReader`, which should spit out the same object types as a result, whether it's reading from the original database or the stolen one. Right, his idea is to choose the appropriate implementation based on a file extension.

So, you'll need to help him with that. Think of which kinds of objects are there in these databases and how they can be represented in code. Then, write an interface `DBReader` and two implementations of it - one for reading JSON and one for reading XML. Both of them should return the object of the same type as a result.

To check that his idea works, make the code print JSON version of the database when it's reading from XML and vice versa. Both XML and JSON fields should be indented with 4 spaces ("pretty-printing").


<h3 id="ex01">Exercise 01: Assessing Damage</h3>

Okay, so now the owner decided to compare the databases. You've seen that the stolen database has modified versions of the same recipes, meaning there are several possible cases:

1) New cake is added or old one removed
2) Cooking time is different for the same cake
3) New ingredient is added or removed for the same cake. *Important:* the order of ingredients doesn't matter. Only the names are.
4) The count of units for the same ingredient has changed.
5) The unit itself for measuring the ingredient has changed.
6) Ingredient unit is missing or added

Quickly looking through the database, the owner also noticed that nobody bothered renaming the cakes or the ingredients (possibly, only added some new ones), so you may assume names are the same across both databases.

Your application should be runnable like this:

`~$ ./compareDB --old original_database.xml --new stolen_database.json`

It should work with both formats (JSON and XML) for original AND new database, reusing the code from Exercise 00.

The output should look like this (the same cases explained above):

```
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
```


<h3 id="ex02">Exercise 02: Afterparty</h3>

After this whole mess they both agreed to use your code, but asked you to do one last task for them. They were quite impressed by how you've managed to do the comparison between the databases, so they've also asked you to do the same thing with their server filesystem backups, so neither bakery would run into any technical issues in the future.

So, your program should take two filesystem dumps.

`~$ ./compareFS --old snapshot1.txt --new snapshot2.txt`

They are both plain text files, unsorted, and each of them includes a filepath on every like, like this:

```
/etc/stove/config.xml
/Users/baker/recipes/database.xml
/Users/baker/recipes/database_version3.yaml
/var/log/orders.log
/Users/baker/pokemon.avi
```

Your tool should output the very similar thing to a previous code (without CHANGED case though):

```
ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
```

There is one issue though - the files can be really big, so you can assume both of them won't fit into RAM on the same time. There are two possible ways to overcome this - either to compress the file in memory somehow, or just read one of them and then avoid reading the other. **Side note:** this is actually a very popular interview question.


