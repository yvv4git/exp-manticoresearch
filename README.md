# exp-manticoresearch


## Multi nodes
On first node:
````
CREATE TABLE testrt ( title text, content text, gid integer);
CREATE CLUSTER posts;
ALTER CLUSTER posts ADD testrt;
SELECT * FROM testrt;
````

On second node:
````
JOIN CLUSTER posts AT 'manticore1:9312';
INSERT INTO posts:testrt(title,content,gid)  VALUES('hello','world',1);
````