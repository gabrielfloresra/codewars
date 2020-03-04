#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* longest(char* s1, char* s2) {
	char leters[26]="00000000000000000000000000";
	int posLeter;
	int pos = 0;
	char car = s1[pos];
	while(car != '\0'){
		
		posLeter = car - 97;
		leters[posLeter] = posLeter + 97;
		pos++;
		car = s1[pos];
	}
	pos = 0;
	car = s2[pos];
	while(car != '\0'){
		posLeter = car - 97;
		leters[posLeter] = posLeter + 97;
		pos++;
		car = s2[pos];
	}

	pos = 0;

	for(int i = 0; i < 26; i++){
		if(leters[i] != '0'){
			leters[pos] = leters[i];
			pos++;
		}
	}

	leters[pos] = '\0';
	return leters;
}

void dotest(char* s1, char* s2, char *expr)
{
    char *act = longest(s1, s2);
    if(strcmp(act, expr) != 0)
        printf("\n%s\n%s\n\n", expr, act);
}

int main(){
    dotest("aretheyhere", "yestheyarehere", "aehrsty");
	dotest("loopingisfunbutdangerous", "lessdangerousthancoding", "abcdefghilnoprstu");
    dotest("inmanylanguages", "theresapairoffunctions", "acefghilmnoprstuy");
    dotest("lordsofthefallen", "gamekult", "adefghklmnorstu");
    dotest("codewars", "codewars", "acdeorsw");
	return 0;
}