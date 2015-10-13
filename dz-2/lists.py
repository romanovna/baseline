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

list_of_students = []

while len(list_of_students) < 5:
    name = input('type name: ')
    if name.isalpha() is True:
        list_of_students.append(name)
        print('Ok, sounds good')
    else:
        print('You should type only letters')

print('{name} has {elem} elements'.format(name=list_of_students, elem=(len(list_of_students))))

single_choice = input('Type index of the {0}: '.format(list_of_students))

print('You have selected {0}'.format(list_of_students[int(single_choice)]))

slice_choice_1 = ''
slice_choice_2 = ''

while True:
    slice_choice_1 = input("type start's index of the {0} ".format(list_of_students))
    if slice_choice_1.isdigit() is True:
        print('ok, passed')
        break
    else:
        print('you have to type numbers only')

while True:
    slice_choice_2 = input("type end's index of the {0} ".format(list_of_students))
    if slice_choice_2.isdigit() is True:
        print('ok, passed')
        break
    else:
        print('You have to type numbers only')

for item in list_of_students[int(slice_choice_1):int(slice_choice_2)]:
    print(item, end=' ')
