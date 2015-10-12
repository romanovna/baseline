__author__ = 'romandeles'

"""создать файл со списком чисел, прочитать этот файл,
первые два числа надо сложить, третье и четвертое перемножить и взять корень этого числа,
взять максимальное из этих чисел и от 5ого числа взять косинус"""

import random, math

list = [i for i in random.sample(range(1, 99), 5)]
print('5 randomly generated numbers : ', str(list)[1:-1])

numbers = open("numbers.txt", mode="w+")

for item in list:
    numbers.write("%s\n" % item)

summa = list[0] + list[1]
print('Сумма первых двух чисел = ', summa,file=numbers)
multiplication = math.sqrt(list[2] * list[3])
print('Произведение и корень из 3его и 4его чисел =', multiplication,file=numbers)
print('Максимальное из первых двух результатов', max(summa,multiplication),file=numbers)
cosinus_5 = math.cos(list[4])
print('Косинус пятого числа =', cosinus_5,file=numbers)

numbers.close()