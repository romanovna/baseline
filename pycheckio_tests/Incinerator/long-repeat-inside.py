import re


def repeat_inside(line):
    res = [i[0] + i[1] for i in re.findall(r'(?=(.+?)(\1+))', line)]
    return max(res, key=len) if res else ''


if __name__ == '__main__':
    # These "asserts" using only for self-checking and not necessary for auto-testing
    assert repeat_inside('aaaaa') == 'aaaaa', "First"
    assert repeat_inside('aabbff') == 'aa', "Second"
    assert repeat_inside('aababcc') == 'abab', "Third"
    assert repeat_inside('abc') == '', "Forth"
    assert repeat_inside('abcabcabab') == 'abcabc', "Fifth"
    print('"Run" is good. How is "Check"?')
