class Solution:
    def eventualSafeNodes(self, graph: List[List[int]]) -> List[int]:
        
        def dfs(node, adj, visit, in_stack):

            if in_stack[node]: return True
            if visit[node]: return False

            visit[node], in_stack[node] = True, True

            for neighbor in adj[node]:
                if dfs(neighbor, adj, visit, in_stack): return True

            
            in_stack[node] = False
            return False

        
        n = len(graph)
        visit = [False] * n
        in_stack = [False] * n

        for i in range(n):
            dfs(i, graph, visit, in_stack)

        safe_nodes = []
        for i in range(n):
            if not in_stack[i]: safe_nodes.append(i)

        return safe_nodes
