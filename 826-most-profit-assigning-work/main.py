class Solution:
    def maxProfitAssignment(self, difficulty: List[int], profit: List[int], worker: List[int]) -> int:
        
        jobs = list(zip(difficulty, profit))
        
        jobs.sort()
        worker.sort()

        max_profit, best_profit = 0, 0
        job_index = 0
        n = len(jobs)
        for ability in worker:
            while job_index < n and ability >= jobs[job_index][0]:
                best_profit = max(best_profit, jobs[job_index][1])
                job_index += 1  

            max_profit += best_profit
        
        return max_profit
