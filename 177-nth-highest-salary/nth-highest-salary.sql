CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) 
RETURNS TABLE (Salary INT) AS 
$$
BEGIN
  RETURN QUERY (
    SELECT CTE.Salary
    FROM (
        SELECT   
            E.Salary,
            DENSE_RANK() OVER (ORDER BY E.Salary DESC) AS Rank
        FROM
            Employee E
    ) AS CTE
    WHERE CTE.Rank = N
    LIMIT 1
  );
END;
$$ LANGUAGE plpgsql;
