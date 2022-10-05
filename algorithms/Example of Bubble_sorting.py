a=[22,44,1,5,3,77,5,40,1325,76]
b=len(a)
for j in range(0,b):
    for i in range(0,b):
        if i==b-1:
            pass
        else:  
            if a[i]>a[i+1]:
                a[i],a[i+1]=a[i+1],a[i]
            else:
                pass
print(a)    
