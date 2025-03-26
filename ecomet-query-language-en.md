ECOMET QUERY LANGUAGE SPECIFICATION 

=QUERY LANGUAGE=
Here is the **translation of the full Ecomet Query Language (EQL) documentation** to English:

---

**DOCUMENTATION IN PROGRESS**  
Version: 20250305-1913

## Overview of the Ecomet Query Language

This language is designed for operations such as data retrieval, modification, subscription management, and deletion, providing a structured and expressive syntax. The following sections document the available query types, their syntax, and examples.

---

## 1. GET Queries

**Purpose**  
Retrieve data from one or more databases (or objects) using filtering, grouping, sorting, pagination, and aggregation.

**General Syntax**  
```plaintext
get <field list> from <database list> where <conditions> [options]
```

- **Field list**: One or more fields to retrieve. Function calls (`$concat`, `$term`) and aliases (`AS`) are supported.
- **Database list**: Names of databases or object identifiers (e.g., `root`, `db1`).
- **Conditions**: Logical expression with operators `=`, `:=`, `:>`, `:<`, `:LIKE` and logical connectors `AND`, `OR`, `ANDNOT`.
- **Options** (optional): Additional parameters like `group by`, `order by`, `page`, `lock`, `format`.

**Examples**

Simple GET query:
```plaintext
get foo AS 'alias' from root where AND ( buz :> 123 )
```

Query with multiple databases and fields:
```plaintext
get foo1, foo2 from db1,db2 where OR ( buz :LIKE 'f', bar = -42 )
```

Negation and sorting:
```plaintext
get foo from 'TEST' where ANDNOT (buz := 45.5, bar :< 4) order by foo3 DESC
```

Pagination:
```plaintext
get foo from root where AND (buz := 1) page 5:10
```

Read lock:
```plaintext
get foo from * where AND (bar :< 4) page 5:10 lock read
```

Grouping and sorting:
```plaintext
get foo1 from d1,d2,'D3' where AND (bar :< 4) group by foo2, buz2 order by foo3, buz3 DESC
```

Custom formatting:
```plaintext
get foo1 from root where AND (bar :< 4) format $json
```

With a custom formatting function:
```plaintext
get foo1 from root where AND (bar :< 4) format $my_lib:my_fun
```

---

## 2. SUBSCRIBE and UNSUBSCRIBE Queries

**Purpose**  
Subscribe to data updates and manage subscriptions.

**Syntax**

Subscribe:
```plaintext
subscribe <query> [options]
```

Unsubscribe:
```plaintext
unsubscribe <subscription-id>
```

Example:
```plaintext
subscribe get foo from root where AND ( buz :> 123 )
```

Unsubscribe:
```plaintext
unsubscribe <subscription-id>
```

---

## 3. SET Queries

**Purpose**  
Update field values in existing objects.

**Syntax**
```plaintext
set <field assignments> in <database list> where <conditions> [options]
```

**Examples**

Simple update:
```plaintext
set foo=$hello in root where AND (bar :< 1) lock write
```

Update with a function:
```plaintext
set foo=$concat($buz,'_test') in * where OR (bar :> 1) lock read
```

---

## 4. INSERT Queries

**Purpose**  
Insert new objects into the database.

**Syntax**
```plaintext
insert <field assignments> [format <formatter>]
```

**Examples**

Simple insert:
```plaintext
insert foo1=123, foo2=$term('buz')
```

With formatting:
```plaintext
insert foo1=123, foo2=$term('buz') format $from_string
```

---

## 5. DELETE Queries

**Purpose**  
Delete objects from the database.

**Syntax**
```plaintext
delete from <database list> where <conditions> [options]
```

**Examples**

Delete with lock:
```plaintext
delete from root where AND (bar :< 1) lock write
```

Delete with complex condition:
```plaintext
delete from db1,db2 where ANDNOT (bar :LIKE '1', buz := foo)
```

---

## 6. Transaction Management

**Purpose**  
Begin, commit, or roll back transactions.

**Syntax**

Start transaction:
```plaintext
transaction_start
```

Commit:
```plaintext
transaction_commit
```

Rollback:
```plaintext
transaction_rollback
```

**Example**

Commit:
```plaintext
transaction_start
set foo=$hello in root where AND (bar :< 1) lock write
transaction_commit
```

Rollback:
```plaintext
transaction_start
... (some queries) ...
transaction_rollback
```

---

## Summary

Ecomet Query Language supports:

- **Data retrieval (`get`)** with filtering, sorting, grouping
- **Data modification (`set`)**
- **Data insertion (`insert`)**
- **Data deletion (`delete`)**
- **Subscriptions (`subscribe`/`unsubscribe`)**
- **Transaction control**

Each query can include options like locking, formatting, grouping, and pagination — making EQL a powerful tool for distributed data systems.

Use the examples to build your own queries! 🚀

---

## EQL: `$count` Function Usage

**$count** is an aggregate function in EQL used to count the number of elements (objects, records, etc.) matching a query or condition. It's commonly used with `group by`.

### 1. Basic Usage
Returns a single number representing the count of all matching elements.

### 2. With `group by`
Returns a count for each group.

### 3. With `AS`
Alias can be used to name the result (useful for programmatic use).

### 4. Multiple Aggregate Functions
Multiple aggregate functions can be used in one query.

### 5. Restrictions
- Only in `get` and `subscribe`
- Must be part of the field list in `get`
- Syntax: `$count()` — no arguments allowed

