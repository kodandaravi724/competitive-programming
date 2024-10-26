# Write your MySQL query statement below

select s.score, DENSE_RANK() over(order by score desc) as 'rank'
From Scores s order by 'rank';