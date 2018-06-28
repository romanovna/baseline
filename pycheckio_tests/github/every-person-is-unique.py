import datetime


class Person:
    def __init__(self, first_name, last_name, birth_date, job, working_years, salary, country, city, gender='unknown'):
        self.fname = first_name
        self.lname = last_name
        self.bdate = birth_date
        self.job = job
        self.wyears = working_years
        self.salary = salary
        self.country = country
        self.city = city
        self.gender = gender

    def name(self):
        return '{} {}'.format(self.fname, self.lname)

    def age(self):
        cur_date = datetime.datetime.strptime('01.01.2018', '%d.%m.%Y')
        bbdate = datetime.datetime.strptime(self.bdate, '%d.%m.%Y')
        return cur_date.year - bbdate.year - ((cur_date.month, cur_date.day) < (bbdate.month, bbdate.day))

    def work(self):
        genders = {'male': 'He is', 'female': 'She is', 'unknown': 'Is'}
        return '{} a {}'.format(genders[self.gender], self.job)

    def money(self):
        return f'{12 * self.wyears * self.salary:,}'.replace(',', ' ')

    def home(self):
        return 'Lives in {}, {}'.format(self.city, self.country)


if __name__ == '__main__':
    # These "asserts" using only for self-checking and not necessary for auto-testing

    p1 = Person("John", "Smith", "19.09.1979", "welder", 15, 3600, "Canada", "Vancouver", "male")
    p2 = Person("Hanna Rose", "May", "05.12.1995", "designer", 2.2, 2150, "Austria", "Vienna")
    assert p1.name() == "John Smith", "Name"
    assert p1.age() == 38, "Age"
    assert p2.work() == "Is a designer", "Job"
    assert p1.money() == "648 000", "Money"
    assert p2.home() == "Lives in Vienna, Austria", "Home"
    print("Coding complete? Let's try tests!")
