def sum(n):
    if n==1: return 4 % 1009
    elif n==2: return (4+43) % 1009
    elif n==3: return (4+43+49) % 1009
    else: return 2*((3**n)-1) % 1009

    print(sum(1))