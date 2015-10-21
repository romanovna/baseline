__author__ = 'romandeles'
# coding: utf-8

import pickle, os, sys, debug
from collections import OrderedDict

auto_key = {}


def entry_point():
    users_input = input('Type the command of the list {0}: \n'.format(debug.load_commands()))
    if users_input == debug.load_commands()[0]:
        print('Ok, you have chosen to {0} some data'.format(debug.load_commands()[0]))
        add_data()
    elif users_input == debug.load_commands()[1]:
        print('Ok, you have chosen to {0} some data'.format(debug.load_commands()[1]))
        debug.load_data()
    elif users_input == debug.load_commands()[2]:
        print('Ok, you have chosen to {0} from the program'.format(debug.load_commands()[2]))
        sys.exit()
    elif users_input == debug.load_commands()[3]:
        print('Ok, you have chosen to {0} from the program'.format(debug.load_commands()[3]))
        search_smthn = input('Type brand or HP of a car you wanna search in the Database: ')
        search_data(search_smthn)
    elif users_input == debug.load_commands()[4]:
        print('Ok, you have chosen to {0} from the program'.format(debug.load_commands()[4]))
        edit_smthn = input('type the item you wanna to edit: ')
        edit_data(edit_smthn)
    elif users_input == debug.load_commands()[5]:
        print('Ok, you have chosen to {0} from the program'.format(debug.load_commands()[5]))
        delete_smthn = input('type the item you wanna to delete: ')
        delete_data(delete_smthn)
    else:
        print('U have to choose one of the command from the list')
        entry_point()


def add_data():
    global auto_key
    cars = open('CARS.p', 'wb')
    input_data = input('Type CAR and HP , for example VAZ:300 \n')
    auto_key[input_data.split(':')[0]] = input_data.split(':')[1]
    for key, value in auto_key.items():
        if key.isalnum() and value.isdigit():
            pickle.dump(auto_key, cars)
        else:
            debug.typo_error()
            add_data()
    cars.close()
    debug.repeat()


def search_data(input_data):
    for key, value in auto_key.items():
        if key == input_data or value == input_data:
            print(key + ':' + value)
    entry_point()


def edit_data(inputdata):
    for key, value in auto_key.items():
        if key in inputdata:
            print('ok, founded!')
            new_value = input('Type new value for the car :')
            auto_key[key] = new_value
    cars = open('CARS.p', 'wb')
    pickle.dump(auto_key, cars)
    cars.close()
    entry_point()


def delete_data(inputdata):
    flag = False
    for key, value in auto_key.items():
        if key in inputdata:
            print('ok, the element named %s will be deleted' % inputdata)
            flag = True
    if flag:
        del auto_key[inputdata]
        flag = False
    cars = open('CARS.p', 'wb')
    pickle.dump(auto_key, cars)
    cars.close()
    entry_point()
