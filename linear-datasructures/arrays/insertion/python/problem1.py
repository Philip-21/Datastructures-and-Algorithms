
def insert_item(A,N, item,pos):
      I = N
      
      while I >= pos:
           
           #Move Ith element downward.
           #This step shifts the value at index I one position to the right, 
           # effectively moving the Ith element downward
           A[I+1]=A[I]
           #In the given algorithm, the counter is a variable I used to keep track of the current position while traversing the array.
           # The counter starts with the last index of the array N, which is N-1, and then moves down the array one index at a time until it reaches the insertion position Pos.
           # The counter is decremented by one in each iteration of the loop until the insertion position is reached.
           I = I-1 #decrease the conter
      A[pos]=item #insert the element    
      N = N+1 #reset n
      return A,N     
      
      
           
           