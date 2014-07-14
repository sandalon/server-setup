use wordfinder;
copy word (headword, content) from 'word.tsv' with DELIMITER = '\t' and HEADER = FALSE;
copy lookup (wordformdisplay, headword, content) from 'lookup.tsv' with DELIMITER = '\t' and HEADER = FALSE;
