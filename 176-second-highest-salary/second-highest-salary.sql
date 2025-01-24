-- Write your PostgreSQL query statement below
SELECT 
    MAX(salary) AS "SecondHighestSalary"
FROM Employee
WHERE salary <> (
    SELECT 
        MAX(salary) AS SALARY
    FROM Employee
);s
