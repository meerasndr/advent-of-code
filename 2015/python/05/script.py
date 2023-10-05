import re
def nice_count1(lines):
    nice = 0
    vowel_regex = re.compile(r'[aeiou]')
    doubles_regex = re.compile(r'(.)\1')
    forbid_regex = re.compile(r'(ab)|(cd)|(pq)|(xy)')
    for string in lines:
        if forbid_regex.search(string) == None and len(vowel_regex.findall(string)) >= 3 and doubles_regex.search(string):
            nice += 1
            #print(string)
    return nice

def nice_count2(lines):
    nice = 0
    for string in lines:
        pair = re.search(r'(..).*\1', string)
        between = re.search(r'(.).\1', string)
        if pair and between:
            nice += 1
    return nice

with open("input.txt") as file:
    lines = [line.rstrip().lower() for line in file]
    
print(nice_count1(lines))
print(nice_count2(lines))
