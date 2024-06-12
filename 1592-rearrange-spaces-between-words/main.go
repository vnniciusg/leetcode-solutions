import "strings"

func reorderSpaces(text string) string {
    
    if len(text) == 0 || !strings.Contains(text," ") {
        return text
    }

    countSpaces := strings.Count(text, " ")
    withoutSpacesList := strings.Fields(text)

    m := len(withoutSpacesList)

    if m  == 1 {
        return withoutSpacesList[0] + strings.Repeat(" ", countSpaces)
    }

    divSpaces := countSpaces / (m - 1)
    remainingSpaces := countSpaces % (m - 1)

    for i := range withoutSpacesList[:m-1] {
        withoutSpacesList[i] += strings.Repeat(" ", divSpaces)
    }

    withoutSpacesList[m-1] += strings.Repeat(" ", remainingSpaces)

    return strings.Join(withoutSpacesList, "")

}
