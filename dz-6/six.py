__author__ = 'romandeles'
# coding: utf-8

import os

def find(rash='.txt'):
    for filename in os.listdir('.'):
        if filename.endswith(rash):
            for i,line in enumerate(open(filename)):
                yield filename,i,line


def grep(gen,substr):
    for name,i,s in gen:

        yield name,i,s



for name,i,s in grep(find('.txt'),'def'):
    print(name,i,s)