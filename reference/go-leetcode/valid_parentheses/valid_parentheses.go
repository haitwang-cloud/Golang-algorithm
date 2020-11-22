package valid_parentheses

func isValid(s string) bool {
    m := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }
    stack := make([]rune, len(s))
    top := 0
    for _, c := range s {
        switch c {
        case '(', '[', '{':
            stack[top] = m[c]
            top++
        case ')', ']', '}':
            if top > 0 && stack[top-1] == c {
                top--
            } else {
                return false
            }
        }
    }

    return top == 0
}
