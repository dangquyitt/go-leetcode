# 0196. Delete Duplicate Emails

```sql
DELETE p1
FROM Person p1
INNER JOIN Person p2 ON p2.Email = p1.Email AND p1.Id > p2.Id;
```

**Input**:
Person table:
| id |email |
|--- |--- |
|1 | john@example.com|
|2 | bob@example.com|
|3 | john@example.com|

**Output**:
| id |email |
|--- |--- |
|1 | john@example.com|
|2 | bob@example.com|

**Explanation**: john@example.com is repeated two times. We keep the row with the smallest Id = 1.
