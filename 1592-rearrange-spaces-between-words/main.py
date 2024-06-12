class Solution:
    def reorderSpaces(self, text: str) -> str:
        
        if not text or " " not in text: 
            return text

        count_spaces = text.count(" ")
        without_spaces_list = text.split()

        if len(without_spaces_list) == 1:
            return without_spaces_list[0] + " " * count_spaces
        
        div_spaces = count_spaces // (len(without_spaces_list) - 1)
        remaining_spaces = count_spaces % (len(without_spaces_list) - 1)
        
        for i in range(len(without_spaces_list) - 1):
            without_spaces_list[i] += " " * div_spaces

        without_spaces_list[-1] += " " * remaining_spaces
    
        return "".join(without_spaces_list)
