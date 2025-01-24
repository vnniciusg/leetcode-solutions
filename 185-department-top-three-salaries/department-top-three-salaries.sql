-- Write your PostgreSQL query statement below
WITH CTE AS (
    SELECT
        e.name AS Employee,
        e.salary AS Salary,
        d.name AS Department,
        DENSE_RANK() OVER (PARTITION BY d.name ORDER BY e.salary DESC) AS salary_rank
    FROM
        Employee e
    JOIN
        Department d ON e.departmentId = d.id
)

SELECT
    Department,
    Employee,
    Salary
FROM
    CTE
WHERE
    salary_rank <= 3;
