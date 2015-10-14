__author__ = 'roman.deles'
# coding: utf-8
"""
1) Создаем список, заполняем (имя + фамилия)

2)выводим имя одного студента на экран:
получаем индекс через функцию input().
Выводим на экран студента по этому индексу

3)Выводим на экран имена нескольких студентов:
получаем через input начало и конец среза
выводим на экран студентов из такого среза

4) находим кол-во студентов , в именах котоых есть буква 'Р'

5)находим группы студентов с одинаковыми именами и создаем списки этих групп
"""

from collections import Counter

#1
list_of_students = []

while len(list_of_students) < 5:
    name = input('type name: ')
    if name.isdigit() is False:
        list_of_students.append(name)
        print('Ok, sounds good')
    else:
        print('You should type only letters')

print('{name} has {elem} elements'.format(name=list_of_students, elem=(len(list_of_students))))

#2
"""
single_choice = ''

while True:
    single_choice = input('Type index of the {0}: '.format(list_of_students))
    if single_choice.isdigit() is True:
        print('ok, Passed')
        break
    else:
        print('Type the numbers only!')
print('You have selected {0}'.format(list_of_students[int(single_choice)]))

#3

slice_choice_1 = ''
slice_choice_2 = ''

while True:
    slice_choice_1 = input("Type start's index of the {0} ".format(list_of_students))
    if slice_choice_1.isdigit() is True:
        print('Ok, Passed ! ')
        break
    else:
        print('You have to type numbers only')

while True:
    slice_choice_2 = input("Type end's index of the {0} ".format(list_of_students))
    if slice_choice_2.isdigit() is True:
        print('Ok, passed')
        break
    else:
        print('You have to type numbers only')

# print('You have selected %s' % list_of_students[int(slice_choice_1):int(slice_choice_2)])
for item in list_of_students[int(slice_choice_1):int(slice_choice_2)]:
    print('%s' % item, end=', ')

#4
find_letter = ''
letters_list = []
while True:
    find_letter = input("Type the letter you wanna select from {0} ".format(list_of_students))
    if find_letter.isdigit() is True:
        print('Nope, Try again ! ')
    else:
        print('Ok, sounds good')
        break

print('You have selected the {0}'.format(find_letter), end='\n')

for item in list_of_students:
    if item.find(find_letter) != -1:
        letters_list.append(item)

print('Result: {0}'.format(letters_list), end='\n')
"""
#5
names = []
uniq = []

for i in list_of_students:
    name = i.split()[0]
    if name not in names:
        names.append(name)
        uniq.append([i])
    else:
        uniq[ names.index(name) ].append(i)

print(uniq)