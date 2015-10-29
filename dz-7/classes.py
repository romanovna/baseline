__author__ = 'romandeles'
# coding: utf-8

class Tank:
    def __init__(self):
        self.shassi = 10
        self.gusenici = True
        self.model = 'M46 Patton'
        self.speed = 50
    def status(self):
        print(self.shassi,self.gusenici,self.speed,self.model)
        if self.speed > 0 and self.gusenici is True:
            print('ITS ALIVE!')
        else:
            print('something going wrong')


class Car:
    def __init__(self):
        self.speed = 150
        self.kolesa = 3
        self.model = 'LADA Sedan - Baklazhan'
    def status(self):
        print(self.speed,self.kolesa,self.model)
        if self.speed > 0 and self.kolesa == 4:
            print('Ok, its moving')
        else:
            print('Vizivay evakuator')

class Telega:
    def __init__(self):
        self.kolesa = 2
        self.speed = 3
    def status(self):
        print(self.kolesa,self.speed)
        if self.kolesa == 2 and self.speed > 0:
            print('ok, its alive')
        else:
            print('something goes wrong')


spisok = [Car(),Telega(),Tank()]

for i in spisok:
    i.status()