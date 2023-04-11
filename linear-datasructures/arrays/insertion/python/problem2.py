
def insert(I,N,Pos,ITEM):
      I = N-1
      while I>=Pos:
            A[I+1]=A[I]
            I =I-1
      
      A[Pos]=ITEM
      N=N+1      