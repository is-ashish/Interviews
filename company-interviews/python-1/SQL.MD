# SQL Interview Practice Questions

## Table Structure Reference

### Customers Table
| customer_id | first_name | last_name | age | country |
|---|---|---|---|---|
| 1 | John | Doe | 31 | USA |
| 2 | Robert | Luna | 22 | USA |
| 3 | David | Robinson | 22 | UK |
| 4 | John | Reinhardt | 25 | UK |
| 5 | Betty | Doe | 28 | UAE |

### Orders Table
| order_id | item | amount | customer_id |
|---|---|---|---|
| 1 | Keyboard | 400 | 4 |
| 2 | Mouse | 300 | 4 |
| 3 | Monitor | 12000 | 3 |
| 4 | Keyboard | 400 | 1 |
| 5 | Mousepad | 250 | 2 |

### Shippings Table
| shipping_id | status | customer |
|---|---|---|
| 1 | Pending | 2 |
| 2 | Pending | 4 |
| 3 | Delivered | 3 |
| 4 | Pending | 5 |
| 5 | Delivered | 1 |

---

## INTERMEDIATE LEVEL

### 11. **INNER JOIN**
**Question:** Get customer names along with their orders.
```sql
SELECT c.first_name, c.last_name, o.item, o.amount
FROM Customers c
INNER JOIN Orders o ON c.customer_id = o.customer_id;
```

