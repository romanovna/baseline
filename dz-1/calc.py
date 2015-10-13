__author__ = 'romandeles'
#coding: utf-8

"""Дефолтный калькулятор
замена запятой на точку - присутствует
"""

math_ops = ['+', '-', '*', '/', '%']

users_input = input("type number #1: ")
users_input = users_input.replace(",", ".")
for i in users_input:
    if i in math_ops:
        num_index = users_input.index(i)
number1 = users_input[:num_index]
sign = users_input[num_index:num_index + 1]
number2 = users_input[num_index + 1:]

if sign is "+":
    result = float(number1) + float(number2)
if sign is "-":
    result = float(number1) - float(number2)
if sign is "*":
    result = float(number1) * float(number2)
if sign is '/':
    result = float(number1) / float(number2)
if sign is '%':
    result = float(number1) % float(number2)

print(result)
