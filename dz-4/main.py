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

# Добавлены исключения


import os,sys
sys.path.append(sys._home)
import db,debug

auto_key = {}

if __name__ == '__main__':
    try:
        print('Hello! Initializing data...')
        debug.load_data()
    except (FileNotFoundError,KeyboardInterrupt):
        print('System crashed! Please contact your system administrator (づ｡◕‿‿◕｡)づ ')
    else:
        db.entry_point()