__author__ = 'romandeles'
# coding: utf-8

"""
1)
из 3его ДЗ выделить следующие функции:
- Ввод команды - отдельная функция
- Сообщение в случае ошибки ввода команды - отдельная функция
- Ввести и вывести -  2 отдельные функции
- поиски по условию - 3 отдельные функции соответственно
- сохранение в pickle и загрузка из pickle - 2 отдельные функции

2) функции ввести и вывести добавляем в словарь след. образом:
funcs = { 'input':input_func, 'output':output_func,} и меняем if-else на поиск в этом словаре и запуск функции
по ключу словаря

3)добавить редактирование - для изменения уже введеных данных, а также "удалить".
"""

import os,sys,pickle
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
        search_smthn = input('Type brand or HP of a car you wanna search in the Database: ')
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

