data = open('day-3-input.txt', 'r').read()
def solosanta():
    coords = [0,0]
    coord_list = ["0,0"]
    for i in data:
        if i == ">":
            coords[0] += 1
        elif i == '<':
            coords[0] -= 1
        elif i == "^":
            coords[1] += 1
        elif i == "v":
            coords[1] -= 1
        coord_list.append(str(coords[0])+","+str(coords[1]))
    return len(set(coord_list))

def robosanta():
    santa_coords = [0,0]
    robo_coords = [0,0]
    coord_list = ["0,0"]
    for i in range(0, len(data), 2):
        j = data[i]
        if j == ">":
            santa_coords[0] += 1
        elif j == "<":
            santa_coords[0] -= 1
        elif j == "^":
            santa_coords[1] += 1
        elif j == "v":
            santa_coords[1] -= 1
        coord_list.append(str(santa_coords[0])+","+str(santa_coords[1]))
        
    for x in range(1, len(data), 2):
        y = data[x]
        if y == ">":
            robo_coords[0] += 1
        elif y == "<":
            robo_coords[0] -= 1
        elif y == "^":
            robo_coords[1] += 1
        elif y == "v":
            robo_coords[1] -= 1
        coord_list.append(str(robo_coords[0])+","+str(robo_coords[1]))
    return len(set(coord_list))
    

print(solosanta())
print(robosanta())