-- Write your PostgreSQL query statement below
WITH CUSTOMER_AMOUNT_MEAN AS (
    SELECT  
        visited_on,
        SUM(amount) AS SUM_AMOUNT
    FROM CUSTOMER
    GROUP BY visited_on
    ORDER BY visited_on ASC
)

SELECT 
    visited_on,
    SUM(SUM_AMOUNT) OVER (ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) AS amount,
    ROUND(AVG(SUM_AMOUNT) OVER (ROWS BETWEEN 6 PRECEDING AND CURRENT ROW), 2) AS average_amount
FROM CUSTOMER_AMOUNT_MEAN
OFFSET 6;
