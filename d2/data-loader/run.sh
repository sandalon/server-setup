#!/bin/bash

cqlsh db 9160 -f /create.cql
cqlsh db 9160 -f /word.cql
