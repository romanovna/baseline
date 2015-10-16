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


def load_commands():
    commands = ['input', 'output', 'exit']
    return commands


def input_command():
    input1 = input('Type the command of the list {0}: '.format(load_commands()))
    if input1 == load_commands()[0]:
        print('Ok, you have chosen to {0} some data'.format(load_commands()[0]))
        _input_data()
    elif input1 == load_commands()[1]:
        print('Ok, you have chosen to {0} some data'.format(load_commands()[1]))
        load_data()
    elif input1 == load_commands()[2]:
        print('Ok, you have chosen to {0} from the program'.format(load_commands()[2]))
        sys.exit()
    else:
        print('U have to choose one of the command from the list')
        input_command()


def _input_data():
    input_data = input('Type CAR and HP ')
    auto_key = dict(input_data.split() for s in input_data)
    for key, value in auto_key.items():
        if key.isalpha() and value.isdigit():
            with open('CARS.db', 'wb') as f:
                pickle.dump(auto_key, f)
        else:
            print('Please type correct items')
            _input_data()
    repeat = input('do you wanna add some more? Y or N ')
    if repeat == 'Y':
        _input_data()
    else:
        input_command()


def load_data():
    with open('CARS.db', 'rb') as f:
        print(pickle.load(f))


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

input_command()
