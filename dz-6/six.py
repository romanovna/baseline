# coding: utf-8
import os

def find(ext='.py'):
    for filename in os.listdir('.'):
        if filename.endswith(ext):
            for i,line in enumerate(open(filename),1):
                if 'def' in line:
                    yield filename,i,line


def grep(gen,substr):
    for name,i,s in gen:
        yield name,i,s



for name,i,s in grep(gen=find('.py'),substr='def'):
    print(name,i,s)