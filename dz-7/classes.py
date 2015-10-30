__author__ = 'romandeles'
# coding: utf-8
"""
1. создаём 3 отдельных класса (танк, машина,телега)
2. у каждого класса метод status, выводящий состояине объекта на данный момент
3. Создаём и собираем сколько-то новых объектов этих классов в список cars
4. Делаем несколько действий с этими объектами ( например, назначили машине audi скорость 90, у танка т-34 сняли шасси)
5. в конце программы выводим состояния всех объектов из сars
"""


class Tank:
    def __init__(self, shassi=int, model=str, speed=int, gusenici=bool):
        self.shassi = shassi
        self.gusenici = gusenici
        self.model = model
        self.speed = speed

    def status(self):
        print("Shassi = {0}, gusenici is {1}, speed = {2} Km/H, model = {3}".format(self.shassi, self.gusenici, self.speed,
                                                                               self.model))
        if self.speed > 0 and self.gusenici is True:
            print('Health check is ok')
        else:
            print('Health check is failed')


class Car:
    def __init__(self, speed=int, model=str, kolesa=4):
        self.speed = speed
        self.kolesa = kolesa
        self.model = model

    def status(self):
        print('Speed = {0} Km/H, kolesa = {1}, model is {2}'.format(self.speed, self.kolesa, self.model))
        if self.speed > 0 and self.kolesa == 4:
            print('Health check is ok')
        else:
            print('Health check is failed')


class Telega:
    def __init__(self, kolesa=int, speed=int):
        self.kolesa = kolesa
        self.speed = speed

    def status(self):
        print('Kolesa = {0} Km/H, speed = {1}'.format(self.kolesa, self.speed))
        if self.kolesa == 2 and self.speed > 0:
            print('Health check is ok')
        else:
            print('Health check is failed')


Patton = Tank(10, 'M46 Patton 2', 50, True)
BMW = Car(120, 'E46')
Telega_1 = Telega(2, 5)

cars = [Patton, BMW, Telega_1]

print('Default values is: ')
for i in cars:
    i.status()


BMW.kolesa = 3
Telega_1.speed = 0
Patton.gusenici = False

print('Some changes was happened')
for i in cars:
    i.status()