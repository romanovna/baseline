from datetime import datetime


def time_converter(time):
    time = time.replace('a.m.', 'AM').replace('p.m.', 'PM')
    ct = datetime.strptime(time, '%I:%M %p')
    ct = ct.strftime('%H:%M')
    return ct


if __name__ == '__main__':
    print("Example:")
    print(time_converter('12:30 p.m.'))

    # These "asserts" using only for self-checking and not necessary for auto-testing
    assert time_converter('12:30 p.m.') == '12:30'
    assert time_converter('9:00 a.m.') == '09:00'
    assert time_converter('11:15 p.m.') == '23:15'
    print("Coding complete? Click 'Check' to earn cool rewards!")
