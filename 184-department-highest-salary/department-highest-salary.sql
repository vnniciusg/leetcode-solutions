-- Write your PostgreSQL query statement below
WITH CTE AS (
    SELECT 
        d.name as Department,
        e.name as Employee,
        e.salary as Salary,
        RANK() OVER (PARTITION BY d.name ORDER BY e.salary DESC) AS rank
    FROM Employee e
    JOIN Department d ON e.departmentId = d.id
)

SELECT 
    Department,
    Employee,
    Salary
FROM CTE
WHERE rank = 1;
