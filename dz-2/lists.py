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

print(list_of_students)

single_choice = input('Type index of the {0}: '.format(list_of_students))

print('You have selected {0}'.format(list_of_students[int(single_choice)]))
