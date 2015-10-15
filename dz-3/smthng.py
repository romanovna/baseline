__author__ = 'romandeles'
"""Создать базу данных:
юзер вводит одну из комманд: input, output
input : вводит марку авто и его мощность.
необходимо проверить, что марка состоит из букв, а мощность только из цирф

output : выводятся все алфавиту со стандартной сортирвовкой.

реализовать поиск в бд

по мощности
по названию
"""

import pickle

commands = ['input','output']
listen_command = ''
auto_key ={}

while True:
    listen_command = input('Type the command of the list {0}'.format(commands))
    if listen_command in commands:
        break
    else:
        print('Please type {0} or {1}'.format(commands[0],commands[1]))

input_data = ''

if listen_command == commands[0]:
    print('Ok, you have chosen {0} some data'.format(commands[0]))
    input_data = input('Type something ').split()
    print(input_data)
    auto_key = dict(input_data)

print(auto_key)
