__author__ = 'romandeles'
# coding: utf-8

import pickle, db, main
from collections import OrderedDict


def load_commands():
    commands = ['input', 'output', 'exit', 'search', 'edit', 'delete']
    return commands


def typo_error():
    return print('You have inserted an incorrect value')


def repeat():
    type_repeat = input('Do you wanna add some more? Y or N \n')
    if type_repeat in ['y', 'Y', 'YES', 'yes']:
        db.add_data()
    else:
        db.entry_point()


def load_data():
    with open('CARS.p', 'rb') as f:
        cars = pickle.load(f)
        main.auto_key = cars
        cars_sorted = OrderedDict(sorted(cars.items()))
        for key, value in cars_sorted.items():
            print(key + ':' + value)
    db.entry_point()
