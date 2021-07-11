# Language Definitions

## Basic Keywords

### GRAB - SELECTS DATA

NEEDS:

- DATA

### REMOVE - REMOVES DATA

Removes json documents

NEEDS: 

- IDENTIFIER (what document(s) to remove)

### CHANGE 

NEEDS:

- NODE, ATTRIBUTE, NEW VALUE

### MAKE

NEEDS:

 - NODE, VALUES

## Query Keywords

## GET IDENTIFIER

### WHERE 

NEEDS:

- A key and a value 

CAN HAVE:

- AND (WHERE name = "value" AND age = 144)
- OR  (WHERE name = "value" OR age = 144)
- NOT (WHERE NOT name = "value" AND NOT age = 144)

## LOAD DATA

### LOAD

NEEDS: LOADABLE ARRAY OF LINKS

# Examples

//OBJECTID will be automatically assigned to a json document (like primary key in SQL)
GRAB USER WHERE OBJECTID = 323

GRAB USER WHERE username = "kanyewest" OR firstName = "kanye"

// LOAD will fill json with data from other document
GRAB COMPANY WHERE name = "companyname" LOAD EMPLOYEES
