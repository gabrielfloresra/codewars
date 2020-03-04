#include <stdio.h>
#include <stdlib.h>
#define lcm(a, b) (a * b) / gcd(a, b)

typedef struct Laps_Pairing
{
    int bob;
    int charles;
} laps;

int gcd(int a, int b){
    return !b ? a : gcd(b, a % b);
}

laps nbr_of_laps(int bob, int charles)
{
    laps lapsFormated;

    int min = lcm(bob, charles);
    lapsFormated.bob = min/bob;
    lapsFormated.charles = min/charles;
    return lapsFormated;
}
int main(int argc, char const *argv[])
{
    nbr_of_laps(5, 3);
    nbr_of_laps(4, 6);
    return 0;
}