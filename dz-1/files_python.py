__author__ = 'romandeles'
#+79819711529 - Ирина
#a9711529@yandex.ru

"""
1)найти в папке все файлы в которых содержится слово python
и вывести все файлы с совпадениями
2) запиисать в файл result.txt список найденных файлов и число
найденных слов python
"""



import os,sys,glob

home_directory = sys._home
list_of_files = os.listdir(home_directory)
results = open("results","w+")
words_count = 0

for i in list_of_files and glob.glob("list*.txt"):
    data = open(i,"r").read()
    words_count += data.count("python")
    print(data.count("python"),i,sep=" matches in ",file=results)

print("Общее кол-во найденных слов = ",words_count, file=results)

results.close()
