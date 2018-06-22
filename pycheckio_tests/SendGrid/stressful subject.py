import re

STOP_WORDS = ('help', 'asap', 'urgent')


def check_normalized(text):
    return any(word in re.sub(r'(\w)\1+', r'\1', text) for word in STOP_WORDS)


def check_explicit(text):
    if '!!!' in text[-3:] or text.isupper():
        return True


def is_stressful(subj):
    clean_subj = re.sub('[!-.]', '', subj).lower()

    if check_explicit(subj) or check_normalized(clean_subj):
        return True

    return False


if __name__ == '__main__':
    assert is_stressful("Hi") == False, "First"
    assert is_stressful("I neeed HELP") == True, "Second"
    print('Done! Go Check it!')
