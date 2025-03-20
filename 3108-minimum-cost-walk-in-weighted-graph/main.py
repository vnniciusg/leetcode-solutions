class UnionFind:

    def __init__(self, n: int):
        self.parent = list(range(n))
        self.size = [1] * n
    
    def find(self, x):
        if x != self.parent[x]:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        x = self.find(x)
        y = self.find(y)
        if x != y:
            if self.size[x] < self.size[y]:
                self.parent[x] = y
                self.size[y] += self.size[x]
            else:
                self.parent[y] = x
                self.size[x] += self.size[y]


class Solution:
    def minimumCost(self, n: int, edges: List[List[int]], query: List[List[int]]) -> List[int]:
        
        union_find = UnionFind(n)

        for u, v, w in edges:
            union_find.union(u, v)
        
        component_cost = {}
        for u, v, w in edges:
            root = union_find.find(u)
            if root not in component_cost:
                component_cost[root] = w
            else:
                component_cost[root] &= w

        res = []
        for src, dst in query:
            r1, r2 = union_find.find(src), union_find.find(dst)
            if r1 != r2:
                res.append(-1)
            else:
                res.append(component_cost[r1])
        return res
        
