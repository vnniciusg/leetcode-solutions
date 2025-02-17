class Solution:
    def numTilePossibilities(self, tiles: str) -> int:
    
        def backtrack(path, counter) -> None:
            if path:
                result.add(path)
            
            for tile in counter:
                if counter[tile] > 0:
                    counter[tile] -= 1
                    backtrack(path + tile, counter)
                    counter[tile] += 1

        
        result = set()
        counter = {}

        for tile in tiles:
            if tile in counter:
                counter[tile] += 1
            else:
                counter[tile] = 1

        backtrack("", counter)
        return len(result)
            
