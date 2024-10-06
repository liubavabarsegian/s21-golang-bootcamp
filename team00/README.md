# Team 00 - Go Boot camp

<h3 id="ex00">Task 00: Transmitter</h3>

Create gRPC with each message consisting of just three fields - 'session_id' as a string, 'frequency' as a double and also a current timestamp in UTC.

We don't know much about distribution here, so let's implement it in such way that whenever new client connects [expected value](https://en.wikipedia.org/wiki/Expected_value) and [standard deviation](https://en.wikipedia.org/wiki/Standard_deviation)" are picked at random. For this experiment, let's pick mean from [-10, 10] interval and standard deviation from [0.3, 1.5].

On each new connection server should generate a random UUID (sent as session_id) and new random values for mean and STD. All generated values should be written to a server log (stdout or file). After that it should send a stream of entries with fields explained above, where for each message 'frequency' would be a value picked at random (sampled) from a normal distribution with these standard deviation and expected value.

It is required to describe the schema in a *.proto* file and generate the code itself from it. Also, you shouldn't modify generated code manually, just import it.

<h2 id="chapter-v" >Chapter V</h2>
<h3 id="ex01">Task 01: Anomaly Detection</h3>

"Now to the interesting part! While others are working on gRPC server, let's think of a client. I expect that gRPC client should be handled by the same guys writing the server to test it, so let's focus on a different thing. We need to detect anomalies in a frequency distribution!"

So, you know you're getting a stream of values. With each new incoming entry from a stream your code should be able to approximate mean and STD from the random distribution generated on a server. Of course it's not really possible to predict it looking only on 3-5 values, but after 50-100 it should be precise enough. Keep in mind that mean and STD are generated for each new connection, so you shouldn't restart the client during the process. Also, values shouldn't keep piling up in memory, so you may consider using sync.Pool for easy reuse.

While working on this task, you can temporarily forget about gRPC and test the code by just sending it a sequence of values to stdin.

Your client code should write into a log periodically, how many values are processed so far as well as predicted values of mean and STD.

After some time, when your client decides that the predicted distribution parameters are good enough (feel free to choose this moment by yourself), it should switch automatically into an Anomaly Detection stage. Here there is one more parameter which comes into play - an *STD anomaly coefficient*. So, your client should accept a command-line parameter (let it be '-k') with a float-typed coefficient.

An incoming frequency is considered an anomaly, if it differs from the expected value by more than *k \* STD* to any side (to the left or to the right, as the distribution is symmetric). You can read more about how it works by following links from Chapter 4.

For now you should just write found anomalies into a log.

<h3 id="ex02">Task 02: Report</h3>

"As general knows nothing about our *sciency gizmo*, let's store all anomalies that we encounter in a database and then he'll be able to look at it through some interface they have" - Louise seems to be a lot more concerned about the data rather than the general.

So, let's learn how to write data entries to PostgreSQL. Usually it is considered a bad practice to just write plain SQL queries in code when dealing with highly secure environments (you can read about SQL Injections by following links from Chapter 4). Let's use an ORM. In case of PostgreSQL there are two most obvious choices (these links are below as well), but you can choose any other. The main idea here is to not have any strings with SQL code in your sources.

You'll have to describe your entry (session_id, frequency and a timestamp) as a structure in Go and then use it together with ORM to map it into database columns.


