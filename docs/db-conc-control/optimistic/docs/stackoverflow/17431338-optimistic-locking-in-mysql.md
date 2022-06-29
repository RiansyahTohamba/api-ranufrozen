17431338-optimistic-locking-in-mysql.md

# Optimistic locking in MySQL

I can't find any details on optimistic locking in MySQL. 

I read that starting a transaction keep updates on two entities synced, however - it doesn't stop two users updating the data at the same time causing a conflict.

Apparently optimistic locking will solve this issue? 

How is this applied in MySQL. Is there SQL syntax / keyword for this? 
Or does MySQL have default behavior?

# singkatnya:
1. apakah optimis locking sudah ada didalam internal MySQL secara defaul 
atau harus programmer tuliskan syntax khusus?
- jawaban:
Ya, ada syntax khususnya juga, seperti berikut:
SQL + Host Code (Go/PHP)


# OPTIMIS LOCKING (WITHOUT TRANSACTION)
The OPTIMISTIC LOCKING way is:
- SELECT iD, val1, val2 FROM theTable WHERE iD = @theId;

 - {code that calculates new values}

 - UPDATE theTable
       SET val1 = @newVal1,
           val2 = @newVal2
       WHERE iD = @theId
           AND val1 = @oldVal1
           AND val2 = @oldVal2;
 - {if AffectedRows == 1 }
 -     {go on with your other code}
 - {else}
- sebenarnya tidak akan perubahan update juga? ngapain harus rollback ya?
 -     {decide what to do since it has gone bad... in your code}
 - {endif}

Note that the key point is 
1. in the structure of the UPDATE instruction and 
2. the subsequent number of affected rows check. 

It is these two things together that let your code realize that someone has already modified the data in between when you have executed the SELECT and UPDATE. 

Notice that all has been done without transactions! 

This has been possible (absence of transactions) only because this is a very simple example but this tells also that the key point for Optimistic locking is not in transactions themselves.


# OPTIMIS LOCKING WITH TRANSACTION
pada contoh control untuk transaction, case yang diberikan adalah pengubahan 2 table.

jika table_1 terdapat affected rows maka table_2 juga harus ada affected rows (?)
jangan hanya salah satu saja yang berhasil update, yang lain gagal.

What about TRANSACTIONS then?
 - SELECT iD, val1, val2 FROM theTable WHERE iD = @theId;

 - {code that calculates new values}

 - BEGIN TRANSACTION;
 
 - UPDATE table_1
       SET col1 = @newCol1,
           col2 = @newCol2
       WHERE iD = @theId;

 - UPDATE table_2
       SET val1 = @newVal1,
           val2 = @newVal2
       WHERE iD = @theId
           AND val1 = @oldVal1
           AND val2 = @oldVal2;
 
 - {if AffectedRows == 1 }
 -     COMMIT TRANSACTION;
 -     {go on with your other code}
 - {else}
 -     ROLLBACK TRANSACTION;
 -     {decide what to do since it has gone bad... in your code}
 - {endif}



This last example shows that if you check for collisions at some point and discover a collision has happened when you have already modified other tables/rows.
then with transactions you are able to rollback ALL the changes that you've done since the beginning. 

Pada contoh terakhir menunjukkan bahwa jika kamu mau cek collisions at some at point and discover a collison has happened ketika kamu telah memodifikasi table/row lain.

collison disini bermakna tubrukan antara 2 Transaksi. 
Misalkan terdapat 2 user yang melakukan transaksi (T1 DAN T2) dengan membeli produk dengan id=1.

T1:
    stok-awal= 2
    membeli product id=1 sebanyak 2
    UPDATE table order

    stock berkurang menjadi 0
    UPDATE table product

T2:
    stok-awal=2
    -- ditengah jalan, T1 sudah sukses
    UPDATE table order

    membeli product id=1 sebanyak 2
    stock berkurang menjadi -2 (aneh?)
    ROLLBACK(`UPDATE table order`)

