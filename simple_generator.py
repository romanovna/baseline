import argparse

import requests

BASE_URL = 'https://randomuser.me/api/'

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-c", "--count", type=int,
                        help="Count of records to be generated, max = 5000")
    parser.add_argument('-f', '--file', type=str, help='Output file', default='data.csv')

    args = parser.parse_args()

    args.file = args.file + '.csv' if not args.file.endswith('.csv') else args.file

    if args.count in range(1, 5001):
        data = requests.get(BASE_URL + f'?results={args.count}&noinfo&exc=login&format=csv')
        with open(args.file, 'wb') as f:
            f.write(data.content)
    else:
        print('Count must be in range of 1-5000')
