class Solution:
    def maxSatisfied(self, customers: List[int], grumpy: List[int], minutes: int) -> int:
        
        n = len(customers)

        satisfied_customers = sum([costumer for idx, costumer in enumerate(customers) if grumpy[idx] == 0])
        additional_satisfied = sum([customers[idx] for idx in range(minutes) if grumpy[idx] == 1])

        max_additional_satisfied = additional_satisfied

        for idx in range(minutes, n):
            if grumpy[idx] == 1:
                additional_satisfied += customers[idx]
            if grumpy[idx - minutes] == 1:
                additional_satisfied -= customers[idx - minutes]

            max_additional_satisfied = max(max_additional_satisfied, additional_satisfied)

        return satisfied_customers + max_additional_satisfied

