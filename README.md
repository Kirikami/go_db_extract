Test exercise for extracting records from database


Introduction:

We are running a multi country setup for any client. 


Each of the countries runs  a dedicated MySQL instance containing two tables. 

We have been asked to dump the content of each db instance into its respective table files. 

 

 To make things a bit more challenging, the dump must be implement via "SELECT" query. 

  

  The DB schema for each instance (country) would more or less look like:

   

   Table users ( 

   user_id  int, auto increment, primary key

   name      varchar(255)

   )

    

    Table sales (

    order_id  int, auto increment, primary key

    user_id   int

    order_amount float

    )

     

     Your Task:

     Implement a GoLang application matching the following requirements:

     1. Dump both tables from each instance into the files "users.csv" and "sales.csv" files respectively. You can ignore any order order with the file. 

     2. Dumps must be implemented via "SELECT" query. 

     2. For performance reasons, both db instances have to be processed concurrently (would be threads in case of Java)

     3. At the end of the process the tool must prompt feedback on lines of entries for each csv file. 

     4. Both files should eventually be archived and stored in the relative directory "./archive"

      

      Good luck
