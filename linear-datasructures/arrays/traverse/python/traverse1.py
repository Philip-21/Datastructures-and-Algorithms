#traversing a data structure is to access or modify its elements, or to search 
#for a specific element or set of elements

arr = [12,34,56,77,88]

for i in range(len(arr)):
      print(arr[0:2])
      print(arr[i])
      
for a in range(len(arr)):
      u = arr.append(99) 
      print(u)     
      
new = [23,45,67,88,12]
new[2]= 99
print(new)  
# it will print out [23,45,99,88,12] but it will store in a temporary instance
# but wont be permanently modified , unless the initial array will be readjusted again
    