pada case diatas T2 tubrukan dengan T1. Maka kita dapat menggunakan fitur rollback pada transaction
---
# nah sebelum lanjut, harus tahu dulu apa itu collision
collision

----

Obviously it is up to you (that knows what your application is doing) to decide how large the amount of operations to rollback is for each possible collision and based on this decide where to put the transactions boundaries and where to check for collisions with the special `UPDATE + AffectedRows check`.
In this case with transactions we have `separated` the moment 
1. when we perform the UPDATE 
2. from the moment when it is committed. 

terserah kita, seberapa besar total operasi yang akan di rollback untuk setiap kemungkinan collision dan berdasarkan keputusan total operasi ini, dimana kita akan tempatkan boundaries dan dimana kita akan mengecek collision pada UPDATE + affected rows check.

Pada contoh kasus ini, beliau telah memisahkan momen dimana
1. update berlangsung
2. moment `commit or rollback`

So what happens when an "other process" performs an update in this time frame? To know what happens exactly requires delving into the details of isolation levels (and how they are managed on each engine). 

As an example in the case of Microsoft SQL Server with READ_COMMITTED the updated rows are locked until the COMMIT, so "other process" can't do nothing (is kept waiting) on that rows, neither a SELECT (in fact it can only READ_COMMITTED). 

So since the "other process" activity is deferred it's UPDATE will fail.

# versioning
The VERSIONING OPTIMISTIC LOCKING option:
 - SELECT iD, val1, val2, version
       FROM theTable
       WHERE iD = @theId;
 - {code that calculates new values}
 - UPDATE theTable
       SET val1 = @newVal1,
           val2 = @newVal2,
           version = version + 1
       WHERE iD = @theId
           AND version = @oldversion;
 - {if AffectedRows == 1 }
 -     {go on with your other code}
 - {else}
 -     {decide what to do since it has gone bad... in your code}
 - {endif}
 
Here it is shown that instead of checking if the value is still the same for all the fields we can use a dedicated field (that is modified each time we do an UPDATE) to see if anyone was quicker than us and changed the row between our SELECT and UPDATE. Here the absence of transactions is due to the simplicity as in the first example and is not related with the version column use. 

Again this column use is up to the implementation in the application code and not a database engine feature.

More than this there are other points which I think would make this answer too long (is already much too long) so I only mention them by now with some references:

transaction isolation level (here for MySQL) about transaction effect on SELECTs.

for the INSERT on tables with primary keys not autogenerated (or unique constraints) it will automatically fail with no need of particular checking if two processes try to insert the same values where it must be unique.

if you have no id column (primary key or unique constraints) also a single SELECT + UPDATE require transactions because you could have the surprise that after modifications made by others there are more rows than expected matching the criteria of the UPDATE's WHERE clause.

# how to check
How to check in practice and get confident
Since the isolation level value and implementation may be different the best advice (as usual in this site) is to perform a test on the used platform / environment.

It may seem difficult but in reality it can be done quite easily from any DB development environment using two separate windows and starting on each one a transaction then executing the commands one by one.

At some point you will see that the the command execution continues indefinitely. Then when on the other window it is called COMMIT or ROLLBACK it completes the execution.

Here are some very basic commands ready to be tested as just described.

Use these for creating the table and one useful row:

CREATE TABLE theTable(
    iD int NOT NULL,
    val1 int NOT NULL,
    val2 int NOT NULL
);
INSERT INTO theTable (iD, val1, val2) VALUES (1, 2 ,3);
Then the following on two different windows and step by step:

BEGIN TRAN

SELECT val1, val2 FROM theTable WHERE iD = 1;

UPDATE theTable
  SET val1=11
  WHERE iD = 1 AND val1 = 2 AND val2 = 3;

COMMIT TRAN
Then change the order of commands and order of execution in any order you may think.


# pertanyaan selanjutnya, apakah ada software untuk menguji transaction 





