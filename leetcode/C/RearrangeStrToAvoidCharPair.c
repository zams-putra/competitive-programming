#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char* rearrangeString(char* s, char x, char y) {
    int len = strlen(s);
    char* result = (char*)malloc(len + 1);
    int index = 0;
    for(int i = 0; i < len; i++){
        if(s[i] != x){
            result[index++] = s[i];
        }
    }
    for(int i = 0; i < len; i++){
        if(s[i] == x){
            result[index++] = s[i];
        }
    }
    result[index] = '\0';
    return result;
}