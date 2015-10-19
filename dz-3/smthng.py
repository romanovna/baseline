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

import pickle, sys
from collections import OrderedDict

auto_key = {}


def load_commands():
    commands = ['input', 'output', 'exit','search']
    return commands


def repeat():
    type_repeat = input('Do you wanna add some more? Y or N \n')
    if type_repeat in ['y', 'Y', 'YES', 'yes']:
        add_data()
    else:
        entry_point()


def entry_point():
    users_input = input('Type the command of the list {0}: \n'.format(load_commands()))
    if users_input == load_commands()[0]:
        print('Ok, you have chosen to {0} some data'.format(load_commands()[0]))
        add_data()
    elif users_input == load_commands()[1]:
        print('Ok, you have chosen to {0} some data'.format(load_commands()[1]))
        load_data()
    elif users_input == load_commands()[2]:
        print('Ok, you have chosen to {0} from the program'.format(load_commands()[2]))
        sys.exit()
    elif users_input == load_commands()[3]:
        print('Ok, you have chosen to {0} from the program'.format(load_commands()[3]))
        search_smthn = input('Type some data ')
        search_data(search_smthn)
    else:
        print('U have to choose one of the command from the list')
        entry_point()


def add_data():
    global auto_key
    cars = open('CARS.db','wb')
    input_data = input('Type CAR and HP , for example VAZ:300 \n')
    #auto_key = dict(input_data.split(':') for s in input_data)
    auto_key[input_data.split(':')[0]] = input_data.split(':')[1]
    for key, value in auto_key.items():
        if key.isalnum() and value.isdigit():
            print(auto_key)
            pickle.dump(auto_key,cars)
        else:
            print('Please type correct items')
            add_data()
    cars.close()
    repeat()


def load_data():
    global auto_key
    with open('CARS.db','rb') as f:
        cars = pickle.load(f)
        auto_key = cars
        cars_sorted = OrderedDict(sorted(cars.items()))
        for key,value in cars_sorted.items():
            print(key +':'+value)
    entry_point()


def search_data(input_data):
    for key,value in auto_key.items():
        if key == input_data or value == input_data:
            print(key + ':' + value)
    entry_point()

if __name__ == '__main__':
    entry_point()




"""
while True:
    listen_command = input('Type the command of the list {0}: '.format(commands))
    if listen_command == commands[0]:
        print('Ok, you have chosen to {0} some data: '.format(commands[0]))
        input_data = input('Type CAR and HP ')
        auto_key = dict(input_data.split( ) for s in input_data)
        for key, value in auto_key.items():
            if key.isalpha() and value.isdigit():
                with open('CARS.db','wb+') as f:
                    pickle.dump(auto_key,f)
                    f.close()
            else:
                print('Please type correct items')

    if listen_command == commands[1]:
        cars = pickle.load(open('CARS.db','rb'))
        print(cars)
    if listen_command == commands[2]:
        sys.exit()
"""


