class Solution:
    def checkIfPrerequisite(self, numCourses: int, prerequisites: List[List[int]], queries: List[List[int]]) -> List[bool]:
        
        adj = defaultdict(list)
        for prereq, crs in prerequisites:
            adj[crs].append(prereq)

        def dfs(crs):
            if crs not in prereq_map:
                prereq_map[crs] = set()
                for prereq in adj[crs]:
                    prereq_map[crs] |= dfs(prereq)
                prereq_map[crs].add(crs)

            return prereq_map[crs]

        prereq_map = {}
        for crs in range(numCourses):
            dfs(crs)

        res = []
        for u, v in queries:
            res.append(u in prereq_map[v])

        return res