### Summary
- Powerful counting tool in EQL
- Useful for analytics and summaries
- Combine with `group by` for aggregate insights

---

## EQL: `group by` Clause Usage

Used to group rows with identical values in one or more columns into summary rows (e.g., "count users per country").

### 1. Basic Usage
Syntax: list one or more fields to group by. Each unique combo creates a group.

### 2. With Aggregate Functions
Group-specific aggregates like `$count`, `$sum`, `$avg`.

### 3. Multiple Fields
Multi-level grouping is supported.

### 4. With `where`
Filter before grouping.

### 5. With `order by`
Group and sort together.

### 6. Restrictions
Fields not in `group by` must be inside an aggregate function.  
Only works in `get` and `subscribe`.

---

## LIKE Operator in EQL

**Purpose**  
Pattern matching on string fields (specifically `binary()` in Erlang).

**Supported Wildcards**
- `^`: Match from the start of the string.
- `_`: Match any single character.
- `%`: Match any sequence of characters (zero or more).

**Restrictions**
- Only the above are supported.
- No support for full regex syntax (`*`, `+`, `$`, etc.)

**Syntax**
```erlang
{ Field, 'LIKE', Pattern }
```

**Examples**
- Starts with:
  ```erlang
  {<<"fp_path">>, 'LIKE', <<^/root/FP/PROJECT/>>}
  ```
- Single character:
  ```erlang
  {<<"fp_path">>, 'LIKE', <<"/root/_/">>}
  ```
- Any sequence:
  ```erlang
  {<<"fp_path">>, 'LIKE', <<"/root/%/">>}
  ```

**Usage in code**  
Used in `get` and `subscribe` functions to filter objects based on path matching.

---

## Practical Examples from the BPM Journals Project

### Current Shift
```plaintext
GET status, .name, .folder
FROM *
WHERE AND (
  .pattern = $oid('/root/.patterns/shift'),
  .path :like '/root/JOURNALS/OPERATIONAL/technical/',
  OR ( status = 'active' )
)
PAGE 1:50
format $to_json
```

### User Groups (Multiple Roles)
```plaintext
GET .name, .oid, .pattern, .ts
FROM *
WHERE AND (
  .folder = $oid('/root/.users'),
  OR (
    usergroups = $oid('/root/.usergroups/oper_chief_dispatcher'),
    usergroups = $oid('/root/.usergroups/oper_dispatcher'),
    usergroups = $oid('/root/.usergroups/oper_engineer')
  )
)
ORDER BY .ts DESC
PAGE 1:200
format $to_json
```

### Range
```plaintext
GET .name, .folder
FROM *
WHERE AND (
  .pattern = $oid('/root/.patterns/shift'),
  AND (
    .name :>= '159',
    .name :=< '180'
  )
)
ORDER BY .name DESC
PAGE 1:100
format $to_json
```

### Exact Match
```plaintext
GET .oid, .path, .folder, .pattern, .name
FROM *
WHERE AND (
  .pattern = $oid('/root/.patterns/shift'),
  .name = '3122'
)
format $to_json
```

### LIKE Partial Match
```plaintext
GET .folder, message, .oid, created_at
FROM *
WHERE AND (
  .pattern = $oid('/root/.patterns/record'),
  message :like 'Suburban-East'
)
ORDER BY created_at DESC
format $to_json
```

---

## Final Summary

```plaintext
GET <fields> FROM * WHERE AND(<conditions>) 
[ORDER BY <field> ASC|DESC] 
[PAGE <page_number>:<page_size>] 
[format $to_json]
```

Common usages:
- Filter by folder or status using `.path :like` and `status = 'active'`
- Use `OR()` for multiple conditions
- Filter by user groups with `usergroups = $oid(...)` or `:in [...]`
- Use range (`:>=`, `:=<`) and exact match (`=`) for `.name`
- Use `:like` for partial matches

--- 

Let me know if you’d like this exported to a PDF, formatted for publishing, or explained in diagram form!

=Usage in bash=
```bash
wscat -c wss://localhost:9443/websocket --no-check
# Connected (press Ctrl+C to quit)

# Login (replace with your credentials)
> {"id": 0, "action": "login", "params": {"login": "admin", "pass": "admin"}} 
# Wait for login response (should be {"id":0,"type":"ok","result":null})

# Create alarm (replace with your data. Make sure folder exists and has correct permissions)
> {"id": 1, "action": "create", "params": {
    "fields": {
      ".folder": "/root/.databases/project/LOG",  # Correct folder OID!
      ".pattern": "/root/.patterns/alarm",
      "type": "alarm",
      "point": "WscatPoint",
      "text": "Alarm from wscat",
      "dt_on": 1700000000000 # Millisecond timestamp
    }
  }
}
# Wait for create response (should be {"id":1,"type":"ok","result":"<OID of created alarm>"})

# Get server time
> {"id": 2, "action": "application", "params": {
    "module": "fp_json",
    "function": "get_server_time",
    "function_params": {}
  }
}
# Wait for response


# Query for alarms (example – adjust as needed)
> {"id": 3, "action": "query", "params": {
      "statement": "GET * FROM /root/.databases/project/LOG WHERE .pattern=$oid('/root/.patterns/alarm') and active=true format $to_json"
  }}


# Logout (optional)
> {"id": 4, "action": "logout", "params": {}}

```

