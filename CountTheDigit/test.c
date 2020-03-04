#include <stdio.h>
#include <stdlib.h>

unsigned int getNumDigitos(unsigned int num){
  unsigned int digitos = 1;
  while(num/10 > 0){
    digitos++;
    num /= 10;
  }
  return digitos;
}

int nbDig(int n, int d) {
  if(n < 0 || d < 0 || d > 9) return 0;
  int numD = 0;
  int i, j;
  unsigned int numDigitos;
  unsigned int num;
  for(i = 0; i <= n; i++){
    num = i*i;
    numDigitos = getNumDigitos(num);
    char numChar[numDigitos];
    snprintf(numChar, numDigitos, "%d", num);
    
    // printf("%d*%d = %s\n",i,i,numChar);
    for(j = 0; j < numDigitos; j++){
      int b = numChar[j] - 48;
      if(d == b) numD++;
    }
    
  }
    // printf("encontrado %d digitos\n", numD);
  return numD;
}

void dotest(int n, int d, int exp)
{
    int act = nbDig(n, d);
    // if(act != exp)
        printf("Error. Expected %d but got %d\n", exp, act);
}

int main(int argc, char const *argv[])
{
  dotest(5750, 0, 4700);
  dotest(11011, 2, 9481);
  dotest(12224, 8, 7733);
  dotest(11549, 1, 11905);
  return 0;
}