### 12. **LEFT JOIN**
**Question:** Get all customers and their orders (even if they haven't placed any orders).
```sql
SELECT c.first_name, c.last_name, o.item, o.amount
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id;
```

### 13. **GROUP BY with COUNT**
**Question:** How many orders did each customer place?
```sql
SELECT c.first_name, c.last_name, COUNT(o.order_id) AS total_orders
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name
ORDER BY total_orders DESC;
```

### 14. **GROUP BY with SUM**
**Question:** What is the total amount spent by each customer?
```sql
SELECT c.first_name, c.last_name, SUM(o.amount) AS total_spent
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name
ORDER BY total_spent DESC;
```

### 15. **GROUP BY with AVG**
**Question:** What is the average order value for each customer?
```sql
SELECT c.first_name, c.last_name, AVG(o.amount) AS avg_order_value
FROM Customers c
INNER JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name;
```

### 16. **HAVING Clause**
**Question:** Find customers who have spent more than 500 in total.
```sql
SELECT c.first_name, c.last_name, SUM(o.amount) AS total_spent
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name
HAVING SUM(o.amount) > 500
ORDER BY total_spent DESC;
```

### 17. **Multiple JOINs**
**Question:** Get customer names, their orders, and shipping status.
```sql
SELECT c.first_name, c.last_name, o.item, o.amount, s.status
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
LEFT JOIN Shippings s ON c.customer_id = s.customer
ORDER BY c.customer_id;
```

### 18. **Subquery in WHERE Clause**
**Question:** Find customers who have placed orders (using subquery).
```sql
SELECT * FROM Customers
WHERE customer_id IN (SELECT DISTINCT customer_id FROM Orders);
```

### 19. **Subquery with COUNT**
**Question:** Find customers whose total spending is above the average spending.
```sql
SELECT c.first_name, c.last_name, SUM(o.amount) AS total_spent
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name
HAVING SUM(o.amount) > (SELECT AVG(total) FROM (SELECT SUM(amount) AS total FROM Orders GROUP BY customer_id) AS avg_spent);
```

### 20. **CASE Statement**
**Question:** Categorize customers based on their spending (Low: <500, Medium: 500-1000, High: >1000).
```sql
SELECT c.first_name, c.last_name,
       SUM(o.amount) AS total_spent,
       CASE 
           WHEN SUM(o.amount) IS NULL THEN 'No Orders'
           WHEN SUM(o.amount) < 500 THEN 'Low Spender'
           WHEN SUM(o.amount) BETWEEN 500 AND 1000 THEN 'Medium Spender'
           ELSE 'High Spender'
       END AS spending_category
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name;
```

---

## ADVANCED LEVEL

### 21. **Ranking with ROW_NUMBER()**
**Question:** Rank customers by their total spending.
```sql
SELECT c.first_name, c.last_name,
       SUM(o.amount) AS total_spent,
       ROW_NUMBER() OVER (ORDER BY SUM(o.amount) DESC) AS rank
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name;
```

### 22. **RANK() vs DENSE_RANK()**
**Question:** Rank customers by age (handle ties with RANK and DENSE_RANK).
```sql
SELECT first_name, last_name, age,
       RANK() OVER (ORDER BY age DESC) AS rank_with_gaps,
       DENSE_RANK() OVER (ORDER BY age DESC) AS rank_without_gaps
FROM Customers;
```

### 23. **Running Total (SUM Window Function)**
**Question:** Show a running total of order amounts for each customer.
```sql
SELECT c.first_name, c.last_name, o.item, o.amount,
       SUM(o.amount) OVER (PARTITION BY c.customer_id ORDER BY o.order_id) AS running_total
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
ORDER BY c.customer_id, o.order_id;
```

### 24. **Finding Customers with No Orders**
**Question:** Get customers who haven't placed any orders.
```sql
SELECT c.* FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
WHERE o.order_id IS NULL;
```

### 25. **Finding Duplicate Customers (Same First and Last Name)**
**Question:** Find if there are customers with the same first name and last name.
```sql
SELECT first_name, last_name, COUNT(*) AS count
FROM Customers
GROUP BY first_name, last_name
HAVING COUNT(*) > 1;
```

### 26. **UNION - Combine Results**
**Question:** Get a list of all unique first names from both customers and their orders.
```sql
SELECT first_name FROM Customers
UNION
SELECT item FROM Orders;
```

### 27. **NOT IN with Subquery**
**Question:** Find customers who have not placed any orders.
```sql
SELECT * FROM Customers
WHERE customer_id NOT IN (SELECT DISTINCT customer_id FROM Orders);
```

### 28. **EXISTS Subquery**
**Question:** Find customers who have at least one pending shipment.
```sql
SELECT DISTINCT c.*
FROM Customers c
WHERE EXISTS (SELECT 1 FROM Shippings s WHERE s.customer = c.customer_id AND s.status = 'Pending');
```

### 29. **NOT EXISTS Subquery**
**Question:** Find customers who don't have any pending shipments.
```sql
SELECT c.*
FROM Customers c
WHERE NOT EXISTS (SELECT 1 FROM Shippings s WHERE s.customer = c.customer_id AND s.status = 'Pending');
```

### 30. **Self JOIN (Finding Customers from Same Country)**
**Question:** Find pairs of customers from the same country.
```sql
SELECT c1.first_name AS customer1, c2.first_name AS customer2, c1.country
FROM Customers c1
INNER JOIN Customers c2 ON c1.country = c2.country AND c1.customer_id < c2.customer_id;
```

### 31. **CTE (Common Table Expression)**
**Question:** Find the top 3 customers by spending using CTE.
```sql
WITH customer_spending AS (
    SELECT c.customer_id, c.first_name, c.last_name, SUM(o.amount) AS total_spent
    FROM Customers c
    LEFT JOIN Orders o ON c.customer_id = o.customer_id
    GROUP BY c.customer_id, c.first_name, c.last_name
)
SELECT * FROM customer_spending
ORDER BY total_spent DESC
LIMIT 3;
```

### 32. **Multiple CTEs**
**Question:** Find customers who have spent more than average AND have pending shipments.
```sql
WITH avg_spending AS (
    SELECT AVG(amount) AS avg_amount FROM Orders
),
customer_totals AS (
    SELECT c.customer_id, c.first_name, c.last_name, SUM(o.amount) AS total_spent
    FROM Customers c
    LEFT JOIN Orders o ON c.customer_id = o.customer_id
    GROUP BY c.customer_id, c.first_name, c.last_name
)
SELECT DISTINCT ct.first_name, ct.last_name, ct.total_spent
FROM customer_totals ct
CROSS JOIN avg_spending av
WHERE ct.total_spent > av.avg_amount
AND EXISTS (SELECT 1 FROM Shippings s WHERE s.customer = ct.customer_id AND s.status = 'Pending');
```

### 33. **Recursive CTE (Hierarchical Data)**
**Question:** Generate series of numbers (common pattern in real queries).
```sql
WITH RECURSIVE numbers AS (
    SELECT 1 AS num
    UNION ALL
    SELECT num + 1 FROM numbers WHERE num < 5
)
SELECT * FROM numbers;
```

### 34. **String Functions - CONCAT**
**Question:** Get customer full names with their country.
```sql
SELECT CONCAT(first_name, ' ', last_name, ' (', country, ')') AS full_customer_info
FROM Customers;
```

### 35. **String Functions - SUBSTRING/SUBSTR**
**Question:** Extract the first 3 characters of country name.
```sql
SELECT first_name, last_name, SUBSTRING(country, 1, 3) AS country_code
FROM Customers;
```

### 36. **String Functions - UPPER/LOWER**
**Question:** Get customer names in uppercase with items in lowercase.
```sql
SELECT UPPER(CONCAT(c.first_name, ' ', c.last_name)) AS customer_name, LOWER(o.item) AS item
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id;
```

### 37. **String Functions - LENGTH**
**Question:** Find customers whose last name is longer than 5 characters.
```sql
SELECT * FROM Customers WHERE LENGTH(last_name) > 5;
```

### 38. **COALESCE Function**
**Question:** Show order details, but if no amount, show 0.
```sql
SELECT order_id, item, COALESCE(amount, 0) AS amount, customer_id
FROM Orders;
```

### 39. **NULLIF Function**
**Question:** Return NULL if amount equals 400 (Keyboard price), otherwise return the amount.
```sql
SELECT order_id, item, NULLIF(amount, 400) AS amount
FROM Orders;
```

### 40. **CASE with Aggregation**
**Question:** Count orders by type (delivered vs pending shipments).
```sql
SELECT 
    SUM(CASE WHEN s.status = 'Delivered' THEN 1 ELSE 0 END) AS delivered_count,
    SUM(CASE WHEN s.status = 'Pending' THEN 1 ELSE 0 END) AS pending_count
FROM Shippings s;
```

---

## BONUS: Real-World Scenario Questions

### 41. **Data Integrity Check**
**Question:** Identify orders from customers who don't exist in the Customers table.
```sql
SELECT o.* FROM Orders o
LEFT JOIN Customers c ON o.customer_id = c.customer_id
WHERE c.customer_id IS NULL;
```

### 42. **Cross-Table Consistency**
**Question:** Find shipping records for customers who don't have orders.
```sql
SELECT DISTINCT s.customer
FROM Shippings s
WHERE s.customer NOT IN (SELECT DISTINCT customer_id FROM Orders);
```

### 43. **Summary Report**
**Question:** Create a summary report showing customer info, order count, total spent, and latest shipping status.
```sql
SELECT 
    c.customer_id,
    CONCAT(c.first_name, ' ', c.last_name) AS customer_name,
    c.country,
    COUNT(DISTINCT o.order_id) AS total_orders,
    COALESCE(SUM(o.amount), 0) AS total_spent,
    COALESCE(s.status, 'No Shipment') AS latest_status
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
LEFT JOIN Shippings s ON c.customer_id = s.customer
GROUP BY c.customer_id, c.first_name, c.last_name, c.country, s.status
ORDER BY total_spent DESC;
```

### 44. **Percentile Calculation**
**Question:** Find customers in the top 25% by spending.
```sql
SELECT c.first_name, c.last_name, SUM(o.amount) AS total_spent,
       PERCENT_RANK() OVER (ORDER BY SUM(o.amount)) AS percentile_rank
FROM Customers c
LEFT JOIN Orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id, c.first_name, c.last_name
ORDER BY percentile_rank DESC;
```

### 45. **Pivot-like Result (Multiple Columns from Single Column)**
**Question:** Show shipment status count by country.
```sql
SELECT 
    c.country,
    SUM(CASE WHEN s.status = 'Delivered' THEN 1 ELSE 0 END) AS delivered,
    SUM(CASE WHEN s.status = 'Pending' THEN 1 ELSE 0 END) AS pending
FROM Customers c
LEFT JOIN Shippings s ON c.customer_id = s.customer
GROUP BY c.country;
```

---

## Interview Tips

1. **Read the question carefully** - Understand what data you need to return
2. **Start simple** - Use basic SELECT before jumping to complex JOINs
3. **Test with the sample data** - Verify your results make sense
4. **Explain your approach** - Walk the interviewer through your logic
5. **Consider edge cases** - What about NULL values? Empty results?
6. **Optimize when asked** - Use indexes, avoid SELECT *, optimize JOINs
7. **Practice syntax** - Different databases (MySQL, PostgreSQL, SQL Server) have slight variations
8. **Know your window functions** - They're increasingly common in modern interviews

---

## Common SQL Keywords Reference

| Keyword | Purpose |
|---------|---------|
| SELECT | Choose columns to return |
| FROM | Specify table(s) |
| WHERE | Filter rows (conditions) |
| JOIN | Combine tables |
| GROUP BY | Aggregate data |
| HAVING | Filter groups |
| ORDER BY | Sort results |
| LIMIT | Limit number of rows |
| COUNT, SUM, AVG, MIN, MAX | Aggregate functions |
| DISTINCT | Remove duplicates |
| CASE | Conditional logic |
| UNION | Combine result sets |
| EXISTS, IN, NOT IN | Subqueries |
| OVER(), PARTITION BY | Window functions |

Good luck with your interview! 🚀