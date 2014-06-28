CREATE KEYSPACE wordfinder WITH REPLICATION = {
  'class' : 'SimpleStrategy',
  'replication_factor' : 1
};

use wordfinder;

CREATE TABLE word (
  display  VARCHAR,
  content  VARCHAR,
  primary key (display)
);

CREATE TABLE lookup (
  wordformDisplay VARCHAR,
  headword VARCHAR,
  primary key (wordformDisplay)
);
