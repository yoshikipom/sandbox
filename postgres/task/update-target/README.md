# update-target

before
```
$ psql -h localhost -p 5432 -U root person 
psql (14.6 (Homebrew), server 12.13 (Debian 12.13-1.pgdg110+1))
Type "help" for help.

person=# select * from person;
 id | name | deleted 
----+------+---------
  1 | a    | f
  2 | b    | f
  3 | c    | f
  4 | d    | f
  5 | e    | f
(5 rows)
```

run
```
$ export PGPASSWORD=root && sh run.sh
1
UPDATE 1
3
UPDATE 1
5
UPDATE 1
```

after
```
$ psql -h localhost -p 5432 -U root person 
psql (14.6 (Homebrew), server 12.13 (Debian 12.13-1.pgdg110+1))
Type "help" for help.

person=# select * from person;
 id | name | deleted 
----+------+---------
  2 | b    | f
  4 | d    | f
  1 | a    | t
  3 | c    | t
  5 | e    | t
(5 rows)
```
