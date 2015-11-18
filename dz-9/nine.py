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
4. Далее запускается метод класса Ферма "Сводная инфа" , который расскажет нам об изменениях на ферме. """

import random
from abc import ABCMeta, abstractmethod


# изменен базовый класс на абстрактный, с абстрактным методом bon_appetit
class base(metaclass=ABCMeta):
    def __init__(self, prod=str, K_prod=int, prod_items=str, voice_per_day=0, speed=0, moves_per_day=0, food=float):
        ## counters ============================
        self.things_current = 0
        self.voices_current = 0
        self.dist_current = 0
        self.food_current = 0

        self.things_at_all = 0
        self.voices_at_all = 0
        self.dist_at_all = 0
        self.food_at_all = 0
        ##=======================================


        ## some attributes
        self.prod = prod
        self.k_prod = K_prod
        self.prod_items = prod_items
        self.voices_per_day = voice_per_day
        self.speed = speed
        self.moved_per_day = moves_per_day
        self.food = food

        self.voice_k = random.randint(1, 5)
        self.move_k = random.randint(1, 5)
        self.things_k = random.randint(1, 5)
        self.summary_in_month = 0

    def move(self, days=1):
        self.dist_current = (self.speed * self.moved_per_day * days * self.move_k)
        self.dist_at_all += self.dist_current

    def voice(self, days=1):
        self.voices_current = (self.voices_per_day * days * self.voice_k)
        self.voices_at_all += self.voices_current

    def items(self, month):
        self.things_current = (self.things_k * self.k_prod * month)
        self.things_at_all += self.things_current

    ## counting...
    def deal_with_it(self, month=1):
        self.voice(month * 31)
        self.items(month)
        self.move(month * 31)
        self.bon_appetit(month * 31)
        self.summary_in_month += month

    @abstractmethod
    def bon_appetit(self, days):
        pass


class CrazyChicken(base):
    def __init__(self, prod='Eggs', K_prod=8, prod_items='items', voice_per_day=25, speed=2, moves_per_day=2, food=0.3):
        super().__init__(prod, K_prod, prod_items, voice_per_day, speed, moves_per_day, food)

    def bon_appetit(self, days=1):
        self.food_current = (self.food * days)
        self.food_at_all += self.food_current


class CrazyCow(base):
    def __init__(self, prod='Vodka', K_prod=5, prod_items='liters', voice_per_day=50, speed=15, moves_per_day=4,
                 food=1.2):
        super().__init__(prod, K_prod, prod_items, voice_per_day, speed, moves_per_day, food)

    def bon_appetit(self, days=1):
        self.food_current = (self.food * days)
        self.food_at_all += self.food_current


class CrazyDog(base):
    def __init__(self, prod='Anger', K_prod=15, prod_items='tonnes', voice_per_day=30, speed=30, moves_per_day=8,
                 food=0.8):
        super().__init__(prod, K_prod, prod_items, voice_per_day, speed, moves_per_day, food)

    def bon_appetit(self, days=1):
        self.food_current = (self.food * days)
        self.food_at_all += self.food_current


class the_Farm:
    def __init__(self):

        self.cows = []
        self.dogs = []
        self.chickens = []

        for i in range(0, 3):
            self.chickens.append(CrazyChicken())
        for i in range(0, 2):
            self.cows.append(CrazyCow())
        for i in range(0, 1):
            self.dogs.append(CrazyDog())

        self.animals = {'CRAZY_DOG': self.dogs, 'CRAZY_COW': self.cows, 'CRAZY_CHICKEN': self.chickens}

    def past_month(self, month=1):
        for animal in self.animals.keys():
            for x in self.animals[animal]:
                x.deal_with_it(month)
                x.bon_appetit(month)

    def status(self):
        with open('farm9.txt', mode='a+', encoding='utf-8') as f:
            print('============================================START=================================================',
                  file=f)
            for animal in self.animals.keys():
                if self.animals[animal]:
                    total_count = 0
                    for i, x in enumerate(self.animals[animal]):
                        print('%s' % animal, i + 1, file=f)
                        print('produced', x.prod, 'in an amount of ', x.things_at_all, x.prod_items, file=f)
                        print('runned out', x.dist_at_all, 'KM', file=f)
                        print('did a sound', x.voices_at_all, 'times', file=f)
                        print('eaten a food', x.food_at_all, 'kg', file=f)
                        total_count += x.things_at_all

                    print(animal, 'has produced at all', total_count, x.prod_items, 'products', x.prod, 'for a',
                          x.summary_in_month,
                          'months on the farm', file=f)
            print('===========================================END====================================================',
                  file=f)


if __name__ == '__main__':
    farm = the_Farm()
    farm.past_month(1)
    farm.status()
