class Solution:
    def findAllRecipes(self, recipes: List[str], ingredients: List[List[str]], supplies: List[str]) -> List[str]:
        
        disponible = set(supplies)
        needs = {recipe: set(ingredient) for recipe, ingredient in zip(recipes, ingredients)}

        updated = True
        while updated:
            updated = False

            for recipe in recipes:
                if recipe in disponible: continue
            
                if needs[recipe].issubset(disponible):
                    disponible.add(recipe)
                    updated = True
        
        return [recipe for recipe in recipes if recipe in disponible]

    
