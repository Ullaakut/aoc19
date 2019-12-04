def create_list_of_positions(wire):
    wire = wire.split(",")
    positions = dict()
    x = 0
    y = 0
    for move in wire:
        direction = move[0]
        steps = int(move[1:])
        if direction == 'R':
            for i in range(0, steps):
                x += 1
                positions[(x,y)] = True
        elif direction == 'L':
            for i in range(0, steps):
                x -= 1
                positions[(x,y)] = True
        elif direction == 'U':
            for i in range(0, steps):
                y -= 1
                positions[(x,y)] = True
        else:
            for i in range(0, steps):
                y += 1
                positions[(x,y)] = True
    return positions


with open(f'input.txt') as f:
    content = f.readlines()
    wires = [line.strip() for line in content]

    first_list_positions = set(create_list_of_positions(wires[0]).keys())
    second_list_positions = set(create_list_of_positions(wires[1]).keys())
    intersections_distance = []

    intersections = first_list_positions.intersection(second_list_positions)
    print(intersections)
    for intersection in intersections:
        # print(intersection)
        intersections_distance.append(abs(intersection[0]) + abs(intersection[1]))

    print(min(intersections_distance))


