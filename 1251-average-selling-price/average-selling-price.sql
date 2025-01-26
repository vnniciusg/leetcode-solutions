# Write your MySQL query statement below
SELECT
    p.product_id,
    IFNULL(ROUND(SUM(units*price)/SUM(units), 2), 0) as average_price
FROM Prices p
LEFT JOIN UnitsSold us  
            ON p.product_id = us.product_id
            AND us.purchase_date BETWEEN start_date AND end_date
GROUP BY p.product_id;
