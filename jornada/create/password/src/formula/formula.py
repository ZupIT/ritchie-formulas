# Password Generator Project
import random

letter = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y',
           'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
number = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
symbol = ['!', '#', '$', '%', '&', '(', ')', '*', '+']



def Run(number_letters, number_symbols, number_numbers):
    password_easy = ""
    
    letters = int(number_letters)
    symbols = int(number_symbols)
    numbers = int(number_numbers)
    
    while letters > 0:
        password_easy += random.choice(letter)
        letters -= 1
    while symbols > 0:
        password_easy += random.choice(symbol)
        symbols -= 1
    while numbers > 0:
        password_easy += random.choice(number)
        numbers -= 1


    password_hard = ''.join(random.sample(password_easy, len(password_easy)))
    print(f"Your strongest password is: {password_hard}")

