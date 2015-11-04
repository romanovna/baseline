__author__ = 'roman.deles'
# coding: utf-8
""" Жизнь на ферме
- Курица
- Корова
- Собака
1. У этих классов есть следующие функциональности:
- бежать
- голос
- продукт ( яйцо, молоко)      все эти классы унаследованы от базового "животное"
2. Также нужен класс ферма
Программа инициализирует ферму с заданным числом каждого животного.
3. Далее запускается метод класса ферма "прошел_месяц"
Там циклом проходим по всем животным, запуская их собственный метод "прошел месяц"
(какое животное сколько раз делает продукт, как успешно, где использовать random,
какие случайные факторы внести в жизнь фермы, решайте сами)
4. Далее запускается метод класса Ферма "Сводная инфа" , который расскажет нам об изменениях на ферме."""

import time,random


class base:
    def __init__(self, prod=str, K_prod=int, prod_items=str, voice_per_day=0, speed=0, moved_per_day=0):
## counters ============================
        self.things_in_last_month = 0
        self.voices_in_last_month = 0
        self.dist_in_last_month = 0

        self.things_at_all = 0
        self.voices_at_all = 0
        self.dist_at_all = 0
##=======================================


## some attributes
        self.prod = prod
        self.k_prod = K_prod
        self.prod_items = prod_items
        self.voices_per_day = voice_per_day
        self.speed = speed
        self.moved_per_day = moved_per_day

        self.voice_k= random.randint(1,5)
        self.move_k = random.randint(1,5)
        self.things_k = random.randint(1,5)
        self.summary_in_month = 0

    def move(self,days=1):
        self.dist_in_last_month=(self.speed * self.moved_per_day * days * self.moved_per_day * self.move_k)
        self.dist_at_all += self.dist_in_last_month

    def voice(self,days=1):
        self.voices_in_last_month = (self.voices_per_day * days * self.voice_k)
        self.voices_at_all += self.voices_in_last_month

    def items(self,month):
        self.things_in_last_month = (self.things_k * self.k_prod * month)
        self.things_at_all += self.things_in_last_month


    ## counting...
    def deal_with_it(self, month=1):
        self.move(month*30)
        self.voice(month*30)
        self.items(month)
        self.summary_in_month += month



class Chicken_from_the_Hell(base):
    def __init__(self, prod='Eggs',K_prod=30,prod_items='items',speed=2,moved_per_day=5):
        super().__init__(prod,K_prod,prod_items,speed,moved_per_day)
        self.voices_per_day = 50




class Cow_from_the_Hell(base):
    def __init__(self,prod='Vodka',K_prod=5,prod_items='liters',speed=15,moved_per_day=80):
        super().__init__(prod,K_prod,prod_items,speed,moved_per_day)
        self.voices_per_day = 500


class crazy_Dog(base):
    def __init__(self,prod='Anger',K_prod=50,prod_items='tonnes',speed=30,moved_per_day=150):
        super().__init__(prod,K_prod,prod_items,speed,moved_per_day)
        self.voices_per_day = 900

class the_Hell:

