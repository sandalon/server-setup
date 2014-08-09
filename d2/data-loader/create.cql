CREATE KEYSPACE wordfinder WITH REPLICATION = {
  'class' : 'SimpleStrategy',
  'replication_factor' : 1
};

use wordfinder;

CREATE TABLE word (
  headword  VARCHAR,
  content  VARCHAR,
  primary key (headword)
);

CREATE TABLE lookup (
  wordformDisplay VARCHAR,
  headword VARCHAR,
  content VARCHAR,
  title VARCHAR,
  meta VARCHAR,
  primary key (wordformDisplay)
);
